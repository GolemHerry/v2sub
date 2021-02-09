package utils

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

var (
	invalidObjError = errors.New("invalid object")
	notExitError    = errors.New("not exit")
)

func GetFormBytes(bytes []byte, path string) (interface{}, error) {
	var data interface{}
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}
	return getFromUnmarshal(data, path)
}

func GetFromReader(reader io.Reader, path string) (interface{}, error) {
	var data interface{}
	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return nil, err
	}
	return getFromUnmarshal(data, path)
}

func WriteJson(jsonFile, filedExp, value string) error {
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	var jsonData map[string]interface{}
	if err = json.Unmarshal(byteValue, &jsonData); err != nil {
		return err
	}

	cursor := jsonData
	token, err := parseExp(filedExp)
	if err != nil {
		return err
	}
	for _, t := range token[:len(token)-1] {
		subObj, err := getSubFromObj(cursor, t)
		if err != nil {
			return err
		}
		cursor = subObj.(map[string]interface{})
	}
	cursor[token[len(token)-1]] = value

	byteValue, err = json.Marshal(jsonData)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(jsonFile, byteValue, 0644)
}

func getFromUnmarshal(data interface{}, path string) (interface{}, error) {
	token, err := parseExp(path)
	if err != nil {
		return nil, err
	}
	return get(data, token)
}

func parseExp(path string) ([]string, error) {
	token := make([]string, 0)
	for _, v := range strings.Split(path, ".") {
		if !strings.Contains(v, "[") {
			token = append(token, v)
		} else {
			token = append(token, v[:strings.Index(v, "[")], v[strings.Index(v, "[")+1:strings.Index(v, "]")])
		}
	}
	return token, nil
}

func get(data interface{}, token []string) (interface{}, error) {
	var err error
	obj := data

	for _, index := range token {
		obj, err = getSubFromObj(obj, index)
		if err != nil {
			return nil, err
		}
	}

	return obj, nil
}

func getSubFromObj(obj interface{}, index string) (interface{}, error) {
	if reflect.TypeOf(obj) == nil {
		return nil, invalidObjError
	}

	switch reflect.ValueOf(obj).Kind() {
	case reflect.Map:
		for _, v := range reflect.ValueOf(obj).MapKeys() {
			if v.String() == index {
				return reflect.ValueOf(obj).MapIndex(v).Interface(), nil
			}
		}
		return nil, notExitError
	case reflect.Slice:
		i, _ := strconv.Atoi(index)
		if i > -1 && i < reflect.ValueOf(obj).Len() {
			return reflect.ValueOf(obj).Index(i).Interface(), nil
		}
		return nil, notExitError
	default:
		return nil, invalidObjError
	}
}
