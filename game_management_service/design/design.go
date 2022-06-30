package design

import (
	"net/http"

	shareddesign "github.com/akasinan/game-library/shared/design"
	microerrors "github.com/flexera/micro/errors"
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"
)

var _ = Service("GameManagement", func() {
	Error(microerrors.Unauthorized)
	Error(microerrors.Forbidden)
	Error(microerrors.BadRequest)
	Error(microerrors.BadGateway)
	Error(microerrors.InternalError)

	HTTP(func() {
		Response(microerrors.Unauthorized, http.StatusUnauthorized)
		Response(microerrors.Forbidden, http.StatusForbidden)
		Response(microerrors.BadRequest, http.StatusBadRequest)
		Response(microerrors.BadGateway, http.StatusBadGateway)
		Response(microerrors.InternalError, http.StatusInternalServerError)
		Path("/games")
	})

	Method("create", func() {
		Description("Creates a game in the postgres database.")

		Payload(func() {
			Reference(GameParameters)
			Attribute("name")
			Attribute("release_date")
			Attribute("genre")
			Attribute("platforms")
			Attribute("developer")
			Attribute("publisher")
			Required("name")
		})

		Result(func() {
			Attribute("location", String, "Location of the newly created game")
		})
		Error(microerrors.Conflict)

		HTTP(func() {
			POST("")
			Response(http.StatusCreated, func() {
				Header("location:Location")
				Description("The game was created successfully")
			})
			Response(microerrors.Conflict, http.StatusConflict)
		})
	})

	Method("update", func() {
		Description("Update a game’s fields in the database. Would be used if new information on the game was added such as finalized release dates, ports to new platforms etc.")

		Payload(func() {
			Reference(GameParameters)
			Attribute("id")
			Attribute("release_date")
			Attribute("genre")
			Attribute("platforms")
			Attribute("developer")
			Attribute("publisher")
			Required("id")
		})
		Result(Empty)
		Error(microerrors.NotFound)
		Error(microerrors.NotFound)

		HTTP(func() {
			PUT("/{id}")
			Response(http.StatusNoContent)
			Response(microerrors.NotFound, http.StatusNotFound)
		})
	})

	Method("show", func() {
		Description("Retrieve and show a game from the database. Will be used to show a game’s full profile when on the game’s page on the site.")

		Payload(func() {
			Reference(GameParameters)
			Attribute("id", UInt64, "The ID of the game")
			Required("id")
		})
		Result(Game)
		Error(microerrors.NotFound)

		HTTP(func() {
			GET("/{id}")
			Response(http.StatusOK)
			Response(microerrors.NotFound, http.StatusNotFound)
		})
	})

	Method("index", func() {
		Description("Retrieve a list of all games from the database. Filter will be used to filter out games for searching by values like name. Can be used to retrieve a game’s ID to use the Show endpoint.")

		Payload(func() {
			Reference(GameParameters)
			Attribute("filter")
			Required("filter")
		})
		Result(gameList)
		Error(microerrors.UnprocessableEntity)

		HTTP(func() {
			GET("")
			Body(Empty)
			Param("filter")
			Response(http.StatusOK)
			Response(microerrors.UnprocessableEntity, http.StatusUnprocessableEntity)
		})
	})

	Method("delete", func() {
		Description("Delete a game from the database by ID. Could be used by site admins in case of an upcoming game’s cancellation or a mistake.")

		Payload(func() {
			Reference(GameParameters)
			Attribute("id")
			Required("id")
		})
		Result(Empty)
		Error(microerrors.NotFound)

		HTTP(func() {
			DELETE("/{id}")
			Response(http.StatusNoContent)
			Response(microerrors.NotFound, http.StatusNotFound)
		})
	})
})

var GameParameters = Type("GameParameters", func() {
	Field(1, "id", UInt64, "ID of the game in the Postgres database", func() {
		Example(shareddesign.IDExample)
		Minimum(shareddesign.IDMinValue)
	})
	Field(2, "name", String, "Name of the game", func() {
		Example(shareddesign.GameNameExample)
	})
	Field(3, "release_date", String, "Initial release date of game", func() {
		Format(FormatDateTime)
	})
	Field(4, "genre", ArrayOf(String), "An array of genres the game falls under", func() {
		Elem(func() {
			Example(shareddesign.GenreExample)
		})
	})
	Field(5, "platforms", ArrayOf(String), "A list of Platforms the game has been released on. Consoles/PC etc.", func() {
		Elem(func() {
			Example(shareddesign.PlatformExample)
		})
	})
	Field(6, "developer", String, "Name of Game’s developer", func() {
		Example(shareddesign.DeveloperExample)
	})
	Field(7, "publisher", String, "Name of Game’s publisher", func() {
		Example(shareddesign.PublisherExample)
	})
	Field(8, "created_at", String, "Game creation timestamp", func() {
		Format(FormatDateTime)
	})
	Field(9, "updated_at", String, "Game last updated timestamp", func() {
		Format(FormatDateTime)
	})
	Field(10, "filter", String, `Filter for filtering list of games returned. 
The following filters are supported:
| Filter | Description | Allowed Operator | Behavior | Example |
| --- | ---| --- | --- | --- |
| name | Filters on the game name | co | Contains - The entire operator value must be a substring of the attribute value for a match. |  name co 'Remaster' |
|	|	| eq | Equal - The attribute and operator values must be identical for a match. | name eq 'Super Mario 64'  |
* All operators are case-sensitive and only lowercase is allowed. 
		e.g: 
			valid	: eq
			invalid : EQ

* Filters can be combined using the and/or logical operator. 
		e.g: (name co 'Mario' or name eq 'Super Mario 64')
	
* When used as a query string, the filter parameter value must be URL encoded.		 
		e.g: ?filter=(name%20co%20'Mario'%20or%20name%20eq%20'Super Mario 64')

	When using cURL, the same expression can be URL encoded as follows: 
		$ curl -i -H "$AUTH_HEADER" {$BASE_URL}/games \
			-G --data-urlencode "filter=(name co 'Mario' or name eq 'Super Mario 64')"
`, func() {
		Example(shareddesign.GameFilterExample)
		MinLength(shareddesign.GameFilterMinLength)
	})
})

// gameList defines the collection of games
var gameList = ResultType("game_list", func() {
	TypeName("GameList")
	Description(`A collection of games.`)
	Field(1, "values", CollectionOf(Game), func() {
		Description("List of games")
	})
	Required("values")
})

var Game = ResultType("game", func() {
	TypeName("Game")
	Description(`A game is a collection of values representing a game in the database.`)

	Attributes(func() {
		Reference(GameParameters)
		Attribute("id")
		Attribute("name")
		Attribute("release_date")
		Attribute("genre")
		Attribute("platforms")
		Attribute("developer")
		Attribute("publisher")
		Attribute("created_at")
		Attribute("updated_at")
	})
	Required("id", "name", "created_at")
})
