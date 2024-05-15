package models
// Possible values for the EdgeOpensWith setting.
type EdgeOpenOptions int

const (
    // Not configured.
    NOTCONFIGURED_EDGEOPENOPTIONS EdgeOpenOptions = iota
    // StartPage.
    STARTPAGE_EDGEOPENOPTIONS
    // NewTabPage.
    NEWTABPAGE_EDGEOPENOPTIONS
    // PreviousPages.
    PREVIOUSPAGES_EDGEOPENOPTIONS
    // SpecificPages.
    SPECIFICPAGES_EDGEOPENOPTIONS
)

func (i EdgeOpenOptions) String() string {
    return []string{"notConfigured", "startPage", "newTabPage", "previousPages", "specificPages"}[i]
}
func ParseEdgeOpenOptions(v string) (any, error) {
    result := NOTCONFIGURED_EDGEOPENOPTIONS
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_EDGEOPENOPTIONS
        case "startPage":
            result = STARTPAGE_EDGEOPENOPTIONS
        case "newTabPage":
            result = NEWTABPAGE_EDGEOPENOPTIONS
        case "previousPages":
            result = PREVIOUSPAGES_EDGEOPENOPTIONS
        case "specificPages":
            result = SPECIFICPAGES_EDGEOPENOPTIONS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEdgeOpenOptions(values []EdgeOpenOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EdgeOpenOptions) isMultiValue() bool {
    return false
}
