package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type LinkScopeAbilities struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewLinkScopeAbilities instantiates a new LinkScopeAbilities and sets the default values.
func NewLinkScopeAbilities()(*LinkScopeAbilities) {
    m := &LinkScopeAbilities{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLinkScopeAbilitiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLinkScopeAbilitiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLinkScopeAbilities(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LinkScopeAbilities) GetAdditionalData()(map[string]any) {
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
func (m *LinkScopeAbilities) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetBlockDownloadLinkAbilities gets the blockDownloadLinkAbilities property value. The blockDownload link abilities.
// returns a LinkRoleAbilitiesable when successful
func (m *LinkScopeAbilities) GetBlockDownloadLinkAbilities()(LinkRoleAbilitiesable) {
    val, err := m.GetBackingStore().Get("blockDownloadLinkAbilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LinkRoleAbilitiesable)
    }
    return nil
}
// GetEditLinkAbilities gets the editLinkAbilities property value. The editLinkAbilities property
// returns a LinkRoleAbilitiesable when successful
func (m *LinkScopeAbilities) GetEditLinkAbilities()(LinkRoleAbilitiesable) {
    val, err := m.GetBackingStore().Get("editLinkAbilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LinkRoleAbilitiesable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LinkScopeAbilities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["blockDownloadLinkAbilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkRoleAbilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlockDownloadLinkAbilities(val.(LinkRoleAbilitiesable))
        }
        return nil
    }
    res["editLinkAbilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkRoleAbilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEditLinkAbilities(val.(LinkRoleAbilitiesable))
        }
        return nil
    }
    res["manageListLinkAbilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkRoleAbilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManageListLinkAbilities(val.(LinkRoleAbilitiesable))
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
    res["readLinkAbilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkRoleAbilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadLinkAbilities(val.(LinkRoleAbilitiesable))
        }
        return nil
    }
    res["reviewLinkAbilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkRoleAbilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewLinkAbilities(val.(LinkRoleAbilitiesable))
        }
        return nil
    }
    res["submitOnlyLinkAbilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLinkRoleAbilitiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmitOnlyLinkAbilities(val.(LinkRoleAbilitiesable))
        }
        return nil
    }
    return res
}
// GetManageListLinkAbilities gets the manageListLinkAbilities property value. The manageList link abilities.
// returns a LinkRoleAbilitiesable when successful
func (m *LinkScopeAbilities) GetManageListLinkAbilities()(LinkRoleAbilitiesable) {
    val, err := m.GetBackingStore().Get("manageListLinkAbilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LinkRoleAbilitiesable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *LinkScopeAbilities) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReadLinkAbilities gets the readLinkAbilities property value. The readLinkAbilities property
// returns a LinkRoleAbilitiesable when successful
func (m *LinkScopeAbilities) GetReadLinkAbilities()(LinkRoleAbilitiesable) {
    val, err := m.GetBackingStore().Get("readLinkAbilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LinkRoleAbilitiesable)
    }
    return nil
}
// GetReviewLinkAbilities gets the reviewLinkAbilities property value. The review link abilities.
// returns a LinkRoleAbilitiesable when successful
func (m *LinkScopeAbilities) GetReviewLinkAbilities()(LinkRoleAbilitiesable) {
    val, err := m.GetBackingStore().Get("reviewLinkAbilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LinkRoleAbilitiesable)
    }
    return nil
}
// GetSubmitOnlyLinkAbilities gets the submitOnlyLinkAbilities property value. The submitOnly link abilities.
// returns a LinkRoleAbilitiesable when successful
func (m *LinkScopeAbilities) GetSubmitOnlyLinkAbilities()(LinkRoleAbilitiesable) {
    val, err := m.GetBackingStore().Get("submitOnlyLinkAbilities")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(LinkRoleAbilitiesable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *LinkScopeAbilities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("blockDownloadLinkAbilities", m.GetBlockDownloadLinkAbilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("editLinkAbilities", m.GetEditLinkAbilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("manageListLinkAbilities", m.GetManageListLinkAbilities())
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
        err := writer.WriteObjectValue("readLinkAbilities", m.GetReadLinkAbilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("reviewLinkAbilities", m.GetReviewLinkAbilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("submitOnlyLinkAbilities", m.GetSubmitOnlyLinkAbilities())
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
func (m *LinkScopeAbilities) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *LinkScopeAbilities) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetBlockDownloadLinkAbilities sets the blockDownloadLinkAbilities property value. The blockDownload link abilities.
func (m *LinkScopeAbilities) SetBlockDownloadLinkAbilities(value LinkRoleAbilitiesable)() {
    err := m.GetBackingStore().Set("blockDownloadLinkAbilities", value)
    if err != nil {
        panic(err)
    }
}
// SetEditLinkAbilities sets the editLinkAbilities property value. The editLinkAbilities property
func (m *LinkScopeAbilities) SetEditLinkAbilities(value LinkRoleAbilitiesable)() {
    err := m.GetBackingStore().Set("editLinkAbilities", value)
    if err != nil {
        panic(err)
    }
}
// SetManageListLinkAbilities sets the manageListLinkAbilities property value. The manageList link abilities.
func (m *LinkScopeAbilities) SetManageListLinkAbilities(value LinkRoleAbilitiesable)() {
    err := m.GetBackingStore().Set("manageListLinkAbilities", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LinkScopeAbilities) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetReadLinkAbilities sets the readLinkAbilities property value. The readLinkAbilities property
func (m *LinkScopeAbilities) SetReadLinkAbilities(value LinkRoleAbilitiesable)() {
    err := m.GetBackingStore().Set("readLinkAbilities", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewLinkAbilities sets the reviewLinkAbilities property value. The review link abilities.
func (m *LinkScopeAbilities) SetReviewLinkAbilities(value LinkRoleAbilitiesable)() {
    err := m.GetBackingStore().Set("reviewLinkAbilities", value)
    if err != nil {
        panic(err)
    }
}
// SetSubmitOnlyLinkAbilities sets the submitOnlyLinkAbilities property value. The submitOnly link abilities.
func (m *LinkScopeAbilities) SetSubmitOnlyLinkAbilities(value LinkRoleAbilitiesable)() {
    err := m.GetBackingStore().Set("submitOnlyLinkAbilities", value)
    if err != nil {
        panic(err)
    }
}
type LinkScopeAbilitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetBlockDownloadLinkAbilities()(LinkRoleAbilitiesable)
    GetEditLinkAbilities()(LinkRoleAbilitiesable)
    GetManageListLinkAbilities()(LinkRoleAbilitiesable)
    GetOdataType()(*string)
    GetReadLinkAbilities()(LinkRoleAbilitiesable)
    GetReviewLinkAbilities()(LinkRoleAbilitiesable)
    GetSubmitOnlyLinkAbilities()(LinkRoleAbilitiesable)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetBlockDownloadLinkAbilities(value LinkRoleAbilitiesable)()
    SetEditLinkAbilities(value LinkRoleAbilitiesable)()
    SetManageListLinkAbilities(value LinkRoleAbilitiesable)()
    SetOdataType(value *string)()
    SetReadLinkAbilities(value LinkRoleAbilitiesable)()
    SetReviewLinkAbilities(value LinkRoleAbilitiesable)()
    SetSubmitOnlyLinkAbilities(value LinkRoleAbilitiesable)()
}
