// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support - Donaderoyan",
            "email": "donaderoyan@gmail.com"
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
        "/login": {
            "post": {
                "description": "Logs in a user by email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "Login Credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/login.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful, returns access token",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request, invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register a new user with email, password, first name, and last name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "Register Input",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/register.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User registered successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request, invalid input",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "409": {
                        "description": "Conflict, user already exists",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/musicalinfo/{id}": {
            "post": {
                "description": "Create musical information for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create musical information",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Musical information to create",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/musicalinfo.MusicalInfoInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Musical information created successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "patch": {
                "description": "Update musical information for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update musical information (partial update)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Musical information to update",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/musicalinfo.MusicalInfoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Musical information updated successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/user/profile/{id}": {
            "get": {
                "description": "Get the profile of a user by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/profile.UpdateProfileInput"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "put": {
                "description": "Update the profile of a user by their ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Profile Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/profile.UpdateProfileInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile updated successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "patch": {
                "description": "Update the profile of a user by their ID. Only the fields that are provided will be updated.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user profile (partial update)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Profile Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/profile.UpdateProfileInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile updated successfully",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "login.LoginInput": {
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
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "musicalinfo.MusicalInfoInput": {
            "type": "object",
            "required": [
                "primaryInstrument",
                "skillLevel"
            ],
            "properties": {
                "favoriteArtists": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "genres": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "learningGoals": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "primaryInstrument": {
                    "type": "string"
                },
                "secondaryInstruments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "skillLevel": {
                    "type": "string"
                }
            }
        },
        "profile.Address": {
            "type": "object",
            "required": [
                "city",
                "country",
                "postalCode",
                "state"
            ],
            "properties": {
                "city": {
                    "type": "string",
                    "minLength": 3
                },
                "country": {
                    "type": "string"
                },
                "postalCode": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "profile.UpdateProfileInput": {
            "type": "object",
            "required": [
                "address",
                "birthday",
                "firstName",
                "gender",
                "lastName",
                "phone"
            ],
            "properties": {
                "address": {
                    "$ref": "#/definitions/profile.Address"
                },
                "bio": {
                    "type": "string"
                },
                "birthday": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "gender": {
                    "type": "string",
                    "enum": [
                        "male",
                        "female"
                    ]
                },
                "lastName": {
                    "type": "string"
                },
                "nationality": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "profilePicture": {
                    "type": "string"
                }
            }
        },
        "register.RegisterInput": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7890",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Talentgrowth API",
	Description:      "This is the API documentation for Talentgrowth",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
