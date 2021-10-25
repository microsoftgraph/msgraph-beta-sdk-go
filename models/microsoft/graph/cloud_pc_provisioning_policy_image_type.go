package graph
import (
    "strings"
    "errors"
)
type CloudPcProvisioningPolicyImageType int

const (
    GALLERY_CLOUDPCPROVISIONINGPOLICYIMAGETYPE CloudPcProvisioningPolicyImageType = iota
    CUSTOM_CLOUDPCPROVISIONINGPOLICYIMAGETYPE
)

func (i CloudPcProvisioningPolicyImageType) String() string {
    return []string{"GALLERY", "CUSTOM"}[i]
}
func ParseCloudPcProvisioningPolicyImageType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "GALLERY":
            return GALLERY_CLOUDPCPROVISIONINGPOLICYIMAGETYPE, nil
        case "CUSTOM":
            return CUSTOM_CLOUDPCPROVISIONINGPOLICYIMAGETYPE, nil
    }
    return 0, errors.New("Unknown CloudPcProvisioningPolicyImageType value: " + v)
}
func SerializeCloudPcProvisioningPolicyImageType(values []CloudPcProvisioningPolicyImageType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
