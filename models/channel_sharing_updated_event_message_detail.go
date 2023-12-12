package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelSharingUpdatedEventMessageDetail 
type ChannelSharingUpdatedEventMessageDetail struct {
    EventMessageDetail
}
// NewChannelSharingUpdatedEventMessageDetail instantiates a new channelSharingUpdatedEventMessageDetail and sets the default values.
func NewChannelSharingUpdatedEventMessageDetail()(*ChannelSharingUpdatedEventMessageDetail) {
    m := &ChannelSharingUpdatedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.channelSharingUpdatedEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateChannelSharingUpdatedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelSharingUpdatedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannelSharingUpdatedEventMessageDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChannelSharingUpdatedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
    res["initiator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiator(val.(IdentitySetable))
        }
        return nil
    }
    res["ownerTeamId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerTeamId(val)
        }
        return nil
    }
    res["ownerTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerTenantId(val)
        }
        return nil
    }
    res["sharedChannelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedChannelId(val)
        }
        return nil
    }
    return res
}
// GetInitiator gets the initiator property value. Initiator of the event.
func (m *ChannelSharingUpdatedEventMessageDetail) GetInitiator()(IdentitySetable) {
    val, err := m.GetBackingStore().Get("initiator")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(IdentitySetable)
    }
    return nil
}
// GetOwnerTeamId gets the ownerTeamId property value. The ID of the team to which the shared channel belongs.
func (m *ChannelSharingUpdatedEventMessageDetail) GetOwnerTeamId()(*string) {
    val, err := m.GetBackingStore().Get("ownerTeamId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOwnerTenantId gets the ownerTenantId property value. The ID of the tenant to which the shared channel belongs.
func (m *ChannelSharingUpdatedEventMessageDetail) GetOwnerTenantId()(*string) {
    val, err := m.GetBackingStore().Get("ownerTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSharedChannelId gets the sharedChannelId property value. The ID of the shared channel.
func (m *ChannelSharingUpdatedEventMessageDetail) GetSharedChannelId()(*string) {
    val, err := m.GetBackingStore().Get("sharedChannelId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ChannelSharingUpdatedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("initiator", m.GetInitiator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerTeamId", m.GetOwnerTeamId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerTenantId", m.GetOwnerTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sharedChannelId", m.GetSharedChannelId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *ChannelSharingUpdatedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    err := m.GetBackingStore().Set("initiator", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerTeamId sets the ownerTeamId property value. The ID of the team to which the shared channel belongs.
func (m *ChannelSharingUpdatedEventMessageDetail) SetOwnerTeamId(value *string)() {
    err := m.GetBackingStore().Set("ownerTeamId", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerTenantId sets the ownerTenantId property value. The ID of the tenant to which the shared channel belongs.
func (m *ChannelSharingUpdatedEventMessageDetail) SetOwnerTenantId(value *string)() {
    err := m.GetBackingStore().Set("ownerTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetSharedChannelId sets the sharedChannelId property value. The ID of the shared channel.
func (m *ChannelSharingUpdatedEventMessageDetail) SetSharedChannelId(value *string)() {
    err := m.GetBackingStore().Set("sharedChannelId", value)
    if err != nil {
        panic(err)
    }
}
// ChannelSharingUpdatedEventMessageDetailable 
type ChannelSharingUpdatedEventMessageDetailable interface {
    EventMessageDetailable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInitiator()(IdentitySetable)
    GetOwnerTeamId()(*string)
    GetOwnerTenantId()(*string)
    GetSharedChannelId()(*string)
    SetInitiator(value IdentitySetable)()
    SetOwnerTeamId(value *string)()
    SetOwnerTenantId(value *string)()
    SetSharedChannelId(value *string)()
}
