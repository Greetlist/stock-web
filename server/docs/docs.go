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
        "/api/auth/login": {
            "post": {
                "description": "User Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User Login",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "login",
                        "name": "request_json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResponse"
                        }
                    }
                }
            }
        },
        "/api/auth/logout": {
            "post": {
                "description": "User Logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User Logout",
                "operationId": "logout",
                "parameters": [
                    {
                        "description": "Logout",
                        "name": "request_json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LogoutRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.LogoutResponse"
                        }
                    }
                }
            }
        },
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
                            "$ref": "#/definitions/model.MarketRawData"
                        }
                    }
                }
            }
        },
        "/api/stock/getMarketDistribution": {
            "post": {
                "description": "Return Market Distribution",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query Market Distribution",
                "operationId": "getMarketDistribution",
                "parameters": [
                    {
                        "description": "Query Index ",
                        "name": "request_json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetMarketDistributionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetMarketDistributionResponse"
                        }
                    }
                }
            }
        },
        "/api/stock/getRecommandStockPrediction": {
            "post": {
                "description": "Return Recommand Stock Prediction Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query Recommand Stock Prediction Data",
                "operationId": "getRecommandStockPrediction",
                "parameters": [
                    {
                        "description": "Query Date",
                        "name": "request_json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetRecommandStockRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetRecommandStockResponse"
                        }
                    }
                }
            }
        },
        "/api/stock/getTotalMarketIndexData": {
            "post": {
                "description": "Return Strategy Sepcific Computed Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query Sepcific All Index Computed Data",
                "operationId": "getTotalMarketIndexData",
                "parameters": [
                    {
                        "description": "Query Index ",
                        "name": "request_json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetTotalMarketIndexDataRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetTotalMarketIndexDataResponse"
                        }
                    }
                }
            }
        },
        "/api/stock/queryStockData": {
            "post": {
                "description": "Return Strategy Sepcific Computed Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query Sepcific Stock Computed Data",
                "operationId": "queryStockData",
                "parameters": [
                    {
                        "description": "Query Stock List",
                        "name": "request_json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.QueryStockDataRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.QueryStockDataResponse"
                        }
                    }
                }
            }
        },
        "/api/user/addPreferStock": {
            "post": {
                "description": "User AddPreferStock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User AddPreferStock",
                "operationId": "addPreferStock",
                "parameters": [
                    {
                        "description": "login",
                        "name": "request_json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddPreferStockRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AddPreferStockResponse"
                        }
                    }
                }
            }
        },
        "/api/user/deletePreferStock": {
            "post": {
                "description": "User DeletePreferStock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User DeletePreferStock",
                "operationId": "deletePreferStock",
                "parameters": [
                    {
                        "description": "login",
                        "name": "request_json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DeletePreferStockRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.DeletePreferStockResponse"
                        }
                    }
                }
            }
        },
        "/api/user/getPreferStock": {
            "post": {
                "description": "User GetPreferStock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "User GetPreferStock",
                "operationId": "getPreferStock",
                "parameters": [
                    {
                        "description": "login",
                        "name": "request_json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.GetPreferStockRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetPreferStockResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AddPreferStockRequest": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "stock_code": {
                    "type": "string"
                }
            }
        },
        "model.AddPreferStockResponse": {
            "type": "object",
            "properties": {
                "error_msg": {
                    "type": "string"
                },
                "is_success": {
                    "type": "boolean"
                }
            }
        },
        "model.DeletePreferStockRequest": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "stock_code": {
                    "type": "string"
                }
            }
        },
        "model.DeletePreferStockResponse": {
            "type": "object",
            "properties": {
                "error_msg": {
                    "type": "string"
                },
                "is_success": {
                    "type": "boolean"
                }
            }
        },
        "model.DistributionData": {
            "type": "object",
            "properties": {
                "large": {
                    "type": "integer"
                },
                "mid": {
                    "type": "integer"
                },
                "small": {
                    "type": "integer"
                }
            }
        },
        "model.FactorRawData": {
            "type": "object",
            "properties": {
                "ema13": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "ema34": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "ema55": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "ma13": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "ma34": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "ma55": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                }
            }
        },
        "model.FollowStock": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "exchange": {
                    "type": "string"
                },
                "industry": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "stock_code": {
                    "type": "string"
                }
            }
        },
        "model.GetAllStockCodeResponse": {
            "type": "object",
            "properties": {
                "stock_info_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.StockBasicInfo"
                    }
                }
            }
        },
        "model.GetMarketDistributionRequest": {
            "type": "object",
            "properties": {
                "query_date": {
                    "type": "string",
                    "example": "2021-07-01"
                }
            }
        },
        "model.GetMarketDistributionResponse": {
            "type": "object",
            "properties": {
                "distribution_data_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.DistributionData"
                    }
                }
            }
        },
        "model.GetPreferStockRequest": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                }
            }
        },
        "model.GetPreferStockResponse": {
            "type": "object",
            "properties": {
                "error_msg": {
                    "type": "string"
                },
                "follow_stock_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.FollowStock"
                    }
                },
                "is_success": {
                    "type": "boolean"
                }
            }
        },
        "model.GetRecommandStockRequest": {
            "type": "object",
            "properties": {
                "query_data_len": {
                    "type": "integer"
                },
                "query_date": {
                    "type": "string",
                    "example": "2021-07-01"
                },
                "stock_code": {
                    "type": "string"
                }
            }
        },
        "model.GetRecommandStockResponse": {
            "type": "object",
            "properties": {
                "stock_datas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SingleStockData"
                    }
                },
                "stock_prediction_datas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.StockPredictItem"
                    }
                }
            }
        },
        "model.GetTotalMarketIndexDataRequest": {
            "type": "object",
            "properties": {
                "query_date": {
                    "type": "string",
                    "example": "2021-07-01"
                }
            }
        },
        "model.GetTotalMarketIndexDataResponse": {
            "type": "object",
            "properties": {
                "index_data_list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.IndexData"
                    }
                }
            }
        },
        "model.IndexData": {
            "type": "object",
            "properties": {
                "index_name": {
                    "type": "string"
                },
                "index_pred_data": {
                    "$ref": "#/definitions/model.SinglePredictRecord"
                },
                "index_raw_data": {
                    "$ref": "#/definitions/model.SinglePredictRecord"
                },
                "show_msg": {
                    "type": "string"
                }
            }
        },
        "model.LoginRequest": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "passwd": {
                    "type": "string"
                }
            }
        },
        "model.LoginResponse": {
            "type": "object",
            "properties": {
                "err_msg": {
                    "type": "string"
                },
                "login_succ": {
                    "type": "boolean"
                }
            }
        },
        "model.LogoutRequest": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                }
            }
        },
        "model.LogoutResponse": {
            "type": "object",
            "properties": {
                "err_msg": {
                    "type": "string"
                },
                "logout_succ": {
                    "type": "boolean"
                }
            }
        },
        "model.MarketRawData": {
            "type": "object",
            "properties": {
                "close": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "date": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "high": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "low": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "money": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "open": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "volume": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                }
            }
        },
        "model.QueryStockDataRequest": {
            "type": "object",
            "properties": {
                "query_data_len": {
                    "type": "integer"
                },
                "stock_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.QueryStockDataResponse": {
            "type": "object",
            "properties": {
                "stock_datas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SingleStockData"
                    }
                }
            }
        },
        "model.SinglePredictRecord": {
            "type": "object",
            "properties": {
                "close": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "date": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "high": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "low": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "open": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                }
            }
        },
        "model.SingleStockData": {
            "type": "object",
            "properties": {
                "stock_info": {
                    "$ref": "#/definitions/model.StockBasicInfo"
                },
                "stock_raw_datas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.StockSinglePeriodData"
                    }
                }
            }
        },
        "model.StockBasicInfo": {
            "type": "object",
            "properties": {
                "exchange": {
                    "type": "string",
                    "example": "BJSE/SZSE/SSE"
                },
                "stock_code": {
                    "type": "string",
                    "example": "002142.SZ"
                },
                "stock_industry": {
                    "type": "string",
                    "example": "石油"
                },
                "stock_name": {
                    "type": "string",
                    "example": "宁波银行"
                }
            }
        },
        "model.StockPredictItem": {
            "type": "object",
            "properties": {
                "prediction_record": {
                    "$ref": "#/definitions/model.SinglePredictRecord"
                },
                "show_msg": {
                    "type": "string"
                },
                "stock_info": {
                    "$ref": "#/definitions/model.StockBasicInfo"
                }
            }
        },
        "model.StockSinglePeriodData": {
            "type": "object",
            "properties": {
                "factor_raw_data": {
                    "$ref": "#/definitions/model.FactorRawData"
                },
                "market_raw_data": {
                    "$ref": "#/definitions/model.MarketRawData"
                },
                "period_type": {
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
