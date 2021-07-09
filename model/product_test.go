// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

import (
	"encoding/json"
	"reflect"
	"testing"
)

var prod = Product{
	// ID:                1,
	Name:              "Product-Model-Tookit",
	Version:           "0.1.0",
	HomepageURL:       "osr.cs.fau.de/",
	Description:       "Manage OSS dependencies in software products",
	ExternalReference: "https://github.com/osrgroup/product-model-toolkit",
	Comment:           "WIP",
	VCS:               "https://github.com/osrgroup/product-model-toolkit",
	UsageTypes:        []UsageType{OnPremise, CloudService},
}

const prodStr string = `
  {
	"id": 1,
	"name": "Product-Model-Tookit",
	"description": "Manage OSS dependencies in software products",
	"comment": "WIP",
	"homepageUrl": "osr.cs.fau.de/",
	"externalRef": "https://github.com/osrgroup/product-model-toolkit",
	"version": "0.1.0",
	"vcs": "https://github.com/osrgroup/product-model-toolkit",
	"components": null,
	"usageTypes": [
	  "on-premise",
	  "cloud-service"
	]
  }
`

func TestJSON_Marshal(t *testing.T) {
	_, err := json.MarshalIndent(prod, "", "  ")
	if err != nil {
		t.Errorf("Expected marshaling product to JSON without errors, but got %v", err)
	}
}

func TestJSON_Unmarshal(t *testing.T) {
	var result Product
	err := json.Unmarshal([]byte(prodStr), &result)
	if err != nil {
		t.Errorf("Expected unmarshaling JSON to product without errors, but got %v", err)
	}

	if !reflect.DeepEqual(prod, result) {
		t.Error("Expected unmarshaled product to be equal with reference")
	}
}
