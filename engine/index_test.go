package engine

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	"vms-be/infra/database"

	"github.com/stretchr/testify/suite"
)

type engineSuite struct {
	suite.Suite
	engine *Engine

	countTest int
}

func (suite *engineSuite) SetupTest() {

	var engineOpt = &EngineOpts{
		LogPath: "logs/engine-test.txt",
		DatabaseOpts: &database.DatabaseOpts{DriverName: "sqlite3", DatabaseOpts_SQL: database.DatabaseOpts_SQL{
			Path: fmt.Sprintf("%s%s-%03d%s", "storage/main", time.Now().Format("2006_01_02"), suite.countTest, ".db"),
		}},
	}

	suite.engine = InitEngine(engineOpt)

	suite.countTest++
}

func (suite *engineSuite) TestSmth() {
	log.Println(os.Getwd())
}

func TestHook(t *testing.T) {

	thisSuite := new(engineSuite)

	suite.Run(t, thisSuite)
}
