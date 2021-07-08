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
        "termsOfService": "http://github.com/greetlist/",
        "contact": {
            "name": "Greetlist",
            "url": "http://github.com/greetlist",
            "email": "lrt12250914@outlook.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/stock/getAllStockCode": {
            "get": {
                "description": "Return All Stock Code",
                "produces": [
                    "application/json"
                ],
                "summary": "Query All Stock Code",
                "operationId": "getAllStockCode",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetAllStockCodeResponse"
                        }
                    }
                }
            }
        },
        "/api/stock/getDailyCalcStockData": {
            "get": {
                "description": "Return All Daily Calc Stock Data",
                "produces": [
                    "application/json"
                ],
                "summary": "Query Daily Calc Stock Data",
                "operationId": "getDailyCalcStockData",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetDailyCalcStockDataResponse"
                        }
                    }
                }
            }
        },
        "/api/stock/getQueryStockData": {
            "post": {
                "description": "Return Strategy Sepcific Computed Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query Sepcific Stock Computed Data",
                "operationId": "getQueryStockData",
                "parameters": [
                    {
                        "description": "Query Stock List",
                        "name": "request_json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetQueryStockDataRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetQueryStockDataResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.GetAllStockCodeResponse": {
            "type": "object",
            "properties": {
                "stock_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "['000001'",
                        " '002142'",
                        " ...]"
                    ]
                }
            }
        },
        "model.GetDailyCalcStockDataResponse": {
            "type": "object",
            "properties": {
                "stock_datas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.StockDataItem"
                    }
                }
            }
        },
        "model.GetQueryStockDataRequest": {
            "type": "object",
            "properties": {
                "stock_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.GetQueryStockDataResponse": {
            "type": "object",
            "properties": {
                "stock_datas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.StockDataItem"
                    }
                }
            }
        },
        "model.SingleRecord": {
            "type": "object",
            "properties": {
                "close": {
                    "type": "number"
                },
                "date": {
                    "type": "string"
                },
                "high": {
                    "type": "number"
                },
                "low": {
                    "type": "number"
                },
                "ma13": {
                    "type": "number"
                },
                "ma34": {
                    "type": "number"
                },
                "ma55": {
                    "type": "number"
                },
                "money": {
                    "type": "number"
                },
                "open": {
                    "type": "number"
                },
                "volume": {
                    "type": "number"
                }
            }
        },
        "model.StockDataItem": {
            "type": "object",
            "properties": {
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SingleRecord"
                    }
                },
                "stock_code": {
                    "type": "string",
                    "example": "002142"
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
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Stock-Show Web API",
	Description: "server API for stock web",
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
