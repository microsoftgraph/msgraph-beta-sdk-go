package models
// Action type on Configuration Manager client
type ConfigurationManagerActionType int

const (
    // Refresh machine policy on Configuration Manager client
    REFRESHMACHINEPOLICY_CONFIGURATIONMANAGERACTIONTYPE ConfigurationManagerActionType = iota
    // Refresh user policy on Configuration Manager client
    REFRESHUSERPOLICY_CONFIGURATIONMANAGERACTIONTYPE
    // Wake up Configuration Manager client
    WAKEUPCLIENT_CONFIGURATIONMANAGERACTIONTYPE
    // Evaluation application policy on Configuration Manager client
    APPEVALUATION_CONFIGURATIONMANAGERACTIONTYPE
    // Evaluation application policy on Configuration Manager client
    QUICKSCAN_CONFIGURATIONMANAGERACTIONTYPE
    // Evaluation application policy on Configuration Manager client
    FULLSCAN_CONFIGURATIONMANAGERACTIONTYPE
    // Evaluation application policy on Configuration Manager client
    WINDOWSDEFENDERUPDATESIGNATURES_CONFIGURATIONMANAGERACTIONTYPE
)

func (i ConfigurationManagerActionType) String() string {
    return []string{"refreshMachinePolicy", "refreshUserPolicy", "wakeUpClient", "appEvaluation", "quickScan", "fullScan", "windowsDefenderUpdateSignatures"}[i]
}
func ParseConfigurationManagerActionType(v string) (any, error) {
    result := REFRESHMACHINEPOLICY_CONFIGURATIONMANAGERACTIONTYPE
    switch v {
        case "refreshMachinePolicy":
            result = REFRESHMACHINEPOLICY_CONFIGURATIONMANAGERACTIONTYPE
        case "refreshUserPolicy":
            result = REFRESHUSERPOLICY_CONFIGURATIONMANAGERACTIONTYPE
        case "wakeUpClient":
            result = WAKEUPCLIENT_CONFIGURATIONMANAGERACTIONTYPE
        case "appEvaluation":
            result = APPEVALUATION_CONFIGURATIONMANAGERACTIONTYPE
        case "quickScan":
            result = QUICKSCAN_CONFIGURATIONMANAGERACTIONTYPE
        case "fullScan":
            result = FULLSCAN_CONFIGURATIONMANAGERACTIONTYPE
        case "windowsDefenderUpdateSignatures":
            result = WINDOWSDEFENDERUPDATESIGNATURES_CONFIGURATIONMANAGERACTIONTYPE
        default:
            return nil, nil
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
func (i ConfigurationManagerActionType) isMultiValue() bool {
    return false
}
