package infra

import (
	"net"
	"time"
)

type Machine interface {
	ID() string
	Class() string
	Flavor() string
	Family() string
	CreationTime() time.Time
	Delete() (bool, error)
	PublicIP() net.IP
	PublicDNS() string
	PrivateIP() net.IP
	PrivateDNS() string
}

type Network interface {
	ID() string
	Name() string
}
