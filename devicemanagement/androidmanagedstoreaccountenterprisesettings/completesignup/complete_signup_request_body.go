package completesignup

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CompleteSignupRequestBody 
type CompleteSignupRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    enterpriseToken *string;
}
// NewCompleteSignupRequestBody instantiates a new completeSignupRequestBody and sets the default values.
func NewCompleteSignupRequestBody()(*CompleteSignupRequestBody) {
    m := &CompleteSignupRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompleteSignupRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnterpriseToken gets the enterpriseToken property value. 
func (m *CompleteSignupRequestBody) GetEnterpriseToken()(*string) {
    if m == nil {
        return nil
    } else {
        return m.enterpriseToken
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CompleteSignupRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enterpriseToken"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseToken(val)
        }
        return nil
    }
    return res
}
func (m *CompleteSignupRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompleteSignupRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnterpriseToken sets the enterpriseToken property value. 
func (m *CompleteSignupRequestBody) SetEnterpriseToken(value *string)() {
    if m != nil {
        m.enterpriseToken = value
    }
}
