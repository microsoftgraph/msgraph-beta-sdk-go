package cloudlicensing
import (
    "math"
    "strings"
)
type AssigneeTypes int

const (
    NONE_ASSIGNEETYPES = 1
    USER_ASSIGNEETYPES = 2
    GROUP_ASSIGNEETYPES = 4
    DEVICE_ASSIGNEETYPES = 8
    UNKNOWNFUTUREVALUE_ASSIGNEETYPES = 16
)

func (i AssigneeTypes) String() string {
    var values []string
    options := []string{"none", "user", "group", "device", "unknownFutureValue"}
    for p := 0; p < 5; p++ {
        mantis := AssigneeTypes(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAssigneeTypes(v string) (any, error) {
    var result AssigneeTypes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_ASSIGNEETYPES
            case "user":
                result |= USER_ASSIGNEETYPES
            case "group":
                result |= GROUP_ASSIGNEETYPES
            case "device":
                result |= DEVICE_ASSIGNEETYPES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ASSIGNEETYPES
            default:
                return nil, nil
        }
    }
    return &result, nil
}
func SerializeAssigneeTypes(values []AssigneeTypes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AssigneeTypes) isMultiValue() bool {
    return true
}
