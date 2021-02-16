terraform {
  required_providers {
    heroku = {
      source = "heroku/heroku"
    }
  }
}

resource "heroku_app" "outdoorapi" {
  name = "outdoor-api"
  region = "us"
  stack = "container"
}

resource "heroku_build" "outdoorapi" {
  app = heroku_app.outdoorapi.name

  source = {
    path = "api"
  }
}

resource "heroku_app" "outdoorweb" {
  name = "go-outdoors"
  region = "us"
  stack = "container"
}


resource "heroku_build" "outdoorweb" {
  app = heroku_app.outdoorweb.name

  source = {
    "path" = "client"
  }
}