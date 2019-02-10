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

type ContentPackageResponseContent struct {
	License  string `json:"license,omitempty"`
	Location string `json:"location,omitempty"`
	Origin   string `json:"origin,omitempty"`
	Package_ string `json:"package,omitempty"`
	Size     string `json:"size,omitempty"`
	Type_    string `json:"type,omitempty"`
	Version  string `json:"version,omitempty"`
}