package tables

import (
	"github.com/jinzhu/gorm"
)

type Transfer struct {
	gorm.Model
	Token   string
	Account string
}
