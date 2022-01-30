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

// WindowsSecurityContextOptions contain Windows-specific options and credentials.
//+kubebuilder:object:generate=true
type V1WindowsSecurityContextOptions struct {
	// GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.
	GmsaCredentialSpec string `json:"gmsaCredentialSpec,omitempty"`
	// GMSACredentialSpecName is the name of the GMSA credential spec to use.
	GmsaCredentialSpecName string `json:"gmsaCredentialSpecName,omitempty"`
	// The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.
	RunAsUserName string `json:"runAsUserName,omitempty"`
}
