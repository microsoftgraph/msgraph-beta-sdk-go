package models
// The Enum to specify the level of display for the Installation Progress Setup UI on the Device.
type OfficeSuiteInstallProgressDisplayLevel int

const (
    NONE_OFFICESUITEINSTALLPROGRESSDISPLAYLEVEL OfficeSuiteInstallProgressDisplayLevel = iota
    FULL_OFFICESUITEINSTALLPROGRESSDISPLAYLEVEL
)

func (i OfficeSuiteInstallProgressDisplayLevel) String() string {
    return []string{"none", "full"}[i]
}
func ParseOfficeSuiteInstallProgressDisplayLevel(v string) (any, error) {
    result := NONE_OFFICESUITEINSTALLPROGRESSDISPLAYLEVEL
    switch v {
        case "none":
            result = NONE_OFFICESUITEINSTALLPROGRESSDISPLAYLEVEL
        case "full":
            result = FULL_OFFICESUITEINSTALLPROGRESSDISPLAYLEVEL
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeOfficeSuiteInstallProgressDisplayLevel(values []OfficeSuiteInstallProgressDisplayLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i OfficeSuiteInstallProgressDisplayLevel) isMultiValue() bool {
    return false
}
