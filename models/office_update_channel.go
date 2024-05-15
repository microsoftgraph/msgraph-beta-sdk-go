package models
// The Enum to specify the Office365 Updates Channel.
type OfficeUpdateChannel int

const (
    NONE_OFFICEUPDATECHANNEL OfficeUpdateChannel = iota
    CURRENT_OFFICEUPDATECHANNEL
    DEFERRED_OFFICEUPDATECHANNEL
    FIRSTRELEASECURRENT_OFFICEUPDATECHANNEL
    FIRSTRELEASEDEFERRED_OFFICEUPDATECHANNEL
    MONTHLYENTERPRISE_OFFICEUPDATECHANNEL
)

func (i OfficeUpdateChannel) String() string {
    return []string{"none", "current", "deferred", "firstReleaseCurrent", "firstReleaseDeferred", "monthlyEnterprise"}[i]
}
func ParseOfficeUpdateChannel(v string) (any, error) {
    result := NONE_OFFICEUPDATECHANNEL
    switch v {
        case "none":
            result = NONE_OFFICEUPDATECHANNEL
        case "current":
            result = CURRENT_OFFICEUPDATECHANNEL
        case "deferred":
            result = DEFERRED_OFFICEUPDATECHANNEL
        case "firstReleaseCurrent":
            result = FIRSTRELEASECURRENT_OFFICEUPDATECHANNEL
        case "firstReleaseDeferred":
            result = FIRSTRELEASEDEFERRED_OFFICEUPDATECHANNEL
        case "monthlyEnterprise":
            result = MONTHLYENTERPRISE_OFFICEUPDATECHANNEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOfficeUpdateChannel(values []OfficeUpdateChannel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OfficeUpdateChannel) isMultiValue() bool {
    return false
}
