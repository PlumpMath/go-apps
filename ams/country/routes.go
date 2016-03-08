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
}
