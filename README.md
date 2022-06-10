# Set up

1. Be sure to have all the dependencies:
```bash
go mod tidy
```
2. Run the application:
```bash
go run main.go
```

## Endpoints:
By default, the application will listen on port 4200.

| Method | Path                                                                    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
|--------|-------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------| 
| GET    | /boardgame/{boardgame_id}                                               | Get info of the boardgame, if not exist in our csv file it will return 404 error, the *boaradgame_id* should be an int                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| GET    | /pokemon/{pokemon_id}                                                   | Get info of the pokemon, if not exist in the pokemonAPI it will return 404 error, the *pokemon_id* could be the name of the pokemon or the pokedex number                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | 
| GET    | /pokemon?type={type}&items={items}&items_per_workers={items_per_worker} | - _type_: should be **odd** or **even** and it's mandatory. This means if the ID should be odd or even<br/>- _items_: is optional, the default value is 5. This means the number of items that you want to return<br/> - _items_per_worker_: is optional, default value is 1. This means the number of success items each worker will process. Note The API is running with 8 workers<br/><br/>The Endpoint will return whatever of the conditions that occur first:<br/>- If the EOF is reached, return the items that were processed<br/>- If the limit of items is reached<br/>- If each worker reached the limit of items per worker | 