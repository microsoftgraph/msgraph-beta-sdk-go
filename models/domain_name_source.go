// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models
// Domainname source.
type DomainNameSource int

const (
    // Full domain name.
    FULLDOMAINNAME_DOMAINNAMESOURCE DomainNameSource = iota
    // net bios domain name.
    NETBIOSDOMAINNAME_DOMAINNAMESOURCE
)

func (i DomainNameSource) String() string {
    return []string{"fullDomainName", "netBiosDomainName"}[i]
}
func ParseDomainNameSource(v string) (any, error) {
    result := FULLDOMAINNAME_DOMAINNAMESOURCE
    switch v {
        case "fullDomainName":
            result = FULLDOMAINNAME_DOMAINNAMESOURCE
        case "netBiosDomainName":
            result = NETBIOSDOMAINNAME_DOMAINNAMESOURCE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDomainNameSource(values []DomainNameSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DomainNameSource) isMultiValue() bool {
    return false
}
