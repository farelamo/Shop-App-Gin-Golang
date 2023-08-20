package UserService

import (
	"shop/models"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserImpl struct {
	DB *sql.DB
}

func NewUserService(DB *sql.DB) UserService {
	return &UserImpl {
		DB: DB,
	}
}


func (u *UserImpl) FindAll() (*[]models.User, error) {
	var users = []models.User{}

	sql 		:= `SELECT * FROM users`
	rows, err 	:= u.DB.Query(sql)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user = models.User{}

		err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Pass, &user.Age, &user.CreatedAt, &user.UpdatedAt,)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

func (u *UserImpl) Save(user *models.AddUser) (*models.User, error) {
	var newUser = models.User{}

	if user.Name == "" {
		return nil, errors.New("Name Must Be Filled")
	}else if strconv.Itoa(user.Age) == "" {
		return nil, errors.New("Age Must Be Filled")
	}else if user.Username == "" {
		return nil, errors.New("username Must Be Filled")
	}else if user.Pass == "" {
		return nil, errors.New("Password Must Be Filled")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Pass = string(hashedPassword)

	sql := `INSERT INTO users (name, username, password, age) VALUES ($1, $2, $3, $4) Returning *`
	err = u.DB.QueryRow(sql, user.Name, user.Username, user.Pass, user.Age,).Scan(
		&newUser.Id, &newUser.Name, &newUser.Username, &user.Pass, &newUser.Age, 
		&newUser.CreatedAt, &newUser.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (u *UserImpl) FindById(id int) (*models.User, error) {
	var user = models.User{}
	sql := `SELECT * FROM users WHERE id=($1)`
	err := u.DB.QueryRow(sql, id).Scan(&user.Id, &user.Name, &user.Username, &user.Pass, &user.Age, &user.CreatedAt, &user.UpdatedAt,)
	if err != nil {
		return nil, err		
	}
	return &user, err
}

func (u *UserImpl) Update(id int, user *models.User) (int, error) {
	if user.Name == "" {
		return 0, errors.New("Name Must Be Filled")
	}else if strconv.Itoa(user.Age) == "" {
		return 0, errors.New("Age Must Be Filled")
	}else if user.Username == "" {
		return 0, errors.New("Username Must Be Filled")
	}else if user.Pass == "" {
		return 0, errors.New("Password Must Be Filled")
	}

	sqlStatement := `UPDATE users SET name=$2, username=$3, password=$4, age=$5 WHERE id=$1;`
	
	result, err := u.DB.Exec(sqlStatement, id, user.Name,  user.Username, user.Pass, user.Age,)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating user record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the user, please try again later"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (u *UserImpl) Delete(id int) (int, error) {
	sql := `DELETE FROM users WHERE id=$1;`
	res, err := u.DB.Exec(sql, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete user record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the user, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil	
}


