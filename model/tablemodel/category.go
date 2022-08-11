package tablemodel

type Category struct {
	Id   int    `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

func (category *Category) TableName() string {
	return "categories"
}
