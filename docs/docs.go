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
        "termsOfService": "https://github.com/zzlpeter/dawn-go",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/assassin/account/permission": {
            "put": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Update special permission",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "role_id",
                        "name": "role_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "permission_id",
                        "name": "permission_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "permission_desc",
                        "name": "permission_desc",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Add special permission",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role_id",
                        "name": "role_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "permission_id",
                        "name": "permission_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "permission_desc",
                        "name": "permission_desc",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Delete special permission",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            }
        },
        "/assassin/account/permissions": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get permissions list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "size",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "role_id",
                        "name": "role_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "permission_id",
                        "name": "permission_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "permission_desc",
                        "name": "permission_desc",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            }
        },
        "/assassin/account/role": {
            "put": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Update special role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "role_id",
                        "name": "role_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "role_desc",
                        "name": "role_desc",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Add special role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "role_id",
                        "name": "role_id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "role_desc",
                        "name": "role_desc",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Delete special role",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            }
        },
        "/assassin/account/roles": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get roles list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "size",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "role_id",
                        "name": "role_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "role_desc",
                        "name": "role_desc",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            }
        },
        "/assassin/cas/": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get Auth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "userName",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            }
        },
        "/assassin/task/task": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get special task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Update special task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "任务ID",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "任务key",
                        "name": "task_key",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "执行方法",
                        "name": "execute_func",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "任务描述",
                        "name": "desc",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "调度类型",
                        "name": "trigger",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "调度频次",
                        "name": "spec",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "执行参数",
                        "name": "args",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "是否有效",
                        "name": "is_valid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "额外信息(json格式)",
                        "name": "extra",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "执行状态",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Add special task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "任务key",
                        "name": "task_key",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "执行方法",
                        "name": "execute_func",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "任务描述",
                        "name": "desc",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "调度类型",
                        "name": "trigger",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "调度频次",
                        "name": "spec",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "执行参数",
                        "name": "args",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "是否有效",
                        "name": "is_valid",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "执行状态",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "额外信息",
                        "name": "extra",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            }
        },
        "/assassin/task/tasks": {
            "get": {
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tasks"
                ],
                "summary": "Get tasks list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "size",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.SwaggerResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.SwaggerResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "message": {
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
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Dawn-go API By Gin",
	Description: "Dawn-go API By Golang Gin",
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