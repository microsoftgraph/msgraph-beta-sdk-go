package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsRole 
type AwsRole struct {
    AwsIdentity
}
// NewAwsRole instantiates a new awsRole and sets the default values.
func NewAwsRole()(*AwsRole) {
    m := &AwsRole{
        AwsIdentity: *NewAwsIdentity(),
    }
    odataTypeValue := "#microsoft.graph.awsRole"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsRole(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsIdentity.GetFieldDeserializers()
    res["roleType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAwsRoleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleType(val.(*AwsRoleType))
        }
        return nil
    }
    res["trustEntityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAwsRoleTrustEntityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustEntityType(val.(*AwsRoleTrustEntityType))
        }
        return nil
    }
    return res
}
// GetRoleType gets the roleType property value. The roleType property
func (m *AwsRole) GetRoleType()(*AwsRoleType) {
    val, err := m.GetBackingStore().Get("roleType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AwsRoleType)
    }
    return nil
}
// GetTrustEntityType gets the trustEntityType property value. The trustEntityType property
func (m *AwsRole) GetTrustEntityType()(*AwsRoleTrustEntityType) {
    val, err := m.GetBackingStore().Get("trustEntityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AwsRoleTrustEntityType)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRoleType() != nil {
        cast := (*m.GetRoleType()).String()
        err = writer.WriteStringValue("roleType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTrustEntityType() != nil {
        cast := (*m.GetTrustEntityType()).String()
        err = writer.WriteStringValue("trustEntityType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRoleType sets the roleType property value. The roleType property
func (m *AwsRole) SetRoleType(value *AwsRoleType)() {
    err := m.GetBackingStore().Set("roleType", value)
    if err != nil {
        panic(err)
    }
}
// SetTrustEntityType sets the trustEntityType property value. The trustEntityType property
func (m *AwsRole) SetTrustEntityType(value *AwsRoleTrustEntityType)() {
    err := m.GetBackingStore().Set("trustEntityType", value)
    if err != nil {
        panic(err)
    }
}
// AwsRoleable 
type AwsRoleable interface {
    AwsIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRoleType()(*AwsRoleType)
    GetTrustEntityType()(*AwsRoleTrustEntityType)
    SetRoleType(value *AwsRoleType)()
    SetTrustEntityType(value *AwsRoleTrustEntityType)()
}
