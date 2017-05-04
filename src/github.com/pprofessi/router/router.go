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

	redirectableResult := redirectableUri(r.URL.Path[1:])
	//&& (r.Host == config.Config.ToDomain)
	if redirectableResult != "" {

		final_url_raw := "https://" + config.Config.ToDomain + "/" + redirectableResult

		destinyRouteURL, err := url.Parse(final_url_raw)
		if err != nil {
			config.LOG.Errorf("failed parse destinyRouteString")
			config.LOG.Errorf("Error: %s", err)
		}

		r.URL = destinyRouteURL
		r.Host = destinyRouteURL.Host
		r.RequestURI = ""
		client.ClientBackendRequester(w, r)

	} else {
		response_writer.Response_writer(w)
	}

}

//redireccion del host solamente sin uri
func redirectableUri(fromUrl string) string {
	var toUrl string

	urlsdb := get_db()

	stmt, err := urlsdb.Prepare("SELECT toUrl FROM target_keywords where fromUrl=? LIMIT 1")
	if err != nil {
		config.LOG.Errorf("Error preparing query redirectableUri: %s", err)
	}

	err = stmt.QueryRow(fromUrl).Scan(&toUrl)
	if err != nil {
		config.LOG.Errorf("Error quering toUrl: %s", err)
	}

	config.LOG.Debugf("Route to host obtained from DB query: %s", toUrl)
	return toUrl
}
