package plugins

import "myoo/models"

type User struct {
	Id          int64
	Name        string
	Avatar      string
	Description string
}

func CheckUser(value models.User) User {
	template := User{
		Id:          value.Id,
		Name:        value.Name,
		Avatar:      CheckAvatar(value.Avatar),
		Description: CheckDescription(value.Description),
	}
	return template
}

func CheckUsers(value []models.User) []User {
	template := []User{}
	for v := range value {
		template = append(template, User{
			Id:          value[v].Id,
			Name:        value[v].Name,
			Avatar:      CheckAvatar(value[v].Avatar),
			Description: CheckDescription(value[v].Description),
		})
	}
	return template
}

func CheckAvatar(value string) string {
	if value == "" {
		return "/src/avatar/default.jpg"
	}
	return value
}

func CheckDescription(value string) string {
	if value == "" {
		return "暂无签名..."
	}
	return value
}
