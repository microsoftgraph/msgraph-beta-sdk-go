package models
// The publishing cadence of the quality update. Possible values are: monthly, outOfBand. This property cannot be modified and is automatically populated when the catalog is created.
type WindowsQualityUpdateCadence int

const (
    // Default. Indicates the quality update is released in a regular monthly pattern.
    MONTHLY_WINDOWSQUALITYUPDATECADENCE WindowsQualityUpdateCadence = iota
    // Indicates the quality update is released in an out-of-band pattern.
    OUTOFBAND_WINDOWSQUALITYUPDATECADENCE
    // Evolvable enumeration sentinel value. Do not use.
    UNKNOWNFUTUREVALUE_WINDOWSQUALITYUPDATECADENCE
)

func (i WindowsQualityUpdateCadence) String() string {
    return []string{"monthly", "outOfBand", "unknownFutureValue"}[i]
}
func ParseWindowsQualityUpdateCadence(v string) (any, error) {
    result := MONTHLY_WINDOWSQUALITYUPDATECADENCE
    switch v {
        case "monthly":
            result = MONTHLY_WINDOWSQUALITYUPDATECADENCE
        case "outOfBand":
            result = OUTOFBAND_WINDOWSQUALITYUPDATECADENCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_WINDOWSQUALITYUPDATECADENCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindowsQualityUpdateCadence(values []WindowsQualityUpdateCadence) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i WindowsQualityUpdateCadence) isMultiValue() bool {
    return false
}
