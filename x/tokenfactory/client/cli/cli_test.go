package cli_test

import (
	"testing"

	"github.com/osmosis-labs/osmosis/osmoutils/osmocli"
	"github.com/tessornetwork/mokita/x/tokenfactory/client/cli"
	"github.com/tessornetwork/mokita/x/tokenfactory/types"
)

func TestGetCmdDenomAuthorityMetadata(t *testing.T) {
	desc, _ := cli.GetCmdDenomAuthorityMetadata()
	tcs := map[string]osmocli.QueryCliTestCase[*types.QueryDenomAuthorityMetadataRequest]{
		"basic test": {
			Cmd: "uatom",
			ExpectedQuery: &types.QueryDenomAuthorityMetadataRequest{
				Denom: "uatom",
			},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdDenomsFromCreator(t *testing.T) {
	desc, _ := cli.GetCmdDenomsFromCreator()
	tcs := map[string]osmocli.QueryCliTestCase[*types.QueryDenomsFromCreatorRequest]{
		"basic test": {
			Cmd: "moki1test",
			ExpectedQuery: &types.QueryDenomsFromCreatorRequest{
				Creator: "moki1test",
			},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}
