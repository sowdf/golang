package model

type Spectrum struct {
	Id          string
	Type        string
	Title       string
	From        string
	FromUrl     string
	Composition string
	Upload      string
	Views       int
	UploadDate  string
	ImageUrls   []string
	ClickNum    int
}
