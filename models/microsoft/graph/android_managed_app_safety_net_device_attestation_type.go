package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE, nil
        case "BASICINTEGRITY":
            return BASICINTEGRITY_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE, nil
        case "BASICINTEGRITYANDDEVICECERTIFICATION":
            return BASICINTEGRITYANDDEVICECERTIFICATION_ANDROIDMANAGEDAPPSAFETYNETDEVICEATTESTATIONTYPE, nil
    }
    return 0, errors.New("Unknown AndroidManagedAppSafetyNetDeviceAttestationType value: " + v)
}
func SerializeAndroidManagedAppSafetyNetDeviceAttestationType(values []AndroidManagedAppSafetyNetDeviceAttestationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
