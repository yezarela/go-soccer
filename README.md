# go-soccer

Simple soccer REST API with Echo & MongoDB

## How to run this project

This project is using go module, so you don't have to clone it inside GOPATH directory.

### Run the app

```
# Clone the repository
git clone https://github.com/yezarela/go-soccer.git

# Navigate to project directory
cd go-soccer

# Install dependencies
make deps

# Run the application
make run

# Hit the endpoint
curl localhost:1323/teams
```

If you want to use your own mongodb uri, you can update the .env file with your own mongodb uri

### Run the test
```
make test
```

## API Reference

These are the endpoints available from the app

### `GET /teams`

Returns list of teams 

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": [
    {
      "id": "5f6a5d6129b2289c40b7444b",
      "name": "AC Milan 2",
      "description": "some-description",
      "location": "Italy",
      "players": [
        {
          "id": "5f6a5d6129b2289c40b74448",
          "name": "John Doe 1",
          "nickname": "Lolo",
          "position": "forward",
          "created_at": "2020-09-22T20:24:01.872Z"
        }
      ],
      "created_at": "2020-09-22T20:24:01.846Z"
    }
  ]
}
```

</p>
</details>

---

### `GET /teams/:id`

Returns a team by id

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": {
    "id": "5f6a5d6129b2289c40b7444b",
    "name": "AC Milan 2",
    "description": "some-description",
    "location": "Italy",
    "players": [
      {
        "id": "5f6a5d6129b2289c40b74448",
        "name": "John Doe 1",
        "nickname": "Lolo",
        "position": "forward",
        "created_at": "2020-09-22T20:24:01.872Z"
      }
    ],
    "created_at": "2020-09-22T20:24:01.846Z"
  }
}
```

</p>
</details>

---

### `POST /teams`

Creates a new team and it's players

#### Request 

This request requires body payload, you can find the example below.

<details><summary>Show example payload</summary>
<p>

```json
{
  "name": "AC Milan 2",
  "description": "some-description",
  "location": "Italy",
  "players": [
    {
      "name": "John Doe 1",
      "nickname": "Lolo",
      "position": "forward"
    }
  ]
}
```
</p>
</details>

#### Response

The request will return the created data in JSON response like this:

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": {
    "id": "5f6a5d6129b2289c40b7444b",
    "name": "AC Milan 2",
    "description": "some-description",
    "location": "Italy",
    "players": [
      {
        "id": "5f6a5d6129b2289c40b74448",
        "name": "John Doe 1",
        "nickname": "Lolo",
        "position": "forward",
        "created_at": "2020-09-22T20:24:01.872Z"
      }
    ],
    "created_at": "2020-09-22T20:24:01.846Z"
  }
}
```

</p>
</details>

---

### `GET /players`

Returns list of players 

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": [
    {
      "id": "5f6a5c31d7c451c369802c02",
      "name": "John Doe 1",
      "nickname": "Lolo",
      "position": "forward",
      "created_at": "2020-09-22T20:18:57.957Z"
    }
  ]
}
```

</p>
</details>

---


### `GET /players/:id`

Returns a player by id

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": {
    "id": "5f6a5c31d7c451c369802c02",
    "name": "John Doe 1",
    "nickname": "Lolo",
    "position": "forward",
    "created_at": "2020-09-22T20:18:57.957Z"
  }
}
```

</p>
</details>

---