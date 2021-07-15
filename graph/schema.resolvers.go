package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/dumunari/gophql/graph/generated"
	"github.com/dumunari/gophql/graph/model"
)

func (r *mutationResolver) CreateDog(ctx context.Context, input model.NewDog) (*model.Dog, error) {
	var breed *model.Breed

	for _, b := range r.Resolver.Breeds {
		if b.ID == input.BreedID {
			breed = b
		}
	}

	dog := &model.Dog{
		ID:      fmt.Sprintf("T%d", rand.Int()),
		Name:    input.Name,
		Age:     input.Age,
		Breed:   breed,
		Puppies: nil,
	}

	r.Resolver.Dogs = append(r.Resolver.Dogs, dog)

	return dog, nil
}

func (r *mutationResolver) CreateBreed(ctx context.Context, input model.NewBreed) (*model.Breed, error) {
	breed := &model.Breed{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: input.Name,
	}

	r.Resolver.Breeds = append(r.Resolver.Breeds, breed)

	return breed, nil
}

func (r *mutationResolver) CreatePuppy(ctx context.Context, input model.NewPuppy) (*model.Puppy, error) {
	var parents []*model.Dog
	var breed *model.Breed

	for _, b := range r.Resolver.Breeds {
		if b.ID == input.BreedID {
			breed = b
		}
	}

	for _, d := range r.Resolver.Dogs {
		if d.ID == input.ParentsID[0] || d.ID == input.ParentsID[1] {
			parents = append(parents, d)
		}
	}

	puppy := &model.Puppy{
		ID:      fmt.Sprintf("T%d", rand.Int()),
		Name:    input.Name,
		Color:   input.Color,
		Breed:   breed,
		Parents: parents,
	}

	r.Resolver.Puppies = append(r.Resolver.Puppies, puppy)

	return puppy, nil
}

func (r *mutationResolver) AddPuppyToParents(ctx context.Context, input model.AddPuppy) ([]*model.Dog, error) {
	var puppy model.Puppy
	var dogs []*model.Dog

	for _, p := range r.Resolver.Puppies {
		if p.ID == input.PuppyID {
			puppy = *p
		}
	}

	for _, d := range r.Resolver.Dogs {
		if d.ID == input.ParentsID[0] || d.ID == input.ParentsID[1] {
			dogs = append(dogs, d)
			d.Puppies = append(d.Puppies, &puppy)
		}
	}

	return dogs, nil
}

func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	return r.Resolver.Dogs, nil
}

func (r *queryResolver) Breeds(ctx context.Context) ([]*model.Breed, error) {
	return r.Resolver.Breeds, nil
}

func (r *queryResolver) Puppies(ctx context.Context) ([]*model.Puppy, error) {
	return r.Resolver.Puppies, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
