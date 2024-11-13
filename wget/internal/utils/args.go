package utils

type Args struct {
	Bdownload     bool
	Output        string
	Path          string
	LimitRate     string
	MultiDownload string
	Mirror        Mirror
	Url           []string
	UrlFile       []string
}

type Mirror struct {
	Active       bool
	Reject       string
	Exclude      string
	ConvertLinks bool
}
