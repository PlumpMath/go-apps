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
