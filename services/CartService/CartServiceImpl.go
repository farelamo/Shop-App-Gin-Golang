package CartService

import (
	"shop/models"
	"database/sql"
	"errors"
	"fmt"
)

type CartServiceImpl struct {
	DB *sql.DB
}

func NewCartService(DB *sql.DB) CartService {
	return &CartServiceImpl {
		DB: DB,
	}
}


func (f *CartServiceImpl) FindAll() (*[]models.Cart, error) {
	var Carts = []models.Cart{}

	sql 		:= `SELECT * FROM carts`
	rows, err 	:= f.DB.Query(sql)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Cart = models.Cart{}

		err = rows.Scan(&Cart.Id, &Cart.UserId, &Cart.ProductId, &Cart.Paid, &Cart.Checkout, &Cart.Total, &Cart.CreatedAt, &Cart.UpdatedAt)
		if err != nil {
			return nil, err
		}

		Carts = append(Carts, Cart)
	}

	return &Carts, nil
}

func (f *CartServiceImpl) Save(Cart *models.AddCart) (*models.Cart, error) {
	var newCart = models.Cart{}

	sql := `INSERT INTO carts (user_id, product_id, paid, checkout, total) VALUES ($1) Returning *`
	err := f.DB.QueryRow(sql, Cart.UserId, Cart.ProductId, Cart.Paid, Cart.Checkout, Cart.Total,).Scan(
		&newCart.Id, &newCart.UserId, &newCart.ProductId, &newCart.Paid, &newCart.Checkout, &newCart.Total, 
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
	err := f.DB.QueryRow(sql, id).Scan(&Cart.Id, &Cart.UserId, &Cart.ProductId, &Cart.Paid, &Cart.Checkout, &Cart.Total, &Cart.CreatedAt, &Cart.UpdatedAt,)
	if err != nil {
		return nil, err		
	}
	return &Cart, err
}

func (f *CartServiceImpl) Update(id int, Cart *models.Cart) (int, error) {
	sqlStatement := `UPDATE carts SET user_id=$2, product_id=$3, paid=$4, checkout=$5, total=$6 WHERE id=$1;`
	
	result, err := f.DB.Exec(sqlStatement, id, Cart.UserId, Cart.ProductId, Cart.Paid, Cart.Checkout, Cart.Total)
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

func (f *CartServiceImpl) Delete(id int) (int, error) {
	sql := `DELETE FROM carts WHERE id=$1;`
	res, err := f.DB.Exec(sql, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete Cart record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the Cart, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil	
}


