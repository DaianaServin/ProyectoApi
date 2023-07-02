
{
    "swagger": "2.0",
    "info": {
        "description": "This API Handle Products.",
        "title": "Odontologo - Final"
        "contact": {
            "name": "API Support",
            "url": "https://developers.ctd.com.ar/support"
        },
        "version": "1.0"
    },
    "paths": {
        "/odontologos": {
            "post": {
                "description": "create odontologo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "Create odontologo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "odontologo to create",
                        "name": "domain.odontologo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.odontologo"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "odontologo successfully created",
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
                        "description": "error: the odontologo already exist",
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
        "/odontologos/{id}": {
            "get": {
                "description": "get odontologo by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "Get odontologo by ID",
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
                "description": "update all odontologo by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "Update all odontologo by id",
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
                        "description": "odontologo to update",
                        "name": "domain.odontologo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.odontologo"
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
                "description": "delete odontologo by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "Delete odontologo by id",
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
                "description": "update odontologo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "Update odontologo",
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
                        "description": "odontologo to update",
                        "name": "domain.odontologo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.odontologo"
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
        "/pacientes": {
            "post": {
                "description": "create paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pacientes"
                ],
                "summary": "Create paciente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "paciente to create",
                        "name": "domain.paciente",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.paciente"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "paciente successfully created",
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
                        "description": "error: the paciente already exist",
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
        "/pacientes/{id}": {
            "get": {
                "description": "get paciente by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pacientes"
                ],
                "summary": "Get paciente by ID",
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
                "description": "delete paciente by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pacientes"
                ],
                "summary": "Delete paciente by id",
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
                "description": "update paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pacientes"
                ],
                "summary": "Update paciente",
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
                        "description": "paciente to update",
                        "name": "domain.paciente",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.paciente"
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
                "description": "get Turno by DNI paciente",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turnos"
                ],
                "summary": "Get Turno by DNI paciente",
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
        "domain.odontologo": {
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
        "domain.paciente": {
            "type": "object",
            "properties": {
                "DNI": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "fechaTurno": {
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
                "odontologo_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "paciente_id": {
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
}