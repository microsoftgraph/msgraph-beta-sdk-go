package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MfaDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates the MFA auth detail for the corresponding Sign-in activity when the MFA Required is 'Yes'.
    authDetail *string;
    // Indicates the MFA Auth methods (SMS, Phone, Authenticator App are some of the value) for the corresponding sign-in activity when the MFA Required field is 'Yes'.
    authMethod *string;
}
// Instantiates a new mfaDetail and sets the default values.
func NewMfaDetail()(*MfaDetail) {
    m := &MfaDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MfaDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the authDetail property value. Indicates the MFA auth detail for the corresponding Sign-in activity when the MFA Required is 'Yes'.
func (m *MfaDetail) GetAuthDetail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authDetail
    }
}
// Gets the authMethod property value. Indicates the MFA Auth methods (SMS, Phone, Authenticator App are some of the value) for the corresponding sign-in activity when the MFA Required field is 'Yes'.
func (m *MfaDetail) GetAuthMethod()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authMethod
    }
}
// The deserialization information for the current model
func (m *MfaDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthDetail(val)
        return nil
    }
    res["authMethod"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthMethod(val)
        return nil
    }
    return res
}
func (m *MfaDetail) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *MfaDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the authDetail property value. Indicates the MFA auth detail for the corresponding Sign-in activity when the MFA Required is 'Yes'.
// Parameters:
//  - value : Value to set for the authDetail property.
func (m *MfaDetail) SetAuthDetail(value *string)() {
    m.authDetail = value
}
// Sets the authMethod property value. Indicates the MFA Auth methods (SMS, Phone, Authenticator App are some of the value) for the corresponding sign-in activity when the MFA Required field is 'Yes'.
// Parameters:
//  - value : Value to set for the authMethod property.
func (m *MfaDetail) SetAuthMethod(value *string)() {
    m.authMethod = value
}
