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

// Java package content listings from images
type ContentJavaPackageResponse struct {
	Content     []ContentJavaPackageResponseContent `json:"content,omitempty"`
	ContentType string                              `json:"content_type,omitempty"`
	ImageDigest string                              `json:"imageDigest,omitempty"`
}
