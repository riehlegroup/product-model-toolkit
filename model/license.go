// SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
//
// SPDX-License-Identifier: Apache-2.0

package model

// License represents a open source license.
type License struct {
	SPDXID           string `json:"spdxId,omitempty"`
	DeclaredLicense  string `json:"declaredLicense,omitempty"`
	ConcludedLicense string `json:"concludedLicense,omitempty"`
}

func (l *License) toString() string {
	if l.SPDXID != "" {
		return l.SPDXID
	}

	if l.ConcludedLicense != "" {
		return l.ConcludedLicense
	}

	return l.DeclaredLicense
}
