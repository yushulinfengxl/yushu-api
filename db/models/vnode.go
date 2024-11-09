package mod

type Vnode struct {
	Id    int    `gorm:"type:int;auto_increment;primaryKey;comment:id" json:"id"`
	Name  string `gorm:"type:varchar(32);default:vnode-name;not null;comment:name" json:"name"`
	Value string `gorm:"type:longblob;comment:data" json:"value"`
}
