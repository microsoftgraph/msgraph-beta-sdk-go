package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type LinkRoleAbilities struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewLinkRoleAbilities instantiates a new LinkRoleAbilities and sets the default values.
func NewLinkRoleAbilities()(*LinkRoleAbilities) {
    m := &LinkRoleAbilities{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLinkRoleAbilitiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLinkRoleAbilitiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLinkRoleAbilities(), nil
}
// GetAddExistingExternalUsers gets the addExistingExternalUsers property value. Indicates if the current user can add existing external user recipients to this sharing link.
// returns a SharingOperationStatusable when successful
func (m *LinkRoleAbilities) GetAddExistingExternalUsers()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("addExistingExternalUsers")
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
func (m *LinkRoleAbilities) GetAdditionalData()(map[string]any) {
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
// GetAddNewExternalUsers gets the addNewExternalUsers property value. Indicates if the current user can add new external user recipients to this sharing link.
// returns a SharingOperationStatusable when successful
func (m *LinkRoleAbilities) GetAddNewExternalUsers()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("addNewExternalUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// GetApplyVariants gets the applyVariants property value. Indicates the status of the potential sharing link variants. If selected, it generates a separate sharing link from the sharing link that would otherwise be generated without the variant, yet with an identical role and scope.
// returns a SharingLinkVariantsable when successful
func (m *LinkRoleAbilities) GetApplyVariants()(SharingLinkVariantsable) {
    val, err := m.GetBackingStore().Get("applyVariants")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingLinkVariantsable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *LinkRoleAbilities) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCreateLink gets the createLink property value. Indicates if links of this role can be created.
// returns a SharingOperationStatusable when successful
func (m *LinkRoleAbilities) GetCreateLink()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("createLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// GetDeleteLink gets the deleteLink property value. Indicates if links of this role can be deleted.
// returns a SharingOperationStatusable when successful
func (m *LinkRoleAbilities) GetDeleteLink()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("deleteLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LinkRoleAbilities) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["applyVariants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingLinkVariantsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyVariants(val.(SharingLinkVariantsable))
        }
        return nil
    }
    res["createLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreateLink(val.(SharingOperationStatusable))
        }
        return nil
    }
    res["deleteLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleteLink(val.(SharingOperationStatusable))
        }
        return nil
    }
    res["linkAllowsExternalUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkAllowsExternalUsers(val.(SharingOperationStatusable))
        }
        return nil
    }
    res["linkExpiration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingLinkExpirationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkExpiration(val.(SharingLinkExpirationStatusable))
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
    res["retrieveLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRetrieveLink(val.(SharingOperationStatusable))
        }
        return nil
    }
    res["updateLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateLink(val.(SharingOperationStatusable))
        }
        return nil
    }
    return res
}
// GetLinkAllowsExternalUsers gets the linkAllowsExternalUsers property value. Indicates if this link can include external users.
// returns a SharingOperationStatusable when successful
func (m *LinkRoleAbilities) GetLinkAllowsExternalUsers()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("linkAllowsExternalUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// GetLinkExpiration gets the linkExpiration property value. Indicates if links must expire, meaning the link is no longer usable after a specified time. If link expiration is enabled, a default link expiration time is provided.
// returns a SharingLinkExpirationStatusable when successful
func (m *LinkRoleAbilities) GetLinkExpiration()(SharingLinkExpirationStatusable) {
    val, err := m.GetBackingStore().Get("linkExpiration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingLinkExpirationStatusable)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *LinkRoleAbilities) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRetrieveLink gets the retrieveLink property value. Indicates if links of this role can be retrieved.
// returns a SharingOperationStatusable when successful
func (m *LinkRoleAbilities) GetRetrieveLink()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("retrieveLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// GetUpdateLink gets the updateLink property value. Indicates if links of this role can be updated.
// returns a SharingOperationStatusable when successful
func (m *LinkRoleAbilities) GetUpdateLink()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("updateLink")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *LinkRoleAbilities) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("addExistingExternalUsers", m.GetAddExistingExternalUsers())
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
        err := writer.WriteObjectValue("applyVariants", m.GetApplyVariants())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("createLink", m.GetCreateLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deleteLink", m.GetDeleteLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("linkAllowsExternalUsers", m.GetLinkAllowsExternalUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("linkExpiration", m.GetLinkExpiration())
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
        err := writer.WriteObjectValue("retrieveLink", m.GetRetrieveLink())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("updateLink", m.GetUpdateLink())
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
// SetAddExistingExternalUsers sets the addExistingExternalUsers property value. Indicates if the current user can add existing external user recipients to this sharing link.
func (m *LinkRoleAbilities) SetAddExistingExternalUsers(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("addExistingExternalUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LinkRoleAbilities) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAddNewExternalUsers sets the addNewExternalUsers property value. Indicates if the current user can add new external user recipients to this sharing link.
func (m *LinkRoleAbilities) SetAddNewExternalUsers(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("addNewExternalUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetApplyVariants sets the applyVariants property value. Indicates the status of the potential sharing link variants. If selected, it generates a separate sharing link from the sharing link that would otherwise be generated without the variant, yet with an identical role and scope.
func (m *LinkRoleAbilities) SetApplyVariants(value SharingLinkVariantsable)() {
    err := m.GetBackingStore().Set("applyVariants", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *LinkRoleAbilities) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCreateLink sets the createLink property value. Indicates if links of this role can be created.
func (m *LinkRoleAbilities) SetCreateLink(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("createLink", value)
    if err != nil {
        panic(err)
    }
}
// SetDeleteLink sets the deleteLink property value. Indicates if links of this role can be deleted.
func (m *LinkRoleAbilities) SetDeleteLink(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("deleteLink", value)
    if err != nil {
        panic(err)
    }
}
// SetLinkAllowsExternalUsers sets the linkAllowsExternalUsers property value. Indicates if this link can include external users.
func (m *LinkRoleAbilities) SetLinkAllowsExternalUsers(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("linkAllowsExternalUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetLinkExpiration sets the linkExpiration property value. Indicates if links must expire, meaning the link is no longer usable after a specified time. If link expiration is enabled, a default link expiration time is provided.
func (m *LinkRoleAbilities) SetLinkExpiration(value SharingLinkExpirationStatusable)() {
    err := m.GetBackingStore().Set("linkExpiration", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *LinkRoleAbilities) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRetrieveLink sets the retrieveLink property value. Indicates if links of this role can be retrieved.
func (m *LinkRoleAbilities) SetRetrieveLink(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("retrieveLink", value)
    if err != nil {
        panic(err)
    }
}
// SetUpdateLink sets the updateLink property value. Indicates if links of this role can be updated.
func (m *LinkRoleAbilities) SetUpdateLink(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("updateLink", value)
    if err != nil {
        panic(err)
    }
}
type LinkRoleAbilitiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddExistingExternalUsers()(SharingOperationStatusable)
    GetAddNewExternalUsers()(SharingOperationStatusable)
    GetApplyVariants()(SharingLinkVariantsable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCreateLink()(SharingOperationStatusable)
    GetDeleteLink()(SharingOperationStatusable)
    GetLinkAllowsExternalUsers()(SharingOperationStatusable)
    GetLinkExpiration()(SharingLinkExpirationStatusable)
    GetOdataType()(*string)
    GetRetrieveLink()(SharingOperationStatusable)
    GetUpdateLink()(SharingOperationStatusable)
    SetAddExistingExternalUsers(value SharingOperationStatusable)()
    SetAddNewExternalUsers(value SharingOperationStatusable)()
    SetApplyVariants(value SharingLinkVariantsable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCreateLink(value SharingOperationStatusable)()
    SetDeleteLink(value SharingOperationStatusable)()
    SetLinkAllowsExternalUsers(value SharingOperationStatusable)()
    SetLinkExpiration(value SharingLinkExpirationStatusable)()
    SetOdataType(value *string)()
    SetRetrieveLink(value SharingOperationStatusable)()
    SetUpdateLink(value SharingOperationStatusable)()
}
