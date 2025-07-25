package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

// ArticleCategory тип для enum
type ArticleCategory string

// Допустимые значения категорий
const (
	CategoryGeneral  ArticleCategory = "general"
	CategoryTech     ArticleCategory = "tech"
	CategoryScience  ArticleCategory = "science"
	CategoryPolitics ArticleCategory = "politics"
	CategoryHealth   ArticleCategory = "health"
)

func (ac ArticleCategory) IsValid() bool {
	switch ac {
	case CategoryGeneral, CategoryTech, CategoryScience, CategoryPolitics, CategoryHealth:
		return true
	}
	return false
}

// Value для интерфейса driver.Valuer (запись в БД)
func (ac ArticleCategory) Value() (driver.Value, error) {
	if !ac.IsValid() {
		return nil, errors.New("invalid article category")
	}
	return string(ac), nil
}

// Scan для интерфейса sql.Scanner (чтение из БД)
func (ac *ArticleCategory) Scan(value interface{}) error {
	if value == nil {
		*ac = CategoryGeneral
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan ArticleCategory")
	}

	*ac = ArticleCategory(str)
	if !ac.IsValid() {
		return fmt.Errorf("invalid ArticleCategory value: %s", str)
	}
	return nil
}
