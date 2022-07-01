/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.11.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// CloudProvider Cloud provider.
type CloudProvider struct {
	// Indicates the type of this object. Will be 'CloudProvider' link.
	Kind string `json:"kind,omitempty"`
	// Unique identifier of the object.
	Id string `json:"id,omitempty"`
	// Name of the cloud provider for display purposes.
	DisplayName string `json:"display_name,omitempty"`
	// Human friendly identifier of the cloud provider, for example `aws`.
	Name string `json:"name,omitempty"`
	// Whether the cloud provider is enabled for deploying an OSD cluster.
	Enabled bool `json:"enabled"`
}
