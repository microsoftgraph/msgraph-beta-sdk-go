package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudCertificationAuthorityLeafCertificate entity that represents a leaf certificate issued by a cloud certification authority.
type CloudCertificationAuthorityLeafCertificate struct {
    Entity
}
// NewCloudCertificationAuthorityLeafCertificate instantiates a new CloudCertificationAuthorityLeafCertificate and sets the default values.
func NewCloudCertificationAuthorityLeafCertificate()(*CloudCertificationAuthorityLeafCertificate) {
    m := &CloudCertificationAuthorityLeafCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCloudCertificationAuthorityLeafCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCloudCertificationAuthorityLeafCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCloudCertificationAuthorityLeafCertificate(), nil
}
// GetCertificateStatus gets the certificateStatus property value. Enum type of possible leaf certificate statuses. These statuses indicate whether certificates are active and usable or unusable if they have been revoked or expired.
// returns a *CloudCertificationAuthorityLeafCertificateStatus when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetCertificateStatus()(*CloudCertificationAuthorityLeafCertificateStatus) {
    val, err := m.GetBackingStore().Get("certificateStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*CloudCertificationAuthorityLeafCertificateStatus)
    }
    return nil
}
// GetCertificationAuthorityIssuerUri gets the certificationAuthorityIssuerUri property value. The URI of the certification authority that issued the certificate. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetCertificationAuthorityIssuerUri()(*string) {
    val, err := m.GetBackingStore().Get("certificationAuthorityIssuerUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCrlDistributionPointUrl gets the crlDistributionPointUrl property value. URL to find the relevant Certificate Revocation List for this certificate. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetCrlDistributionPointUrl()(*string) {
    val, err := m.GetBackingStore().Get("crlDistributionPointUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceId gets the deviceId property value. The unique identifier of the managed device for which the certificate was created. This ID is assigned at device enrollment time. Read-only. Supports $select.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("deviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. Name of the device for which the certificate was created. Read-only. Supports $select.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDevicePlatform gets the devicePlatform property value. The platform of the device for which the certificate was created. Possible values are: Android, AndroidForWork, iOS, MacOS, WindowsPhone81, Windows81AndLater, Windows10AndLater, AndroidWorkProfile, Unknown, AndroidAOSP, AndroidMobileApplicationManagement, iOSMobileApplicationManagement. Default value: Unknown. Read-only. Supports $select.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetDevicePlatform()(*string) {
    val, err := m.GetBackingStore().Get("devicePlatform")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. Certificate extensions that further define the purpose of the public key contained in a certificate. Data is formatted as a comma-separated list of object identifiers (OID). For example a possible value is '1.3.6.1.5.5.7.3.2'. Read-only. Nullable.
// returns a []string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetExtendedKeyUsages()([]string) {
    val, err := m.GetBackingStore().Get("extendedKeyUsages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudCertificationAuthorityLeafCertificateStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateStatus(val.(*CloudCertificationAuthorityLeafCertificateStatus))
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
    res["crlDistributionPointUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrlDistributionPointUrl(val)
        }
        return nil
    }
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["devicePlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicePlatform(val)
        }
        return nil
    }
    res["extendedKeyUsages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetExtendedKeyUsages(res)
        }
        return nil
    }
    res["issuerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerId(val)
        }
        return nil
    }
    res["issuerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerName(val)
        }
        return nil
    }
    res["keyUsages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetKeyUsages(res)
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
    res["revocationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRevocationDateTime(val)
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
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
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
    return res
}
// GetIssuerId gets the issuerId property value. The globally unique identifier of the certification authority that issued the leaf certificate. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetIssuerId()(*string) {
    val, err := m.GetBackingStore().Get("issuerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetIssuerName gets the issuerName property value. The name of the certification authority that issued the leaf certificate. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetIssuerName()(*string) {
    val, err := m.GetBackingStore().Get("issuerName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetKeyUsages gets the keyUsages property value. Certificate extensions that define the purpose of the public key contained in a certificate. For example possible values are 'Key Encipherment' and 'Digital Signature'. Read-only. Nullable.
// returns a []string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetKeyUsages()([]string) {
    val, err := m.GetBackingStore().Get("keyUsages")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetOcspResponderUri gets the ocspResponderUri property value. The Online Certificate Status Protocol (OCSP) responder URI that can be used to determine certificate status. Read-only.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetOcspResponderUri()(*string) {
    val, err := m.GetBackingStore().Get("ocspResponderUri")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRevocationDateTime gets the revocationDateTime property value. The date and time a certificate was revoked. If the certificate was not revoked, this will be null. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only.
// returns a *Time when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetRevocationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("revocationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetSerialNumber gets the serialNumber property value. The serial number used to uniquely identify a certificate with its issuing certification authority. Read-only. Supports $select.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetSerialNumber()(*string) {
    val, err := m.GetBackingStore().Get("serialNumber")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubjectName gets the subjectName property value. The subject name of the certificate. The subject is the target or intended beneficiary of the security being provided, such as a user or device. Read-only. Supports $select and $orderby.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetSubjectName()(*string) {
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
func (m *CloudCertificationAuthorityLeafCertificate) GetThumbprint()(*string) {
    val, err := m.GetBackingStore().Get("thumbprint")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserId gets the userId property value. The unique identifier of the user for which the certificate was created. Null for userless devices. This is an Intune user ID. Nullable. Read-only. Supports $select.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name of the user for which the certificate was created. Null for userless devices. Nullable. Read-only. Supports $select.
// returns a *string when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("userPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetValidityEndDateTime gets the validityEndDateTime property value. The end date time of the validity period of a certificate. Certificates cannot be used after this date time as they are longer valid. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only. Supports $orderby.
// returns a *Time when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetValidityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("validityEndDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetValidityStartDateTime gets the validityStartDateTime property value. The start date time of the validity period of a certificate. Certificates cannot be used before this date time as they are not yet valid. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only. Supports $orderby.
// returns a *Time when successful
func (m *CloudCertificationAuthorityLeafCertificate) GetValidityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("validityStartDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CloudCertificationAuthorityLeafCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateStatus() != nil {
        cast := (*m.GetCertificateStatus()).String()
        err = writer.WriteStringValue("certificateStatus", &cast)
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
    {
        err = writer.WriteStringValue("crlDistributionPointUrl", m.GetCrlDistributionPointUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("devicePlatform", m.GetDevicePlatform())
        if err != nil {
            return err
        }
    }
    if m.GetExtendedKeyUsages() != nil {
        err = writer.WriteCollectionOfStringValues("extendedKeyUsages", m.GetExtendedKeyUsages())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuerId", m.GetIssuerId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuerName", m.GetIssuerName())
        if err != nil {
            return err
        }
    }
    if m.GetKeyUsages() != nil {
        err = writer.WriteCollectionOfStringValues("keyUsages", m.GetKeyUsages())
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
        err = writer.WriteTimeValue("revocationDateTime", m.GetRevocationDateTime())
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
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
        err = writer.WriteTimeValue("validityStartDateTime", m.GetValidityStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateStatus sets the certificateStatus property value. Enum type of possible leaf certificate statuses. These statuses indicate whether certificates are active and usable or unusable if they have been revoked or expired.
func (m *CloudCertificationAuthorityLeafCertificate) SetCertificateStatus(value *CloudCertificationAuthorityLeafCertificateStatus)() {
    err := m.GetBackingStore().Set("certificateStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetCertificationAuthorityIssuerUri sets the certificationAuthorityIssuerUri property value. The URI of the certification authority that issued the certificate. Read-only.
func (m *CloudCertificationAuthorityLeafCertificate) SetCertificationAuthorityIssuerUri(value *string)() {
    err := m.GetBackingStore().Set("certificationAuthorityIssuerUri", value)
    if err != nil {
        panic(err)
    }
}
// SetCrlDistributionPointUrl sets the crlDistributionPointUrl property value. URL to find the relevant Certificate Revocation List for this certificate. Read-only.
func (m *CloudCertificationAuthorityLeafCertificate) SetCrlDistributionPointUrl(value *string)() {
    err := m.GetBackingStore().Set("crlDistributionPointUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceId sets the deviceId property value. The unique identifier of the managed device for which the certificate was created. This ID is assigned at device enrollment time. Read-only. Supports $select.
func (m *CloudCertificationAuthorityLeafCertificate) SetDeviceId(value *string)() {
    err := m.GetBackingStore().Set("deviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. Name of the device for which the certificate was created. Read-only. Supports $select.
func (m *CloudCertificationAuthorityLeafCertificate) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetDevicePlatform sets the devicePlatform property value. The platform of the device for which the certificate was created. Possible values are: Android, AndroidForWork, iOS, MacOS, WindowsPhone81, Windows81AndLater, Windows10AndLater, AndroidWorkProfile, Unknown, AndroidAOSP, AndroidMobileApplicationManagement, iOSMobileApplicationManagement. Default value: Unknown. Read-only. Supports $select.
func (m *CloudCertificationAuthorityLeafCertificate) SetDevicePlatform(value *string)() {
    err := m.GetBackingStore().Set("devicePlatform", value)
    if err != nil {
        panic(err)
    }
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Certificate extensions that further define the purpose of the public key contained in a certificate. Data is formatted as a comma-separated list of object identifiers (OID). For example a possible value is '1.3.6.1.5.5.7.3.2'. Read-only. Nullable.
func (m *CloudCertificationAuthorityLeafCertificate) SetExtendedKeyUsages(value []string)() {
    err := m.GetBackingStore().Set("extendedKeyUsages", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuerId sets the issuerId property value. The globally unique identifier of the certification authority that issued the leaf certificate. Read-only.
func (m *CloudCertificationAuthorityLeafCertificate) SetIssuerId(value *string)() {
    err := m.GetBackingStore().Set("issuerId", value)
    if err != nil {
        panic(err)
    }
}
// SetIssuerName sets the issuerName property value. The name of the certification authority that issued the leaf certificate. Read-only.
func (m *CloudCertificationAuthorityLeafCertificate) SetIssuerName(value *string)() {
    err := m.GetBackingStore().Set("issuerName", value)
    if err != nil {
        panic(err)
    }
}
// SetKeyUsages sets the keyUsages property value. Certificate extensions that define the purpose of the public key contained in a certificate. For example possible values are 'Key Encipherment' and 'Digital Signature'. Read-only. Nullable.
func (m *CloudCertificationAuthorityLeafCertificate) SetKeyUsages(value []string)() {
    err := m.GetBackingStore().Set("keyUsages", value)
    if err != nil {
        panic(err)
    }
}
// SetOcspResponderUri sets the ocspResponderUri property value. The Online Certificate Status Protocol (OCSP) responder URI that can be used to determine certificate status. Read-only.
func (m *CloudCertificationAuthorityLeafCertificate) SetOcspResponderUri(value *string)() {
    err := m.GetBackingStore().Set("ocspResponderUri", value)
    if err != nil {
        panic(err)
    }
}
// SetRevocationDateTime sets the revocationDateTime property value. The date and time a certificate was revoked. If the certificate was not revoked, this will be null. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only.
func (m *CloudCertificationAuthorityLeafCertificate) SetRevocationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("revocationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSerialNumber sets the serialNumber property value. The serial number used to uniquely identify a certificate with its issuing certification authority. Read-only. Supports $select.
func (m *CloudCertificationAuthorityLeafCertificate) SetSerialNumber(value *string)() {
    err := m.GetBackingStore().Set("serialNumber", value)
    if err != nil {
        panic(err)
    }
}
// SetSubjectName sets the subjectName property value. The subject name of the certificate. The subject is the target or intended beneficiary of the security being provided, such as a user or device. Read-only. Supports $select and $orderby.
func (m *CloudCertificationAuthorityLeafCertificate) SetSubjectName(value *string)() {
    err := m.GetBackingStore().Set("subjectName", value)
    if err != nil {
        panic(err)
    }
}
// SetThumbprint sets the thumbprint property value. Secure Hash Algorithm 1 digest of the certificate that can be used to identify it. Read-only. Supports $select.
func (m *CloudCertificationAuthorityLeafCertificate) SetThumbprint(value *string)() {
    err := m.GetBackingStore().Set("thumbprint", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The unique identifier of the user for which the certificate was created. Null for userless devices. This is an Intune user ID. Nullable. Read-only. Supports $select.
func (m *CloudCertificationAuthorityLeafCertificate) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name of the user for which the certificate was created. Null for userless devices. Nullable. Read-only. Supports $select.
func (m *CloudCertificationAuthorityLeafCertificate) SetUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("userPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetValidityEndDateTime sets the validityEndDateTime property value. The end date time of the validity period of a certificate. Certificates cannot be used after this date time as they are longer valid. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only. Supports $orderby.
func (m *CloudCertificationAuthorityLeafCertificate) SetValidityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("validityEndDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetValidityStartDateTime sets the validityStartDateTime property value. The start date time of the validity period of a certificate. Certificates cannot be used before this date time as they are not yet valid. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Nullable. Read-only. Supports $orderby.
func (m *CloudCertificationAuthorityLeafCertificate) SetValidityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("validityStartDateTime", value)
    if err != nil {
        panic(err)
    }
}
type CloudCertificationAuthorityLeafCertificateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateStatus()(*CloudCertificationAuthorityLeafCertificateStatus)
    GetCertificationAuthorityIssuerUri()(*string)
    GetCrlDistributionPointUrl()(*string)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetDevicePlatform()(*string)
    GetExtendedKeyUsages()([]string)
    GetIssuerId()(*string)
    GetIssuerName()(*string)
    GetKeyUsages()([]string)
    GetOcspResponderUri()(*string)
    GetRevocationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSerialNumber()(*string)
    GetSubjectName()(*string)
    GetThumbprint()(*string)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetValidityEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetValidityStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCertificateStatus(value *CloudCertificationAuthorityLeafCertificateStatus)()
    SetCertificationAuthorityIssuerUri(value *string)()
    SetCrlDistributionPointUrl(value *string)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetDevicePlatform(value *string)()
    SetExtendedKeyUsages(value []string)()
    SetIssuerId(value *string)()
    SetIssuerName(value *string)()
    SetKeyUsages(value []string)()
    SetOcspResponderUri(value *string)()
    SetRevocationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSerialNumber(value *string)()
    SetSubjectName(value *string)()
    SetThumbprint(value *string)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetValidityEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetValidityStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
