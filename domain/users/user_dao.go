package users

import (
	"fmt"
	"main/datasources/mysql/users_db"
	"main/utils/date_utils"
	"main/utils/errors"
	"strings"
)
const(
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser = "INSERT INTO users(first_name , last_name , email , date_created) VALUES(?,?,?,?);"
)
var(
	usersDB = make(map[int64]*User)
)
func (user *User) Get() *errors.RestErr{

	if err := users_db.Client.Ping(); err != nil{
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found" , user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil 
}

func (user *User) Save() *errors.RestErr {
	statement , err := users_db.Client.Prepare(queryInsertUser)
	if err != nil{
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()
	user.DateCreated =  date_utils.GetNowString()
	insertResult , err := statement.Exec(user.FirstName , user.LastName , user.Email , user.DateCreated)
	if err != nil{
		if strings.Contains(err.Error() , indexUniqueEmail){
			return errors.NewBadRequestError(fmt.Sprintf("email %s has already exists" , user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s" ,err.Error() ))
	}
	userId , err := insertResult.LastInsertId()
	if err!= nil{
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s" ,err.Error() ))
	}
	user.Id = userId
	return nil

	//current := usersDB[user.Id]
	//if current != nil {
	//	if current.Email == user.Email{
	//		return errors.NewBadRequestError("fuck")
	//	}
	//	return errors.NewBadRequestError("already exsit")
	//}
	//usersDB[user.Id] = user
	//return nil
}