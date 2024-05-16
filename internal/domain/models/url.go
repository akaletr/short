package models

type URL struct {
	ID    int64  `json:"id" gorm:"column:id;primaryKey"`
	Alias string `json:"alias" gorm:"column:alias;unique"`
	URL   string `json:"url" gorm:"column:url"`
}
