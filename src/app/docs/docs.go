// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "soberkoder@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/addbeer": {
            "post": {
                "description": "add by json beer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "beers"
                ],
                "summary": "Add a Beer",
                "parameters": [
                    {
                        "description": "Add Beer",
                        "name": "beer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Beer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Beer"
                        }
                    }
                }
            }
        },
        "/api/getbeer/{id}": {
            "get": {
                "description": "Get Beer for by beer_id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "beer"
                ],
                "summary": "Get Beer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Beer Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        }
    },
    "definitions": {
        "main.Beer": {
            "type": "object",
            "properties": {
                "abv": {
                    "type": "number"
                },
                "ibu": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "srm": {
                    "type": "number"
                },
                "style": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "upc": {
                    "type": "number"
                },
                "update": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Sample API",
	Description: "This is a sample serice for managing orders",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
