package models
import (
    "errors"
)
// Android Device Owner possible values for states of the device's plugged-in power modes.
type AndroidDeviceOwnerBatteryPluggedMode int

const (
    // Not configured; this value is ignored.
    NOTCONFIGURED_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE AndroidDeviceOwnerBatteryPluggedMode = iota
    // Power source is an AC charger.
    AC_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
    // Power source is a USB port.
    USB_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
    // Power source is wireless.
    WIRELESS_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
)

func (i AndroidDeviceOwnerBatteryPluggedMode) String() string {
    return []string{"notConfigured", "ac", "usb", "wireless"}[i]
}
func ParseAndroidDeviceOwnerBatteryPluggedMode(v string) (any, error) {
    result := NOTCONFIGURED_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
        case "ac":
            result = AC_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
        case "usb":
            result = USB_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
        case "wireless":
            result = WIRELESS_ANDROIDDEVICEOWNERBATTERYPLUGGEDMODE
        default:
            return 0, errors.New("Unknown AndroidDeviceOwnerBatteryPluggedMode value: " + v)
    }
    return &result, nil
}
func SerializeAndroidDeviceOwnerBatteryPluggedMode(values []AndroidDeviceOwnerBatteryPluggedMode) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidDeviceOwnerBatteryPluggedMode) isMultiValue() bool {
    return false
}
