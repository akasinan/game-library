package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"
)

var _ = API("game_library_service", func() {
	Title("Game Library Service")
	Description(`A microservice allowing us to create add new games, retrieve games, update game information and delete games`)
})

/*
// JWTAuth defines a security scheme that uses JWT tokens.
var JWTAuth = JWTSecurity("jwt", func() {
	Description("Secures requests by requiring a valid JWT token retrieved via the Flexera Authorization Service.")
	Scope("iamos:read", "Allows reading from Onboarding Service")
	Scope("iamos:write", "Allows writing to Onboarding Service")
})
*/
// CommonAttributes defines attributes used by all endpoints.
var CommonAttributes = Type("Common", func() {
	TokenField(1, "auth_token", String, "JWT token used to perform authorization", func() {
		Example("1111aaaa.2222bbbb.3333cccc")
	})
})
