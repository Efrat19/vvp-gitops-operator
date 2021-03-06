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

// TCPSocketAction describes an action based on opening a socket
//+kubebuilder:object:generate=true
type V1TcpSocketAction struct {
	// Optional: Host name to connect to, defaults to the pod IP.
	Host string `json:"host,omitempty"`
	// IntOrString is a type that can hold an int32 or a string.  When used in JSON or YAML marshalling and unmarshalling, it produces or consumes the inner type.  This allows you to have, for example, a JSON field that can accept a name or number.
	Port *IntOrString `json:"port"`
}
