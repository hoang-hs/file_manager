package resources

type File struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

type Files struct {
	ListDir []File
}
