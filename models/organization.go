package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Organization struct {
    DirectoryObject
}
// NewOrganization instantiates a new Organization and sets the default values.
func NewOrganization()(*Organization) {
    m := &Organization{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.organization"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOrganizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganization(), nil
}
// GetAssignedPlans gets the assignedPlans property value. The collection of service plans associated with the tenant. Not nullable.
// returns a []AssignedPlanable when successful
func (m *Organization) GetAssignedPlans()([]AssignedPlanable) {
    val, err := m.GetBackingStore().Get("assignedPlans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AssignedPlanable)
    }
    return nil
}
// GetBranding gets the branding property value. Resource to manage the default branding for the organization. Nullable.
// returns a OrganizationalBrandingable when successful
func (m *Organization) GetBranding()(OrganizationalBrandingable) {
    val, err := m.GetBackingStore().Get("branding")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OrganizationalBrandingable)
    }
    return nil
}
// GetBusinessPhones gets the businessPhones property value. Telephone number for the organization. Although this is a string collection, only one number can be set for this property.
// returns a []string when successful
func (m *Organization) GetBusinessPhones()([]string) {
    val, err := m.GetBackingStore().Get("businessPhones")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetCertificateBasedAuthConfiguration gets the certificateBasedAuthConfiguration property value. Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection.
// returns a []CertificateBasedAuthConfigurationable when successful
func (m *Organization) GetCertificateBasedAuthConfiguration()([]CertificateBasedAuthConfigurationable) {
    val, err := m.GetBackingStore().Get("certificateBasedAuthConfiguration")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CertificateBasedAuthConfigurationable)
    }
    return nil
}
// GetCertificateConnectorSetting gets the certificateConnectorSetting property value. Certificate connector setting.
// returns a CertificateConnectorSettingable when successful
func (m *Organization) GetCertificateConnectorSetting()(CertificateConnectorSettingable) {
    val, err := m.GetBackingStore().Get("certificateConnectorSetting")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CertificateConnectorSettingable)
    }
    return nil
}
// GetCity gets the city property value. City name of the address for the organization.
// returns a *string when successful
func (m *Organization) GetCity()(*string) {
    val, err := m.GetBackingStore().Get("city")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCountry gets the country property value. Country/region name of the address for the organization.
// returns a *string when successful
func (m *Organization) GetCountry()(*string) {
    val, err := m.GetBackingStore().Get("country")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCountryLetterCode gets the countryLetterCode property value. Country or region abbreviation for the organization in ISO 3166-2 format.
// returns a *string when successful
func (m *Organization) GetCountryLetterCode()(*string) {
    val, err := m.GetBackingStore().Get("countryLetterCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// returns a *Time when successful
func (m *Organization) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDefaultUsageLocation gets the defaultUsageLocation property value. Two-letter ISO 3166 country code indicating the default service usage location of an organization.
// returns a *string when successful
func (m *Organization) GetDefaultUsageLocation()(*string) {
    val, err := m.GetBackingStore().Get("defaultUsageLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDirectorySizeQuota gets the directorySizeQuota property value. The directory size quota information of an organization.
// returns a DirectorySizeQuotaable when successful
func (m *Organization) GetDirectorySizeQuota()(DirectorySizeQuotaable) {
    val, err := m.GetBackingStore().Get("directorySizeQuota")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DirectorySizeQuotaable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The display name for the tenant.
// returns a *string when successful
func (m *Organization) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the organization resource. Nullable.
// returns a []Extensionable when successful
func (m *Organization) GetExtensions()([]Extensionable) {
    val, err := m.GetBackingStore().Get("extensions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Extensionable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Organization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["assignedPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedPlanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AssignedPlanable)
                }
            }
            m.SetAssignedPlans(res)
        }
        return nil
    }
    res["branding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationalBrandingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranding(val.(OrganizationalBrandingable))
        }
        return nil
    }
    res["businessPhones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetBusinessPhones(res)
        }
        return nil
    }
    res["certificateBasedAuthConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificateBasedAuthConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateBasedAuthConfigurationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CertificateBasedAuthConfigurationable)
                }
            }
            m.SetCertificateBasedAuthConfiguration(res)
        }
        return nil
    }
    res["certificateConnectorSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificateConnectorSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateConnectorSetting(val.(CertificateConnectorSettingable))
        }
        return nil
    }
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
    res["country"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
        }
        return nil
    }
    res["countryLetterCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryLetterCode(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["defaultUsageLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultUsageLocation(val)
        }
        return nil
    }
    res["directorySizeQuota"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectorySizeQuotaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectorySizeQuota(val.(DirectorySizeQuotaable))
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
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Extensionable)
                }
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["isMultipleDataLocationsForServicesEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMultipleDataLocationsForServicesEnabled(val)
        }
        return nil
    }
    res["marketingNotificationEmails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetMarketingNotificationEmails(res)
        }
        return nil
    }
    res["mobileDeviceManagementAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMdmAuthority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileDeviceManagementAuthority(val.(*MdmAuthority))
        }
        return nil
    }
    res["onPremisesLastPasswordSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastPasswordSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["partnerInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePartnerInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerInformation(val.(PartnerInformationable))
        }
        return nil
    }
    res["partnerTenantType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePartnerTenantType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerTenantType(val.(*PartnerTenantType))
        }
        return nil
    }
    res["postalCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostalCode(val)
        }
        return nil
    }
    res["preferredLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguage(val)
        }
        return nil
    }
    res["privacyProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivacyProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyProfile(val.(PrivacyProfileable))
        }
        return nil
    }
    res["provisionedPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisionedPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisionedPlanable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ProvisionedPlanable)
                }
            }
            m.SetProvisionedPlans(res)
        }
        return nil
    }
    res["securityComplianceNotificationMails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSecurityComplianceNotificationMails(res)
        }
        return nil
    }
    res["securityComplianceNotificationPhones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetSecurityComplianceNotificationPhones(res)
        }
        return nil
    }
    res["settings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(OrganizationSettingsable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["street"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreet(val)
        }
        return nil
    }
    res["technicalNotificationMails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetTechnicalNotificationMails(res)
        }
        return nil
    }
    res["verifiedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVerifiedDomainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VerifiedDomainable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VerifiedDomainable)
                }
            }
            m.SetVerifiedDomains(res)
        }
        return nil
    }
    return res
}
// GetIsMultipleDataLocationsForServicesEnabled gets the isMultipleDataLocationsForServicesEnabled property value. true if organization is Multi-Geo enabled; false if organization is not Multi-Geo enabled; null (default). Read-only. For more information, see OneDrive Online Multi-Geo.
// returns a *bool when successful
func (m *Organization) GetIsMultipleDataLocationsForServicesEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isMultipleDataLocationsForServicesEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMarketingNotificationEmails gets the marketingNotificationEmails property value. Not nullable.
// returns a []string when successful
func (m *Organization) GetMarketingNotificationEmails()([]string) {
    val, err := m.GetBackingStore().Get("marketingNotificationEmails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetMobileDeviceManagementAuthority gets the mobileDeviceManagementAuthority property value. Mobile device management authority.
// returns a *MdmAuthority when successful
func (m *Organization) GetMobileDeviceManagementAuthority()(*MdmAuthority) {
    val, err := m.GetBackingStore().Get("mobileDeviceManagementAuthority")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*MdmAuthority)
    }
    return nil
}
// GetOnPremisesLastPasswordSyncDateTime gets the onPremisesLastPasswordSyncDateTime property value. The last time a password sync request was received for the tenant.
// returns a *Time when successful
func (m *Organization) GetOnPremisesLastPasswordSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("onPremisesLastPasswordSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// returns a *Time when successful
func (m *Organization) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("onPremisesLastSyncDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; Nullable. null if this object has never been synced from an on-premises directory (default).
// returns a *bool when successful
func (m *Organization) GetOnPremisesSyncEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("onPremisesSyncEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetPartnerInformation gets the partnerInformation property value. The partnerInformation property
// returns a PartnerInformationable when successful
func (m *Organization) GetPartnerInformation()(PartnerInformationable) {
    val, err := m.GetBackingStore().Get("partnerInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PartnerInformationable)
    }
    return nil
}
// GetPartnerTenantType gets the partnerTenantType property value. The type of partnership this tenant has with Microsoft. The possible values are: microsoftSupport, syndicatePartner, breadthPartner, breadthPartnerDelegatedAdmin, resellerPartnerDelegatedAdmin, valueAddedResellerPartnerDelegatedAdmin, unknownFutureValue. Nullable. For more information about the possible types, see partnerTenantType values.
// returns a *PartnerTenantType when successful
func (m *Organization) GetPartnerTenantType()(*PartnerTenantType) {
    val, err := m.GetBackingStore().Get("partnerTenantType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*PartnerTenantType)
    }
    return nil
}
// GetPostalCode gets the postalCode property value. Postal code of the address for the organization.
// returns a *string when successful
func (m *Organization) GetPostalCode()(*string) {
    val, err := m.GetBackingStore().Get("postalCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPreferredLanguage gets the preferredLanguage property value. The preferred language for the organization. Should follow ISO 639-1 Code; for example en.
// returns a *string when successful
func (m *Organization) GetPreferredLanguage()(*string) {
    val, err := m.GetBackingStore().Get("preferredLanguage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrivacyProfile gets the privacyProfile property value. The privacy profile of an organization.
// returns a PrivacyProfileable when successful
func (m *Organization) GetPrivacyProfile()(PrivacyProfileable) {
    val, err := m.GetBackingStore().Get("privacyProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PrivacyProfileable)
    }
    return nil
}
// GetProvisionedPlans gets the provisionedPlans property value. Not nullable.
// returns a []ProvisionedPlanable when successful
func (m *Organization) GetProvisionedPlans()([]ProvisionedPlanable) {
    val, err := m.GetBackingStore().Get("provisionedPlans")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ProvisionedPlanable)
    }
    return nil
}
// GetSecurityComplianceNotificationMails gets the securityComplianceNotificationMails property value. Not nullable.
// returns a []string when successful
func (m *Organization) GetSecurityComplianceNotificationMails()([]string) {
    val, err := m.GetBackingStore().Get("securityComplianceNotificationMails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSecurityComplianceNotificationPhones gets the securityComplianceNotificationPhones property value. Not nullable.
// returns a []string when successful
func (m *Organization) GetSecurityComplianceNotificationPhones()([]string) {
    val, err := m.GetBackingStore().Get("securityComplianceNotificationPhones")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetSettings gets the settings property value. Retrieve the properties and relationships of organizationSettings object. Nullable.
// returns a OrganizationSettingsable when successful
func (m *Organization) GetSettings()(OrganizationSettingsable) {
    val, err := m.GetBackingStore().Get("settings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(OrganizationSettingsable)
    }
    return nil
}
// GetState gets the state property value. State name of the address for the organization.
// returns a *string when successful
func (m *Organization) GetState()(*string) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStreet gets the street property value. Street name of the address for organization.
// returns a *string when successful
func (m *Organization) GetStreet()(*string) {
    val, err := m.GetBackingStore().Get("street")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTechnicalNotificationMails gets the technicalNotificationMails property value. Not nullable.
// returns a []string when successful
func (m *Organization) GetTechnicalNotificationMails()([]string) {
    val, err := m.GetBackingStore().Get("technicalNotificationMails")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetVerifiedDomains gets the verifiedDomains property value. The collection of domains associated with this tenant. Not nullable.
// returns a []VerifiedDomainable when successful
func (m *Organization) GetVerifiedDomains()([]VerifiedDomainable) {
    val, err := m.GetBackingStore().Get("verifiedDomains")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VerifiedDomainable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Organization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedPlans()))
        for i, v := range m.GetAssignedPlans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("assignedPlans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("branding", m.GetBranding())
        if err != nil {
            return err
        }
    }
    if m.GetBusinessPhones() != nil {
        err = writer.WriteCollectionOfStringValues("businessPhones", m.GetBusinessPhones())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateBasedAuthConfiguration() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificateBasedAuthConfiguration()))
        for i, v := range m.GetCertificateBasedAuthConfiguration() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("certificateBasedAuthConfiguration", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("certificateConnectorSetting", m.GetCertificateConnectorSetting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("country", m.GetCountry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryLetterCode", m.GetCountryLetterCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultUsageLocation", m.GetDefaultUsageLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("directorySizeQuota", m.GetDirectorySizeQuota())
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
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isMultipleDataLocationsForServicesEnabled", m.GetIsMultipleDataLocationsForServicesEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetMarketingNotificationEmails() != nil {
        err = writer.WriteCollectionOfStringValues("marketingNotificationEmails", m.GetMarketingNotificationEmails())
        if err != nil {
            return err
        }
    }
    if m.GetMobileDeviceManagementAuthority() != nil {
        cast := (*m.GetMobileDeviceManagementAuthority()).String()
        err = writer.WriteStringValue("mobileDeviceManagementAuthority", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastPasswordSyncDateTime", m.GetOnPremisesLastPasswordSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("partnerInformation", m.GetPartnerInformation())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerTenantType() != nil {
        cast := (*m.GetPartnerTenantType()).String()
        err = writer.WriteStringValue("partnerTenantType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredLanguage", m.GetPreferredLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("privacyProfile", m.GetPrivacyProfile())
        if err != nil {
            return err
        }
    }
    if m.GetProvisionedPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProvisionedPlans()))
        for i, v := range m.GetProvisionedPlans() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("provisionedPlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityComplianceNotificationMails() != nil {
        err = writer.WriteCollectionOfStringValues("securityComplianceNotificationMails", m.GetSecurityComplianceNotificationMails())
        if err != nil {
            return err
        }
    }
    if m.GetSecurityComplianceNotificationPhones() != nil {
        err = writer.WriteCollectionOfStringValues("securityComplianceNotificationPhones", m.GetSecurityComplianceNotificationPhones())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("street", m.GetStreet())
        if err != nil {
            return err
        }
    }
    if m.GetTechnicalNotificationMails() != nil {
        err = writer.WriteCollectionOfStringValues("technicalNotificationMails", m.GetTechnicalNotificationMails())
        if err != nil {
            return err
        }
    }
    if m.GetVerifiedDomains() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVerifiedDomains()))
        for i, v := range m.GetVerifiedDomains() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("verifiedDomains", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedPlans sets the assignedPlans property value. The collection of service plans associated with the tenant. Not nullable.
func (m *Organization) SetAssignedPlans(value []AssignedPlanable)() {
    err := m.GetBackingStore().Set("assignedPlans", value)
    if err != nil {
        panic(err)
    }
}
// SetBranding sets the branding property value. Resource to manage the default branding for the organization. Nullable.
func (m *Organization) SetBranding(value OrganizationalBrandingable)() {
    err := m.GetBackingStore().Set("branding", value)
    if err != nil {
        panic(err)
    }
}
// SetBusinessPhones sets the businessPhones property value. Telephone number for the organization. Although this is a string collection, only one number can be set for this property.
func (m *Organization) SetBusinessPhones(value []string)() {
    err := m.GetBackingStore().Set("businessPhones", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateBasedAuthConfiguration sets the certificateBasedAuthConfiguration property value. Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection.
func (m *Organization) SetCertificateBasedAuthConfiguration(value []CertificateBasedAuthConfigurationable)() {
    err := m.GetBackingStore().Set("certificateBasedAuthConfiguration", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateConnectorSetting sets the certificateConnectorSetting property value. Certificate connector setting.
func (m *Organization) SetCertificateConnectorSetting(value CertificateConnectorSettingable)() {
    err := m.GetBackingStore().Set("certificateConnectorSetting", value)
    if err != nil {
        panic(err)
    }
}
// SetCity sets the city property value. City name of the address for the organization.
func (m *Organization) SetCity(value *string)() {
    err := m.GetBackingStore().Set("city", value)
    if err != nil {
        panic(err)
    }
}
// SetCountry sets the country property value. Country/region name of the address for the organization.
func (m *Organization) SetCountry(value *string)() {
    err := m.GetBackingStore().Set("country", value)
    if err != nil {
        panic(err)
    }
}
// SetCountryLetterCode sets the countryLetterCode property value. Country or region abbreviation for the organization in ISO 3166-2 format.
func (m *Organization) SetCountryLetterCode(value *string)() {
    err := m.GetBackingStore().Set("countryLetterCode", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Organization) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDefaultUsageLocation sets the defaultUsageLocation property value. Two-letter ISO 3166 country code indicating the default service usage location of an organization.
func (m *Organization) SetDefaultUsageLocation(value *string)() {
    err := m.GetBackingStore().Set("defaultUsageLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetDirectorySizeQuota sets the directorySizeQuota property value. The directory size quota information of an organization.
func (m *Organization) SetDirectorySizeQuota(value DirectorySizeQuotaable)() {
    err := m.GetBackingStore().Set("directorySizeQuota", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The display name for the tenant.
func (m *Organization) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the organization resource. Nullable.
func (m *Organization) SetExtensions(value []Extensionable)() {
    err := m.GetBackingStore().Set("extensions", value)
    if err != nil {
        panic(err)
    }
}
// SetIsMultipleDataLocationsForServicesEnabled sets the isMultipleDataLocationsForServicesEnabled property value. true if organization is Multi-Geo enabled; false if organization is not Multi-Geo enabled; null (default). Read-only. For more information, see OneDrive Online Multi-Geo.
func (m *Organization) SetIsMultipleDataLocationsForServicesEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isMultipleDataLocationsForServicesEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetMarketingNotificationEmails sets the marketingNotificationEmails property value. Not nullable.
func (m *Organization) SetMarketingNotificationEmails(value []string)() {
    err := m.GetBackingStore().Set("marketingNotificationEmails", value)
    if err != nil {
        panic(err)
    }
}
// SetMobileDeviceManagementAuthority sets the mobileDeviceManagementAuthority property value. Mobile device management authority.
func (m *Organization) SetMobileDeviceManagementAuthority(value *MdmAuthority)() {
    err := m.GetBackingStore().Set("mobileDeviceManagementAuthority", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesLastPasswordSyncDateTime sets the onPremisesLastPasswordSyncDateTime property value. The last time a password sync request was received for the tenant.
func (m *Organization) SetOnPremisesLastPasswordSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("onPremisesLastPasswordSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Organization) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("onPremisesLastSyncDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced; Nullable. null if this object has never been synced from an on-premises directory (default).
func (m *Organization) SetOnPremisesSyncEnabled(value *bool)() {
    err := m.GetBackingStore().Set("onPremisesSyncEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetPartnerInformation sets the partnerInformation property value. The partnerInformation property
func (m *Organization) SetPartnerInformation(value PartnerInformationable)() {
    err := m.GetBackingStore().Set("partnerInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetPartnerTenantType sets the partnerTenantType property value. The type of partnership this tenant has with Microsoft. The possible values are: microsoftSupport, syndicatePartner, breadthPartner, breadthPartnerDelegatedAdmin, resellerPartnerDelegatedAdmin, valueAddedResellerPartnerDelegatedAdmin, unknownFutureValue. Nullable. For more information about the possible types, see partnerTenantType values.
func (m *Organization) SetPartnerTenantType(value *PartnerTenantType)() {
    err := m.GetBackingStore().Set("partnerTenantType", value)
    if err != nil {
        panic(err)
    }
}
// SetPostalCode sets the postalCode property value. Postal code of the address for the organization.
func (m *Organization) SetPostalCode(value *string)() {
    err := m.GetBackingStore().Set("postalCode", value)
    if err != nil {
        panic(err)
    }
}
// SetPreferredLanguage sets the preferredLanguage property value. The preferred language for the organization. Should follow ISO 639-1 Code; for example en.
func (m *Organization) SetPreferredLanguage(value *string)() {
    err := m.GetBackingStore().Set("preferredLanguage", value)
    if err != nil {
        panic(err)
    }
}
// SetPrivacyProfile sets the privacyProfile property value. The privacy profile of an organization.
func (m *Organization) SetPrivacyProfile(value PrivacyProfileable)() {
    err := m.GetBackingStore().Set("privacyProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisionedPlans sets the provisionedPlans property value. Not nullable.
func (m *Organization) SetProvisionedPlans(value []ProvisionedPlanable)() {
    err := m.GetBackingStore().Set("provisionedPlans", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityComplianceNotificationMails sets the securityComplianceNotificationMails property value. Not nullable.
func (m *Organization) SetSecurityComplianceNotificationMails(value []string)() {
    err := m.GetBackingStore().Set("securityComplianceNotificationMails", value)
    if err != nil {
        panic(err)
    }
}
// SetSecurityComplianceNotificationPhones sets the securityComplianceNotificationPhones property value. Not nullable.
func (m *Organization) SetSecurityComplianceNotificationPhones(value []string)() {
    err := m.GetBackingStore().Set("securityComplianceNotificationPhones", value)
    if err != nil {
        panic(err)
    }
}
// SetSettings sets the settings property value. Retrieve the properties and relationships of organizationSettings object. Nullable.
func (m *Organization) SetSettings(value OrganizationSettingsable)() {
    err := m.GetBackingStore().Set("settings", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. State name of the address for the organization.
func (m *Organization) SetState(value *string)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SetStreet sets the street property value. Street name of the address for organization.
func (m *Organization) SetStreet(value *string)() {
    err := m.GetBackingStore().Set("street", value)
    if err != nil {
        panic(err)
    }
}
// SetTechnicalNotificationMails sets the technicalNotificationMails property value. Not nullable.
func (m *Organization) SetTechnicalNotificationMails(value []string)() {
    err := m.GetBackingStore().Set("technicalNotificationMails", value)
    if err != nil {
        panic(err)
    }
}
// SetVerifiedDomains sets the verifiedDomains property value. The collection of domains associated with this tenant. Not nullable.
func (m *Organization) SetVerifiedDomains(value []VerifiedDomainable)() {
    err := m.GetBackingStore().Set("verifiedDomains", value)
    if err != nil {
        panic(err)
    }
}
type Organizationable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedPlans()([]AssignedPlanable)
    GetBranding()(OrganizationalBrandingable)
    GetBusinessPhones()([]string)
    GetCertificateBasedAuthConfiguration()([]CertificateBasedAuthConfigurationable)
    GetCertificateConnectorSetting()(CertificateConnectorSettingable)
    GetCity()(*string)
    GetCountry()(*string)
    GetCountryLetterCode()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDefaultUsageLocation()(*string)
    GetDirectorySizeQuota()(DirectorySizeQuotaable)
    GetDisplayName()(*string)
    GetExtensions()([]Extensionable)
    GetIsMultipleDataLocationsForServicesEnabled()(*bool)
    GetMarketingNotificationEmails()([]string)
    GetMobileDeviceManagementAuthority()(*MdmAuthority)
    GetOnPremisesLastPasswordSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOnPremisesSyncEnabled()(*bool)
    GetPartnerInformation()(PartnerInformationable)
    GetPartnerTenantType()(*PartnerTenantType)
    GetPostalCode()(*string)
    GetPreferredLanguage()(*string)
    GetPrivacyProfile()(PrivacyProfileable)
    GetProvisionedPlans()([]ProvisionedPlanable)
    GetSecurityComplianceNotificationMails()([]string)
    GetSecurityComplianceNotificationPhones()([]string)
    GetSettings()(OrganizationSettingsable)
    GetState()(*string)
    GetStreet()(*string)
    GetTechnicalNotificationMails()([]string)
    GetVerifiedDomains()([]VerifiedDomainable)
    SetAssignedPlans(value []AssignedPlanable)()
    SetBranding(value OrganizationalBrandingable)()
    SetBusinessPhones(value []string)()
    SetCertificateBasedAuthConfiguration(value []CertificateBasedAuthConfigurationable)()
    SetCertificateConnectorSetting(value CertificateConnectorSettingable)()
    SetCity(value *string)()
    SetCountry(value *string)()
    SetCountryLetterCode(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDefaultUsageLocation(value *string)()
    SetDirectorySizeQuota(value DirectorySizeQuotaable)()
    SetDisplayName(value *string)()
    SetExtensions(value []Extensionable)()
    SetIsMultipleDataLocationsForServicesEnabled(value *bool)()
    SetMarketingNotificationEmails(value []string)()
    SetMobileDeviceManagementAuthority(value *MdmAuthority)()
    SetOnPremisesLastPasswordSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOnPremisesSyncEnabled(value *bool)()
    SetPartnerInformation(value PartnerInformationable)()
    SetPartnerTenantType(value *PartnerTenantType)()
    SetPostalCode(value *string)()
    SetPreferredLanguage(value *string)()
    SetPrivacyProfile(value PrivacyProfileable)()
    SetProvisionedPlans(value []ProvisionedPlanable)()
    SetSecurityComplianceNotificationMails(value []string)()
    SetSecurityComplianceNotificationPhones(value []string)()
    SetSettings(value OrganizationSettingsable)()
    SetState(value *string)()
    SetStreet(value *string)()
    SetTechnicalNotificationMails(value []string)()
    SetVerifiedDomains(value []VerifiedDomainable)()
}
