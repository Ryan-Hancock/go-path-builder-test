package builder

import (
	"fmt"
	"reflect"
	"strconv"
)

// PathParser repersents a path parser.
type PathParser struct {
	Separator string
}

// ParseArray will take a slice of interfaces along the lines of ['abc', nil]
func (p PathParser) ParseArray(arr []interface{}) (path string) {
	var strArray []string

	for _, str := range arr {
		if str == nil {
			continue
		}

		if reflect.TypeOf(str).Kind() == reflect.Int {
			str = strconv.Itoa(str.(int))
		}

		if reflect.TypeOf(str).Kind() != reflect.String {
			continue
		}

		if p.isEmpty(str.(string)) {
			continue
		}

		strArray = append(strArray, str.(string))
	}

	if len(strArray) != 0 {
		path = p.JoinArray(strArray)
	}

	if len(strArray) == 1 {
		return strArray[0]
	}

	return
}

// JoinArray joins the strings together by the separator
func (p PathParser) JoinArray(strArr []string) (joined string) {
	for _, str := range strArr {
		if p.checkStringIndexIs(str, 0) {
			joined = fmt.Sprintf("%s%s", joined, str)
			continue
		}

		if joined != "" {
			if !p.checkStringIndexIs(joined, len(joined)-1) {
				joined = fmt.Sprintf("%s%s%s", joined, p.Separator, str)
				continue
			}
		}

		if p.checkStringIndexIs(str, len(str)-1) {
			joined = fmt.Sprintf("%s%s", joined, str)
			continue
		}

		joined = fmt.Sprintf("%s%s", joined, str)
	}

	return
}

func (p PathParser) checkStringIndexIs(str string, index int) bool {
	return string(str[index]) == p.Separator
}

func (p PathParser) isEmpty(str string) bool {
	return str == ""
}
