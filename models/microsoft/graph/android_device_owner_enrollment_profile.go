package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AndroidDeviceOwnerEnrollmentProfile struct {
    Entity
    // Tenant GUID the enrollment profile belongs to.
    accountId *string;
    // Date time the enrollment profile was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Description for the enrollment profile.
    description *string;
    // Display name for the enrollment profile.
    displayName *string;
    // Total number of Android devices that have enrolled using this enrollment profile.
    enrolledDeviceCount *int32;
    // The enrollment mode of devices that use this enrollment profile. Possible values are: corporateOwnedDedicatedDevice, corporateOwnedFullyManaged, corporateOwnedWorkProfile, corporateOwnedAOSPUserlessDevice, corporateOwnedAOSPUserAssociatedDevice.
    enrollmentMode *AndroidDeviceOwnerEnrollmentMode;
    // The enrollment token type for an enrollment profile. Possible values are: default, corporateOwnedDedicatedDeviceWithAzureADSharedMode.
    enrollmentTokenType *AndroidDeviceOwnerEnrollmentTokenType;
    // Total number of AOSP devices that have enrolled using the current token.
    enrollmentTokenUsageCount *int32;
    // Date time the enrollment profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // String used to generate a QR code for the token.
    qrCodeContent *string;
    // String used to generate a QR code for the token.
    qrCodeImage *MimeContent;
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string;
    // Date time the most recently created token was created.
    tokenCreationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Date time the most recently created token will expire.
    tokenExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Value of the most recently created token for this enrollment profile.
    tokenValue *string;
    // Boolean that indicates if hidden wifi networks are enabled
    wifiHidden *bool;
    // String that contains the wi-fi login password
    wifiPassword *string;
    // String that contains the wi-fi security type. Possible values are: none, wpa, wep.
    wifiSecurityType *AospWifiSecurityType;
    // String that contains the wi-fi login ssid
    wifiSsid *string;
}
// Instantiates a new androidDeviceOwnerEnrollmentProfile and sets the default values.
func NewAndroidDeviceOwnerEnrollmentProfile()(*AndroidDeviceOwnerEnrollmentProfile) {
    m := &AndroidDeviceOwnerEnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// Gets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Description for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Display name for the enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrolledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDeviceCount
    }
}
// Gets the enrollmentMode property value. The enrollment mode of devices that use this enrollment profile. Possible values are: corporateOwnedDedicatedDevice, corporateOwnedFullyManaged, corporateOwnedWorkProfile, corporateOwnedAOSPUserlessDevice, corporateOwnedAOSPUserAssociatedDevice.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentMode()(*AndroidDeviceOwnerEnrollmentMode) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentMode
    }
}
// Gets the enrollmentTokenType property value. The enrollment token type for an enrollment profile. Possible values are: default, corporateOwnedDedicatedDeviceWithAzureADSharedMode.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentTokenType()(*AndroidDeviceOwnerEnrollmentTokenType) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentTokenType
    }
}
// Gets the enrollmentTokenUsageCount property value. Total number of AOSP devices that have enrolled using the current token.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetEnrollmentTokenUsageCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enrollmentTokenUsageCount
    }
}
// Gets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetQrCodeContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeContent
    }
}
// Gets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetQrCodeImage()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeImage
    }
}
// Gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetRoleScopeTagIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roleScopeTagIds
    }
}
// Gets the tokenCreationDateTime property value. Date time the most recently created token was created.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.tokenCreationDateTime
    }
}
// Gets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.tokenExpirationDateTime
    }
}
// Gets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetTokenValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenValue
    }
}
// Gets the wifiHidden property value. Boolean that indicates if hidden wifi networks are enabled
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiHidden()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.wifiHidden
    }
}
// Gets the wifiPassword property value. String that contains the wi-fi login password
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiPassword
    }
}
// Gets the wifiSecurityType property value. String that contains the wi-fi security type. Possible values are: none, wpa, wep.
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiSecurityType()(*AospWifiSecurityType) {
    if m == nil {
        return nil
    } else {
        return m.wifiSecurityType
    }
}
// Gets the wifiSsid property value. String that contains the wi-fi login ssid
func (m *AndroidDeviceOwnerEnrollmentProfile) GetWifiSsid()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiSsid
    }
}
// The deserialization information for the current model
func (m *AndroidDeviceOwnerEnrollmentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["enrolledDeviceCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrolledDeviceCount(val)
        }
        return nil
    }
    res["enrollmentMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerEnrollmentMode)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AndroidDeviceOwnerEnrollmentMode)
            m.SetEnrollmentMode(&cast)
        }
        return nil
    }
    res["enrollmentTokenType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidDeviceOwnerEnrollmentTokenType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AndroidDeviceOwnerEnrollmentTokenType)
            m.SetEnrollmentTokenType(&cast)
        }
        return nil
    }
    res["enrollmentTokenUsageCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentTokenUsageCount(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["qrCodeContent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQrCodeContent(val)
        }
        return nil
    }
    res["qrCodeImage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMimeContent() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQrCodeImage(val.(*MimeContent))
        }
        return nil
    }
    res["roleScopeTagIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["tokenCreationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenCreationDateTime(val)
        }
        return nil
    }
    res["tokenExpirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenExpirationDateTime(val)
        }
        return nil
    }
    res["tokenValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenValue(val)
        }
        return nil
    }
    res["wifiHidden"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiHidden(val)
        }
        return nil
    }
    res["wifiPassword"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiPassword(val)
        }
        return nil
    }
    res["wifiSecurityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAospWifiSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AospWifiSecurityType)
            m.SetWifiSecurityType(&cast)
        }
        return nil
    }
    res["wifiSsid"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *AndroidDeviceOwnerEnrollmentProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AndroidDeviceOwnerEnrollmentProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        cast := m.GetEnrollmentMode().String()
        err = writer.WriteStringValue("enrollmentMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentTokenType() != nil {
        cast := m.GetEnrollmentTokenType().String()
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
    {
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
        cast := m.GetWifiSecurityType().String()
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
// Sets the accountId property value. Tenant GUID the enrollment profile belongs to.
// Parameters:
//  - value : Value to set for the accountId property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetAccountId(value *string)() {
    m.accountId = value
}
// Sets the createdDateTime property value. Date time the enrollment profile was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Description for the enrollment profile.
// Parameters:
//  - value : Value to set for the description property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Display name for the enrollment profile.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
// Parameters:
//  - value : Value to set for the enrolledDeviceCount property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrolledDeviceCount(value *int32)() {
    m.enrolledDeviceCount = value
}
// Sets the enrollmentMode property value. The enrollment mode of devices that use this enrollment profile. Possible values are: corporateOwnedDedicatedDevice, corporateOwnedFullyManaged, corporateOwnedWorkProfile, corporateOwnedAOSPUserlessDevice, corporateOwnedAOSPUserAssociatedDevice.
// Parameters:
//  - value : Value to set for the enrollmentMode property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentMode(value *AndroidDeviceOwnerEnrollmentMode)() {
    m.enrollmentMode = value
}
// Sets the enrollmentTokenType property value. The enrollment token type for an enrollment profile. Possible values are: default, corporateOwnedDedicatedDeviceWithAzureADSharedMode.
// Parameters:
//  - value : Value to set for the enrollmentTokenType property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentTokenType(value *AndroidDeviceOwnerEnrollmentTokenType)() {
    m.enrollmentTokenType = value
}
// Sets the enrollmentTokenUsageCount property value. Total number of AOSP devices that have enrolled using the current token.
// Parameters:
//  - value : Value to set for the enrollmentTokenUsageCount property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetEnrollmentTokenUsageCount(value *int32)() {
    m.enrollmentTokenUsageCount = value
}
// Sets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the qrCodeContent property value. String used to generate a QR code for the token.
// Parameters:
//  - value : Value to set for the qrCodeContent property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetQrCodeContent(value *string)() {
    m.qrCodeContent = value
}
// Sets the qrCodeImage property value. String used to generate a QR code for the token.
// Parameters:
//  - value : Value to set for the qrCodeImage property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetQrCodeImage(value *MimeContent)() {
    m.qrCodeImage = value
}
// Sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
// Parameters:
//  - value : Value to set for the roleScopeTagIds property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
// Sets the tokenCreationDateTime property value. Date time the most recently created token was created.
// Parameters:
//  - value : Value to set for the tokenCreationDateTime property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.tokenCreationDateTime = value
}
// Sets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
// Parameters:
//  - value : Value to set for the tokenExpirationDateTime property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.tokenExpirationDateTime = value
}
// Sets the tokenValue property value. Value of the most recently created token for this enrollment profile.
// Parameters:
//  - value : Value to set for the tokenValue property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetTokenValue(value *string)() {
    m.tokenValue = value
}
// Sets the wifiHidden property value. Boolean that indicates if hidden wifi networks are enabled
// Parameters:
//  - value : Value to set for the wifiHidden property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiHidden(value *bool)() {
    m.wifiHidden = value
}
// Sets the wifiPassword property value. String that contains the wi-fi login password
// Parameters:
//  - value : Value to set for the wifiPassword property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiPassword(value *string)() {
    m.wifiPassword = value
}
// Sets the wifiSecurityType property value. String that contains the wi-fi security type. Possible values are: none, wpa, wep.
// Parameters:
//  - value : Value to set for the wifiSecurityType property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiSecurityType(value *AospWifiSecurityType)() {
    m.wifiSecurityType = value
}
// Sets the wifiSsid property value. String that contains the wi-fi login ssid
// Parameters:
//  - value : Value to set for the wifiSsid property.
func (m *AndroidDeviceOwnerEnrollmentProfile) SetWifiSsid(value *string)() {
    m.wifiSsid = value
}
