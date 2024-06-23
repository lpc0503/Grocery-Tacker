package graph

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	resolver := &Resolver{}

	user, err := resolver.Mutation().RegisterUser(context.TODO(), "Test123")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Test123", user.UserID)
}

func TestLoginUser(t *testing.T) {
	resolver := &Resolver{}
	_, _ = resolver.Mutation().RegisterUser(context.TODO(), "Test234")

	user, err := resolver.Mutation().LoginUser(context.TODO(), "Test234")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Test234", user.UserID)
}

func TestAddGroceryItem(t *testing.T) {
	resolver := &Resolver{}

	var quantity = 2
	var q = &quantity
	item, err := resolver.Mutation().AddGroceryItem(context.TODO(), "Milk", q, nil, nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, "Milk", item.Name)
	assert.Equal(t, 2, *item.Quantity)
}

func TestDeleteGroceryItem(t *testing.T) {
	resolver := &Resolver{}

	// Add item first
	var quantity = 2
	var q = &quantity
	item, _ := resolver.Mutation().AddGroceryItem(context.TODO(), "Coke", q, nil, nil, nil, nil, nil)

	// Now delete the item
	success, err := resolver.Mutation().DeleteGroceryItem(context.TODO(), item.ID)
	assert.NoError(t, err)
	assert.True(t, success)

	// Try deleting again, should fail
	success, err = resolver.Mutation().DeleteGroceryItem(context.TODO(), item.ID)
	assert.Error(t, err)
	assert.False(t, success)
}
