// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package commands

import (
	"log"
)

func logServerVersion(c *Client) {
	log.Print("[REST-Client] Server version: 1.0.0\n")
}
