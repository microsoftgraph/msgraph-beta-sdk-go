package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PartnerInformation 
type PartnerInformation struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPartnerInformation instantiates a new partnerInformation and sets the default values.
func NewPartnerInformation()(*PartnerInformation) {
    m := &PartnerInformation{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePartnerInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePartnerInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPartnerInformation(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PartnerInformation) GetAdditionalData()(map[string]any) {
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
func (m *PartnerInformation) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCommerceUrl gets the commerceUrl property value. The commerceUrl property
func (m *PartnerInformation) GetCommerceUrl()(*string) {
    val, err := m.GetBackingStore().Get("commerceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCompanyName gets the companyName property value. The companyName property
func (m *PartnerInformation) GetCompanyName()(*string) {
    val, err := m.GetBackingStore().Get("companyName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCompanyType gets the companyType property value. The companyType property
func (m *PartnerInformation) GetCompanyType()(*PartnerTenantType) {
    val, err := m.GetBackingStore().Get("companyType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PartnerTenantType)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PartnerInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commerceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommerceUrl(val)
        }
        return nil
    }
    res["companyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
        }
        return nil
    }
    res["companyType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePartnerTenantType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyType(val.(*PartnerTenantType))
        }
        return nil
    }
    res["helpUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpUrl(val)
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
    res["partnerTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerTenantId(val)
        }
        return nil
    }
    res["supportEmails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSupportEmails(res)
        }
        return nil
    }
    res["supportTelephones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetSupportTelephones(res)
        }
        return nil
    }
    res["supportUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportUrl(val)
        }
        return nil
    }
    return res
}
// GetHelpUrl gets the helpUrl property value. The helpUrl property
func (m *PartnerInformation) GetHelpUrl()(*string) {
    val, err := m.GetBackingStore().Get("helpUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PartnerInformation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPartnerTenantId gets the partnerTenantId property value. The partnerTenantId property
func (m *PartnerInformation) GetPartnerTenantId()(*string) {
    val, err := m.GetBackingStore().Get("partnerTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSupportEmails gets the supportEmails property value. The supportEmails property
func (m *PartnerInformation) GetSupportEmails()([]string) {
    val, err := m.GetBackingStore().Get("supportEmails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSupportTelephones gets the supportTelephones property value. The supportTelephones property
func (m *PartnerInformation) GetSupportTelephones()([]string) {
    val, err := m.GetBackingStore().Get("supportTelephones")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSupportUrl gets the supportUrl property value. The supportUrl property
func (m *PartnerInformation) GetSupportUrl()(*string) {
    val, err := m.GetBackingStore().Get("supportUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PartnerInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("commerceUrl", m.GetCommerceUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    if m.GetCompanyType() != nil {
        cast := (*m.GetCompanyType()).String()
        err := writer.WriteStringValue("companyType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("helpUrl", m.GetHelpUrl())
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
        err := writer.WriteStringValue("partnerTenantId", m.GetPartnerTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetSupportEmails() != nil {
        err := writer.WriteCollectionOfStringValues("supportEmails", m.GetSupportEmails())
        if err != nil {
            return err
        }
    }
    if m.GetSupportTelephones() != nil {
        err := writer.WriteCollectionOfStringValues("supportTelephones", m.GetSupportTelephones())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("supportUrl", m.GetSupportUrl())
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
func (m *PartnerInformation) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *PartnerInformation) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCommerceUrl sets the commerceUrl property value. The commerceUrl property
func (m *PartnerInformation) SetCommerceUrl(value *string)() {
    err := m.GetBackingStore().Set("commerceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCompanyName sets the companyName property value. The companyName property
func (m *PartnerInformation) SetCompanyName(value *string)() {
    err := m.GetBackingStore().Set("companyName", value)
    if err != nil {
        panic(err)
    }
}
// SetCompanyType sets the companyType property value. The companyType property
func (m *PartnerInformation) SetCompanyType(value *PartnerTenantType)() {
    err := m.GetBackingStore().Set("companyType", value)
    if err != nil {
        panic(err)
    }
}
// SetHelpUrl sets the helpUrl property value. The helpUrl property
func (m *PartnerInformation) SetHelpUrl(value *string)() {
    err := m.GetBackingStore().Set("helpUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PartnerInformation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPartnerTenantId sets the partnerTenantId property value. The partnerTenantId property
func (m *PartnerInformation) SetPartnerTenantId(value *string)() {
    err := m.GetBackingStore().Set("partnerTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportEmails sets the supportEmails property value. The supportEmails property
func (m *PartnerInformation) SetSupportEmails(value []string)() {
    err := m.GetBackingStore().Set("supportEmails", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportTelephones sets the supportTelephones property value. The supportTelephones property
func (m *PartnerInformation) SetSupportTelephones(value []string)() {
    err := m.GetBackingStore().Set("supportTelephones", value)
    if err != nil {
        panic(err)
    }
}
// SetSupportUrl sets the supportUrl property value. The supportUrl property
func (m *PartnerInformation) SetSupportUrl(value *string)() {
    err := m.GetBackingStore().Set("supportUrl", value)
    if err != nil {
        panic(err)
    }
}
// PartnerInformationable 
type PartnerInformationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCommerceUrl()(*string)
    GetCompanyName()(*string)
    GetCompanyType()(*PartnerTenantType)
    GetHelpUrl()(*string)
    GetOdataType()(*string)
    GetPartnerTenantId()(*string)
    GetSupportEmails()([]string)
    GetSupportTelephones()([]string)
    GetSupportUrl()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCommerceUrl(value *string)()
    SetCompanyName(value *string)()
    SetCompanyType(value *PartnerTenantType)()
    SetHelpUrl(value *string)()
    SetOdataType(value *string)()
    SetPartnerTenantId(value *string)()
    SetSupportEmails(value []string)()
    SetSupportTelephones(value []string)()
    SetSupportUrl(value *string)()
}
