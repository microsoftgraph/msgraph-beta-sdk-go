package models
// Extensible Authentication Protocol (EAP) Configuration Types.
type AndroidEapType int

const (
    // Extensible Authentication Protocol-Transport Layer Security (EAP-TLS).
    EAPTLS_ANDROIDEAPTYPE AndroidEapType = iota
    // Extensible Authentication Protocol-Tunneled Transport Layer Security (EAP-TTLS).
    EAPTTLS_ANDROIDEAPTYPE
    // Protected Extensible Authentication Protocol (PEAP).
    PEAP_ANDROIDEAPTYPE
)

func (i AndroidEapType) String() string {
    return []string{"eapTls", "eapTtls", "peap"}[i]
}
func ParseAndroidEapType(v string) (any, error) {
    result := EAPTLS_ANDROIDEAPTYPE
    switch v {
        case "eapTls":
            result = EAPTLS_ANDROIDEAPTYPE
        case "eapTtls":
            result = EAPTTLS_ANDROIDEAPTYPE
        case "peap":
            result = PEAP_ANDROIDEAPTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAndroidEapType(values []AndroidEapType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AndroidEapType) isMultiValue() bool {
    return false
}
