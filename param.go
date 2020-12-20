package mws

import (
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	keySellerID         = "SellerId"
	keyMWSAuthToken     = "MWSAuthToken"
	keyAWSAccessKeyID   = "AWSAccessKeyId"
	keyVersion          = "Version"
	keyAction           = "Action"
	keySignature        = "Signature"
	keySignatureMethod  = "SignatureMethod"
	keySignatureVersion = "SignatureVersion"
	keyTimestamp        = "Timestamp"
)

type Valuer interface {
	String() string
}

func ParamStruct(paramStructs ...interface{}) Param {
	p := Param{}
	for _, paramStruct := range paramStructs {
		p = p.From(paramStruct)
	}
	return p
}

func ParamSet(key string, value interface{}) Param {
	return Param{}.Set(key, value)
}

func ParamNexToken(nextToken string) Param {
	return ParamSet("NextToken", nextToken)
}

type Param map[string]string

func (p Param) Set(key string, value interface{}) Param {
	if value == nil {
		return p
	}

	var val string
	switch v := value.(type) {
	case string:
		val = v
	case int:
	case int8:
		if v != 0 {
			val = strconv.FormatInt(int64(v), 10)
		}
	case int16:
		if v != 0 {
			val = strconv.FormatInt(int64(v), 10)
		}
	case int32:
		if v != 0 {
			val = strconv.FormatInt(int64(v), 10)
		}
	case int64:
		if v != 0 {
			val = strconv.FormatInt(v, 10)
		}
	case uint:
		if v != 0 {
			val = strconv.FormatInt(int64(v), 10)
		}
	case uint8:
		if v != 0 {
			val = strconv.FormatInt(int64(v), 10)
		}
	case uint16:
		if v != 0 {
			val = strconv.FormatInt(int64(v), 10)
		}
	case uint32:
		if v != 0 {
			val = strconv.FormatInt(int64(v), 10)
		}
	case uint64:
		if v != 0 {
			val = strconv.FormatInt(int64(v), 10)
		}
	case float64:
		if v != 0 {
			val = strconv.FormatFloat(v, 'f', -1, 64)
		}
	case float32:
		if v != 0 {
			val = strconv.FormatFloat(float64(v), 'f', -1, 64)
		}
	case time.Time:
		if !v.IsZero() {
			val = v.Format(time.RFC3339)
		}
	case []string:
		for i, s := range v {
			p[fmt.Sprintf("%s.%d", key, i+1)] = s
		}
		return p
	default:
		if valuer, ok := value.(Valuer); ok {
			val = valuer.String()
		}
	}

	if val != "" {
		p[key] = val
	}

	return p
}

func (p Param) Get(key string) string {
	return p[key]
}

func (p Param) Del(key string) {
	delete(p, key)
	for k := range p {
		if strings.HasPrefix(k, key+".") {
			delete(p, k)
		}
	}
}

func (p Param) Encode() string {
	var keys []string
	for key := range p {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	var buf strings.Builder
	for i, key := range keys {
		value := p[key]
		if i > 0 {
			buf.WriteRune('&')
		}
		buf.WriteString(url.QueryEscape(key))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(value))
	}

	return buf.String()
}

func (p Param) From(paramStruct interface{}) Param {
	if paramStruct == nil {
		return p
	}
	v := reflect.Indirect(reflect.ValueOf(paramStruct))
	if v.Kind() != reflect.Struct {
		//todo log.Warnf("param struct must be struct or struct pointer")
		return p
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag, find := f.Tag.Lookup("mws")
		var (
			name     = f.Name
			required bool
		)

		if find {
			tags := strings.Split(tag, ",")
			if len(tags) > 0 && tags[0] != "" {
				name = tags[0]
			}
			if len(tags) > 1 && tags[1] != "" {
				required = tags[1] == "required"
			}
		}
		val := v.Field(i)
		if required {
			switch val.Kind() {
			case reflect.Bool:
				p.Set(name, strconv.FormatBool(val.Bool()))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				p.Set(name, strconv.FormatInt(val.Int(), 10))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				p.Set(name, strconv.FormatUint(val.Uint(), 10))
			case reflect.Float32, reflect.Float64:
				p.Set(name, strconv.FormatFloat(val.Float(), 'f', -1, 64))
			default:
				p.Set(name, val.Interface())
			}
		} else {
			p.Set(name, val.Interface())
		}
	}
	return p
}

func (p Param) SetAction(action string) Param {
	p.Set(keyAction, action)
	return p
}
