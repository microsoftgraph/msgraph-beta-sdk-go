package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// TenantDetailedInformation provides operations to manage the collection of accessReview entities.
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
    odataTypeValue := "#microsoft.graph.managedTenants.tenantDetailedInformation";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateTenantDetailedInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantDetailedInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantDetailedInformation(), nil
}
// GetCity gets the city property value. The city where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.city
    }
}
// GetCountryCode gets the countryCode property value. The code for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCountryCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryCode
    }
}
// GetCountryName gets the countryName property value. The name for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetCountryName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.countryName
    }
}
// GetDefaultDomainName gets the defaultDomainName property value. The default domain name for the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetDefaultDomainName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.defaultDomainName
    }
}
// GetDisplayName gets the displayName property value. The display name for the managed tenant.
func (m *TenantDetailedInformation) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
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
    if m == nil {
        return nil
    } else {
        return m.industryName
    }
}
// GetRegion gets the region property value. The region where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) GetRegion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.region
    }
}
// GetSegmentName gets the segmentName property value. The business segment associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetSegmentName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.segmentName
    }
}
// GetTenantId gets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant.
func (m *TenantDetailedInformation) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetVerticalName gets the verticalName property value. The vertical associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) GetVerticalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.verticalName
    }
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
    if m != nil {
        m.city = value
    }
}
// SetCountryCode sets the countryCode property value. The code for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) SetCountryCode(value *string)() {
    if m != nil {
        m.countryCode = value
    }
}
// SetCountryName sets the countryName property value. The name for the country where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) SetCountryName(value *string)() {
    if m != nil {
        m.countryName = value
    }
}
// SetDefaultDomainName sets the defaultDomainName property value. The default domain name for the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetDefaultDomainName(value *string)() {
    if m != nil {
        m.defaultDomainName = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the managed tenant.
func (m *TenantDetailedInformation) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetIndustryName sets the industryName property value. The business industry associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetIndustryName(value *string)() {
    if m != nil {
        m.industryName = value
    }
}
// SetRegion sets the region property value. The region where the managed tenant is located. Optional. Read-only.
func (m *TenantDetailedInformation) SetRegion(value *string)() {
    if m != nil {
        m.region = value
    }
}
// SetSegmentName sets the segmentName property value. The business segment associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetSegmentName(value *string)() {
    if m != nil {
        m.segmentName = value
    }
}
// SetTenantId sets the tenantId property value. The Azure Active Directory tenant identifier for the managed tenant.
func (m *TenantDetailedInformation) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
// SetVerticalName sets the verticalName property value. The vertical associated with the managed tenant. Optional. Read-only.
func (m *TenantDetailedInformation) SetVerticalName(value *string)() {
    if m != nil {
        m.verticalName = value
    }
}
