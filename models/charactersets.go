package models

type CHARACTERSETS struct {
	CHARACTERSETNAME   string `gorm:"column:CHARACTER_SET_NAME;NOT NULL"`
	DEFAULTCOLLATENAME string `gorm:"column:DEFAULT_COLLATE_NAME;NOT NULL"`
	DESCRIPTION        string `gorm:"column:DESCRIPTION;NOT NULL"`
	MAXLEN             int64  `gorm:"column:MAXLEN;default:0;NOT NULL"`
}

func (CHARACTERSETS) TableName() string {
	return "CHARACTER_SETS"
}
