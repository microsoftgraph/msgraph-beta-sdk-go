package requestsignupurl

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RequestSignupUrlRequestBody 
type RequestSignupUrlRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    hostName *string;
}
// NewRequestSignupUrlRequestBody instantiates a new requestSignupUrlRequestBody and sets the default values.
func NewRequestSignupUrlRequestBody()(*RequestSignupUrlRequestBody) {
    m := &RequestSignupUrlRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RequestSignupUrlRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetHostName gets the hostName property value. 
func (m *RequestSignupUrlRequestBody) GetHostName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hostName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RequestSignupUrlRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hostName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostName(val)
        }
        return nil
    }
    return res
}
func (m *RequestSignupUrlRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *RequestSignupUrlRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("hostName", m.GetHostName())
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
func (m *RequestSignupUrlRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetHostName sets the hostName property value. 
func (m *RequestSignupUrlRequestBody) SetHostName(value *string)() {
    if m != nil {
        m.hostName = value
    }
}
