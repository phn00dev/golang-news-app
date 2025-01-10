package constructor

import (
	"github.com/phn00dev/golang-news-app/internal/app"
	categoryConstructor "github.com/phn00dev/golang-news-app/internal/domain/category/constructor"
)

func Build(dependencies *app.Dependencies) {
	categoryConstructor.InitCategoryRequirementCreator(dependencies.DB, dependencies.Config)
}
