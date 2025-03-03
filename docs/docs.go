// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "post": {
                "description": "Get access token by login",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authentication",
                "parameters": [
                    {
                        "description": "Auth body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/klikform_src_interfaces_v1_schemas_auths.AuthBodySchema"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/schemas.ResponseSchema"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/klikform_src_interfaces_v1_schemas_auths.AuthResponseSchema"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Failure response",
                        "schema": {
                            "$ref": "#/definitions/schemas.ResponseSchema"
                        }
                    }
                }
            }
        },
        "/welcome": {
            "get": {
                "description": "Welcome entry point to test API run",
                "tags": [
                    "Welcome"
                ],
                "summary": "Welcome point",
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/schemas.ResponseSchema"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/schemas.WelcomeResponseSchema"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "klikform_src_interfaces_v1_schemas_auths.AuthBodySchema": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "klikform_src_interfaces_v1_schemas_auths.AuthResponseSchema": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "schemas.ResponseSchema": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.WelcomeResponseSchema": {
            "type": "object",
            "properties": {
                "about": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "KlikForm Service",
	Description:      "KlikForm Service API Documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
