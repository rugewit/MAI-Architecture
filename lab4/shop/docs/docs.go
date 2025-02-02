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
        "/auth/login": {
            "post": {
                "description": "Login an user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth API"
                ],
                "summary": "Login an user",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "user",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LogInUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Bad request",
                        "schema": {}
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Resiger an user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth API"
                ],
                "summary": "Resiger an user",
                "operationId": "register",
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
                    "200": {
                        "description": "OK",
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
        "/basket/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a basket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basket API"
                ],
                "summary": "Update a basket",
                "operationId": "update-basket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "basket ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated basket",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Basket"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Basket"
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
            }
        },
        "/baskets": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get Baskets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basket API"
                ],
                "summary": "Get Baskets",
                "operationId": "get-baskets",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Basket"
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
        "/baskets/add/{idbasket}/{idproduct}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basket API"
                ],
                "summary": "Add a product",
                "operationId": "add-product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Basket ID",
                        "name": "idbasket",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "idproduct",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created"
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    }
                }
            }
        },
        "/baskets/remove/{idbasket}/{idproduct}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Remove a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basket API"
                ],
                "summary": "Remove a product",
                "operationId": "remove-product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Basket ID",
                        "name": "idbasket",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "idproduct",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Basket"
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
            }
        },
        "/baskets/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a basket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basket API"
                ],
                "summary": "Get a basket",
                "operationId": "get-basket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "basket ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Basket"
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
            }
        },
        "/products": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get Products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product API"
                ],
                "summary": "Get Products",
                "operationId": "get-products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
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
        "/products/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product API"
                ],
                "summary": "Create a product",
                "operationId": "create-product",
                "parameters": [
                    {
                        "description": "product",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "created",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {}
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product API"
                ],
                "summary": "Get a product",
                "operationId": "get-product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product API"
                ],
                "summary": "Update a product",
                "operationId": "update-product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "updated product",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product API"
                ],
                "summary": "Delete a product",
                "operationId": "delete-product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
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
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
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
        "/users/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
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
        "/users/basket/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get an user basket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Get an user basket",
                "operationId": "get-user-basket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Basket"
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
            }
        },
        "/users/login-search": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get an user by login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User API"
                ],
                "summary": "Get an user by login",
                "operationId": "get-user-by-login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user login",
                        "name": "login",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
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
            }
        },
        "/users/pattern-search": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
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
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
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
        "models.Basket": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Product"
                    }
                },
                "totalPrice": {
                    "type": "number"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.LogInUser": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
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
        "models.Product": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
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
                "login": {
                    "type": "string",
                    "example": "AlexHere"
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
                "_id": {
                    "type": "string"
                },
                "basket_id": {
                    "type": "string"
                },
                "creationDate": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "login": {
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
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "X-Token",
            "in": "header"
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
