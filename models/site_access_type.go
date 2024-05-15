package models
type SiteAccessType int

const (
    BLOCK_SITEACCESSTYPE SiteAccessType = iota
    FULL_SITEACCESSTYPE
    LIMITED_SITEACCESSTYPE
)

func (i SiteAccessType) String() string {
    return []string{"block", "full", "limited"}[i]
}
func ParseSiteAccessType(v string) (any, error) {
    result := BLOCK_SITEACCESSTYPE
    switch v {
        case "block":
            result = BLOCK_SITEACCESSTYPE
        case "full":
            result = FULL_SITEACCESSTYPE
        case "limited":
            result = LIMITED_SITEACCESSTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeSiteAccessType(values []SiteAccessType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i SiteAccessType) isMultiValue() bool {
    return false
}
