package models
import (
    "errors"
)
type AzureAccessType int

const (
    PUBLIC_AZUREACCESSTYPE AzureAccessType = iota
    PRIVATE_AZUREACCESSTYPE
    UNKNOWNFUTUREVALUE_AZUREACCESSTYPE
)

func (i AzureAccessType) String() string {
    return []string{"public", "private", "unknownFutureValue"}[i]
}
func ParseAzureAccessType(v string) (any, error) {
    result := PUBLIC_AZUREACCESSTYPE
    switch v {
        case "public":
            result = PUBLIC_AZUREACCESSTYPE
        case "private":
            result = PRIVATE_AZUREACCESSTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AZUREACCESSTYPE
        default:
            return 0, errors.New("Unknown AzureAccessType value: " + v)
    }
    return &result, nil
}
func SerializeAzureAccessType(values []AzureAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AzureAccessType) isMultiValue() bool {
    return false
}
