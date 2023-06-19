package keeper_test

import (
	"testing"

	testkeeper "ck47/testutil/keeper"
	"ck47/x/ck47/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Ck47Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
