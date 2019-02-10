/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.9
 * Contact: nurmi@anchore.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Api Version string
type ServiceVersionApi struct {
	// Semantic version of the api
	Version string `json:"version,omitempty"`
}
