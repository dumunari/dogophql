package graph

import "github.com/dumunari/gophql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Dogs    []*model.Dog
	Breeds  []*model.Breed
	Puppies []*model.Puppy
}