package file

import (
	"context"
	"github.com/carlosrodriguesf/dfile/pkg/model"
	"github.com/carlosrodriguesf/dfile/pkg/repository/file"
	"github.com/carlosrodriguesf/dfile/pkg/tool/array"
	"github.com/carlosrodriguesf/dfile/pkg/tool/counter"
	"github.com/carlosrodriguesf/dfile/pkg/tool/stacktrace"
	"github.com/carlosrodriguesf/dfile/pkg/tool/sumcalc"
	"github.com/carlosrodriguesf/dfile/pkg/tool/threadmanager"
	"golang.org/x/sync/errgroup"
	"log"
	"sync"
)

type (
	Params struct {
		SumCalc        sumcalc.Calculator
		FileRepository file.Repository
	}
	Service interface {
		GenerateChecksum(ctx context.Context) error
	}
	service struct {
		sumCalc        sumcalc.Calculator
		fileRepository file.Repository
	}
)

func NewService(params Params) Service {
	return &service{
		sumCalc:        params.SumCalc,
		fileRepository: params.FileRepository,
	}
}

func (s service) GenerateChecksum(ctx context.Context) error {
	allFiles, err := s.fileRepository.ListAll(ctx)
	if err != nil {
		return stacktrace.WrapError(err)
	}

	var (
		errGroup  *errgroup.Group
		mutexSave sync.Mutex

		countTotal      = len(allFiles)
		countGenerating = counter.New(countTotal)
		countGenerated  = counter.New(countTotal)

		threadManager = threadmanager.NewManager(100)
	)
	errGroup, ctx = errgroup.WithContext(ctx)
	array.Each(ctx, allFiles, func(item model.File) {
		threadManager.Lock()
		errGroup.Go(func() error {
			defer threadManager.Release()

			generating, generatingPercent := countGenerating.Increment()
			log.Printf("[%d/%d:%.2f%%] generating checksum for '%s'", generating, countTotal, generatingPercent, item.Path)

			sum, err := s.sumCalc.Calculate(item.Path)
			if err != nil {
				item.Error = stacktrace.ExtractCause(err)
			} else {
				item.Checksum = sum
			}

			generated, generatedPercent := countGenerated.Increment()
			log.Printf("[%d/%d:%.2f%%] checksum generated for '%s': '%s'", generated, countTotal, generatedPercent, item.Path, sum)

			mutexSave.Lock()
			defer mutexSave.Unlock()
			err = s.fileRepository.Save(ctx, item)
			if err != nil {
				return stacktrace.WrapError(err)
			}
			return nil
		})
	})

	if err = errGroup.Wait(); err != nil {
		return stacktrace.WrapError(err)
	}
	return nil
}
