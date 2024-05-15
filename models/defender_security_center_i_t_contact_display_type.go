package models
// Possible values for defenderSecurityCenterITContactDisplay
type DefenderSecurityCenterITContactDisplayType int

const (
    // Not Configured
    NOTCONFIGURED_DEFENDERSECURITYCENTERITCONTACTDISPLAYTYPE DefenderSecurityCenterITContactDisplayType = iota
    // Display in app and in notifications
    DISPLAYINAPPANDINNOTIFICATIONS_DEFENDERSECURITYCENTERITCONTACTDISPLAYTYPE
    // Display only in app
    DISPLAYONLYINAPP_DEFENDERSECURITYCENTERITCONTACTDISPLAYTYPE
    // Display only in notifications
    DISPLAYONLYINNOTIFICATIONS_DEFENDERSECURITYCENTERITCONTACTDISPLAYTYPE
)

func (i DefenderSecurityCenterITContactDisplayType) String() string {
    return []string{"notConfigured", "displayInAppAndInNotifications", "displayOnlyInApp", "displayOnlyInNotifications"}[i]
}
func ParseDefenderSecurityCenterITContactDisplayType(v string) (any, error) {
    result := NOTCONFIGURED_DEFENDERSECURITYCENTERITCONTACTDISPLAYTYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_DEFENDERSECURITYCENTERITCONTACTDISPLAYTYPE
        case "displayInAppAndInNotifications":
            result = DISPLAYINAPPANDINNOTIFICATIONS_DEFENDERSECURITYCENTERITCONTACTDISPLAYTYPE
        case "displayOnlyInApp":
            result = DISPLAYONLYINAPP_DEFENDERSECURITYCENTERITCONTACTDISPLAYTYPE
        case "displayOnlyInNotifications":
            result = DISPLAYONLYINNOTIFICATIONS_DEFENDERSECURITYCENTERITCONTACTDISPLAYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDefenderSecurityCenterITContactDisplayType(values []DefenderSecurityCenterITContactDisplayType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DefenderSecurityCenterITContactDisplayType) isMultiValue() bool {
    return false
}
