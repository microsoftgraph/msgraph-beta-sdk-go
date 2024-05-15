package networkaccess
type BandwidthCapacityInMbps int

const (
    MBPS250_BANDWIDTHCAPACITYINMBPS BandwidthCapacityInMbps = iota
    MBPS500_BANDWIDTHCAPACITYINMBPS
    MBPS750_BANDWIDTHCAPACITYINMBPS
    MBPS1000_BANDWIDTHCAPACITYINMBPS
    UNKNOWNFUTUREVALUE_BANDWIDTHCAPACITYINMBPS
)

func (i BandwidthCapacityInMbps) String() string {
    return []string{"mbps250", "mbps500", "mbps750", "mbps1000", "unknownFutureValue"}[i]
}
func ParseBandwidthCapacityInMbps(v string) (any, error) {
    result := MBPS250_BANDWIDTHCAPACITYINMBPS
    switch v {
        case "mbps250":
            result = MBPS250_BANDWIDTHCAPACITYINMBPS
        case "mbps500":
            result = MBPS500_BANDWIDTHCAPACITYINMBPS
        case "mbps750":
            result = MBPS750_BANDWIDTHCAPACITYINMBPS
        case "mbps1000":
            result = MBPS1000_BANDWIDTHCAPACITYINMBPS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_BANDWIDTHCAPACITYINMBPS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeBandwidthCapacityInMbps(values []BandwidthCapacityInMbps) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i BandwidthCapacityInMbps) isMultiValue() bool {
    return false
}
