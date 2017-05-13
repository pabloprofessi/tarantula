package router

import (
	"github.com/pprofessi/client"
	"github.com/pprofessi/config"
	"github.com/pprofessi/response_writer"
	"net/http"
	"net/url"
)

func Router(w http.ResponseWriter, r *http.Request) {
	config.LOG.Infof("host to be proxied:  %s", r.Host)
	config.LOG.Infof("path recieved to be evaluated: %s", r.URL.String())

	if r.URL.Path[1:] == "ping" {
		response_writer.Response_writer(w, "pong")
		return
	}
	//redirectableResult := redirectableUri(r.Host + r.URL.Path)
	redirectableResult := redirectableUri(r.URL.Path[1:])
	//&& (r.Host == config.Config.ToDomain)

	if redirectableResult != "" {
		final_url_raw := "https://" + config.Config.ToDomain + "/" + redirectableResult
		//final_url_raw := "https://" + redirectableResult
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
		response_writer.Response_writer(w, "Source not found!\n")
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
