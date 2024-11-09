package mod

type UserPermit struct {
	Id    int    `gorm:"type:int;autoIncrement:100000000;comment:id" json:"id"`
	Name  string `gorm:"type:varchar(32);not null;not null;comment:name" json:"name"`
	Title string `gorm:"type:varchar(32);comment:title" json:"title"`
	Value string `gorm:"type:enum('true', 'false');not null;default:true;comment:value" json:"value"`
}
