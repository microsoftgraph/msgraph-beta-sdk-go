package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationContextClassReference provides operations to manage the collection of activityStatistics entities.
type AuthenticationContextClassReference struct {
    Entity
    // A short explanation of the policies that are enforced by authenticationContextClassReference. This value should be used to provide secondary text to describe the authentication context class reference when building user facing admin experiences. For example, selection UX.
    description *string
    // A friendly name that identifies the authenticationContextClassReference object when building user-facing admin experiences. For example, a selection UX.
    displayName *string
    // Indicates whether the authenticationContextClassReference has been published by the security admin and is ready for use by apps. When it is set to false, it should not be shown in selection UX used to tag resources with authentication context class values. It will still be shown in the Conditionall Access policy authoring experience.  Supports $filter (eq).
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
    return m.description
}
// GetDisplayName gets the displayName property value. A friendly name that identifies the authenticationContextClassReference object when building user-facing admin experiences. For example, a selection UX.
func (m *AuthenticationContextClassReference) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationContextClassReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["isAvailable"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAvailable)
    return res
}
// GetIsAvailable gets the isAvailable property value. Indicates whether the authenticationContextClassReference has been published by the security admin and is ready for use by apps. When it is set to false, it should not be shown in selection UX used to tag resources with authentication context class values. It will still be shown in the Conditionall Access policy authoring experience.  Supports $filter (eq).
func (m *AuthenticationContextClassReference) GetIsAvailable()(*bool) {
    return m.isAvailable
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
    m.description = value
}
// SetDisplayName sets the displayName property value. A friendly name that identifies the authenticationContextClassReference object when building user-facing admin experiences. For example, a selection UX.
func (m *AuthenticationContextClassReference) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsAvailable sets the isAvailable property value. Indicates whether the authenticationContextClassReference has been published by the security admin and is ready for use by apps. When it is set to false, it should not be shown in selection UX used to tag resources with authentication context class values. It will still be shown in the Conditionall Access policy authoring experience.  Supports $filter (eq).
func (m *AuthenticationContextClassReference) SetIsAvailable(value *bool)() {
    m.isAvailable = value
}
