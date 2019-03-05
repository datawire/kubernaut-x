package broker

import (
	"fmt"
	"github.com/datawire/kubernaut/pkg/version"
	"net/http"
)

func (b *Broker) healthz() http.HandlerFunc {
	versionJSON := version.GetVersionJSON()
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("content-type", "application/json")

		if _, err := resp.Write([]byte(versionJSON)); err != nil {
			fmt.Println(err)
		}
	}
}

func (b *Broker) readyz() http.HandlerFunc {
	return nil
}
