package models
import (
    "errors"
)
// 
type AzureEncryption int

const (
    MICROSOFTSTORAGE_AZUREENCRYPTION AzureEncryption = iota
    MICROSOFTKEYVAULT_AZUREENCRYPTION
    CUSTOMER_AZUREENCRYPTION
    UNKNOWNFUTUREVALUE_AZUREENCRYPTION
)

func (i AzureEncryption) String() string {
    return []string{"microsoftStorage", "microsoftKeyVault", "customer", "unknownFutureValue"}[i]
}
func ParseAzureEncryption(v string) (any, error) {
    result := MICROSOFTSTORAGE_AZUREENCRYPTION
    switch v {
        case "microsoftStorage":
            result = MICROSOFTSTORAGE_AZUREENCRYPTION
        case "microsoftKeyVault":
            result = MICROSOFTKEYVAULT_AZUREENCRYPTION
        case "customer":
            result = CUSTOMER_AZUREENCRYPTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AZUREENCRYPTION
        default:
            return 0, errors.New("Unknown AzureEncryption value: " + v)
    }
    return &result, nil
}
func SerializeAzureEncryption(values []AzureEncryption) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AzureEncryption) isMultiValue() bool {
    return false
}
