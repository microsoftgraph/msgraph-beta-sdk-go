package models
type ConnectorGroupType int

const (
    APPLICATIONPROXY_CONNECTORGROUPTYPE ConnectorGroupType = iota
)

func (i ConnectorGroupType) String() string {
    return []string{"applicationProxy"}[i]
}
func ParseConnectorGroupType(v string) (any, error) {
    result := APPLICATIONPROXY_CONNECTORGROUPTYPE
    switch v {
        case "applicationProxy":
            result = APPLICATIONPROXY_CONNECTORGROUPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeConnectorGroupType(values []ConnectorGroupType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConnectorGroupType) isMultiValue() bool {
    return false
}
