package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkTeamsClientConfiguration 
type TeamworkTeamsClientConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkTeamsClientConfiguration instantiates a new teamworkTeamsClientConfiguration and sets the default values.
func NewTeamworkTeamsClientConfiguration()(*TeamworkTeamsClientConfiguration) {
    m := &TeamworkTeamsClientConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkTeamsClientConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkTeamsClientConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkTeamsClientConfiguration(), nil
}
// GetAccountConfiguration gets the accountConfiguration property value. The configuration of the Microsoft Teams client user account for a device.
func (m *TeamworkTeamsClientConfiguration) GetAccountConfiguration()(TeamworkAccountConfigurationable) {
    val, err := m.GetBackingStore().Get("accountConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkAccountConfigurationable)
    }
    return nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkTeamsClientConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *TeamworkTeamsClientConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFeaturesConfiguration gets the featuresConfiguration property value. The configuration of Microsoft Teams client features for a device.
func (m *TeamworkTeamsClientConfiguration) GetFeaturesConfiguration()(TeamworkFeaturesConfigurationable) {
    val, err := m.GetBackingStore().Get("featuresConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkFeaturesConfigurationable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkTeamsClientConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkAccountConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountConfiguration(val.(TeamworkAccountConfigurationable))
        }
        return nil
    }
    res["featuresConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkFeaturesConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeaturesConfiguration(val.(TeamworkFeaturesConfigurationable))
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkTeamsClientConfiguration) GetOdataType()(*string) {
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
func (m *TeamworkTeamsClientConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("accountConfiguration", m.GetAccountConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("featuresConfiguration", m.GetFeaturesConfiguration())
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
// SetAccountConfiguration sets the accountConfiguration property value. The configuration of the Microsoft Teams client user account for a device.
func (m *TeamworkTeamsClientConfiguration) SetAccountConfiguration(value TeamworkAccountConfigurationable)() {
    err := m.GetBackingStore().Set("accountConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkTeamsClientConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *TeamworkTeamsClientConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetFeaturesConfiguration sets the featuresConfiguration property value. The configuration of Microsoft Teams client features for a device.
func (m *TeamworkTeamsClientConfiguration) SetFeaturesConfiguration(value TeamworkFeaturesConfigurationable)() {
    err := m.GetBackingStore().Set("featuresConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkTeamsClientConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkTeamsClientConfigurationable 
type TeamworkTeamsClientConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountConfiguration()(TeamworkAccountConfigurationable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetFeaturesConfiguration()(TeamworkFeaturesConfigurationable)
    GetOdataType()(*string)
    SetAccountConfiguration(value TeamworkAccountConfigurationable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetFeaturesConfiguration(value TeamworkFeaturesConfigurationable)()
    SetOdataType(value *string)()
}
