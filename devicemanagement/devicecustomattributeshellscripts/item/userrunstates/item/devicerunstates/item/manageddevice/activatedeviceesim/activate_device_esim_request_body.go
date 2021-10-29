package activatedeviceesim

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ActivateDeviceEsimRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    carrierUrl *string;
}
// Instantiates a new activateDeviceEsimRequestBody and sets the default values.
func NewActivateDeviceEsimRequestBody()(*ActivateDeviceEsimRequestBody) {
    m := &ActivateDeviceEsimRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActivateDeviceEsimRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the carrierUrl property value. 
func (m *ActivateDeviceEsimRequestBody) GetCarrierUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.carrierUrl
    }
}
// The deserialization information for the current model
func (m *ActivateDeviceEsimRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["carrierUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCarrierUrl(val)
        return nil
    }
    return res
}
func (m *ActivateDeviceEsimRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ActivateDeviceEsimRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("carrierUrl", m.GetCarrierUrl())
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
func (m *ActivateDeviceEsimRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the carrierUrl property value. 
// Parameters:
//  - value : Value to set for the carrierUrl property.
func (m *ActivateDeviceEsimRequestBody) SetCarrierUrl(value *string)() {
    m.carrierUrl = value
}
