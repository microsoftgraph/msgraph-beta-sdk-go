package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type StatusBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Possible values are: success, warning, failure, skipped, unknownFutureValue.
    status *ProvisioningResult;
}
// Instantiates a new statusBase and sets the default values.
func NewStatusBase()(*StatusBase) {
    m := &StatusBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *StatusBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the status property value. Possible values are: success, warning, failure, skipped, unknownFutureValue.
func (m *StatusBase) GetStatus()(*ProvisioningResult) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *StatusBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningResult)
        if err != nil {
            return err
        }
        cast := val.(ProvisioningResult)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *StatusBase) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *StatusBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *StatusBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the status property value. Possible values are: success, warning, failure, skipped, unknownFutureValue.
// Parameters:
//  - value : Value to set for the status property.
func (m *StatusBase) SetStatus(value *ProvisioningResult)() {
    m.status = value
}
