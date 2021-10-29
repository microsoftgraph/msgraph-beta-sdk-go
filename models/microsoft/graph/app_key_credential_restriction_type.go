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
    switch strings.ToUpper(v) {
        case "ASYMMETRICKEYLIFETIME":
            return ASYMMETRICKEYLIFETIME_APPKEYCREDENTIALRESTRICTIONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_APPKEYCREDENTIALRESTRICTIONTYPE, nil
    }
    return 0, errors.New("Unknown AppKeyCredentialRestrictionType value: " + v)
}
func SerializeAppKeyCredentialRestrictionType(values []AppKeyCredentialRestrictionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
