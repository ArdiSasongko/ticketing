package buyer_query_builder

import (
	"github.com/ArdiSasongko/ticketing_app/query_builder"
	"gorm.io/gorm"
)

type OrderQueryBuilder interface {
	querybuilder.BaseQueryBuilder
	GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error)
	getAllowedFilters() []string
	getAllowedSorts() []string
}

type OrderQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilder
}

func NewOrderQueryBuilder(db *gorm.DB) *OrderQueryBuilderImpl {
	return &OrderQueryBuilderImpl{
		BaseQueryBuilder: querybuilder.NewBaseQueryBuilder(db),
	}
}

func (orderQueryBuilder *OrderQueryBuilderImpl) GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error) {
	allowedFilters := orderQueryBuilder.getAllowedFilters()
	allowedSorts := orderQueryBuilder.getAllowedSorts()
	query, err := orderQueryBuilder.GetQueryBuilder(filters, sort, allowedFilters, allowedSorts, limit, page)
	if err != nil {
		return nil, err
	}
	query = query.Preload("HistoryItems.Event").Preload("Buyer")

	return query, nil
}

func (orderQueryBuilder *OrderQueryBuilderImpl) getAllowedFilters() []string {
	return []string{
		"number",
		"payment_status",
	}
}

func (orderQueryBuilder *OrderQueryBuilderImpl) getAllowedSorts() []string {
	return []string{
		"paid_at",
		"created_at",
		"total",
	}
}
