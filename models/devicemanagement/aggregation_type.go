package devicemanagement
type AggregationType int

const (
    COUNT_AGGREGATIONTYPE AggregationType = iota
    PERCENTAGE_AGGREGATIONTYPE
    AFFECTEDCLOUDPCCOUNT_AGGREGATIONTYPE
    AFFECTEDCLOUDPCPERCENTAGE_AGGREGATIONTYPE
    UNKNOWNFUTUREVALUE_AGGREGATIONTYPE
)

func (i AggregationType) String() string {
    return []string{"count", "percentage", "affectedCloudPcCount", "affectedCloudPcPercentage", "unknownFutureValue"}[i]
}
func ParseAggregationType(v string) (any, error) {
    result := COUNT_AGGREGATIONTYPE
    switch v {
        case "count":
            result = COUNT_AGGREGATIONTYPE
        case "percentage":
            result = PERCENTAGE_AGGREGATIONTYPE
        case "affectedCloudPcCount":
            result = AFFECTEDCLOUDPCCOUNT_AGGREGATIONTYPE
        case "affectedCloudPcPercentage":
            result = AFFECTEDCLOUDPCPERCENTAGE_AGGREGATIONTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_AGGREGATIONTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAggregationType(values []AggregationType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AggregationType) isMultiValue() bool {
    return false
}
