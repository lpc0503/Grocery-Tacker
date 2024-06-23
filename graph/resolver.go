package graph

import "github.com/lpc0503/Grocery-Tracker/graph/model"

//go:generate go run -mod=mod github.com/99designs/gqlgen generate --verbose

type Resolver struct {
	users        []*model.User
	groceryItems []*model.GroceryItem
}
