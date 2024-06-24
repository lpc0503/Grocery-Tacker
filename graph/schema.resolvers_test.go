package graph

import (
	"context"
	"testing"

	"github.com/lpc0503/Grocery-Tracker/graph/model"
	"github.com/stretchr/testify/assert"
)

var mockUsers = []*model.User{
	{UserID: "user1"},
	{UserID: "user2"},
	{UserID: "user3"},
}

func createTestResovler() *Resolver {
	return &Resolver{
		users: mockUsers,
	}
}

func TestRegisterUser(t *testing.T) {
	resolver := &Resolver{}

	for index, user := range mockUsers {

		registerUser, err := resolver.Mutation().RegisterUser(context.TODO(), user.UserID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, mockUsers[index].UserID, registerUser.UserID)
	}
}

func TestLoginUser(t *testing.T) {
	resolver := createTestResovler()

	for index, user := range mockUsers {

		loginUser, err := resolver.Mutation().LoginUser(context.TODO(), user.UserID)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, mockUsers[index].UserID, loginUser.UserID)
	}
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
