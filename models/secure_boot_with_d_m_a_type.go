package models
// Possible values of Secure Boot with DMA
type SecureBootWithDMAType int

const (
    // Not configured, no operation
    NOTCONFIGURED_SECUREBOOTWITHDMATYPE SecureBootWithDMAType = iota
    // Turns on VBS with Secure Boot
    WITHOUTDMA_SECUREBOOTWITHDMATYPE
    // Turns on VBS with Secure Boot and DMA
    WITHDMA_SECUREBOOTWITHDMATYPE
)

func (i SecureBootWithDMAType) String() string {
    return []string{"notConfigured", "withoutDMA", "withDMA"}[i]
}
func ParseSecureBootWithDMAType(v string) (any, error) {
    result := NOTCONFIGURED_SECUREBOOTWITHDMATYPE
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_SECUREBOOTWITHDMATYPE
        case "withoutDMA":
            result = WITHOUTDMA_SECUREBOOTWITHDMATYPE
        case "withDMA":
            result = WITHDMA_SECUREBOOTWITHDMATYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSecureBootWithDMAType(values []SecureBootWithDMAType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SecureBootWithDMAType) isMultiValue() bool {
    return false
}
