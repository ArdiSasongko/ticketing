package seller_query_builder

import (
	querybuilder "github.com/ArdiSasongko/ticketing_app/query_builder"
	"gorm.io/gorm"
)

type EventQueryBuilder interface {
	querybuilder.BaseQueryBuilder
	GetBuilder(sellerId int, filters map[string]string, sort string, limit int, page int) (*gorm.DB, error)
	getAllowedFilters() []string
	getAllowedSorts() []string
}

type EventQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilder
}

func NewEventQueryBuilder(db *gorm.DB) *EventQueryBuilderImpl {
	return &EventQueryBuilderImpl{
		BaseQueryBuilder: querybuilder.NewBaseQueryBuilder(db),
	}
}

func (eventQueryBuilder *EventQueryBuilderImpl) GetBuilder(sellerId int, filters map[string]string, sort string, limit int, page int) (*gorm.DB, error) {
	allowedFilters := eventQueryBuilder.getAllowedFilters()
	allowedSorts := eventQueryBuilder.getAllowedSorts()
	query, err := eventQueryBuilder.GetQueryBuilder(filters, sort, allowedFilters, allowedSorts, limit, page)
	if err != nil {
		return nil, err
	}
	query = query.Where("seller_id = ?", sellerId)

	return query, nil
}

func (eventQueryBuilder *EventQueryBuilderImpl) getAllowedFilters() []string {
	return []string{
		"seller_id",
		"category",
		"location",
		"name",
	}
}

func (eventQueryBuilder *EventQueryBuilderImpl) getAllowedSorts() []string {
	return []string{
		"qty",
		"date",
		"category",
	}
}
