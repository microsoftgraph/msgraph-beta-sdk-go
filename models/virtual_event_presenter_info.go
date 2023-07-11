package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventPresenterInfo 
type VirtualEventPresenterInfo struct {
    MeetingParticipantInfo
}
// NewVirtualEventPresenterInfo instantiates a new virtualEventPresenterInfo and sets the default values.
func NewVirtualEventPresenterInfo()(*VirtualEventPresenterInfo) {
    m := &VirtualEventPresenterInfo{
        MeetingParticipantInfo: *NewMeetingParticipantInfo(),
    }
    odataTypeValue := "#microsoft.graph.virtualEventPresenterInfo"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateVirtualEventPresenterInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventPresenterInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventPresenterInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEventPresenterInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MeetingParticipantInfo.GetFieldDeserializers()
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VirtualEventPresenterInfo) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPresenterDetails gets the presenterDetails property value. The presenterDetails property
func (m *VirtualEventPresenterInfo) GetPresenterDetails()(VirtualEventPresenterDetailsable) {
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
func (m *VirtualEventPresenterInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MeetingParticipantInfo.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
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
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VirtualEventPresenterInfo) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPresenterDetails sets the presenterDetails property value. The presenterDetails property
func (m *VirtualEventPresenterInfo) SetPresenterDetails(value VirtualEventPresenterDetailsable)() {
    err := m.GetBackingStore().Set("presenterDetails", value)
    if err != nil {
        panic(err)
    }
}
// VirtualEventPresenterInfoable 
type VirtualEventPresenterInfoable interface {
    MeetingParticipantInfoable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetPresenterDetails()(VirtualEventPresenterDetailsable)
    SetOdataType(value *string)()
    SetPresenterDetails(value VirtualEventPresenterDetailsable)()
}
