package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://developers.ctd.com.ar/support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/Odontologos": {
            "post": {
                "description": "create Odontologo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologos"
                ],
                "summary": "Create Odontologo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Odontologo to create",
                        "name": "domain.Odontologo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Odontologo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Odontologo successfully created",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "409": {
                        "description": "error: the Odontologo already exist",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "422": {
                        "description": "error: ¡incomplete fields!",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "500": {
                        "description": "error while saving",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/Odontologos/{id}": {
            "get": {
                "description": "get Odontologo by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologos"
                ],
                "summary": "Get Odontologo by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "update all Odontologo by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologos"
                ],
                "summary": "Update all Odontologo by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Odontologo to update",
                        "name": "domain.Odontologo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Odontologo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete Odontologo by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologos"
                ],
                "summary": "Delete Odontologo by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "update Odontologo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Odontologos"
                ],
                "summary": "Update Odontologo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Odontologo to update",
                        "name": "domain.Odontologo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Odontologo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/Pacientes": {
            "post": {
                "description": "create Paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pacientes"
                ],
                "summary": "Create Paciente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Paciente to create",
                        "name": "domain.Paciente",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Paciente successfully created",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "409": {
                        "description": "error: the Paciente already exist",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "422": {
                        "description": "error: ¡incomplete fields!",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "500": {
                        "description": "error while saving",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/Pacientes/{id}": {
            "get": {
                "description": "get Paciente by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pacientes"
                ],
                "summary": "Get Paciente by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete Paciente by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pacientes"
                ],
                "summary": "Delete Paciente by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "update Paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pacientes"
                ],
                "summary": "Update Paciente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Paciente to update",
                        "name": "domain.Paciente",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Paciente"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/Turnos": {
            "post": {
                "description": "create Turno",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turnos"
                ],
                "summary": "Create Turno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Turno to create",
                        "name": "domain.Turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Turno"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Turno successfully created",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "409": {
                        "description": "error: the Turno already exist",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "422": {
                        "description": "error: ¡incomplete fields!",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "500": {
                        "description": "error while saving",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/Turnos/bydni": {
            "get": {
                "description": "get Turno by DNI Paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turnos"
                ],
                "summary": "Get Turno by DNI Paciente",
                "parameters": [
                    {
                        "type": "string",
                        "description": "dni",
                        "name": "dni",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/Turnos/{id}": {
            "get": {
                "description": "get Turno by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turnos"
                ],
                "summary": "Get Turno by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "update all Turno by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turnos"
                ],
                "summary": "Update all Turno by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Turno to update",
                        "name": "domain.Turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Turno"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete Turno by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turnos"
                ],
                "summary": "Delete Turno by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "patch": {
                "description": "update Turno",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turnos"
                ],
                "summary": "Update Turno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Turno to update",
                        "name": "domain.Turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Turno"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    },
                    "404": {
                        "description": "not found",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Odontologo": {
            "type": "object",
            "properties": {
                "enrollment": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "domain.Paciente": {
            "type": "object",
            "properties": {
                "DNI": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "discharge_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "domain.Turno": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "Odontologo_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "Paciente_id": {
                    "type": "integer"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "web.Response": {
            "type": "object",
            "properties": {
                "data": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Odontologo - Final",
	Description:      "This API Handle Products.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}