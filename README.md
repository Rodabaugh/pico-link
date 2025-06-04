# About
pico-link is a lightweight URL shortener.  There is no authentication yet, and the admin page is shared by all users.

# This Project Involved
- Getting a list of links from the database and rendering them on the page using templ.
- Getting a new link from the user using HTMX, passing it to the backend, validating it, and storing it in the database.
- Having HTMX automatically update the list of links after writing to the database.
- Dynamically passing the ID of a link to the DELETE /api/links endpoint so the backend can delete the link.
- Updating the links list after the delete has been performed.

The new link and delete link endpoints check to see if the response should be in JSON and will reply in JSON if it is. This allows the use of a custom client over the API. If the Accept is not set to "application/json", the backend will reply with HTML so HTMX can render it on the page.

# Hosting
## Get the repo

1. Clone the repo `git clone https://github.com/Rodabaugh/pico-link/`
2. Navigate to the program dir `cd pico-link`

## Configuration

Environment variables are used for configuration. Create a .env file in the root of the project dir. You need to specify your DB_URL, PLATFORM. DB_URL is the url for your database. Platform can either be "prod" or "dev". Your `.env` file should look something like the one below. Please be sure to use your own DB_URL.
```
PLATFORM=prod
DB_URL="postgres://postgresUser:postgresPass@localhost:5432/pico_link?sslmode=disable"
```

A port may also be specified using ```PORT=1234```. If a port is not specified, it will default to 8080.

## Setting up the database

Goose is used to manage the database migrations. Install goose with `go install github.com/pressly/goose/v3/cmd/goose@latest`

Navigate to the sql/schema dir `cd sql/schema`

Setup the database using goose `goose postgres <connection_string> up` e.g `goose postgres postgres://postgresUser:postgresPass@localhost:5432/pico_link?sslmode=disable up`

## Compile and run the application

Once your .env has been configured, and your database is setup, it is time to build and run the application.

Build the application with `make build`

Run the backend application with `./pico-link`

Once the application is running, you can setup your server to run the application as a service.

# API Endpoints
Be sure to set your Accept header to "application/json", so the backend responds with JSON and not HTML.

## POST /api/links
Request body:
```json
{
  "link_name": "test",
  "link_url": "https://erikrodabaugh.com/"
}
```

Response body:
```json
{
	"ID": "953b569d-4a26-43c9-ba29-2c9c431bf6f7",
	"CreatedAt": "2025-06-03T20:41:36.035423Z",
	"UpdatedAt": "2025-06-03T20:41:36.035423Z",
	"LinkName": "test",
	"LinkUrl": "https://erikrodabaugh.com/"
}
```

## GET /api/links
Response body:
```json
[
	{
		"ID": "953b569d-4a26-43c9-ba29-2c9c431bf6f7",
		"CreatedAt": "2025-06-03T20:41:36.035423Z",
		"UpdatedAt": "2025-06-03T20:41:36.035423Z",
		"LinkName": "test",
		"LinkUrl": "https://erikrodabaugh.com/"
	},
	{
		"ID": "2d6249b6-726e-49b0-a122-51630f7ecdf1",
		"CreatedAt": "2025-06-03T20:42:34.758915Z",
		"UpdatedAt": "2025-06-03T20:42:34.758915Z",
		"LinkName": "test2",
		"LinkUrl": "https://erikrodabaugh.com"
	}
]
```

## DELETE /api/posts/{link_id}

No content is required for this request. An empty link object will be returned, along with a 200.

Response body:
```json
{
	"ID": "00000000-0000-0000-0000-000000000000",
	"CreatedAt": "0001-01-01T00:00:00Z",
	"UpdatedAt": "0001-01-01T00:00:00Z",
	"LinkName": "",
	"LinkUrl": ""
}
```