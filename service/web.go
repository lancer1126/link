package service

var (
	httpServers = newServerPool[*httpServer]()
)

func startWebService() Service {
	return nil
}
