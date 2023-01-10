package models
import (
    "errors"
)
// 
type ManagedDeviceRemoteAction int

const (
    RETIRE_MANAGEDDEVICEREMOTEACTION ManagedDeviceRemoteAction = iota
    DELETE_MANAGEDDEVICEREMOTEACTION
    FULLSCAN_MANAGEDDEVICEREMOTEACTION
    QUICKSCAN_MANAGEDDEVICEREMOTEACTION
    SIGNATUREUPDATE_MANAGEDDEVICEREMOTEACTION
    WIPE_MANAGEDDEVICEREMOTEACTION
    CUSTOMTEXTNOTIFICATION_MANAGEDDEVICEREMOTEACTION
    REBOOTNOW_MANAGEDDEVICEREMOTEACTION
    SETDEVICENAME_MANAGEDDEVICEREMOTEACTION
    SYNCDEVICE_MANAGEDDEVICEREMOTEACTION
    // Name of the deprovision action.
    DEPROVISION_MANAGEDDEVICEREMOTEACTION
    // Name of the disable action.
    DISABLE_MANAGEDDEVICEREMOTEACTION
    // Name of the reenable action.
    REENABLE_MANAGEDDEVICEREMOTEACTION
    // Name of the moveDevicesToOU action.
    MOVEDEVICETOORGANIZATIONALUNIT_MANAGEDDEVICEREMOTEACTION
    // Name of action to Activate eSIM on the device.
    ACTIVATEDEVICEESIM_MANAGEDDEVICEREMOTEACTION
    // Name of the collectDiagnostics action.
    COLLECTDIAGNOSTICS_MANAGEDDEVICEREMOTEACTION
    // Name of action to initiate MDM key recovery
    INITIATEMOBILEDEVICEMANAGEMENTKEYRECOVERY_MANAGEDDEVICEREMOTEACTION
)

func (i ManagedDeviceRemoteAction) String() string {
    return []string{"retire", "delete", "fullScan", "quickScan", "signatureUpdate", "wipe", "customTextNotification", "rebootNow", "setDeviceName", "syncDevice", "deprovision", "disable", "reenable", "moveDeviceToOrganizationalUnit", "activateDeviceEsim", "collectDiagnostics", "initiateMobileDeviceManagementKeyRecovery"}[i]
}
func ParseManagedDeviceRemoteAction(v string) (interface{}, error) {
    result := RETIRE_MANAGEDDEVICEREMOTEACTION
    switch v {
        case "retire":
            result = RETIRE_MANAGEDDEVICEREMOTEACTION
        case "delete":
            result = DELETE_MANAGEDDEVICEREMOTEACTION
        case "fullScan":
            result = FULLSCAN_MANAGEDDEVICEREMOTEACTION
        case "quickScan":
            result = QUICKSCAN_MANAGEDDEVICEREMOTEACTION
        case "signatureUpdate":
            result = SIGNATUREUPDATE_MANAGEDDEVICEREMOTEACTION
        case "wipe":
            result = WIPE_MANAGEDDEVICEREMOTEACTION
        case "customTextNotification":
            result = CUSTOMTEXTNOTIFICATION_MANAGEDDEVICEREMOTEACTION
        case "rebootNow":
            result = REBOOTNOW_MANAGEDDEVICEREMOTEACTION
        case "setDeviceName":
            result = SETDEVICENAME_MANAGEDDEVICEREMOTEACTION
        case "syncDevice":
            result = SYNCDEVICE_MANAGEDDEVICEREMOTEACTION
        case "deprovision":
            result = DEPROVISION_MANAGEDDEVICEREMOTEACTION
        case "disable":
            result = DISABLE_MANAGEDDEVICEREMOTEACTION
        case "reenable":
            result = REENABLE_MANAGEDDEVICEREMOTEACTION
        case "moveDeviceToOrganizationalUnit":
            result = MOVEDEVICETOORGANIZATIONALUNIT_MANAGEDDEVICEREMOTEACTION
        case "activateDeviceEsim":
            result = ACTIVATEDEVICEESIM_MANAGEDDEVICEREMOTEACTION
        case "collectDiagnostics":
            result = COLLECTDIAGNOSTICS_MANAGEDDEVICEREMOTEACTION
        case "initiateMobileDeviceManagementKeyRecovery":
            result = INITIATEMOBILEDEVICEMANAGEMENTKEYRECOVERY_MANAGEDDEVICEREMOTEACTION
        default:
            return 0, errors.New("Unknown ManagedDeviceRemoteAction value: " + v)
    }
    return &result, nil
}
func SerializeManagedDeviceRemoteAction(values []ManagedDeviceRemoteAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
