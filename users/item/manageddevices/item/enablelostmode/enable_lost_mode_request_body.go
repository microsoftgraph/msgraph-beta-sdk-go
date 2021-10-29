package enablelostmode

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EnableLostModeRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    footer *string;
    // 
    message *string;
    // 
    phoneNumber *string;
}
// Instantiates a new enableLostModeRequestBody and sets the default values.
func NewEnableLostModeRequestBody()(*EnableLostModeRequestBody) {
    m := &EnableLostModeRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnableLostModeRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the footer property value. 
func (m *EnableLostModeRequestBody) GetFooter()(*string) {
    if m == nil {
        return nil
    } else {
        return m.footer
    }
}
// Gets the message property value. 
func (m *EnableLostModeRequestBody) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// Gets the phoneNumber property value. 
func (m *EnableLostModeRequestBody) GetPhoneNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneNumber
    }
}
// The deserialization information for the current model
func (m *EnableLostModeRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["footer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFooter(val)
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    res["phoneNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPhoneNumber(val)
        return nil
    }
    return res
}
func (m *EnableLostModeRequestBody) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EnableLostModeRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("footer", m.GetFooter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("phoneNumber", m.GetPhoneNumber())
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
func (m *EnableLostModeRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the footer property value. 
// Parameters:
//  - value : Value to set for the footer property.
func (m *EnableLostModeRequestBody) SetFooter(value *string)() {
    m.footer = value
}
// Sets the message property value. 
// Parameters:
//  - value : Value to set for the message property.
func (m *EnableLostModeRequestBody) SetMessage(value *string)() {
    m.message = value
}
// Sets the phoneNumber property value. 
// Parameters:
//  - value : Value to set for the phoneNumber property.
func (m *EnableLostModeRequestBody) SetPhoneNumber(value *string)() {
    m.phoneNumber = value
}
