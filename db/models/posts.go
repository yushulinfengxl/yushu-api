package mod

type Posts struct {
	Id       int    `gorm:"type:int;auto_increment;primaryKey;comment:id" json:"id"`
	Name     string `gorm:"type:varchar(32);default:posts-name;not null;comment:name" json:"name"`
	Types    string `gorm:"type:varchar(32);default:article;not null;comment:ilk" json:"ilk"`
	Title    string `gorm:"type:varchar(10240);null;comment:title" json:"title"`
	Data     string `gorm:"type:longblob;comment:data" json:"data"`
	Datetime string `gorm:"type:datetime;default:(CURRENT_TIMESTAMP);comment:datetime" json:"datetime"`
	Date     string `gorm:"type:date;default:(CURRENT_DATE);comment:date" json:"date"`
	//OtherUser int
	User   User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID;comment:user_id"`
	UserID int  `gorm:"type:int;comment:user_id" json:"user_id"`
}
