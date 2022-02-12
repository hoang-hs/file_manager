package notice

type Template struct {
	Job       string
	Message   string
	ParseMode string
}

func NewTemplate(job, message, mode string) Template {
	return Template{
		Job:       job,
		Message:   message,
		ParseMode: mode,
	}
}
