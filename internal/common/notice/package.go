package notice

type Package struct {
	Channel  Channel
	Template Template
}

func NewPackage(channel Channel, template Template) Package {
	return Package{
		Channel:  channel,
		Template: template,
	}
}
