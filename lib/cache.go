package lib

import (
	"encoding/json"
	"os"
	"os/user"
	"path"

	"github.com/dishbreak/terraform-cloud-launcher/models"
)

type ItemCache struct {
	Items           []models.ListItem `json:"items"`
	ExpiryTime      int               `json:"expiryTime"`
	refreshCallback func() ([]models.ListItem, error)
	filePath        string
}

func NewCache(name string, timeout int, callback func() []models.ListItem) (*ItemCache, error) {
	c := &ItemCache{}

	current, err := user.Current()
	if err != nil {
		return c, err
	}

	c.filePath = path.Join(current.HomeDir, SettingsDir, CacheDir, name+".json")
	return c, nil
}

func (c *ItemCache) Get() ([]models.ListItem, error) {

}

func (c *ItemCache) save() error {
	err := os.MkdirAll(path.Base(c.filePath), 0700)
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

func (c *ItemCache) refresh(force bool) error {
	if fp, err := os.Open(c.filePath); err == nil {
		defer fp.Close()
		decoder := json.NewDecoder(fp)
		err := decoder.Decode(c)
		if err != nil {
			return err
		}
	} else {
		items, err := c.refreshCallback()
		if err != nil {
			return err
		}
		c.Items = items
		return c.save()
	}
	return nil
}
