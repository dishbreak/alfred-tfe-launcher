package lib

import (
	"encoding/json"
	"os"
	"os/user"
	"path"
	"time"
)

type ItemCache struct {
	Items           []ListItem `json:"items"`
	ExpiryTime      int64      `json:"expiryTime"`
	timeout         int
	refreshCallback func() ([]ListItem, error)
	filePath        string
}

func NewCache(name string, timeout int, callback func() ([]ListItem, error)) (*ItemCache, error) {
	c := &ItemCache{
		timeout:         timeout,
		refreshCallback: callback,
		ExpiryTime:      -1,
	}

	current, err := user.Current()
	if err != nil {
		return c, err
	}

	c.filePath = path.Join(current.HomeDir, SettingsDir, CacheDir, name+".json")
	return c, nil
}

func (c *ItemCache) Get() ([]ListItem, error) {
	err := c.refresh(false)
	if err != nil {
		return nil, err
	}
	return c.Items, nil
}

func (c *ItemCache) save() error {
	c.ExpiryTime = (int64(c.timeout)) + time.Now().Unix()
	err := os.MkdirAll(path.Dir(c.filePath), 0700)
	if err != nil {
		return err
	}

	fp, err := os.Create(c.filePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	encoder := json.NewEncoder(fp)
	return encoder.Encode(c)
}

func (c *ItemCache) expired() bool {
	return c.timeout != 0 && c.ExpiryTime < time.Now().Unix()
}

func (c *ItemCache) refresh(force bool) error {
	if fp, err := os.Open(c.filePath); err == nil {
		decoder := json.NewDecoder(fp)
		err := decoder.Decode(c)
		fp.Close()
		if err != nil {
			return err
		}
	}
	if force || c.expired() {
		items, err := c.refreshCallback()
		if err != nil {
			return err
		}
		c.Items = items
		return c.save()
	}

	return nil
}
