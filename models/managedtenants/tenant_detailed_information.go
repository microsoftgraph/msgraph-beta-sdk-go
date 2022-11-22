package managedtenants

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TenantDetailedInformation provides operations to manage the collection of accessReviewDecision entities.
type TenantDetailedInformation struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The city where the managed tenant is located. Optional. Read-only.
    city *string
    // The code for the country where the managed tenant is located. Optional. Read-only.
    countryCode *string
    // The name for the country where the managed tenant is located. Optional. Read-only.
    countryName *string
    // The default domain name for the managed tenant. Optional. Read-only.
    defaultDomainName *string
    // The display name for the managed tenant.
    displayName *string
    // The business industry associated with the managed tenant. Optional. Read-only.
    industryName *string
    // The region where the managed tenant is located. Optional. Read-only.
    region *string
    // The business segment associated with the managed tenant. Optional. Read-only.
    segmentName *string
    // The Azure Active Directory tenant identifier for the managed tenant.
    tenantId *string
    // The vertical associated with the managed tenant. Optional. Read-only.
    verticalName *string
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
    return m.city
}
// GetCountryCode gets the countryCode property value. The code for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCountryCode()(*string) {
    return m.countryCode
}
// GetCountryName gets the countryName property value. The name for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCountryName()(*string) {
    return m.countryName
}
// GetDefaultDomainName gets the defaultDomainName property value. The default domain name for the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetDefaultDomainName()(*string) {
    return m.defaultDomainName
}
// GetDisplayName gets the displayName property value. The display name for the managed tenant.
func (m *TenantDetailedInformation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantDetailedInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["city"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCity)
    res["countryCode"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCountryCode)
    res["countryName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCountryName)
    res["defaultDomainName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDefaultDomainName)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["industryName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIndustryName)
    res["region"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetRegion)
    res["segmentName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetSegmentName)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    res["verticalName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetVerticalName)
    return res
}
// GetIndustryName gets the industryName property value. The business industry associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetIndustryName()(*string) {
    return m.industryName
}
// GetRegion gets the region property value. The region where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetRegion()(*string) {
    return m.region
}
// GetSegmentName gets the segmentName property value. The business segment associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetSegmentName()(*string) {
    return m.segmentName
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant.
func (m *TenantDetailedInformation) GetTenantId()(*string) {
    return m.tenantId
}
// GetVerticalName gets the verticalName property value. The vertical associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetVerticalName()(*string) {
    return m.verticalName
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
    m.city = value
}
// SetCountryCode sets the countryCode property value. The code for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) SetCountryCode(value *string)() {
    m.countryCode = value
}
// SetCountryName sets the countryName property value. The name for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) SetCountryName(value *string)() {
    m.countryName = value
}
// SetDefaultDomainName sets the defaultDomainName property value. The default domain name for the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetDefaultDomainName(value *string)() {
    m.defaultDomainName = value
}
// SetDisplayName sets the displayName property value. The display name for the managed tenant.
func (m *TenantDetailedInformation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIndustryName sets the industryName property value. The business industry associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetIndustryName(value *string)() {
    m.industryName = value
}
// SetRegion sets the region property value. The region where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) SetRegion(value *string)() {
    m.region = value
}
// SetSegmentName sets the segmentName property value. The business segment associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetSegmentName(value *string)() {
    m.segmentName = value
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant.
func (m *TenantDetailedInformation) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetVerticalName sets the verticalName property value. The vertical associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetVerticalName(value *string)() {
    m.verticalName = value
}
