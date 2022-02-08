package externalconnectors
import (
    "strings"
    "errors"
)
// 
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
    result := AZUREACTIVEDIRECTORY_IDENTITYSOURCETYPE
    switch strings.ToUpper(v) {
        case "AZUREACTIVEDIRECTORY":
            result = AZUREACTIVEDIRECTORY_IDENTITYSOURCETYPE
        case "EXTERNAL":
            result = EXTERNAL_IDENTITYSOURCETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_IDENTITYSOURCETYPE
        default:
            return 0, errors.New("Unknown IdentitySourceType value: " + v)
    }
    return &result, nil
}
func SerializeIdentitySourceType(values []IdentitySourceType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
