package security
import (
    "errors"
    "math"
    "strings"
)
// 
type UserMailboxSetting int

const (
    NONE_USERMAILBOXSETTING = 1
    JUNKMAILDELETION_USERMAILBOXSETTING = 2
    ISFROMADDRESSINADDRESSBOOK_USERMAILBOXSETTING = 4
    ISFROMADDRESSINADDRESSSAFELIST_USERMAILBOXSETTING = 8
    ISFROMADDRESSINADDRESSBLOCKLIST_USERMAILBOXSETTING = 16
    ISFROMADDRESSINADDRESSIMPLICITSAFELIST_USERMAILBOXSETTING = 32
    ISFROMADDRESSINADDRESSIMPLICITJUNKLIST_USERMAILBOXSETTING = 64
    ISFROMDOMAININDOMAINSAFELIST_USERMAILBOXSETTING = 128
    ISFROMDOMAININDOMAINBLOCKLIST_USERMAILBOXSETTING = 256
    ISRECIPIENTINRECIPIENTSAFELIST_USERMAILBOXSETTING = 512
    CUSTOMRULE_USERMAILBOXSETTING = 1024
    JUNKMAILRULE_USERMAILBOXSETTING = 2048
    SENDERPRAPRESENT_USERMAILBOXSETTING = 4096
    FROMFIRSTTIMESENDER_USERMAILBOXSETTING = 8192
    EXCLUSIVE_USERMAILBOXSETTING = 16384
    PRIORSEENPASS_USERMAILBOXSETTING = 32768
    SENDERAUTHENTICATIONSUCCEEDED_USERMAILBOXSETTING = 65536
    ISJUNKMAILRULEENABLED_USERMAILBOXSETTING = 131072
    UNKNOWNFUTUREVALUE_USERMAILBOXSETTING = 262144
)

func (i UserMailboxSetting) String() string {
    var values []string
    options := []string{"none", "junkMailDeletion", "isFromAddressInAddressBook", "isFromAddressInAddressSafeList", "isFromAddressInAddressBlockList", "isFromAddressInAddressImplicitSafeList", "isFromAddressInAddressImplicitJunkList", "isFromDomainInDomainSafeList", "isFromDomainInDomainBlockList", "isRecipientInRecipientSafeList", "customRule", "junkMailRule", "senderPraPresent", "fromFirstTimeSender", "exclusive", "priorSeenPass", "senderAuthenticationSucceeded", "isJunkMailRuleEnabled", "unknownFutureValue"}
    for p := 0; p < 19; p++ {
        mantis := UserMailboxSetting(int(math.Pow(2, float64(p))))
        if i&mantis == mantis {
            values = append(values, options[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseUserMailboxSetting(v string) (any, error) {
    var result UserMailboxSetting
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_USERMAILBOXSETTING
            case "junkMailDeletion":
                result |= JUNKMAILDELETION_USERMAILBOXSETTING
            case "isFromAddressInAddressBook":
                result |= ISFROMADDRESSINADDRESSBOOK_USERMAILBOXSETTING
            case "isFromAddressInAddressSafeList":
                result |= ISFROMADDRESSINADDRESSSAFELIST_USERMAILBOXSETTING
            case "isFromAddressInAddressBlockList":
                result |= ISFROMADDRESSINADDRESSBLOCKLIST_USERMAILBOXSETTING
            case "isFromAddressInAddressImplicitSafeList":
                result |= ISFROMADDRESSINADDRESSIMPLICITSAFELIST_USERMAILBOXSETTING
            case "isFromAddressInAddressImplicitJunkList":
                result |= ISFROMADDRESSINADDRESSIMPLICITJUNKLIST_USERMAILBOXSETTING
            case "isFromDomainInDomainSafeList":
                result |= ISFROMDOMAININDOMAINSAFELIST_USERMAILBOXSETTING
            case "isFromDomainInDomainBlockList":
                result |= ISFROMDOMAININDOMAINBLOCKLIST_USERMAILBOXSETTING
            case "isRecipientInRecipientSafeList":
                result |= ISRECIPIENTINRECIPIENTSAFELIST_USERMAILBOXSETTING
            case "customRule":
                result |= CUSTOMRULE_USERMAILBOXSETTING
            case "junkMailRule":
                result |= JUNKMAILRULE_USERMAILBOXSETTING
            case "senderPraPresent":
                result |= SENDERPRAPRESENT_USERMAILBOXSETTING
            case "fromFirstTimeSender":
                result |= FROMFIRSTTIMESENDER_USERMAILBOXSETTING
            case "exclusive":
                result |= EXCLUSIVE_USERMAILBOXSETTING
            case "priorSeenPass":
                result |= PRIORSEENPASS_USERMAILBOXSETTING
            case "senderAuthenticationSucceeded":
                result |= SENDERAUTHENTICATIONSUCCEEDED_USERMAILBOXSETTING
            case "isJunkMailRuleEnabled":
                result |= ISJUNKMAILRULEENABLED_USERMAILBOXSETTING
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_USERMAILBOXSETTING
            default:
                return 0, errors.New("Unknown UserMailboxSetting value: " + v)
        }
    }
    return &result, nil
}
func SerializeUserMailboxSetting(values []UserMailboxSetting) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i UserMailboxSetting) isMultiValue() bool {
    return true
}
