package models
// System Update Types for Android Device Owner.
type AndroidDeviceOwnerSystemUpdateInstallType int

const (
    // Device default behavior, which typically prompts the user to accept system updates.
    DEVICEDEFAULT_ANDROIDDEVICEOWNERSYSTEMUPDATEINSTALLTYPE AndroidDeviceOwnerSystemUpdateInstallType = iota
    // Postpone automatic install of updates up to 30 days.
    POSTPONE_ANDROIDDEVICEOWNERSYSTEMUPDATEINSTALLTYPE
    // Install automatically inside a daily maintenance window.
    WINDOWED_ANDROIDDEVICEOWNERSYSTEMUPDATEINSTALLTYPE
    // Automatically install updates as soon as possible.
    AUTOMATIC_ANDROIDDEVICEOWNERSYSTEMUPDATEINSTALLTYPE
)

func (i AndroidDeviceOwnerSystemUpdateInstallType) String() string {
    return []string{"deviceDefault", "postpone", "windowed", "automatic"}[i]
}
func ParseAndroidDeviceOwnerSystemUpdateInstallType(v string) (any, error) {
    result := DEVICEDEFAULT_ANDROIDDEVICEOWNERSYSTEMUPDATEINSTALLTYPE
    switch v {
        case "deviceDefault":
            result = DEVICEDEFAULT_ANDROIDDEVICEOWNERSYSTEMUPDATEINSTALLTYPE
        case "postpone":
            result = POSTPONE_ANDROIDDEVICEOWNERSYSTEMUPDATEINSTALLTYPE
        case "windowed":
            result = WINDOWED_ANDROIDDEVICEOWNERSYSTEMUPDATEINSTALLTYPE
        case "automatic":
            result = AUTOMATIC_ANDROIDDEVICEOWNERSYSTEMUPDATEINSTALLTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerSystemUpdateInstallType(values []AndroidDeviceOwnerSystemUpdateInstallType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerSystemUpdateInstallType) isMultiValue() bool {
    return false
}
