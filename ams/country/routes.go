package country

import "github.com/jittakal/go-apps/ams/common"

type CountryRoutes []common.Route

var Routes = CountryRoutes{
	common.Route{
		"Index",
		"GET",
		"/v1",
		Index,
	},
	common.Route{
		"FindId",
		"GET",
		"/v1/{id:[a-z0-9]+}",
		FindId,
	},
	common.Route{
		"Create",
		"POST",
		"/v1",
		Create,
	},
}
