package graph
import (
    "strings"
    "errors"
)
// 
type ExternalAuthenticationType int

const (
    PASSTHRU_EXTERNALAUTHENTICATIONTYPE ExternalAuthenticationType = iota
    AADPREAUTHENTICATION_EXTERNALAUTHENTICATIONTYPE
)

func (i ExternalAuthenticationType) String() string {
    return []string{"PASSTHRU", "AADPREAUTHENTICATION"}[i]
}
func ParseExternalAuthenticationType(v string) (interface{}, error) {
    result := PASSTHRU_EXTERNALAUTHENTICATIONTYPE
    switch strings.ToUpper(v) {
        case "PASSTHRU":
            result = PASSTHRU_EXTERNALAUTHENTICATIONTYPE
        case "AADPREAUTHENTICATION":
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
