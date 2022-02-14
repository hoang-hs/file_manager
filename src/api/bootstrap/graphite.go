package bootstrap

import (
	"file_manager/src/configs"
	"github.com/marpaia/graphite-golang"
	"go.uber.org/fx"
)

func LoadGraphite(cf *configs.Config) []fx.Option {
	return []fx.Option{
		fx.Supply(newGraphite(cf.GraphiteHost, cf.GraphitePort)),
	}
}

func newGraphite(host string, port int) *graphite.Graphite {
	//graphiteClient, err := graphite.NewGraphite(host, port)
	//if err != nil {
	//	log.Fatalf("Can not connect to graphite, err: [%s]\n", err)
	//} else {
	//	log.Println("Connected to graphite successfully")
	//}
	graphiteClient := &graphite.Graphite{}
	return graphiteClient
}
