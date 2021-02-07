# Localsearch Home Assignment Backend

This repository contains the backend part of the localsearch home assignment.

### Building and running the code

To build and run the backend app, you should navigate your terminal to the following path

```
  cmd/backend
```

You need to create a .env file with the following parameters:

```
  host=host
  port=port
  places_url=https://url.to/places
```

Then run these commands

```
  $ go mod download
  $ go build
  $ ./backend -env your.env
```

### Functionalities

The goal of the application is provide data of places for the frontend part of the localsearch home assignment. The app
implements a REST API to serve the frontend and also an HTTP client to fetch the data from the provided upstream API.

The only implemented endpoint is a GET method with one path parameter for the id of the place. Example:

```
    http://host:port/places/:id
```

The app will fetch the place from the upstream API and will return the place with the provided id, if it exists, with the following 
response in JSON:

```
    data: {
        "name": "place",
        "address": "address_of_the_place",
        "openingHours": [
            "dayRange": "monday - friday" // (day range can be multiple or single days of the week
            "timeIntervals": [
                {
                    "start":"15:00"
                    "end":"20:00"
                    "type":"OPEN" // second type is "CLOSED"
                }
            ]
        ]
    }
```

The place from the upstream API is converted to a model that is adapted to the frontend.

If the place does not exist, the API will respond with a `400 Bad Request` error. And a 500 if something unexpected 
happens with the upstream API.