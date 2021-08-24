package lib

import (
	"encoding/json"
	"os"
	"os/user"
	"path"

	"github.com/pkg/errors"
)

type Settings struct {
	CacheTimeout int
	filePath     string
}

func NewSettings() (*Settings, error) {
	s := &Settings{
		CacheTimeout: 3600,
	}

	current, err := user.Current()
	if err != nil {
		return s, err
	}

	s.filePath = path.Join(current.HomeDir, SettingsDir, SettingsFilename)
	return s, nil
}

func (s *Settings) Load() error {
	fp, err := os.Open(s.filePath)
	if err != nil {
		if errors.Is(os.ErrExist, err) {
			return err
		} else {
			return nil
		}
	}
	defer fp.Close()

	decoder := json.NewDecoder(fp)
	return decoder.Decode(s)
}

func (s *Settings) Save() error {
	if err := os.MkdirAll(path.Base(s.filePath), 0700); err != nil {
		return err
	}

	fp, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	encoder := json.NewEncoder(fp)
	return encoder.Encode(s)
}
