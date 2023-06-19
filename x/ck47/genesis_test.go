package ck47_test

import (
	"testing"

	keepertest "ck47/testutil/keeper"
	"ck47/testutil/nullify"
	"ck47/x/ck47"
	"ck47/x/ck47/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Ck47Keeper(t)
	ck47.InitGenesis(ctx, *k, genesisState)
	got := ck47.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
