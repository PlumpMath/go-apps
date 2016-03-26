// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package state

import "github.com/jittakal/go-apps/ams/common"

type StateRoutes []common.Route

var Routes = StateRoutes{
	common.Route{
		"Index",
		"GET",
		"/v1",
		Index,
	},
	common.Route{
		"IndexWithLink",
		"GET",
		"/v1/link",
		IndexWithLink,
	},
	common.Route{
		"Create",
		"POST",
		"/v1",
		Create,
	},
}
