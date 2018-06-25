package extend

import (
	"hello/models"
)

/**
获取用户id
 */
func UserData(art []models.Article) []int{
	var idArr []int
	for _,val := range art{
		idArr = append(idArr,val.User_id)
	}
	return idArr
}
type ArticleAndUser struct {
	models.Article
	Phone string
	User_id int
}
/**
组合数据
 */
func ArticleAddUserData(art *[]models.Article,user *[]models.User) []ArticleAndUser{
	var re  []ArticleAndUser
	var g ArticleAndUser
	for _,val :=range *art{
		for _,u_val := range *user{
			if val.User_id == u_val.Id{
				g = ArticleAndUser{val,u_val.Phone,u_val.Id}
				re = append(re,g)
			}
		}
	}
	return re
}
