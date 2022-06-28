package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CredentialUserRegistrationCount 
type CredentialUserRegistrationCount struct {
    Entity
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
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
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCredentialUserRegistrationCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCredentialUserRegistrationCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCredentialUserRegistrationCount(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CredentialUserRegistrationCount) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CredentialUserRegistrationCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["totalUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUserCount(val)
        }
        return nil
    }
    res["userRegistrationCounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserRegistrationCountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserRegistrationCountable, len(val))
            for i, v := range val {
                res[i] = v.(UserRegistrationCountable)
            }
            m.SetUserRegistrationCounts(res)
        }
        return nil
    }
    return res
}
// GetTotalUserCount gets the totalUserCount property value. Provides the total user count in the tenant.
func (m *CredentialUserRegistrationCount) GetTotalUserCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.totalUserCount
    }
}
// GetUserRegistrationCounts gets the userRegistrationCounts property value. A collection of registration count and status information for users in your tenant.
func (m *CredentialUserRegistrationCount) GetUserRegistrationCounts()([]UserRegistrationCountable) {
    if m == nil {
        return nil
    } else {
        return m.userRegistrationCounts
    }
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserRegistrationCounts()))
        for i, v := range m.GetUserRegistrationCounts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userRegistrationCounts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CredentialUserRegistrationCount) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTotalUserCount sets the totalUserCount property value. Provides the total user count in the tenant.
func (m *CredentialUserRegistrationCount) SetTotalUserCount(value *int64)() {
    if m != nil {
        m.totalUserCount = value
    }
}
// SetUserRegistrationCounts sets the userRegistrationCounts property value. A collection of registration count and status information for users in your tenant.
func (m *CredentialUserRegistrationCount) SetUserRegistrationCounts(value []UserRegistrationCountable)() {
    if m != nil {
        m.userRegistrationCounts = value
    }
}
