package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_CONFIGURATIONMANAGERCLIENTSTATE, nil
        case "INSTALLED":
            return INSTALLED_CONFIGURATIONMANAGERCLIENTSTATE, nil
        case "HEALTHY":
            return HEALTHY_CONFIGURATIONMANAGERCLIENTSTATE, nil
        case "INSTALLFAILED":
            return INSTALLFAILED_CONFIGURATIONMANAGERCLIENTSTATE, nil
        case "UPDATEFAILED":
            return UPDATEFAILED_CONFIGURATIONMANAGERCLIENTSTATE, nil
        case "COMMUNICATIONERROR":
            return COMMUNICATIONERROR_CONFIGURATIONMANAGERCLIENTSTATE, nil
    }
    return 0, errors.New("Unknown ConfigurationManagerClientState value: " + v)
}
func SerializeConfigurationManagerClientState(values []ConfigurationManagerClientState) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
