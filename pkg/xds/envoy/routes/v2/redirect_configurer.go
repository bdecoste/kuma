package v2

import (
	envoy_route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
)

type RedirectConfigurer struct {
	MatchPath    string
	NewPath      string
	Port         uint32
	AllowGetOnly bool
}

func (c RedirectConfigurer) Configure(virtualHost *envoy_route.VirtualHost) error {
	var headersMatcher []*envoy_route.HeaderMatcher
	if c.AllowGetOnly {
		headersMatcher = []*envoy_route.HeaderMatcher{
			{
				Name: ":method",
				HeaderMatchSpecifier: &envoy_route.HeaderMatcher_ExactMatch{
					ExactMatch: "GET",
				},
			},
		}
	}
	virtualHost.Routes = append(virtualHost.Routes, &envoy_route.Route{
		Match: &envoy_route.RouteMatch{
			PathSpecifier: &envoy_route.RouteMatch_Path{
				Path: c.MatchPath,
			},
			Headers: headersMatcher,
		},
		Action: &envoy_route.Route_Redirect{
			Redirect: &envoy_route.RedirectAction{
				PortRedirect: c.Port,
				PathRewriteSpecifier: &envoy_route.RedirectAction_PathRedirect{
					PathRedirect: c.NewPath,
				},
			},
		},
	})
	return nil
}
