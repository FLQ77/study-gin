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
        "/api/image_upload": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "公共"
                ],
                "summary": "图片上传",
                "parameters": [
                    {
                        "description": "params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ImageUpload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/auth/info": {
            "get": {
                "description": "获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "Info",
                "responses": {
                    "200": {
                        "description": "{\"error_code\": 0, \"message\": \"OK\", \"data\": {}}",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"error_code\": 0, \"message\": \"OK\", \"data\": {}}",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "description": "用户登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "{\"error_code\": 0, \"message\": \"OK\", \"data\": {}}",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "params",
                        "name": "params",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Register"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.ImageUpload": {
            "type": "object"
        },
        "request.Login": {
            "type": "object",
            "required": [
                "mobile",
                "password"
            ],
            "properties": {
                "mobile": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.Register": {
            "type": "object",
            "required": [
                "mobile",
                "name",
                "password"
            ],
            "properties": {
                "mobile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "数据"
                },
                "error_code": {
                    "description": "自定义错误码",
                    "type": "integer"
                },
                "message": {
                    "description": "信息",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "api接口文档",
	Description:      "本文档主要是提供app使用的接口, 接口域名：https://localhost:8888/",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}