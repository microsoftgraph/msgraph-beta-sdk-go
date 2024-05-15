package windowsupdates
type QualityUpdateCadence int

const (
    MONTHLY_QUALITYUPDATECADENCE QualityUpdateCadence = iota
    OUTOFBAND_QUALITYUPDATECADENCE
    UNKNOWNFUTUREVALUE_QUALITYUPDATECADENCE
)

func (i QualityUpdateCadence) String() string {
    return []string{"monthly", "outOfBand", "unknownFutureValue"}[i]
}
func ParseQualityUpdateCadence(v string) (any, error) {
    result := MONTHLY_QUALITYUPDATECADENCE
    switch v {
        case "monthly":
            result = MONTHLY_QUALITYUPDATECADENCE
        case "outOfBand":
            result = OUTOFBAND_QUALITYUPDATECADENCE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_QUALITYUPDATECADENCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeQualityUpdateCadence(values []QualityUpdateCadence) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i QualityUpdateCadence) isMultiValue() bool {
    return false
}
