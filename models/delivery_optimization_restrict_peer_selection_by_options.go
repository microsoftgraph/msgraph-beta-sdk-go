package models
import (
    "errors"
)
// Values to restrict peer selection by.
type DeliveryOptimizationRestrictPeerSelectionByOptions int

const (
    // Not configured.
    NOTCONFIGURED_DELIVERYOPTIMIZATIONRESTRICTPEERSELECTIONBYOPTIONS DeliveryOptimizationRestrictPeerSelectionByOptions = iota
    // Subnet mask.
    SUBNETMASK_DELIVERYOPTIMIZATIONRESTRICTPEERSELECTIONBYOPTIONS
)

func (i DeliveryOptimizationRestrictPeerSelectionByOptions) String() string {
    return []string{"notConfigured", "subnetMask"}[i]
}
func ParseDeliveryOptimizationRestrictPeerSelectionByOptions(v string) (any, error) {
    result := NOTCONFIGURED_DELIVERYOPTIMIZATIONRESTRICTPEERSELECTIONBYOPTIONS
    switch v {
        case "notConfigured":
            result = NOTCONFIGURED_DELIVERYOPTIMIZATIONRESTRICTPEERSELECTIONBYOPTIONS
        case "subnetMask":
            result = SUBNETMASK_DELIVERYOPTIMIZATIONRESTRICTPEERSELECTIONBYOPTIONS
        default:
            return 0, errors.New("Unknown DeliveryOptimizationRestrictPeerSelectionByOptions value: " + v)
    }
    return &result, nil
}
func SerializeDeliveryOptimizationRestrictPeerSelectionByOptions(values []DeliveryOptimizationRestrictPeerSelectionByOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
