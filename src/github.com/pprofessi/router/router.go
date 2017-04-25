package router

import (
	"github.com/pprofessi/client"
	"github.com/pprofessi/config"
	"github.com/pprofessi/response_writer"
	"net/http"
	"net/url"
)

func Router(w http.ResponseWriter, r *http.Request) {
	config.LOG.Debugf("r.URL.String(): %s", r.URL.String())
	config.LOG.Debugf("r.URL.Path[1:]: %s", r.URL.Path[1:])
	//config.LOG.Debugf("r.URL.Scheme: %s", r.URL.Scheme) no mustra nada, no se q onda
	config.LOG.Debugf("r.Host:  %s", r.Host)
	config.LOG.Debugf("r.RequestURI:  %s", r.RequestURI)

	route_to_host := redirectableHost(r.Host)
	config.LOG.Debugf("route_to_host.SourceHost:  %s", route_to_host.SourceHost)
	config.LOG.Debugf("route_to_host.DestinyHost:  %s", route_to_host.DestinyHost)
	config.LOG.Debugf("route_to_host.ID:  %s", route_to_host.ID)

	if route_to_host.DestinyHost != "" {
		route_to_uri := redirectableUri(route_to_host, r.URL.String())
		if route_to_uri.DestinyU != "" {

			final_url_raw := route_to_host.DestinyHost
			final_url_raw = final_url_raw + route_to_uri.DestinyUri

			config.LOG.Debugf("route_to_uri.ID:  %s", route_to_uri.ID)
			config.LOG.Debugf("source_uri.ID:  %s", route_to_uri.SourceUri)
			config.LOG.Debugf("destiny_uri.ID:  %s", route_to_uri.DestinyUri)
			config.LOG.Debugf("route_to_host_id.ID:  %s", route_to_uri.RouteToHostID)

			destinyRouteURL, err := url.Parse(final_url_raw)
			if err != nil {
				config.LOG.Errorf("failed parse destinyRouteString")
				config.LOG.Errorf("Error: %s", err)
			}

			r.URL = destinyRouteURL
			r.Host = destinyRouteURL.Host
			r.RequestURI = ""
			client.ClientBackendRequester(w, r)
		}

	} else {
		response_writer.Response_writer(w)
	}

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
func redirectableUri(route_to_host RouteToHost, source_uri string) RouteToUri {
	var route_to_uri RouteToUri
	db := get_db()
	db.Where("route_to_host_id = ?  and source_uri = ?", route_to_host.ID, source_uri).First(&route_to_uri)
	config.LOG.Debugf("Route to key-word obtained from DB query: %s", route_to_uri.DestinyUri)
	return route_to_uri
}
