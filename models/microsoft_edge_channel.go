package models
import (
    "errors"
)
// Provides operations to manage the collection of accessReview entities.
type MicrosoftEdgeChannel int

const (
    DEV_MICROSOFTEDGECHANNEL MicrosoftEdgeChannel = iota
    BETA_MICROSOFTEDGECHANNEL
    STABLE_MICROSOFTEDGECHANNEL
)

func (i MicrosoftEdgeChannel) String() string {
    return []string{"dev", "beta", "stable"}[i]
}
func ParseMicrosoftEdgeChannel(v string) (interface{}, error) {
    result := DEV_MICROSOFTEDGECHANNEL
    switch v {
        case "dev":
            result = DEV_MICROSOFTEDGECHANNEL
        case "beta":
            result = BETA_MICROSOFTEDGECHANNEL
        case "stable":
            result = STABLE_MICROSOFTEDGECHANNEL
        default:
            return 0, errors.New("Unknown MicrosoftEdgeChannel value: " + v)
    }
    return &result, nil
}
func SerializeMicrosoftEdgeChannel(values []MicrosoftEdgeChannel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
