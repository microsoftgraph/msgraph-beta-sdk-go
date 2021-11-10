package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AndroidForWorkEnrollmentProfile struct {
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
    // Date time the enrollment profile was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // String used to generate a QR code for the token.
    qrCodeContent *string;
    // String used to generate a QR code for the token.
    qrCodeImage *MimeContent;
    // Date time the most recently created token will expire.
    tokenExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Value of the most recently created token for this enrollment profile.
    tokenValue *string;
}
// Instantiates a new androidForWorkEnrollmentProfile and sets the default values.
func NewAndroidForWorkEnrollmentProfile()(*AndroidForWorkEnrollmentProfile) {
    m := &AndroidForWorkEnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidForWorkEnrollmentProfile) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// Gets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidForWorkEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Description for the enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Display name for the enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetEnrolledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDeviceCount
    }
}
// Gets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidForWorkEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidForWorkEnrollmentProfile) GetQrCodeContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeContent
    }
}
// Gets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidForWorkEnrollmentProfile) GetQrCodeImage()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeImage
    }
}
// Gets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidForWorkEnrollmentProfile) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.tokenExpirationDateTime
    }
}
// Gets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetTokenValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenValue
    }
}
// The deserialization information for the current model
func (m *AndroidForWorkEnrollmentProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    return res
}
func (m *AndroidForWorkEnrollmentProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AndroidForWorkEnrollmentProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    return nil
}
// Sets the accountId property value. Tenant GUID the enrollment profile belongs to.
// Parameters:
//  - value : Value to set for the accountId property.
func (m *AndroidForWorkEnrollmentProfile) SetAccountId(value *string)() {
    m.accountId = value
}
// Sets the createdDateTime property value. Date time the enrollment profile was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *AndroidForWorkEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Description for the enrollment profile.
// Parameters:
//  - value : Value to set for the description property.
func (m *AndroidForWorkEnrollmentProfile) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Display name for the enrollment profile.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AndroidForWorkEnrollmentProfile) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
// Parameters:
//  - value : Value to set for the enrolledDeviceCount property.
func (m *AndroidForWorkEnrollmentProfile) SetEnrolledDeviceCount(value *int32)() {
    m.enrolledDeviceCount = value
}
// Sets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *AndroidForWorkEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the qrCodeContent property value. String used to generate a QR code for the token.
// Parameters:
//  - value : Value to set for the qrCodeContent property.
func (m *AndroidForWorkEnrollmentProfile) SetQrCodeContent(value *string)() {
    m.qrCodeContent = value
}
// Sets the qrCodeImage property value. String used to generate a QR code for the token.
// Parameters:
//  - value : Value to set for the qrCodeImage property.
func (m *AndroidForWorkEnrollmentProfile) SetQrCodeImage(value *MimeContent)() {
    m.qrCodeImage = value
}
// Sets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
// Parameters:
//  - value : Value to set for the tokenExpirationDateTime property.
func (m *AndroidForWorkEnrollmentProfile) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.tokenExpirationDateTime = value
}
// Sets the tokenValue property value. Value of the most recently created token for this enrollment profile.
// Parameters:
//  - value : Value to set for the tokenValue property.
func (m *AndroidForWorkEnrollmentProfile) SetTokenValue(value *string)() {
    m.tokenValue = value
}
