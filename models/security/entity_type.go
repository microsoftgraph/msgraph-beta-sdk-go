package security
type EntityType int

const (
    USERNAME_ENTITYTYPE EntityType = iota
    IPADDRESS_ENTITYTYPE
    MACHINENAME_ENTITYTYPE
    OTHER_ENTITYTYPE
    UNKNOWN_ENTITYTYPE
    UNKNOWNFUTUREVALUE_ENTITYTYPE
)

func (i EntityType) String() string {
    return []string{"userName", "ipAddress", "machineName", "other", "unknown", "unknownFutureValue"}[i]
}
func ParseEntityType(v string) (any, error) {
    result := USERNAME_ENTITYTYPE
    switch v {
        case "userName":
            result = USERNAME_ENTITYTYPE
        case "ipAddress":
            result = IPADDRESS_ENTITYTYPE
        case "machineName":
            result = MACHINENAME_ENTITYTYPE
        case "other":
            result = OTHER_ENTITYTYPE
        case "unknown":
            result = UNKNOWN_ENTITYTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ENTITYTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeEntityType(values []EntityType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i EntityType) isMultiValue() bool {
    return false
}
