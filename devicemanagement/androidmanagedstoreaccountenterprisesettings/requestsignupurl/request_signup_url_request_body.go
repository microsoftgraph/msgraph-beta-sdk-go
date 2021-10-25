package requestsignupurl

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RequestSignupUrlRequestBody struct {
    additionalData map[string]interface{};
    hostName *string;
}
func NewRequestSignupUrlRequestBody()(*RequestSignupUrlRequestBody) {
    m := &RequestSignupUrlRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RequestSignupUrlRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RequestSignupUrlRequestBody) GetHostName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hostName
    }
}
func (m *RequestSignupUrlRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hostName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHostName(val)
        return nil
    }
    return res
}
func (m *RequestSignupUrlRequestBody) IsNil()(bool) {
    return m == nil
}
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
func (m *RequestSignupUrlRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RequestSignupUrlRequestBody) SetHostName(value *string)() {
    m.hostName = value
}
