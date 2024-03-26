package models
import (
    "errors"
)
type ConditionalAccessWhatIfReasons int

const (
    NOTSET_CONDITIONALACCESSWHATIFREASONS ConditionalAccessWhatIfReasons = iota
    NOTENOUGHINFORMATION_CONDITIONALACCESSWHATIFREASONS
    INVALIDCONDITION_CONDITIONALACCESSWHATIFREASONS
    USERS_CONDITIONALACCESSWHATIFREASONS
    WORKLOADIDENTITIES_CONDITIONALACCESSWHATIFREASONS
    APPLICATION_CONDITIONALACCESSWHATIFREASONS
    USERACTIONS_CONDITIONALACCESSWHATIFREASONS
    AUTHENTICATIONCONTEXT_CONDITIONALACCESSWHATIFREASONS
    DEVICEPLATFORM_CONDITIONALACCESSWHATIFREASONS
    DEVICES_CONDITIONALACCESSWHATIFREASONS
    CLIENTAPPS_CONDITIONALACCESSWHATIFREASONS
    LOCATION_CONDITIONALACCESSWHATIFREASONS
    SIGNINRISK_CONDITIONALACCESSWHATIFREASONS
    EMPTYPOLICY_CONDITIONALACCESSWHATIFREASONS
    INVALIDPOLICY_CONDITIONALACCESSWHATIFREASONS
    POLICYNOTENABLED_CONDITIONALACCESSWHATIFREASONS
    USERRISK_CONDITIONALACCESSWHATIFREASONS
    TIME_CONDITIONALACCESSWHATIFREASONS
    INSIDERRISK_CONDITIONALACCESSWHATIFREASONS
    AUTHENTICATIONFLOW_CONDITIONALACCESSWHATIFREASONS
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSWHATIFREASONS
)

func (i ConditionalAccessWhatIfReasons) String() string {
    return []string{"notSet", "notEnoughInformation", "invalidCondition", "users", "workloadIdentities", "application", "userActions", "authenticationContext", "devicePlatform", "devices", "clientApps", "location", "signInRisk", "emptyPolicy", "invalidPolicy", "policyNotEnabled", "userRisk", "time", "insiderRisk", "authenticationFlow", "unknownFutureValue"}[i]
}
func ParseConditionalAccessWhatIfReasons(v string) (any, error) {
    result := NOTSET_CONDITIONALACCESSWHATIFREASONS
    switch v {
        case "notSet":
            result = NOTSET_CONDITIONALACCESSWHATIFREASONS
        case "notEnoughInformation":
            result = NOTENOUGHINFORMATION_CONDITIONALACCESSWHATIFREASONS
        case "invalidCondition":
            result = INVALIDCONDITION_CONDITIONALACCESSWHATIFREASONS
        case "users":
            result = USERS_CONDITIONALACCESSWHATIFREASONS
        case "workloadIdentities":
            result = WORKLOADIDENTITIES_CONDITIONALACCESSWHATIFREASONS
        case "application":
            result = APPLICATION_CONDITIONALACCESSWHATIFREASONS
        case "userActions":
            result = USERACTIONS_CONDITIONALACCESSWHATIFREASONS
        case "authenticationContext":
            result = AUTHENTICATIONCONTEXT_CONDITIONALACCESSWHATIFREASONS
        case "devicePlatform":
            result = DEVICEPLATFORM_CONDITIONALACCESSWHATIFREASONS
        case "devices":
            result = DEVICES_CONDITIONALACCESSWHATIFREASONS
        case "clientApps":
            result = CLIENTAPPS_CONDITIONALACCESSWHATIFREASONS
        case "location":
            result = LOCATION_CONDITIONALACCESSWHATIFREASONS
        case "signInRisk":
            result = SIGNINRISK_CONDITIONALACCESSWHATIFREASONS
        case "emptyPolicy":
            result = EMPTYPOLICY_CONDITIONALACCESSWHATIFREASONS
        case "invalidPolicy":
            result = INVALIDPOLICY_CONDITIONALACCESSWHATIFREASONS
        case "policyNotEnabled":
            result = POLICYNOTENABLED_CONDITIONALACCESSWHATIFREASONS
        case "userRisk":
            result = USERRISK_CONDITIONALACCESSWHATIFREASONS
        case "time":
            result = TIME_CONDITIONALACCESSWHATIFREASONS
        case "insiderRisk":
            result = INSIDERRISK_CONDITIONALACCESSWHATIFREASONS
        case "authenticationFlow":
            result = AUTHENTICATIONFLOW_CONDITIONALACCESSWHATIFREASONS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONDITIONALACCESSWHATIFREASONS
        default:
            return 0, errors.New("Unknown ConditionalAccessWhatIfReasons value: " + v)
    }
    return &result, nil
}
func SerializeConditionalAccessWhatIfReasons(values []ConditionalAccessWhatIfReasons) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConditionalAccessWhatIfReasons) isMultiValue() bool {
    return false
}
