package updatewindowsdeviceaccount

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// UpdateWindowsDeviceAccountRequestBody 
type UpdateWindowsDeviceAccountRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    updateWindowsDeviceAccountActionParameter *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdateWindowsDeviceAccountActionParameter;
}
// NewUpdateWindowsDeviceAccountRequestBody instantiates a new updateWindowsDeviceAccountRequestBody and sets the default values.
func NewUpdateWindowsDeviceAccountRequestBody()(*UpdateWindowsDeviceAccountRequestBody) {
    m := &UpdateWindowsDeviceAccountRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateWindowsDeviceAccountRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetUpdateWindowsDeviceAccountActionParameter gets the updateWindowsDeviceAccountActionParameter property value. 
func (m *UpdateWindowsDeviceAccountRequestBody) GetUpdateWindowsDeviceAccountActionParameter()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdateWindowsDeviceAccountActionParameter) {
    if m == nil {
        return nil
    } else {
        return m.updateWindowsDeviceAccountActionParameter
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateWindowsDeviceAccountRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["updateWindowsDeviceAccountActionParameter"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewUpdateWindowsDeviceAccountActionParameter() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateWindowsDeviceAccountActionParameter(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdateWindowsDeviceAccountActionParameter))
        }
        return nil
    }
    return res
}
func (m *UpdateWindowsDeviceAccountRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UpdateWindowsDeviceAccountRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("updateWindowsDeviceAccountActionParameter", m.GetUpdateWindowsDeviceAccountActionParameter())
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
func (m *UpdateWindowsDeviceAccountRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetUpdateWindowsDeviceAccountActionParameter sets the updateWindowsDeviceAccountActionParameter property value. 
func (m *UpdateWindowsDeviceAccountRequestBody) SetUpdateWindowsDeviceAccountActionParameter(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.UpdateWindowsDeviceAccountActionParameter)() {
    if m != nil {
        m.updateWindowsDeviceAccountActionParameter = value
    }
}
