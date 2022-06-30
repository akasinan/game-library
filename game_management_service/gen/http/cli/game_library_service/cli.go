// Code generated by goa v3.7.6, DO NOT EDIT.
//
// game_library_service HTTP client CLI support package
//
// Command:
// $ goa gen github.com/akasinan/game-library/game_management_service/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	gamemanagementc "github.com/akasinan/game-library/game_management_service/gen/http/game_management/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `game-management (create|update|show|index|delete)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` game-management create --body '{
      "developer": "some-developer",
      "genre": [
         "some-genre",
         "some-genre",
         "some-genre"
      ],
      "name": "some-game",
      "platforms": [
         "some-platform",
         "some-platform"
      ],
      "publisher": "some-publisher",
      "release_date": "1994-05-17T20:31:16Z"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		gameManagementFlags = flag.NewFlagSet("game-management", flag.ContinueOnError)

		gameManagementCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		gameManagementCreateBodyFlag = gameManagementCreateFlags.String("body", "REQUIRED", "")

		gameManagementUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		gameManagementUpdateBodyFlag = gameManagementUpdateFlags.String("body", "REQUIRED", "")
		gameManagementUpdateIDFlag   = gameManagementUpdateFlags.String("id", "REQUIRED", "ID of the game in the Postgres database")

		gameManagementShowFlags  = flag.NewFlagSet("show", flag.ExitOnError)
		gameManagementShowIDFlag = gameManagementShowFlags.String("id", "REQUIRED", "The ID of the game")

		gameManagementIndexFlags      = flag.NewFlagSet("index", flag.ExitOnError)
		gameManagementIndexFilterFlag = gameManagementIndexFlags.String("filter", "REQUIRED", "")

		gameManagementDeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		gameManagementDeleteIDFlag = gameManagementDeleteFlags.String("id", "REQUIRED", "ID of the game in the Postgres database")
	)
	gameManagementFlags.Usage = gameManagementUsage
	gameManagementCreateFlags.Usage = gameManagementCreateUsage
	gameManagementUpdateFlags.Usage = gameManagementUpdateUsage
	gameManagementShowFlags.Usage = gameManagementShowUsage
	gameManagementIndexFlags.Usage = gameManagementIndexUsage
	gameManagementDeleteFlags.Usage = gameManagementDeleteUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "game-management":
			svcf = gameManagementFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "game-management":
			switch epn {
			case "create":
				epf = gameManagementCreateFlags

			case "update":
				epf = gameManagementUpdateFlags

			case "show":
				epf = gameManagementShowFlags

			case "index":
				epf = gameManagementIndexFlags

			case "delete":
				epf = gameManagementDeleteFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "game-management":
			c := gamemanagementc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = gamemanagementc.BuildCreatePayload(*gameManagementCreateBodyFlag)
			case "update":
				endpoint = c.Update()
				data, err = gamemanagementc.BuildUpdatePayload(*gameManagementUpdateBodyFlag, *gameManagementUpdateIDFlag)
			case "show":
				endpoint = c.Show()
				data, err = gamemanagementc.BuildShowPayload(*gameManagementShowIDFlag)
			case "index":
				endpoint = c.Index()
				data, err = gamemanagementc.BuildIndexPayload(*gameManagementIndexFilterFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = gamemanagementc.BuildDeletePayload(*gameManagementDeleteIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// game-managementUsage displays the usage of the game-management command and
// its subcommands.
func gameManagementUsage() {
	fmt.Fprintf(os.Stderr, `Service is the GameManagement service interface.
Usage:
    %[1]s [globalflags] game-management COMMAND [flags]

COMMAND:
    create: Creates a game in the postgres database.
    update: Update a game’s fields in the database. Would be used if new information on the game was added such as finalized release dates, ports to new platforms etc.
    show: Retrieve and show a game from the database. Will be used to show a game’s full profile when on the game’s page on the site.
    index: Retrieve a list of all games from the database. Filter will be used to filter out games for searching by values like name. Can be used to retrieve a game’s ID to use the Show endpoint.
    delete: Delete a game from the database by ID. Could be used by site admins in case of an upcoming game’s cancellation or a mistake.

Additional help:
    %[1]s game-management COMMAND --help
`, os.Args[0])
}
func gameManagementCreateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] game-management create -body JSON

Creates a game in the postgres database.
    -body JSON: 

Example:
    %[1]s game-management create --body '{
      "developer": "some-developer",
      "genre": [
         "some-genre",
         "some-genre",
         "some-genre"
      ],
      "name": "some-game",
      "platforms": [
         "some-platform",
         "some-platform"
      ],
      "publisher": "some-publisher",
      "release_date": "1994-05-17T20:31:16Z"
   }'
`, os.Args[0])
}

func gameManagementUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] game-management update -body JSON -id UINT64

Update a game’s fields in the database. Would be used if new information on the game was added such as finalized release dates, ports to new platforms etc.
    -body JSON: 
    -id UINT64: ID of the game in the Postgres database

Example:
    %[1]s game-management update --body '{
      "developer": "some-developer",
      "genre": [
         "some-genre",
         "some-genre",
         "some-genre"
      ],
      "platforms": [
         "some-platform",
         "some-platform",
         "some-platform",
         "some-platform"
      ],
      "publisher": "some-publisher",
      "release_date": "1971-01-04T03:55:57Z"
   }' --id 1234
`, os.Args[0])
}

func gameManagementShowUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] game-management show -id UINT64

Retrieve and show a game from the database. Will be used to show a game’s full profile when on the game’s page on the site.
    -id UINT64: The ID of the game

Example:
    %[1]s game-management show --id 1234
`, os.Args[0])
}

func gameManagementIndexUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] game-management index -filter STRING

Retrieve a list of all games from the database. Filter will be used to filter out games for searching by values like name. Can be used to retrieve a game’s ID to use the Show endpoint.
    -filter STRING: 

Example:
    %[1]s game-management index --filter "(name co 'Mario' or name co 'Super Mario 64')"
`, os.Args[0])
}

func gameManagementDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] game-management delete -id UINT64

Delete a game from the database by ID. Could be used by site admins in case of an upcoming game’s cancellation or a mistake.
    -id UINT64: ID of the game in the Postgres database

Example:
    %[1]s game-management delete --id 1234
`, os.Args[0])
}
