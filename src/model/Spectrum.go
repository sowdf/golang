package model

type Spectrum struct {
	Id          string
	Type        string
	Title       string
	From        string
	FromUrl     string
	Composition string
	Uploader    string
	Views       int
	UploadDate  string
	ImageUrls   []string
	ClickNum    int
	Lyrics      string
}
