package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOsVppAppAssignedLicense macOS Volume Purchase Program license assignment. This class does not support Create, Delete, or Update.
type MacOsVppAppAssignedLicense struct {
    Entity
    // The user email address.
    userEmailAddress *string
    // The user ID.
    userId *string
    // The user name.
    userName *string
    // The user principal name.
    userPrincipalName *string
}
// NewMacOsVppAppAssignedLicense instantiates a new macOsVppAppAssignedLicense and sets the default values.
func NewMacOsVppAppAssignedLicense()(*MacOsVppAppAssignedLicense) {
    m := &MacOsVppAppAssignedLicense{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMacOsVppAppAssignedLicenseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOsVppAppAssignedLicenseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOsVppAppAssignedLicense(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOsVppAppAssignedLicense) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["userEmailAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserEmailAddress)
    res["userId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserId)
    res["userName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserName)
    res["userPrincipalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetUserPrincipalName)
    return res
}
// GetUserEmailAddress gets the userEmailAddress property value. The user email address.
func (m *MacOsVppAppAssignedLicense) GetUserEmailAddress()(*string) {
    return m.userEmailAddress
}
// GetUserId gets the userId property value. The user ID.
func (m *MacOsVppAppAssignedLicense) GetUserId()(*string) {
    return m.userId
}
// GetUserName gets the userName property value. The user name.
func (m *MacOsVppAppAssignedLicense) GetUserName()(*string) {
    return m.userName
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name.
func (m *MacOsVppAppAssignedLicense) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *MacOsVppAppAssignedLicense) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userEmailAddress", m.GetUserEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserEmailAddress sets the userEmailAddress property value. The user email address.
func (m *MacOsVppAppAssignedLicense) SetUserEmailAddress(value *string)() {
    m.userEmailAddress = value
}
// SetUserId sets the userId property value. The user ID.
func (m *MacOsVppAppAssignedLicense) SetUserId(value *string)() {
    m.userId = value
}
// SetUserName sets the userName property value. The user name.
func (m *MacOsVppAppAssignedLicense) SetUserName(value *string)() {
    m.userName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name.
func (m *MacOsVppAppAssignedLicense) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
