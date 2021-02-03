# Localsearch Home Assignment Backend

This repository contains the backend part of the localsearch home assignment.

# Building and running the code

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
