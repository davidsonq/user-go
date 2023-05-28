// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "davidsonquaresma@hotmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/davidsonq/user-go/blame/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/users": {
            "post": {
                "description": "Create new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ExampleInputUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponse"
                        }
                    },
                    "400": {
                        "description": "The body of the request is empty or lack of Nickname, Email, Password in the body of the request, the email has to be a valid email and nickname has to have at least 3 characters and password at least 6 ",
                        "schema": {
                            "$ref": "#/definitions/models.ErrosNoBody"
                        }
                    },
                    "401": {
                        "description": "This error is generated when trying to register a nickname or email already registered",
                        "schema": {
                            "$ref": "#/definitions/models.ErrosNoBody"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrosNoBody": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.ExampleInputUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "models.UserResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "https://usergo.onrender.com",
	BasePath:         "/api/users",
	Schemes:          []string{},
	Title:            "User Management Microservice with Login System",
	Description:      "This is a microservice built to manage users, with authentication and login features.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
