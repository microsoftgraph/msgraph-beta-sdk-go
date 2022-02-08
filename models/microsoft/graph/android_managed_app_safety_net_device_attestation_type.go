package graph
import (
    "strings"
    "errors"
)
// 
type AndroidManagedAppSafetyNetDeviceAttestationType int

const (
    NONE_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE AndroidManagedAppSafetyNetDeviceAttestationType = iota
    BASICINTEGRITY_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
    BASICINTEGRITYANDDEVICECERTIFICATION_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
)

func (i AndroidManagedAppSafetyNetDeviceAttestationType) String() string {
    return []string{"NONE", "BASICINTEGRITY", "BASICINTEGRITYANDDEVICECERTIFICATION"}[i]
}
func ParseAndroidManagedAppSafetyNetDeviceAttestationType(v string) (interface{}, error) {
    result := NONE_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
        case "BASICINTEGRITY":
            result = BASICINTEGRITY_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
        case "BASICINTEGRITYANDDEVICECERTIFICATION":
            result = BASICINTEGRITYANDDEVICECERTIFICATION_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE
        default:
            return 0, errors.New("Unknown AndroidManagedAppSafetyNetDeviceAttestationType value: " + v)
    }
    return &result, nil
}
func SerializeAndroidManagedAppSafetyNetDeviceAttestationType(values []AndroidManagedAppSafetyNetDeviceAttestationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
