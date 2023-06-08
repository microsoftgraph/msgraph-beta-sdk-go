package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventPresenter 
type VirtualEventPresenter struct {
    Entity
}
// NewVirtualEventPresenter instantiates a new virtualEventPresenter and sets the default values.
func NewVirtualEventPresenter()(*VirtualEventPresenter) {
    m := &VirtualEventPresenter{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEventPresenterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventPresenterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventPresenter(), nil
}
// GetEmail gets the email property value. Email address of the presenter.
func (m *VirtualEventPresenter) GetEmail()(*string) {
    val, err := m.GetBackingStore().Get("email")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEventPresenter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommunicationsUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(CommunicationsUserIdentityable))
        }
        return nil
    }
    res["presenterDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVirtualEventPresenterDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPresenterDetails(val.(VirtualEventPresenterDetailsable))
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. Identity information of the presenter.
func (m *VirtualEventPresenter) GetIdentity()(CommunicationsUserIdentityable) {
    val, err := m.GetBackingStore().Get("identity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CommunicationsUserIdentityable)
    }
    return nil
}
// GetPresenterDetails gets the presenterDetails property value. Other detail information of the presenter.
func (m *VirtualEventPresenter) GetPresenterDetails()(VirtualEventPresenterDetailsable) {
    val, err := m.GetBackingStore().Get("presenterDetails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(VirtualEventPresenterDetailsable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEventPresenter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("presenterDetails", m.GetPresenterDetails())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmail sets the email property value. Email address of the presenter.
func (m *VirtualEventPresenter) SetEmail(value *string)() {
    err := m.GetBackingStore().Set("email", value)
    if err != nil {
        panic(err)
    }
}
// SetIdentity sets the identity property value. Identity information of the presenter.
func (m *VirtualEventPresenter) SetIdentity(value CommunicationsUserIdentityable)() {
    err := m.GetBackingStore().Set("identity", value)
    if err != nil {
        panic(err)
    }
}
// SetPresenterDetails sets the presenterDetails property value. Other detail information of the presenter.
func (m *VirtualEventPresenter) SetPresenterDetails(value VirtualEventPresenterDetailsable)() {
    err := m.GetBackingStore().Set("presenterDetails", value)
    if err != nil {
        panic(err)
    }
}
// VirtualEventPresenterable 
type VirtualEventPresenterable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetIdentity()(CommunicationsUserIdentityable)
    GetPresenterDetails()(VirtualEventPresenterDetailsable)
    SetEmail(value *string)()
    SetIdentity(value CommunicationsUserIdentityable)()
    SetPresenterDetails(value VirtualEventPresenterDetailsable)()
}
