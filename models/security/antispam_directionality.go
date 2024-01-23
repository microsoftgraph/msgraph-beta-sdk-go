package security
import (
    "errors"
)
// 
type AntispamDirectionality int

const (
    UNKNOWN_ANTISPAMDIRECTIONALITY AntispamDirectionality = iota
    INBOUND_ANTISPAMDIRECTIONALITY
    OUTBOUND_ANTISPAMDIRECTIONALITY
    INTRAORG_ANTISPAMDIRECTIONALITY
    UNKNOWNFUTUREVALUE_ANTISPAMDIRECTIONALITY
)

func (i AntispamDirectionality) String() string {
    return []string{"unknown", "inbound", "outbound", "intraOrg", "unknownFutureValue"}[i]
}
func ParseAntispamDirectionality(v string) (any, error) {
    result := UNKNOWN_ANTISPAMDIRECTIONALITY
    switch v {
        case "unknown":
            result = UNKNOWN_ANTISPAMDIRECTIONALITY
        case "inbound":
            result = INBOUND_ANTISPAMDIRECTIONALITY
        case "outbound":
            result = OUTBOUND_ANTISPAMDIRECTIONALITY
        case "intraOrg":
            result = INTRAORG_ANTISPAMDIRECTIONALITY
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ANTISPAMDIRECTIONALITY
        default:
            return 0, errors.New("Unknown AntispamDirectionality value: " + v)
    }
    return &result, nil
}
func SerializeAntispamDirectionality(values []AntispamDirectionality) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AntispamDirectionality) isMultiValue() bool {
    return false
}
