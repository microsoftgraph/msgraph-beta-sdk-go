package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudCertificationAuthority entity that represents a collection of metadata of a cloud certification authority.
type CloudCertificationAuthority struct {
    Entity
}
// NewCloudCertificationAuthority instantiates a new CloudCertificationAuthority and sets the default values.
func NewCloudCertificationAuthority()(*CloudCertificationAuthority) {
    m := &CloudCertificationAuthority{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudCertificationAuthorityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudCertificationAuthorityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudCertificationAuthority(), nil
}
// GetCertificateDownloadUrl gets the certificateDownloadUrl property value. The URL to download the certification authority certificate. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetCertificateDownloadUrl()(*string) {
    val, err := m.GetBackingStore().Get("certificateDownloadUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateKeySize gets the certificateKeySize property value. Enum of possible cloud certification authority certificate cryptography and key size combinations.
// returns a *CloudCertificationAuthorityCertificateKeySize when successful
func (m *CloudCertificationAuthority) GetCertificateKeySize()(*CloudCertificationAuthorityCertificateKeySize) {
    val, err := m.GetBackingStore().Get("certificateKeySize")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudCertificationAuthorityCertificateKeySize)
    }
    return nil
}
// GetCertificateRevocationListUrl gets the certificateRevocationListUrl property value. The cloud certification authority's Certificate Revocation List URL that can be used to determine revocation status. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetCertificateRevocationListUrl()(*string) {
    val, err := m.GetBackingStore().Get("certificateRevocationListUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificateSigningRequest gets the certificateSigningRequest property value. The certificate signing request used to create an issuing certification authority with a root certification authority external to Microsoft Cloud PKI. The based-64 encoded certificate signing request can be downloaded through this property. After downloading the certificate signing request, it must be signed by the external root certifcation authority. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetCertificateSigningRequest()(*string) {
    val, err := m.GetBackingStore().Get("certificateSigningRequest")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificationAuthorityIssuerId gets the certificationAuthorityIssuerId property value. Issuer (parent) certification authority identifier. Nullable. Read-only. Supports $orderby and $select.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetCertificationAuthorityIssuerId()(*string) {
    val, err := m.GetBackingStore().Get("certificationAuthorityIssuerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificationAuthorityIssuerUri gets the certificationAuthorityIssuerUri property value. The URI of the issuing certification authority of a subordinate certification authority. Returns null if a root certification authority. Nullable. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetCertificationAuthorityIssuerUri()(*string) {
    val, err := m.GetBackingStore().Get("certificationAuthorityIssuerUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCertificationAuthorityStatus gets the certificationAuthorityStatus property value. Enum type of possible certification authority statuses. These statuses indicate whether a certification authority is currently able to issue certificates or temporarily paused or permanently revoked.
// returns a *CloudCertificationAuthorityStatus when successful
func (m *CloudCertificationAuthority) GetCertificationAuthorityStatus()(*CloudCertificationAuthorityStatus) {
    val, err := m.GetBackingStore().Get("certificationAuthorityStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudCertificationAuthorityStatus)
    }
    return nil
}
// GetCloudCertificationAuthorityHashingAlgorithm gets the cloudCertificationAuthorityHashingAlgorithm property value. Enum type of possible certificate hashing algorithms used by the certification authority to create certificates.
// returns a *CloudCertificationAuthorityHashingAlgorithm when successful
func (m *CloudCertificationAuthority) GetCloudCertificationAuthorityHashingAlgorithm()(*CloudCertificationAuthorityHashingAlgorithm) {
    val, err := m.GetBackingStore().Get("cloudCertificationAuthorityHashingAlgorithm")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudCertificationAuthorityHashingAlgorithm)
    }
    return nil
}
// GetCloudCertificationAuthorityLeafCertificate gets the cloudCertificationAuthorityLeafCertificate property value. Required OData property to expose leaf certificate API.
// returns a []CloudCertificationAuthorityLeafCertificateable when successful
func (m *CloudCertificationAuthority) GetCloudCertificationAuthorityLeafCertificate()([]CloudCertificationAuthorityLeafCertificateable) {
    val, err := m.GetBackingStore().Get("cloudCertificationAuthorityLeafCertificate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CloudCertificationAuthorityLeafCertificateable)
    }
    return nil
}
// GetCloudCertificationAuthorityType gets the cloudCertificationAuthorityType property value. Enum type of possible certificate authority types. This feature supports a two-tier certification authority model with a root certification authority and one or more child issuing (intermediate) certification authorities.
// returns a *CloudCertificationAuthorityType when successful
func (m *CloudCertificationAuthority) GetCloudCertificationAuthorityType()(*CloudCertificationAuthorityType) {
    val, err := m.GetBackingStore().Get("cloudCertificationAuthorityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudCertificationAuthorityType)
    }
    return nil
}
// GetCommonName gets the commonName property value. The common name of the certificate subject name, which must be unique. This property is a relative distinguished name used to compose the certificate subject name. Read-only. Supports $select.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetCommonName()(*string) {
    val, err := m.GetBackingStore().Get("commonName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCountryName gets the countryName property value. The country name that is used to compose the subject name of a certification authority certificate in the form 'C='. Nullable. Example: US. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetCountryName()(*string) {
    val, err := m.GetBackingStore().Get("countryName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Creation date of this cloud certification authority entity instance. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only.
// returns a *Time when successful
func (m *CloudCertificationAuthority) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. The certification authority description displayed in the Intune admin console. Nullable. Read/write. Returns null if not set.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The certification authority display name the Intune admin console. Read/write. Supports $select and $orderby.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetETag gets the eTag property value. ETag for optimistic concurrency control. Read/write.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetETag()(*string) {
    val, err := m.GetBackingStore().Get("eTag")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. The certificate extended key usages, which specify the usage capabilities of the certificate. Read-only.
// returns a []ExtendedKeyUsageable when successful
func (m *CloudCertificationAuthority) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
    val, err := m.GetBackingStore().Get("extendedKeyUsages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExtendedKeyUsageable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudCertificationAuthority) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificateDownloadUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateDownloadUrl(val)
        }
        return nil
    }
    res["certificateKeySize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudCertificationAuthorityCertificateKeySize)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKeySize(val.(*CloudCertificationAuthorityCertificateKeySize))
        }
        return nil
    }
    res["certificateRevocationListUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateRevocationListUrl(val)
        }
        return nil
    }
    res["certificateSigningRequest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSigningRequest(val)
        }
        return nil
    }
    res["certificationAuthorityIssuerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityIssuerId(val)
        }
        return nil
    }
    res["certificationAuthorityIssuerUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityIssuerUri(val)
        }
        return nil
    }
    res["certificationAuthorityStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudCertificationAuthorityStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityStatus(val.(*CloudCertificationAuthorityStatus))
        }
        return nil
    }
    res["cloudCertificationAuthorityHashingAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudCertificationAuthorityHashingAlgorithm)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudCertificationAuthorityHashingAlgorithm(val.(*CloudCertificationAuthorityHashingAlgorithm))
        }
        return nil
    }
    res["cloudCertificationAuthorityLeafCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCloudCertificationAuthorityLeafCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CloudCertificationAuthorityLeafCertificateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(CloudCertificationAuthorityLeafCertificateable)
                }
            }
            m.SetCloudCertificationAuthorityLeafCertificate(res)
        }
        return nil
    }
    res["cloudCertificationAuthorityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudCertificationAuthorityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudCertificationAuthorityType(val.(*CloudCertificationAuthorityType))
        }
        return nil
    }
    res["commonName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommonName(val)
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["eTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetETag(val)
        }
        return nil
    }
    res["extendedKeyUsages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtendedKeyUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExtendedKeyUsageable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExtendedKeyUsageable)
                }
            }
            m.SetExtendedKeyUsages(res)
        }
        return nil
    }
    res["issuerCommonName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerCommonName(val)
        }
        return nil
    }
    res["keyPlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudCertificationAuthorityKeyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyPlatform(val.(*CloudCertificationAuthorityKeyPlatformType))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["localityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocalityName(val)
        }
        return nil
    }
    res["ocspResponderUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOcspResponderUri(val)
        }
        return nil
    }
    res["organizationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationName(val)
        }
        return nil
    }
    res["organizationUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationUnit(val)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["rootCertificateCommonName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateCommonName(val)
        }
        return nil
    }
    res["scepServerUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScepServerUrl(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["stateName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateName(val)
        }
        return nil
    }
    res["subjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectName(val)
        }
        return nil
    }
    res["thumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbprint(val)
        }
        return nil
    }
    res["validityEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidityEndDateTime(val)
        }
        return nil
    }
    res["validityPeriodInYears"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidityPeriodInYears(val)
        }
        return nil
    }
    res["validityStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidityStartDateTime(val)
        }
        return nil
    }
    res["versionNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionNumber(val)
        }
        return nil
    }
    return res
}
// GetIssuerCommonName gets the issuerCommonName property value. The issuerCommonName property
// returns a *string when successful
func (m *CloudCertificationAuthority) GetIssuerCommonName()(*string) {
    val, err := m.GetBackingStore().Get("issuerCommonName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKeyPlatform gets the keyPlatform property value. Enum type of possible key platforms used by the certification authority.
// returns a *CloudCertificationAuthorityKeyPlatformType when successful
func (m *CloudCertificationAuthority) GetKeyPlatform()(*CloudCertificationAuthorityKeyPlatformType) {
    val, err := m.GetBackingStore().Get("keyPlatform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudCertificationAuthorityKeyPlatformType)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last modification date and time of this certification authority entity instance. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read/write.
// returns a *Time when successful
func (m *CloudCertificationAuthority) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetLocalityName gets the localityName property value. The locality (town, city, etc.) name that is used to compose the subject name of a certification authority certificate in the form 'L='. This is Nullable. Example: Redmond. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetLocalityName()(*string) {
    val, err := m.GetBackingStore().Get("localityName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOcspResponderUri gets the ocspResponderUri property value. The Online Certificate Status Protocol (OCSP) responder URI that can be used to determine certificate status. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetOcspResponderUri()(*string) {
    val, err := m.GetBackingStore().Get("ocspResponderUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrganizationName gets the organizationName property value. The organization name that is used as a distinguished name in the subject name of a certification authority certificate in the form 'O='. Nullable. Example: Microsoft. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetOrganizationName()(*string) {
    val, err := m.GetBackingStore().Get("organizationName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOrganizationUnit gets the organizationUnit property value. The organization unit name that is used as a distinguished name in the subject name of a certification authority certificate in the form 'OU='. Nullable. Example: Security. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetOrganizationUnit()(*string) {
    val, err := m.GetBackingStore().Get("organizationUnit")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this entity instance. Scope tags limit access to an entity instance. Nullable. Read/write.
// returns a []string when successful
func (m *CloudCertificationAuthority) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetRootCertificateCommonName gets the rootCertificateCommonName property value. The common name of the certificate subject name of the certification authority issuer. This property can be used to identify the certification authority that issued the current certification authority. For issuing certification authorities, this is the common name of the certificate subject name of the root certification authority to which it is anchored. For externally signed certification authorities, this is the common name of the certificate subject name of the signing certification authority. For root certification authorities, this is the common name of the certification authority's own certificate subject name. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetRootCertificateCommonName()(*string) {
    val, err := m.GetBackingStore().Get("rootCertificateCommonName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScepServerUrl gets the scepServerUrl property value. The SCEP server URL for device SCEP connections to request certificates. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetScepServerUrl()(*string) {
    val, err := m.GetBackingStore().Get("scepServerUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. The serial number used to uniquely identify a certificate with its issuing certification authority. Read-only. Supports $select.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("serialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStateName gets the stateName property value. The state or province name that is used to compose the subject name of a certification authority certificate in the form 'ST='. Nullable. Example: Washington. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetStateName()(*string) {
    val, err := m.GetBackingStore().Get("stateName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubjectName gets the subjectName property value. The subject name of the certificate. The subject is the target or intended beneficiary of the security being provided, such as a company or government entity. Read-only. Supports $orderby and $select.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetSubjectName()(*string) {
    val, err := m.GetBackingStore().Get("subjectName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThumbprint gets the thumbprint property value. Secure Hash Algorithm 1 digest of the certificate that can be used to identify it. Read-only. Supports $select.
// returns a *string when successful
func (m *CloudCertificationAuthority) GetThumbprint()(*string) {
    val, err := m.GetBackingStore().Get("thumbprint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValidityEndDateTime gets the validityEndDateTime property value. The end date time of the validity period of a certification authority certificate. Certificates cannot be used after this date time as they are longer valid. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only. Supports $orderby.
// returns a *Time when successful
func (m *CloudCertificationAuthority) GetValidityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("validityEndDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetValidityPeriodInYears gets the validityPeriodInYears property value. The certification authority validity period in years configured by admins.
// returns a *int32 when successful
func (m *CloudCertificationAuthority) GetValidityPeriodInYears()(*int32) {
    val, err := m.GetBackingStore().Get("validityPeriodInYears")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetValidityStartDateTime gets the validityStartDateTime property value. The start date time of the validity period of a certification authority certificate. Certificates cannot be used before this date time as they are not yet valid. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only. Supports $orderby.
// returns a *Time when successful
func (m *CloudCertificationAuthority) GetValidityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("validityStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetVersionNumber gets the versionNumber property value. The certification authority version, which is incremented each time the certification authority is renewed. Read-only.
// returns a *int32 when successful
func (m *CloudCertificationAuthority) GetVersionNumber()(*int32) {
    val, err := m.GetBackingStore().Get("versionNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudCertificationAuthority) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certificateDownloadUrl", m.GetCertificateDownloadUrl())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateKeySize() != nil {
        cast := (*m.GetCertificateKeySize()).String()
        err = writer.WriteStringValue("certificateKeySize", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateRevocationListUrl", m.GetCertificateRevocationListUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSigningRequest", m.GetCertificateSigningRequest())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificationAuthorityIssuerId", m.GetCertificationAuthorityIssuerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificationAuthorityIssuerUri", m.GetCertificationAuthorityIssuerUri())
        if err != nil {
            return err
        }
    }
    if m.GetCertificationAuthorityStatus() != nil {
        cast := (*m.GetCertificationAuthorityStatus()).String()
        err = writer.WriteStringValue("certificationAuthorityStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudCertificationAuthorityHashingAlgorithm() != nil {
        cast := (*m.GetCloudCertificationAuthorityHashingAlgorithm()).String()
        err = writer.WriteStringValue("cloudCertificationAuthorityHashingAlgorithm", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudCertificationAuthorityLeafCertificate() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCloudCertificationAuthorityLeafCertificate()))
        for i, v := range m.GetCloudCertificationAuthorityLeafCertificate() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("cloudCertificationAuthorityLeafCertificate", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCloudCertificationAuthorityType() != nil {
        cast := (*m.GetCloudCertificationAuthorityType()).String()
        err = writer.WriteStringValue("cloudCertificationAuthorityType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("commonName", m.GetCommonName())
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteStringValue("eTag", m.GetETag())
        if err != nil {
            return err
        }
    }
    if m.GetExtendedKeyUsages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtendedKeyUsages()))
        for i, v := range m.GetExtendedKeyUsages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("extendedKeyUsages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuerCommonName", m.GetIssuerCommonName())
        if err != nil {
            return err
        }
    }
    if m.GetKeyPlatform() != nil {
        cast := (*m.GetKeyPlatform()).String()
        err = writer.WriteStringValue("keyPlatform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("localityName", m.GetLocalityName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ocspResponderUri", m.GetOcspResponderUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizationName", m.GetOrganizationName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("organizationUnit", m.GetOrganizationUnit())
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rootCertificateCommonName", m.GetRootCertificateCommonName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scepServerUrl", m.GetScepServerUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("stateName", m.GetStateName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectName", m.GetSubjectName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbprint", m.GetThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("validityEndDateTime", m.GetValidityEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("validityPeriodInYears", m.GetValidityPeriodInYears())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("validityStartDateTime", m.GetValidityStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("versionNumber", m.GetVersionNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateDownloadUrl sets the certificateDownloadUrl property value. The URL to download the certification authority certificate. Read-only.
func (m *CloudCertificationAuthority) SetCertificateDownloadUrl(value *string)() {
    err := m.GetBackingStore().Set("certificateDownloadUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateKeySize sets the certificateKeySize property value. Enum of possible cloud certification authority certificate cryptography and key size combinations.
func (m *CloudCertificationAuthority) SetCertificateKeySize(value *CloudCertificationAuthorityCertificateKeySize)() {
    err := m.GetBackingStore().Set("certificateKeySize", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateRevocationListUrl sets the certificateRevocationListUrl property value. The cloud certification authority's Certificate Revocation List URL that can be used to determine revocation status. Read-only.
func (m *CloudCertificationAuthority) SetCertificateRevocationListUrl(value *string)() {
    err := m.GetBackingStore().Set("certificateRevocationListUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificateSigningRequest sets the certificateSigningRequest property value. The certificate signing request used to create an issuing certification authority with a root certification authority external to Microsoft Cloud PKI. The based-64 encoded certificate signing request can be downloaded through this property. After downloading the certificate signing request, it must be signed by the external root certifcation authority. Read-only.
func (m *CloudCertificationAuthority) SetCertificateSigningRequest(value *string)() {
    err := m.GetBackingStore().Set("certificateSigningRequest", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificationAuthorityIssuerId sets the certificationAuthorityIssuerId property value. Issuer (parent) certification authority identifier. Nullable. Read-only. Supports $orderby and $select.
func (m *CloudCertificationAuthority) SetCertificationAuthorityIssuerId(value *string)() {
    err := m.GetBackingStore().Set("certificationAuthorityIssuerId", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificationAuthorityIssuerUri sets the certificationAuthorityIssuerUri property value. The URI of the issuing certification authority of a subordinate certification authority. Returns null if a root certification authority. Nullable. Read-only.
func (m *CloudCertificationAuthority) SetCertificationAuthorityIssuerUri(value *string)() {
    err := m.GetBackingStore().Set("certificationAuthorityIssuerUri", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificationAuthorityStatus sets the certificationAuthorityStatus property value. Enum type of possible certification authority statuses. These statuses indicate whether a certification authority is currently able to issue certificates or temporarily paused or permanently revoked.
func (m *CloudCertificationAuthority) SetCertificationAuthorityStatus(value *CloudCertificationAuthorityStatus)() {
    err := m.GetBackingStore().Set("certificationAuthorityStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudCertificationAuthorityHashingAlgorithm sets the cloudCertificationAuthorityHashingAlgorithm property value. Enum type of possible certificate hashing algorithms used by the certification authority to create certificates.
func (m *CloudCertificationAuthority) SetCloudCertificationAuthorityHashingAlgorithm(value *CloudCertificationAuthorityHashingAlgorithm)() {
    err := m.GetBackingStore().Set("cloudCertificationAuthorityHashingAlgorithm", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudCertificationAuthorityLeafCertificate sets the cloudCertificationAuthorityLeafCertificate property value. Required OData property to expose leaf certificate API.
func (m *CloudCertificationAuthority) SetCloudCertificationAuthorityLeafCertificate(value []CloudCertificationAuthorityLeafCertificateable)() {
    err := m.GetBackingStore().Set("cloudCertificationAuthorityLeafCertificate", value)
    if err != nil {
        panic(err)
    }
}
// SetCloudCertificationAuthorityType sets the cloudCertificationAuthorityType property value. Enum type of possible certificate authority types. This feature supports a two-tier certification authority model with a root certification authority and one or more child issuing (intermediate) certification authorities.
func (m *CloudCertificationAuthority) SetCloudCertificationAuthorityType(value *CloudCertificationAuthorityType)() {
    err := m.GetBackingStore().Set("cloudCertificationAuthorityType", value)
    if err != nil {
        panic(err)
    }
}
// SetCommonName sets the commonName property value. The common name of the certificate subject name, which must be unique. This property is a relative distinguished name used to compose the certificate subject name. Read-only. Supports $select.
func (m *CloudCertificationAuthority) SetCommonName(value *string)() {
    err := m.GetBackingStore().Set("commonName", value)
    if err != nil {
        panic(err)
    }
}
// SetCountryName sets the countryName property value. The country name that is used to compose the subject name of a certification authority certificate in the form 'C='. Nullable. Example: US. Read-only.
func (m *CloudCertificationAuthority) SetCountryName(value *string)() {
    err := m.GetBackingStore().Set("countryName", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Creation date of this cloud certification authority entity instance. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only.
func (m *CloudCertificationAuthority) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The certification authority description displayed in the Intune admin console. Nullable. Read/write. Returns null if not set.
func (m *CloudCertificationAuthority) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The certification authority display name the Intune admin console. Read/write. Supports $select and $orderby.
func (m *CloudCertificationAuthority) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetETag sets the eTag property value. ETag for optimistic concurrency control. Read/write.
func (m *CloudCertificationAuthority) SetETag(value *string)() {
    err := m.GetBackingStore().Set("eTag", value)
    if err != nil {
        panic(err)
    }
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. The certificate extended key usages, which specify the usage capabilities of the certificate. Read-only.
func (m *CloudCertificationAuthority) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    err := m.GetBackingStore().Set("extendedKeyUsages", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuerCommonName sets the issuerCommonName property value. The issuerCommonName property
func (m *CloudCertificationAuthority) SetIssuerCommonName(value *string)() {
    err := m.GetBackingStore().Set("issuerCommonName", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyPlatform sets the keyPlatform property value. Enum type of possible key platforms used by the certification authority.
func (m *CloudCertificationAuthority) SetKeyPlatform(value *CloudCertificationAuthorityKeyPlatformType)() {
    err := m.GetBackingStore().Set("keyPlatform", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last modification date and time of this certification authority entity instance. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read/write.
func (m *CloudCertificationAuthority) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetLocalityName sets the localityName property value. The locality (town, city, etc.) name that is used to compose the subject name of a certification authority certificate in the form 'L='. This is Nullable. Example: Redmond. Read-only.
func (m *CloudCertificationAuthority) SetLocalityName(value *string)() {
    err := m.GetBackingStore().Set("localityName", value)
    if err != nil {
        panic(err)
    }
}
// SetOcspResponderUri sets the ocspResponderUri property value. The Online Certificate Status Protocol (OCSP) responder URI that can be used to determine certificate status. Read-only.
func (m *CloudCertificationAuthority) SetOcspResponderUri(value *string)() {
    err := m.GetBackingStore().Set("ocspResponderUri", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationName sets the organizationName property value. The organization name that is used as a distinguished name in the subject name of a certification authority certificate in the form 'O='. Nullable. Example: Microsoft. Read-only.
func (m *CloudCertificationAuthority) SetOrganizationName(value *string)() {
    err := m.GetBackingStore().Set("organizationName", value)
    if err != nil {
        panic(err)
    }
}
// SetOrganizationUnit sets the organizationUnit property value. The organization unit name that is used as a distinguished name in the subject name of a certification authority certificate in the form 'OU='. Nullable. Example: Security. Read-only.
func (m *CloudCertificationAuthority) SetOrganizationUnit(value *string)() {
    err := m.GetBackingStore().Set("organizationUnit", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this entity instance. Scope tags limit access to an entity instance. Nullable. Read/write.
func (m *CloudCertificationAuthority) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetRootCertificateCommonName sets the rootCertificateCommonName property value. The common name of the certificate subject name of the certification authority issuer. This property can be used to identify the certification authority that issued the current certification authority. For issuing certification authorities, this is the common name of the certificate subject name of the root certification authority to which it is anchored. For externally signed certification authorities, this is the common name of the certificate subject name of the signing certification authority. For root certification authorities, this is the common name of the certification authority's own certificate subject name. Read-only.
func (m *CloudCertificationAuthority) SetRootCertificateCommonName(value *string)() {
    err := m.GetBackingStore().Set("rootCertificateCommonName", value)
    if err != nil {
        panic(err)
    }
}
// SetScepServerUrl sets the scepServerUrl property value. The SCEP server URL for device SCEP connections to request certificates. Read-only.
func (m *CloudCertificationAuthority) SetScepServerUrl(value *string)() {
    err := m.GetBackingStore().Set("scepServerUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. The serial number used to uniquely identify a certificate with its issuing certification authority. Read-only. Supports $select.
func (m *CloudCertificationAuthority) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetStateName sets the stateName property value. The state or province name that is used to compose the subject name of a certification authority certificate in the form 'ST='. Nullable. Example: Washington. Read-only.
func (m *CloudCertificationAuthority) SetStateName(value *string)() {
    err := m.GetBackingStore().Set("stateName", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectName sets the subjectName property value. The subject name of the certificate. The subject is the target or intended beneficiary of the security being provided, such as a company or government entity. Read-only. Supports $orderby and $select.
func (m *CloudCertificationAuthority) SetSubjectName(value *string)() {
    err := m.GetBackingStore().Set("subjectName", value)
    if err != nil {
        panic(err)
    }
}
// SetThumbprint sets the thumbprint property value. Secure Hash Algorithm 1 digest of the certificate that can be used to identify it. Read-only. Supports $select.
func (m *CloudCertificationAuthority) SetThumbprint(value *string)() {
    err := m.GetBackingStore().Set("thumbprint", value)
    if err != nil {
        panic(err)
    }
}
// SetValidityEndDateTime sets the validityEndDateTime property value. The end date time of the validity period of a certification authority certificate. Certificates cannot be used after this date time as they are longer valid. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only. Supports $orderby.
func (m *CloudCertificationAuthority) SetValidityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("validityEndDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetValidityPeriodInYears sets the validityPeriodInYears property value. The certification authority validity period in years configured by admins.
func (m *CloudCertificationAuthority) SetValidityPeriodInYears(value *int32)() {
    err := m.GetBackingStore().Set("validityPeriodInYears", value)
    if err != nil {
        panic(err)
    }
}
// SetValidityStartDateTime sets the validityStartDateTime property value. The start date time of the validity period of a certification authority certificate. Certificates cannot be used before this date time as they are not yet valid. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only. Supports $orderby.
func (m *CloudCertificationAuthority) SetValidityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("validityStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetVersionNumber sets the versionNumber property value. The certification authority version, which is incremented each time the certification authority is renewed. Read-only.
func (m *CloudCertificationAuthority) SetVersionNumber(value *int32)() {
    err := m.GetBackingStore().Set("versionNumber", value)
    if err != nil {
        panic(err)
    }
}
type CloudCertificationAuthorityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateDownloadUrl()(*string)
    GetCertificateKeySize()(*CloudCertificationAuthorityCertificateKeySize)
    GetCertificateRevocationListUrl()(*string)
    GetCertificateSigningRequest()(*string)
    GetCertificationAuthorityIssuerId()(*string)
    GetCertificationAuthorityIssuerUri()(*string)
    GetCertificationAuthorityStatus()(*CloudCertificationAuthorityStatus)
    GetCloudCertificationAuthorityHashingAlgorithm()(*CloudCertificationAuthorityHashingAlgorithm)
    GetCloudCertificationAuthorityLeafCertificate()([]CloudCertificationAuthorityLeafCertificateable)
    GetCloudCertificationAuthorityType()(*CloudCertificationAuthorityType)
    GetCommonName()(*string)
    GetCountryName()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetETag()(*string)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    GetIssuerCommonName()(*string)
    GetKeyPlatform()(*CloudCertificationAuthorityKeyPlatformType)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLocalityName()(*string)
    GetOcspResponderUri()(*string)
    GetOrganizationName()(*string)
    GetOrganizationUnit()(*string)
    GetRoleScopeTagIds()([]string)
    GetRootCertificateCommonName()(*string)
    GetScepServerUrl()(*string)
    GetSerialNumber()(*string)
    GetStateName()(*string)
    GetSubjectName()(*string)
    GetThumbprint()(*string)
    GetValidityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetValidityPeriodInYears()(*int32)
    GetValidityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVersionNumber()(*int32)
    SetCertificateDownloadUrl(value *string)()
    SetCertificateKeySize(value *CloudCertificationAuthorityCertificateKeySize)()
    SetCertificateRevocationListUrl(value *string)()
    SetCertificateSigningRequest(value *string)()
    SetCertificationAuthorityIssuerId(value *string)()
    SetCertificationAuthorityIssuerUri(value *string)()
    SetCertificationAuthorityStatus(value *CloudCertificationAuthorityStatus)()
    SetCloudCertificationAuthorityHashingAlgorithm(value *CloudCertificationAuthorityHashingAlgorithm)()
    SetCloudCertificationAuthorityLeafCertificate(value []CloudCertificationAuthorityLeafCertificateable)()
    SetCloudCertificationAuthorityType(value *CloudCertificationAuthorityType)()
    SetCommonName(value *string)()
    SetCountryName(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetETag(value *string)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
    SetIssuerCommonName(value *string)()
    SetKeyPlatform(value *CloudCertificationAuthorityKeyPlatformType)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLocalityName(value *string)()
    SetOcspResponderUri(value *string)()
    SetOrganizationName(value *string)()
    SetOrganizationUnit(value *string)()
    SetRoleScopeTagIds(value []string)()
    SetRootCertificateCommonName(value *string)()
    SetScepServerUrl(value *string)()
    SetSerialNumber(value *string)()
    SetStateName(value *string)()
    SetSubjectName(value *string)()
    SetThumbprint(value *string)()
    SetValidityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetValidityPeriodInYears(value *int32)()
    SetValidityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVersionNumber(value *int32)()
}
