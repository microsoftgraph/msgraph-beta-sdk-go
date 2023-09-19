package security
import (
    "errors"
)
// 
type TenantAllowBlockListAction int

const (
    ALLOW_TENANTALLOWBLOCKLISTACTION TenantAllowBlockListAction = iota
    BLOCK_TENANTALLOWBLOCKLISTACTION
    UNKNOWNFUTUREVALUE_TENANTALLOWBLOCKLISTACTION
)

func (i TenantAllowBlockListAction) String() string {
    return []string{"allow", "block", "unknownFutureValue"}[i]
}
func ParseTenantAllowBlockListAction(v string) (any, error) {
    result := ALLOW_TENANTALLOWBLOCKLISTACTION
    switch v {
        case "allow":
            result = ALLOW_TENANTALLOWBLOCKLISTACTION
        case "block":
            result = BLOCK_TENANTALLOWBLOCKLISTACTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_TENANTALLOWBLOCKLISTACTION
        default:
            return 0, errors.New("Unknown TenantAllowBlockListAction value: " + v)
    }
    return &result, nil
}
func SerializeTenantAllowBlockListAction(values []TenantAllowBlockListAction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i TenantAllowBlockListAction) isMultiValue() bool {
    return false
}
