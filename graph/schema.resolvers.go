package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"errors"
	"strconv"

	"github.com/lpc0503/Grocery-Tacker/graph/model"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, userID string) (*model.User, error) {

	for _, user := range r.users {
		if user.UserID == userID {
			return nil, errors.New("User exist")
		}
	}

	user := &model.User{
		UserID: userID,
	}
	r.users = append(r.users, user)
	return user, nil
}

// LoginUser is the resolver for the loginUser field.
func (r *mutationResolver) LoginUser(ctx context.Context, userID string) (*model.User, error) {
	for _, user := range r.users {
		if user.UserID == userID {
			return user, nil
		}
	}
	return nil, errors.New("invalid credentials")
}

// AddGroceryItem is the resolver for the addGroceryItem field.
func (r *mutationResolver) AddGroceryItem(ctx context.Context, name string, quantity *int, purchaseDate *string, expirationDate *string, price *float64, materials []*string, category *string) (*model.GroceryItem, error) {
	item := &model.GroceryItem{
		ID:             strconv.Itoa(len(r.groceryItems) + 1),
		Name:           name,
		Quantity:       quantity,
		PurchaseDate:   purchaseDate,
		ExpirationDate: expirationDate,
		Price:          price,
		Materials:      materials,
		Category:       category,
	}
	r.groceryItems = append(r.groceryItems, item)
	return item, nil
}

// UpdateGroceryItem is the resolver for the updateGroceryItem field.
func (r *mutationResolver) UpdateGroceryItem(ctx context.Context, id string, name *string, quantity *int, purchaseDate *string, expirationDate *string, price *float64, materials []*string, category *string) (*model.GroceryItem, error) {
	for _, item := range r.groceryItems {
		if item.ID == id {
			if name != nil {
				item.Name = *name
			}
			if quantity != nil {
				item.Quantity = quantity
			}
			if purchaseDate != nil {
				item.PurchaseDate = purchaseDate
			}
			if expirationDate != nil {
				item.ExpirationDate = expirationDate
			}
			if price != nil {
				item.Price = price
			}
			if materials != nil {
				item.Materials = materials
			}
			if category != nil {
				item.Category = category
			}
			return item, nil
		}
	}
	return nil, errors.New("item not found")
}

// DeleteGroceryItem is the resolver for the deleteGroceryItem field.
func (r *mutationResolver) DeleteGroceryItem(ctx context.Context, id string) (bool, error) {
	for i, item := range r.groceryItems {
		if item.ID == id {
			r.groceryItems = append(r.groceryItems[:i], r.groceryItems[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("item not found")
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	for _, user := range r.users {
		if user.UserID == id {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

// GetGroceryItem is the resolver for the getGroceryItem field.
func (r *queryResolver) GetGroceryItem(ctx context.Context, id string) (*model.GroceryItem, error) {
	for _, item := range r.groceryItems {
		if item.ID == id {
			return item, nil
		}
	}
	return nil, errors.New("item not found")
}

// GetGroceryItems is the resolver for the getGroceryItems field.
func (r *queryResolver) GetGroceryItems(ctx context.Context) ([]*model.GroceryItem, error) {
	return r.groceryItems, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }