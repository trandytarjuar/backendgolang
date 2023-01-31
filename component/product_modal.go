package model

import "restapigo/component/config"

type Product struct{
	product_id int `json:"id" form:"id" gorm:"primaryKey"`
	product_name string `json:"name" form:"name"`
	product_description string `json:"description" form:"description"`
	created_at string `json:"by" form:"by"`
}

func getAll(keyword string)([]product,error)
{
	var product []users
	result := config.DB.Where("id LIKE ? OR name LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&product)
	return product, result.Error

}