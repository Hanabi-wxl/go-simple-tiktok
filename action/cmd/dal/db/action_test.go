package db

import "testing"

func TestCreate(t *testing.T) {
	Init()
	CreateFavorite(1, 2)
}
