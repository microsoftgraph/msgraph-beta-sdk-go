package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// WindowsDriverUpdateProfileInventorySyncStatus a complex type to store the status of a driver and firmware profile inventory sync. The status includes the last successful sync date time and the state of the last sync.
type WindowsDriverUpdateProfileInventorySyncStatus struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewWindowsDriverUpdateProfileInventorySyncStatus instantiates a new WindowsDriverUpdateProfileInventorySyncStatus and sets the default values.
func NewWindowsDriverUpdateProfileInventorySyncStatus()(*WindowsDriverUpdateProfileInventorySyncStatus) {
    m := &WindowsDriverUpdateProfileInventorySyncStatus{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsDriverUpdateProfileInventorySyncStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWindowsDriverUpdateProfileInventorySyncStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDriverUpdateProfileInventorySyncStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDriverInventorySyncState gets the driverInventorySyncState property value. Windows DnF update inventory sync state.
// returns a *WindowsDriverUpdateProfileInventorySyncState when successful
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetDriverInventorySyncState()(*WindowsDriverUpdateProfileInventorySyncState) {
    val, err := m.GetBackingStore().Get("driverInventorySyncState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*WindowsDriverUpdateProfileInventorySyncState)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["driverInventorySyncState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsDriverUpdateProfileInventorySyncState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriverInventorySyncState(val.(*WindowsDriverUpdateProfileInventorySyncState))
        }
        return nil
    }
    res["lastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
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
    return res
}
// GetLastSuccessfulSyncDateTime gets the lastSuccessfulSyncDateTime property value. The last successful sync date and time in UTC.
// returns a *Time when successful
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastSuccessfulSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *WindowsDriverUpdateProfileInventorySyncStatus) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *WindowsDriverUpdateProfileInventorySyncStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDriverInventorySyncState() != nil {
        cast := (*m.GetDriverInventorySyncState()).String()
        err := writer.WriteStringValue("driverInventorySyncState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastSuccessfulSyncDateTime", m.GetLastSuccessfulSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsDriverUpdateProfileInventorySyncStatus) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *WindowsDriverUpdateProfileInventorySyncStatus) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDriverInventorySyncState sets the driverInventorySyncState property value. Windows DnF update inventory sync state.
func (m *WindowsDriverUpdateProfileInventorySyncStatus) SetDriverInventorySyncState(value *WindowsDriverUpdateProfileInventorySyncState)() {
    err := m.GetBackingStore().Set("driverInventorySyncState", value)
    if err != nil {
        panic(err)
    }
}
// SetLastSuccessfulSyncDateTime sets the lastSuccessfulSyncDateTime property value. The last successful sync date and time in UTC.
func (m *WindowsDriverUpdateProfileInventorySyncStatus) SetLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastSuccessfulSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsDriverUpdateProfileInventorySyncStatus) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
type WindowsDriverUpdateProfileInventorySyncStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDriverInventorySyncState()(*WindowsDriverUpdateProfileInventorySyncState)
    GetLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDriverInventorySyncState(value *WindowsDriverUpdateProfileInventorySyncState)()
    SetLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOdataType(value *string)()
}
