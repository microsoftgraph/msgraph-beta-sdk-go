package models
// Windows 10 VPN connection types.
type Windows10VpnAuthenticationMethod int

const (
    // Authenticate with a certificate.
    CERTIFICATE_WINDOWS10VPNAUTHENTICATIONMETHOD Windows10VpnAuthenticationMethod = iota
    // Use username and password for authentication.
    USERNAMEANDPASSWORD_WINDOWS10VPNAUTHENTICATIONMETHOD
    // Authentication method is specified in custom EAP XML.
    CUSTOMEAPXML_WINDOWS10VPNAUTHENTICATIONMETHOD
    // Use Derived Credential for authentication.
    DERIVEDCREDENTIAL_WINDOWS10VPNAUTHENTICATIONMETHOD
)

func (i Windows10VpnAuthenticationMethod) String() string {
    return []string{"certificate", "usernameAndPassword", "customEapXml", "derivedCredential"}[i]
}
func ParseWindows10VpnAuthenticationMethod(v string) (any, error) {
    result := CERTIFICATE_WINDOWS10VPNAUTHENTICATIONMETHOD
    switch v {
        case "certificate":
            result = CERTIFICATE_WINDOWS10VPNAUTHENTICATIONMETHOD
        case "usernameAndPassword":
            result = USERNAMEANDPASSWORD_WINDOWS10VPNAUTHENTICATIONMETHOD
        case "customEapXml":
            result = CUSTOMEAPXML_WINDOWS10VPNAUTHENTICATIONMETHOD
        case "derivedCredential":
            result = DERIVEDCREDENTIAL_WINDOWS10VPNAUTHENTICATIONMETHOD
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeWindows10VpnAuthenticationMethod(values []Windows10VpnAuthenticationMethod) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Windows10VpnAuthenticationMethod) isMultiValue() bool {
    return false
}
