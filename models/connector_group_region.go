package models
type ConnectorGroupRegion int

const (
    NAM_CONNECTORGROUPREGION ConnectorGroupRegion = iota
    EUR_CONNECTORGROUPREGION
    AUS_CONNECTORGROUPREGION
    ASIA_CONNECTORGROUPREGION
    IND_CONNECTORGROUPREGION
    UNKNOWNFUTUREVALUE_CONNECTORGROUPREGION
)

func (i ConnectorGroupRegion) String() string {
    return []string{"nam", "eur", "aus", "asia", "ind", "unknownFutureValue"}[i]
}
func ParseConnectorGroupRegion(v string) (any, error) {
    result := NAM_CONNECTORGROUPREGION
    switch v {
        case "nam":
            result = NAM_CONNECTORGROUPREGION
        case "eur":
            result = EUR_CONNECTORGROUPREGION
        case "aus":
            result = AUS_CONNECTORGROUPREGION
        case "asia":
            result = ASIA_CONNECTORGROUPREGION
        case "ind":
            result = IND_CONNECTORGROUPREGION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONNECTORGROUPREGION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConnectorGroupRegion(values []ConnectorGroupRegion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConnectorGroupRegion) isMultiValue() bool {
    return false
}
