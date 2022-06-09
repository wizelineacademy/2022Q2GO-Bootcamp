package pokemon

import (
	"context"
	"strconv"
	"testing"
)

var (
	pokemons = [][]string{
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
		{"150", "mewtwo"},
	}
)

func TestWorkerPool_Odd(t *testing.T) {
	pool := NewWorkerPool("odd")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	itemsWorker := 20
	items := 10
	expectedItems := 12

	var tasks []Work
	for _, p := range pokemons {
		id, _ := strconv.Atoi(p[0])
		tasks = append(tasks, Work{
			pokemon: Pokemon{
				ID:   id,
				Name: p[1],
			},
		})
	}

	go pool.AddWork(tasks)
	go pool.Run(ctx, itemsWorker)
	pkmns := []Pokemon{}
	for {
		select {
		case r, ok := <-pool.result:
			if !ok {
				continue
			}
			pkmns = append(pkmns, r.Value)
			if len(pkmns) == items {
				return
			}
		case <-pool.done:
			return
		}
	}

	if len(pkmns) != expectedItems {
		t.Errorf("expected %d items, got %d", expectedItems, len(pkmns))
	}
}

func Test_workerPool_foo(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var tasks []Work
	for _, pokmn := range pokemons {
		id, _ := strconv.Atoi(pokmn[0])
		tasks = append(tasks, Work{
			pokemon: Pokemon{
				ID:   id,
				Name: pokmn[1],
			},
		})
	}

	tests := map[string]struct {
		tpe           string
		items         int
		itemsWorker   int
		expectedItems int
	}{
		"get odd pokemons": {
			tpe:           "odd",
			items:         100,
			itemsWorker:   5,
			expectedItems: 12,
		},
		"get even pokemons": {
			tpe:           "even",
			items:         100,
			itemsWorker:   5,
			expectedItems: 7,
		},
		"reach limit workers": {
			tpe:           "odd",
			items:         10,
			itemsWorker:   1,
			expectedItems: 8,
		},
		"get total items": {
			tpe:           "odd",
			items:         5,
			itemsWorker:   3,
			expectedItems: 5,
		},
		"reach EOF": {
			tpe:           "even",
			items:         150,
			itemsWorker:   20,
			expectedItems: 7,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pool := NewWorkerPool(test.tpe)
			go pool.AddWork(tasks)
			go pool.Run(ctx, test.itemsWorker)
			pkmns := []Pokemon{}
			for {
				select {
				case r, ok := <-pool.result:
					if !ok {
						continue
					}
					pkmns = append(pkmns, r.Value)
					if len(pkmns) == test.items {
						if len(pkmns) != test.expectedItems {
							t.Errorf("expected %d items, got %d", test.expectedItems, len(pkmns))
						}
						return
					}
				case <-pool.done:
					if len(pkmns) != test.expectedItems {
						t.Errorf("expected %d items, got %d", test.expectedItems, len(pkmns))
					}
					return
				}
			}
		})
	}
}
