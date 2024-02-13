package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualMachineDetails struct {
    Entity
}
// NewVirtualMachineDetails instantiates a new VirtualMachineDetails and sets the default values.
func NewVirtualMachineDetails()(*VirtualMachineDetails) {
    m := &VirtualMachineDetails{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualMachineDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualMachineDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualMachineDetails(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualMachineDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["virtualMachine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthorizationSystemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualMachine(val.(AuthorizationSystemResourceable))
        }
        return nil
    }
    return res
}
// GetVirtualMachine gets the virtualMachine property value. The virtualMachine property
// returns a AuthorizationSystemResourceable when successful
func (m *VirtualMachineDetails) GetVirtualMachine()(AuthorizationSystemResourceable) {
    val, err := m.GetBackingStore().Get("virtualMachine")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AuthorizationSystemResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualMachineDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("virtualMachine", m.GetVirtualMachine())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetVirtualMachine sets the virtualMachine property value. The virtualMachine property
func (m *VirtualMachineDetails) SetVirtualMachine(value AuthorizationSystemResourceable)() {
    err := m.GetBackingStore().Set("virtualMachine", value)
    if err != nil {
        panic(err)
    }
}
type VirtualMachineDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetVirtualMachine()(AuthorizationSystemResourceable)
    SetVirtualMachine(value AuthorizationSystemResourceable)()
}
