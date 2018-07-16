package actions

import (
	"fmt"
	"testing"

	"github.com/gobuffalo/suite"
	"github.com/markbates/willie"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type ActionSuite struct {
	*suite.Action
}

var as *ActionSuite

func init() {
	action := suite.NewAction(App())
	action.Willie = willie.New(action.App)
	as = &ActionSuite{action}
}

func Test_ActionSuite(t *testing.T) {
	fmt.Printf("init test....., %T, %T\n\n", as, &as)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Action Suite")
}

var _ = BeforeSuite(func() {
	fmt.Println("truncate all\n\n")
	as.DB.TruncateAll()
})
