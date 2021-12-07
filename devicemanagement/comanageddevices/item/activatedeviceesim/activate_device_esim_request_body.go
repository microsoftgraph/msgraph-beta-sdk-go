package activatedeviceesim

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ActivateDeviceEsimRequestBody 
type ActivateDeviceEsimRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    carrierUrl *string;
}
// NewActivateDeviceEsimRequestBody instantiates a new activateDeviceEsimRequestBody and sets the default values.
func NewActivateDeviceEsimRequestBody()(*ActivateDeviceEsimRequestBody) {
    m := &ActivateDeviceEsimRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActivateDeviceEsimRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCarrierUrl gets the carrierUrl property value. 
func (m *ActivateDeviceEsimRequestBody) GetCarrierUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.carrierUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivateDeviceEsimRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["carrierUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCarrierUrl(val)
        }
        return nil
    }
    return res
}
func (m *ActivateDeviceEsimRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActivateDeviceEsimRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCarrierUrl sets the carrierUrl property value. 
func (m *ActivateDeviceEsimRequestBody) SetCarrierUrl(value *string)() {
    if m != nil {
        m.carrierUrl = value
    }
}
