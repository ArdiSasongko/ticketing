package admin_query_builder

import (
	"github.com/ArdiSasongko/ticketing_app/query_builder"
	"gorm.io/gorm"
)

type SellerQueryBuilder interface {
	querybuilder.BaseQueryBuilder
	GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error)
	getAllowedFilters() []string
	getAllowedSorts() []string
}

type SellerQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilder
}

func NewSellerQueryBuilder(db *gorm.DB) *SellerQueryBuilderImpl {
	return &SellerQueryBuilderImpl{
		BaseQueryBuilder: querybuilder.NewBaseQueryBuilder(db),
	}
}

func (sellerQueryBuilder *SellerQueryBuilderImpl) GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error) {
	allowedFilters := sellerQueryBuilder.getAllowedFilters()
	allowedSorts := sellerQueryBuilder.getAllowedSorts()
	query, err := sellerQueryBuilder.GetQueryBuilder(filters, sort, allowedFilters, allowedSorts, limit, page)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (sellerQueryBuilder *SellerQueryBuilderImpl) getAllowedFilters() []string {
	return []string{
		"name",
		"email",
	}
}

func (sellerQueryBuilder *SellerQueryBuilderImpl) getAllowedSorts() []string {
	return []string{
		"name",
		"email",
		"created_at",
	}
}
