package activateserviceplan

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ActivateServicePlanRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    servicePlanId *string;
    // 
    skuId *string;
}
// Instantiates a new activateServicePlanRequestBody and sets the default values.
func NewActivateServicePlanRequestBody()(*ActivateServicePlanRequestBody) {
    m := &ActivateServicePlanRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActivateServicePlanRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the servicePlanId property value. 
func (m *ActivateServicePlanRequestBody) GetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanId
    }
}
// Gets the skuId property value. 
func (m *ActivateServicePlanRequestBody) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
// The deserialization information for the current model
func (m *ActivateServicePlanRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["servicePlanId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetServicePlanId(val)
        return nil
    }
    res["skuId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSkuId(val)
        return nil
    }
    return res
}
func (m *ActivateServicePlanRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ActivateServicePlanRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the servicePlanId property value. 
// Parameters:
//  - value : Value to set for the servicePlanId property.
func (m *ActivateServicePlanRequestBody) SetServicePlanId(value *string)() {
    m.servicePlanId = value
}
// Sets the skuId property value. 
// Parameters:
//  - value : Value to set for the skuId property.
func (m *ActivateServicePlanRequestBody) SetSkuId(value *string)() {
    m.skuId = value
}
