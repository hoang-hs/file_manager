package notice

type Template struct {
	Job     string
	Message string
}

func NewTemplate(job, message string) Template {
	return Template{
		Job:     job,
		Message: message,
	}
}
