package models

type Product struct {
	ProId       string
	VersionId   string
	MainTitle   string
	Subtitle    string
	MainPicture string
	Price       float64
	Status      int
	Active      int
}
