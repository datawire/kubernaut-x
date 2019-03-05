package broker

func (b *Broker) configureRouter() {
	b.router.Get("/healthz", b.healthz())
}
