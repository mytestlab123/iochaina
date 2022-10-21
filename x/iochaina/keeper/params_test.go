package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iochaina/testutil/keeper"
	"github.com/mytestlab123/iochaina/x/iochaina/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IochainaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
