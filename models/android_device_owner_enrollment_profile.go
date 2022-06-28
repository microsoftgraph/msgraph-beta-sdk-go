package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerEnrollmentProfile enrollment Profile used to enroll Android Enterprise devices using Google's Cloud Management.
type AndroidDeviceOwnerEnrollmentProfile struct {
    Entity
    // Tenant GUID the enrollment profile belongs to.
    accountId *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Date time the enrollment profile was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description for the enrollment profile.
    description *string
    // Display name for the enrollment profile.
    displayName *string
    // Total number of Android devices that have enrolled using this enrollment profile.
    enrolledDeviceCount *int32
    // The enrollment mode of devices that use this enrollment profile. Possible values are: corporateOwnedDedicatedDevice, corporateOwnedFullyManaged, corporateOwnedWorkProfile, corporateOwnedAOSPUserlessDevice, corporateOwnedAOSPUserAssociatedDevice.
    enrollmentMode *AndroidDeviceOwnerEnrollmentMode
    // The enrollment token type for an enrollment profile. Possible values are: default, corporateOwnedDedicatedDeviceWithAzureADSharedMode.
    enrollmentTokenType *AndroidDeviceOwnerEnrollmentTokenType
    // Total number of AOSP devices that have enrolled using the current token.
    enrollmentTokenUsageCount *int32
    // Date time the enrollment profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // String used to generate a QR code for the token.
    qrCodeContent *string
    // String used to generate a QR code for the token.
    qrCodeImage MimeContentable
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
    // Date time the most recently created token was created.
    tokenCreationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date time the most recently created token will expire.
    tokenExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Value of the most recently created token for this enrollment profile.
    tokenValue *string
    // Boolean that indicates if hidden wifi networks are enabled
    wifiHidden *bool
    // String that contains the wi-fi login password
    wifiPassword *string
    // String that contains the wi-fi security type. Possible values are: none, wpa, wep.
    wifiSecurityType *AospWifiSecurityType
    // String that contains the wi-fi login ssid
    wifiSsid *string
}
// NewAndroidDeviceOwnerEnrollmentProfile instantiates a new androidDeviceOwnerEnrollmentProfile and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfile()(*AndroidDeviceOwnerEnrollmentProfile) {
    m := &AndroidDeviceOwnerEnrollmentProfile{
        Entity: *NewEntity(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerEnrollmentProfile(), nil
}
// GetAccountId gets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Description for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display name for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnrolledDeviceCount gets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrolledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDeviceCount
    }
}
// GetEnrollmentMode gets the enrollmentMode property value. The enrollment mode of devices that use this enrollment profile. Possible values are: corporateOwnedDedicatedDevice, corporateOwnedFullyManaged, corporateOwnedWorkProfile, corporateOwnedAOSPUserlessDevice, corporateOwnedAOSPUserAssociatedDevice.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentMode()(*AndroidDeviceOwnerEnrollmentMode) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentMode
    }
}
// GetEnrollmentTokenType gets the enrollmentTokenType property value. The enrollment token type for an enrollment profile. Possible values are: default, corporateOwnedDedicatedDeviceWithAzureADSharedMode.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentTokenType()(*AndroidDeviceOwnerEnrollmentTokenType) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentTokenType
    }
}
// GetEnrollmentTokenUsageCount gets the enrollmentTokenUsageCount property value. Total number of AOSP devices that have enrolled using the current token.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentTokenUsageCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentTokenUsageCount
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerEnrollmentProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
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
    res["enrolledDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDeviceCount(val)
        }
        return nil
    }
    res["enrollmentMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerEnrollmentMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentMode(val.(*AndroidDeviceOwnerEnrollmentMode))
        }
        return nil
    }
    res["enrollmentTokenType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerEnrollmentTokenType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentTokenType(val.(*AndroidDeviceOwnerEnrollmentTokenType))
        }
        return nil
    }
    res["enrollmentTokenUsageCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentTokenUsageCount(val)
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
    res["qrCodeContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQrCodeContent(val)
        }
        return nil
    }
    res["qrCodeImage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMimeContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQrCodeImage(val.(MimeContentable))
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
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    res["tokenCreationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenCreationDateTime(val)
        }
        return nil
    }
    res["tokenExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenExpirationDateTime(val)
        }
        return nil
    }
    res["tokenValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenValue(val)
        }
        return nil
    }
    res["wifiHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiHidden(val)
        }
        return nil
    }
    res["wifiPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiPassword(val)
        }
        return nil
    }
    res["wifiSecurityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAospWifiSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiSecurityType(val.(*AospWifiSecurityType))
        }
        return nil
    }
    res["wifiSsid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiSsid(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetQrCodeContent gets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetQrCodeContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeContent
    }
}
// GetQrCodeImage gets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetQrCodeImage()(MimeContentable) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeImage
    }
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// GetTokenCreationDateTime gets the tokenCreationDateTime property value. Date time the most recently created token was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.tokenCreationDateTime
    }
}
// GetTokenExpirationDateTime gets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.tokenExpirationDateTime
    }
}
// GetTokenValue gets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenValue
    }
}
// GetWifiHidden gets the wifiHidden property value. Boolean that indicates if hidden wifi networks are enabled
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wifiHidden
    }
}
// GetWifiPassword gets the wifiPassword property value. String that contains the wi-fi login password
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiPassword
    }
}
// GetWifiSecurityType gets the wifiSecurityType property value. String that contains the wi-fi security type. Possible values are: none, wpa, wep.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiSecurityType()(*AospWifiSecurityType) {
    if m == nil {
        return nil
    } else {
        return m.wifiSecurityType
    }
}
// GetWifiSsid gets the wifiSsid property value. String that contains the wi-fi login ssid
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiSsid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiSsid
    }
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerEnrollmentProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accountId", m.GetAccountId())
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
        err = writer.WriteInt32Value("enrolledDeviceCount", m.GetEnrolledDeviceCount())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentMode() != nil {
        cast := (*m.GetEnrollmentMode()).String()
        err = writer.WriteStringValue("enrollmentMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentTokenType() != nil {
        cast := (*m.GetEnrollmentTokenType()).String()
        err = writer.WriteStringValue("enrollmentTokenType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("enrollmentTokenUsageCount", m.GetEnrollmentTokenUsageCount())
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
        err = writer.WriteStringValue("qrCodeContent", m.GetQrCodeContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("qrCodeImage", m.GetQrCodeImage())
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
        err = writer.WriteTimeValue("tokenCreationDateTime", m.GetTokenCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("tokenExpirationDateTime", m.GetTokenExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tokenValue", m.GetTokenValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wifiHidden", m.GetWifiHidden())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("wifiPassword", m.GetWifiPassword())
        if err != nil {
            return err
        }
    }
    if m.GetWifiSecurityType() != nil {
        cast := (*m.GetWifiSecurityType()).String()
        err = writer.WriteStringValue("wifiSecurityType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("wifiSsid", m.GetWifiSsid())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountId sets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetAccountId(value *string)() {
    if m != nil {
        m.accountId = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Description for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display name for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnrolledDeviceCount sets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrolledDeviceCount(value *int32)() {
    if m != nil {
        m.enrolledDeviceCount = value
    }
}
// SetEnrollmentMode sets the enrollmentMode property value. The enrollment mode of devices that use this enrollment profile. Possible values are: corporateOwnedDedicatedDevice, corporateOwnedFullyManaged, corporateOwnedWorkProfile, corporateOwnedAOSPUserlessDevice, corporateOwnedAOSPUserAssociatedDevice.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentMode(value *AndroidDeviceOwnerEnrollmentMode)() {
    if m != nil {
        m.enrollmentMode = value
    }
}
// SetEnrollmentTokenType sets the enrollmentTokenType property value. The enrollment token type for an enrollment profile. Possible values are: default, corporateOwnedDedicatedDeviceWithAzureADSharedMode.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentTokenType(value *AndroidDeviceOwnerEnrollmentTokenType)() {
    if m != nil {
        m.enrollmentTokenType = value
    }
}
// SetEnrollmentTokenUsageCount sets the enrollmentTokenUsageCount property value. Total number of AOSP devices that have enrolled using the current token.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentTokenUsageCount(value *int32)() {
    if m != nil {
        m.enrollmentTokenUsageCount = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetQrCodeContent sets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetQrCodeContent(value *string)() {
    if m != nil {
        m.qrCodeContent = value
    }
}
// SetQrCodeImage sets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetQrCodeImage(value MimeContentable)() {
    if m != nil {
        m.qrCodeImage = value
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetRoleScopeTagIds(value []string)() {
    if m != nil {
        m.roleScopeTagIds = value
    }
}
// SetTokenCreationDateTime sets the tokenCreationDateTime property value. Date time the most recently created token was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.tokenCreationDateTime = value
    }
}
// SetTokenExpirationDateTime sets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.tokenExpirationDateTime = value
    }
}
// SetTokenValue sets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenValue(value *string)() {
    if m != nil {
        m.tokenValue = value
    }
}
// SetWifiHidden sets the wifiHidden property value. Boolean that indicates if hidden wifi networks are enabled
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiHidden(value *bool)() {
    if m != nil {
        m.wifiHidden = value
    }
}
// SetWifiPassword sets the wifiPassword property value. String that contains the wi-fi login password
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiPassword(value *string)() {
    if m != nil {
        m.wifiPassword = value
    }
}
// SetWifiSecurityType sets the wifiSecurityType property value. String that contains the wi-fi security type. Possible values are: none, wpa, wep.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiSecurityType(value *AospWifiSecurityType)() {
    if m != nil {
        m.wifiSecurityType = value
    }
}
// SetWifiSsid sets the wifiSsid property value. String that contains the wi-fi login ssid
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiSsid(value *string)() {
    if m != nil {
        m.wifiSsid = value
    }
}
