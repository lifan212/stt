package stt_test

import (
	"testing"

	keepertest "stt/testutil/keeper"
	"stt/testutil/nullify"
	stt "stt/x/stt/module"
	"stt/x/stt/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SttKeeper(t)
	stt.InitGenesis(ctx, k, genesisState)
	got := stt.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
