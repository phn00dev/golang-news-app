package constructor

import (
	"github.com/phn00dev/golang-news-app/internal/app"
	adminConstructor "github.com/phn00dev/golang-news-app/internal/domain/admin/constructor"
	categoryConstructor "github.com/phn00dev/golang-news-app/internal/domain/category/constructor"
	postConstructor "github.com/phn00dev/golang-news-app/internal/domain/post/constructor"
)

func Build(dependencies *app.Dependencies) {
	adminConstructor.InitAdminRequirementCreator(dependencies.DB)
	categoryConstructor.InitCategoryRequirementCreator(dependencies.DB, dependencies.Config)
	postConstructor.InitPostRequirementCreator(dependencies.DB, dependencies.Config)
}
