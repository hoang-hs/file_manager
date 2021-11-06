package notice

type Package struct {
	Channel  Channel
	Template NoticeTemplate
}

func NewPackage(channel Channel, template NoticeTemplate) Package {
	return Package{
		Channel:  channel,
		Template: template,
	}
}
