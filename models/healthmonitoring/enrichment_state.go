package healthmonitoring
type EnrichmentState int

const (
    NONE_ENRICHMENTSTATE EnrichmentState = iota
    INPROGRESS_ENRICHMENTSTATE
    ENRICHED_ENRICHMENTSTATE
    UNKNOWNFUTUREVALUE_ENRICHMENTSTATE
)

func (i EnrichmentState) String() string {
    return []string{"none", "inProgress", "enriched", "unknownFutureValue"}[i]
}
func ParseEnrichmentState(v string) (any, error) {
    result := NONE_ENRICHMENTSTATE
    switch v {
        case "none":
            result = NONE_ENRICHMENTSTATE
        case "inProgress":
            result = INPROGRESS_ENRICHMENTSTATE
        case "enriched":
            result = ENRICHED_ENRICHMENTSTATE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENRICHMENTSTATE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEnrichmentState(values []EnrichmentState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EnrichmentState) isMultiValue() bool {
    return false
}
