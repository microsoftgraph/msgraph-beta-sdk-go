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
    UNKNOWNFUTUREVALUE_APPCREDENTIALRESTRICTIONTYPE
)

func (i AppCredentialRestrictionType) String() string {
    return []string{"PASSWORDADDITION", "PASSWORDLIFETIME", "SYMMETRICKEYADDITION", "SYMMETRICKEYLIFETIME", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAppCredentialRestrictionType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "PASSWORDADDITION":
            return PASSWORDADDITION_APPCREDENTIALRESTRICTIONTYPE, nil
        case "PASSWORDLIFETIME":
            return PASSWORDLIFETIME_APPCREDENTIALRESTRICTIONTYPE, nil
        case "SYMMETRICKEYADDITION":
            return SYMMETRICKEYADDITION_APPCREDENTIALRESTRICTIONTYPE, nil
        case "SYMMETRICKEYLIFETIME":
            return SYMMETRICKEYLIFETIME_APPCREDENTIALRESTRICTIONTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_APPCREDENTIALRESTRICTIONTYPE, nil
    }
    return 0, errors.New("Unknown AppCredentialRestrictionType value: " + v)
}
func SerializeAppCredentialRestrictionType(values []AppCredentialRestrictionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
