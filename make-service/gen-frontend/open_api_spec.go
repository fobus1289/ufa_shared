package genfrontend

// OpenAPISpec represents the structure of an OpenAPI 2.0 specification.
type OpenAPISpec struct {
	Swagger             string                        `json:"swagger"`
	Info                Info                          `json:"info"`
	Host                string                        `json:"host"`
	BasePath            string                        `json:"basePath"`
	Schemes             []string                      `json:"schemes"`
	Paths               map[string]PathItem           `json:"paths"`
	Definitions         map[string]Definition         `json:"definitions"`
	Parameters          map[string]Parameter          `json:"parameters,omitempty"`
	Responses           map[string]Response           `json:"responses,omitempty"`
	Security            []SecurityRequirement         `json:"security,omitempty"`
	SecurityDefinitions map[string]SecurityDefinition `json:"securityDefinitions,omitempty"`
}

// Info holds API metadata.
type Info struct {
	Version        string   `json:"version"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	TermsOfService string   `json:"termsOfService,omitempty"`
	Contact        *Contact `json:"contact,omitempty"`
	License        *License `json:"license,omitempty"`
}

// Contact holds contact information.
type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// License holds license information.
type License struct {
	Name string `json:"name"`
	URL  string `json:"url,omitempty"`
}

// PathItem represents the available operations for a single API path.
type PathItem struct {
	Get     *Operation `json:"get,omitempty"`
	Post    *Operation `json:"post,omitempty"`
	Put     *Operation `json:"put,omitempty"`
	Delete  *Operation `json:"delete,omitempty"`
	Options *Operation `json:"options,omitempty"`
	Head    *Operation `json:"head,omitempty"`
	Patch   *Operation `json:"patch,omitempty"`
}

// Operation describes a single API operation.
type Operation struct {
	Summary     string                `json:"summary,omitempty"`
	Description string                `json:"description,omitempty"`
	OperationID string                `json:"operationId,omitempty"`
	Produces    []string              `json:"produces,omitempty"`
	Consumes    []string              `json:"consumes,omitempty"`
	Tags        []string              `json:"tags,omitempty"`
	Parameters  []Parameter           `json:"parameters,omitempty"`
	Responses   map[string]Response   `json:"responses"`
	Security    []SecurityRequirement `json:"security,omitempty"`
}

// Parameter describes a single operation parameter.
type Parameter struct {
	Name        string  `json:"name"`
	In          string  `json:"in"`
	Description string  `json:"description,omitempty"`
	Required    bool    `json:"required,omitempty"`
	Type        string  `json:"type,omitempty"`
	Format      string  `json:"format,omitempty"`
	Schema      *Schema `json:"schema,omitempty"`
}

// Response describes a single API response.
type Response struct {
	Description string            `json:"description"`
	Schema      *Schema           `json:"schema,omitempty"`
	Headers     map[string]Header `json:"headers,omitempty"`
}

// Header describes a single response header.
type Header struct {
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

// Schema represents a data model schema.
type Schema struct {
	Type       string            `json:"type,omitempty"`
	Properties map[string]Schema `json:"properties,omitempty"`
	Items      *Schema           `json:"items,omitempty"`
}

// Definition represents a reusable data model schema.
type Definition struct {
	Description string            `json:"description,omitempty"`
	Type        string            `json:"type"`
	Properties  map[string]Schema `json:"properties,omitempty"`
	Required    []string          `json:"required,omitempty"`
}

// SecurityRequirement represents the security requirements for an API operation.
type SecurityRequirement struct {
	Requirement map[string][]string `json:"securityRequirement"`
}

// SecurityDefinition describes a security scheme.
type SecurityDefinition struct {
	Type             string `json:"type"`
	Description      string `json:"description,omitempty"`
	Name             string `json:"name,omitempty"`
	In               string `json:"in,omitempty"`
	Flow             string `json:"flow,omitempty"`
	AuthorizationURL string `json:"authorizationUrl,omitempty"`
	TokenURL         string `json:"tokenUrl,omitempty"`
}

// GetSwagger returns the Swagger version.
func (s *OpenAPISpec) GetSwagger() string {
	return s.Swagger
}

// GetInfo returns the Info struct.
func (s *OpenAPISpec) GetInfo() Info {
	return s.Info
}

// GetHost returns the host string.
func (s *OpenAPISpec) GetHost() string {
	return s.Host
}

// GetBasePath returns the base path string.
func (s *OpenAPISpec) GetBasePath() string {
	return s.BasePath
}

// GetSchemes returns a copy of the schemes slice or an empty slice if nil.
func (s *OpenAPISpec) GetSchemes() []string {
	if s.Schemes == nil {
		return []string{}
	}
	schemesCopy := make([]string, len(s.Schemes))
	copy(schemesCopy, s.Schemes)
	return schemesCopy
}

// GetPaths returns a copy of the paths map or an empty map if nil.
// func (s *OpenAPISpec) GetPaths() map[string]PathItem {
// 	if s.Paths == nil {
// 		return make(map[string]PathItem)
// 	}
// 	pathsCopy := make(map[string]PathItem)
// 	for k, v := range s.Paths {
// 		pathsCopy[k] = v
// 	}
// 	return pathsCopy
// }

// GetDefinitions returns a copy of the definitions map or an empty map if nil.
func (s *OpenAPISpec) GetDefinitions() map[string]Definition {
	if s.Definitions == nil {
		return make(map[string]Definition)
	}
	definitionsCopy := make(map[string]Definition)
	for k, v := range s.Definitions {
		definitionsCopy[k] = v
	}
	return definitionsCopy
}

// GetParameters returns a copy of the parameters map or an empty map if nil.
func (s *OpenAPISpec) GetParameters() map[string]Parameter {
	if s.Parameters == nil {
		return make(map[string]Parameter)
	}
	parametersCopy := make(map[string]Parameter)
	for k, v := range s.Parameters {
		parametersCopy[k] = v
	}
	return parametersCopy
}

// GetResponses returns a copy of the responses map or an empty map if nil.
func (s *OpenAPISpec) GetResponses() map[string]Response {
	if s.Responses == nil {
		return make(map[string]Response)
	}
	responsesCopy := make(map[string]Response)
	for k, v := range s.Responses {
		responsesCopy[k] = v
	}
	return responsesCopy
}

// GetSecurity returns a copy of the security slice or an empty slice if nil.
func (s *OpenAPISpec) GetSecurity() []SecurityRequirement {
	if s.Security == nil {
		return []SecurityRequirement{}
	}
	securityCopy := make([]SecurityRequirement, len(s.Security))
	copy(securityCopy, s.Security)
	return securityCopy
}

// GetSecurityDefinitions returns a copy of the security definitions map or an empty map if nil.
func (s *OpenAPISpec) GetSecurityDefinitions() map[string]SecurityDefinition {
	if s.SecurityDefinitions == nil {
		return make(map[string]SecurityDefinition)
	}
	securityDefinitionsCopy := make(map[string]SecurityDefinition)
	for k, v := range s.SecurityDefinitions {
		securityDefinitionsCopy[k] = v
	}
	return securityDefinitionsCopy
}
