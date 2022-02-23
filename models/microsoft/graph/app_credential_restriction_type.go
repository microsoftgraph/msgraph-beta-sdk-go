package graph
import (
    "strings"
    "errors"
)
// 
type AppCredentialRestrictionType int

const (
    PASSWORDADDITION_APPCREDENTIALRESTRICTIONTYPE AppCredentialRestrictionType = iota
    PASSWORDLIFETIME_APPCREDENTIALRESTRICTIONTYPE
    SYMMETRICKEYADDITION_APPCREDENTIALRESTRICTIONTYPE
    SYMMETRICKEYLIFETIME_APPCREDENTIALRESTRICTIONTYPE
    CUSTOMPASSWORDADDITION_APPCREDENTIALRESTRICTIONTYPE
    UNKNOWNFUTUREVALUE_APPCREDENTIALRESTRICTIONTYPE
)

func (i AppCredentialRestrictionType) String() string {
    return []string{"PASSWORDADDITION", "PASSWORDLIFETIME", "SYMMETRICKEYADDITION", "SYMMETRICKEYLIFETIME", "CUSTOMPASSWORDADDITION", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAppCredentialRestrictionType(v string) (interface{}, error) {
    result := PASSWORDADDITION_APPCREDENTIALRESTRICTIONTYPE
    switch strings.ToUpper(v) {
        case "PASSWORDADDITION":
            result = PASSWORDADDITION_APPCREDENTIALRESTRICTIONTYPE
        case "PASSWORDLIFETIME":
            result = PASSWORDLIFETIME_APPCREDENTIALRESTRICTIONTYPE
        case "SYMMETRICKEYADDITION":
            result = SYMMETRICKEYADDITION_APPCREDENTIALRESTRICTIONTYPE
        case "SYMMETRICKEYLIFETIME":
            result = SYMMETRICKEYLIFETIME_APPCREDENTIALRESTRICTIONTYPE
        case "CUSTOMPASSWORDADDITION":
            result = CUSTOMPASSWORDADDITION_APPCREDENTIALRESTRICTIONTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_APPCREDENTIALRESTRICTIONTYPE
        default:
            return 0, errors.New("Unknown AppCredentialRestrictionType value: " + v)
    }
    return &result, nil
}
func SerializeAppCredentialRestrictionType(values []AppCredentialRestrictionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
