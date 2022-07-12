package dbfile

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type (
	Options struct {
		AutoPersist      bool
		AutoPersistCount int
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

		Persist() error
		Load() error
	}
	dbFile struct {
		dbFilePath string

		data data

		autoPersist             bool
		autoPersistCount        int
		autoPersistCountCurrent int
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
	log.Println("set: ", file)
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
	return m.data.Files[file]
}

func (m *dbFile) GetFileKeys() []string {
	i, keys := 0, make([]string, len(m.data.Files))
	for key, _ := range m.data.Files {
		keys[i] = key
		i++
	}
	return keys
}

func (m *dbFile) DelFile(file string) {
	log.Println("del: ", file)
	delete(m.data.Files, file)
}

func (m *dbFile) CreateEntry(file string) {
	m.SetFile(file, FileEntry{})
}

func (m *dbFile) Persist() error {
	dt, err := json.MarshalIndent(m.data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(m.dbFilePath, dt, 0700)
	if err != nil {
		return err
	}
	m.autoPersistCountCurrent = 0
	return nil
}

func (m *dbFile) Load() error {
	dataBytes, err := os.ReadFile(m.dbFilePath)
	if err != nil {
		if err == os.ErrNotExist || err.Error() == fmt.Sprintf("open %s: no such file or directory", m.dbFilePath) {
			return nil
		}
		log.Printf("error: %v", err)
		return err
	}

	err = json.Unmarshal(dataBytes, &m.data)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	return nil
}
