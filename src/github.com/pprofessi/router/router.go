package router

import (
	"github.com/pprofessi/client"
	"github.com/pprofessi/config"
	//"github.com/pprofessi/response_writer"
	"net/http"
)

func Router(w http.ResponseWriter, r *http.Request) {
	config.LOG.Debugf("r.URL.String(): %s", r.URL.String())
	config.LOG.Debugf("r.URL.Path[1:]: %s", r.URL.Path[1:])
	//config.LOG.Debugf("r.URL.Scheme: %s", r.URL.Scheme) no mustra nada, no se q onda
	config.LOG.Debugf("r.Host:  %s", r.Host)

	route_to_host := redirectableHost(r.Host)
	config.LOG.Debugf("route_to_host.SourceHost:  %s", route_to_host.SourceHost)
	config.LOG.Debugf("route_to_host.DestinyHost:  %s", route_to_host.DestinyHost)
	//config.LOG.Debugf("route_to_host.DestinyHost:  %s", route_to_host.RouteToUris[0])

	client.ClientBackendRequester(w, r)
	//response_writer.Response_writer(w)

}

//redireccion del host solamente sin uri
func redirectableHost(sourceHost string) RouteToHost {
	var route_to_host RouteToHost
	db := get_db()
	db.Where("source_host = ?", sourceHost).First(&route_to_host)
	config.LOG.Debugf("Route to host obtained from DB query: %s", route_to_host.DestinyHost)
	return route_to_host
}

// se debe redireccionar despues del host, es lo que viene desp de /
func redirectableUri(route_to_host RouteToHost) RouteToUri {
	var route_to_uri RouteToUri
	//db := get_db()
	route_to_uri.SourceUri = ""
	route_to_uri.DestinyUri = ""
	//db.Where("key_word = ?", uri).First(&rtu)
	//onfig.LOG.Debugf("Route to key-word obtained from DB query: %s", rtu.DestinyUri)
	//return rtkw
	return route_to_uri
}
