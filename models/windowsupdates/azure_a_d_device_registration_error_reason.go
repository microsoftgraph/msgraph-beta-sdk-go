package windowsupdates
type AzureADDeviceRegistrationErrorReason int

const (
    INVALIDGLOBALDEVICEID_AZUREADDEVICEREGISTRATIONERRORREASON AzureADDeviceRegistrationErrorReason = iota
    INVALIDAZUREADDEVICEID_AZUREADDEVICEREGISTRATIONERRORREASON
    MISSINGTRUSTTYPE_AZUREADDEVICEREGISTRATIONERRORREASON
    INVALIDAZUREADJOIN_AZUREADDEVICEREGISTRATIONERRORREASON
    UNKNOWNFUTUREVALUE_AZUREADDEVICEREGISTRATIONERRORREASON
)

func (i AzureADDeviceRegistrationErrorReason) String() string {
    return []string{"invalidGlobalDeviceId", "invalidAzureADDeviceId", "missingTrustType", "invalidAzureADJoin", "unknownFutureValue"}[i]
}
func ParseAzureADDeviceRegistrationErrorReason(v string) (any, error) {
    result := INVALIDGLOBALDEVICEID_AZUREADDEVICEREGISTRATIONERRORREASON
    switch v {
        case "invalidGlobalDeviceId":
            result = INVALIDGLOBALDEVICEID_AZUREADDEVICEREGISTRATIONERRORREASON
        case "invalidAzureADDeviceId":
            result = INVALIDAZUREADDEVICEID_AZUREADDEVICEREGISTRATIONERRORREASON
        case "missingTrustType":
            result = MISSINGTRUSTTYPE_AZUREADDEVICEREGISTRATIONERRORREASON
        case "invalidAzureADJoin":
            result = INVALIDAZUREADJOIN_AZUREADDEVICEREGISTRATIONERRORREASON
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AZUREADDEVICEREGISTRATIONERRORREASON
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAzureADDeviceRegistrationErrorReason(values []AzureADDeviceRegistrationErrorReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AzureADDeviceRegistrationErrorReason) isMultiValue() bool {
    return false
}
