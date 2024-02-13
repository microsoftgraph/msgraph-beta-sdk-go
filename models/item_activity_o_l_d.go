package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemActivityOLD struct {
    Entity
}
// NewItemActivityOLD instantiates a new ItemActivityOLD and sets the default values.
func NewItemActivityOLD()(*ItemActivityOLD) {
    m := &ItemActivityOLD{
        Entity: *NewEntity(),
    }
    return m
}
// CreateItemActivityOLDFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemActivityOLDFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemActivityOLD(), nil
}
// GetAction gets the action property value. The action property
// returns a ItemActionSetable when successful
func (m *ItemActivityOLD) GetAction()(ItemActionSetable) {
    val, err := m.GetBackingStore().Get("action")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemActionSetable)
    }
    return nil
}
// GetActor gets the actor property value. The actor property
// returns a IdentitySetable when successful
func (m *ItemActivityOLD) GetActor()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("actor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetDriveItem gets the driveItem property value. The driveItem property
// returns a DriveItemable when successful
func (m *ItemActivityOLD) GetDriveItem()(DriveItemable) {
    val, err := m.GetBackingStore().Get("driveItem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DriveItemable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemActivityOLD) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActionSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(ItemActionSetable))
        }
        return nil
    }
    res["actor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActor(val.(IdentitySetable))
        }
        return nil
    }
    res["driveItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveItem(val.(DriveItemable))
        }
        return nil
    }
    res["listItem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItem(val.(ListItemable))
        }
        return nil
    }
    res["times"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemActivityTimeSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimes(val.(ItemActivityTimeSetable))
        }
        return nil
    }
    return res
}
// GetListItem gets the listItem property value. The listItem property
// returns a ListItemable when successful
func (m *ItemActivityOLD) GetListItem()(ListItemable) {
    val, err := m.GetBackingStore().Get("listItem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ListItemable)
    }
    return nil
}
// GetTimes gets the times property value. The times property
// returns a ItemActivityTimeSetable when successful
func (m *ItemActivityOLD) GetTimes()(ItemActivityTimeSetable) {
    val, err := m.GetBackingStore().Get("times")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemActivityTimeSetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemActivityOLD) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("action", m.GetAction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("actor", m.GetActor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("driveItem", m.GetDriveItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("listItem", m.GetListItem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("times", m.GetTimes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. The action property
func (m *ItemActivityOLD) SetAction(value ItemActionSetable)() {
    err := m.GetBackingStore().Set("action", value)
    if err != nil {
        panic(err)
    }
}
// SetActor sets the actor property value. The actor property
func (m *ItemActivityOLD) SetActor(value IdentitySetable)() {
    err := m.GetBackingStore().Set("actor", value)
    if err != nil {
        panic(err)
    }
}
// SetDriveItem sets the driveItem property value. The driveItem property
func (m *ItemActivityOLD) SetDriveItem(value DriveItemable)() {
    err := m.GetBackingStore().Set("driveItem", value)
    if err != nil {
        panic(err)
    }
}
// SetListItem sets the listItem property value. The listItem property
func (m *ItemActivityOLD) SetListItem(value ListItemable)() {
    err := m.GetBackingStore().Set("listItem", value)
    if err != nil {
        panic(err)
    }
}
// SetTimes sets the times property value. The times property
func (m *ItemActivityOLD) SetTimes(value ItemActivityTimeSetable)() {
    err := m.GetBackingStore().Set("times", value)
    if err != nil {
        panic(err)
    }
}
type ItemActivityOLDable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAction()(ItemActionSetable)
    GetActor()(IdentitySetable)
    GetDriveItem()(DriveItemable)
    GetListItem()(ListItemable)
    GetTimes()(ItemActivityTimeSetable)
    SetAction(value ItemActionSetable)()
    SetActor(value IdentitySetable)()
    SetDriveItem(value DriveItemable)()
    SetListItem(value ListItemable)()
    SetTimes(value ItemActivityTimeSetable)()
}
