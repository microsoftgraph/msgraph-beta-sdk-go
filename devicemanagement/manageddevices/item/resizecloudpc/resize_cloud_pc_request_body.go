package resizecloudpc

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResizeCloudPcRequestBody 
type ResizeCloudPcRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    targetServicePlanId *string;
}
// NewResizeCloudPcRequestBody instantiates a new resizeCloudPcRequestBody and sets the default values.
func NewResizeCloudPcRequestBody()(*ResizeCloudPcRequestBody) {
    m := &ResizeCloudPcRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResizeCloudPcRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetTargetServicePlanId gets the targetServicePlanId property value. 
func (m *ResizeCloudPcRequestBody) GetTargetServicePlanId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetServicePlanId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResizeCloudPcRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["targetServicePlanId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetServicePlanId(val)
        }
        return nil
    }
    return res
}
func (m *ResizeCloudPcRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResizeCloudPcRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTargetServicePlanId sets the targetServicePlanId property value. 
func (m *ResizeCloudPcRequestBody) SetTargetServicePlanId(value *string)() {
    if m != nil {
        m.targetServicePlanId = value
    }
}
