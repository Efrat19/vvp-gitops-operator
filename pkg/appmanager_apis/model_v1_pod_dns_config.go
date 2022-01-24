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

// PodDNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.
//+kubebuilder:object:generate=true
type V1PodDnsConfig struct {
	// A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.
	Nameservers []string `json:"nameservers,omitempty"`
	// A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.
	Options []V1PodDnsConfigOption `json:"options,omitempty"`
	// A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.
	Searches []string `json:"searches,omitempty"`
}
