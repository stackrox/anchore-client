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

type EventResponseEventResource struct {
	Id     string `json:"id,omitempty"`
	Type_  string `json:"type,omitempty"`
	UserId string `json:"user_id,omitempty"`
}
