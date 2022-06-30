package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"
)

var _ = API("game_library_service", func() {
	Title("Game Library Service")
	Description(`A microservice allowing us to create add new games, retrieve games, update game information and delete games`)
})
