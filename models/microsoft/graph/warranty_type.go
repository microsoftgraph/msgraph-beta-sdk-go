package graph
import (
    "strings"
    "errors"
)
// 
type WarrantyType int

const (
    UNKNOWN_WARRANTYTYPE WarrantyType = iota
    MANUFACTURER_WARRANTYTYPE
    CONTRACTUAL_WARRANTYTYPE
    UNKNOWNFUTUREVALUE_WARRANTYTYPE
)

func (i WarrantyType) String() string {
    return []string{"UNKNOWN", "MANUFACTURER", "CONTRACTUAL", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWarrantyType(v string) (interface{}, error) {
    result := UNKNOWN_WARRANTYTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_WARRANTYTYPE
        case "MANUFACTURER":
            result = MANUFACTURER_WARRANTYTYPE
        case "CONTRACTUAL":
            result = CONTRACTUAL_WARRANTYTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WARRANTYTYPE
        default:
            return 0, errors.New("Unknown WarrantyType value: " + v)
    }
    return &result, nil
}
func SerializeWarrantyType(values []WarrantyType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
