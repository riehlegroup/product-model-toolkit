package commands

import (
	"log"
)

func logServerVersion(c *Client) {
	log.Print("[REST-Client] Server version: 1.0.0\n")
}
