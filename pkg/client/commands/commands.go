package commands

import (
	"github.com/osrgroup/product-model-toolkit/pkg/client/http/rest"
	"log"
)


func logServerVersion(c *rest.Client) {
	v, err := c.GetServerVersion()
	if err != nil {
		log.Printf("[REST-Client] Unable to read server version: %s", err)
		return
	}

	log.Printf("[REST-Client] Server version: %s", v)
}
