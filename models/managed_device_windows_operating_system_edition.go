package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ManagedDeviceWindowsOperatingSystemEdition different Windows' versions have Edition specific support timelines. This complex type defines the date until which a particular edition is supported in a Windows' version.
type ManagedDeviceWindowsOperatingSystemEdition struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewManagedDeviceWindowsOperatingSystemEdition instantiates a new ManagedDeviceWindowsOperatingSystemEdition and sets the default values.
func NewManagedDeviceWindowsOperatingSystemEdition()(*ManagedDeviceWindowsOperatingSystemEdition) {
    m := &ManagedDeviceWindowsOperatingSystemEdition{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedDeviceWindowsOperatingSystemEditionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateManagedDeviceWindowsOperatingSystemEditionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceWindowsOperatingSystemEdition(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ManagedDeviceWindowsOperatingSystemEdition) GetAdditionalData()(map[string]any) {
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
func (m *ManagedDeviceWindowsOperatingSystemEdition) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEditionType gets the editionType property value. Windows Operating System is available in different editions, which have a specific set of features available. This enum type defines the corresponding edition.
// returns a *ManagedDeviceWindowsOperatingSystemEditionType when successful
func (m *ManagedDeviceWindowsOperatingSystemEdition) GetEditionType()(*ManagedDeviceWindowsOperatingSystemEditionType) {
    val, err := m.GetBackingStore().Get("editionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ManagedDeviceWindowsOperatingSystemEditionType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ManagedDeviceWindowsOperatingSystemEdition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["editionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedDeviceWindowsOperatingSystemEditionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEditionType(val.(*ManagedDeviceWindowsOperatingSystemEditionType))
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
    res["supportEndDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportEndDate(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *ManagedDeviceWindowsOperatingSystemEdition) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSupportEndDate gets the supportEndDate property value. Indicates the Date until which this Operating System edition type is officially supported. The Timestamp type represents date and time information using ISO 8601 format and is always in Pacific Time Zone (PT). For example, 2014-01-01 would mean '2014-01-01T07:00:00Z' in UTC time. Returned by default. Read-only.
// returns a *DateOnly when successful
func (m *ManagedDeviceWindowsOperatingSystemEdition) GetSupportEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("supportEndDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceWindowsOperatingSystemEdition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEditionType() != nil {
        cast := (*m.GetEditionType()).String()
        err := writer.WriteStringValue("editionType", &cast)
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
        err := writer.WriteDateOnlyValue("supportEndDate", m.GetSupportEndDate())
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
func (m *ManagedDeviceWindowsOperatingSystemEdition) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ManagedDeviceWindowsOperatingSystemEdition) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEditionType sets the editionType property value. Windows Operating System is available in different editions, which have a specific set of features available. This enum type defines the corresponding edition.
func (m *ManagedDeviceWindowsOperatingSystemEdition) SetEditionType(value *ManagedDeviceWindowsOperatingSystemEditionType)() {
    err := m.GetBackingStore().Set("editionType", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagedDeviceWindowsOperatingSystemEdition) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportEndDate sets the supportEndDate property value. Indicates the Date until which this Operating System edition type is officially supported. The Timestamp type represents date and time information using ISO 8601 format and is always in Pacific Time Zone (PT). For example, 2014-01-01 would mean '2014-01-01T07:00:00Z' in UTC time. Returned by default. Read-only.
func (m *ManagedDeviceWindowsOperatingSystemEdition) SetSupportEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("supportEndDate", value)
    if err != nil {
        panic(err)
    }
}
type ManagedDeviceWindowsOperatingSystemEditionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEditionType()(*ManagedDeviceWindowsOperatingSystemEditionType)
    GetOdataType()(*string)
    GetSupportEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEditionType(value *ManagedDeviceWindowsOperatingSystemEditionType)()
    SetOdataType(value *string)()
    SetSupportEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
