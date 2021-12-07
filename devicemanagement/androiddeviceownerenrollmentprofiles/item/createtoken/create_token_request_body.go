package createtoken

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CreateTokenRequestBody 
type CreateTokenRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    tokenValidityInSeconds *int32;
}
// NewCreateTokenRequestBody instantiates a new createTokenRequestBody and sets the default values.
func NewCreateTokenRequestBody()(*CreateTokenRequestBody) {
    m := &CreateTokenRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateTokenRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetTokenValidityInSeconds gets the tokenValidityInSeconds property value. 
func (m *CreateTokenRequestBody) GetTokenValidityInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.tokenValidityInSeconds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateTokenRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["tokenValidityInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenValidityInSeconds(val)
        }
        return nil
    }
    return res
}
func (m *CreateTokenRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CreateTokenRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("tokenValidityInSeconds", m.GetTokenValidityInSeconds())
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
func (m *CreateTokenRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTokenValidityInSeconds sets the tokenValidityInSeconds property value. 
func (m *CreateTokenRequestBody) SetTokenValidityInSeconds(value *int32)() {
    if m != nil {
        m.tokenValidityInSeconds = value
    }
}
