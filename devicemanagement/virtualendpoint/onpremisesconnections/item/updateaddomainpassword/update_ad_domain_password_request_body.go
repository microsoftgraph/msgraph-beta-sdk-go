package updateaddomainpassword

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UpdateAdDomainPasswordRequestBody 
type UpdateAdDomainPasswordRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    adDomainPassword *string;
}
// NewUpdateAdDomainPasswordRequestBody instantiates a new updateAdDomainPasswordRequestBody and sets the default values.
func NewUpdateAdDomainPasswordRequestBody()(*UpdateAdDomainPasswordRequestBody) {
    m := &UpdateAdDomainPasswordRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateAdDomainPasswordRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAdDomainPassword gets the adDomainPassword property value. 
func (m *UpdateAdDomainPasswordRequestBody) GetAdDomainPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.adDomainPassword
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateAdDomainPasswordRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["adDomainPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdDomainPassword(val)
        }
        return nil
    }
    return res
}
func (m *UpdateAdDomainPasswordRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UpdateAdDomainPasswordRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("adDomainPassword", m.GetAdDomainPassword())
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
func (m *UpdateAdDomainPasswordRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAdDomainPassword sets the adDomainPassword property value. 
func (m *UpdateAdDomainPasswordRequestBody) SetAdDomainPassword(value *string)() {
    if m != nil {
        m.adDomainPassword = value
    }
}
