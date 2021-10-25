package graph
import (
    "strings"
    "errors"
)
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
)

func (i ManagedDeviceRemoteAction) String() string {
    return []string{"RETIRE", "DELETE", "FULLSCAN", "QUICKSCAN", "SIGNATUREUPDATE", "WIPE", "CUSTOMTEXTNOTIFICATION", "REBOOTNOW", "SETDEVICENAME", "SYNCDEVICE"}[i]
}
func ParseManagedDeviceRemoteAction(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "RETIRE":
            return RETIRE_MANAGEDDEVICEREMOTEACTION, nil
        case "DELETE":
            return DELETE_MANAGEDDEVICEREMOTEACTION, nil
        case "FULLSCAN":
            return FULLSCAN_MANAGEDDEVICEREMOTEACTION, nil
        case "QUICKSCAN":
            return QUICKSCAN_MANAGEDDEVICEREMOTEACTION, nil
        case "SIGNATUREUPDATE":
            return SIGNATUREUPDATE_MANAGEDDEVICEREMOTEACTION, nil
        case "WIPE":
            return WIPE_MANAGEDDEVICEREMOTEACTION, nil
        case "CUSTOMTEXTNOTIFICATION":
            return CUSTOMTEXTNOTIFICATION_MANAGEDDEVICEREMOTEACTION, nil
        case "REBOOTNOW":
            return REBOOTNOW_MANAGEDDEVICEREMOTEACTION, nil
        case "SETDEVICENAME":
            return SETDEVICENAME_MANAGEDDEVICEREMOTEACTION, nil
        case "SYNCDEVICE":
            return SYNCDEVICE_MANAGEDDEVICEREMOTEACTION, nil
    }
    return 0, errors.New("Unknown ManagedDeviceRemoteAction value: " + v)
}
func SerializeManagedDeviceRemoteAction(values []ManagedDeviceRemoteAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
