package models
type HardwareOathTokenHashFunction int

const (
    HMACSHA1_HARDWAREOATHTOKENHASHFUNCTION HardwareOathTokenHashFunction = iota
    HMACSHA256_HARDWAREOATHTOKENHASHFUNCTION
    UNKNOWNFUTUREVALUE_HARDWAREOATHTOKENHASHFUNCTION
)

func (i HardwareOathTokenHashFunction) String() string {
    return []string{"hmacsha1", "hmacsha256", "unknownFutureValue"}[i]
}
func ParseHardwareOathTokenHashFunction(v string) (any, error) {
    result := HMACSHA1_HARDWAREOATHTOKENHASHFUNCTION
    switch v {
        case "hmacsha1":
            result = HMACSHA1_HARDWAREOATHTOKENHASHFUNCTION
        case "hmacsha256":
            result = HMACSHA256_HARDWAREOATHTOKENHASHFUNCTION
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_HARDWAREOATHTOKENHASHFUNCTION
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeHardwareOathTokenHashFunction(values []HardwareOathTokenHashFunction) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i HardwareOathTokenHashFunction) isMultiValue() bool {
    return false
}
