package dbfile

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

type (
	Options struct {
		AutoPersist      bool
		AutoPersistCount int
	}

	PersistOptions struct {
		Indented bool
	}

	DBFile interface {
		HasPath(path string) bool
		SetPath(path string, entry PathEntry)
		DelPath(path string)
		GetPath(path string) PathEntry
		GetPathKeys() []string

		HasFile(file string) bool
		SetFile(file string, entry FileEntry)
		DelFile(file string)
		GetFile(file string) FileEntry
		GetFileKeys() []string

		Persist(opts ...PersistOptions) error
		Load() error
	}
	dbFile struct {
		dbFilePath string

		data data

		autoPersist             bool
		autoPersistCount        int
		autoPersistCountCurrent int

		persistMutex sync.Mutex
	}
)

func New(path string, opts ...Options) DBFile {
	log.Println("db.path:", path)
	log.Println("db.version: 2")
	var dbFile = dbFile{
		dbFilePath: path,
		data: data{
			Version: 2,
			Paths:   map[string]PathEntry{},
			Files:   map[string]FileEntry{},
		},
	}

	if len(opts) > 0 {
		if opt := opts[0]; opt.AutoPersist {
			dbFile.autoPersist = true
			dbFile.autoPersistCount = opt.AutoPersistCount
			if dbFile.autoPersistCount == 0 {
				dbFile.autoPersistCount = 10000
			}

			log.Println("db.autoPersist: true")
			log.Println("db.autoPersistCount:", dbFile.autoPersistCount)
		}
	}

	return &dbFile
}

func (m *dbFile) HasPath(path string) bool {
	_, ok := m.data.Paths[path]
	return ok
}

func (m *dbFile) SetPath(path string, entry PathEntry) {
	m.data.Paths[path] = entry
}

func (m *dbFile) DelPath(path string) {
	delete(m.data.Paths, path)
}

func (m *dbFile) GetPath(path string) PathEntry {
	return m.data.Paths[path]
}

func (m *dbFile) GetPathKeys() []string {
	i, keys := 0, make([]string, len(m.data.Paths))
	for key, _ := range m.data.Paths {
		keys[i] = key
		i++
	}
	return keys
}

func (m *dbFile) HasFile(file string) bool {
	_, ok := m.data.Files[file]
	return ok
}

func (m *dbFile) SetFile(file string, result FileEntry) {
	log.Println("set file:", file)
	m.data.Files[file] = result
	if m.autoPersist {
		m.autoPersistCountCurrent++
		if m.autoPersistCountCurrent == m.autoPersistCount {
			if err := m.Persist(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (m *dbFile) GetFile(file string) FileEntry {
	log.Println("get file:", file)
	return m.data.Files[file]
}

func (m *dbFile) GetFileKeys() []string {
	log.Println("get file keys")
	i, keys := 0, make([]string, len(m.data.Files))
	for key, _ := range m.data.Files {
		keys[i] = key
		i++
	}
	return keys
}

func (m *dbFile) DelFile(file string) {
	log.Println("del file:", file)
	delete(m.data.Files, file)
}

func (m *dbFile) Persist(opts ...PersistOptions) error {
	m.persistMutex.Lock()
	defer m.persistMutex.Unlock()

	log.Println("recreating db file")
	file, err := os.Create(m.dbFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if len(opts) > 0 && opts[0].Indented {
		encoder.SetIndent("", "  ")
	}

	log.Println("saving db data")
	err = encoder.Encode(m.data)
	if err != nil {
		return err
	}

	return nil
}

func (m *dbFile) Load() error {
	log.Println("opening db file:", m.dbFilePath)
	file, err := os.OpenFile(m.dbFilePath, os.O_RDWR, 0700)
	if err != nil {
		if err == os.ErrNotExist || err.Error() == fmt.Sprintf("open %s: no such file or directory", m.dbFilePath) {
			return nil
		}
		return err
	}
	defer file.Close()

	log.Println("loading db data")
	err = json.NewDecoder(file).Decode(&m.data)
	if err != nil {
		return err
	}

	return nil
}
