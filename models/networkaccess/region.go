package networkaccess
import (
    "errors"
)
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
    AUSTRALIAEAST_REGION
    WESTCENTRALUS_REGION
    CENTRALINDIA_REGION
    SOUTHEASTASIA_REGION
    SWEDENCENTRAL_REGION
    SOUTHINDIA_REGION
    AUSTRALIASOUTHEAST_REGION
    KOREACENTRAL_REGION
    POLANDCENTRAL_REGION
    BRAZILSOUTH_REGION
    JAPANEAST_REGION
    JAPANWEST_REGION
    KOREASOUTH_REGION
    ITALYNORTH_REGION
    FRANCESOUTH_REGION
    ISRAELCENTRAL_REGION
    UNKNOWNFUTUREVALUE_REGION
)

func (i Region) String() string {
    return []string{"eastUS", "eastUS2", "westUS", "westUS2", "westUS3", "centralUS", "northCentralUS", "southCentralUS", "northEurope", "westEurope", "franceCentral", "germanyWestCentral", "switzerlandNorth", "ukSouth", "canadaEast", "canadaCentral", "southAfricaWest", "southAfricaNorth", "uaeNorth", "australiaEast", "westCentralUS", "centralIndia", "southEastAsia", "swedenCentral", "southIndia", "australiaSouthEast", "koreaCentral", "polandCentral", "brazilSouth", "japanEast", "japanWest", "koreaSouth", "italyNorth", "franceSouth", "israelCentral", "unknownFutureValue"}[i]
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
        case "australiaEast":
            result = AUSTRALIAEAST_REGION
        case "westCentralUS":
            result = WESTCENTRALUS_REGION
        case "centralIndia":
            result = CENTRALINDIA_REGION
        case "southEastAsia":
            result = SOUTHEASTASIA_REGION
        case "swedenCentral":
            result = SWEDENCENTRAL_REGION
        case "southIndia":
            result = SOUTHINDIA_REGION
        case "australiaSouthEast":
            result = AUSTRALIASOUTHEAST_REGION
        case "koreaCentral":
            result = KOREACENTRAL_REGION
        case "polandCentral":
            result = POLANDCENTRAL_REGION
        case "brazilSouth":
            result = BRAZILSOUTH_REGION
        case "japanEast":
            result = JAPANEAST_REGION
        case "japanWest":
            result = JAPANWEST_REGION
        case "koreaSouth":
            result = KOREASOUTH_REGION
        case "italyNorth":
            result = ITALYNORTH_REGION
        case "franceSouth":
            result = FRANCESOUTH_REGION
        case "israelCentral":
            result = ISRAELCENTRAL_REGION
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
func (i Region) isMultiValue() bool {
    return false
}
