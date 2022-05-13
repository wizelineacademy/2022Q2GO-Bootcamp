# Capstone Project - Manuel Zavala - V0.0.1

## First deriverable instructions
Perform the following:
* Create an API
* Add an endpoint to read from a CSV file, it should get information from the CSV by some field.
* The CSV should have any information, the items in the CSV must have an ID element (int value), for example:
```csv
1,bulbasaur
2,ivysaur
3,venusaur
``` 
* The result should be displayed as a response

You should follow:
* Clean architecture
* Use best practices
* Handle the Errors 
### Project structure
```bash
├── README.md
└── api
    ├── app
    │   └── app.go
    ├── controllers
    │   └── controllers.go
    ├── data
    │   └── pokemon.csv
    ├── dataclass
    │   └── dataclass.go
    ├── go.mod
    ├── main.go
    ├── models
    │   └── models.go
    ├── repository
    │   └── repository.go
    └── services
        └── services.go
```
### How to run
```bash
go run main.go
```
### First endpoint
This endpoint will retrieve a pokemon list. 
* If it receives a quantity, then it will return that quantity of pokemon. 
```bash
curl --location --request GET 'http://localhost:1000/quantity/1'
```
```json
[
    {
        "id": 1,
        "name": "bulbasaur"
    }
]
```
* If quantity is more than existing pokemon size, then it will return an error. 
```bash
curl --location --request GET 'http://localhost:1000/quantity/10'
```
```json
{
    "code": 422,
    "description": "Max quantity"
}
```
* If it does not receive a quantity, then it will return all existing pokemon.
```bash
curl --location --request GET 'http://localhost:1000/quantity/'
```
```json
[
    {
        "id": 1,
        "name": "bulbasaur"
    },
    {
        "id": 2,
        "name": "ivysaur"
    },
    {
        "id": 3,
        "name": "venusaur"
    }
]
```
* If there is an error procesing the csv data, then it will return an error.
```bash
curl --location --request GET 'http://localhost:1000/quantity/'
```
```json
{
    "code": 404,
    "description": "open ./data/pokemon1.csv: no such file or directory"
}
```
### Second endpoint
This endpoint will return a pokemon by its id
```bash
curl --location --request GET 'http://localhost:1000/id/1'
```
```json
{
    "id": 1,
    "name": "bulbasaur"
}
```
* Id param should be a number
```bash
curl --location --request GET 'http://localhost:1000/id/x'
```
```json
{
    "code": 422,
    "description": "Bad format :id should be a number"
}
```
* There is no pokemon with that id
```bash
curl --location --request GET 'http://localhost:1000/id/10'
```
```json
{
    "code": 404,
    "description": "There is no pokemon with id"
}
```
* If there is an error procesing the csv data, then it will return an error.
```bash
curl --location --request GET 'http://localhost:1000/id/'
```
```json
{
    "code": 404,
    "description": "open ./data/pokemon1.csv: no such file or directory"
}
```
### Third endpoint
This endpoint will return a pokemon by its name
```bash
curl --location --request GET 'http://localhost:1000/name/bulbasaur'
```
```json
{
    "id": 1,
    "name": "bulbasaur"
}
```
* Name param should be a string
```bash
curl --location --request GET 'http://localhost:1000/name/'
```
```json
{
    "code": 422,
    "description": "Bad format :name should be a string"
}
```
* There is no pokemon with that name
```bash
curl --location --request GET 'http://localhost:1000/name/10'
```
```json
{
    "code": 404,
    "description": "There is no pokemon with name"
}
```
* If there is an error procesing the csv data, then it will return an error.
```bash
curl --location --request GET 'http://localhost:1000/name/'
```
```json
{
    "code": 404,
    "description": "open ./data/pokemon1.csv: no such file or directory"
}
```
