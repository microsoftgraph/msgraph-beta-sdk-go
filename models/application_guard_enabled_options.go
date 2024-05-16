package models
// Possible values for ApplicationGuardEnabledOptions
type ApplicationGuardEnabledOptions int

const (
    // Not Configured
    NOTCONFIGURED_APPLICATIONGUARDENABLEDOPTIONS ApplicationGuardEnabledOptions = iota
    // Enabled For Edge
    ENABLEDFOREDGE_APPLICATIONGUARDENABLEDOPTIONS
    // Enabled For Office
    ENABLEDFOROFFICE_APPLICATIONGUARDENABLEDOPTIONS
    // Enabled For Edge And Office
    ENABLEDFOREDGEANDOFFICE_APPLICATIONGUARDENABLEDOPTIONS
)

func (i ApplicationGuardEnabledOptions) String() string {
    return []string{"notConfigured", "enabledForEdge", "enabledForOffice", "enabledForEdgeAndOffice"}[i]
}
func ParseApplicationGuardEnabledOptions(v string) (any, error) {
    result := NOTCONFIGURED_APPLICATIONGUARDENABLEDOPTIONS
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_APPLICATIONGUARDENABLEDOPTIONS
        case "enabledForEdge":
            result = ENABLEDFOREDGE_APPLICATIONGUARDENABLEDOPTIONS
        case "enabledForOffice":
            result = ENABLEDFOROFFICE_APPLICATIONGUARDENABLEDOPTIONS
        case "enabledForEdgeAndOffice":
            result = ENABLEDFOREDGEANDOFFICE_APPLICATIONGUARDENABLEDOPTIONS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeApplicationGuardEnabledOptions(values []ApplicationGuardEnabledOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ApplicationGuardEnabledOptions) isMultiValue() bool {
    return false
}
