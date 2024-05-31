package querybuilder

import (
	"fmt"
	"slices"

	"gorm.io/gorm"
)

type BaseQueryBuilder interface {
	GetQueryBuilder(filters map[string]string, sort string, allowedFilters []string, allowedSorts []string, limit int, page int) (*gorm.DB, error)
}

type BaseQueryBuilderImpl struct {
	db *gorm.DB
}

func NewBaseQueryBuilder(db *gorm.DB) *BaseQueryBuilderImpl {
	return &BaseQueryBuilderImpl{
		db: db,
	}
}

func (baseQueryBuilder *BaseQueryBuilderImpl) GetQueryBuilder(filters map[string]string, sort string, allowedFilters []string, allowedSorts []string, limit int, page int) (*gorm.DB, error) {
	query := baseQueryBuilder.db

	for filter, value := range filters {
		if !slices.Contains(allowedFilters, filter) {
			return nil, fmt.Errorf("%s filter is now allowed", filter)
		}
		query = query.Where(fmt.Sprintf("%s = ?", filter), value)
	}

	if sort != "" {
		orderMode := "ASC"
		if sort[0] == '-' {
			orderMode = "DESC"
			sort = sort[1:]
		}
		if !slices.Contains(allowedSorts, sort) {
			return nil, fmt.Errorf("%s sort is now allowed", sort)
		}

		query = query.Order(fmt.Sprintf("%s %s", sort, orderMode))
	}

	if limit == 0 {
		limit = 15
	}
	query = query.Limit(limit)

	if page == 0 {
		page = 1
	}
	query = query.Offset((page - 1) * limit)

	return query, nil
}
