package externalconnectors
import (
    "strings"
    "errors"
)
type IdentitySourceType int

const (
    AZUREACTIVEDIRECTORY_IDENTITYSOURCETYPE IdentitySourceType = iota
    EXTERNAL_IDENTITYSOURCETYPE
    UNKNOWNFUTUREVALUE_IDENTITYSOURCETYPE
)

func (i IdentitySourceType) String() string {
    return []string{"AZUREACTIVEDIRECTORY", "EXTERNAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseIdentitySourceType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "AZUREACTIVEDIRECTORY":
            return AZUREACTIVEDIRECTORY_IDENTITYSOURCETYPE, nil
        case "EXTERNAL":
            return EXTERNAL_IDENTITYSOURCETYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_IDENTITYSOURCETYPE, nil
    }
    return 0, errors.New("Unknown IdentitySourceType value: " + v)
}
func SerializeIdentitySourceType(values []IdentitySourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
