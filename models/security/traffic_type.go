package security
type TrafficType int

const (
    DOWNLOADEDBYTES_TRAFFICTYPE TrafficType = iota
    UPLOADEDBYTES_TRAFFICTYPE
    UNKNOWN_TRAFFICTYPE
    UNKNOWNFUTUREVALUE_TRAFFICTYPE
)

func (i TrafficType) String() string {
    return []string{"downloadedBytes", "uploadedBytes", "unknown", "unknownFutureValue"}[i]
}
func ParseTrafficType(v string) (any, error) {
    result := DOWNLOADEDBYTES_TRAFFICTYPE
    switch v {
        case "downloadedBytes":
            result = DOWNLOADEDBYTES_TRAFFICTYPE
        case "uploadedBytes":
            result = UPLOADEDBYTES_TRAFFICTYPE
        case "unknown":
            result = UNKNOWN_TRAFFICTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TRAFFICTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeTrafficType(values []TrafficType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TrafficType) isMultiValue() bool {
    return false
}
