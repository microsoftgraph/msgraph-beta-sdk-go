package security
type SubmissionResultCategory int

const (
    NOTJUNK_SUBMISSIONRESULTCATEGORY SubmissionResultCategory = iota
    SPAM_SUBMISSIONRESULTCATEGORY
    PHISHING_SUBMISSIONRESULTCATEGORY
    MALWARE_SUBMISSIONRESULTCATEGORY
    ALLOWEDBYPOLICY_SUBMISSIONRESULTCATEGORY
    BLOCKEDBYPOLICY_SUBMISSIONRESULTCATEGORY
    SPOOF_SUBMISSIONRESULTCATEGORY
    UNKNOWN_SUBMISSIONRESULTCATEGORY
    NORESULTAVAILABLE_SUBMISSIONRESULTCATEGORY
    UNKNOWNFUTUREVALUE_SUBMISSIONRESULTCATEGORY
    BEINGANALYZED_SUBMISSIONRESULTCATEGORY
    NOTSUBMITTEDTOMICROSOFT_SUBMISSIONRESULTCATEGORY
    PHISHINGSIMULATION_SUBMISSIONRESULTCATEGORY
    ALLOWEDDUETOORGANIZATIONOVERRIDE_SUBMISSIONRESULTCATEGORY
    BLOCKEDDUETOORGANIZATIONOVERRIDE_SUBMISSIONRESULTCATEGORY
    ALLOWEDDUETOUSEROVERRIDE_SUBMISSIONRESULTCATEGORY
    BLOCKEDDUETOUSEROVERRIDE_SUBMISSIONRESULTCATEGORY
    ITEMNOTFOUND_SUBMISSIONRESULTCATEGORY
    THREATSFOUND_SUBMISSIONRESULTCATEGORY
    NOTHREATSFOUND_SUBMISSIONRESULTCATEGORY
    DOMAINIMPERSONATION_SUBMISSIONRESULTCATEGORY
    USERIMPERSONATION_SUBMISSIONRESULTCATEGORY
    BRANDIMPERSONATION_SUBMISSIONRESULTCATEGORY
    AUTHENTICATIONFAILURE_SUBMISSIONRESULTCATEGORY
    SPOOFEDBLOCKED_SUBMISSIONRESULTCATEGORY
    SPOOFEDALLOWED_SUBMISSIONRESULTCATEGORY
    REASONLOSTINTRANSIT_SUBMISSIONRESULTCATEGORY
    BULK_SUBMISSIONRESULTCATEGORY
)

func (i SubmissionResultCategory) String() string {
    return []string{"notJunk", "spam", "phishing", "malware", "allowedByPolicy", "blockedByPolicy", "spoof", "unknown", "noResultAvailable", "unknownFutureValue", "beingAnalyzed", "notSubmittedToMicrosoft", "phishingSimulation", "allowedDueToOrganizationOverride", "blockedDueToOrganizationOverride", "allowedDueToUserOverride", "blockedDueToUserOverride", "itemNotfound", "threatsFound", "noThreatsFound", "domainImpersonation", "userImpersonation", "brandImpersonation", "authenticationFailure", "spoofedBlocked", "spoofedAllowed", "reasonLostInTransit", "bulk"}[i]
}
func ParseSubmissionResultCategory(v string) (any, error) {
    result := NOTJUNK_SUBMISSIONRESULTCATEGORY
    switch v {
        case "notJunk":
            result = NOTJUNK_SUBMISSIONRESULTCATEGORY
        case "spam":
            result = SPAM_SUBMISSIONRESULTCATEGORY
        case "phishing":
            result = PHISHING_SUBMISSIONRESULTCATEGORY
        case "malware":
            result = MALWARE_SUBMISSIONRESULTCATEGORY
        case "allowedByPolicy":
            result = ALLOWEDBYPOLICY_SUBMISSIONRESULTCATEGORY
        case "blockedByPolicy":
            result = BLOCKEDBYPOLICY_SUBMISSIONRESULTCATEGORY
        case "spoof":
            result = SPOOF_SUBMISSIONRESULTCATEGORY
        case "unknown":
            result = UNKNOWN_SUBMISSIONRESULTCATEGORY
        case "noResultAvailable":
            result = NORESULTAVAILABLE_SUBMISSIONRESULTCATEGORY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SUBMISSIONRESULTCATEGORY
        case "beingAnalyzed":
            result = BEINGANALYZED_SUBMISSIONRESULTCATEGORY
        case "notSubmittedToMicrosoft":
            result = NOTSUBMITTEDTOMICROSOFT_SUBMISSIONRESULTCATEGORY
        case "phishingSimulation":
            result = PHISHINGSIMULATION_SUBMISSIONRESULTCATEGORY
        case "allowedDueToOrganizationOverride":
            result = ALLOWEDDUETOORGANIZATIONOVERRIDE_SUBMISSIONRESULTCATEGORY
        case "blockedDueToOrganizationOverride":
            result = BLOCKEDDUETOORGANIZATIONOVERRIDE_SUBMISSIONRESULTCATEGORY
        case "allowedDueToUserOverride":
            result = ALLOWEDDUETOUSEROVERRIDE_SUBMISSIONRESULTCATEGORY
        case "blockedDueToUserOverride":
            result = BLOCKEDDUETOUSEROVERRIDE_SUBMISSIONRESULTCATEGORY
        case "itemNotfound":
            result = ITEMNOTFOUND_SUBMISSIONRESULTCATEGORY
        case "threatsFound":
            result = THREATSFOUND_SUBMISSIONRESULTCATEGORY
        case "noThreatsFound":
            result = NOTHREATSFOUND_SUBMISSIONRESULTCATEGORY
        case "domainImpersonation":
            result = DOMAINIMPERSONATION_SUBMISSIONRESULTCATEGORY
        case "userImpersonation":
            result = USERIMPERSONATION_SUBMISSIONRESULTCATEGORY
        case "brandImpersonation":
            result = BRANDIMPERSONATION_SUBMISSIONRESULTCATEGORY
        case "authenticationFailure":
            result = AUTHENTICATIONFAILURE_SUBMISSIONRESULTCATEGORY
        case "spoofedBlocked":
            result = SPOOFEDBLOCKED_SUBMISSIONRESULTCATEGORY
        case "spoofedAllowed":
            result = SPOOFEDALLOWED_SUBMISSIONRESULTCATEGORY
        case "reasonLostInTransit":
            result = REASONLOSTINTRANSIT_SUBMISSIONRESULTCATEGORY
        case "bulk":
            result = BULK_SUBMISSIONRESULTCATEGORY
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSubmissionResultCategory(values []SubmissionResultCategory) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SubmissionResultCategory) isMultiValue() bool {
    return false
}
