package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type TeamworkActivePeripherals struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkActivePeripherals instantiates a new TeamworkActivePeripherals and sets the default values.
func NewTeamworkActivePeripherals()(*TeamworkActivePeripherals) {
    m := &TeamworkActivePeripherals{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkActivePeripheralsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTeamworkActivePeripheralsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkActivePeripherals(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TeamworkActivePeripherals) GetAdditionalData()(map[string]any) {
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
func (m *TeamworkActivePeripherals) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCommunicationSpeaker gets the communicationSpeaker property value. The communicationSpeaker property
// returns a TeamworkPeripheralable when successful
func (m *TeamworkActivePeripherals) GetCommunicationSpeaker()(TeamworkPeripheralable) {
    val, err := m.GetBackingStore().Get("communicationSpeaker")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralable)
    }
    return nil
}
// GetContentCamera gets the contentCamera property value. The contentCamera property
// returns a TeamworkPeripheralable when successful
func (m *TeamworkActivePeripherals) GetContentCamera()(TeamworkPeripheralable) {
    val, err := m.GetBackingStore().Get("contentCamera")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TeamworkActivePeripherals) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["communicationSpeaker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommunicationSpeaker(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["contentCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCamera(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["microphone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophone(val.(TeamworkPeripheralable))
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
    res["roomCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomCamera(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["speaker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeaker(val.(TeamworkPeripheralable))
        }
        return nil
    }
    return res
}
// GetMicrophone gets the microphone property value. The microphone property
// returns a TeamworkPeripheralable when successful
func (m *TeamworkActivePeripherals) GetMicrophone()(TeamworkPeripheralable) {
    val, err := m.GetBackingStore().Get("microphone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *TeamworkActivePeripherals) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoomCamera gets the roomCamera property value. The roomCamera property
// returns a TeamworkPeripheralable when successful
func (m *TeamworkActivePeripherals) GetRoomCamera()(TeamworkPeripheralable) {
    val, err := m.GetBackingStore().Get("roomCamera")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralable)
    }
    return nil
}
// GetSpeaker gets the speaker property value. The speaker property
// returns a TeamworkPeripheralable when successful
func (m *TeamworkActivePeripherals) GetSpeaker()(TeamworkPeripheralable) {
    val, err := m.GetBackingStore().Get("speaker")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkActivePeripherals) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("communicationSpeaker", m.GetCommunicationSpeaker())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentCamera", m.GetContentCamera())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("microphone", m.GetMicrophone())
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
        err := writer.WriteObjectValue("roomCamera", m.GetRoomCamera())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("speaker", m.GetSpeaker())
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
func (m *TeamworkActivePeripherals) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeamworkActivePeripherals) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCommunicationSpeaker sets the communicationSpeaker property value. The communicationSpeaker property
func (m *TeamworkActivePeripherals) SetCommunicationSpeaker(value TeamworkPeripheralable)() {
    err := m.GetBackingStore().Set("communicationSpeaker", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCamera sets the contentCamera property value. The contentCamera property
func (m *TeamworkActivePeripherals) SetContentCamera(value TeamworkPeripheralable)() {
    err := m.GetBackingStore().Set("contentCamera", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrophone sets the microphone property value. The microphone property
func (m *TeamworkActivePeripherals) SetMicrophone(value TeamworkPeripheralable)() {
    err := m.GetBackingStore().Set("microphone", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkActivePeripherals) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRoomCamera sets the roomCamera property value. The roomCamera property
func (m *TeamworkActivePeripherals) SetRoomCamera(value TeamworkPeripheralable)() {
    err := m.GetBackingStore().Set("roomCamera", value)
    if err != nil {
        panic(err)
    }
}
// SetSpeaker sets the speaker property value. The speaker property
func (m *TeamworkActivePeripherals) SetSpeaker(value TeamworkPeripheralable)() {
    err := m.GetBackingStore().Set("speaker", value)
    if err != nil {
        panic(err)
    }
}
type TeamworkActivePeripheralsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCommunicationSpeaker()(TeamworkPeripheralable)
    GetContentCamera()(TeamworkPeripheralable)
    GetMicrophone()(TeamworkPeripheralable)
    GetOdataType()(*string)
    GetRoomCamera()(TeamworkPeripheralable)
    GetSpeaker()(TeamworkPeripheralable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCommunicationSpeaker(value TeamworkPeripheralable)()
    SetContentCamera(value TeamworkPeripheralable)()
    SetMicrophone(value TeamworkPeripheralable)()
    SetOdataType(value *string)()
    SetRoomCamera(value TeamworkPeripheralable)()
    SetSpeaker(value TeamworkPeripheralable)()
}
