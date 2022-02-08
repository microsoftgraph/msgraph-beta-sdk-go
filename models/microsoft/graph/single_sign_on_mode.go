package graph
import (
    "strings"
    "errors"
)
// 
type SingleSignOnMode int

const (
    NONE_SINGLESIGNONMODE SingleSignOnMode = iota
    ONPREMISESKERBEROS_SINGLESIGNONMODE
    SAML_SINGLESIGNONMODE
    PINGHEADERBASED_SINGLESIGNONMODE
    AADHEADERBASED_SINGLESIGNONMODE
    UNKNOWNFUTUREVALUE_SINGLESIGNONMODE
)

func (i SingleSignOnMode) String() string {
    return []string{"NONE", "ONPREMISESKERBEROS", "SAML", "PINGHEADERBASED", "AADHEADERBASED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSingleSignOnMode(v string) (interface{}, error) {
    result := NONE_SINGLESIGNONMODE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_SINGLESIGNONMODE
        case "ONPREMISESKERBEROS":
            result = ONPREMISESKERBEROS_SINGLESIGNONMODE
        case "SAML":
            result = SAML_SINGLESIGNONMODE
        case "PINGHEADERBASED":
            result = PINGHEADERBASED_SINGLESIGNONMODE
        case "AADHEADERBASED":
            result = AADHEADERBASED_SINGLESIGNONMODE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_SINGLESIGNONMODE
        default:
            return 0, errors.New("Unknown SingleSignOnMode value: " + v)
    }
    return &result, nil
}
func SerializeSingleSignOnMode(values []SingleSignOnMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
