package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type DirectSharingAbilities struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewDirectSharingAbilities instantiates a new DirectSharingAbilities and sets the default values.
func NewDirectSharingAbilities()(*DirectSharingAbilities) {
    m := &DirectSharingAbilities{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDirectSharingAbilitiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectSharingAbilitiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectSharingAbilities(), nil
}
// GetAddExistingExternalUsers gets the addExistingExternalUsers property value. Indicates whether the current user can add existing guest recipients to this item using direct sharing.
// returns a SharingOperationStatusable when successful
func (m *DirectSharingAbilities) GetAddExistingExternalUsers()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("addExistingExternalUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// GetAddInternalUsers gets the addInternalUsers property value. Indicates whether the current user can add internal recipients to this item using direct sharing.
// returns a SharingOperationStatusable when successful
func (m *DirectSharingAbilities) GetAddInternalUsers()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("addInternalUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *DirectSharingAbilities) GetAdditionalData()(map[string]any) {
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
// GetAddNewExternalUsers gets the addNewExternalUsers property value. Indicates whether the current user can add new guest recipients to this item using direct sharing.
// returns a SharingOperationStatusable when successful
func (m *DirectSharingAbilities) GetAddNewExternalUsers()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("addNewExternalUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *DirectSharingAbilities) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DirectSharingAbilities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addExistingExternalUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddExistingExternalUsers(val.(SharingOperationStatusable))
        }
        return nil
    }
    res["addInternalUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddInternalUsers(val.(SharingOperationStatusable))
        }
        return nil
    }
    res["addNewExternalUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddNewExternalUsers(val.(SharingOperationStatusable))
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
    res["requestGrantAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestGrantAccess(val.(SharingOperationStatusable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *DirectSharingAbilities) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestGrantAccess gets the requestGrantAccess property value. Indicates whether the user querying this endpoint can request access for the user or on behalf of other users, after which, site admins, can approve or deny the creation of a potential sharing link.
// returns a SharingOperationStatusable when successful
func (m *DirectSharingAbilities) GetRequestGrantAccess()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("requestGrantAccess")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DirectSharingAbilities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("addExistingExternalUsers", m.GetAddExistingExternalUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("addInternalUsers", m.GetAddInternalUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("addNewExternalUsers", m.GetAddNewExternalUsers())
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
        err := writer.WriteObjectValue("requestGrantAccess", m.GetRequestGrantAccess())
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
// SetAddExistingExternalUsers sets the addExistingExternalUsers property value. Indicates whether the current user can add existing guest recipients to this item using direct sharing.
func (m *DirectSharingAbilities) SetAddExistingExternalUsers(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("addExistingExternalUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetAddInternalUsers sets the addInternalUsers property value. Indicates whether the current user can add internal recipients to this item using direct sharing.
func (m *DirectSharingAbilities) SetAddInternalUsers(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("addInternalUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DirectSharingAbilities) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAddNewExternalUsers sets the addNewExternalUsers property value. Indicates whether the current user can add new guest recipients to this item using direct sharing.
func (m *DirectSharingAbilities) SetAddNewExternalUsers(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("addNewExternalUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *DirectSharingAbilities) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DirectSharingAbilities) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestGrantAccess sets the requestGrantAccess property value. Indicates whether the user querying this endpoint can request access for the user or on behalf of other users, after which, site admins, can approve or deny the creation of a potential sharing link.
func (m *DirectSharingAbilities) SetRequestGrantAccess(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("requestGrantAccess", value)
    if err != nil {
        panic(err)
    }
}
type DirectSharingAbilitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddExistingExternalUsers()(SharingOperationStatusable)
    GetAddInternalUsers()(SharingOperationStatusable)
    GetAddNewExternalUsers()(SharingOperationStatusable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetRequestGrantAccess()(SharingOperationStatusable)
    SetAddExistingExternalUsers(value SharingOperationStatusable)()
    SetAddInternalUsers(value SharingOperationStatusable)()
    SetAddNewExternalUsers(value SharingOperationStatusable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetRequestGrantAccess(value SharingOperationStatusable)()
}
