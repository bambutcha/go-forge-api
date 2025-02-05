package store_test

import (
	"testing"

	"github.com/bambutcha/http-rest-api/internal/app/model"
	"github.com/bambutcha/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

var (
	exampleEmail = "user@example.org"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := exampleEmail
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	createdUser, err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)

	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
