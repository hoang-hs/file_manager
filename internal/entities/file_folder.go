package entities

type File struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

func NewFile(Type string, Name string, Size int64) *File {
	return &File{
		Type: Type,
		Name: Name,
		Size: Size,
	}
}

type Files struct {
	ListFile []File
}

func NewFiles(file []File) *Files {
	return &Files{
		ListFile: file,
	}
}
