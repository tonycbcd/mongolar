package jsconfig

import (
	"net/http"
	"text/template"
)

// We are compiling the AngularJs config script into the program for fast loading.  All sites should have the same Angular JS config
// template with different values passed from the site config
const (
	ConfigScript = `
		var mongular = angular.module('mongular', 
			[
				{{ range .AngularModules }}
     					'{{.}}',
				{{ end }}
			]
		);

		mongular.constant('mongularConfig', {
    			mongular_url: '/{{ .APIEndPoint }}/',
    			templates_url: '/{{ .TemplateEndpoint }}'}
		);
		mongular.config(function($sceDelegateProvider) {
   			$sceDelegateProvider.resourceUrlWhitelist([
     				// Allow same origin resource loads.
     				'self',
				{{ range .ForeignDomains }}
     				'http://{{.}}/**',
				{{ end }}
   			]);
 		});`
)

// Current available values.  This may be reconfigured if it gets too large.
type JsConfigs struct {
	APIEndPoint      string
	TemplateEndpoint string
	ForeignDomains   []string
	AngularModules   []string
}

// Serve the config
func (c *JsConfigs) Serve(w http.ResponseWriter) {
	t := template.New("Mongolar Config JS")
	t.Parse(ConfigScript)
	t.Execute(w, c)
}
