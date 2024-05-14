// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// Code generated by internal/generate/acctestconsts/main.go; DO NOT EDIT.

package acctest

import (
	"fmt"
)

const (
	CtOne         = "1"
	CtTagsPercent = "tags.%"
	CtTwo         = "2"
	CtZero        = "0"
)

// ConstOrQuote returns the constant name for the given attribute if it exists.
// Otherwise, it returns the attribute quoted. This is intended for use in
// generated code and templates.
func ConstOrQuote(constant string) string {
	allConstants := map[string]string{
		"1":      "CtOne",
		"tags.%": "CtTagsPercent",
		"2":      "CtTwo",
		"0":      "CtZero",
	}

	if v, ok := allConstants[constant]; ok {
		return fmt.Sprintf("acctest.%s", v)
	}
	return fmt.Sprintf("%q", constant)
}
