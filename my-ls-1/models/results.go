package models

type Result struct {
	Perms    string
	Num      string
	UserOwn  string
	GroupOwn string
	Size     string
	Date     string
	Name     string
}

type ResultLen struct {
	Perms    int
	Num      int
	UserOwn  int
	GroupOwn int
	Size     int
	Date     int
}
