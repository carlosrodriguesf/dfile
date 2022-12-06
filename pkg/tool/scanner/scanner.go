package scanner

import (
	"context"
	"fmt"
	"github.com/carlosrodriguesf/dfile/pkg/tool/array"
	"github.com/carlosrodriguesf/dfile/pkg/tool/stack"
	"github.com/carlosrodriguesf/dfile/pkg/tool/stacktrace"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
)

type (
	File struct {
		Prev *File
		Next *File
		Path string
	}
	Option  func(entry os.DirEntry) bool
	Scanner interface {
		Scan(ctx context.Context, path string, opts ...Option) ([]string, error)
	}
	scanner struct {
	}
)

func New() Scanner {
	return scanner{}
}

func (s scanner) Scan(ctx context.Context, path string, opts ...Option) ([]string, error) {
	files, err := s.scan(ctx, path, opts)
	if err != nil {
		fmt.Println(stacktrace.WrapError(err))
		return nil, stacktrace.WrapError(err)
	}
	return files.ToArray(), nil
}

func (s scanner) scan(ctx context.Context, path string, opts []Option) (stack.Stack[string], error) {
	log.Printf("scanning '%s'\n", path)

	var files stack.Stack[string]

	errGroup, ctx := errgroup.WithContext(ctx)

	entries, err := os.ReadDir(path)
	if err != nil {
		return files, stacktrace.WrapError(err)
	}

	array.Each(ctx, entries, func(entry os.DirEntry) {
		entryPath := fmt.Sprintf("%s/%s", path, entry.Name())

		if !acceptAllOptions(entry, opts) {
			log.Printf("ignoring '%s'\n", entryPath)
			return
		}

		log.Printf("found '%s'\n", entryPath)

		if !entry.IsDir() {
			files.Append(entryPath)
			return
		}

		errGroup.Go(func() error {
			children, err := s.scan(ctx, entryPath, opts)
			if err != nil {
				return stacktrace.WrapError(err)
			}
			files.AppendList(&children)
			return nil
		})
	})

	if err := errGroup.Wait(); err != nil {
		return stack.Stack[string]{}, stacktrace.WrapError(err)
	}

	return files, nil
}
