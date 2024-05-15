package models
type ImportedDeviceIdentityType int

const (
    // Unknown value of importedDeviceIdentityType.
    UNKNOWN_IMPORTEDDEVICEIDENTITYTYPE ImportedDeviceIdentityType = iota
    // Device Identity is of type imei.
    IMEI_IMPORTEDDEVICEIDENTITYTYPE
    // Device Identity is of type serial number.
    SERIALNUMBER_IMPORTEDDEVICEIDENTITYTYPE
    // Device Identity is of type manufacturer + model + serial number semi-colon delimited tuple with enforced order.
    MANUFACTURERMODELSERIAL_IMPORTEDDEVICEIDENTITYTYPE
)

func (i ImportedDeviceIdentityType) String() string {
    return []string{"unknown", "imei", "serialNumber", "manufacturerModelSerial"}[i]
}
func ParseImportedDeviceIdentityType(v string) (any, error) {
    result := UNKNOWN_IMPORTEDDEVICEIDENTITYTYPE
    switch v {
        case "unknown":
            result = UNKNOWN_IMPORTEDDEVICEIDENTITYTYPE
        case "imei":
            result = IMEI_IMPORTEDDEVICEIDENTITYTYPE
        case "serialNumber":
            result = SERIALNUMBER_IMPORTEDDEVICEIDENTITYTYPE
        case "manufacturerModelSerial":
            result = MANUFACTURERMODELSERIAL_IMPORTEDDEVICEIDENTITYTYPE
        default:
            return nil, nil
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
func (i ImportedDeviceIdentityType) isMultiValue() bool {
    return false
}
