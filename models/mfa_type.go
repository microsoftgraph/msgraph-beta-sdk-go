package models
type MfaType int

const (
    EOTP_MFATYPE MfaType = iota
    ONEWAYSMS_MFATYPE
    TWOWAYSMS_MFATYPE
    TWOWAYSMSOTHERMOBILE_MFATYPE
    PHONEAPPNOTIFICATION_MFATYPE
    PHONEAPPOTP_MFATYPE
    TWOWAYVOICEMOBILE_MFATYPE
    TWOWAYVOICEOFFICE_MFATYPE
    TWOWAYVOICEOTHERMOBILE_MFATYPE
    FIDO_MFATYPE
    CERTIFICATE_MFATYPE
    OTHER_MFATYPE
    UNKNOWNFUTUREVALUE_MFATYPE
)

func (i MfaType) String() string {
    return []string{"eotp", "oneWaySms", "twoWaySms", "twoWaySmsOtherMobile", "phoneAppNotification", "phoneAppOtp", "twoWayVoiceMobile", "twoWayVoiceOffice", "twoWayVoiceOtherMobile", "fido", "certificate", "other", "unknownFutureValue"}[i]
}
func ParseMfaType(v string) (any, error) {
    result := EOTP_MFATYPE
    switch v {
        case "eotp":
            result = EOTP_MFATYPE
        case "oneWaySms":
            result = ONEWAYSMS_MFATYPE
        case "twoWaySms":
            result = TWOWAYSMS_MFATYPE
        case "twoWaySmsOtherMobile":
            result = TWOWAYSMSOTHERMOBILE_MFATYPE
        case "phoneAppNotification":
            result = PHONEAPPNOTIFICATION_MFATYPE
        case "phoneAppOtp":
            result = PHONEAPPOTP_MFATYPE
        case "twoWayVoiceMobile":
            result = TWOWAYVOICEMOBILE_MFATYPE
        case "twoWayVoiceOffice":
            result = TWOWAYVOICEOFFICE_MFATYPE
        case "twoWayVoiceOtherMobile":
            result = TWOWAYVOICEOTHERMOBILE_MFATYPE
        case "fido":
            result = FIDO_MFATYPE
        case "certificate":
            result = CERTIFICATE_MFATYPE
        case "other":
            result = OTHER_MFATYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MFATYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeMfaType(values []MfaType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MfaType) isMultiValue() bool {
    return false
}
