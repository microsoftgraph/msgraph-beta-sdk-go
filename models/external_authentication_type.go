package models
import (
    "errors"
)
type ExternalAuthenticationType int

const (
    PASSTHRU_EXTERNALAUTHENTICATIONTYPE ExternalAuthenticationType = iota
    AADPREAUTHENTICATION_EXTERNALAUTHENTICATIONTYPE
)

func (i ExternalAuthenticationType) String() string {
    return []string{"passthru", "aadPreAuthentication"}[i]
}
func ParseExternalAuthenticationType(v string) (any, error) {
    result := PASSTHRU_EXTERNALAUTHENTICATIONTYPE
    switch v {
        case "passthru":
            result = PASSTHRU_EXTERNALAUTHENTICATIONTYPE
        case "aadPreAuthentication":
            result = AADPREAUTHENTICATION_EXTERNALAUTHENTICATIONTYPE
        default:
            return 0, errors.New("Unknown ExternalAuthenticationType value: " + v)
    }
    return &result, nil
}
func SerializeExternalAuthenticationType(values []ExternalAuthenticationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ExternalAuthenticationType) isMultiValue() bool {
    return false
}
