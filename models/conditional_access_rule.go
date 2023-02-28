package models
import (
    "errors"
)
// 
type ConditionalAccessRule int

const (
    ALLAPPS_CONDITIONALACCESSRULE ConditionalAccessRule = iota
    FIRSTPARTYAPPS_CONDITIONALACCESSRULE
    OFFICE365_CONDITIONALACCESSRULE
    APPID_CONDITIONALACCESSRULE
    ACR_CONDITIONALACCESSRULE
    APPFILTER_CONDITIONALACCESSRULE
    ALLUSERS_CONDITIONALACCESSRULE
    GUEST_CONDITIONALACCESSRULE
    GROUPID_CONDITIONALACCESSRULE
    ROLEID_CONDITIONALACCESSRULE
    USERID_CONDITIONALACCESSRULE
    ALLDEVICEPLATFORMS_CONDITIONALACCESSRULE
    DEVICEPLATFORM_CONDITIONALACCESSRULE
    ALLLOCATIONS_CONDITIONALACCESSRULE
    INSIDECORPNET_CONDITIONALACCESSRULE
    ALLTRUSTEDLOCATIONS_CONDITIONALACCESSRULE
    LOCATIONID_CONDITIONALACCESSRULE
    ALLDEVICES_CONDITIONALACCESSRULE
    DEVICEFILTER_CONDITIONALACCESSRULE
    DEVICESTATE_CONDITIONALACCESSRULE
    UNKNOWNFUTUREVALUE_CONDITIONALACCESSRULE
    DEVICEFILTERINCLUDERULENOTMATCHED_CONDITIONALACCESSRULE
    ALLDEVICESTATES_CONDITIONALACCESSRULE
    ANONYMIZEDIPADDRESS_CONDITIONALACCESSRULE
    UNFAMILIARFEATURES_CONDITIONALACCESSRULE
    NATIONSTATEIPADDRESS_CONDITIONALACCESSRULE
    REALTIMETHREATINTELLIGENCE_CONDITIONALACCESSRULE
    INTERNALGUEST_CONDITIONALACCESSRULE
    B2BCOLLABORATIONGUEST_CONDITIONALACCESSRULE
    B2BCOLLABORATIONMEMBER_CONDITIONALACCESSRULE
    B2BDIRECTCONNECTUSER_CONDITIONALACCESSRULE
    OTHEREXTERNALUSER_CONDITIONALACCESSRULE
    SERVICEPROVIDER_CONDITIONALACCESSRULE
)

func (i ConditionalAccessRule) String() string {
    return []string{"allApps", "firstPartyApps", "office365", "appId", "acr", "appFilter", "allUsers", "guest", "groupId", "roleId", "userId", "allDevicePlatforms", "devicePlatform", "allLocations", "insideCorpnet", "allTrustedLocations", "locationId", "allDevices", "deviceFilter", "deviceState", "unknownFutureValue", "deviceFilterIncludeRuleNotMatched", "allDeviceStates", "anonymizedIPAddress", "unfamiliarFeatures", "nationStateIPAddress", "realTimeThreatIntelligence", "internalGuest", "b2bCollaborationGuest", "b2bCollaborationMember", "b2bDirectConnectUser", "otherExternalUser", "serviceProvider"}[i]
}
func ParseConditionalAccessRule(v string) (any, error) {
    result := ALLAPPS_CONDITIONALACCESSRULE
    switch v {
        case "allApps":
            result = ALLAPPS_CONDITIONALACCESSRULE
        case "firstPartyApps":
            result = FIRSTPARTYAPPS_CONDITIONALACCESSRULE
        case "office365":
            result = OFFICE365_CONDITIONALACCESSRULE
        case "appId":
            result = APPID_CONDITIONALACCESSRULE
        case "acr":
            result = ACR_CONDITIONALACCESSRULE
        case "appFilter":
            result = APPFILTER_CONDITIONALACCESSRULE
        case "allUsers":
            result = ALLUSERS_CONDITIONALACCESSRULE
        case "guest":
            result = GUEST_CONDITIONALACCESSRULE
        case "groupId":
            result = GROUPID_CONDITIONALACCESSRULE
        case "roleId":
            result = ROLEID_CONDITIONALACCESSRULE
        case "userId":
            result = USERID_CONDITIONALACCESSRULE
        case "allDevicePlatforms":
            result = ALLDEVICEPLATFORMS_CONDITIONALACCESSRULE
        case "devicePlatform":
            result = DEVICEPLATFORM_CONDITIONALACCESSRULE
        case "allLocations":
            result = ALLLOCATIONS_CONDITIONALACCESSRULE
        case "insideCorpnet":
            result = INSIDECORPNET_CONDITIONALACCESSRULE
        case "allTrustedLocations":
            result = ALLTRUSTEDLOCATIONS_CONDITIONALACCESSRULE
        case "locationId":
            result = LOCATIONID_CONDITIONALACCESSRULE
        case "allDevices":
            result = ALLDEVICES_CONDITIONALACCESSRULE
        case "deviceFilter":
            result = DEVICEFILTER_CONDITIONALACCESSRULE
        case "deviceState":
            result = DEVICESTATE_CONDITIONALACCESSRULE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONDITIONALACCESSRULE
        case "deviceFilterIncludeRuleNotMatched":
            result = DEVICEFILTERINCLUDERULENOTMATCHED_CONDITIONALACCESSRULE
        case "allDeviceStates":
            result = ALLDEVICESTATES_CONDITIONALACCESSRULE
        case "anonymizedIPAddress":
            result = ANONYMIZEDIPADDRESS_CONDITIONALACCESSRULE
        case "unfamiliarFeatures":
            result = UNFAMILIARFEATURES_CONDITIONALACCESSRULE
        case "nationStateIPAddress":
            result = NATIONSTATEIPADDRESS_CONDITIONALACCESSRULE
        case "realTimeThreatIntelligence":
            result = REALTIMETHREATINTELLIGENCE_CONDITIONALACCESSRULE
        case "internalGuest":
            result = INTERNALGUEST_CONDITIONALACCESSRULE
        case "b2bCollaborationGuest":
            result = B2BCOLLABORATIONGUEST_CONDITIONALACCESSRULE
        case "b2bCollaborationMember":
            result = B2BCOLLABORATIONMEMBER_CONDITIONALACCESSRULE
        case "b2bDirectConnectUser":
            result = B2BDIRECTCONNECTUSER_CONDITIONALACCESSRULE
        case "otherExternalUser":
            result = OTHEREXTERNALUSER_CONDITIONALACCESSRULE
        case "serviceProvider":
            result = SERVICEPROVIDER_CONDITIONALACCESSRULE
        default:
            return 0, errors.New("Unknown ConditionalAccessRule value: " + v)
    }
    return &result, nil
}
func SerializeConditionalAccessRule(values []ConditionalAccessRule) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
