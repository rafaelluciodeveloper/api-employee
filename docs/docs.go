// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package main

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
        "/": {
            "get": {
                "description": "get all employees",
                "tags": [
                    "Employee"
                ],
                "summary": "Get all employees",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Employee"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "description": "Update employee",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Update employee based on parameters",
                "parameters": [
                    {
                        "description": "Employee Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new employee",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Employee"
                ],
                "summary": "Create new employee based on paramters",
                "parameters": [
                    {
                        "description": "Employee Data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/{id}": {
            "get": {
                "description": "get employee by ID",
                "tags": [
                    "Employee"
                ],
                "summary": "Get one employee",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete employee by ID",
                "tags": [
                    "Employee"
                ],
                "summary": "Delete one employee",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Employee": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "occupation": {
                    "type": "string"
                },
                "salary": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "https://api-employee-v1.herokuapp.com",
	BasePath:         "/employee",
	Schemes:          []string{},
	Title:            "Employee API documentation",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
