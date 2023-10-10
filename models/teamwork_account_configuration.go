package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkAccountConfiguration 
type TeamworkAccountConfiguration struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkAccountConfiguration instantiates a new teamworkAccountConfiguration and sets the default values.
func NewTeamworkAccountConfiguration()(*TeamworkAccountConfiguration) {
    m := &TeamworkAccountConfiguration{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkAccountConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkAccountConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkAccountConfiguration(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkAccountConfiguration) GetAdditionalData()(map[string]any) {
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
func (m *TeamworkAccountConfiguration) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkAccountConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["onPremisesCalendarSyncConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkOnPremisesCalendarSyncConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesCalendarSyncConfiguration(val.(TeamworkOnPremisesCalendarSyncConfigurationable))
        }
        return nil
    }
    res["supportedClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkSupportedClient)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportedClient(val.(*TeamworkSupportedClient))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkAccountConfiguration) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOnPremisesCalendarSyncConfiguration gets the onPremisesCalendarSyncConfiguration property value. The account used to sync the calendar.
func (m *TeamworkAccountConfiguration) GetOnPremisesCalendarSyncConfiguration()(TeamworkOnPremisesCalendarSyncConfigurationable) {
    val, err := m.GetBackingStore().Get("onPremisesCalendarSyncConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkOnPremisesCalendarSyncConfigurationable)
    }
    return nil
}
// GetSupportedClient gets the supportedClient property value. The supported client for Teams Rooms devices. The possible values are: unknown, skypeDefaultAndTeams, teamsDefaultAndSkype, skypeOnly, teamsOnly, unknownFutureValue.
func (m *TeamworkAccountConfiguration) GetSupportedClient()(*TeamworkSupportedClient) {
    val, err := m.GetBackingStore().Get("supportedClient")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*TeamworkSupportedClient)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkAccountConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("onPremisesCalendarSyncConfiguration", m.GetOnPremisesCalendarSyncConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetSupportedClient() != nil {
        cast := (*m.GetSupportedClient()).String()
        err := writer.WriteStringValue("supportedClient", &cast)
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
func (m *TeamworkAccountConfiguration) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeamworkAccountConfiguration) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkAccountConfiguration) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesCalendarSyncConfiguration sets the onPremisesCalendarSyncConfiguration property value. The account used to sync the calendar.
func (m *TeamworkAccountConfiguration) SetOnPremisesCalendarSyncConfiguration(value TeamworkOnPremisesCalendarSyncConfigurationable)() {
    err := m.GetBackingStore().Set("onPremisesCalendarSyncConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportedClient sets the supportedClient property value. The supported client for Teams Rooms devices. The possible values are: unknown, skypeDefaultAndTeams, teamsDefaultAndSkype, skypeOnly, teamsOnly, unknownFutureValue.
func (m *TeamworkAccountConfiguration) SetSupportedClient(value *TeamworkSupportedClient)() {
    err := m.GetBackingStore().Set("supportedClient", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkAccountConfigurationable 
type TeamworkAccountConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetOnPremisesCalendarSyncConfiguration()(TeamworkOnPremisesCalendarSyncConfigurationable)
    GetSupportedClient()(*TeamworkSupportedClient)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetOnPremisesCalendarSyncConfiguration(value TeamworkOnPremisesCalendarSyncConfigurationable)()
    SetSupportedClient(value *TeamworkSupportedClient)()
}
