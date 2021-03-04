/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConnectorCluster Schema for the request to update a data plane cluster's status
type ConnectorCluster struct {
	Id             string                        `json:"id,omitempty"`
	Kind           string                        `json:"kind,omitempty"`
	Href           string                        `json:"href,omitempty"`
	Metadata       ConnectorClusterAllOfMetadata `json:"metadata,omitempty"`
	ServiceAccount string                        `json:"service_account,omitempty"`
	Status         string                        `json:"status,omitempty"`
}
