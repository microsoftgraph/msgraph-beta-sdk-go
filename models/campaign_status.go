package models
import (
    "errors"
)
type CampaignStatus int

const (
    UNKNOWN_CAMPAIGNSTATUS CampaignStatus = iota
    DRAFT_CAMPAIGNSTATUS
    INPROGRESS_CAMPAIGNSTATUS
    SCHEDULED_CAMPAIGNSTATUS
    COMPLETED_CAMPAIGNSTATUS
    FAILED_CAMPAIGNSTATUS
    CANCELLED_CAMPAIGNSTATUS
    EXCLUDED_CAMPAIGNSTATUS
    DELETED_CAMPAIGNSTATUS
    UNKNOWNFUTUREVALUE_CAMPAIGNSTATUS
)

func (i CampaignStatus) String() string {
    return []string{"unknown", "draft", "inProgress", "scheduled", "completed", "failed", "cancelled", "excluded", "deleted", "unknownFutureValue"}[i]
}
func ParseCampaignStatus(v string) (any, error) {
    result := UNKNOWN_CAMPAIGNSTATUS
    switch v {
        case "unknown":
            result = UNKNOWN_CAMPAIGNSTATUS
        case "draft":
            result = DRAFT_CAMPAIGNSTATUS
        case "inProgress":
            result = INPROGRESS_CAMPAIGNSTATUS
        case "scheduled":
            result = SCHEDULED_CAMPAIGNSTATUS
        case "completed":
            result = COMPLETED_CAMPAIGNSTATUS
        case "failed":
            result = FAILED_CAMPAIGNSTATUS
        case "cancelled":
            result = CANCELLED_CAMPAIGNSTATUS
        case "excluded":
            result = EXCLUDED_CAMPAIGNSTATUS
        case "deleted":
            result = DELETED_CAMPAIGNSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CAMPAIGNSTATUS
        default:
            return 0, errors.New("Unknown CampaignStatus value: " + v)
    }
    return &result, nil
}
func SerializeCampaignStatus(values []CampaignStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i CampaignStatus) isMultiValue() bool {
    return false
}
