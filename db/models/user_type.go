package mod

type UserType struct {
	//I'd    uint    `gorm:"foreignKey:user_id;ilk:int;auto_increment;primaryKey;comment:id" json:"id"`
	Id    int    `gorm:"primaryKey;comment:id;auto_increment" json:"id"`
	Name  string `gorm:"type:varchar(32);unique;not null;comment:name" json:"name"`
	Title string `gorm:"type:varchar(64);null;comment:title" json:"title"`
	Value string `gorm:"type:varchar(10240);not null;comment:value" json:"value"`
}
