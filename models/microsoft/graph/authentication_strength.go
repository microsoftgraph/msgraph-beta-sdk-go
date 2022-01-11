package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuthenticationStrength 
type AuthenticationStrength struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Identifier of the authentication strength.
    authenticationStrengthId *string;
    // The name of the authentication strength.
    displayName *string;
}
// NewAuthenticationStrength instantiates a new authenticationStrength and sets the default values.
func NewAuthenticationStrength()(*AuthenticationStrength) {
    m := &AuthenticationStrength{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationStrength) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAuthenticationStrengthId gets the authenticationStrengthId property value. Identifier of the authentication strength.
func (m *AuthenticationStrength) GetAuthenticationStrengthId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.authenticationStrengthId
    }
}
// GetDisplayName gets the displayName property value. The name of the authentication strength.
func (m *AuthenticationStrength) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationStrength) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["authenticationStrengthId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationStrengthId(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
func (m *AuthenticationStrength) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationStrength) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAuthenticationStrengthId sets the authenticationStrengthId property value. Identifier of the authentication strength.
func (m *AuthenticationStrength) SetAuthenticationStrengthId(value *string)() {
    if m != nil {
        m.authenticationStrengthId = value
    }
}
// SetDisplayName sets the displayName property value. The name of the authentication strength.
func (m *AuthenticationStrength) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
