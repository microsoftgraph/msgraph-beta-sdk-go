package healthmonitoring
type Scenario int

const (
    UNKNOWN_SCENARIO Scenario = iota
    MFA_SCENARIO
    DEVICES_SCENARIO
    UNKNOWNFUTUREVALUE_SCENARIO
)

func (i Scenario) String() string {
    return []string{"unknown", "mfa", "devices", "unknownFutureValue"}[i]
}
func ParseScenario(v string) (any, error) {
    result := UNKNOWN_SCENARIO
    switch v {
        case "unknown":
            result = UNKNOWN_SCENARIO
        case "mfa":
            result = MFA_SCENARIO
        case "devices":
            result = DEVICES_SCENARIO
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_SCENARIO
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeScenario(values []Scenario) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Scenario) isMultiValue() bool {
    return false
}
