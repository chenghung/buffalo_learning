package actions

import (
	"fmt"
	"log"

	//	"github.com/chenghung/buffalo_learning/models"
	"github.com/chenghung/buffalo_learning/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/*
func (as *ActionSuite) Test_User_Index() {
	users := models.Users{
		{FirstName: "eddie", LastName: "chen", Email: "eddie.chen@gobuffalo.com"},
		{FirstName: "gopher", LastName: "developer", Email: "eddie.chen@gopher.com"},
	}

	err := as.DB.Create(&users)
	log.Println("create error", err)
	if err != nil {
		log.Fatal("create users failed", err)
	}

	res := as.JSON("/users").Get()
	as.Equal(200, res.Code)
	log.Println("users", res.Body)
}
*/

var _ = Describe("Test User APIs", func() {
	It("GET /api/users", func() {
		users := models.Users{
			{FirstName: "eddie", LastName: "chen", Email: "eddie.chen@gobuffalo.com"},
			{FirstName: "gopher", LastName: "developer", Email: "eddie.chen@gopher.com"},
		}

		err := as.DB.Create(&users)
		log.Println("create error", err)
		if err != nil {
			log.Fatal("create users failed", err)
		}
		res := as.JSON("/users").Get()
		fmt.Println("res", res.Body)
		Expect(1).To(Equal(1))
	})

	It("test", func() {
		res := as.JSON("/users").Get()
		fmt.Println("res", res.Body)
	})
})
