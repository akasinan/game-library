# Sinan's Game Library

The game_management_service is backend GRPC service for a proposed IMDB/Letterboxd style service for video games. The service should allow a site admin to add new games, retrieve games, update game information and delete games from a database. 

The Games resource type will contain the following variables

Games
* Name
* Release Date
* Genre
* Platforms
* Developer
* Publisher
* CreatedAt
* UpdatedAt

The following endpoints are to be used in this service:

## Create

Creates a game in the postgres database.

### Payload
| Param	| Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| Name | String | * | Name of the game |
| Release Date | DateTime |  | Initial release date of game |
| Genre | listof(string) |  | A list of Genres the game falls under. A list of strings |
| Platforms | listof(string) |  | A list of Platforms the game has been released on. Consoles/PC etc. A list of strings |
| Developer | String |  | Name of Game’s developer |
| Publisher | String |  | Name of Game’s publisher |

Note: Params for create besides name are optional as when a game has been announced the other parameters such as platforms and release date might not have been finalized yet.

### Response
| Param	| Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| ID | Uint64 | * | ID of the game in the Postgres database |

### Errors
|Error Code	| Error Type |
| --------- | ---------- |
| 422 | Unprocessable Entity |

Unprocessable entity error will occur if a game is created in the database with an invalid genre/platform/Developer/Publisher

Note: Might want to add functionality to add Developer/Publisher alongside game if missing from Database

## Update

Update a game’s fields in the database. Would be used if new information on the game was added such as finalized release dates, ports to new platforms etc.

### Payload
| Param	| Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| ID | Uint64 | * | ID of the game in the Postgres database |
| Name | String | * | Name of the game |
| Release Date | DateTime |  | Initial release date of game |
| Genre | listof(string) |  | A list of Genres the game falls under. A list of strings |
| Platforms | listof(string) |  | A list of Platforms the game has been released on. Consoles/PC etc. A list of strings |
| Developer | String |  | Name of Game’s developer |
| Publisher | String |  | Name of Game’s publisher |

### Response
No content

### Errors
|Error Code	| Error Type |
| --------- | ---------- |
| 422 | Unprocessable Entity |
| 404 | NotFound |

Note for unprocessable Entity same as Create

## Show

Retrieve and show a game from the database. Will be used to show a game’s full profile when on the game’s page on the site. 

### Payload
| Param	| Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| ID | Uint64 | * | ID of the game in the Postgres database |

### Response

GameMediaType is returned by the Show Action

| Param	| Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| ID | Uint64 | * | ID of the game in the Postgres database |
| Name | String | * | Name of the game |
| Release Date | DateTime |  | Initial release date of game |
| Genre | listof(string) |  | A list of Genres the game falls under. A list of strings |
| Platforms | listof(string) |  | A list of Platforms the game has been released on. Consoles/PC etc. A list of strings |
| Developer | String |  | Name of Game’s developer |
| Publisher | String |  | Name of Game’s publisher |
| CreatedAt | DateTime |  | DateTime at which the game record was created in the database | 
| UpdatedAt	| DateTime |  | DateTime at which the game record was last updated in the database |

Note: In the future, the show endpoint should be updated to pull a staff list from a different table

### Errors
|Error Code	| Error Type |
| --------- | ---------- |
| 404 | NotFound |

## Index

Retrieve a list of all games from the database. Filter will be used to filter out games for searching by values like name. Can be used to retrieve a game’s ID to use the Show endpoint. 

### Payload
| Param	| Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| Filter | String |  | Optional filter used to filter out games in the index by values like name |

### Response

Array of GameMediaType is returned by the Index Action. Without a filter, all games in the database will be returned.

### Error 
n/a

## Delete

Delete a game from the database by ID. Could be used by site admins in case of an upcoming game’s cancellation or a mistake.

### Payload
| Param	| Type | Required | Description |
| ----- | ---- | -------- | ----------- |
| ID | Uint64 | * | ID of the game in the Postgres database |

### Response
No content


### Errors:
|Error Code	| Error Type |
| --------- | ---------- |
| 404 | NotFound |

## General Errors
|Error Code	| Error Type |
| --------- | ---------- |
| 401 | Unauthorized | 
| 403 | Forbidden | 
| 400 | BadRequest | 
| 502 | BadGateway | 
| 500 | InternalError | 

