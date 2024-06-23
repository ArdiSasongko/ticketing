package buyer_query_builder

import (
	"github.com/ArdiSasongko/ticketing_app/query_builder"
	"gorm.io/gorm"
)

type TicketQueryBuilder interface {
	querybuilder.BaseQueryBuilder
	GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error)
	getAllowedFilters() []string
	getAllowedSorts() []string
}

type TicketQueryBuilderImpl struct {
	querybuilder.BaseQueryBuilder
}

func NewTicketQueryBuilder(db *gorm.DB) *TicketQueryBuilderImpl {
	return &TicketQueryBuilderImpl{
		BaseQueryBuilder: querybuilder.NewBaseQueryBuilder(db),
	}
}

func (eventQueryBuilder *TicketQueryBuilderImpl) GetBuilder(filters map[string]string, sort string, limit int, page int) (*gorm.DB, error) {
	allowedFilters := eventQueryBuilder.getAllowedFilters()
	allowedSorts := eventQueryBuilder.getAllowedSorts()
	query, err := eventQueryBuilder.GetQueryBuilder(filters, sort, allowedFilters, allowedSorts, limit, page)
	if err != nil {
		return nil, err
	}
	query = query.Preload("Event.Seller")

	return query, nil
}

func (eventQueryBuilder *TicketQueryBuilderImpl) getAllowedFilters() []string {
	return []string{
		"status",
	}
}

func (eventQueryBuilder *TicketQueryBuilderImpl) getAllowedSorts() []string {
	return []string{}
}
