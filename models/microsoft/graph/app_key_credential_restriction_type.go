package graph
import (
    "strings"
    "errors"
)
// 
type AppKeyCredentialRestrictionType int

const (
    ASYMMETRICKEYLIFETIME_APPKEYCREDENTIALRESTRICTIONTYPE AppKeyCredentialRestrictionType = iota
    UNKNOWNFUTUREVALUE_APPKEYCREDENTIALRESTRICTIONTYPE
)

func (i AppKeyCredentialRestrictionType) String() string {
    return []string{"ASYMMETRICKEYLIFETIME", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAppKeyCredentialRestrictionType(v string) (interface{}, error) {
    result := ASYMMETRICKEYLIFETIME_APPKEYCREDENTIALRESTRICTIONTYPE
    switch strings.ToUpper(v) {
        case "ASYMMETRICKEYLIFETIME":
            result = ASYMMETRICKEYLIFETIME_APPKEYCREDENTIALRESTRICTIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_APPKEYCREDENTIALRESTRICTIONTYPE
        default:
            return 0, errors.New("Unknown AppKeyCredentialRestrictionType value: " + v)
    }
    return &result, nil
}
func SerializeAppKeyCredentialRestrictionType(values []AppKeyCredentialRestrictionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
