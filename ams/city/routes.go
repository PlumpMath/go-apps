// Copyright 2016 Jittakal Author. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package city

import "github.com/jittakal/go-apps/ams/common"

type CityRoutes []common.Route

var Routes = CityRoutes{
	common.Route{
		"Index",
		"GET",
		"/v1",
		Index,
	},
	common.Route{
		"FindById",
		"GET",
		"/v1/{id:[a-z0-9]+}",
		ById,
	},
	common.Route{
		"FindByName",
		"POST",
		"/v1/find",
		ByName,
	},
	common.Route{
		"Create",
		"POST",
		"/v1",
		Create,
	},
	common.Route{
		"Update",
		"PUT",
		"/v1",
		Update,
	},
}
