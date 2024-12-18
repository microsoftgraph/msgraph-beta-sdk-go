package models
type ChangeAnnouncementChangeType int

const (
    BREAKINGCHANGE_CHANGEANNOUNCEMENTCHANGETYPE ChangeAnnouncementChangeType = iota
    DEPRECATION_CHANGEANNOUNCEMENTCHANGETYPE
    ENDOFSUPPORT_CHANGEANNOUNCEMENTCHANGETYPE
    FEATURECHANGE_CHANGEANNOUNCEMENTCHANGETYPE
    OTHER_CHANGEANNOUNCEMENTCHANGETYPE
    RETIREMENT_CHANGEANNOUNCEMENTCHANGETYPE
    SECURITYINCIDENT_CHANGEANNOUNCEMENTCHANGETYPE
    UXCHANGE_CHANGEANNOUNCEMENTCHANGETYPE
    UNKNOWNFUTUREVALUE_CHANGEANNOUNCEMENTCHANGETYPE
)

func (i ChangeAnnouncementChangeType) String() string {
    return []string{"breakingChange", "deprecation", "endOfSupport", "featureChange", "other", "retirement", "securityIncident", "uxChange", "unknownFutureValue"}[i]
}
func ParseChangeAnnouncementChangeType(v string) (any, error) {
    result := BREAKINGCHANGE_CHANGEANNOUNCEMENTCHANGETYPE
    switch v {
        case "breakingChange":
            result = BREAKINGCHANGE_CHANGEANNOUNCEMENTCHANGETYPE
        case "deprecation":
            result = DEPRECATION_CHANGEANNOUNCEMENTCHANGETYPE
        case "endOfSupport":
            result = ENDOFSUPPORT_CHANGEANNOUNCEMENTCHANGETYPE
        case "featureChange":
            result = FEATURECHANGE_CHANGEANNOUNCEMENTCHANGETYPE
        case "other":
            result = OTHER_CHANGEANNOUNCEMENTCHANGETYPE
        case "retirement":
            result = RETIREMENT_CHANGEANNOUNCEMENTCHANGETYPE
        case "securityIncident":
            result = SECURITYINCIDENT_CHANGEANNOUNCEMENTCHANGETYPE
        case "uxChange":
            result = UXCHANGE_CHANGEANNOUNCEMENTCHANGETYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CHANGEANNOUNCEMENTCHANGETYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeChangeAnnouncementChangeType(values []ChangeAnnouncementChangeType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ChangeAnnouncementChangeType) isMultiValue() bool {
    return false
}
