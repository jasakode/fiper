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

// RESPONSE FOUND
func (r *response) Found(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("Found")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}
func (r *response) Modified(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("Modified")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

// RESPONSE REDIRECTING
func (r *response) Redirect(ctx *fiber.Ctx, url string) error {
	v, err := fiperPlugin.Lookup("Redirect")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, url string) error)(ctx, url)
	return result
}

// RESPONSE ERROR
func (r *response) BadRequest(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("BadRequest")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) Unauthorized(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("Unauthorized")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) PaymentRequired(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("PaymentRequired")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) Forbidden(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("Forbidden")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) NotFound(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("NotFound")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) MethodNotAllowed(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("MethodNotAllowed")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) NotAcceptable(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("NotAcceptable")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) RequestTimeout(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("RequestTimeout")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) LengthRequired(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("LengthRequired")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) RequestEntityTooLarge(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("RequestEntityTooLarge")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) UnsupportedMediaType(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("UnsupportedMediaType")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) UnprocessableEntity(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("UnprocessableEntity")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) Locked(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("Locked")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) UpgradeRequired(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("UpgradeRequired")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

func (r *response) TooManyRequests(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("TooManyRequests")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}

// RESPONSE OTHER
func (r *response) NotImplemented(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("NotImplemented")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}
func (r *response) BadGateway(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("BadGateway")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}
func (r *response) ServiceUnavailable(ctx *fiber.Ctx, message string, data ...interface{}) error {
	v, err := fiperPlugin.Lookup("ServiceUnavailable")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string, data ...interface{}) error)(ctx, message, data...)
	return result
}
func (r *response) InternalServerError(ctx *fiber.Ctx, message string) error {
	v, err := fiperPlugin.Lookup("InternalServerError")
	if err != nil {
		panic(err)
	}
	result := v.(func(ctx *fiber.Ctx, message string) error)(ctx, message)
	return result
}

// var Database
// var Storage
// var Sessions
// var Request
