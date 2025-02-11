package databaseconf

type Offer struct {
	Title         string
	Company       string
	Location      string
	Experience    string
	OperatingMode string
	Tags          string
	URL           string `gorm:"primarykey"`
}
