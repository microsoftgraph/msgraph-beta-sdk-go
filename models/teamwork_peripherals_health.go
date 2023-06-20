package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkPeripheralsHealth 
type TeamworkPeripheralsHealth struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkPeripheralsHealth instantiates a new teamworkPeripheralsHealth and sets the default values.
func NewTeamworkPeripheralsHealth()(*TeamworkPeripheralsHealth) {
    m := &TeamworkPeripheralsHealth{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkPeripheralsHealthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkPeripheralsHealthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkPeripheralsHealth(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkPeripheralsHealth) GetAdditionalData()(map[string]any) {
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
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *TeamworkPeripheralsHealth) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCommunicationSpeakerHealth gets the communicationSpeakerHealth property value. The health details about the communication speaker.
func (m *TeamworkPeripheralsHealth) GetCommunicationSpeakerHealth()(TeamworkPeripheralHealthable) {
    val, err := m.GetBackingStore().Get("communicationSpeakerHealth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralHealthable)
    }
    return nil
}
// GetContentCameraHealth gets the contentCameraHealth property value. The health details about the content camera.
func (m *TeamworkPeripheralsHealth) GetContentCameraHealth()(TeamworkPeripheralHealthable) {
    val, err := m.GetBackingStore().Get("contentCameraHealth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralHealthable)
    }
    return nil
}
// GetDisplayHealthCollection gets the displayHealthCollection property value. The health details about displays.
func (m *TeamworkPeripheralsHealth) GetDisplayHealthCollection()([]TeamworkPeripheralHealthable) {
    val, err := m.GetBackingStore().Get("displayHealthCollection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TeamworkPeripheralHealthable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkPeripheralsHealth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["communicationSpeakerHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommunicationSpeakerHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["contentCameraHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentCameraHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["displayHealthCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkPeripheralHealthable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TeamworkPeripheralHealthable)
                }
            }
            m.SetDisplayHealthCollection(res)
        }
        return nil
    }
    res["microphoneHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrophoneHealth(val.(TeamworkPeripheralHealthable))
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
    res["roomCameraHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomCameraHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    res["speakerHealth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeakerHealth(val.(TeamworkPeripheralHealthable))
        }
        return nil
    }
    return res
}
// GetMicrophoneHealth gets the microphoneHealth property value. The health details about the microphone.
func (m *TeamworkPeripheralsHealth) GetMicrophoneHealth()(TeamworkPeripheralHealthable) {
    val, err := m.GetBackingStore().Get("microphoneHealth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralHealthable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkPeripheralsHealth) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoomCameraHealth gets the roomCameraHealth property value. The health details about the room camera.
func (m *TeamworkPeripheralsHealth) GetRoomCameraHealth()(TeamworkPeripheralHealthable) {
    val, err := m.GetBackingStore().Get("roomCameraHealth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralHealthable)
    }
    return nil
}
// GetSpeakerHealth gets the speakerHealth property value. The health details about the speaker.
func (m *TeamworkPeripheralsHealth) GetSpeakerHealth()(TeamworkPeripheralHealthable) {
    val, err := m.GetBackingStore().Get("speakerHealth")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralHealthable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkPeripheralsHealth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("communicationSpeakerHealth", m.GetCommunicationSpeakerHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentCameraHealth", m.GetContentCameraHealth())
        if err != nil {
            return err
        }
    }
    if m.GetDisplayHealthCollection() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDisplayHealthCollection()))
        for i, v := range m.GetDisplayHealthCollection() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("displayHealthCollection", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("microphoneHealth", m.GetMicrophoneHealth())
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
        err := writer.WriteObjectValue("roomCameraHealth", m.GetRoomCameraHealth())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("speakerHealth", m.GetSpeakerHealth())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkPeripheralsHealth) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *TeamworkPeripheralsHealth) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCommunicationSpeakerHealth sets the communicationSpeakerHealth property value. The health details about the communication speaker.
func (m *TeamworkPeripheralsHealth) SetCommunicationSpeakerHealth(value TeamworkPeripheralHealthable)() {
    err := m.GetBackingStore().Set("communicationSpeakerHealth", value)
    if err != nil {
        panic(err)
    }
}
// SetContentCameraHealth sets the contentCameraHealth property value. The health details about the content camera.
func (m *TeamworkPeripheralsHealth) SetContentCameraHealth(value TeamworkPeripheralHealthable)() {
    err := m.GetBackingStore().Set("contentCameraHealth", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayHealthCollection sets the displayHealthCollection property value. The health details about displays.
func (m *TeamworkPeripheralsHealth) SetDisplayHealthCollection(value []TeamworkPeripheralHealthable)() {
    err := m.GetBackingStore().Set("displayHealthCollection", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrophoneHealth sets the microphoneHealth property value. The health details about the microphone.
func (m *TeamworkPeripheralsHealth) SetMicrophoneHealth(value TeamworkPeripheralHealthable)() {
    err := m.GetBackingStore().Set("microphoneHealth", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkPeripheralsHealth) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRoomCameraHealth sets the roomCameraHealth property value. The health details about the room camera.
func (m *TeamworkPeripheralsHealth) SetRoomCameraHealth(value TeamworkPeripheralHealthable)() {
    err := m.GetBackingStore().Set("roomCameraHealth", value)
    if err != nil {
        panic(err)
    }
}
// SetSpeakerHealth sets the speakerHealth property value. The health details about the speaker.
func (m *TeamworkPeripheralsHealth) SetSpeakerHealth(value TeamworkPeripheralHealthable)() {
    err := m.GetBackingStore().Set("speakerHealth", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkPeripheralsHealthable 
type TeamworkPeripheralsHealthable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCommunicationSpeakerHealth()(TeamworkPeripheralHealthable)
    GetContentCameraHealth()(TeamworkPeripheralHealthable)
    GetDisplayHealthCollection()([]TeamworkPeripheralHealthable)
    GetMicrophoneHealth()(TeamworkPeripheralHealthable)
    GetOdataType()(*string)
    GetRoomCameraHealth()(TeamworkPeripheralHealthable)
    GetSpeakerHealth()(TeamworkPeripheralHealthable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCommunicationSpeakerHealth(value TeamworkPeripheralHealthable)()
    SetContentCameraHealth(value TeamworkPeripheralHealthable)()
    SetDisplayHealthCollection(value []TeamworkPeripheralHealthable)()
    SetMicrophoneHealth(value TeamworkPeripheralHealthable)()
    SetOdataType(value *string)()
    SetRoomCameraHealth(value TeamworkPeripheralHealthable)()
    SetSpeakerHealth(value TeamworkPeripheralHealthable)()
}
