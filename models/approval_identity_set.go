package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ApprovalIdentitySet struct {
    IdentitySet
}
// NewApprovalIdentitySet instantiates a new ApprovalIdentitySet and sets the default values.
func NewApprovalIdentitySet()(*ApprovalIdentitySet) {
    m := &ApprovalIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    odataTypeValue := "#microsoft.graph.approvalIdentitySet"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateApprovalIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateApprovalIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalIdentitySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ApprovalIdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Identityable))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The Microsoft Entra group associated with the approval item.
// returns a Identityable when successful
func (m *ApprovalIdentitySet) GetGroup()(Identityable) {
    val, err := m.GetBackingStore().Get("group")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Identityable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ApprovalIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroup sets the group property value. The Microsoft Entra group associated with the approval item.
func (m *ApprovalIdentitySet) SetGroup(value Identityable)() {
    err := m.GetBackingStore().Set("group", value)
    if err != nil {
        panic(err)
    }
}
type ApprovalIdentitySetable interface {
    IdentitySetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroup()(Identityable)
    SetGroup(value Identityable)()
}
