package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MfaDetail 
type MfaDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates the MFA auth detail for the corresponding Sign-in activity when the MFA Required is 'Yes'.
    authDetail *string;
    // Indicates the MFA Auth methods (SMS, Phone, Authenticator App are some of the value) for the corresponding sign-in activity when the MFA Required field is 'Yes'.
    authMethod *string;
}
// NewMfaDetail instantiates a new mfaDetail and sets the default values.
func NewMfaDetail()(*MfaDetail) {
    m := &MfaDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MfaDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAuthDetail gets the authDetail property value. Indicates the MFA auth detail for the corresponding Sign-in activity when the MFA Required is 'Yes'.
func (m *MfaDetail) GetAuthDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authDetail
    }
}
// GetAuthMethod gets the authMethod property value. Indicates the MFA Auth methods (SMS, Phone, Authenticator App are some of the value) for the corresponding sign-in activity when the MFA Required field is 'Yes'.
func (m *MfaDetail) GetAuthMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authMethod
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MfaDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthDetail(val)
        }
        return nil
    }
    res["authMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthMethod(val)
        }
        return nil
    }
    return res
}
func (m *MfaDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MfaDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authDetail", m.GetAuthDetail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("authMethod", m.GetAuthMethod())
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
func (m *MfaDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAuthDetail sets the authDetail property value. Indicates the MFA auth detail for the corresponding Sign-in activity when the MFA Required is 'Yes'.
func (m *MfaDetail) SetAuthDetail(value *string)() {
    if m != nil {
        m.authDetail = value
    }
}
// SetAuthMethod sets the authMethod property value. Indicates the MFA Auth methods (SMS, Phone, Authenticator App are some of the value) for the corresponding sign-in activity when the MFA Required field is 'Yes'.
func (m *MfaDetail) SetAuthMethod(value *string)() {
    if m != nil {
        m.authMethod = value
    }
}
