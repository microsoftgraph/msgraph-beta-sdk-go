package graph
import (
    "strings"
    "errors"
)
// 
type CloudPcProvisioningPolicyImageType int

const (
    GALLERY_CLOUDPCPROVISIONINGPOLICYIMAGETYPE CloudPcProvisioningPolicyImageType = iota
    CUSTOM_CLOUDPCPROVISIONINGPOLICYIMAGETYPE
)

func (i CloudPcProvisioningPolicyImageType) String() string {
    return []string{"GALLERY", "CUSTOM"}[i]
}
func ParseCloudPcProvisioningPolicyImageType(v string) (interface{}, error) {
    result := GALLERY_CLOUDPCPROVISIONINGPOLICYIMAGETYPE
    switch strings.ToUpper(v) {
        case "GALLERY":
            result = GALLERY_CLOUDPCPROVISIONINGPOLICYIMAGETYPE
        case "CUSTOM":
            result = CUSTOM_CLOUDPCPROVISIONINGPOLICYIMAGETYPE
        default:
            return 0, errors.New("Unknown CloudPcProvisioningPolicyImageType value: " + v)
    }
    return &result, nil
}
func SerializeCloudPcProvisioningPolicyImageType(values []CloudPcProvisioningPolicyImageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
