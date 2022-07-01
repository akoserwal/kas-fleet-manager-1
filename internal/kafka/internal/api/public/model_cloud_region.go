/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.11.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// CloudRegion Description of a region of a cloud provider.
type CloudRegion struct {
	// Indicates the type of this object. Will be 'CloudRegion'.
	Kind string `json:"kind,omitempty"`
	// Unique identifier of the object.
	Id string `json:"id,omitempty"`
	// Name of the region for display purposes, for example `N. Virginia`.
	DisplayName string `json:"display_name,omitempty"`
	// Whether the region is enabled for deploying an OSD cluster.
	Enabled bool `json:"enabled"`
	// Indicates whether there is capacity left per instance type
	Capacity []RegionCapacityListItem `json:"capacity"`
}
