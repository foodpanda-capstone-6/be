package engine

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type engineSuite struct {
	suite.Suite
	Log *GlobalLog

	countTest int
}

func (suite *engineSuite) SetupTest() {

	var engineOpt = &EngineOpts{
		LogPath: "logs/engine-test.txt",
	}

	suite.Log = InitGlobalLog(engineOpt)

	suite.countTest++
}

func (suite *engineSuite) TestSmth() {
	log.Println(os.Getwd())
}

func TestHook(t *testing.T) {

	thisSuite := new(engineSuite)

	suite.Run(t, thisSuite)
}
