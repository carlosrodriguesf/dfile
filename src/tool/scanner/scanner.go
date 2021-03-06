package scanner

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"sync"
)

type (
	ScanOptions struct {
		IgnoreFolders    []string
		AcceptExtensions []string
	}
	Scanner interface {
		Scan(ctx context.Context, path string, opts ...ScanOptions) (allFiles []string, err error)
	}
	scanner struct {
		ignoreFolders    map[string]bool
		acceptExtensions map[string]bool

		errGroup *errgroup.Group
	}
)

func New() Scanner {
	return &scanner{}
}

func (s scanner) Scan(ctx context.Context, path string, opts ...ScanOptions) (allFiles []string, err error) {
	if len(opts) > 0 {
		s.acceptExtensions = mapStringArray(opts[0].AcceptExtensions)
		s.ignoreFolders = mapStringArray(opts[0].IgnoreFolders)
	}

	var mutex sync.Mutex

	allFiles = make([]string, 0)
	s.errGroup, ctx = errgroup.WithContext(ctx)
	s.scan(ctx, path, func(files []string) {
		mutex.Lock()
		allFiles = append(allFiles, files...)
		mutex.Unlock()
	})

	err = s.errGroup.Wait()
	if err != nil {
		log.Printf("error: %v", err)
	}

	return
}

func (s *scanner) scan(ctx context.Context, path string, addFiles func(files []string)) {
	s.errGroup.Go(func() error {
		log.Println("scanning: ", path)

		files := make([]string, 0)

		entries, err := os.ReadDir(path)
		if err != nil {
			log.Printf("error: %v", err)
			return err
		}

		for _, entry := range entries {
			if s.ignore(entry) {
				continue
			}

			info, err := entry.Info()
			if err != nil {
				log.Printf("error: %v", err)
				return err
			}

			entryPath := fmt.Sprintf("%s/%s", path, info.Name())
			if entry.IsDir() {
				s.scan(ctx, entryPath, addFiles)
				continue
			}

			log.Println("found: ", entryPath)
			files = append(files, entryPath)
		}

		addFiles(files)

		log.Println("scanned: ", path)
		return nil
	})
}

func (s *scanner) ignore(entry os.DirEntry) bool {
	name := entry.Name()
	if len(s.ignoreFolders) > 0 && s.ignoreFolders[name] {
		return true
	}
	if len(s.acceptExtensions) > 0 && !entry.IsDir() && !s.acceptExtensions[getExtension(entry.Name())] {
		return true
	}
	return false
}
