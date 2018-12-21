package memconfig

import (
	"testing"

	"github.com/simelo/rextporter/src/core"
	"github.com/simelo/rextporter/src/core/mocks"
	"github.com/stretchr/testify/suite"
)

func newLabelDef(suite *labelDefConfSuit) core.RextLabelDef {
	return NewLabelDef(
		suite.name,
		suite.nodeSolver,
	)
}

type labelDefConfSuit struct {
	suite.Suite
	labelDef   core.RextLabelDef
	name       string
	nodeSolver core.RextNodeSolver
}

func (suite *labelDefConfSuit) SetupTest() {
	suite.name = "MySupperLabel"
	suite.nodeSolver = NewNodeSolver("tr", "pat", NewOptionsMap())
	suite.labelDef = newLabelDef(suite)
}

func TestLabelDefConfSuit(t *testing.T) {
	suite.Run(t, new(labelDefConfSuit))
}

func (suite *labelDefConfSuit) TestNewLabelDef() {
	// NOTE(denisacostaq@gmail.com): Giving

	// NOTE(denisacostaq@gmail.com): When
	labelDef := newLabelDef(suite)

	// NOTE(denisacostaq@gmail.com): Assert
	suite.Equal(suite.name, labelDef.GetName())
	suite.Equal(suite.nodeSolver, labelDef.GetNodeSolver())
}

func (suite *labelDefConfSuit) TestAbleToSetName() {
	// NOTE(denisacostaq@gmail.com): Giving
	orgName := suite.labelDef.GetName()
	name := "fgfg78"
	suite.labelDef.SetName(name)

	// NOTE(denisacostaq@gmail.com): When
	name2 := suite.labelDef.GetName()

	// NOTE(denisacostaq@gmail.com): Assert
	suite.Equal(name, name2)
	suite.NotEqual(orgName, name2)
}

func (suite *labelDefConfSuit) TestAbleToSetNodeSolver() {
	// NOTE(denisacostaq@gmail.com): Giving
	orgNs := suite.labelDef.GetNodeSolver()
	ns := &NodeSolver{MType: "fee"}
	suite.labelDef.SetNodeSolver(ns)

	// NOTE(denisacostaq@gmail.com): When
	ns2 := suite.labelDef.GetNodeSolver()

	// NOTE(denisacostaq@gmail.com): Assert
	suite.Equal(ns, ns2)
	suite.NotEqual(orgNs, ns2)
}

func (suite *labelDefConfSuit) TestValidationClonedShouldBeValid() {
	// NOTE(denisacostaq@gmail.com): Giving

	// NOTE(denisacostaq@gmail.com): When
	labelDef, err := suite.labelDef.Clone()
	suite.Nil(err)
	suite.Equal(suite.labelDef, labelDef)
	hasError := labelDef.Validate()

	// NOTE(denisacostaq@gmail.com): Assert
	suite.False(hasError)
}

func (suite *labelDefConfSuit) TestValidationNameShouldNotBeEmpty() {
	// NOTE(denisacostaq@gmail.com): Giving
	labelDef, err := suite.labelDef.Clone()
	suite.Nil(err)
	setUpFakeValidationOn3rdPartyOverLabel(labelDef)

	// NOTE(denisacostaq@gmail.com): When
	labelDef.SetName("")
	hasError := labelDef.Validate()

	// NOTE(denisacostaq@gmail.com): Assert
	suite.True(hasError)
}

func (suite *labelDefConfSuit) TestValidationNodeSolverShouldNotBeEmpty() {
	// NOTE(denisacostaq@gmail.com): Giving
	labelDef, err := suite.labelDef.Clone()
	suite.Nil(err)

	// NOTE(denisacostaq@gmail.com): When
	labelDef.SetNodeSolver(nil)
	hasError := labelDef.Validate()

	// NOTE(denisacostaq@gmail.com): Assert
	suite.True(hasError)
}

func (suite *labelDefConfSuit) TestValidationShouldGoDownTroughFields() {
	// NOTE(denisacostaq@gmail.com): Giving
	cLabelConf, err := suite.labelDef.Clone()
	suite.Nil(err)
	mockNodeSolver := new(mocks.RextNodeSolver)
	mockNodeSolver.On("Validate").Return(false)
	cLabelConf.SetNodeSolver(mockNodeSolver)

	// NOTE(denisacostaq@gmail.com): When
	cLabelConf.Validate()

	// NOTE(denisacostaq@gmail.com): Assert
	mockNodeSolver.AssertCalled(suite.T(), "Validate")
}

func setUpFakeValidationOn3rdPartyOverLabel(labelDef core.RextLabelDef) {
	nodeSolverStub := new(mocks.RextNodeSolver)
	nodeSolverStub.On("Validate").Return(false)
	labelDef.SetNodeSolver(nodeSolverStub)
}
