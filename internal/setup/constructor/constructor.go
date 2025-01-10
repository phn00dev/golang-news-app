package constructor

import (
	"github.com/phn00dev/golang-news-app/internal/app"
	categoryConstructor "github.com/phn00dev/golang-news-app/internal/domain/category/constructor"
	postConstructor "github.com/phn00dev/golang-news-app/internal/domain/post/constructor"
)

func Build(dependencies *app.Dependencies) {
	categoryConstructor.InitCategoryRequirementCreator(dependencies.DB, dependencies.Config)
	postConstructor.InitPostRequirementCreator(dependencies.DB, dependencies.Config)
}
