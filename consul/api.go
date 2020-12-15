package consulapi

import (
	"github.com/hashicorp/consul/api"
	"os"
)

func Upload(key string, value []byte) error {
	consuladdr := "http://127.0.0.1:8500"
	if os.Getenv("CONSUL_ADDRESS") != "" {
		consuladdr = os.Getenv("CONSUL_ADDRESS")
	}

	c, err := api.NewClient(&api.Config{
		Address: consuladdr,
	})
	if err != nil {
		return err
	}

	kv := c.KV()

	_, err = kv.Put(&api.KVPair{
		Key:   key,
		Value: value,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}
