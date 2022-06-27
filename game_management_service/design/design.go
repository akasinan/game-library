package design

import (
	shareddesign "github.com/sinan/game-library/shared/design"

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

	GRPC(func() {
		Response(microerrors.Unauthorized, CodeUnauthenticated)
		Response(microerrors.Forbidden, CodePermissionDenied)
		Response(microerrors.BadRequest, CodeInvalidArgument)
		Response(microerrors.BadGateway, CodeUnavailable)
		Response(microerrors.InternalError, CodeInternal)
	})

	Method("create", func() {
		Description("Creates a game in the postgres database.")

		Payload(func() {
			Extend(CommonAttributes)
			Reference(GameParameters)
			Attribute("name")
			Attribute("code")
			Attribute("parent_org_id")
			Required("name")
		})

		Result(func() {
			Reference(GameParameters)
			Attribute("id")
			Required("id")
		})
		Error(microerrors.Conflict)

		GRPC(func() {
			Response(CodeOK)
			Response(microerrors.Conflict, CodeAlreadyExists)
		})
	})

	Method("update", func() {
		Description("Update a game’s fields in the database. Would be used if new information on the game was added such as finalized release dates, ports to new platforms etc.")

		Payload(func() {
			Extend(CommonAttributes)
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

		GRPC(func() {
			Response(CodeOK)
			Response(microerrors.NotFound, CodeNotFound)
		})
	})

	Method("show", func() {
		Description("Retrieve and show a game from the database. Will be used to show a game’s full profile when on the game’s page on the site.")

		Payload(func() {
			Extend(CommonAttributes)
			Reference(GameParameters)
			Attribute("id", UInt64, "The ID of the game")
			Required("id")
		})
		Result(Game)
		Error(microerrors.NotFound)

		GRPC(func() {
			Response(CodeOK)
			Response(microerrors.NotFound, CodeNotFound)
		})
	})

	Method("index", func() {
		Description("Retrieve a list of all games from the database. Filter will be used to filter out games for searching by values like name. Can be used to retrieve a game’s ID to use the Show endpoint.")

		Payload(func() {
			Extend(CommonAttributes)
		})
		Result(gameList)

		GRPC(func() {
			Response(CodeOK)
		})
	})

	Method("delete", func() {
		Description("Delete a game from the database by ID. Could be used by site admins in case of an upcoming game’s cancellation or a mistake.")

		Payload(func() {
			Extend(CommonAttributes)
			Reference(GameParameters)
			Attribute("id")
			Required("id")
		})
		Result(Empty)
		Error(microerrors.NotFound)

		GRPC(func() {
			Response(CodeOK)
			Response(microerrors.NotFound, CodeNotFound)
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
		Example(shareddesign.GenreExample)
	})
	Field(5, "platforms", ArrayOf(String), "A list of Platforms the game has been released on. Consoles/PC etc.", func() {
		Example(shareddesign.GenreExample)
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
