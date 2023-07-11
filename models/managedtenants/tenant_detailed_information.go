package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TenantDetailedInformation 
type TenantDetailedInformation struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewTenantDetailedInformation instantiates a new tenantDetailedInformation and sets the default values.
func NewTenantDetailedInformation()(*TenantDetailedInformation) {
    m := &TenantDetailedInformation{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateTenantDetailedInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantDetailedInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantDetailedInformation(), nil
}
// GetCity gets the city property value. The city where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCity()(*string) {
    val, err := m.GetBackingStore().Get("city")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCountryCode gets the countryCode property value. The code for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCountryCode()(*string) {
    val, err := m.GetBackingStore().Get("countryCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCountryName gets the countryName property value. The name for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCountryName()(*string) {
    val, err := m.GetBackingStore().Get("countryName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDefaultDomainName gets the defaultDomainName property value. The default domain name for the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetDefaultDomainName()(*string) {
    val, err := m.GetBackingStore().Get("defaultDomainName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the managed tenant.
func (m *TenantDetailedInformation) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantDetailedInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryCode(val)
        }
        return nil
    }
    res["countryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryName(val)
        }
        return nil
    }
    res["defaultDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDomainName(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["industryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndustryName(val)
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
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["segmentName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSegmentName(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["verticalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerticalName(val)
        }
        return nil
    }
    return res
}
// GetIndustryName gets the industryName property value. The business industry associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetIndustryName()(*string) {
    val, err := m.GetBackingStore().Get("industryName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TenantDetailedInformation) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRegion gets the region property value. The region where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetRegion()(*string) {
    val, err := m.GetBackingStore().Get("region")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSegmentName gets the segmentName property value. The business segment associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetSegmentName()(*string) {
    val, err := m.GetBackingStore().Get("segmentName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant.
func (m *TenantDetailedInformation) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVerticalName gets the verticalName property value. The vertical associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetVerticalName()(*string) {
    val, err := m.GetBackingStore().Get("verticalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *TenantDetailedInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryCode", m.GetCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryName", m.GetCountryName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultDomainName", m.GetDefaultDomainName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("industryName", m.GetIndustryName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("segmentName", m.GetSegmentName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("verticalName", m.GetVerticalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCity sets the city property value. The city where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) SetCity(value *string)() {
    err := m.GetBackingStore().Set("city", value)
    if err != nil {
        panic(err)
    }
}
// SetCountryCode sets the countryCode property value. The code for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) SetCountryCode(value *string)() {
    err := m.GetBackingStore().Set("countryCode", value)
    if err != nil {
        panic(err)
    }
}
// SetCountryName sets the countryName property value. The name for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) SetCountryName(value *string)() {
    err := m.GetBackingStore().Set("countryName", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultDomainName sets the defaultDomainName property value. The default domain name for the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetDefaultDomainName(value *string)() {
    err := m.GetBackingStore().Set("defaultDomainName", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the managed tenant.
func (m *TenantDetailedInformation) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetIndustryName sets the industryName property value. The business industry associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetIndustryName(value *string)() {
    err := m.GetBackingStore().Set("industryName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantDetailedInformation) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRegion sets the region property value. The region where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) SetRegion(value *string)() {
    err := m.GetBackingStore().Set("region", value)
    if err != nil {
        panic(err)
    }
}
// SetSegmentName sets the segmentName property value. The business segment associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetSegmentName(value *string)() {
    err := m.GetBackingStore().Set("segmentName", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant.
func (m *TenantDetailedInformation) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetVerticalName sets the verticalName property value. The vertical associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetVerticalName(value *string)() {
    err := m.GetBackingStore().Set("verticalName", value)
    if err != nil {
        panic(err)
    }
}
// TenantDetailedInformationable 
type TenantDetailedInformationable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCity()(*string)
    GetCountryCode()(*string)
    GetCountryName()(*string)
    GetDefaultDomainName()(*string)
    GetDisplayName()(*string)
    GetIndustryName()(*string)
    GetOdataType()(*string)
    GetRegion()(*string)
    GetSegmentName()(*string)
    GetTenantId()(*string)
    GetVerticalName()(*string)
    SetCity(value *string)()
    SetCountryCode(value *string)()
    SetCountryName(value *string)()
    SetDefaultDomainName(value *string)()
    SetDisplayName(value *string)()
    SetIndustryName(value *string)()
    SetOdataType(value *string)()
    SetRegion(value *string)()
    SetSegmentName(value *string)()
    SetTenantId(value *string)()
    SetVerticalName(value *string)()
}
