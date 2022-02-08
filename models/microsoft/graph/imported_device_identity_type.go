package graph
import (
    "strings"
    "errors"
)
// 
type ImportedDeviceIdentityType int

const (
    UNKNOWN_IMPORTEDDEVICEIDENTITYTYPE ImportedDeviceIdentityType = iota
    IMEI_IMPORTEDDEVICEIDENTITYTYPE
    SERIALNUMBER_IMPORTEDDEVICEIDENTITYTYPE
)

func (i ImportedDeviceIdentityType) String() string {
    return []string{"UNKNOWN", "IMEI", "SERIALNUMBER"}[i]
}
func ParseImportedDeviceIdentityType(v string) (interface{}, error) {
    result := UNKNOWN_IMPORTEDDEVICEIDENTITYTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_IMPORTEDDEVICEIDENTITYTYPE
        case "IMEI":
            result = IMEI_IMPORTEDDEVICEIDENTITYTYPE
        case "SERIALNUMBER":
            result = SERIALNUMBER_IMPORTEDDEVICEIDENTITYTYPE
        default:
            return 0, errors.New("Unknown ImportedDeviceIdentityType value: " + v)
    }
    return &result, nil
}
func SerializeImportedDeviceIdentityType(values []ImportedDeviceIdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
