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
    switch strings.ToUpper(v) {
        case "NAM":
            return NAM_CONNECTORGROUPREGION, nil
        case "EUR":
            return EUR_CONNECTORGROUPREGION, nil
        case "AUS":
            return AUS_CONNECTORGROUPREGION, nil
        case "ASIA":
            return ASIA_CONNECTORGROUPREGION, nil
        case "IND":
            return IND_CONNECTORGROUPREGION, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_CONNECTORGROUPREGION, nil
    }
    return 0, errors.New("Unknown ConnectorGroupRegion value: " + v)
}
func SerializeConnectorGroupRegion(values []ConnectorGroupRegion) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
