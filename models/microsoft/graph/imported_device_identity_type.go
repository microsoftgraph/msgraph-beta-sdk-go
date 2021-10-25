package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_IMPORTEDDEVICEIDENTITYTYPE, nil
        case "IMEI":
            return IMEI_IMPORTEDDEVICEIDENTITYTYPE, nil
        case "SERIALNUMBER":
            return SERIALNUMBER_IMPORTEDDEVICEIDENTITYTYPE, nil
    }
    return 0, errors.New("Unknown ImportedDeviceIdentityType value: " + v)
}
func SerializeImportedDeviceIdentityType(values []ImportedDeviceIdentityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
