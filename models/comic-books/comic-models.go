package comicbooks

type Comic_books struct {
	IdComic  int64  `gorm:"primaryKey;autoIncrement" json:"id_comic" form:"id_comic"`
	NameC    string `gorm:"Varchar(300)" json:"name_c" form:"name_c"`
	Genre    string `gorm:"Varchar(300)" json:"genre" form:"genre"`
	Image    string `gorm:"Varchar(300)" json:"image" form:"image"`
	CountryC string `gorm:"Varchar(300)" json:"country_c" form:"country_c"`
	AccessU  int    `gorm:"default:0" json:"access_u" form:"access_u"` // Changed type to int
	ViewU    string `gorm:"Varchar(10)" json:"view_u" form:"view_u"`
}

type Comic_details struct {
	IdDetComic  int64       `gorm:"primaryKey" json:"id_det_comic" form:"id_det_comic"`
	ComicId     int64       `json:"comic_id" form:"comic_id"`
	Status      string      `gorm:"Varchar(300)" json:"status" form:"status"`
	ReaderAge   string      `gorm:"Varchar(300)" json:"reader_age" form:"reader_age"`
	HowRead     string      `gorm:"Varchar(300)" json:"how_read" form:"how_read"`
	ComicArtist string      `gorm:"Varchar(300)" json:"comic_artist" form:"comic_artist"`
	Description string      `gorm:"text" json:"description" form:"description"`
	Comic       Comic_books `gorm:"foreignKey:ComicId;references:IdComic"`
}

type Comic_subbtns struct {
	IdSub   int64       `gorm:"primaryKey" json:"id_sub" form:"id_sub"`
	ComicID int64       `json:"comic_id" form:"comic_id"`
	Name    string      `gorm:"varchar(200)" json:"name" form:"name"`
	NameSub string      `gorm:"varchar(200)" json:"name_sub" form:"name_sub"`
	Comic   Comic_books `gorm:"foreignKey:ComicID;references:IdComic"`
}

type Subbtns_images struct {
	IdImage int64         `gorm:"primaryKey" json:"id_image" form:"id_image"`
	SubId   int64         `json:"sub_id" form:"sub_id"`
	Images  string        `gorm:"varchar(200)" json:"images" form:"images"`
	Subbtns Comic_subbtns `gorm:"foreignKey:SubId;references:IdSub"`
}
