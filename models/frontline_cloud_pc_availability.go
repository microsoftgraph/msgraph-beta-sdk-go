package models
type FrontlineCloudPcAvailability int

const (
    NOTAPPLICABLE_FRONTLINECLOUDPCAVAILABILITY FrontlineCloudPcAvailability = iota
    AVAILABLE_FRONTLINECLOUDPCAVAILABILITY
    NOTAVAILABLE_FRONTLINECLOUDPCAVAILABILITY
    UNKNOWNFUTUREVALUE_FRONTLINECLOUDPCAVAILABILITY
)

func (i FrontlineCloudPcAvailability) String() string {
    return []string{"notApplicable", "available", "notAvailable", "unknownFutureValue"}[i]
}
func ParseFrontlineCloudPcAvailability(v string) (any, error) {
    result := NOTAPPLICABLE_FRONTLINECLOUDPCAVAILABILITY
    switch v {
        case "notApplicable":
            result = NOTAPPLICABLE_FRONTLINECLOUDPCAVAILABILITY
        case "available":
            result = AVAILABLE_FRONTLINECLOUDPCAVAILABILITY
        case "notAvailable":
            result = NOTAVAILABLE_FRONTLINECLOUDPCAVAILABILITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_FRONTLINECLOUDPCAVAILABILITY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeFrontlineCloudPcAvailability(values []FrontlineCloudPcAvailability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i FrontlineCloudPcAvailability) isMultiValue() bool {
    return false
}
