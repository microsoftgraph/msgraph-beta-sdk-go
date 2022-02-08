package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CustomExtensionClientConfiguration 
type CustomExtensionClientConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    timeoutInMilliseconds *int32;
}
// NewCustomExtensionClientConfiguration instantiates a new customExtensionClientConfiguration and sets the default values.
func NewCustomExtensionClientConfiguration()(*CustomExtensionClientConfiguration) {
    m := &CustomExtensionClientConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CustomExtensionClientConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetTimeoutInMilliseconds gets the timeoutInMilliseconds property value. 
func (m *CustomExtensionClientConfiguration) GetTimeoutInMilliseconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.timeoutInMilliseconds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CustomExtensionClientConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["timeoutInMilliseconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeoutInMilliseconds(val)
        }
        return nil
    }
    return res
}
func (m *CustomExtensionClientConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CustomExtensionClientConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("timeoutInMilliseconds", m.GetTimeoutInMilliseconds())
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
func (m *CustomExtensionClientConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTimeoutInMilliseconds sets the timeoutInMilliseconds property value. 
func (m *CustomExtensionClientConfiguration) SetTimeoutInMilliseconds(value *int32)() {
    if m != nil {
        m.timeoutInMilliseconds = value
    }
}
