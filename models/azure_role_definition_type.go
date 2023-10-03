package models
import (
    "errors"
)
// 
type AzureRoleDefinitionType int

const (
    SYSTEM_AZUREROLEDEFINITIONTYPE AzureRoleDefinitionType = iota
    CUSTOM_AZUREROLEDEFINITIONTYPE
    UNKNOWNFUTUREVALUE_AZUREROLEDEFINITIONTYPE
)

func (i AzureRoleDefinitionType) String() string {
    return []string{"system", "custom", "unknownFutureValue"}[i]
}
func ParseAzureRoleDefinitionType(v string) (any, error) {
    result := SYSTEM_AZUREROLEDEFINITIONTYPE
    switch v {
        case "system":
            result = SYSTEM_AZUREROLEDEFINITIONTYPE
        case "custom":
            result = CUSTOM_AZUREROLEDEFINITIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AZUREROLEDEFINITIONTYPE
        default:
            return 0, errors.New("Unknown AzureRoleDefinitionType value: " + v)
    }
    return &result, nil
}
func SerializeAzureRoleDefinitionType(values []AzureRoleDefinitionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AzureRoleDefinitionType) isMultiValue() bool {
    return false
}
