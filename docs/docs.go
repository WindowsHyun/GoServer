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
        "/app/menu": {
            "get": {
                "description": "App Menu",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "App"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structure.ResDefaultMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request Error"
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "User Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "User login request body",
                        "name": "reqBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structure.ReqLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structure.ResLogin"
                        }
                    },
                    "400": {
                        "description": "Bad Request Error"
                    }
                }
            }
        },
        "/user/regist": {
            "post": {
                "description": "User Regist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "parameters": [
                    {
                        "description": "User registration request body",
                        "name": "reqBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structure.ReqRegist"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structure.ResRegist"
                        }
                    },
                    "400": {
                        "description": "Bad Request Error"
                    },
                    "502": {
                        "description": "Could Not Be Searched In DB Collection"
                    },
                    "506": {
                        "description": "DB Collection Update Error"
                    },
                    "509": {
                        "description": "JWT Token Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "structure.ReqLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "structure.ReqRegist": {
            "type": "object",
            "properties": {
                "birthMonth": {
                    "type": "string"
                },
                "birthYear": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                }
            }
        },
        "structure.ResDefaultMessage": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "structure.ResLogin": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "structure.ResRegist": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
