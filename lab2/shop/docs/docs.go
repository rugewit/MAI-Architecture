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
        "/users": {
            "get": {
                "description": "Get Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Get Users",
                "operationId": "get-users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    }
                }
            }
        },
        "/users/": {
            "post": {
                "description": "Create an user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Create an user",
                "operationId": "create-user",
                "parameters": [
                    {
                        "description": "user",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    }
                }
            }
        },
        "/users/pattern-search": {
            "post": {
                "description": "Pattern Search. % The percent sign represents zero, one, or multiple characters. _ The underscore sign represents one, single character",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Pattern Search",
                "operationId": "pattern-search-users",
                "parameters": [
                    {
                        "description": "pattern search request",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PatternSearchRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Get an user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Get an user",
                "operationId": "get-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {}
                    }
                }
            },
            "put": {
                "description": "Update an user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Update an user",
                "operationId": "update-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated user",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignUpUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Delete a user",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "models.PatternSearchRequest": {
            "type": "object",
            "properties": {
                "lastNamePattern": {
                    "type": "string",
                    "example": "%_%"
                },
                "namePattern": {
                    "type": "string",
                    "example": "%_%"
                }
            }
        },
        "models.SignUpUser": {
            "type": "object",
            "properties": {
                "lastname": {
                    "type": "string",
                    "example": "Ivanov"
                },
                "name": {
                    "type": "string",
                    "example": "Alex"
                },
                "password": {
                    "type": "string",
                    "example": "qwerty"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "creationDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Shop",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
