package models
// An error code indicating the failure reason, when the deployment state is createFailed. Possible values: See zebraFotaErrorCode enum.
type ZebraFotaErrorCode int

const (
    // Default error code indicating success (no error).
    SUCCESS_ZEBRAFOTAERRORCODE ZebraFotaErrorCode = iota
    // This error indicates that no devices were found in the selected Azure Active Directory (AAD) security group(s) when creating a deployment. E.g.: a user selects one or more security groups that doesn't include any devices
    NODEVICESFOUNDINSELECTEDAADGROUPS_ZEBRAFOTAERRORCODE
    // This error indicates that no Intune managed devices were found in the selected Azure Active Directory (AAD) security groups when creating a deployment. E.g.: a user selects a group that includes devices that are no longer managed in Intune.
    NOINTUNEDEVICESFOUNDINSELECTEDAADGROUPS_ZEBRAFOTAERRORCODE
    // This error indicates that no Zebra FOTA enrolled devices were found for the current tenant when creating a deployment. E.g.: a user has not enrolled any devices in Zebra FOTA and attempts to create a deployment targeting only non-Zebra devices.
    NOZEBRAFOTAENROLLEDDEVICESFOUNDFORCURRENTTENANT_ZEBRAFOTAERRORCODE
    // This error indicates that no Zebra FOTA enrolled devices were found in the selected Azure Active Directory (AAD) Groups when creating a deployment. E.g.: a user has enrolled one or more devices in Zebra FOTA, but attempts to create a deployment targeting security groups that don't include any Zebra Devices.
    NOZEBRAFOTAENROLLEDDEVICESFOUNDINSELECTEDAADGROUPS_ZEBRAFOTAERRORCODE
    // This error indicates that no Zebra FOTA devices were found that match the selected device model when creating a deployment. E.g.: a user selects `TC8300` device model, but no Zebra devices with the same device model were found in the targeted Azure Active Directory (AAD) security groups.
    NOZEBRAFOTADEVICESFOUNDFORSELECTEDDEVICEMODEL_ZEBRAFOTAERRORCODE
    // This error indicates that an external request to Zebra APIs failed when creating a deployment. The failure can be caused by a transient issue (e.g.: Network is down) or caused by a permanent server error.
    ZEBRAFOTACREATEDEPLOYMENTREQUESTFAILURE_ZEBRAFOTAERRORCODE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_ZEBRAFOTAERRORCODE
)

func (i ZebraFotaErrorCode) String() string {
    return []string{"success", "noDevicesFoundInSelectedAadGroups", "noIntuneDevicesFoundInSelectedAadGroups", "noZebraFotaEnrolledDevicesFoundForCurrentTenant", "noZebraFotaEnrolledDevicesFoundInSelectedAadGroups", "noZebraFotaDevicesFoundForSelectedDeviceModel", "zebraFotaCreateDeploymentRequestFailure", "unknownFutureValue"}[i]
}
func ParseZebraFotaErrorCode(v string) (any, error) {
    result := SUCCESS_ZEBRAFOTAERRORCODE
    switch v {
        case "success":
            result = SUCCESS_ZEBRAFOTAERRORCODE
        case "noDevicesFoundInSelectedAadGroups":
            result = NODEVICESFOUNDINSELECTEDAADGROUPS_ZEBRAFOTAERRORCODE
        case "noIntuneDevicesFoundInSelectedAadGroups":
            result = NOINTUNEDEVICESFOUNDINSELECTEDAADGROUPS_ZEBRAFOTAERRORCODE
        case "noZebraFotaEnrolledDevicesFoundForCurrentTenant":
            result = NOZEBRAFOTAENROLLEDDEVICESFOUNDFORCURRENTTENANT_ZEBRAFOTAERRORCODE
        case "noZebraFotaEnrolledDevicesFoundInSelectedAadGroups":
            result = NOZEBRAFOTAENROLLEDDEVICESFOUNDINSELECTEDAADGROUPS_ZEBRAFOTAERRORCODE
        case "noZebraFotaDevicesFoundForSelectedDeviceModel":
            result = NOZEBRAFOTADEVICESFOUNDFORSELECTEDDEVICEMODEL_ZEBRAFOTAERRORCODE
        case "zebraFotaCreateDeploymentRequestFailure":
            result = ZEBRAFOTACREATEDEPLOYMENTREQUESTFAILURE_ZEBRAFOTAERRORCODE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ZEBRAFOTAERRORCODE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeZebraFotaErrorCode(values []ZebraFotaErrorCode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ZebraFotaErrorCode) isMultiValue() bool {
    return false
}
