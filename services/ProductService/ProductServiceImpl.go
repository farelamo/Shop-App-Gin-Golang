package ProductService

import (
	"shop/models"
	"database/sql"
	"errors"
	"fmt"
)

type ProductServiceImpl struct {
	DB *sql.DB
}

func NewProductService(DB *sql.DB) ProductService {
	return &ProductServiceImpl {
		DB: DB,
	}
}


func (f *ProductServiceImpl) FindAll() (*[]models.Product, error) {
	var Products = []models.Product{}

	sql 		:= `SELECT * FROM products`
	rows, err 	:= f.DB.Query(sql)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Product = models.Product{}

		err = rows.Scan(&Product.Id, &Product.Name, &Product.Stock, &Product.Price, &Product.Description, &Product.CategoryId, &Product.CreatedAt, &Product.UpdatedAt)
		if err != nil {
			return nil, err
		}

		Products = append(Products, Product)
	}

	return &Products, nil
}

func (f *ProductServiceImpl) FindByCategory(id int) (*[]models.Product, error) {
	var Products = []models.Product{}

	sql 		:= `SELECT * FROM products WHERE category_id=($1)`
	rows, err 	:= f.DB.Query(sql, id)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Product = models.Product{}

		err = rows.Scan(&Product.Id, &Product.Name, &Product.Stock, &Product.Price, &Product.Description, &Product.CategoryId, &Product.CreatedAt, &Product.UpdatedAt)
		if err != nil {
			return nil, err
		}

		Products = append(Products, Product)
	}

	return &Products, nil
}

func (f *ProductServiceImpl) Save(Product *models.AddProduct) (*models.Product, error) {
	var newProduct = models.Product{}

	if Product.Name == "" {
		return nil, errors.New("Product Name Must Be Filled")
	}

	sql := `INSERT INTO products (name, stock, price, description, category_id) VALUES ($1, $2, $3, $4, $5) Returning *`
	err := f.DB.QueryRow(sql, Product.Name, Product.Stock, Product.Price, Product.Description, Product.CategoryId).Scan(
		&newProduct.Id, &Product.Name, &Product.Stock, &Product.Price, &Product.Description, &Product.CategoryId, &newProduct.CreatedAt, &newProduct.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newProduct, nil
}

func (f *ProductServiceImpl) FindById(id int) (*models.Product, error) {
	var Product = models.Product{}
	sql := `SELECT * FROM products WHERE id=($1)`
	err := f.DB.QueryRow(sql, id).Scan(&Product.Id, &Product.Name, &Product.Stock, &Product.Price, &Product.Description, &Product.CategoryId, &Product.CreatedAt, &Product.UpdatedAt,)
	if err != nil {
		return nil, err		
	}
	return &Product, err
}

func (f *ProductServiceImpl) Update(id int, Product *models.Product) (int, error) {
	if Product.Name == "" {
		return 0, errors.New("Products Name Must Be Filled")
	}

	sqlStatement := `UPDATE products SET name=$2, price=$3, stock=$4, description=$5, category_id=$6 WHERE id=$1;`
	
	result, err := f.DB.Exec(sqlStatement, id, Product.Name, Product.Price, Product.Stock, Product.Description, Product.CategoryId)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating product record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the product, please try again later"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (f *ProductServiceImpl) Delete(id int) (int, error) {
	sql := `DELETE FROM products WHERE id=$1;`
	res, err := f.DB.Exec(sql, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete product record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the product, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil	
}


