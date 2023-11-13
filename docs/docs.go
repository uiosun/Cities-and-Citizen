// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://www.uiosun.com/policy/privacy",
        "contact": {
            "name": "C\u0026C API Support",
            "email": "uiosun@outlook.com"
        },
        "license": {
            "name": "GPL-3.0",
            "url": "https://github.com/Sun-FreePort/Cities-and-Citizen/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {}
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:22042",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Cities-and-Citizen API",
	Description:      "Wish nice day for you! 👋",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
