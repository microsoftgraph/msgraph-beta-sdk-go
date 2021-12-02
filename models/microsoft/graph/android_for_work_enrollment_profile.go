package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidForWorkEnrollmentProfile 
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
// NewAndroidForWorkEnrollmentProfile instantiates a new androidForWorkEnrollmentProfile and sets the default values.
func NewAndroidForWorkEnrollmentProfile()(*AndroidForWorkEnrollmentProfile) {
    m := &AndroidForWorkEnrollmentProfile{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccountId gets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidForWorkEnrollmentProfile) GetAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidForWorkEnrollmentProfile) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDescription gets the description property value. Description for the enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Display name for the enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEnrolledDeviceCount gets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetEnrolledDeviceCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.enrolledDeviceCount
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidForWorkEnrollmentProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetQrCodeContent gets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidForWorkEnrollmentProfile) GetQrCodeContent()(*string) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeContent
    }
}
// GetQrCodeImage gets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidForWorkEnrollmentProfile) GetQrCodeImage()(*MimeContent) {
    if m == nil {
        return nil
    } else {
        return m.qrCodeImage
    }
}
// GetTokenExpirationDateTime gets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidForWorkEnrollmentProfile) GetTokenExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.tokenExpirationDateTime
    }
}
// GetTokenValue gets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) GetTokenValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tokenValue
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetAccountId sets the accountId property value. Tenant GUID the enrollment profile belongs to.
func (m *AndroidForWorkEnrollmentProfile) SetAccountId(value *string)() {
    if m != nil {
        m.accountId = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Date time the enrollment profile was created.
func (m *AndroidForWorkEnrollmentProfile) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDescription sets the description property value. Description for the enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Display name for the enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEnrolledDeviceCount sets the enrolledDeviceCount property value. Total number of Android devices that have enrolled using this enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) SetEnrolledDeviceCount(value *int32)() {
    if m != nil {
        m.enrolledDeviceCount = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date time the enrollment profile was last modified.
func (m *AndroidForWorkEnrollmentProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetQrCodeContent sets the qrCodeContent property value. String used to generate a QR code for the token.
func (m *AndroidForWorkEnrollmentProfile) SetQrCodeContent(value *string)() {
    if m != nil {
        m.qrCodeContent = value
    }
}
// SetQrCodeImage sets the qrCodeImage property value. String used to generate a QR code for the token.
func (m *AndroidForWorkEnrollmentProfile) SetQrCodeImage(value *MimeContent)() {
    if m != nil {
        m.qrCodeImage = value
    }
}
// SetTokenExpirationDateTime sets the tokenExpirationDateTime property value. Date time the most recently created token will expire.
func (m *AndroidForWorkEnrollmentProfile) SetTokenExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.tokenExpirationDateTime = value
    }
}
// SetTokenValue sets the tokenValue property value. Value of the most recently created token for this enrollment profile.
func (m *AndroidForWorkEnrollmentProfile) SetTokenValue(value *string)() {
    if m != nil {
        m.tokenValue = value
    }
}
