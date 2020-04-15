// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

// Info represents generic additional information related to a product or component.
type Info struct {
	Description       string `json:"description,omitempty"`
	Comment           string `json:"comment,omitempty"`
	HomepageURL       string `json:"homepage-url,omitempty"`
	ExternalReference string `json:"external-ref,omitempty"`
}
