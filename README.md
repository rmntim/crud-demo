# Simple CRUD demo with Go

## Usage
```console
$ go build
$ ./crud-demo
```

## Endpoints

- `/`: Welcome page
- `/movies`:
    - `GET`: Get all movies
    - `POST` Create a movie
- `/movies/{id}`:
    - `GET`: Get movie by specified `id`
    - `PUT`: Update a movie with specified `id`
    - `DELETE`: Delete a movie with specified `id`

## JSON movie layout

```json
{
  "id": string,
  "isbn": string,
  "title": string,
  "director": {
    "firstname": string,
    "lastname": string
  }
}
```
