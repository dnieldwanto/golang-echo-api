package migrate

import "golang-echo-api/model"

func LoadAllModels() []interface{} {
	return []interface{}{
		&model.Category{},
		&model.Wallpaper{},
	}
}
