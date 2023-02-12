package jwtauth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"io.github.tuanictu97/api/pkg/x/cachex"
)

func TestAuth(t *testing.T) {
	cache := cachex.NewBadgerCache(cachex.BadgerConfig{
		Path: "./tmp/jwt",
	})

	store := NewStoreWithCache(cache)
	ctx := context.Background()
	jwtAuth := New(store)

	userID := "test"
	token, err := jwtAuth.GenerateToken(ctx, userID)
	assert.Nil(t, err)
	assert.NotNil(t, token)

	id, err := jwtAuth.ParseSubject(ctx, token.GetAccessToken())
	assert.Nil(t, err)
	assert.Equal(t, userID, id)

	err = jwtAuth.DestroyToken(ctx, token.GetAccessToken())
	assert.Nil(t, err)

	id, err = jwtAuth.ParseSubject(ctx, token.GetAccessToken())
	assert.NotNil(t, err)
	assert.EqualError(t, err, ErrInvalidToken.Error())
	assert.Empty(t, id)

	err = jwtAuth.Release(ctx)
	assert.Nil(t, err)
}
