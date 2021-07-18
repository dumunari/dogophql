package main

import (
	"github.com/dumunari/dogophql/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dumunari/dogophql/graph"
	"github.com/dumunari/dogophql/graph/generated"
)

const defaultPort = "8080"


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: setupInitialConfiguration()}))

	http.Handle("/", playground.Handler("GraphQL playground", "/dogophql"))
	http.Handle("/dogophql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func setupInitialConfiguration() *graph.Resolver {
	a := 1
	c := "Black"
	breed := &model.Breed{
		ID:   "T5577006791947779410",
		Name: "English bulldog",
	}
	p1 := &model.Dog{
		ID:      "T6129484611666145821",
		Name:    "Spike",
		Age:     &a,
		Breed:   breed,
		Puppies: nil,
	}

	p2 := &model.Dog{
		ID:      "T2775422040480279449",
		Name:    "Spika",
		Age:     &a,
		Breed:   breed,
		Puppies: nil,
	}
	puppy := &model.Puppy{
		ID:      "T4751997750760398084",
		Name:    "Tyke",
		Color:   &c,
		Breed:   breed,
		Parents: []*model.Dog{
			p1,
			p2,
		},
	}

	return &graph.Resolver{
		Dogs:    []*model.Dog{
			p1,
			p2,
		},
		Breeds:  []*model.Breed{
			breed,
		},
		Puppies: []*model.Puppy{
			puppy,
		},
	}
}

