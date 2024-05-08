package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type SharePointSharingAbilities struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSharePointSharingAbilities instantiates a new SharePointSharingAbilities and sets the default values.
func NewSharePointSharingAbilities()(*SharePointSharingAbilities) {
    m := &SharePointSharingAbilities{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSharePointSharingAbilitiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSharePointSharingAbilitiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharePointSharingAbilities(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SharePointSharingAbilities) GetAdditionalData()(map[string]any) {
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
// GetAnyoneLinkAbilities gets the anyoneLinkAbilities property value. The anyone link abilities.
// returns a LinkScopeAbilitiesable when successful
func (m *SharePointSharingAbilities) GetAnyoneLinkAbilities()(LinkScopeAbilitiesable) {
    val, err := m.GetBackingStore().Get("anyoneLinkAbilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LinkScopeAbilitiesable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *SharePointSharingAbilities) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDirectSharingAbilities gets the directSharingAbilities property value. The direct sharing abilities.
// returns a DirectSharingAbilitiesable when successful
func (m *SharePointSharingAbilities) GetDirectSharingAbilities()(DirectSharingAbilitiesable) {
    val, err := m.GetBackingStore().Get("directSharingAbilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DirectSharingAbilitiesable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SharePointSharingAbilities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["anyoneLinkAbilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkScopeAbilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnyoneLinkAbilities(val.(LinkScopeAbilitiesable))
        }
        return nil
    }
    res["directSharingAbilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectSharingAbilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectSharingAbilities(val.(DirectSharingAbilitiesable))
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
    res["organizationLinkAbilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkScopeAbilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationLinkAbilities(val.(LinkScopeAbilitiesable))
        }
        return nil
    }
    res["specificPeopleLinkAbilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkScopeAbilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpecificPeopleLinkAbilities(val.(LinkScopeAbilitiesable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SharePointSharingAbilities) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrganizationLinkAbilities gets the organizationLinkAbilities property value. The organization link abilities.
// returns a LinkScopeAbilitiesable when successful
func (m *SharePointSharingAbilities) GetOrganizationLinkAbilities()(LinkScopeAbilitiesable) {
    val, err := m.GetBackingStore().Get("organizationLinkAbilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LinkScopeAbilitiesable)
    }
    return nil
}
// GetSpecificPeopleLinkAbilities gets the specificPeopleLinkAbilities property value. The specificPeople link abilities.
// returns a LinkScopeAbilitiesable when successful
func (m *SharePointSharingAbilities) GetSpecificPeopleLinkAbilities()(LinkScopeAbilitiesable) {
    val, err := m.GetBackingStore().Get("specificPeopleLinkAbilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LinkScopeAbilitiesable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SharePointSharingAbilities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("anyoneLinkAbilities", m.GetAnyoneLinkAbilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("directSharingAbilities", m.GetDirectSharingAbilities())
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
        err := writer.WriteObjectValue("organizationLinkAbilities", m.GetOrganizationLinkAbilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("specificPeopleLinkAbilities", m.GetSpecificPeopleLinkAbilities())
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
func (m *SharePointSharingAbilities) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAnyoneLinkAbilities sets the anyoneLinkAbilities property value. The anyone link abilities.
func (m *SharePointSharingAbilities) SetAnyoneLinkAbilities(value LinkScopeAbilitiesable)() {
    err := m.GetBackingStore().Set("anyoneLinkAbilities", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *SharePointSharingAbilities) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDirectSharingAbilities sets the directSharingAbilities property value. The direct sharing abilities.
func (m *SharePointSharingAbilities) SetDirectSharingAbilities(value DirectSharingAbilitiesable)() {
    err := m.GetBackingStore().Set("directSharingAbilities", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SharePointSharingAbilities) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationLinkAbilities sets the organizationLinkAbilities property value. The organization link abilities.
func (m *SharePointSharingAbilities) SetOrganizationLinkAbilities(value LinkScopeAbilitiesable)() {
    err := m.GetBackingStore().Set("organizationLinkAbilities", value)
    if err != nil {
        panic(err)
    }
}
// SetSpecificPeopleLinkAbilities sets the specificPeopleLinkAbilities property value. The specificPeople link abilities.
func (m *SharePointSharingAbilities) SetSpecificPeopleLinkAbilities(value LinkScopeAbilitiesable)() {
    err := m.GetBackingStore().Set("specificPeopleLinkAbilities", value)
    if err != nil {
        panic(err)
    }
}
type SharePointSharingAbilitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnyoneLinkAbilities()(LinkScopeAbilitiesable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDirectSharingAbilities()(DirectSharingAbilitiesable)
    GetOdataType()(*string)
    GetOrganizationLinkAbilities()(LinkScopeAbilitiesable)
    GetSpecificPeopleLinkAbilities()(LinkScopeAbilitiesable)
    SetAnyoneLinkAbilities(value LinkScopeAbilitiesable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDirectSharingAbilities(value DirectSharingAbilitiesable)()
    SetOdataType(value *string)()
    SetOrganizationLinkAbilities(value LinkScopeAbilitiesable)()
    SetSpecificPeopleLinkAbilities(value LinkScopeAbilitiesable)()
}
