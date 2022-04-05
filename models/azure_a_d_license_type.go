package models
import (
    "strings"
    "errors"
)
// Provides operations to call the getAzureADLicenseUsage method.
type AzureADLicenseType int

const (
    NONE_AZUREADLICENSETYPE AzureADLicenseType = iota
    FREE_AZUREADLICENSETYPE
    BASIC_AZUREADLICENSETYPE
    PREMIUMP1_AZUREADLICENSETYPE
    PREMIUMP2_AZUREADLICENSETYPE
    UNKNOWNFUTUREVALUE_AZUREADLICENSETYPE
)

func (i AzureADLicenseType) String() string {
    return []string{"NONE", "FREE", "BASIC", "PREMIUMP1", "PREMIUMP2", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAzureADLicenseType(v string) (interface{}, error) {
    result := NONE_AZUREADLICENSETYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_AZUREADLICENSETYPE
        case "FREE":
            result = FREE_AZUREADLICENSETYPE
        case "BASIC":
            result = BASIC_AZUREADLICENSETYPE
        case "PREMIUMP1":
            result = PREMIUMP1_AZUREADLICENSETYPE
        case "PREMIUMP2":
            result = PREMIUMP2_AZUREADLICENSETYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_AZUREADLICENSETYPE
        default:
            return 0, errors.New("Unknown AzureADLicenseType value: " + v)
    }
    return &result, nil
}
func SerializeAzureADLicenseType(values []AzureADLicenseType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
