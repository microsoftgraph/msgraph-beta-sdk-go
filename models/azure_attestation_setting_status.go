package models
import (
    "errors"
)
// A list of possible Azure Attestation states for a device. Azure Attestation setting status is determined by report sent from Microsoft Azure Attestation service. Only Windows 11 devices will have values "enabled" or "disabled". Windows 10 devices will have value "notApplicable".
type AzureAttestationSettingStatus int

const (
    // Indicates that the device is not a Windows 11 device.
    NOTAPPLICABLE_AZUREATTESTATIONSETTINGSTATUS AzureAttestationSettingStatus = iota
    // Indicates that the device has the Azure attestation setting enabled.
    ENABLED_AZUREATTESTATIONSETTINGSTATUS
    // Indicates that the device has the Azure attestation setting disabled.
    DISABLED_AZUREATTESTATIONSETTINGSTATUS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_AZUREATTESTATIONSETTINGSTATUS
)

func (i AzureAttestationSettingStatus) String() string {
    return []string{"notApplicable", "enabled", "disabled", "unknownFutureValue"}[i]
}
func ParseAzureAttestationSettingStatus(v string) (any, error) {
    result := NOTAPPLICABLE_AZUREATTESTATIONSETTINGSTATUS
    switch v {
        case "notApplicable":
            result = NOTAPPLICABLE_AZUREATTESTATIONSETTINGSTATUS
        case "enabled":
            result = ENABLED_AZUREATTESTATIONSETTINGSTATUS
        case "disabled":
            result = DISABLED_AZUREATTESTATIONSETTINGSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AZUREATTESTATIONSETTINGSTATUS
        default:
            return 0, errors.New("Unknown AzureAttestationSettingStatus value: " + v)
    }
    return &result, nil
}
func SerializeAzureAttestationSettingStatus(values []AzureAttestationSettingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AzureAttestationSettingStatus) isMultiValue() bool {
    return false
}
