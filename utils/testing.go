package utils

import "reflect"

func IsEqualIgnoringTags(a, b interface{}, ignoreTags ...string) bool {
	v1 := reflect.ValueOf(a)
	v2 := reflect.ValueOf(b)

	if v1.Type() != v2.Type() {
		return false
	}

	for i := 0; i < v1.NumField(); i++ {
		field := v1.Type().Field(i)
		if Contains(ignoreTags, field.Tag) {
			continue
		}
		if !reflect.DeepEqual(v1.Field(i).Interface(), v2.Field(i).Interface()) {
			return false
		}
	}
	return true
}

func Contains(tags []string, tag reflect.StructTag) bool {
	for _, t := range tags {
		if tag.Get(t) == "true" {
			return true
		}
	}
	return false
}
