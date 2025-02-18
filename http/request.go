package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type tagInfo struct {
	Name         string
	DefaultValue string
	Required     bool
	InValues     []string
}

var timeLayouts = []string{
	time.RFC3339,
	time.RFC3339Nano,
	time.RFC1123Z,
	time.RFC1123,
	time.DateTime,
	time.DateOnly,
}

func MarshalByQuery[T any](s *T, finder func(string) string) error {
	return Marshal(finder, s, "query")
}

func MarshalByParam[T any](s *T, finder func(string) string) error {
	return Marshal(finder, s, "param")
}

func MarshalByForm[T any](s *T, finder func(string) string) error {
	return Marshal(finder, s, "form")
}

func Marshal[T any](fn func(string) string, s *T, tagName string) error {

	if s == nil {
		return errors.New("null ptr")
	}

	val := reflect.ValueOf(s)
	{
		if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
			return errors.New("input is not a pointer to a struct")
		}
	}

	if tagName == "" {
		return errors.New("tag name can't be empty")
	}

	val = val.Elem()

	structType := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag.Get(tagName)

		if tag == "" {
			continue
		}

		tagParams := parseTag(tag)

		el := val.Field(i)
		{
			if !el.CanSet() {
				continue
			}
		}

		result := fn(tagParams.Name)

		if result == "" && tagParams.DefaultValue != "" {
			result = tagParams.DefaultValue
		}

		if result == "" && tagParams.Required {
			return fmt.Errorf("field %s is required but was not provided", tagParams.Name)
		}

		if len(tagParams.InValues) > 0 && !isInList(result, tagParams.InValues) {
			return fmt.Errorf("invalid field: %s", tagParams.Name)
		}

		err := setFieldValue(result, el)
		{
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func isInList(value string, list []string) bool {
	for _, item := range list {
		if value == item {
			return true
		}
	}
	return false
}

func parseTag(tag string) (info tagInfo) {
	parts := strings.Split(tag, ",")
	info.Name = parts[0]

	for _, part := range parts[1:] {
		if strings.HasPrefix(part, "default:") {
			info.DefaultValue = strings.TrimPrefix(part, "default:")
		} else if part == "required" {
			info.Required = true
		} else if strings.HasPrefix(part, "in:[") && strings.HasSuffix(part, "]") {
			inValues := strings.TrimPrefix(part, "in:[")
			inValues = strings.TrimSuffix(inValues, "]")
			info.InValues = strings.Split(inValues, ",")
		}
	}

	return
}

func setFieldValue(val string, field reflect.Value) error {
	var eKind reflect.Type
	var eAddr reflect.Value

	if val == "" {
		return nil
	}

	if field.Kind() == reflect.Ptr {
		eKind, eAddr = getFinal(field)
	} else {
		eKind = field.Type()
		eAddr = field
	}

	convertedVal, err := convertVal(eKind, val)

	if err != nil {

		return err
	}

	addrV := reflect.ValueOf(convertedVal)
	eAddr.Set(addrV)

	return nil
}

func getFinal(field reflect.Value) (reflect.Type, reflect.Value) {
	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		return getFinal(field.Elem())
	}

	return field.Type(), field
}

func parseTime(dateTime string) (*time.Time, error) {
	for _, layout := range timeLayouts {
		timeFormatted, err := time.Parse(layout, dateTime)
		if err != nil || timeFormatted.Year() <= 1 { // Year 1 is the default zero value
			continue
		}
		return &timeFormatted, nil
	}
	return nil, fmt.Errorf("invalid date time: %s", dateTime)
}

func convertVal(t reflect.Type, val string) (any, error) {
	switch t.Kind() {
	case reflect.Bool:
		v, err := strconv.ParseBool(val)
		if err != nil {
			return nil, err
		}
		return v, nil

	case reflect.Int:
		v, err := strconv.ParseInt(val, 10, 0)
		if err != nil {
			return nil, err
		}
		return int(v), nil

	case reflect.Int8:
		v, err := strconv.ParseInt(val, 10, 8)
		if err != nil {
			return nil, err
		}
		return int8(v), nil

	case reflect.Int16:
		v, err := strconv.ParseInt(val, 10, 16)
		if err != nil {
			return nil, err
		}
		return int16(v), nil

	case reflect.Int32:
		v, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return nil, err
		}
		return int32(v), nil

	case reflect.Int64:
		v, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case reflect.Uint:
		v, err := strconv.ParseUint(val, 10, 0)
		if err != nil {
			return nil, err
		}
		return uint(v), nil

	case reflect.Uint8:
		v, err := strconv.ParseUint(val, 10, 8)
		if err != nil {
			return nil, err
		}
		return uint8(v), nil

	case reflect.Uint16:
		v, err := strconv.ParseUint(val, 10, 16)
		if err != nil {
			return nil, err
		}
		return uint16(v), nil

	case reflect.Uint32:
		v, err := strconv.ParseUint(val, 10, 32)
		if err != nil {
			return nil, err
		}
		return uint32(v), nil

	case reflect.Uint64:
		v, err := strconv.ParseUint(val, 10, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case reflect.Float32:
		v, err := strconv.ParseFloat(val, 32)
		if err != nil {
			return nil, err
		}
		return float32(v), nil

	case reflect.Float64:
		v, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, err
		}
		return v, nil

	case reflect.Complex64:
		v, err := strconv.ParseComplex(val, 64)
		if err != nil {
			return nil, err
		}
		return complex64(v), nil

	case reflect.Complex128:
		v, err := strconv.ParseComplex(val, 128)
		if err != nil {
			return nil, err
		}
		return v, nil

	case reflect.Slice:
		sliceElemType := t.Elem()
		slicePtr := reflect.New(reflect.SliceOf(sliceElemType)).Interface()

		err := json.Unmarshal([]byte(val), &slicePtr)
		if err != nil {
			return nil, err
		}

		return reflect.ValueOf(slicePtr).Elem().Interface(), nil

	case reflect.String:
		return val, nil

	case reflect.Struct:
		if t.ConvertibleTo(reflect.TypeOf(time.Time{})) {
			convertedTime, err := parseTime(val)
			if err == nil {
				return *convertedTime, nil
			}
			return nil, err
		}

		structPtr := reflect.New(t).Interface()
		err := json.Unmarshal([]byte(val), &structPtr)
		if err == nil {
			return reflect.ValueOf(structPtr).Elem().Interface(), nil
		}

		return nil, errors.New("unsupported struct type")

	default:
		return nil, errors.New("unsupported type")
	}
}
