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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_SINGLESIGNONMODE, nil
        case "ONPREMISESKERBEROS":
            return ONPREMISESKERBEROS_SINGLESIGNONMODE, nil
        case "SAML":
            return SAML_SINGLESIGNONMODE, nil
        case "PINGHEADERBASED":
            return PINGHEADERBASED_SINGLESIGNONMODE, nil
        case "AADHEADERBASED":
            return AADHEADERBASED_SINGLESIGNONMODE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SINGLESIGNONMODE, nil
    }
    return 0, errors.New("Unknown SingleSignOnMode value: " + v)
}
func SerializeSingleSignOnMode(values []SingleSignOnMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
