package models
import (
    "errors"
)
type AccessPackageSubjectLifecycle int

const (
    NOTDEFINED_ACCESSPACKAGESUBJECTLIFECYCLE AccessPackageSubjectLifecycle = iota
    NOTGOVERNED_ACCESSPACKAGESUBJECTLIFECYCLE
    GOVERNED_ACCESSPACKAGESUBJECTLIFECYCLE
    UNKNOWNFUTUREVALUE_ACCESSPACKAGESUBJECTLIFECYCLE
)

func (i AccessPackageSubjectLifecycle) String() string {
    return []string{"notDefined", "notGoverned", "governed", "unknownFutureValue"}[i]
}
func ParseAccessPackageSubjectLifecycle(v string) (any, error) {
    result := NOTDEFINED_ACCESSPACKAGESUBJECTLIFECYCLE
    switch v {
        case "notDefined":
            result = NOTDEFINED_ACCESSPACKAGESUBJECTLIFECYCLE
        case "notGoverned":
            result = NOTGOVERNED_ACCESSPACKAGESUBJECTLIFECYCLE
        case "governed":
            result = GOVERNED_ACCESSPACKAGESUBJECTLIFECYCLE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ACCESSPACKAGESUBJECTLIFECYCLE
        default:
            return 0, errors.New("Unknown AccessPackageSubjectLifecycle value: " + v)
    }
    return &result, nil
}
func SerializeAccessPackageSubjectLifecycle(values []AccessPackageSubjectLifecycle) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AccessPackageSubjectLifecycle) isMultiValue() bool {
    return false
}
