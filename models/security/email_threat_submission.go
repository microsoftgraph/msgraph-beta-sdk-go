package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailThreatSubmission struct {
    ThreatSubmission
}
// NewEmailThreatSubmission instantiates a new EmailThreatSubmission and sets the default values.
func NewEmailThreatSubmission()(*EmailThreatSubmission) {
    m := &EmailThreatSubmission{
        ThreatSubmission: *NewThreatSubmission(),
    }
    odataTypeValue := "#microsoft.graph.security.emailThreatSubmission"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEmailThreatSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailThreatSubmissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.security.emailContentThreatSubmission":
                        return NewEmailContentThreatSubmission(), nil
                    case "#microsoft.graph.security.emailUrlThreatSubmission":
                        return NewEmailUrlThreatSubmission(), nil
                }
            }
        }
    }
    return NewEmailThreatSubmission(), nil
}
// GetAttackSimulationInfo gets the attackSimulationInfo property value. If the email is phishing simulation, this field won't be null.
// returns a AttackSimulationInfoable when successful
func (m *EmailThreatSubmission) GetAttackSimulationInfo()(AttackSimulationInfoable) {
    val, err := m.GetBackingStore().Get("attackSimulationInfo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AttackSimulationInfoable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailThreatSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ThreatSubmission.GetFieldDeserializers()
    res["attackSimulationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAttackSimulationInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttackSimulationInfo(val.(AttackSimulationInfoable))
        }
        return nil
    }
    res["internetMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetMessageId(val)
        }
        return nil
    }
    res["originalCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubmissionCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalCategory(val.(*SubmissionCategory))
        }
        return nil
    }
    res["receivedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTime(val)
        }
        return nil
    }
    res["recipientEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientEmailAddress(val)
        }
        return nil
    }
    res["sender"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSender(val)
        }
        return nil
    }
    res["senderIP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderIP(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["tenantAllowOrBlockListAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTenantAllowOrBlockListActionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantAllowOrBlockListAction(val.(TenantAllowOrBlockListActionable))
        }
        return nil
    }
    return res
}
// GetInternetMessageId gets the internetMessageId property value. Specifies the internet message ID of the email being submitted. This information is present in the email header.
// returns a *string when successful
func (m *EmailThreatSubmission) GetInternetMessageId()(*string) {
    val, err := m.GetBackingStore().Get("internetMessageId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOriginalCategory gets the originalCategory property value. The original category of the submission. The possible values are: notJunk, spam, phishing, malware and unkownFutureValue.
// returns a *SubmissionCategory when successful
func (m *EmailThreatSubmission) GetOriginalCategory()(*SubmissionCategory) {
    val, err := m.GetBackingStore().Get("originalCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubmissionCategory)
    }
    return nil
}
// GetReceivedDateTime gets the receivedDateTime property value. Specifies the date and time stamp when the email was received.
// returns a *Time when successful
func (m *EmailThreatSubmission) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("receivedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRecipientEmailAddress gets the recipientEmailAddress property value. Specifies the email address (in smtp format) of the recipient who received the email.
// returns a *string when successful
func (m *EmailThreatSubmission) GetRecipientEmailAddress()(*string) {
    val, err := m.GetBackingStore().Get("recipientEmailAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSender gets the sender property value. Specifies the email address of the sender.
// returns a *string when successful
func (m *EmailThreatSubmission) GetSender()(*string) {
    val, err := m.GetBackingStore().Get("sender")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSenderIP gets the senderIP property value. Specifies the IP address of the sender.
// returns a *string when successful
func (m *EmailThreatSubmission) GetSenderIP()(*string) {
    val, err := m.GetBackingStore().Get("senderIP")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubject gets the subject property value. Specifies the subject of the email.
// returns a *string when successful
func (m *EmailThreatSubmission) GetSubject()(*string) {
    val, err := m.GetBackingStore().Get("subject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTenantAllowOrBlockListAction gets the tenantAllowOrBlockListAction property value. It's used to automatically add allows for the components such as URL, file, sender; which are deemed bad by Microsoft so that similar messages in the future can be allowed.
// returns a TenantAllowOrBlockListActionable when successful
func (m *EmailThreatSubmission) GetTenantAllowOrBlockListAction()(TenantAllowOrBlockListActionable) {
    val, err := m.GetBackingStore().Get("tenantAllowOrBlockListAction")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TenantAllowOrBlockListActionable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EmailThreatSubmission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ThreatSubmission.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("attackSimulationInfo", m.GetAttackSimulationInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internetMessageId", m.GetInternetMessageId())
        if err != nil {
            return err
        }
    }
    if m.GetOriginalCategory() != nil {
        cast := (*m.GetOriginalCategory()).String()
        err = writer.WriteStringValue("originalCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientEmailAddress", m.GetRecipientEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sender", m.GetSender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderIP", m.GetSenderIP())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("tenantAllowOrBlockListAction", m.GetTenantAllowOrBlockListAction())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttackSimulationInfo sets the attackSimulationInfo property value. If the email is phishing simulation, this field won't be null.
func (m *EmailThreatSubmission) SetAttackSimulationInfo(value AttackSimulationInfoable)() {
    err := m.GetBackingStore().Set("attackSimulationInfo", value)
    if err != nil {
        panic(err)
    }
}
// SetInternetMessageId sets the internetMessageId property value. Specifies the internet message ID of the email being submitted. This information is present in the email header.
func (m *EmailThreatSubmission) SetInternetMessageId(value *string)() {
    err := m.GetBackingStore().Set("internetMessageId", value)
    if err != nil {
        panic(err)
    }
}
// SetOriginalCategory sets the originalCategory property value. The original category of the submission. The possible values are: notJunk, spam, phishing, malware and unkownFutureValue.
func (m *EmailThreatSubmission) SetOriginalCategory(value *SubmissionCategory)() {
    err := m.GetBackingStore().Set("originalCategory", value)
    if err != nil {
        panic(err)
    }
}
// SetReceivedDateTime sets the receivedDateTime property value. Specifies the date and time stamp when the email was received.
func (m *EmailThreatSubmission) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("receivedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRecipientEmailAddress sets the recipientEmailAddress property value. Specifies the email address (in smtp format) of the recipient who received the email.
func (m *EmailThreatSubmission) SetRecipientEmailAddress(value *string)() {
    err := m.GetBackingStore().Set("recipientEmailAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetSender sets the sender property value. Specifies the email address of the sender.
func (m *EmailThreatSubmission) SetSender(value *string)() {
    err := m.GetBackingStore().Set("sender", value)
    if err != nil {
        panic(err)
    }
}
// SetSenderIP sets the senderIP property value. Specifies the IP address of the sender.
func (m *EmailThreatSubmission) SetSenderIP(value *string)() {
    err := m.GetBackingStore().Set("senderIP", value)
    if err != nil {
        panic(err)
    }
}
// SetSubject sets the subject property value. Specifies the subject of the email.
func (m *EmailThreatSubmission) SetSubject(value *string)() {
    err := m.GetBackingStore().Set("subject", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantAllowOrBlockListAction sets the tenantAllowOrBlockListAction property value. It's used to automatically add allows for the components such as URL, file, sender; which are deemed bad by Microsoft so that similar messages in the future can be allowed.
func (m *EmailThreatSubmission) SetTenantAllowOrBlockListAction(value TenantAllowOrBlockListActionable)() {
    err := m.GetBackingStore().Set("tenantAllowOrBlockListAction", value)
    if err != nil {
        panic(err)
    }
}
type EmailThreatSubmissionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ThreatSubmissionable
    GetAttackSimulationInfo()(AttackSimulationInfoable)
    GetInternetMessageId()(*string)
    GetOriginalCategory()(*SubmissionCategory)
    GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecipientEmailAddress()(*string)
    GetSender()(*string)
    GetSenderIP()(*string)
    GetSubject()(*string)
    GetTenantAllowOrBlockListAction()(TenantAllowOrBlockListActionable)
    SetAttackSimulationInfo(value AttackSimulationInfoable)()
    SetInternetMessageId(value *string)()
    SetOriginalCategory(value *SubmissionCategory)()
    SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecipientEmailAddress(value *string)()
    SetSender(value *string)()
    SetSenderIP(value *string)()
    SetSubject(value *string)()
    SetTenantAllowOrBlockListAction(value TenantAllowOrBlockListActionable)()
}
