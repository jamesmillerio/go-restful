// package swagger implements the structures of the Swagger (https://github.com/wordnik/swagger-core/wiki) specification
package swagger

type ResourceListing struct {
	apiVersion, swaggerVersion, basePath string
	apis                                 []Api
}

type Api struct {
	Path        string      `json:"path"`
	Description string      `json:"description"`
	Operations  []Operation `json:"operations"`
	Models      []Model     `json:"models"`
}

type ApiDeclaration struct {
	ApiVersion     string `json:"apiVersion"`
	SwaggerVersion string `json:"swaggerVersion"`
	BasePath       string `json:"basePath"`
	ResourcePath   string `json:"resourcePath"`
}

type Operation struct {
	HttpMethod     string `json:"httpMethod"`
	Nickname       string `json:"nickname"`
	ResponseClass  string `json:"responseClass"`
	Summary        string `json:"summary"`
	Notes          string `json:"notes"`
	Parameters     []Parameter
	ErrorResponses []ErrorResponse
}

type Parameter struct {
	ParamType       string `json:"paramType"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	DataType        string `json:"dataType"`
	Required        bool   `json:"required"`
	AllowableValues map[string]string
	AllowMultiple   bool
}

type ErrorResponse struct {
	Code   int    `json:"code"`
	Reason string `json:"reason"`
}

type Model struct{}