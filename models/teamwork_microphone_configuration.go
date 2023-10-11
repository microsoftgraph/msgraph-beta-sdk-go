package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkMicrophoneConfiguration 
type TeamworkMicrophoneConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkMicrophoneConfiguration instantiates a new teamworkMicrophoneConfiguration and sets the default values.
func NewTeamworkMicrophoneConfiguration()(*TeamworkMicrophoneConfiguration) {
    m := &TeamworkMicrophoneConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkMicrophoneConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkMicrophoneConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkMicrophoneConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkMicrophoneConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *TeamworkMicrophoneConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDefaultMicrophone gets the defaultMicrophone property value. The defaultMicrophone property
func (m *TeamworkMicrophoneConfiguration) GetDefaultMicrophone()(TeamworkPeripheralable) {
    val, err := m.GetBackingStore().Get("defaultMicrophone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkPeripheralable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkMicrophoneConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultMicrophone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkPeripheralFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultMicrophone(val.(TeamworkPeripheralable))
        }
        return nil
    }
    res["isMicrophoneOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMicrophoneOptional(val)
        }
        return nil
    }
    res["microphones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetMicrophones(res)
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
// GetIsMicrophoneOptional gets the isMicrophoneOptional property value. True if the configured microphone is optional. False if the microphone is not optional and the health state of the device should be computed.
func (m *TeamworkMicrophoneConfiguration) GetIsMicrophoneOptional()(*bool) {
    val, err := m.GetBackingStore().Get("isMicrophoneOptional")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMicrophones gets the microphones property value. The microphones property
func (m *TeamworkMicrophoneConfiguration) GetMicrophones()([]TeamworkPeripheralable) {
    val, err := m.GetBackingStore().Get("microphones")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TeamworkPeripheralable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkMicrophoneConfiguration) GetOdataType()(*string) {
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
func (m *TeamworkMicrophoneConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("defaultMicrophone", m.GetDefaultMicrophone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMicrophoneOptional", m.GetIsMicrophoneOptional())
        if err != nil {
            return err
        }
    }
    if m.GetMicrophones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrophones()))
        for i, v := range m.GetMicrophones() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("microphones", cast)
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
func (m *TeamworkMicrophoneConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeamworkMicrophoneConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDefaultMicrophone sets the defaultMicrophone property value. The defaultMicrophone property
func (m *TeamworkMicrophoneConfiguration) SetDefaultMicrophone(value TeamworkPeripheralable)() {
    err := m.GetBackingStore().Set("defaultMicrophone", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMicrophoneOptional sets the isMicrophoneOptional property value. True if the configured microphone is optional. False if the microphone is not optional and the health state of the device should be computed.
func (m *TeamworkMicrophoneConfiguration) SetIsMicrophoneOptional(value *bool)() {
    err := m.GetBackingStore().Set("isMicrophoneOptional", value)
    if err != nil {
        panic(err)
    }
}
// SetMicrophones sets the microphones property value. The microphones property
func (m *TeamworkMicrophoneConfiguration) SetMicrophones(value []TeamworkPeripheralable)() {
    err := m.GetBackingStore().Set("microphones", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkMicrophoneConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkMicrophoneConfigurationable 
type TeamworkMicrophoneConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDefaultMicrophone()(TeamworkPeripheralable)
    GetIsMicrophoneOptional()(*bool)
    GetMicrophones()([]TeamworkPeripheralable)
    GetOdataType()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDefaultMicrophone(value TeamworkPeripheralable)()
    SetIsMicrophoneOptional(value *bool)()
    SetMicrophones(value []TeamworkPeripheralable)()
    SetOdataType(value *string)()
}
