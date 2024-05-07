package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type SharingLinkVariants struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSharingLinkVariants instantiates a new SharingLinkVariants and sets the default values.
func NewSharingLinkVariants()(*SharingLinkVariants) {
    m := &SharingLinkVariants{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSharingLinkVariantsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSharingLinkVariantsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharingLinkVariants(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SharingLinkVariants) GetAdditionalData()(map[string]any) {
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
// GetAddressBarLinkPermission gets the addressBarLinkPermission property value. Indicates the most permissive role with which an address bar link can be created. The possible values are: none, view, edit, manageList, review, restrictedView, submitOnly, unknownFutureValue.
// returns a *SharingRole when successful
func (m *SharingLinkVariants) GetAddressBarLinkPermission()(*SharingRole) {
    val, err := m.GetBackingStore().Get("addressBarLinkPermission")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SharingRole)
    }
    return nil
}
// GetAllowEmbed gets the allowEmbed property value. Indicates whether a link can be embedded.
// returns a SharingOperationStatusable when successful
func (m *SharingLinkVariants) GetAllowEmbed()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("allowEmbed")
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
func (m *SharingLinkVariants) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SharingLinkVariants) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addressBarLinkPermission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSharingRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressBarLinkPermission(val.(*SharingRole))
        }
        return nil
    }
    res["allowEmbed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowEmbed(val.(SharingOperationStatusable))
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
    res["passwordProtected"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordProtected(val.(SharingOperationStatusable))
        }
        return nil
    }
    res["requiresAuthentication"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingOperationStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiresAuthentication(val.(SharingOperationStatusable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
// returns a *string when successful
func (m *SharingLinkVariants) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPasswordProtected gets the passwordProtected property value. Indicates whether a link can be password protected, meaning that link users would need to enter a password to access the item for which the sharing link is produced. Creating a passwordProtected link for the first time requires providing a password.
// returns a SharingOperationStatusable when successful
func (m *SharingLinkVariants) GetPasswordProtected()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("passwordProtected")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// GetRequiresAuthentication gets the requiresAuthentication property value. Indicates whether a link requires identity authentication for recipients. Users can be verified through either an email address or identity.
// returns a SharingOperationStatusable when successful
func (m *SharingLinkVariants) GetRequiresAuthentication()(SharingOperationStatusable) {
    val, err := m.GetBackingStore().Get("requiresAuthentication")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SharingOperationStatusable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SharingLinkVariants) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAddressBarLinkPermission() != nil {
        cast := (*m.GetAddressBarLinkPermission()).String()
        err := writer.WriteStringValue("addressBarLinkPermission", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("allowEmbed", m.GetAllowEmbed())
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
        err := writer.WriteObjectValue("passwordProtected", m.GetPasswordProtected())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("requiresAuthentication", m.GetRequiresAuthentication())
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
func (m *SharingLinkVariants) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAddressBarLinkPermission sets the addressBarLinkPermission property value. Indicates the most permissive role with which an address bar link can be created. The possible values are: none, view, edit, manageList, review, restrictedView, submitOnly, unknownFutureValue.
func (m *SharingLinkVariants) SetAddressBarLinkPermission(value *SharingRole)() {
    err := m.GetBackingStore().Set("addressBarLinkPermission", value)
    if err != nil {
        panic(err)
    }
}
// SetAllowEmbed sets the allowEmbed property value. Indicates whether a link can be embedded.
func (m *SharingLinkVariants) SetAllowEmbed(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("allowEmbed", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *SharingLinkVariants) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SharingLinkVariants) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPasswordProtected sets the passwordProtected property value. Indicates whether a link can be password protected, meaning that link users would need to enter a password to access the item for which the sharing link is produced. Creating a passwordProtected link for the first time requires providing a password.
func (m *SharingLinkVariants) SetPasswordProtected(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("passwordProtected", value)
    if err != nil {
        panic(err)
    }
}
// SetRequiresAuthentication sets the requiresAuthentication property value. Indicates whether a link requires identity authentication for recipients. Users can be verified through either an email address or identity.
func (m *SharingLinkVariants) SetRequiresAuthentication(value SharingOperationStatusable)() {
    err := m.GetBackingStore().Set("requiresAuthentication", value)
    if err != nil {
        panic(err)
    }
}
type SharingLinkVariantsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddressBarLinkPermission()(*SharingRole)
    GetAllowEmbed()(SharingOperationStatusable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetOdataType()(*string)
    GetPasswordProtected()(SharingOperationStatusable)
    GetRequiresAuthentication()(SharingOperationStatusable)
    SetAddressBarLinkPermission(value *SharingRole)()
    SetAllowEmbed(value SharingOperationStatusable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetOdataType(value *string)()
    SetPasswordProtected(value SharingOperationStatusable)()
    SetRequiresAuthentication(value SharingOperationStatusable)()
}
