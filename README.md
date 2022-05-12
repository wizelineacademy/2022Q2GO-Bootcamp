# 2022 Q2 Go Bootcamp

## First Deliverable


What is delivered:

- The CSV file used is about Pokemons and contains multiple fields including IDs
  - It obtained from [here](http://www.metabates.com/2015/10/15/handling-http-request-errors-in-go/)
- An endpoint that shows all Pokemons
  - /pokemons
- An endpoint that shows a pokemon by its ID
  - /pokemons/:id
- An endpoint that shows only legendary pokemons
  - /pokemons/legendary

Read the API Definition section for more details


&nbsp;

## API Definition

&nbsp;

**GET** ***/pokemons***

Displays *Hello World!*

Sample Query: 
`http://localhost:8080/pokemons`

Response

    [
    {
        "id": 1,
        "name": "Bulbasaur",
        "type1": "Grass",
        "type2": "Poison",
        "total": 318,
        "hp": 45,
        "attack": 49,
        "defendse": 49,
        "sp_atk": 65,
        "sp_def": 65,
        "speed": 45,
        "generation": 1,
        "legendary": false
    },
    {
        "id": 2,
        "name": "Ivysaur",
        .
        .
        .

&nbsp;

**GET** ***/pokemons/:id***

Get a random activity for one participant

Sample Query: 
`http://localhost:8080/pokemons/25`

Response

    {
      "id": 25,
      "name": "Pikachu",
      "type1": "Electric",
      "type2": "",
      "total": 320,
      "hp": 35,
      "attack": 55,
      "defendse": 40,
      "sp_atk": 50,
      "sp_def": 50,
      "speed": 90,
      "generation": 1,
      "legendary": false
    }


&nbsp;

**GET** ***/pokemons/legendary***

Get a random activity for one participant

Sample Query: 
`http://localhost:8080/pokemons/legendary`

Response

    [
    {
        "id": 144,
        "name": "Articuno",
        "type1": "Ice",
        "type2": "Flying",
        "total": 580,
        "hp": 90,
        "attack": 85,
        "defendse": 100,
        "sp_atk": 95,
        "sp_def": 125,
        "speed": 85,
        "generation": 1,
        "legendary": true
    },
    {
        "id": 145,
        "name": "Zapdos",
        "type1": "Electric",
        "type2": "Flying",
        "total": 580,
        "hp": 90,
        "attack": 90,
        "defendse": 85,
        "sp_atk": 125,
        "sp_def": 90,
        "speed": 100,
        "generation": 1,
        "legendary": true
        .
        .
        .


&nbsp;


## Running the application

&nbsp;
From the project root directory run:

    go run main.go  

&nbsp;

The service will run on  on localhost:8080

&nbsp;