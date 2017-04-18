/* 
 * Switch Database REST API
 *
 * Switch API is the primary endpoint of data sevices and Switch DB's platform. You can do adding, editing, deleting or listing data works to your database with query operations by using this low-level API based on HTTP.
 *
 * OpenAPI spec version: 1.2.1
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type ListOrder struct {

	// Order types: ASC, DESC
	Type_ string `json:"type,omitempty"`

	// Column name.
	By string `json:"by,omitempty"`
}