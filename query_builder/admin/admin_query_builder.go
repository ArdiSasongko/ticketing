package admin_query_builder

import (
	"github.com/ArdiSasongko/ticketing_app/query_builder"
	"gorm.io/gorm"
)

type AdminQueryBuilder interface {
	querybuilder.BaseQueryBuilder
	GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error)
	getAllowedFilters() []string
	getAllowedSorts() []string
}

type AdminQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilder
}

func NewAdminQueryBuilder(db *gorm.DB) *AdminQueryBuilderImpl {
	return &AdminQueryBuilderImpl{
		BaseQueryBuilder: querybuilder.NewBaseQueryBuilder(db),
	}
}

func (adminQueryBuilder *AdminQueryBuilderImpl) GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error) {
	allowedFilters := adminQueryBuilder.getAllowedFilters()
	allowedSorts := adminQueryBuilder.getAllowedSorts()
	query, err := adminQueryBuilder.GetQueryBuilder(filters, sort, allowedFilters, allowedSorts, limit, page)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (adminQueryBuilder *AdminQueryBuilderImpl) getAllowedFilters() []string {
	return []string{
		"name",
		"email",
	}
}

func (adminQueryBuilder *AdminQueryBuilderImpl) getAllowedSorts() []string {
	return []string{
		"name",
		"email",
		"created_at",
	}
}
