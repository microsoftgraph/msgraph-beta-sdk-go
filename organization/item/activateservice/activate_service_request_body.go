package activateservice

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ActivateServiceRequestBody provides operations to call the activateService method.
type ActivateServiceRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    service *string;
    // 
    servicePlanId *string;
    // 
    skuId *string;
}
// NewActivateServiceRequestBody instantiates a new activateServiceRequestBody and sets the default values.
func NewActivateServiceRequestBody()(*ActivateServiceRequestBody) {
    m := &ActivateServiceRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateActivateServiceRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActivateServiceRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewActivateServiceRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ActivateServiceRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivateServiceRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
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
// GetService gets the service property value. 
func (m *ActivateServiceRequestBody) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// GetServicePlanId gets the servicePlanId property value. 
func (m *ActivateServiceRequestBody) GetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.servicePlanId
    }
}
// GetSkuId gets the skuId property value. 
func (m *ActivateServiceRequestBody) GetSkuId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.skuId
    }
}
func (m *ActivateServiceRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ActivateServiceRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
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
func (m *ActivateServiceRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetService sets the service property value. 
func (m *ActivateServiceRequestBody) SetService(value *string)() {
    if m != nil {
        m.service = value
    }
}
// SetServicePlanId sets the servicePlanId property value. 
func (m *ActivateServiceRequestBody) SetServicePlanId(value *string)() {
    if m != nil {
        m.servicePlanId = value
    }
}
// SetSkuId sets the skuId property value. 
func (m *ActivateServiceRequestBody) SetSkuId(value *string)() {
    if m != nil {
        m.skuId = value
    }
}
