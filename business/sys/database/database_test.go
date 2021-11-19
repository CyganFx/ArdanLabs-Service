package database

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	log2 "log"
	"testing"
	"unsafe"
)

type testDest struct {
}

func TestOpen(t *testing.T) {
	// check when no container is running
	badCfg := Config{
		User:       "asd",
		Password:   "asd",
		Host:       "1235124123123",
		Name:       "random name",
		DisableTLS: true,
	}

	db, err := Open(badCfg)
	require.NoError(t, err)
	defer db.Close()

	err = db.Ping()
	require.Error(t, err)
}

func TestNamedQuerySlice(t *testing.T) {
	// check when no container is running
	cfg := Config{
		User:       "postgres",
		Password:   "postgres",
		Host:       "0.0.0.0",
		Name:       "postgres",
		DisableTLS: true,
	}

	db, err := Open(cfg)
	require.NoError(t, err)
	defer db.Close()

	ctx := context.Background()
	l := log2.Logger{}
	query := ""

	data := struct {
		Offset      int `db:"offset"`
		RowsPerPage int `db:"rows_per_page"`
	}{
		Offset:      1,
		RowsPerPage: 10,
	}

	var dest1 unsafe.Pointer
	err = NamedQuerySlice(ctx, &l, db, query, data, dest1)
	assert.Error(t, err)

	var dest2 []testDest
	err = NamedQuerySlice(ctx, &l, db, query, data, &dest2)
	assert.Error(t, err)
}
