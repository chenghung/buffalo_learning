package actions

import (
	"log"

	"github.com/chenghung/buffalo_learning/models"
)

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

// func (as *ActionSuite) Test_User_Show() {
// 	as.Fail("Not Implemented!")
// }

// func (as *ActionSuite) Test_User_Create() {
// 	as.Fail("Not Implemented!")
// }

// func (as *ActionSuite) Test_User_Update() {
// 	as.Fail("Not Implemented!")
// }

// func (as *ActionSuite) Test_User_Destroy() {
// 	as.Fail("Not Implemented!")
// }
