/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.6.1
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// HTTPGetAction describes an action based on HTTP Get requests.
//+kubebuilder:object:generate=true
type V1HttpGetAction struct {
	// Host name to connect to, defaults to the pod IP. You probably want to set \"Host\" in httpHeaders instead.
	Host string `json:"host,omitempty"`
	// Custom headers to set in the request. HTTP allows repeated headers.
	HttpHeaders []V1HttpHeader `json:"httpHeaders,omitempty"`
	// Path to access on the HTTP server.
	Path string `json:"path,omitempty"`
	// IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.
	Port *IntOrString `json:"port"`
	// Scheme to use for connecting to the host. Defaults to HTTP.
	Scheme string `json:"scheme,omitempty"`
}
