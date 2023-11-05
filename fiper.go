package fiper

import "github.com/gofiber/fiber/v2"

// RESPONSE
type response struct{}

// Include all Methote to Response Variable
var Response response

func (r *response) CreateJSON() map[string]interface{} {
	v, err := fiperPlugin.Lookup("CreateJSON")
	if err != nil {
		panic(err)
	}
	result := v.(func() map[string]interface{})()
	return result
}

func (r *response) Send(ctx *fiber.Ctx, code int, message string, data ...map[string]interface{}) error {
	v, err := fiperPlugin.Lookup("Send")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, code int, message string, data ...map[string]interface{}) error)(ctx, code, message, data...)
	return result
}

func (r *response) Processing(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("Processing")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) Continue(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("Continue")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

// response Success

func (r *response) Ok(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("Ok")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) Created(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("Created")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) Accepted(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("Accepted")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) NoContent(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("NoContent")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) IMUsed(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("IMUsed")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

// var Database
// var Storage
// var Sessions
// var Request
