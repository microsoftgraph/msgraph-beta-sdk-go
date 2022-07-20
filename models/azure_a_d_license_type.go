package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
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
    return []string{"none", "free", "basic", "premiumP1", "premiumP2", "unknownFutureValue"}[i]
}
func ParseAzureADLicenseType(v string) (interface{}, error) {
    result := NONE_AZUREADLICENSETYPE
    switch v {
        case "none":
            result = NONE_AZUREADLICENSETYPE
        case "free":
            result = FREE_AZUREADLICENSETYPE
        case "basic":
            result = BASIC_AZUREADLICENSETYPE
        case "premiumP1":
            result = PREMIUMP1_AZUREADLICENSETYPE
        case "premiumP2":
            result = PREMIUMP2_AZUREADLICENSETYPE
        case "unknownFutureValue":
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
