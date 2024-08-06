package models
type ContentModelType int

const (
    TEACHINGMETHOD_CONTENTMODELTYPE ContentModelType = iota
    LAYOUTMETHOD_CONTENTMODELTYPE
    FREEFORMSELECTIONMETHOD_CONTENTMODELTYPE
    PREBUILTCONTRACTMODEL_CONTENTMODELTYPE
    PREBUILTINVOICEMODEL_CONTENTMODELTYPE
    PREBUILTRECEIPTMODEL_CONTENTMODELTYPE
    UNKNOWNFUTUREVALUE_CONTENTMODELTYPE
)

func (i ContentModelType) String() string {
    return []string{"teachingMethod", "layoutMethod", "freeformSelectionMethod", "prebuiltContractModel", "prebuiltInvoiceModel", "prebuiltReceiptModel", "unknownFutureValue"}[i]
}
func ParseContentModelType(v string) (any, error) {
    result := TEACHINGMETHOD_CONTENTMODELTYPE
    switch v {
        case "teachingMethod":
            result = TEACHINGMETHOD_CONTENTMODELTYPE
        case "layoutMethod":
            result = LAYOUTMETHOD_CONTENTMODELTYPE
        case "freeformSelectionMethod":
            result = FREEFORMSELECTIONMETHOD_CONTENTMODELTYPE
        case "prebuiltContractModel":
            result = PREBUILTCONTRACTMODEL_CONTENTMODELTYPE
        case "prebuiltInvoiceModel":
            result = PREBUILTINVOICEMODEL_CONTENTMODELTYPE
        case "prebuiltReceiptModel":
            result = PREBUILTRECEIPTMODEL_CONTENTMODELTYPE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_CONTENTMODELTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeContentModelType(values []ContentModelType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ContentModelType) isMultiValue() bool {
    return false
}
