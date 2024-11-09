package mod

type User struct {
	Id       int    `gorm:"type:int;primaryKey;unique;autoIncrement:100000000;comment:id" json:"id"`
	Name     string `gorm:"type:varchar(32);not null;comment:name" json:"name"`
	Account  string `gorm:"type:varchar(32);not null;comment:account" json:"account"`
	Nickname string `gorm:"type:varchar(256);comment:nickname" json:"nickname"`
	Data     string `gorm:"type:longblob;comment:data" json:"data"`
	Password string `gorm:"type:varchar(128);not null;comment:password" json:"password"`
	//Type     []usertype.UserType `gorm:"foreignKey:user_id;ilk:int;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null;comment:ilk" json:"ilk"`
	Type   UserType `gorm:"ForeignKey:TypeID;AssociationForeignKey:ID;comment:type_id"`
	TypeID int      `gorm:"type:int;comment:type_id" json:"type_id"`
}
