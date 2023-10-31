package security
import (
    "errors"
    "strings"
)
// 
type AdditionalOptions int

const (
    NONE_ADDITIONALOPTIONS AdditionalOptions = iota
    TEAMSANDYAMMERCONVERSATIONS_ADDITIONALOPTIONS
    CLOUDATTACHMENTS_ADDITIONALOPTIONS
    ALLDOCUMENTVERSIONS_ADDITIONALOPTIONS
    SUBFOLDERCONTENTS_ADDITIONALOPTIONS
    LISTATTACHMENTS_ADDITIONALOPTIONS
    UNKNOWNFUTUREVALUE_ADDITIONALOPTIONS
)

func (i AdditionalOptions) String() string {
    var values []string
    for p := AdditionalOptions(1); p <= UNKNOWNFUTUREVALUE_ADDITIONALOPTIONS; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"none", "teamsAndYammerConversations", "cloudAttachments", "allDocumentVersions", "subfolderContents", "listAttachments", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAdditionalOptions(v string) (any, error) {
    var result AdditionalOptions
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "none":
                result |= NONE_ADDITIONALOPTIONS
            case "teamsAndYammerConversations":
                result |= TEAMSANDYAMMERCONVERSATIONS_ADDITIONALOPTIONS
            case "cloudAttachments":
                result |= CLOUDATTACHMENTS_ADDITIONALOPTIONS
            case "allDocumentVersions":
                result |= ALLDOCUMENTVERSIONS_ADDITIONALOPTIONS
            case "subfolderContents":
                result |= SUBFOLDERCONTENTS_ADDITIONALOPTIONS
            case "listAttachments":
                result |= LISTATTACHMENTS_ADDITIONALOPTIONS
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_ADDITIONALOPTIONS
            default:
                return 0, errors.New("Unknown AdditionalOptions value: " + v)
        }
    }
    return &result, nil
}
func SerializeAdditionalOptions(values []AdditionalOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AdditionalOptions) isMultiValue() bool {
    return true
}
