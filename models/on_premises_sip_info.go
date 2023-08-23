package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// OnPremisesSipInfo 
type OnPremisesSipInfo struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewOnPremisesSipInfo instantiates a new onPremisesSipInfo and sets the default values.
func NewOnPremisesSipInfo()(*OnPremisesSipInfo) {
    m := &OnPremisesSipInfo{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnPremisesSipInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesSipInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesSipInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesSipInfo) GetAdditionalData()(map[string]any) {
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
func (m *OnPremisesSipInfo) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesSipInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isSipEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSipEnabled(val)
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
    res["sipDeploymentLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSipDeploymentLocation(val)
        }
        return nil
    }
    res["sipPrimaryAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSipPrimaryAddress(val)
        }
        return nil
    }
    return res
}
// GetIsSipEnabled gets the isSipEnabled property value. Indicates whether the user is currently enabled for on-premises Skype for Business.
func (m *OnPremisesSipInfo) GetIsSipEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isSipEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnPremisesSipInfo) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSipDeploymentLocation gets the sipDeploymentLocation property value. Indicates a fully qualified DNS name of the Microsoft Online Communications Server deployment.
func (m *OnPremisesSipInfo) GetSipDeploymentLocation()(*string) {
    val, err := m.GetBackingStore().Get("sipDeploymentLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSipPrimaryAddress gets the sipPrimaryAddress property value. Serves as a unique identifier for each user on the on-premises Skype for Business.
func (m *OnPremisesSipInfo) GetSipPrimaryAddress()(*string) {
    val, err := m.GetBackingStore().Get("sipPrimaryAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OnPremisesSipInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isSipEnabled", m.GetIsSipEnabled())
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
        err := writer.WriteStringValue("sipDeploymentLocation", m.GetSipDeploymentLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sipPrimaryAddress", m.GetSipPrimaryAddress())
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
func (m *OnPremisesSipInfo) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *OnPremisesSipInfo) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetIsSipEnabled sets the isSipEnabled property value. Indicates whether the user is currently enabled for on-premises Skype for Business.
func (m *OnPremisesSipInfo) SetIsSipEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isSipEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesSipInfo) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSipDeploymentLocation sets the sipDeploymentLocation property value. Indicates a fully qualified DNS name of the Microsoft Online Communications Server deployment.
func (m *OnPremisesSipInfo) SetSipDeploymentLocation(value *string)() {
    err := m.GetBackingStore().Set("sipDeploymentLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetSipPrimaryAddress sets the sipPrimaryAddress property value. Serves as a unique identifier for each user on the on-premises Skype for Business.
func (m *OnPremisesSipInfo) SetSipPrimaryAddress(value *string)() {
    err := m.GetBackingStore().Set("sipPrimaryAddress", value)
    if err != nil {
        panic(err)
    }
}
// OnPremisesSipInfoable 
type OnPremisesSipInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetIsSipEnabled()(*bool)
    GetOdataType()(*string)
    GetSipDeploymentLocation()(*string)
    GetSipPrimaryAddress()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetIsSipEnabled(value *bool)()
    SetOdataType(value *string)()
    SetSipDeploymentLocation(value *string)()
    SetSipPrimaryAddress(value *string)()
}
