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
        "/edits": {
            "post": {
                "description": "Alter model parameters for text completion models.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "Modify model parameters",
                "parameters": [
                    {
                        "description": "Define which model to modify.",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Define the initial prompt for the model.",
                        "name": "instruction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Initial input prompt for model.",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Define stop words for the model to reply after.",
                        "name": "stop",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.OpenAIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    }
                }
            }
        },
        "/models": {
            "get": {
                "description": "List all currently loaded models.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "Get a list of all currently loaded models",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.OpenAIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    }
                }
            }
        },
        "/v1/chat/completions": {
            "post": {
                "description": "Generates text completions based on the provided prompt and previous messages, using a pre-trained language model.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "Generate Model-based Text Completions",
                "parameters": [
                    {
                        "description": "The name of the pre-trained language model to use for generating text completions.",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The list of previous messages exchanged with the language model, including the user's messages and the model's responses.",
                        "name": "messages",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Message"
                            }
                        }
                    },
                    {
                        "default": 0.5,
                        "description": "The sampling temperature to use when generating text completions. Must be between 0 and 1. Higher values result in more diverse completions, while lower values result in more conservative completions.",
                        "name": "temperature",
                        "in": "body",
                        "schema": {
                            "type": "number"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.OpenAIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.APIError"
                        }
                    }
                }
            }
        },
        "/v1/completions": {
            "post": {
                "description": "Allows to generate completions for a given prompt and model.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Chat completions.",
                "parameters": [
                    {
                        "description": "The prompt to generate completions for.",
                        "name": "prompt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The ID of the model to use.",
                        "name": "model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The maximum number of tokens to generate in the completion.",
                        "name": "max_tokens",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "How many completions to generate for each prompt.",
                        "name": "n",
                        "in": "body",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "description": "The sampling temperature to use when generating completions.",
                        "name": "temperature",
                        "in": "body",
                        "schema": {
                            "type": "number"
                        }
                    },
                    {
                        "description": "Sequence where the API will stop generating further tokens",
                        "name": "stop",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.APIError": {
            "description": "Error returned by the API",
            "type": "object",
            "properties": {
                "code": {},
                "message": {
                    "type": "string"
                },
                "param": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "api.Choice": {
            "type": "object",
            "properties": {
                "delta": {
                    "$ref": "#/definitions/api.Message"
                },
                "finish_reason": {
                    "type": "string"
                },
                "index": {
                    "type": "integer"
                },
                "message": {
                    "$ref": "#/definitions/api.Message"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "api.Message": {
            "description": "Message with a content and a role",
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "api.OpenAIResponse": {
            "type": "object",
            "properties": {
                "choices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Choice"
                    }
                },
                "created": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "object": {
                    "type": "string"
                },
                "usage": {
                    "$ref": "#/definitions/api.OpenAIUsage"
                }
            }
        },
        "api.OpenAIUsage": {
            "type": "object",
            "properties": {
                "completion_tokens": {
                    "type": "integer"
                },
                "prompt_tokens": {
                    "type": "integer"
                },
                "total_tokens": {
                    "type": "integer"
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