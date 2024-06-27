# Movies API

This is a simple RESTful API for managing a collection of movies using Go, Gorilla Mux, and Swagger for API documentation.

## Features
- Get all movies
- Get a movie by ID
- Create a new movie
- Update a movie by ID
- Delete a movie by ID
- Swagger documentation

The server will start on port 8000. You can access the API at http://localhost:8000.

# API Endpoints

## Get all movies

- **URL:** `/movies`
- **Method:** `GET`
- **Response:**
  - `200 OK` with a list of movies.

## Get a movie by ID

- **URL:** `/movies/{id}`
- **Method:** `GET`
- **URL Params:** `id=[string]`
- **Response:**
  - `200 OK` with the movie object.
  - `404 Not Found` if the movie does not exist.

## Create a new movie

- **URL:** `/movies`
- **Method:** `POST`
- **Body:** JSON object with movie details.
- **Response:**
  - `201 Created` with the created movie object.
  - `400 Bad Request` if the request body is invalid.

## Update a movie by ID

- **URL:** `/movies/{id}`
- **Method:** `PUT`
- **URL Params:** `id=[string]`
- **Body:** JSON object with updated movie details.
- **Response:**
  - `200 OK` with the updated movie object.
  - `400 Bad Request` if the request body is invalid.
  - `404 Not Found` if the movie does not exist.

## Delete a movie by ID

- **URL:** `/movies/{id}`
- **Method:** `DELETE`
- **URL Params:** `id=[string]`
- **Response:**
  - `200 OK` with the remaining movies.
  - `404 Not Found` if the movie does not exist.

## Swagger Documentation

Swagger documentation is available at [http://localhost:8000/docs](http://localhost:8000/docs).
