// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/search": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "查询任务",
                "parameters": [
                    {
                        "description": "2",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.DeleteTaskService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    },
                    "500": {
                        "description": "status\":500,\"data\":{},\"Msg\":{},\"Error\":\"error\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/task": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "修改任务",
                "parameters": [
                    {
                        "description": "2",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.DeleteTaskService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    },
                    "500": {
                        "description": "status\":500,\"data\":{},\"Msg\":{},\"Error\":\"error\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "创建任务",
                "parameters": [
                    {
                        "description": "title",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.CreateTaskService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/serializer.ResponseTask"
                        }
                    },
                    "500": {
                        "description": "status\":500,\"data\":{},\"Msg\":{},\"Error\":\"error\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/task/:id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "展示任务详细信息",
                "parameters": [
                    {
                        "description": "rush",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ShowTaskService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/serializer.ResponseTask"
                        }
                    },
                    "500": {
                        "description": "status\":500,\"data\":{},\"Msg\":{},\"Error\":\"error\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "删除任务",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.DeleteTaskService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    },
                    "500": {
                        "description": "status\":500,\"data\":{},\"Msg\":{},\"Error\":\"error\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "获取任务列表",
                "parameters": [
                    {
                        "description": "rush",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ListTasksService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/serializer.ResponseTask"
                        }
                    },
                    "500": {
                        "description": "status\":500,\"data\":{},\"Msg\":{},\"Error\":\"error\"}",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "user_name, password",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.UserService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"登陆成功\"}",
                        "schema": {
                            "$ref": "#/definitions/serializer.ResponseUser"
                        }
                    },
                    "500": {
                        "description": "{\"status\":500,\"data\":{},\"Msg\":{},\"Error\":\"error\"}",
                        "schema": {
                            "$ref": "#/definitions/serializer.ResponseUser"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户名, 密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.UserService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"status\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "$ref": "#/definitions/serializer.ResponseUser"
                        }
                    },
                    "500": {
                        "description": "{\"status\":500,\"data\":{},\"Msg\":{},\"Error\":\"error\"}",
                        "schema": {
                            "$ref": "#/definitions/serializer.ResponseUser"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serializer.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "serializer.ResponseTask": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/serializer.Task"
                },
                "error": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "serializer.ResponseUser": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/serializer.User"
                },
                "error": {
                    "type": "string",
                    "example": ""
                },
                "msg": {
                    "type": "string",
                    "example": "ok"
                },
                "status": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "serializer.Task": {
            "type": "object",
            "properties": {
                "content": {
                    "description": "内容",
                    "type": "string",
                    "example": "睡觉"
                },
                "created_at": {
                    "type": "integer"
                },
                "end_time": {
                    "type": "integer"
                },
                "id": {
                    "description": "任务ID",
                    "type": "integer",
                    "example": 1
                },
                "start_time": {
                    "type": "integer"
                },
                "status": {
                    "description": "状态(0未完成，1已完成)",
                    "type": "integer",
                    "example": 0
                },
                "title": {
                    "description": "题目",
                    "type": "string",
                    "example": "吃饭"
                },
                "view": {
                    "description": "浏览量",
                    "type": "integer",
                    "example": 32
                }
            }
        },
        "serializer.User": {
            "type": "object",
            "properties": {
                "create_at": {
                    "description": "创建",
                    "type": "integer"
                },
                "id": {
                    "description": "用户ID",
                    "type": "integer",
                    "example": 1
                },
                "user_name": {
                    "description": "用户名",
                    "type": "string",
                    "example": "FanOne"
                }
            }
        },
        "service.CreateTaskService": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string",
                    "maxLength": 1000
                },
                "status": {
                    "description": "0 待办   1已完成",
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                }
            }
        },
        "service.DeleteTaskService": {
            "type": "object"
        },
        "service.ListTasksService": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "start": {
                    "type": "integer"
                }
            }
        },
        "service.ShowTaskService": {
            "type": "object"
        },
        "service.UserService": {
            "type": "object",
            "required": [
                "password",
                "user_name"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 5,
                    "example": "FanOne666"
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 15,
                    "minLength": 3,
                    "example": "FanOne"
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
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "ToDoList API",
	Description:      "This is a sample Server pets",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
