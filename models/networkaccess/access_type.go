package networkaccess
import (
    "errors"
)
// 
type AccessType int

const (
    QUICKACCESS_ACCESSTYPE AccessType = iota
    PRIVATEACCESS_ACCESSTYPE
    UNKNOWNFUTUREVALUE_ACCESSTYPE
)

func (i AccessType) String() string {
    return []string{"quickAccess", "privateAccess", "unknownFutureValue"}[i]
}
func ParseAccessType(v string) (any, error) {
    result := QUICKACCESS_ACCESSTYPE
    switch v {
        case "quickAccess":
            result = QUICKACCESS_ACCESSTYPE
        case "privateAccess":
            result = PRIVATEACCESS_ACCESSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCESSTYPE
        default:
            return 0, errors.New("Unknown AccessType value: " + v)
    }
    return &result, nil
}
func SerializeAccessType(values []AccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AccessType) isMultiValue() bool {
    return false
}
