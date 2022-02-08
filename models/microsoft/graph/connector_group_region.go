package graph
import (
    "strings"
    "errors"
)
// 
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
    return []string{"NAM", "EUR", "AUS", "ASIA", "IND", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseConnectorGroupRegion(v string) (interface{}, error) {
    result := NAM_CONNECTORGROUPREGION
    switch strings.ToUpper(v) {
        case "NAM":
            result = NAM_CONNECTORGROUPREGION
        case "EUR":
            result = EUR_CONNECTORGROUPREGION
        case "AUS":
            result = AUS_CONNECTORGROUPREGION
        case "ASIA":
            result = ASIA_CONNECTORGROUPREGION
        case "IND":
            result = IND_CONNECTORGROUPREGION
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_CONNECTORGROUPREGION
        default:
            return 0, errors.New("Unknown ConnectorGroupRegion value: " + v)
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
