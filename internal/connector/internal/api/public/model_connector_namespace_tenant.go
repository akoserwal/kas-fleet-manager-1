/*
 * Connector Management API
 *
 * Connector Management API is a REST API to manage connectors.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// ConnectorNamespaceTenant struct for ConnectorNamespaceTenant
type ConnectorNamespaceTenant struct {
	Kind ConnectorNamespaceTenantKind `json:"kind"`
	// Either user or organisation id depending on the value of kind
	Id string `json:"id"`
}
