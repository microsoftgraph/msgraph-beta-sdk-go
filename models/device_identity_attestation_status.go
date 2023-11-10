package models
import (
    "errors"
)
// Indicates the state of the anomaly. Eg: anomaly severity can be new, active, disabled, removed or other.
type DeviceIdentityAttestationStatus int

const (
    // Default. Set to unknown if attestation has not yet been calculated
    UNKNOWN_DEVICEIDENTITYATTESTATIONSTATUS DeviceIdentityAttestationStatus = iota
    // Indicates that the Device attestation is supported on the device, it was attempted on the device and the attestation has passed. The device is trusted.
    TRUSTED_DEVICEIDENTITYATTESTATIONSTATUS
    // Indicates that the Device attestation is supported on the device, it was attempted on the device and the attestation has failed. The device is untrusted
    UNTRUSTED_DEVICEIDENTITYATTESTATIONSTATUS
    // Indicates that the device does not support Attestation. This could be because of missing hardware or software support.
    NOTSUPPORTED_DEVICEIDENTITYATTESTATIONSTATUS
    // Indicates that the device did not provide with the data that were required to perform attestation.
    INCOMPLETEDATA_DEVICEIDENTITYATTESTATIONSTATUS
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_DEVICEIDENTITYATTESTATIONSTATUS
)

func (i DeviceIdentityAttestationStatus) String() string {
    return []string{"unknown", "trusted", "unTrusted", "notSupported", "incompleteData", "unknownFutureValue"}[i]
}
func ParseDeviceIdentityAttestationStatus(v string) (any, error) {
    result := UNKNOWN_DEVICEIDENTITYATTESTATIONSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_DEVICEIDENTITYATTESTATIONSTATUS
        case "trusted":
            result = TRUSTED_DEVICEIDENTITYATTESTATIONSTATUS
        case "unTrusted":
            result = UNTRUSTED_DEVICEIDENTITYATTESTATIONSTATUS
        case "notSupported":
            result = NOTSUPPORTED_DEVICEIDENTITYATTESTATIONSTATUS
        case "incompleteData":
            result = INCOMPLETEDATA_DEVICEIDENTITYATTESTATIONSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEVICEIDENTITYATTESTATIONSTATUS
        default:
            return 0, errors.New("Unknown DeviceIdentityAttestationStatus value: " + v)
    }
    return &result, nil
}
func SerializeDeviceIdentityAttestationStatus(values []DeviceIdentityAttestationStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeviceIdentityAttestationStatus) isMultiValue() bool {
    return false
}
