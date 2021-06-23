package commands

import (
	"log"
)

func logServerVersion(c *Client) {
	v, err := c.getServerVersion()
	if err != nil {
		log.Printf("[REST-Client] Unable to read server version: %s", err)
		return
	}

	log.Printf("[REST-Client] Server version: %s", v)
}
