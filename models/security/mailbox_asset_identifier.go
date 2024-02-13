package security
import (
    "errors"
)
type MailboxAssetIdentifier int

const (
    ACCOUNTUPN_MAILBOXASSETIDENTIFIER MailboxAssetIdentifier = iota
    FILEOWNERUPN_MAILBOXASSETIDENTIFIER
    INITIATINGPROCESSACCOUNTUPN_MAILBOXASSETIDENTIFIER
    LASTMODIFYINGACCOUNTUPN_MAILBOXASSETIDENTIFIER
    TARGETACCOUNTUPN_MAILBOXASSETIDENTIFIER
    SENDERFROMADDRESS_MAILBOXASSETIDENTIFIER
    SENDERDISPLAYNAME_MAILBOXASSETIDENTIFIER
    RECIPIENTEMAILADDRESS_MAILBOXASSETIDENTIFIER
    SENDERMAILFROMADDRESS_MAILBOXASSETIDENTIFIER
    UNKNOWNFUTUREVALUE_MAILBOXASSETIDENTIFIER
)

func (i MailboxAssetIdentifier) String() string {
    return []string{"accountUpn", "fileOwnerUpn", "initiatingProcessAccountUpn", "lastModifyingAccountUpn", "targetAccountUpn", "senderFromAddress", "senderDisplayName", "recipientEmailAddress", "senderMailFromAddress", "unknownFutureValue"}[i]
}
func ParseMailboxAssetIdentifier(v string) (any, error) {
    result := ACCOUNTUPN_MAILBOXASSETIDENTIFIER
    switch v {
        case "accountUpn":
            result = ACCOUNTUPN_MAILBOXASSETIDENTIFIER
        case "fileOwnerUpn":
            result = FILEOWNERUPN_MAILBOXASSETIDENTIFIER
        case "initiatingProcessAccountUpn":
            result = INITIATINGPROCESSACCOUNTUPN_MAILBOXASSETIDENTIFIER
        case "lastModifyingAccountUpn":
            result = LASTMODIFYINGACCOUNTUPN_MAILBOXASSETIDENTIFIER
        case "targetAccountUpn":
            result = TARGETACCOUNTUPN_MAILBOXASSETIDENTIFIER
        case "senderFromAddress":
            result = SENDERFROMADDRESS_MAILBOXASSETIDENTIFIER
        case "senderDisplayName":
            result = SENDERDISPLAYNAME_MAILBOXASSETIDENTIFIER
        case "recipientEmailAddress":
            result = RECIPIENTEMAILADDRESS_MAILBOXASSETIDENTIFIER
        case "senderMailFromAddress":
            result = SENDERMAILFROMADDRESS_MAILBOXASSETIDENTIFIER
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MAILBOXASSETIDENTIFIER
        default:
            return 0, errors.New("Unknown MailboxAssetIdentifier value: " + v)
    }
    return &result, nil
}
func SerializeMailboxAssetIdentifier(values []MailboxAssetIdentifier) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MailboxAssetIdentifier) isMultiValue() bool {
    return false
}
