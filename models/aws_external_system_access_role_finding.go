package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AwsExternalSystemAccessRoleFinding 
type AwsExternalSystemAccessRoleFinding struct {
    Finding
}
// NewAwsExternalSystemAccessRoleFinding instantiates a new awsExternalSystemAccessRoleFinding and sets the default values.
func NewAwsExternalSystemAccessRoleFinding()(*AwsExternalSystemAccessRoleFinding) {
    m := &AwsExternalSystemAccessRoleFinding{
        Finding: *NewFinding(),
    }
    return m
}
// CreateAwsExternalSystemAccessRoleFindingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAwsExternalSystemAccessRoleFindingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsExternalSystemAccessRoleFinding(), nil
}
// GetAccessibleSystemIds gets the accessibleSystemIds property value. The accessibleSystemIds property
func (m *AwsExternalSystemAccessRoleFinding) GetAccessibleSystemIds()([]string) {
    val, err := m.GetBackingStore().Get("accessibleSystemIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AwsExternalSystemAccessRoleFinding) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Finding.GetFieldDeserializers()
    res["accessibleSystemIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAccessibleSystemIds(res)
        }
        return nil
    }
    res["permissionsCreepIndex"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePermissionsCreepIndexFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionsCreepIndex(val.(PermissionsCreepIndexable))
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAwsRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(AwsRoleable))
        }
        return nil
    }
    return res
}
// GetPermissionsCreepIndex gets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *AwsExternalSystemAccessRoleFinding) GetPermissionsCreepIndex()(PermissionsCreepIndexable) {
    val, err := m.GetBackingStore().Get("permissionsCreepIndex")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PermissionsCreepIndexable)
    }
    return nil
}
// GetRole gets the role property value. The role property
func (m *AwsExternalSystemAccessRoleFinding) GetRole()(AwsRoleable) {
    val, err := m.GetBackingStore().Get("role")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AwsRoleable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AwsExternalSystemAccessRoleFinding) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Finding.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessibleSystemIds() != nil {
        err = writer.WriteCollectionOfStringValues("accessibleSystemIds", m.GetAccessibleSystemIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("permissionsCreepIndex", m.GetPermissionsCreepIndex())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessibleSystemIds sets the accessibleSystemIds property value. The accessibleSystemIds property
func (m *AwsExternalSystemAccessRoleFinding) SetAccessibleSystemIds(value []string)() {
    err := m.GetBackingStore().Set("accessibleSystemIds", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsCreepIndex sets the permissionsCreepIndex property value. The permissionsCreepIndex property
func (m *AwsExternalSystemAccessRoleFinding) SetPermissionsCreepIndex(value PermissionsCreepIndexable)() {
    err := m.GetBackingStore().Set("permissionsCreepIndex", value)
    if err != nil {
        panic(err)
    }
}
// SetRole sets the role property value. The role property
func (m *AwsExternalSystemAccessRoleFinding) SetRole(value AwsRoleable)() {
    err := m.GetBackingStore().Set("role", value)
    if err != nil {
        panic(err)
    }
}
// AwsExternalSystemAccessRoleFindingable 
type AwsExternalSystemAccessRoleFindingable interface {
    Findingable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessibleSystemIds()([]string)
    GetPermissionsCreepIndex()(PermissionsCreepIndexable)
    GetRole()(AwsRoleable)
    SetAccessibleSystemIds(value []string)()
    SetPermissionsCreepIndex(value PermissionsCreepIndexable)()
    SetRole(value AwsRoleable)()
}
