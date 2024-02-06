package models
import (
    "errors"
    "math"
    "strings"
)
// 
type ConditionalAccessConditions int

const (
    NONE_CONDITIONALACCESSCONDITIONS = 1
    APPLICATION_CONDITIONALACCESSCONDITIONS = 2
    USERS_CONDITIONALACCESSCONDITIONS = 4
    DEVICEPLATFORM_CONDITIONALACCESSCONDITIONS = 8
    LOCATION_CONDITIONALACCESSCONDITIONS = 16
    CLIENTTYPE_CONDITIONALACCESSCONDITIONS = 32
    SIGNINRISK_CONDITIONALACCESSCONDITIONS = 64
    USERRISK_CONDITIONALACCESSCONDITIONS = 128
    TIME_CONDITIONALACCESSCONDITIONS = 256
    DEVICESTATE_CONDITIONALACCESSCONDITIONS = 512
    CLIENT_CONDITIONALACCESSCONDITIONS = 1024
    IPADDRESSSEENBYAZUREAD_CONDITIONALACCESSCONDITIONS = 2048
    IPADDRESSSEENBYRESOURCEPROVIDER_CONDITIONALACCESSCONDITIONS = 4096
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSCONDITIONS = 8192
    SERVICEPRINCIPALS_CONDITIONALACCESSCONDITIONS = 16384
    SERVICEPRINCIPALRISK_CONDITIONALACCESSCONDITIONS = 32768
    AUTHENTICATIONFLOWS_CONDITIONALACCESSCONDITIONS = 65536
    INSIDERRISK_CONDITIONALACCESSCONDITIONS = 131072
)

func (i ConditionalAccessConditions) String() string {
    var values []string
    options := []string{"none", "application", "users", "devicePlatform", "location", "clientType", "signInRisk", "userRisk", "time", "deviceState", "client", "ipAddressSeenByAzureAD", "ipAddressSeenByResourceProvider", "unknownFutureValue", "servicePrincipals", "servicePrincipalRisk", "authenticationFlows", "insiderRisk"}
    for p := 0; p < 18; p++ {
        mantis := ConditionalAccessConditions(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseConditionalAccessConditions(v string) (any, error) {
    var result ConditionalAccessConditions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_CONDITIONALACCESSCONDITIONS
            case "application":
                result |= APPLICATION_CONDITIONALACCESSCONDITIONS
            case "users":
                result |= USERS_CONDITIONALACCESSCONDITIONS
            case "devicePlatform":
                result |= DEVICEPLATFORM_CONDITIONALACCESSCONDITIONS
            case "location":
                result |= LOCATION_CONDITIONALACCESSCONDITIONS
            case "clientType":
                result |= CLIENTTYPE_CONDITIONALACCESSCONDITIONS
            case "signInRisk":
                result |= SIGNINRISK_CONDITIONALACCESSCONDITIONS
            case "userRisk":
                result |= USERRISK_CONDITIONALACCESSCONDITIONS
            case "time":
                result |= TIME_CONDITIONALACCESSCONDITIONS
            case "deviceState":
                result |= DEVICESTATE_CONDITIONALACCESSCONDITIONS
            case "client":
                result |= CLIENT_CONDITIONALACCESSCONDITIONS
            case "ipAddressSeenByAzureAD":
                result |= IPADDRESSSEENBYAZUREAD_CONDITIONALACCESSCONDITIONS
            case "ipAddressSeenByResourceProvider":
                result |= IPADDRESSSEENBYRESOURCEPROVIDER_CONDITIONALACCESSCONDITIONS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_CONDITIONALACCESSCONDITIONS
            case "servicePrincipals":
                result |= SERVICEPRINCIPALS_CONDITIONALACCESSCONDITIONS
            case "servicePrincipalRisk":
                result |= SERVICEPRINCIPALRISK_CONDITIONALACCESSCONDITIONS
            case "authenticationFlows":
                result |= AUTHENTICATIONFLOWS_CONDITIONALACCESSCONDITIONS
            case "insiderRisk":
                result |= INSIDERRISK_CONDITIONALACCESSCONDITIONS
            default:
                return 0, errors.New("Unknown ConditionalAccessConditions value: " + v)
        }
    }
    return &result, nil
}
func SerializeConditionalAccessConditions(values []ConditionalAccessConditions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ConditionalAccessConditions) isMultiValue() bool {
    return true
}
