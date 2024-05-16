package security
type DeviceAssetIdentifier int

const (
    DEVICEID_DEVICEASSETIDENTIFIER DeviceAssetIdentifier = iota
    DEVICENAME_DEVICEASSETIDENTIFIER
    REMOTEDEVICENAME_DEVICEASSETIDENTIFIER
    TARGETDEVICENAME_DEVICEASSETIDENTIFIER
    DESTINATIONDEVICENAME_DEVICEASSETIDENTIFIER
    UNKNOWNFUTUREVALUE_DEVICEASSETIDENTIFIER
)

func (i DeviceAssetIdentifier) String() string {
    return []string{"deviceId", "deviceName", "remoteDeviceName", "targetDeviceName", "destinationDeviceName", "unknownFutureValue"}[i]
}
func ParseDeviceAssetIdentifier(v string) (any, error) {
    result := DEVICEID_DEVICEASSETIDENTIFIER
    switch v {
        case "deviceId":
            result = DEVICEID_DEVICEASSETIDENTIFIER
        case "deviceName":
            result = DEVICENAME_DEVICEASSETIDENTIFIER
        case "remoteDeviceName":
            result = REMOTEDEVICENAME_DEVICEASSETIDENTIFIER
        case "targetDeviceName":
            result = TARGETDEVICENAME_DEVICEASSETIDENTIFIER
        case "destinationDeviceName":
            result = DESTINATIONDEVICENAME_DEVICEASSETIDENTIFIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEASSETIDENTIFIER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeviceAssetIdentifier(values []DeviceAssetIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceAssetIdentifier) isMultiValue() bool {
    return false
}
