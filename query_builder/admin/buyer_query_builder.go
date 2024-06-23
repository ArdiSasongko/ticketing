package admin_query_builder

import (
	"github.com/ArdiSasongko/ticketing_app/query_builder"
	"gorm.io/gorm"
)

type BuyerQueryBuilder interface {
	querybuilder.BaseQueryBuilder
	GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error)
	getAllowedFilters() []string
	getAllowedSorts() []string
}

type BuyerQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilder
}

func NewBuyerQueryBuilder(db *gorm.DB) *BuyerQueryBuilderImpl {
	return &BuyerQueryBuilderImpl{
		BaseQueryBuilder: querybuilder.NewBaseQueryBuilder(db),
	}
}

func (buyerQueryBuilder *BuyerQueryBuilderImpl) GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error) {
	allowedFilters := buyerQueryBuilder.getAllowedFilters()
	allowedSorts := buyerQueryBuilder.getAllowedSorts()
	query, err := buyerQueryBuilder.GetQueryBuilder(filters, sort, allowedFilters, allowedSorts, limit, page)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (buyerQueryBuilder *BuyerQueryBuilderImpl) getAllowedFilters() []string {
	return []string{
		"name",
		"email",
	}
}

func (buyerQueryBuilder *BuyerQueryBuilderImpl) getAllowedSorts() []string {
	return []string{
		"name",
		"email",
		"created_at",
	}
}
