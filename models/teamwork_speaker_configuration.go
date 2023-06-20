package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkSpeakerConfiguration 
type TeamworkSpeakerConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkSpeakerConfiguration instantiates a new teamworkSpeakerConfiguration and sets the default values.
func NewTeamworkSpeakerConfiguration()(*TeamworkSpeakerConfiguration) {
    m := &TeamworkSpeakerConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkSpeakerConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkSpeakerConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkSpeakerConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkSpeakerConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *TeamworkSpeakerConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDefaultCommunicationSpeaker gets the defaultCommunicationSpeaker property value. The defaultCommunicationSpeaker property
func (m *TeamworkSpeakerConfiguration) GetDefaultCommunicationSpeaker()(TeamworkPeripheralable) {
    val, err := m.GetBackingStore().Get("defaultCommunicationSpeaker")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralable)
    }
    return nil
}
// GetDefaultSpeaker gets the defaultSpeaker property value. The defaultSpeaker property
func (m *TeamworkSpeakerConfiguration) GetDefaultSpeaker()(TeamworkPeripheralable) {
    val, err := m.GetBackingStore().Get("defaultSpeaker")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkSpeakerConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultCommunicationSpeaker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultCommunicationSpeaker(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["defaultSpeaker"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultSpeaker(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["isCommunicationSpeakerOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCommunicationSpeakerOptional(val)
        }
        return nil
    }
    res["isSpeakerOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSpeakerOptional(val)
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
    res["speakers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamworkPeripheralable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TeamworkPeripheralable)
                }
            }
            m.SetSpeakers(res)
        }
        return nil
    }
    return res
}
// GetIsCommunicationSpeakerOptional gets the isCommunicationSpeakerOptional property value. True if the communication speaker is optional. Used to compute the health state if the communication speaker is not optional.
func (m *TeamworkSpeakerConfiguration) GetIsCommunicationSpeakerOptional()(*bool) {
    val, err := m.GetBackingStore().Get("isCommunicationSpeakerOptional")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsSpeakerOptional gets the isSpeakerOptional property value. True if the configured speaker is optional. Used to compute the health state if the speaker is not optional.
func (m *TeamworkSpeakerConfiguration) GetIsSpeakerOptional()(*bool) {
    val, err := m.GetBackingStore().Get("isSpeakerOptional")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkSpeakerConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSpeakers gets the speakers property value. The speakers property
func (m *TeamworkSpeakerConfiguration) GetSpeakers()([]TeamworkPeripheralable) {
    val, err := m.GetBackingStore().Get("speakers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TeamworkPeripheralable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkSpeakerConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("defaultCommunicationSpeaker", m.GetDefaultCommunicationSpeaker())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("defaultSpeaker", m.GetDefaultSpeaker())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isCommunicationSpeakerOptional", m.GetIsCommunicationSpeakerOptional())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSpeakerOptional", m.GetIsSpeakerOptional())
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
    if m.GetSpeakers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSpeakers()))
        for i, v := range m.GetSpeakers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("speakers", cast)
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
func (m *TeamworkSpeakerConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *TeamworkSpeakerConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDefaultCommunicationSpeaker sets the defaultCommunicationSpeaker property value. The defaultCommunicationSpeaker property
func (m *TeamworkSpeakerConfiguration) SetDefaultCommunicationSpeaker(value TeamworkPeripheralable)() {
    err := m.GetBackingStore().Set("defaultCommunicationSpeaker", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultSpeaker sets the defaultSpeaker property value. The defaultSpeaker property
func (m *TeamworkSpeakerConfiguration) SetDefaultSpeaker(value TeamworkPeripheralable)() {
    err := m.GetBackingStore().Set("defaultSpeaker", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCommunicationSpeakerOptional sets the isCommunicationSpeakerOptional property value. True if the communication speaker is optional. Used to compute the health state if the communication speaker is not optional.
func (m *TeamworkSpeakerConfiguration) SetIsCommunicationSpeakerOptional(value *bool)() {
    err := m.GetBackingStore().Set("isCommunicationSpeakerOptional", value)
    if err != nil {
        panic(err)
    }
}
// SetIsSpeakerOptional sets the isSpeakerOptional property value. True if the configured speaker is optional. Used to compute the health state if the speaker is not optional.
func (m *TeamworkSpeakerConfiguration) SetIsSpeakerOptional(value *bool)() {
    err := m.GetBackingStore().Set("isSpeakerOptional", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkSpeakerConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSpeakers sets the speakers property value. The speakers property
func (m *TeamworkSpeakerConfiguration) SetSpeakers(value []TeamworkPeripheralable)() {
    err := m.GetBackingStore().Set("speakers", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkSpeakerConfigurationable 
type TeamworkSpeakerConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDefaultCommunicationSpeaker()(TeamworkPeripheralable)
    GetDefaultSpeaker()(TeamworkPeripheralable)
    GetIsCommunicationSpeakerOptional()(*bool)
    GetIsSpeakerOptional()(*bool)
    GetOdataType()(*string)
    GetSpeakers()([]TeamworkPeripheralable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDefaultCommunicationSpeaker(value TeamworkPeripheralable)()
    SetDefaultSpeaker(value TeamworkPeripheralable)()
    SetIsCommunicationSpeakerOptional(value *bool)()
    SetIsSpeakerOptional(value *bool)()
    SetOdataType(value *string)()
    SetSpeakers(value []TeamworkPeripheralable)()
}
