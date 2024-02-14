package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AwsUser struct {
    AwsIdentity
}
// NewAwsUser instantiates a new AwsUser and sets the default values.
func NewAwsUser()(*AwsUser) {
    m := &AwsUser{
        AwsIdentity: *NewAwsIdentity(),
    }
    odataTypeValue := "#microsoft.graph.awsUser"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAwsUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAwsUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAwsUser(), nil
}
// GetAssumableRoles gets the assumableRoles property value. Roles assumed by the user.
// returns a []AwsRoleable when successful
func (m *AwsUser) GetAssumableRoles()([]AwsRoleable) {
    val, err := m.GetBackingStore().Get("assumableRoles")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AwsRoleable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AwsUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AwsIdentity.GetFieldDeserializers()
    res["assumableRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAwsRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AwsRoleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AwsRoleable)
                }
            }
            m.SetAssumableRoles(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AwsUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AwsIdentity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssumableRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssumableRoles()))
        for i, v := range m.GetAssumableRoles() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assumableRoles", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssumableRoles sets the assumableRoles property value. Roles assumed by the user.
func (m *AwsUser) SetAssumableRoles(value []AwsRoleable)() {
    err := m.GetBackingStore().Set("assumableRoles", value)
    if err != nil {
        panic(err)
    }
}
type AwsUserable interface {
    AwsIdentityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssumableRoles()([]AwsRoleable)
    SetAssumableRoles(value []AwsRoleable)()
}
