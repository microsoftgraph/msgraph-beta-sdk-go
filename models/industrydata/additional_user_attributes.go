package industrydata
import (
    "errors"
)
type AdditionalUserAttributes int

const (
    USERGRADELEVEL_ADDITIONALUSERATTRIBUTES AdditionalUserAttributes = iota
    USERNUMBER_ADDITIONALUSERATTRIBUTES
    UNKNOWNFUTUREVALUE_ADDITIONALUSERATTRIBUTES
)

func (i AdditionalUserAttributes) String() string {
    return []string{"userGradeLevel", "userNumber", "unknownFutureValue"}[i]
}
func ParseAdditionalUserAttributes(v string) (any, error) {
    result := USERGRADELEVEL_ADDITIONALUSERATTRIBUTES
    switch v {
        case "userGradeLevel":
            result = USERGRADELEVEL_ADDITIONALUSERATTRIBUTES
        case "userNumber":
            result = USERNUMBER_ADDITIONALUSERATTRIBUTES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ADDITIONALUSERATTRIBUTES
        default:
            return 0, errors.New("Unknown AdditionalUserAttributes value: " + v)
    }
    return &result, nil
}
func SerializeAdditionalUserAttributes(values []AdditionalUserAttributes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AdditionalUserAttributes) isMultiValue() bool {
    return false
}
