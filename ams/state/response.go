// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package state

type HyperLink struct {
	Relation string `json:"rel"`
	URL      string `json:"url"`
	Method   string `json:"method"`
}

type ResponseState struct {
	State
	Links []HyperLink `json:"links"`
}
