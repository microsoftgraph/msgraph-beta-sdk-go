package getglobalscripthighestavailableversion

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetGlobalScriptHighestAvailableVersionResponse provides operations to call the getGlobalScriptHighestAvailableVersion method.
type GetGlobalScriptHighestAvailableVersionResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    value *string;
}
// NewGetGlobalScriptHighestAvailableVersionResponse instantiates a new getGlobalScriptHighestAvailableVersionResponse and sets the default values.
func NewGetGlobalScriptHighestAvailableVersionResponse()(*GetGlobalScriptHighestAvailableVersionResponse) {
    m := &GetGlobalScriptHighestAvailableVersionResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetGlobalScriptHighestAvailableVersionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetGlobalScriptHighestAvailableVersionResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewGetGlobalScriptHighestAvailableVersionResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetGlobalScriptHighestAvailableVersionResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetGlobalScriptHighestAvailableVersionResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. 
func (m *GetGlobalScriptHighestAvailableVersionResponse) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *GetGlobalScriptHighestAvailableVersionResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetGlobalScriptHighestAvailableVersionResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *GetGlobalScriptHighestAvailableVersionResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. 
func (m *GetGlobalScriptHighestAvailableVersionResponse) SetValue(value *string)() {
    if m != nil {
        m.value = value
    }
}