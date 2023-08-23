package CartService

import (
	"database/sql"
	"errors"
	"fmt"
	"shop/models"
	product "shop/services/ProductService"
	user "shop/services/UserService"
	"strconv"
)

type CartServiceImpl struct {
	DB *sql.DB
}

func NewCartService(DB *sql.DB) CartService {
	return &CartServiceImpl {
		DB: DB,
	}
}


func (f *CartServiceImpl) FindAll(userId int) (*[]models.Cart, error) {
	var Carts = []models.Cart{}

	sql 		:= `SELECT * FROM carts where user_id = ($1)`
	rows, err 	:= f.DB.Query(sql, strconv.Itoa(userId))
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Cart = models.Cart{}
		var hehe *bool

		err = rows.Scan(&Cart.Id, &Cart.UserId, &Cart.ProductId, &Cart.Checkout, &Cart.Amount, &Cart.Total, &Cart.CreatedAt, &Cart.UpdatedAt)
		if err != nil {
			return nil, err
		}
		fmt.Println(hehe)
		Carts = append(Carts, Cart)
	}

	return &Carts, nil
}

func (f *CartServiceImpl) Save(userId int, Cart *models.AddCart) (*models.Cart, error) {
	var newCart = models.Cart{}

	productService 	:= product.NewProductService(f.DB)
	_, err 			:= productService.FindById(Cart.ProductId)
	if err != nil {
		return nil, err
	}

	userService 	:= user.NewUserService(f.DB)
	_, err 			= userService.FindById(userId)
	if err != nil {
		return nil, err
	}

	sql := `INSERT INTO carts (user_id, product_id, amount, checkout, total) VALUES ($1, $2, $3, $4, $5) Returning *`
	err = f.DB.QueryRow(sql, userId, Cart.ProductId, Cart.Amount, false, Cart.Total,).Scan(
		&newCart.Id, &newCart.UserId, &newCart.ProductId, &newCart.Checkout, &newCart.Amount, &newCart.Total, 
		&newCart.CreatedAt, &newCart.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newCart, nil
}

func (f *CartServiceImpl) FindById(id int) (*models.Cart, error) {
	var Cart = models.Cart{}
	sql := `SELECT * FROM carts WHERE id=($1)`
	err := f.DB.QueryRow(sql, id).Scan(&Cart.Id, &Cart.UserId, &Cart.ProductId, &Cart.Checkout, &Cart.Amount, &Cart.Total, &Cart.CreatedAt, &Cart.UpdatedAt,)
	if err != nil {
		return nil, err		
	}
	return &Cart, err
}

func (f *CartServiceImpl) Update(id int, Cart *models.Cart) (int, error) {
	sqlStatement := `UPDATE carts SET user_id=$2, product_id=$3, checkout=$4, amount=$5, total=$6 WHERE id=$1;`
	
	result, err := f.DB.Exec(sqlStatement, id, Cart.UserId, Cart.ProductId,  Cart.Checkout, Cart.Amount, Cart.Total)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating Cart record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the Cart, please try again later"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (f *CartServiceImpl) Delete(userId int, id int) (int, error) {
	data, err := f.FindById(id)
	if err != nil {
		e := fmt.Sprintf("error: %v product with id %d not found", err, id)
		return 0, errors.New(e)
	}

	if data.UserId != userId {
		return 0, errors.New("Sorry, this is not your product cart")
	}

	sql := `DELETE FROM carts WHERE id=$1;`
	res, err := f.DB.Exec(sql, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete product cart record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the product cart, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil	
}


