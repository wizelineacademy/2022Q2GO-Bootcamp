package pokemon

import (
	"testing"
)

//TODO: add tests for the service
func Test_FindByID(t *testing.T) {
	/*
		s := NewService()

		_, e := s.FindByID("1")
		fmt.Println(e)
	*/
}

type pokemonTask struct {
	name          string
	pokemonInfo   []string
	taskProcessor func([]string)
}

func (t pokemonTask) Run() {
	t.taskProcessor(t.pokemonInfo)
}

func Test_workerPool(t *testing.T) {
	/*
		pool := NewWorkerPool("")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		pokemons := [][]string{
			{"99", "kingler"},
			{"113", "chansey"},
			{"17", "pidgeotto"},
			{"7", "squirtle"},
			{"2", "ivysaur"},
			{"89", "muk"},
			{"67", "machoke"},
			{"94", "gengar"},
			{"92", "gastly"},
			{"94", "gengar"},
			{"93", "haunter"},
			{"666", "vivillon"},
			{"3", "venusaur"},
			{"17", "pidgeotto"},
			{"101", "electrode"},
			{"100", "voltorb"},
			{"301", "delcatty"},
			{"201", "unown"},
			{"150", "mewtwo"}}

		var tasks []Work

		simpleTask := func(ctx context.Context, pokemon Pokemon) (Pokemon, error) {
			fmt.Println("Validating ", pokemon)
			return pokemon, nil
		}

		for _, p := range pokemons {
			tasks = append(tasks, Work{
				fn:      simpleTask,
				pokemon: p,
			})
		}
		go pool.AddWork(tasks)

		go pool.Run(ctx)

		pkmns := []Pokemon{}
		for {
			select {
			case r, ok := <-pool.result:
				if !ok {
					continue
				}
				fmt.Printf("%#v\n", r.Value)
				id, _ := strconv.Atoi(r.Value[0])
				pkmns = append(pkmns, Pokemon{
					ID:   id,
					Name: r.Value[1],
				})
				if len(pkmns) == 3 {
					return
				}
			case <-pool.done:
				return
			}
		}

		fmt.Printf("%#v\n", pkmns)
	*/
}
