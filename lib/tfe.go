package lib

import "github.com/hashicorp/go-tfe"

func NewTfeClient() (*tfe.Client, error) {
	token, err := getToken()
	if err != nil {
		return nil, err
	}

	config := &tfe.Config{
		Token: token,
	}

	return tfe.NewClient(config)
}
