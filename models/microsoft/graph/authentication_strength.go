package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuthenticationStrength struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    authenticationStrengthId *string;
    // 
    displayName *string;
}
// Instantiates a new authenticationStrength and sets the default values.
func NewAuthenticationStrength()(*AuthenticationStrength) {
    m := &AuthenticationStrength{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationStrength) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the authenticationStrengthId property value. 
func (m *AuthenticationStrength) GetAuthenticationStrengthId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStrengthId
    }
}
// Gets the displayName property value. 
func (m *AuthenticationStrength) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// The deserialization information for the current model
func (m *AuthenticationStrength) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authenticationStrengthId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAuthenticationStrengthId(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    return res
}
func (m *AuthenticationStrength) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AuthenticationStrength) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("authenticationStrengthId", m.GetAuthenticationStrengthId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
func (m *AuthenticationStrength) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the authenticationStrengthId property value. 
// Parameters:
//  - value : Value to set for the authenticationStrengthId property.
func (m *AuthenticationStrength) SetAuthenticationStrengthId(value *string)() {
    m.authenticationStrengthId = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AuthenticationStrength) SetDisplayName(value *string)() {
    m.displayName = value
}
