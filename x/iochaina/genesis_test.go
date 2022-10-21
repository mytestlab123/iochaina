package iochaina_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iochaina/testutil/keeper"
	"github.com/mytestlab123/iochaina/testutil/nullify"
	"github.com/mytestlab123/iochaina/x/iochaina"
	"github.com/mytestlab123/iochaina/x/iochaina/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IochainaKeeper(t)
	iochaina.InitGenesis(ctx, *k, genesisState)
	got := iochaina.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
