package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/PaddyMc/osmosis/v8/app/apptesting"
	"github.com/PaddyMc/osmosis/v8/x/incentives/keeper"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	querier keeper.Querier
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()

	suite.querier = keeper.NewQuerier(*suite.App.IncentivesKeeper)

	lockableDurations := suite.App.IncentivesKeeper.GetLockableDurations(suite.Ctx)
	lockableDurations = append(lockableDurations, 2*time.Second)
	suite.App.IncentivesKeeper.SetLockableDurations(suite.Ctx, lockableDurations)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
