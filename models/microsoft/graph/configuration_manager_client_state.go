package graph
import (
    "strings"
    "errors"
)
// 
type ConfigurationManagerClientState int

const (
    UNKNOWN_CONFIGURATIONMANAGERCLIENTSTATE ConfigurationManagerClientState = iota
    INSTALLED_CONFIGURATIONMANAGERCLIENTSTATE
    HEALTHY_CONFIGURATIONMANAGERCLIENTSTATE
    INSTALLFAILED_CONFIGURATIONMANAGERCLIENTSTATE
    UPDATEFAILED_CONFIGURATIONMANAGERCLIENTSTATE
    COMMUNICATIONERROR_CONFIGURATIONMANAGERCLIENTSTATE
)

func (i ConfigurationManagerClientState) String() string {
    return []string{"UNKNOWN", "INSTALLED", "HEALTHY", "INSTALLFAILED", "UPDATEFAILED", "COMMUNICATIONERROR"}[i]
}
func ParseConfigurationManagerClientState(v string) (interface{}, error) {
    result := UNKNOWN_CONFIGURATIONMANAGERCLIENTSTATE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_CONFIGURATIONMANAGERCLIENTSTATE
        case "INSTALLED":
            result = INSTALLED_CONFIGURATIONMANAGERCLIENTSTATE
        case "HEALTHY":
            result = HEALTHY_CONFIGURATIONMANAGERCLIENTSTATE
        case "INSTALLFAILED":
            result = INSTALLFAILED_CONFIGURATIONMANAGERCLIENTSTATE
        case "UPDATEFAILED":
            result = UPDATEFAILED_CONFIGURATIONMANAGERCLIENTSTATE
        case "COMMUNICATIONERROR":
            result = COMMUNICATIONERROR_CONFIGURATIONMANAGERCLIENTSTATE
        default:
            return 0, errors.New("Unknown ConfigurationManagerClientState value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationManagerClientState(values []ConfigurationManagerClientState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
