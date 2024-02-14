package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TeamworkSoftwareUpdateStatus struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkSoftwareUpdateStatus instantiates a new TeamworkSoftwareUpdateStatus and sets the default values.
func NewTeamworkSoftwareUpdateStatus()(*TeamworkSoftwareUpdateStatus) {
    m := &TeamworkSoftwareUpdateStatus{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamworkSoftwareUpdateStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkSoftwareUpdateStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamworkSoftwareUpdateStatus) GetAdditionalData()(map[string]any) {
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
// GetAvailableVersion gets the availableVersion property value. The available software version to update.
// returns a *string when successful
func (m *TeamworkSoftwareUpdateStatus) GetAvailableVersion()(*string) {
    val, err := m.GetBackingStore().Get("availableVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *TeamworkSoftwareUpdateStatus) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCurrentVersion gets the currentVersion property value. The current software version.
// returns a *string when successful
func (m *TeamworkSoftwareUpdateStatus) GetCurrentVersion()(*string) {
    val, err := m.GetBackingStore().Get("currentVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamworkSoftwareUpdateStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availableVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailableVersion(val)
        }
        return nil
    }
    res["currentVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentVersion(val)
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
    res["softwareFreshness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkSoftwareFreshness)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareFreshness(val.(*TeamworkSoftwareFreshness))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TeamworkSoftwareUpdateStatus) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSoftwareFreshness gets the softwareFreshness property value. The update status of the software. The possible values are: unknown, latest, updateAvailable, unknownFutureValue.
// returns a *TeamworkSoftwareFreshness when successful
func (m *TeamworkSoftwareUpdateStatus) GetSoftwareFreshness()(*TeamworkSoftwareFreshness) {
    val, err := m.GetBackingStore().Get("softwareFreshness")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TeamworkSoftwareFreshness)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkSoftwareUpdateStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("availableVersion", m.GetAvailableVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("currentVersion", m.GetCurrentVersion())
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
    if m.GetSoftwareFreshness() != nil {
        cast := (*m.GetSoftwareFreshness()).String()
        err := writer.WriteStringValue("softwareFreshness", &cast)
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
func (m *TeamworkSoftwareUpdateStatus) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAvailableVersion sets the availableVersion property value. The available software version to update.
func (m *TeamworkSoftwareUpdateStatus) SetAvailableVersion(value *string)() {
    err := m.GetBackingStore().Set("availableVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeamworkSoftwareUpdateStatus) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCurrentVersion sets the currentVersion property value. The current software version.
func (m *TeamworkSoftwareUpdateStatus) SetCurrentVersion(value *string)() {
    err := m.GetBackingStore().Set("currentVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkSoftwareUpdateStatus) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSoftwareFreshness sets the softwareFreshness property value. The update status of the software. The possible values are: unknown, latest, updateAvailable, unknownFutureValue.
func (m *TeamworkSoftwareUpdateStatus) SetSoftwareFreshness(value *TeamworkSoftwareFreshness)() {
    err := m.GetBackingStore().Set("softwareFreshness", value)
    if err != nil {
        panic(err)
    }
}
type TeamworkSoftwareUpdateStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableVersion()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCurrentVersion()(*string)
    GetOdataType()(*string)
    GetSoftwareFreshness()(*TeamworkSoftwareFreshness)
    SetAvailableVersion(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCurrentVersion(value *string)()
    SetOdataType(value *string)()
    SetSoftwareFreshness(value *TeamworkSoftwareFreshness)()
}
