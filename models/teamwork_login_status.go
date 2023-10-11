package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// TeamworkLoginStatus 
type TeamworkLoginStatus struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewTeamworkLoginStatus instantiates a new teamworkLoginStatus and sets the default values.
func NewTeamworkLoginStatus()(*TeamworkLoginStatus) {
    m := &TeamworkLoginStatus{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkLoginStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkLoginStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkLoginStatus(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkLoginStatus) GetAdditionalData()(map[string]any) {
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
func (m *TeamworkLoginStatus) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetExchangeConnection gets the exchangeConnection property value. Information about the Exchange connection.
func (m *TeamworkLoginStatus) GetExchangeConnection()(TeamworkConnectionable) {
    val, err := m.GetBackingStore().Get("exchangeConnection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkConnectionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkLoginStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["exchangeConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeConnection(val.(TeamworkConnectionable))
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
    res["skypeConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkypeConnection(val.(TeamworkConnectionable))
        }
        return nil
    }
    res["teamsConnection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamworkConnectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsConnection(val.(TeamworkConnectionable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkLoginStatus) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSkypeConnection gets the skypeConnection property value. Information about the Skype for Business connection.
func (m *TeamworkLoginStatus) GetSkypeConnection()(TeamworkConnectionable) {
    val, err := m.GetBackingStore().Get("skypeConnection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkConnectionable)
    }
    return nil
}
// GetTeamsConnection gets the teamsConnection property value. Information about the Teams connection.
func (m *TeamworkLoginStatus) GetTeamsConnection()(TeamworkConnectionable) {
    val, err := m.GetBackingStore().Get("teamsConnection")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TeamworkConnectionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TeamworkLoginStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("exchangeConnection", m.GetExchangeConnection())
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
        err := writer.WriteObjectValue("skypeConnection", m.GetSkypeConnection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("teamsConnection", m.GetTeamsConnection())
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
func (m *TeamworkLoginStatus) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *TeamworkLoginStatus) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetExchangeConnection sets the exchangeConnection property value. Information about the Exchange connection.
func (m *TeamworkLoginStatus) SetExchangeConnection(value TeamworkConnectionable)() {
    err := m.GetBackingStore().Set("exchangeConnection", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkLoginStatus) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSkypeConnection sets the skypeConnection property value. Information about the Skype for Business connection.
func (m *TeamworkLoginStatus) SetSkypeConnection(value TeamworkConnectionable)() {
    err := m.GetBackingStore().Set("skypeConnection", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsConnection sets the teamsConnection property value. Information about the Teams connection.
func (m *TeamworkLoginStatus) SetTeamsConnection(value TeamworkConnectionable)() {
    err := m.GetBackingStore().Set("teamsConnection", value)
    if err != nil {
        panic(err)
    }
}
// TeamworkLoginStatusable 
type TeamworkLoginStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetExchangeConnection()(TeamworkConnectionable)
    GetOdataType()(*string)
    GetSkypeConnection()(TeamworkConnectionable)
    GetTeamsConnection()(TeamworkConnectionable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetExchangeConnection(value TeamworkConnectionable)()
    SetOdataType(value *string)()
    SetSkypeConnection(value TeamworkConnectionable)()
    SetTeamsConnection(value TeamworkConnectionable)()
}
