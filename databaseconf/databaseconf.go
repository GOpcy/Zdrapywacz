package databaseconf

type Offer struct {
	Title         string
	Company       string
	Location      string
	Experience    string
	Image         string
	OperatingMode string
	Tags          string
	URL           string `gorm:"primarykey"`
}
