package balancer_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/PaddyMc/osmosis/v8/app/apptesting"
	"github.com/PaddyMc/osmosis/v8/x/gamm/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
}
