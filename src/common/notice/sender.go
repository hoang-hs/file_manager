package notice

type Sender interface {
	Send(p Package) Response
}
