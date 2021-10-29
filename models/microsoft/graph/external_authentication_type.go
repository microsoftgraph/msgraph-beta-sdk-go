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
    switch strings.ToUpper(v) {
        case "PASSTHRU":
            return PASSTHRU_EXTERNALAUTHENTICATIONTYPE, nil
        case "AADPREAUTHENTICATION":
            return AADPREAUTHENTICATION_EXTERNALAUTHENTICATIONTYPE, nil
    }
    return 0, errors.New("Unknown ExternalAuthenticationType value: " + v)
}
func SerializeExternalAuthenticationType(values []ExternalAuthenticationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
