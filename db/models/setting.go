package mod

type Setting struct {
	Id    int    `gorm:"type:int;auto_increment;unique;comment:id" json:"id"`
	Name  string `gorm:"type:varchar(32);not null;comment:name" json:"name"`
	Title string `gorm:"type:varchar(1024);comment:title" json:"title"`
	Key   string `gorm:"type:varchar(32);comment:key_name" json:"key"`
	Value string `gorm:"type:longblob;comment:value" json:"value"`
}
