package helper

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var validate = validator.New()

// CustomValidator คือ struct ที่ใช้เพื่อสร้าง custom validator
type CustomValidator struct {
	validator *validator.Validate
}

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse represents the structure of a JSON error response
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func NewSuccessMessage() *Message {
	return &Message{
		Code:    200,
		Message: "Success",
	}
}

func RespJson(c *fiber.Ctx, fn interface{}, input interface{}) error {
	// ตรวจสอบและดึงข้อมูลจาก request ตาม content type
	if err := parseInputData(c, input); err != nil {
		return RenderJSON(c, err, nil)
	}

	// ตรวจสอบความถูกต้องของข้อมูล
	if err := validateInput(input); err != nil {
		return RenderJSON(c, err, nil)
	}

	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(c),
		reflect.ValueOf(input),
	})

	errObj := out[1].Interface()
	if errObj != nil {
		logrus.Errorf("call service error: %s", errObj)
		return RenderJSON(c, errObj.(error), nil)
	}

	ResponseResult := ResponseResult{
		Code:    200,
		Message: "Success",
		Result:  out[0].Interface(),
	}

	return RenderJSON(c, nil, ResponseResult)
}

func RespJsonNoReq(c *fiber.Ctx, fn interface{}) error {
	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(c),
	})

	errObj := out[1].Interface()
	if errObj != nil {
		logrus.Errorf("call service error: %s", errObj)
		return RenderJSON(c, errObj.(error), nil)
	}

	ResponseResult := ResponseResult{
		Code:    200,
		Message: "Success",
		Result:  out[0].Interface(),
	}

	return RenderJSON(c, nil, ResponseResult)
}

func RespSuccess(c *fiber.Ctx, fn interface{}, input interface{}) error {
	// ตรวจสอบและดึงข้อมูลจาก request ตาม content type
	if err := parseInputData(c, input); err != nil {
		return RenderJSON(c, err, nil)
	}

	// ตรวจสอบความถูกต้องของข้อมูล
	if err := validateInput(input); err != nil {
		return RenderJSON(c, err, nil)
	}

	out := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(c),
		reflect.ValueOf(input),
	})

	errObj := out[0].Interface()
	if errObj != nil {
		logrus.Errorf("call service error: %s", errObj)
		return RenderJSON(c, errObj.(error), nil)
	}

	return RenderJSON(c, nil, NewSuccessMessage())
}

// ฟังก์ชันนี้ใช้ในการเช็คและดึงข้อมูลจาก request ตาม content type
func parseInputData(c *fiber.Ctx, input interface{}) error {
	Method := c.Method()
	switch {
	case strings.HasPrefix(Method, "POST"), strings.HasPrefix(c.Method(), "DELETE"):
		// เรียกใช้ฟังก์ชันแปลง
		if err := parseRequestBody(c, &input); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
	case strings.HasPrefix(Method, "GET"):
		if err := c.QueryParser(input); err != nil {
			return err
		}
	case c.Is("multipart/form-data"):
		if err := mapFormValues(c, input); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported content type: %s", Method)
	}
	return nil
}

func mapFormValues(c *fiber.Ctx, input interface{}) error {
	val := reflect.ValueOf(input).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name

		value := c.FormValue(fieldName)
		if value != "" {
			// If the field is a slice, assign the value as a single element
			if field.Type().Kind() == reflect.Slice {
				slice := reflect.MakeSlice(field.Type(), 1, 1)
				slice.Index(0).SetString(value)
				field.Set(slice)
			} else {
				field.SetString(value)
			}
		}
	}

	return nil
}

func parseRequestBody(c *fiber.Ctx, inputStruct interface{}) error {
	// ใช้คำสั่ง ShouldBind ในที่นี้
	if err := c.BodyParser(inputStruct); err != nil {
		return fmt.Errorf("failed to parse request body: %v", err)
	}

	return nil
}

// ฟังก์ชันที่จัดการ error response และ success response
func RenderJSON(c *fiber.Ctx, err error, successResponse interface{}) error {
	if err != nil {
		// Create an ErrorResponse instance
		errorResponse := ErrorResponse{
			Code:    err.(*fiber.Error).Code,
			Message: err.Error(),
		}

		// Set the HTTP status code
		c.Status(err.(*fiber.Error).Code)

		// Return the error response as JSON
		return c.JSON(errorResponse)
	}

	// Return the success response as JSON
	return c.JSON(successResponse)
}

// ValidateDate คือ custom validation function สำหรับการตรวจสอบวันที่
func ValidateDate(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	// ตรวจสอบว่าค่าวันที่ไม่เป็นค่าว่าง
	if dateStr == "" {
		return true
	}
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}

// NewValidator คือฟังก์ชั่นที่ใช้สร้าง custom validator
func NewValidator() *CustomValidator {
	v := validator.New()
	v.RegisterValidation("date", ValidateDate)

	return &CustomValidator{validator: v}
}

// ฟังก์ชันนี้ใช้ในการตรวจสอบความถูกต้องของข้อมูล
func validateInput(input interface{}) error {
	// สร้าง custom validator
	validate := NewValidator()

	// ใช้ custom validator เพื่อ validate ข้อมูล
	if err := validate.validator.Struct(input); err != nil {
		// กรณีมี error ในการ validate แสดงข้อความเพิ่มเติม
		errs := err.(validator.ValidationErrors)
		errorMsg := "Invalid request data:"
		for _, e := range errs {
			errorMsg += fmt.Sprintf("\n- Field: %s, Type: %T, Error: %s", e.Field(), e.Value(), e.Tag())
		}
		// ใช้ fiber.NewError เพื่อสร้าง error ที่เข้ากับ Fiber framework
		return fiber.NewError(fiber.StatusBadRequest, errorMsg)
	}
	return nil
}
