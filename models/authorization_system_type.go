package models
import (
    "errors"
)
// 
type AuthorizationSystemType int

const (
    AZURE_AUTHORIZATIONSYSTEMTYPE AuthorizationSystemType = iota
    GCP_AUTHORIZATIONSYSTEMTYPE
    AWS_AUTHORIZATIONSYSTEMTYPE
    UNKNOWNFUTUREVALUE_AUTHORIZATIONSYSTEMTYPE
)

func (i AuthorizationSystemType) String() string {
    return []string{"azure", "gcp", "aws", "unknownFutureValue"}[i]
}
func ParseAuthorizationSystemType(v string) (any, error) {
    result := AZURE_AUTHORIZATIONSYSTEMTYPE
    switch v {
        case "azure":
            result = AZURE_AUTHORIZATIONSYSTEMTYPE
        case "gcp":
            result = GCP_AUTHORIZATIONSYSTEMTYPE
        case "aws":
            result = AWS_AUTHORIZATIONSYSTEMTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AUTHORIZATIONSYSTEMTYPE
        default:
            return 0, errors.New("Unknown AuthorizationSystemType value: " + v)
    }
    return &result, nil
}
func SerializeAuthorizationSystemType(values []AuthorizationSystemType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthorizationSystemType) isMultiValue() bool {
    return false
}
