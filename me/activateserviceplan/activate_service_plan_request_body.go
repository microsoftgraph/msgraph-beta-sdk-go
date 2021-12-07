package activateserviceplan

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ActivateServicePlanRequestBody 
type ActivateServicePlanRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    servicePlanId *string;
    // 
    skuId *string;
}
// NewActivateServicePlanRequestBody instantiates a new activateServicePlanRequestBody and sets the default values.
func NewActivateServicePlanRequestBody()(*ActivateServicePlanRequestBody) {
    m := &ActivateServicePlanRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActivateServicePlanRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetServicePlanId gets the servicePlanId property value. 
func (m *ActivateServicePlanRequestBody) GetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanId
    }
}
// GetSkuId gets the skuId property value. 
func (m *ActivateServicePlanRequestBody) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivateServicePlanRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["servicePlanId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    res["skuId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    return res
}
func (m *ActivateServicePlanRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ActivateServicePlanRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("servicePlanId", m.GetServicePlanId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("skuId", m.GetSkuId())
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
func (m *ActivateServicePlanRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetServicePlanId sets the servicePlanId property value. 
func (m *ActivateServicePlanRequestBody) SetServicePlanId(value *string)() {
    if m != nil {
        m.servicePlanId = value
    }
}
// SetSkuId sets the skuId property value. 
func (m *ActivateServicePlanRequestBody) SetSkuId(value *string)() {
    if m != nil {
        m.skuId = value
    }
}
