package resizecloudpc

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ResizeCloudPcRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    targetServicePlanId *string;
}
// Instantiates a new resizeCloudPcRequestBody and sets the default values.
func NewResizeCloudPcRequestBody()(*ResizeCloudPcRequestBody) {
    m := &ResizeCloudPcRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResizeCloudPcRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the targetServicePlanId property value. 
func (m *ResizeCloudPcRequestBody) GetTargetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetServicePlanId
    }
}
// The deserialization information for the current model
func (m *ResizeCloudPcRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["targetServicePlanId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetServicePlanId(val)
        return nil
    }
    return res
}
func (m *ResizeCloudPcRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ResizeCloudPcRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("targetServicePlanId", m.GetTargetServicePlanId())
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
func (m *ResizeCloudPcRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the targetServicePlanId property value. 
// Parameters:
//  - value : Value to set for the targetServicePlanId property.
func (m *ResizeCloudPcRequestBody) SetTargetServicePlanId(value *string)() {
    m.targetServicePlanId = value
}
