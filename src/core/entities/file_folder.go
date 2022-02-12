package entities

type File struct {
	Type string
	Name string
	Size int64
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
