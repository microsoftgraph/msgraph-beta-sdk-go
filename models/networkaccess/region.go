package networkaccess
import (
    "errors"
)
// 
type Region int

const (
    EASTUS_REGION Region = iota
    EASTUS2_REGION
    WESTUS_REGION
    WESTUS2_REGION
    WESTUS3_REGION
    CENTRALUS_REGION
    NORTHCENTRALUS_REGION
    SOUTHCENTRALUS_REGION
    NORTHEUROPE_REGION
    WESTEUROPE_REGION
    FRANCECENTRAL_REGION
    GERMANYWESTCENTRAL_REGION
    SWITZERLANDNORTH_REGION
    UKSOUTH_REGION
    CANADAEAST_REGION
    CANADACENTRAL_REGION
    SOUTHAFRICAWEST_REGION
    SOUTHAFRICANORTH_REGION
    UAENORTH_REGION
    UNKNOWNFUTUREVALUE_REGION
)

func (i Region) String() string {
    return []string{"eastUS", "eastUS2", "westUS", "westUS2", "westUS3", "centralUS", "northCentralUS", "southCentralUS", "northEurope", "westEurope", "franceCentral", "germanyWestCentral", "switzerlandNorth", "ukSouth", "canadaEast", "canadaCentral", "southAfricaWest", "southAfricaNorth", "uaeNorth", "unknownFutureValue"}[i]
}
func ParseRegion(v string) (any, error) {
    result := EASTUS_REGION
    switch v {
        case "eastUS":
            result = EASTUS_REGION
        case "eastUS2":
            result = EASTUS2_REGION
        case "westUS":
            result = WESTUS_REGION
        case "westUS2":
            result = WESTUS2_REGION
        case "westUS3":
            result = WESTUS3_REGION
        case "centralUS":
            result = CENTRALUS_REGION
        case "northCentralUS":
            result = NORTHCENTRALUS_REGION
        case "southCentralUS":
            result = SOUTHCENTRALUS_REGION
        case "northEurope":
            result = NORTHEUROPE_REGION
        case "westEurope":
            result = WESTEUROPE_REGION
        case "franceCentral":
            result = FRANCECENTRAL_REGION
        case "germanyWestCentral":
            result = GERMANYWESTCENTRAL_REGION
        case "switzerlandNorth":
            result = SWITZERLANDNORTH_REGION
        case "ukSouth":
            result = UKSOUTH_REGION
        case "canadaEast":
            result = CANADAEAST_REGION
        case "canadaCentral":
            result = CANADACENTRAL_REGION
        case "southAfricaWest":
            result = SOUTHAFRICAWEST_REGION
        case "southAfricaNorth":
            result = SOUTHAFRICANORTH_REGION
        case "uaeNorth":
            result = UAENORTH_REGION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_REGION
        default:
            return 0, errors.New("Unknown Region value: " + v)
    }
    return &result, nil
}
func SerializeRegion(values []Region) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
