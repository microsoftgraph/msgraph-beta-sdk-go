package ediscovery
type LegalHoldStatus int

const (
    PENDING_LEGALHOLDSTATUS LegalHoldStatus = iota
    ERROR_LEGALHOLDSTATUS
    SUCCESS_LEGALHOLDSTATUS
    UNKNOWNFUTUREVALUE_LEGALHOLDSTATUS
)

func (i LegalHoldStatus) String() string {
    return []string{"Pending", "Error", "Success", "UnknownFutureValue"}[i]
}
func ParseLegalHoldStatus(v string) (any, error) {
    result := PENDING_LEGALHOLDSTATUS
    switch v {
        case "Pending":
            result = PENDING_LEGALHOLDSTATUS
        case "Error":
            result = ERROR_LEGALHOLDSTATUS
        case "Success":
            result = SUCCESS_LEGALHOLDSTATUS
        case "UnknownFutureValue":
            result = UNKNOWNFUTUREVALUE_LEGALHOLDSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeLegalHoldStatus(values []LegalHoldStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i LegalHoldStatus) isMultiValue() bool {
    return false
}
