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
	Entry struct {
		Ready bool   `json:"ready"`
		Hash  string `json:"content,omitempty"`
		Error string `json:"error,omitempty"`
	}
	DBFile interface {
		CreateEntry(path string)
		Set(file string, entry Entry)
		Get(file string) (Entry, bool)
		Has(file string) bool
		GetCopy() map[string]Entry
		Persist() error
		Load() error
	}
	dbFile struct {
		dbFilePath string

		data map[string]Entry

		autoPersist             bool
		autoPersistCount        int
		autoPersistCountCurrent int
	}
)

func New(path string, opts ...Options) DBFile {
	log.Println("db.path:", path)
	var dbFile = dbFile{
		dbFilePath: path,
		data:       make(map[string]Entry),
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

func (m *dbFile) Has(file string) bool {
	_, ok := m.data[file]
	return ok
}

func (m *dbFile) Get(file string) (Entry, bool) {
	result, ok := m.data[file]
	return result, ok
}

func (m *dbFile) Set(file string, result Entry) {
	log.Println("set: ", file, result)
	m.data[file] = result
	if m.autoPersist {
		m.autoPersistCountCurrent++
		if m.autoPersistCountCurrent == m.autoPersistCount {
			if err := m.Persist(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (m *dbFile) CreateEntry(file string) {
	m.Set(file, Entry{})
}

func (m *dbFile) GetPreviousDataKeyMap() map[string]bool {
	fileMap := make(map[string]bool)
	for key, _ := range m.data {
		fileMap[key] = true
	}
	return fileMap
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

func (m *dbFile) GetCopy() map[string]Entry {
	copied := make(map[string]Entry)
	for path, entry := range m.data {
		copied[path] = entry
	}
	return copied
}
