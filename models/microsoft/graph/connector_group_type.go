package graph
import (
    "strings"
    "errors"
)
// 
type ConnectorGroupType int

const (
    APPLICATIONPROXY_CONNECTORGROUPTYPE ConnectorGroupType = iota
)

func (i ConnectorGroupType) String() string {
    return []string{"APPLICATIONPROXY"}[i]
}
func ParseConnectorGroupType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "APPLICATIONPROXY":
            return APPLICATIONPROXY_CONNECTORGROUPTYPE, nil
    }
    return 0, errors.New("Unknown ConnectorGroupType value: " + v)
}
func SerializeConnectorGroupType(values []ConnectorGroupType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
