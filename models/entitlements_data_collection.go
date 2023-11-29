package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EntitlementsDataCollection 
type EntitlementsDataCollection struct {
    EntitlementsDataCollectionInfo
}
// NewEntitlementsDataCollection instantiates a new entitlementsDataCollection and sets the default values.
func NewEntitlementsDataCollection()(*EntitlementsDataCollection) {
    m := &EntitlementsDataCollection{
        EntitlementsDataCollectionInfo: *NewEntitlementsDataCollectionInfo(),
    }
    odataTypeValue := "#microsoft.graph.entitlementsDataCollection"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEntitlementsDataCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementsDataCollectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEntitlementsDataCollection(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementsDataCollection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EntitlementsDataCollectionInfo.GetFieldDeserializers()
    res["lastCollectionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCollectionDateTime(val)
        }
        return nil
    }
    res["permissionsModificationCapability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePermissionsModificationCapability)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionsModificationCapability(val.(*PermissionsModificationCapability))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataCollectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DataCollectionStatus))
        }
        return nil
    }
    return res
}
// GetLastCollectionDateTime gets the lastCollectionDateTime property value. Last transformation time of entitlements.
func (m *EntitlementsDataCollection) GetLastCollectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastCollectionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetPermissionsModificationCapability gets the permissionsModificationCapability property value. The permissionsModificationCapability property
func (m *EntitlementsDataCollection) GetPermissionsModificationCapability()(*PermissionsModificationCapability) {
    val, err := m.GetBackingStore().Get("permissionsModificationCapability")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PermissionsModificationCapability)
    }
    return nil
}
// GetStatus gets the status property value. The status property
func (m *EntitlementsDataCollection) GetStatus()(*DataCollectionStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DataCollectionStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EntitlementsDataCollection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EntitlementsDataCollectionInfo.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastCollectionDateTime", m.GetLastCollectionDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionsModificationCapability() != nil {
        cast := (*m.GetPermissionsModificationCapability()).String()
        err = writer.WriteStringValue("permissionsModificationCapability", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastCollectionDateTime sets the lastCollectionDateTime property value. Last transformation time of entitlements.
func (m *EntitlementsDataCollection) SetLastCollectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastCollectionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetPermissionsModificationCapability sets the permissionsModificationCapability property value. The permissionsModificationCapability property
func (m *EntitlementsDataCollection) SetPermissionsModificationCapability(value *PermissionsModificationCapability)() {
    err := m.GetBackingStore().Set("permissionsModificationCapability", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The status property
func (m *EntitlementsDataCollection) SetStatus(value *DataCollectionStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// EntitlementsDataCollectionable 
type EntitlementsDataCollectionable interface {
    EntitlementsDataCollectionInfoable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLastCollectionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPermissionsModificationCapability()(*PermissionsModificationCapability)
    GetStatus()(*DataCollectionStatus)
    SetLastCollectionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPermissionsModificationCapability(value *PermissionsModificationCapability)()
    SetStatus(value *DataCollectionStatus)()
}
