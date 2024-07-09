package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerEnrollmentProfile enrollment Profile used to enroll Android Enterprise devices using Google's Cloud Management.
type AndroidDeviceOwnerEnrollmentProfile struct {
    Entity
}
// NewAndroidDeviceOwnerEnrollmentProfile instantiates a new AndroidDeviceOwnerEnrollmentProfile and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfile()(*AndroidDeviceOwnerEnrollmentProfile) {
    m := &AndroidDeviceOwnerEnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAndroidDeviceOwnerEnrollmentProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidDeviceOwnerEnrollmentProfile(), nil
}
// GetAccountId gets the accountId property value. Tenant GUID the enrollment profile belongs to.
// returns a *string when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetAccountId()(*string) {
    val, err := m.GetBackingStore().Get("accountId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetConfigureWifi gets the configureWifi property value. Boolean that indicates that the Wi-Fi network should be configured during device provisioning. When set to TRUE, device provisioning will use Wi-Fi related properties to automatically connect to Wi-Fi networks. When set to FALSE or undefined, other Wi-Fi related properties will be ignored. Default value is TRUE. Returned by default.
// returns a *bool when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetConfigureWifi()(*bool) {
    val, err := m.GetBackingStore().Get("configureWifi")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Date time the enrollment profile was created.
// returns a *Time when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetDescription gets the description property value. Description for the enrollment profile.
// returns a *string when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Display name for the enrollment profile.
// returns a *string when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEnrolledDeviceCount gets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
// returns a *int32 when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrolledDeviceCount()(*int32) {
    val, err := m.GetBackingStore().Get("enrolledDeviceCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetEnrollmentMode gets the enrollmentMode property value. The enrollment mode for an enrollment profile.
// returns a *AndroidDeviceOwnerEnrollmentMode when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentMode()(*AndroidDeviceOwnerEnrollmentMode) {
    val, err := m.GetBackingStore().Get("enrollmentMode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerEnrollmentMode)
    }
    return nil
}
// GetEnrollmentTokenType gets the enrollmentTokenType property value. The enrollment token type for an enrollment profile.
// returns a *AndroidDeviceOwnerEnrollmentTokenType when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentTokenType()(*AndroidDeviceOwnerEnrollmentTokenType) {
    val, err := m.GetBackingStore().Get("enrollmentTokenType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AndroidDeviceOwnerEnrollmentTokenType)
    }
    return nil
}
// GetEnrollmentTokenUsageCount gets the enrollmentTokenUsageCount property value. Total number of AOSP devices that have enrolled using the current token. Valid values 0 to 20000
// returns a *int32 when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentTokenUsageCount()(*int32) {
    val, err := m.GetBackingStore().Get("enrollmentTokenUsageCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
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
    res["configureWifi"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigureWifi(val)
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
    res["isTeamsDeviceProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTeamsDeviceProfile(val)
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
                if v != nil {
                    res[i] = *(v.(*string))
                }
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
// GetIsTeamsDeviceProfile gets the isTeamsDeviceProfile property value. Boolean indicating if this profile is an Android AOSP for Teams device profile.
// returns a *bool when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetIsTeamsDeviceProfile()(*bool) {
    val, err := m.GetBackingStore().Get("isTeamsDeviceProfile")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
// returns a *Time when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetQrCodeContent gets the qrCodeContent property value. String used to generate a QR code for the token.
// returns a *string when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetQrCodeContent()(*string) {
    val, err := m.GetBackingStore().Get("qrCodeContent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetQrCodeImage gets the qrCodeImage property value. String used to generate a QR code for the token.
// returns a MimeContentable when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetQrCodeImage()(MimeContentable) {
    val, err := m.GetBackingStore().Get("qrCodeImage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(MimeContentable)
    }
    return nil
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
// returns a []string when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetRoleScopeTagIds()([]string) {
    val, err := m.GetBackingStore().Get("roleScopeTagIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetTokenCreationDateTime gets the tokenCreationDateTime property value. Date time the most recently created token was created.
// returns a *Time when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("tokenCreationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetTokenExpirationDateTime gets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
// returns a *Time when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("tokenExpirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetTokenValue gets the tokenValue property value. Value of the most recently created token for this enrollment profile.
// returns a *string when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenValue()(*string) {
    val, err := m.GetBackingStore().Get("tokenValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWifiHidden gets the wifiHidden property value. Boolean that indicates if hidden wifi networks are enabled
// returns a *bool when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiHidden()(*bool) {
    val, err := m.GetBackingStore().Get("wifiHidden")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetWifiPassword gets the wifiPassword property value. String that contains the wi-fi login password
// returns a *string when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiPassword()(*string) {
    val, err := m.GetBackingStore().Get("wifiPassword")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWifiSecurityType gets the wifiSecurityType property value. This enum represents Wi-Fi Security Types for Android Device Owner AOSP Scenarios.
// returns a *AospWifiSecurityType when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiSecurityType()(*AospWifiSecurityType) {
    val, err := m.GetBackingStore().Get("wifiSecurityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AospWifiSecurityType)
    }
    return nil
}
// GetWifiSsid gets the wifiSsid property value. String that contains the wi-fi login ssid
// returns a *string when successful
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiSsid()(*string) {
    val, err := m.GetBackingStore().Get("wifiSsid")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
        err = writer.WriteBoolValue("configureWifi", m.GetConfigureWifi())
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
        err = writer.WriteBoolValue("isTeamsDeviceProfile", m.GetIsTeamsDeviceProfile())
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
    return nil
}
// SetAccountId sets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetAccountId(value *string)() {
    err := m.GetBackingStore().Set("accountId", value)
    if err != nil {
        panic(err)
    }
}
// SetConfigureWifi sets the configureWifi property value. Boolean that indicates that the Wi-Fi network should be configured during device provisioning. When set to TRUE, device provisioning will use Wi-Fi related properties to automatically connect to Wi-Fi networks. When set to FALSE or undefined, other Wi-Fi related properties will be ignored. Default value is TRUE. Returned by default.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetConfigureWifi(value *bool)() {
    err := m.GetBackingStore().Set("configureWifi", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Display name for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrolledDeviceCount sets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrolledDeviceCount(value *int32)() {
    err := m.GetBackingStore().Set("enrolledDeviceCount", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentMode sets the enrollmentMode property value. The enrollment mode for an enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentMode(value *AndroidDeviceOwnerEnrollmentMode)() {
    err := m.GetBackingStore().Set("enrollmentMode", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentTokenType sets the enrollmentTokenType property value. The enrollment token type for an enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentTokenType(value *AndroidDeviceOwnerEnrollmentTokenType)() {
    err := m.GetBackingStore().Set("enrollmentTokenType", value)
    if err != nil {
        panic(err)
    }
}
// SetEnrollmentTokenUsageCount sets the enrollmentTokenUsageCount property value. Total number of AOSP devices that have enrolled using the current token. Valid values 0 to 20000
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentTokenUsageCount(value *int32)() {
    err := m.GetBackingStore().Set("enrollmentTokenUsageCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIsTeamsDeviceProfile sets the isTeamsDeviceProfile property value. Boolean indicating if this profile is an Android AOSP for Teams device profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetIsTeamsDeviceProfile(value *bool)() {
    err := m.GetBackingStore().Set("isTeamsDeviceProfile", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetQrCodeContent sets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetQrCodeContent(value *string)() {
    err := m.GetBackingStore().Set("qrCodeContent", value)
    if err != nil {
        panic(err)
    }
}
// SetQrCodeImage sets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetQrCodeImage(value MimeContentable)() {
    err := m.GetBackingStore().Set("qrCodeImage", value)
    if err != nil {
        panic(err)
    }
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetRoleScopeTagIds(value []string)() {
    err := m.GetBackingStore().Set("roleScopeTagIds", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenCreationDateTime sets the tokenCreationDateTime property value. Date time the most recently created token was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("tokenCreationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenExpirationDateTime sets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("tokenExpirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetTokenValue sets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenValue(value *string)() {
    err := m.GetBackingStore().Set("tokenValue", value)
    if err != nil {
        panic(err)
    }
}
// SetWifiHidden sets the wifiHidden property value. Boolean that indicates if hidden wifi networks are enabled
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiHidden(value *bool)() {
    err := m.GetBackingStore().Set("wifiHidden", value)
    if err != nil {
        panic(err)
    }
}
// SetWifiPassword sets the wifiPassword property value. String that contains the wi-fi login password
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiPassword(value *string)() {
    err := m.GetBackingStore().Set("wifiPassword", value)
    if err != nil {
        panic(err)
    }
}
// SetWifiSecurityType sets the wifiSecurityType property value. This enum represents Wi-Fi Security Types for Android Device Owner AOSP Scenarios.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiSecurityType(value *AospWifiSecurityType)() {
    err := m.GetBackingStore().Set("wifiSecurityType", value)
    if err != nil {
        panic(err)
    }
}
// SetWifiSsid sets the wifiSsid property value. String that contains the wi-fi login ssid
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiSsid(value *string)() {
    err := m.GetBackingStore().Set("wifiSsid", value)
    if err != nil {
        panic(err)
    }
}
type AndroidDeviceOwnerEnrollmentProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountId()(*string)
    GetConfigureWifi()(*bool)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEnrolledDeviceCount()(*int32)
    GetEnrollmentMode()(*AndroidDeviceOwnerEnrollmentMode)
    GetEnrollmentTokenType()(*AndroidDeviceOwnerEnrollmentTokenType)
    GetEnrollmentTokenUsageCount()(*int32)
    GetIsTeamsDeviceProfile()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetQrCodeContent()(*string)
    GetQrCodeImage()(MimeContentable)
    GetRoleScopeTagIds()([]string)
    GetTokenCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTokenValue()(*string)
    GetWifiHidden()(*bool)
    GetWifiPassword()(*string)
    GetWifiSecurityType()(*AospWifiSecurityType)
    GetWifiSsid()(*string)
    SetAccountId(value *string)()
    SetConfigureWifi(value *bool)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEnrolledDeviceCount(value *int32)()
    SetEnrollmentMode(value *AndroidDeviceOwnerEnrollmentMode)()
    SetEnrollmentTokenType(value *AndroidDeviceOwnerEnrollmentTokenType)()
    SetEnrollmentTokenUsageCount(value *int32)()
    SetIsTeamsDeviceProfile(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetQrCodeContent(value *string)()
    SetQrCodeImage(value MimeContentable)()
    SetRoleScopeTagIds(value []string)()
    SetTokenCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTokenValue(value *string)()
    SetWifiHidden(value *bool)()
    SetWifiPassword(value *string)()
    SetWifiSecurityType(value *AospWifiSecurityType)()
    SetWifiSsid(value *string)()
}
