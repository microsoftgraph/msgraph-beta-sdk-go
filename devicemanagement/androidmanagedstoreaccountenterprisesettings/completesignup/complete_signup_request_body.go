package completesignup

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type CompleteSignupRequestBody struct {
    additionalData map[string]interface{};
    enterpriseToken *string;
}
func NewCompleteSignupRequestBody()(*CompleteSignupRequestBody) {
    m := &CompleteSignupRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *CompleteSignupRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *CompleteSignupRequestBody) GetEnterpriseToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseToken
    }
}
func (m *CompleteSignupRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enterpriseToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEnterpriseToken(val)
        return nil
    }
    return res
}
func (m *CompleteSignupRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *CompleteSignupRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("enterpriseToken", m.GetEnterpriseToken())
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
func (m *CompleteSignupRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *CompleteSignupRequestBody) SetEnterpriseToken(value *string)() {
    m.enterpriseToken = value
}
