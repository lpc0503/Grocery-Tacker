package graph

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"math/rand"

	"github.com/lpc0503/Grocery-Tracker/graph/model"
	"github.com/stretchr/testify/assert"
)

var mockUsers = []*model.User{
	{UserID: "user1"},
	{UserID: "user2"},
	{UserID: "user3"},
	{UserID: "user4"},
}

var mockGroceryItems = []*model.GroceryItem{
	{UserID: "user1", Name: "Milk"},
	{UserID: "user2", Name: "Hotdog"},
	{UserID: "user1", Name: "Cat"},
	{UserID: "user3", Name: "Water"},
	{UserID: "user1", Name: "Beer"},
	{UserID: "user2", Name: "Apple"},
}

func createTestResovler() *Resolver {
	return &Resolver{
		users:        mockUsers,
		groceryItems: make(map[string][]*model.GroceryItem),
	}
}

func TestRegisterUser(t *testing.T) {
	resolver := &Resolver{}

	for index, user := range mockUsers {

		registerUser, err := resolver.Mutation().RegisterUser(context.TODO(), user.UserID)
		assert.NoError(t, err)
		assert.NotNil(t, registerUser)
		assert.Equal(t, mockUsers[index].UserID, registerUser.UserID)
	}
}

func TestLoginUser(t *testing.T) {
	resolver := createTestResovler()

	for index, user := range mockUsers {

		loginUser, err := resolver.Mutation().LoginUser(context.TODO(), user.UserID)
		assert.NoError(t, err)
		assert.NotNil(t, loginUser)
		assert.Equal(t, mockUsers[index].UserID, loginUser.UserID)
	}
}

func TestAddGroceryItem(t *testing.T) {
	resolver := createTestResovler()

	for index, item := range mockGroceryItems {

		quantity := rand.Intn(100)
		addItem, err := resolver.Mutation().AddUserGroceryItem(context.TODO(), item.UserID, item.Name, &quantity, nil, nil, nil, nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, addItem)
		assert.Equal(t, mockGroceryItems[index].Name, addItem.Name)
		assert.Equal(t, quantity, *addItem.Quantity)
	}
}

func TestDeleteGroceryItem(t *testing.T) {
	resolver := createTestResovler()

	// Add item first
	for index, item := range mockGroceryItems {

		quantity := rand.Intn(100)
		addItem, err := resolver.Mutation().AddUserGroceryItem(context.TODO(), item.UserID, item.Name, &quantity, nil, nil, nil, nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, addItem)
		assert.NotEmpty(t, addItem.ID)
		assert.Equal(t, mockGroceryItems[index].Name, addItem.Name)
		assert.Equal(t, quantity, *addItem.Quantity)
	}

	randUser := rand.Intn(len(mockUsers))
	userID := mockUsers[randUser].UserID
	deleteItem := strconv.Itoa(rand.Intn(len(resolver.groceryItems[userID]) + 1))

	fmt.Println(randUser)
	fmt.Println(userID)
	fmt.Println(deleteItem)

	fmt.Println("-------------")
	// Now delete the item
	success, err := resolver.Mutation().DeleteUserGroceryItem(context.TODO(), userID, deleteItem)
	assert.NoError(t, err)
	assert.True(t, success)

	// Try deleting again, should fail
	success, err = resolver.Mutation().DeleteUserGroceryItem(context.TODO(), userID, deleteItem)
	assert.Error(t, err)
	assert.False(t, success)
}
