package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"net/http"
    "github.com/gin-gonic/gin"
)

type pokemon struct {
	ID			int		`json:"id"`
	NAME 		string	`json:"name"`
	TYPE1 		string	`json:"type1"`
	TYPE2 		string	`json:"type2"`
	TOTAL		int		`json:"total"`
	HP			int		`json:"hp"`
	ATTACK		int		`json:"attack"`
	DEFENSE		int		`json:"defendse"`
	SPATK		int		`json:"sp_atk"`
	SPDEF		int		`json:"sp_def"`
	SPEED		int		`json:"speed"`
	GENERATION	int		`json:"generation"`
	LEGENDARY	bool	`json:"legendary"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func parseFile(file string) []pokemon {

	pokemons := make([]pokemon,0)
	
	f, err := os.Open(file)
    check(err)

	fileScanner := bufio.NewScanner(f)
 
    fileScanner.Split(bufio.ScanLines)
  
	index := 0
    for fileScanner.Scan() {
		index++
		if index > 1 {
			line := fileScanner.Text()
			myPokemon := textLineToPokemon(line)
			pokemons = append(pokemons, myPokemon)
		}
    }
  
    f.Close()
	return pokemons
}

func textLineToPokemon(line string) pokemon{

	values := strings.Split(line, ",")
	id, err := strconv.Atoi(values[0])
	check(err)

	name := values[1]
	type1 := values[2]
	type2 := values[3]

	total, err := strconv.Atoi(values[4])
	check(err)

	hp, err := strconv.Atoi(values[5])
	check(err)

	attack, err := strconv.Atoi(values[6])
	check(err)

	defense, err := strconv.Atoi(values[7])
	check(err)

	spAtk, err := strconv.Atoi(values[8])
	check(err)

	spDef, err := strconv.Atoi(values[9])
	check(err)

	speed, err := strconv.Atoi(values[10])
	check(err)

	generation, err := strconv.Atoi(values[11])
	check(err)

	legendary, err := strconv.ParseBool(values[12])
	check(err)

	myPokemon := pokemon {
		ID: id,
		NAME: name,
		TYPE1: type1,
		TYPE2: type2,
		TOTAL: total,
		HP: hp,
		ATTACK: attack,
		DEFENSE: defense,
		SPATK: spAtk,
		SPDEF: spDef,
		SPEED: speed,
		GENERATION: generation,
		LEGENDARY: legendary,
	}

	return myPokemon
}

var pokemons = parseFile("data/pokemon.csv")

func getPokemons(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, pokemons)
}

func getPokemonById(c *gin.Context) {
    idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	check(err)
	for _, p := range pokemons {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pokemon not found"})
}

func getLegendaryPokemons(c *gin.Context) {
	legendaryPokemons := make([]pokemon,0)
	for _, p := range pokemons {
		if p.LEGENDARY {
			legendaryPokemons = append(legendaryPokemons, p)
		}
	}
	if len(legendaryPokemons) > 0 {
		c.IndentedJSON(http.StatusOK, legendaryPokemons)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Legendary Pokemons found"})
	}
}

func main(){
	router := gin.Default()
    router.GET("/pokemons", getPokemons)
	router.GET("/pokemons/:id", getPokemonById)
	router.GET("/pokemons/legendary", getLegendaryPokemons)
    router.Run("localhost:8080")
}