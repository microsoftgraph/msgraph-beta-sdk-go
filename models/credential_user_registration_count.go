package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialUserRegistrationCount 
type CredentialUserRegistrationCount struct {
    Entity
    // Provides the total user count in the tenant.
    totalUserCount *int64
    // A collection of registration count and status information for users in your tenant.
    userRegistrationCounts []UserRegistrationCountable
}
// NewCredentialUserRegistrationCount instantiates a new CredentialUserRegistrationCount and sets the default values.
func NewCredentialUserRegistrationCount()(*CredentialUserRegistrationCount) {
    m := &CredentialUserRegistrationCount{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCredentialUserRegistrationCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCredentialUserRegistrationCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredentialUserRegistrationCount(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialUserRegistrationCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["totalUserCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetTotalUserCount)
    res["userRegistrationCounts"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateUserRegistrationCountFromDiscriminatorValue , m.SetUserRegistrationCounts)
    return res
}
// GetTotalUserCount gets the totalUserCount property value. Provides the total user count in the tenant.
func (m *CredentialUserRegistrationCount) GetTotalUserCount()(*int64) {
    return m.totalUserCount
}
// GetUserRegistrationCounts gets the userRegistrationCounts property value. A collection of registration count and status information for users in your tenant.
func (m *CredentialUserRegistrationCount) GetUserRegistrationCounts()([]UserRegistrationCountable) {
    return m.userRegistrationCounts
}
// Serialize serializes information the current object
func (m *CredentialUserRegistrationCount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("totalUserCount", m.GetTotalUserCount())
        if err != nil {
            return err
        }
    }
    if m.GetUserRegistrationCounts() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetUserRegistrationCounts())
        err = writer.WriteCollectionOfObjectValues("userRegistrationCounts", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTotalUserCount sets the totalUserCount property value. Provides the total user count in the tenant.
func (m *CredentialUserRegistrationCount) SetTotalUserCount(value *int64)() {
    m.totalUserCount = value
}
// SetUserRegistrationCounts sets the userRegistrationCounts property value. A collection of registration count and status information for users in your tenant.
func (m *CredentialUserRegistrationCount) SetUserRegistrationCounts(value []UserRegistrationCountable)() {
    m.userRegistrationCounts = value
}
