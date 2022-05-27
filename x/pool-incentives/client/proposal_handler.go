package client

import (
	"github.com/PaddyMc/osmosis/v8/x/pool-incentives/client/cli"
	"github.com/PaddyMc/osmosis/v8/x/pool-incentives/client/rest"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var UpdatePoolIncentivesHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpdatePoolIncentivesProposal, rest.ProposalUpdatePoolIncentivesRESTHandler)
