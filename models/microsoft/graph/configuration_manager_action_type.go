package graph
import (
    "strings"
    "errors"
)
type ConfigurationManagerActionType int

const (
    REFRESHMACHINEPOLICY_CONFIGURATIONMANAGERACTIONTYPE ConfigurationManagerActionType = iota
    REFRESHUSERPOLICY_CONFIGURATIONMANAGERACTIONTYPE
    WAKEUPCLIENT_CONFIGURATIONMANAGERACTIONTYPE
    APPEVALUATION_CONFIGURATIONMANAGERACTIONTYPE
)

func (i ConfigurationManagerActionType) String() string {
    return []string{"REFRESHMACHINEPOLICY", "REFRESHUSERPOLICY", "WAKEUPCLIENT", "APPEVALUATION"}[i]
}
func ParseConfigurationManagerActionType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "REFRESHMACHINEPOLICY":
            return REFRESHMACHINEPOLICY_CONFIGURATIONMANAGERACTIONTYPE, nil
        case "REFRESHUSERPOLICY":
            return REFRESHUSERPOLICY_CONFIGURATIONMANAGERACTIONTYPE, nil
        case "WAKEUPCLIENT":
            return WAKEUPCLIENT_CONFIGURATIONMANAGERACTIONTYPE, nil
        case "APPEVALUATION":
            return APPEVALUATION_CONFIGURATIONMANAGERACTIONTYPE, nil
    }
    return 0, errors.New("Unknown ConfigurationManagerActionType value: " + v)
}
func SerializeConfigurationManagerActionType(values []ConfigurationManagerActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
