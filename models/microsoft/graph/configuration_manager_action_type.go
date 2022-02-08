package graph
import (
    "strings"
    "errors"
)
// 
type ConfigurationManagerActionType int

const (
    REFRESHMACHINEPOLICY_CONFIGURATIONMANAGERACTIONTYPE ConfigurationManagerActionType = iota
    REFRESHUSERPOLICY_CONFIGURATIONMANAGERACTIONTYPE
    WAKEUPCLIENT_CONFIGURATIONMANAGERACTIONTYPE
    APPEVALUATION_CONFIGURATIONMANAGERACTIONTYPE
    QUICKSCAN_CONFIGURATIONMANAGERACTIONTYPE
    FULLSCAN_CONFIGURATIONMANAGERACTIONTYPE
    WINDOWSDEFENDERUPDATESIGNATURES_CONFIGURATIONMANAGERACTIONTYPE
)

func (i ConfigurationManagerActionType) String() string {
    return []string{"REFRESHMACHINEPOLICY", "REFRESHUSERPOLICY", "WAKEUPCLIENT", "APPEVALUATION", "QUICKSCAN", "FULLSCAN", "WINDOWSDEFENDERUPDATESIGNATURES"}[i]
}
func ParseConfigurationManagerActionType(v string) (interface{}, error) {
    result := REFRESHMACHINEPOLICY_CONFIGURATIONMANAGERACTIONTYPE
    switch strings.ToUpper(v) {
        case "REFRESHMACHINEPOLICY":
            result = REFRESHMACHINEPOLICY_CONFIGURATIONMANAGERACTIONTYPE
        case "REFRESHUSERPOLICY":
            result = REFRESHUSERPOLICY_CONFIGURATIONMANAGERACTIONTYPE
        case "WAKEUPCLIENT":
            result = WAKEUPCLIENT_CONFIGURATIONMANAGERACTIONTYPE
        case "APPEVALUATION":
            result = APPEVALUATION_CONFIGURATIONMANAGERACTIONTYPE
        case "QUICKSCAN":
            result = QUICKSCAN_CONFIGURATIONMANAGERACTIONTYPE
        case "FULLSCAN":
            result = FULLSCAN_CONFIGURATIONMANAGERACTIONTYPE
        case "WINDOWSDEFENDERUPDATESIGNATURES":
            result = WINDOWSDEFENDERUPDATESIGNATURES_CONFIGURATIONMANAGERACTIONTYPE
        default:
            return 0, errors.New("Unknown ConfigurationManagerActionType value: " + v)
    }
    return &result, nil
}
func SerializeConfigurationManagerActionType(values []ConfigurationManagerActionType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
