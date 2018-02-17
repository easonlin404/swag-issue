package models

type TagTemplate struct {
	ID         string
	TagId      string
	RemotePath string
	HtmlData   string
}

type Tag struct {
	ID            string
	Name          string
	Template      TagTemplate
	LabelType     LabelType
}

type LabelType struct{
	LabelName string
}