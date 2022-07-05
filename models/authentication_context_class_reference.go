package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationContextClassReference provides operations to manage the collection of accessReview entities.
type AuthenticationContextClassReference struct {
    Entity
    // A short explanation of the policies that are enforced by authenticationContextClassReference. This value should be used to provide secondary text to describe the authentication context class reference when building user facing admin experiences. For example, selection UX.
    description *string
    // The display name is the friendly name of the authenticationContextClassReference. This value should be used to identify the authentication context class reference when building user facing admin experiences. For example, selection UX.
    displayName *string
    // Indicates whether the authenticationContextClassReference has been published by the security admin and is ready for use by apps. When it is set to false it should not be shown in admin UX experiences because the value is not currently available for selection.
    isAvailable *bool
}
// NewAuthenticationContextClassReference instantiates a new authenticationContextClassReference and sets the default values.
func NewAuthenticationContextClassReference()(*AuthenticationContextClassReference) {
    m := &AuthenticationContextClassReference{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationContextClassReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationContextClassReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationContextClassReference(), nil
}
// GetDescription gets the description property value. A short explanation of the policies that are enforced by authenticationContextClassReference. This value should be used to provide secondary text to describe the authentication context class reference when building user facing admin experiences. For example, selection UX.
func (m *AuthenticationContextClassReference) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. The display name is the friendly name of the authenticationContextClassReference. This value should be used to identify the authentication context class reference when building user facing admin experiences. For example, selection UX.
func (m *AuthenticationContextClassReference) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationContextClassReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["isAvailable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAvailable(val)
        }
        return nil
    }
    return res
}
// GetIsAvailable gets the isAvailable property value. Indicates whether the authenticationContextClassReference has been published by the security admin and is ready for use by apps. When it is set to false it should not be shown in admin UX experiences because the value is not currently available for selection.
func (m *AuthenticationContextClassReference) GetIsAvailable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAvailable
    }
}
// Serialize serializes information the current object
func (m *AuthenticationContextClassReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetDescription sets the description property value. A short explanation of the policies that are enforced by authenticationContextClassReference. This value should be used to provide secondary text to describe the authentication context class reference when building user facing admin experiences. For example, selection UX.
func (m *AuthenticationContextClassReference) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. The display name is the friendly name of the authenticationContextClassReference. This value should be used to identify the authentication context class reference when building user facing admin experiences. For example, selection UX.
func (m *AuthenticationContextClassReference) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIsAvailable sets the isAvailable property value. Indicates whether the authenticationContextClassReference has been published by the security admin and is ready for use by apps. When it is set to false it should not be shown in admin UX experiences because the value is not currently available for selection.
func (m *AuthenticationContextClassReference) SetIsAvailable(value *bool)() {
    if m != nil {
        m.isAvailable = value
    }
}
