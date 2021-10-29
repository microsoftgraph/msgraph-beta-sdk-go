package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AttackSimulationUser struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Display name of the user.
    displayName *string;
    // Email address of the user.
    email *string;
    // This is the id property value of the user resource that represents the user in the Azure AD tenant.
    userId *string;
}
// Instantiates a new attackSimulationUser and sets the default values.
func NewAttackSimulationUser()(*AttackSimulationUser) {
    m := &AttackSimulationUser{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttackSimulationUser) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. Display name of the user.
func (m *AttackSimulationUser) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the email property value. Email address of the user.
func (m *AttackSimulationUser) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the userId property value. This is the id property value of the user resource that represents the user in the Azure AD tenant.
func (m *AttackSimulationUser) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// The deserialization information for the current model
func (m *AttackSimulationUser) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEmail(val)
        return nil
    }
    res["userId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserId(val)
        return nil
    }
    return res
}
func (m *AttackSimulationUser) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AttackSimulationUser) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
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
func (m *AttackSimulationUser) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. Display name of the user.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AttackSimulationUser) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the email property value. Email address of the user.
// Parameters:
//  - value : Value to set for the email property.
func (m *AttackSimulationUser) SetEmail(value *string)() {
    m.email = value
}
// Sets the userId property value. This is the id property value of the user resource that represents the user in the Azure AD tenant.
// Parameters:
//  - value : Value to set for the userId property.
func (m *AttackSimulationUser) SetUserId(value *string)() {
    m.userId = value
}
