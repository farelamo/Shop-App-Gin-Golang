package CategoryService

import (
	"shop/models"
	"database/sql"
	"errors"
	"fmt"
)

type CategoryServiceImpl struct {
	DB *sql.DB
}

func NewCategoryService(DB *sql.DB) CategoryService {
	return &CategoryServiceImpl {
		DB: DB,
	}
}


func (f *CategoryServiceImpl) FindAll() (*[]models.Category, error) {
	var Categories = []models.Category{}

	sql 		:= `SELECT * FROM categories`
	rows, err 	:= f.DB.Query(sql)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Category = models.Category{}

		err = rows.Scan(&Category.Id, &Category.Name, &Category.CreatedAt, &Category.UpdatedAt)
		if err != nil {
			return nil, err
		}

		Categories = append(Categories, Category)
	}

	return &Categories, nil
}

func (f *CategoryServiceImpl) Save(Category *models.AddCategory) (*models.Category, error) {
	var newCategory = models.Category{}

	if Category.Name == "" {
		return nil, errors.New("Category Name Must Be Filled")
	}

	sql := `INSERT INTO categories (name) VALUES ($1) Returning *`
	err := f.DB.QueryRow(sql, Category.Name).Scan(
		&newCategory.Id, &newCategory.Name, &newCategory.CreatedAt, &newCategory.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newCategory, nil
}

func (f *CategoryServiceImpl) FindById(id int) (*models.Category, error) {
	var Category = models.Category{}
	sql := `SELECT * FROM categories WHERE id=($1)`
	err := f.DB.QueryRow(sql, id).Scan(&Category.Id, &Category.Name, &Category.CreatedAt, &Category.UpdatedAt,)
	if err != nil {
		return nil, err		
	}
	return &Category, err
}

func (f *CategoryServiceImpl) Update(id int, Category *models.Category) (int, error) {
	if Category.Name == "" {
		return 0, errors.New("Category Name Must Be Filled")
	}

	sqlStatement := `UPDATE categories SET name=$2 WHERE id=$1;`
	
	result, err := f.DB.Exec(sqlStatement, id, Category.Name)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while updating category record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := result.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after update data: %v", err)
		return 0, errors.New(e) 
	}

	if count == 0 {
		e := "could not update the category, please try again later"
		return 0, errors.New(e) 
	}
	return int(count), nil
}

func (f *CategoryServiceImpl) Delete(id int) (int, error) {
	sql := `DELETE FROM categories WHERE id=$1;`
	res, err := f.DB.Exec(sql, id)
	if err != nil {
		e := fmt.Sprintf("error: %v occurred while delete category record with id: %d", err, id)
		return 0, errors.New(e)
	}
	count, err := res.RowsAffected()
	if err != nil {
		e := fmt.Sprintf("error occurred from database after delete data: %v", err)
		return 0, errors.New(e)		
	}

	if count == 0 {
		e := fmt.Sprintf("could not delete the category, there might be no data for ID %d", id)
		return 0, errors.New(e) 
	}
	return int(count), nil	
}


