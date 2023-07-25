package security
import (
    "errors"
)
// 
type UserAssetIdentifier int

const (
    ACCOUNTOBJECTID_USERASSETIDENTIFIER UserAssetIdentifier = iota
    ACCOUNTSID_USERASSETIDENTIFIER
    ACCOUNTUPN_USERASSETIDENTIFIER
    ACCOUNTNAME_USERASSETIDENTIFIER
    ACCOUNTDOMAIN_USERASSETIDENTIFIER
    ACCOUNTID_USERASSETIDENTIFIER
    REQUESTACCOUNTSID_USERASSETIDENTIFIER
    REQUESTACCOUNTNAME_USERASSETIDENTIFIER
    REQUESTACCOUNTDOMAIN_USERASSETIDENTIFIER
    RECIPIENTOBJECTID_USERASSETIDENTIFIER
    PROCESSACCOUNTOBJECTID_USERASSETIDENTIFIER
    INITIATINGACCOUNTSID_USERASSETIDENTIFIER
    INITIATINGPROCESSACCOUNTUPN_USERASSETIDENTIFIER
    INITIATINGACCOUNTNAME_USERASSETIDENTIFIER
    INITIATINGACCOUNTDOMAIN_USERASSETIDENTIFIER
    SERVICEPRINCIPALID_USERASSETIDENTIFIER
    SERVICEPRINCIPALNAME_USERASSETIDENTIFIER
    TARGETACCOUNTUPN_USERASSETIDENTIFIER
    UNKNOWNFUTUREVALUE_USERASSETIDENTIFIER
)

func (i UserAssetIdentifier) String() string {
    return []string{"accountObjectId", "accountSid", "accountUpn", "accountName", "accountDomain", "accountId", "requestAccountSid", "requestAccountName", "requestAccountDomain", "recipientObjectId", "processAccountObjectId", "initiatingAccountSid", "initiatingProcessAccountUpn", "initiatingAccountName", "initiatingAccountDomain", "servicePrincipalId", "servicePrincipalName", "targetAccountUpn", "unknownFutureValue"}[i]
}
func ParseUserAssetIdentifier(v string) (any, error) {
    result := ACCOUNTOBJECTID_USERASSETIDENTIFIER
    switch v {
        case "accountObjectId":
            result = ACCOUNTOBJECTID_USERASSETIDENTIFIER
        case "accountSid":
            result = ACCOUNTSID_USERASSETIDENTIFIER
        case "accountUpn":
            result = ACCOUNTUPN_USERASSETIDENTIFIER
        case "accountName":
            result = ACCOUNTNAME_USERASSETIDENTIFIER
        case "accountDomain":
            result = ACCOUNTDOMAIN_USERASSETIDENTIFIER
        case "accountId":
            result = ACCOUNTID_USERASSETIDENTIFIER
        case "requestAccountSid":
            result = REQUESTACCOUNTSID_USERASSETIDENTIFIER
        case "requestAccountName":
            result = REQUESTACCOUNTNAME_USERASSETIDENTIFIER
        case "requestAccountDomain":
            result = REQUESTACCOUNTDOMAIN_USERASSETIDENTIFIER
        case "recipientObjectId":
            result = RECIPIENTOBJECTID_USERASSETIDENTIFIER
        case "processAccountObjectId":
            result = PROCESSACCOUNTOBJECTID_USERASSETIDENTIFIER
        case "initiatingAccountSid":
            result = INITIATINGACCOUNTSID_USERASSETIDENTIFIER
        case "initiatingProcessAccountUpn":
            result = INITIATINGPROCESSACCOUNTUPN_USERASSETIDENTIFIER
        case "initiatingAccountName":
            result = INITIATINGACCOUNTNAME_USERASSETIDENTIFIER
        case "initiatingAccountDomain":
            result = INITIATINGACCOUNTDOMAIN_USERASSETIDENTIFIER
        case "servicePrincipalId":
            result = SERVICEPRINCIPALID_USERASSETIDENTIFIER
        case "servicePrincipalName":
            result = SERVICEPRINCIPALNAME_USERASSETIDENTIFIER
        case "targetAccountUpn":
            result = TARGETACCOUNTUPN_USERASSETIDENTIFIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_USERASSETIDENTIFIER
        default:
            return 0, errors.New("Unknown UserAssetIdentifier value: " + v)
    }
    return &result, nil
}
func SerializeUserAssetIdentifier(values []UserAssetIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
