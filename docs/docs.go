// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/": {
            "get": {
                "description": "获取所有图书信息",
                "tags": [
                    "library"
                ],
                "summary": "获取所有图书信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "需要查询的页数",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tools.HttpCode"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dao.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tools.HttpCode"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dao.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/log/login": {
            "post": {
                "description": "用户登录",
                "tags": [
                    "login"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "pwd",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            }
        },
        "/log/logout": {
            "get": {
                "description": "用户退出登录",
                "tags": [
                    "login"
                ],
                "summary": "用户退出登录",
                "responses": {}
            }
        },
        "/log/register": {
            "post": {
                "description": "用户注册",
                "tags": [
                    "login"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "pwd",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "0:普通用户，1:超级用户",
                        "name": "type",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            }
        },
        "/manager": {
            "get": {
                "description": "查看所有借阅记录",
                "tags": [
                    "manager"
                ],
                "summary": "查看所有借阅记录",
                "parameters": [
                    {
                        "type": "string",
                        "default": "123",
                        "description": "Debug header",
                        "name": "Debug",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            }
        },
        "/manager/book": {
            "post": {
                "description": "管理员添加图书信息",
                "tags": [
                    "manager"
                ],
                "summary": "添加图书",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "图书ID",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "书名",
                        "name": "book_name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "作者",
                        "name": "author",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "出版社",
                        "name": "publishing_house",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "译者",
                        "name": "translator",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "日期：2002-03-02",
                        "name": "publish_data",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "页数",
                        "name": "pages",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "isbn编号",
                        "name": "isbn",
                        "in": "formData"
                    },
                    {
                        "type": "number",
                        "description": "价格",
                        "name": "price",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "图书简介",
                        "name": "brief_introduction",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "作者简介",
                        "name": "author_introduction",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "图片链接",
                        "name": "img_url",
                        "in": "formData"
                    },
                    {
                        "type": "integer",
                        "description": "图书数量",
                        "name": "count",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            }
        },
        "/manager/book/{book_id}": {
            "delete": {
                "description": "管理员删除指定书名的图书信息",
                "tags": [
                    "manager"
                ],
                "summary": "删除图书信息",
                "parameters": [
                    {
                        "type": "string",
                        "default": "123",
                        "description": "Debug header",
                        "name": "Debug",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "书名",
                        "name": "book_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            }
        },
        "/manager/user": {
            "get": {
                "description": "查看所有用户信息",
                "tags": [
                    "manager"
                ],
                "summary": "查看所有用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "default": "123",
                        "description": "Debug header",
                        "name": "Debug",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            }
        },
        "/manager/user/{id}": {
            "get": {
                "description": "查看指定用户的借阅信息",
                "tags": [
                    "manager"
                ],
                "summary": "查看指定用户的借阅信息",
                "parameters": [
                    {
                        "type": "string",
                        "default": "123",
                        "description": "Debug header",
                        "name": "Debug",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "用户Id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "查看用户信息",
                "tags": [
                    "user"
                ],
                "summary": "查看用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "default": "123",
                        "description": "Debug header",
                        "name": "Debug",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            },
            "put": {
                "description": "修改用户信息",
                "tags": [
                    "user"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "default": "123",
                        "description": "Debug header",
                        "name": "Debug",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "pwd",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            }
        },
        "/user/book": {
            "get": {
                "description": "查看自己的借阅记录",
                "tags": [
                    "user"
                ],
                "summary": "查看自己的借阅记录",
                "parameters": [
                    {
                        "type": "string",
                        "default": "123",
                        "description": "Debug header",
                        "name": "Debug",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            }
        },
        "/user/book/{book_id}": {
            "get": {
                "description": "归还指定书名的图书",
                "tags": [
                    "user"
                ],
                "summary": "归还图书",
                "parameters": [
                    {
                        "type": "string",
                        "default": "123",
                        "description": "Debug",
                        "name": "Debug",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "图书ID",
                        "name": "book_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            },
            "post": {
                "description": "借阅指定书名的图书",
                "tags": [
                    "user"
                ],
                "summary": "借阅图书",
                "parameters": [
                    {
                        "type": "string",
                        "default": "123",
                        "description": "Debug header",
                        "name": "Debug",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "图书ID",
                        "name": "book_id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/tools.HttpCode"
                        }
                    }
                }
            }
        },
        "/{book_name}": {
            "get": {
                "description": "通过模糊查询查看图书信息学",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "library"
                ],
                "summary": "获取图书信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "书名",
                        "name": "book_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "页数",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tools.HttpCode"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dao.Book"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tools.HttpCode"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dao.Book"
                                            }
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
        "dao.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "description": "作者",
                    "type": "string"
                },
                "author_introduction": {
                    "description": "作者简介",
                    "type": "string"
                },
                "book_name": {
                    "description": "书名",
                    "type": "string"
                },
                "brief_introduction": {
                    "description": "内容简介",
                    "type": "string"
                },
                "count": {
                    "description": "库存",
                    "type": "integer"
                },
                "del_flg": {
                    "description": "删除标识",
                    "type": "integer"
                },
                "id": {
                    "description": "书的ID",
                    "type": "integer"
                },
                "img_url": {
                    "description": "封面地址",
                    "type": "string"
                },
                "isbn": {
                    "description": "ISBN号码",
                    "type": "string"
                },
                "pages": {
                    "description": "页数",
                    "type": "integer"
                },
                "price": {
                    "description": "价格",
                    "type": "number"
                },
                "publish_date": {
                    "description": "出版时间",
                    "type": "string"
                },
                "publishing_house": {
                    "description": "出版社",
                    "type": "string"
                },
                "translator": {
                    "description": "译者",
                    "type": "string"
                }
            }
        },
        "tools.HttpCode": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "图书馆管理系统",
	Description:      "该项目用户设有游客，普通用户和管理员，但是只有普通用户和管理员可以借阅书籍",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
