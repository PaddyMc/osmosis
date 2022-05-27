package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/PaddyMc/osmosis/v8/app/apptesting"
	"github.com/PaddyMc/osmosis/v8/x/epochs/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
