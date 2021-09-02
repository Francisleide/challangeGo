// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "get": {
                "description": "List all accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Account"
                            }
                        }
                    },
                    "400": {
                        "description": "Failed to decode"
                    },
                    "404": {
                        "description": "Accounts not found"
                    },
                    "500": {
                        "description": "Unexpected internal server error"
                    }
                }
            },
            "post": {
                "description": "Create an account with the basic information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create an account",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.AccountInput"
                        }
                    }
                ]
            }
        },
        "/accounts/{id}/balance": {
            "get": {
                "description": "show the balance of a specific account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "account balance",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/account.AccountBalance"
                        }
                    },
                    "400": {
                        "description": "Failed to decode"
                    },
                    "404": {
                        "description": "Account not found"
                    },
                    "500": {
                        "description": "Unexpected internal server error"
                    }
                }
            }
        },
        "/deposit": {
            "post": {
                "description": "Make a deposit from an authentic account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Make a deposit",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.DepositInput"
                        }
                    }
                ]
            }
        },
        "/login": {
            "post": {
                "description": "It takes a token to authenticate yourself to the application",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get a Auth",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.Login"
                        }
                    }
                ]
            }
        },
        "/transfer": {
            "post": {
                "description": "Transfer between accounts. The account that will make the transfer must be authenticated with a token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Make a transfer",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.Transfer"
                            }
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer Authorization Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ]
            }
        },
        "/withdraw": {
            "post": {
                "description": "Make a Withdraw from an authentic account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Make a Withdraw",
                "parameters": [
                    {
                        "description": "Body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/account.Withdraw"
                        }
                    }
                ]
            }
        }
    },
    "definitions": {
        "account.AccountBalance": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                }
            }
        },
        "account.DepositInput": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                }
            }
        },
        "account.Withdraw": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                }
            }
        },
        "auth.Login": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                }
            }
        },
        "entities.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "cpf": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                }
            }
        },
        "entities.AccountInput": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                }
            }
        },
        "entities.Transfer": {
            "type": "object",
            "properties": {
                "accountDestinationID": {
                    "type": "string"
                },
                "accountOriginID": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "entities.TransferInput": {
            "type": "object",
            "properties": {
                "accountDestinationID": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
