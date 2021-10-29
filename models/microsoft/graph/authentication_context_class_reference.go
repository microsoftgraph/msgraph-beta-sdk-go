package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuthenticationContextClassReference struct {
    Entity
    // A short explanation of the policies that are enforced by authenticationContextClassReference. This value should be used to provide secondary text to describe the authentication context class reference when building user facing admin experiences. For example, selection UX.
    description *string;
    // The display name is the friendly name of the authenticationContextClassReference. This value should be used to identify the authentication context class reference when building user facing admin experiences. For example, selection UX.
    displayName *string;
    // Indicates whether the authenticationContextClassReference has been published by the security admin and is ready for use by apps. When it is set to false it should not be shown in admin UX experiences because the value is not currently available for selection.
    isAvailable *bool;
}
// Instantiates a new authenticationContextClassReference and sets the default values.
func NewAuthenticationContextClassReference()(*AuthenticationContextClassReference) {
    m := &AuthenticationContextClassReference{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. A short explanation of the policies that are enforced by authenticationContextClassReference. This value should be used to provide secondary text to describe the authentication context class reference when building user facing admin experiences. For example, selection UX.
func (m *AuthenticationContextClassReference) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The display name is the friendly name of the authenticationContextClassReference. This value should be used to identify the authentication context class reference when building user facing admin experiences. For example, selection UX.
func (m *AuthenticationContextClassReference) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isAvailable property value. Indicates whether the authenticationContextClassReference has been published by the security admin and is ready for use by apps. When it is set to false it should not be shown in admin UX experiences because the value is not currently available for selection.
func (m *AuthenticationContextClassReference) GetIsAvailable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAvailable
    }
}
// The deserialization information for the current model
func (m *AuthenticationContextClassReference) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
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
    res["isAvailable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAvailable(val)
        return nil
    }
    return res
}
func (m *AuthenticationContextClassReference) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AuthenticationContextClassReference) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAvailable", m.GetIsAvailable())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. A short explanation of the policies that are enforced by authenticationContextClassReference. This value should be used to provide secondary text to describe the authentication context class reference when building user facing admin experiences. For example, selection UX.
// Parameters:
//  - value : Value to set for the description property.
func (m *AuthenticationContextClassReference) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The display name is the friendly name of the authenticationContextClassReference. This value should be used to identify the authentication context class reference when building user facing admin experiences. For example, selection UX.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AuthenticationContextClassReference) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isAvailable property value. Indicates whether the authenticationContextClassReference has been published by the security admin and is ready for use by apps. When it is set to false it should not be shown in admin UX experiences because the value is not currently available for selection.
// Parameters:
//  - value : Value to set for the isAvailable property.
func (m *AuthenticationContextClassReference) SetIsAvailable(value *bool)() {
    m.isAvailable = value
}
