package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// EmailThreatSubmissionPolicy 
type EmailThreatSubmissionPolicy struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewEmailThreatSubmissionPolicy instantiates a new emailThreatSubmissionPolicy and sets the default values.
func NewEmailThreatSubmissionPolicy()(*EmailThreatSubmissionPolicy) {
    m := &EmailThreatSubmissionPolicy{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateEmailThreatSubmissionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailThreatSubmissionPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailThreatSubmissionPolicy(), nil
}
// GetCustomizedNotificationSenderEmailAddress gets the customizedNotificationSenderEmailAddress property value. Specifies the email address of the sender from which email notifications will be sent to end users to inform them whether an email is spam, phish or clean. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetCustomizedNotificationSenderEmailAddress()(*string) {
    val, err := m.GetBackingStore().Get("customizedNotificationSenderEmailAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetCustomizedReportRecipientEmailAddress gets the customizedReportRecipientEmailAddress property value. Specifies the destination where the reported messages from end users will land whenever they report something as phish, junk or not junk. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetCustomizedReportRecipientEmailAddress()(*string) {
    val, err := m.GetBackingStore().Get("customizedReportRecipientEmailAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailThreatSubmissionPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customizedNotificationSenderEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomizedNotificationSenderEmailAddress(val)
        }
        return nil
    }
    res["customizedReportRecipientEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomizedReportRecipientEmailAddress(val)
        }
        return nil
    }
    res["isAlwaysReportEnabledForUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAlwaysReportEnabledForUsers(val)
        }
        return nil
    }
    res["isAskMeEnabledForUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAskMeEnabledForUsers(val)
        }
        return nil
    }
    res["isCustomizedMessageEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomizedMessageEnabled(val)
        }
        return nil
    }
    res["isCustomizedMessageEnabledForPhishing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomizedMessageEnabledForPhishing(val)
        }
        return nil
    }
    res["isCustomizedNotificationSenderEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCustomizedNotificationSenderEnabled(val)
        }
        return nil
    }
    res["isNeverReportEnabledForUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsNeverReportEnabledForUsers(val)
        }
        return nil
    }
    res["isOrganizationBrandingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizationBrandingEnabled(val)
        }
        return nil
    }
    res["isReportFromQuarantineEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReportFromQuarantineEnabled(val)
        }
        return nil
    }
    res["isReportToCustomizedEmailAddressEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReportToCustomizedEmailAddressEnabled(val)
        }
        return nil
    }
    res["isReportToMicrosoftEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReportToMicrosoftEnabled(val)
        }
        return nil
    }
    res["isReviewEmailNotificationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReviewEmailNotificationEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetIsAlwaysReportEnabledForUsers gets the isAlwaysReportEnabledForUsers property value. Indicates whether end users can report a message as spam, phish or junk directly without a confirmation(popup). The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsAlwaysReportEnabledForUsers()(*bool) {
    val, err := m.GetBackingStore().Get("isAlwaysReportEnabledForUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsAskMeEnabledForUsers gets the isAskMeEnabledForUsers property value. Indicates whether end users can confirm using a popup before reporting messages as spam, phish or not junk. The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsAskMeEnabledForUsers()(*bool) {
    val, err := m.GetBackingStore().Get("isAskMeEnabledForUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsCustomizedMessageEnabled gets the isCustomizedMessageEnabled property value. Indicates whether the email notifications sent to end users to inform them if an email is phish, spam or junk is customized or not. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedMessageEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isCustomizedMessageEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsCustomizedMessageEnabledForPhishing gets the isCustomizedMessageEnabledForPhishing property value. If enabled, customized message only shows when email is reported as phishing. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedMessageEnabledForPhishing()(*bool) {
    val, err := m.GetBackingStore().Get("isCustomizedMessageEnabledForPhishing")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsCustomizedNotificationSenderEnabled gets the isCustomizedNotificationSenderEnabled property value. Indicates whether to use the sender email address set using customizedNotificationSenderEmailAddress for sending email notifications to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedNotificationSenderEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isCustomizedNotificationSenderEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsNeverReportEnabledForUsers gets the isNeverReportEnabledForUsers property value. Indicates whether end users can simply move the message from one folder to another based on the action of spam, phish or not junk without actually reporting it. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsNeverReportEnabledForUsers()(*bool) {
    val, err := m.GetBackingStore().Get("isNeverReportEnabledForUsers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsOrganizationBrandingEnabled gets the isOrganizationBrandingEnabled property value. Indicates whether the branding logo should be used in the email notifications sent to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsOrganizationBrandingEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isOrganizationBrandingEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsReportFromQuarantineEnabled gets the isReportFromQuarantineEnabled property value. Indicates whether end users can submit from the quarantine page. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReportFromQuarantineEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isReportFromQuarantineEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsReportToCustomizedEmailAddressEnabled gets the isReportToCustomizedEmailAddressEnabled property value. Indicates whether emails reported by end users should be send to the custom mailbox configured using customizedReportRecipientEmailAddress.  The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReportToCustomizedEmailAddressEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isReportToCustomizedEmailAddressEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsReportToMicrosoftEnabled gets the isReportToMicrosoftEnabled property value. If enabled, the email will be sent to Microsoft for analysis. The default value is false. Required for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReportToMicrosoftEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isReportToMicrosoftEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsReviewEmailNotificationEnabled gets the isReviewEmailNotificationEnabled property value. Indicates whether an email notification is sent to the end user who reported the email when it has been reviewed by the admin. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReviewEmailNotificationEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isReviewEmailNotificationEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EmailThreatSubmissionPolicy) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EmailThreatSubmissionPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("customizedNotificationSenderEmailAddress", m.GetCustomizedNotificationSenderEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customizedReportRecipientEmailAddress", m.GetCustomizedReportRecipientEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAlwaysReportEnabledForUsers", m.GetIsAlwaysReportEnabledForUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAskMeEnabledForUsers", m.GetIsAskMeEnabledForUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCustomizedMessageEnabled", m.GetIsCustomizedMessageEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCustomizedMessageEnabledForPhishing", m.GetIsCustomizedMessageEnabledForPhishing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCustomizedNotificationSenderEnabled", m.GetIsCustomizedNotificationSenderEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isNeverReportEnabledForUsers", m.GetIsNeverReportEnabledForUsers())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOrganizationBrandingEnabled", m.GetIsOrganizationBrandingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReportFromQuarantineEnabled", m.GetIsReportFromQuarantineEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReportToCustomizedEmailAddressEnabled", m.GetIsReportToCustomizedEmailAddressEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReportToMicrosoftEnabled", m.GetIsReportToMicrosoftEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReviewEmailNotificationEnabled", m.GetIsReviewEmailNotificationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomizedNotificationSenderEmailAddress sets the customizedNotificationSenderEmailAddress property value. Specifies the email address of the sender from which email notifications will be sent to end users to inform them whether an email is spam, phish or clean. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetCustomizedNotificationSenderEmailAddress(value *string)() {
    err := m.GetBackingStore().Set("customizedNotificationSenderEmailAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomizedReportRecipientEmailAddress sets the customizedReportRecipientEmailAddress property value. Specifies the destination where the reported messages from end users will land whenever they report something as phish, junk or not junk. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetCustomizedReportRecipientEmailAddress(value *string)() {
    err := m.GetBackingStore().Set("customizedReportRecipientEmailAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAlwaysReportEnabledForUsers sets the isAlwaysReportEnabledForUsers property value. Indicates whether end users can report a message as spam, phish or junk directly without a confirmation(popup). The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsAlwaysReportEnabledForUsers(value *bool)() {
    err := m.GetBackingStore().Set("isAlwaysReportEnabledForUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetIsAskMeEnabledForUsers sets the isAskMeEnabledForUsers property value. Indicates whether end users can confirm using a popup before reporting messages as spam, phish or not junk. The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsAskMeEnabledForUsers(value *bool)() {
    err := m.GetBackingStore().Set("isAskMeEnabledForUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCustomizedMessageEnabled sets the isCustomizedMessageEnabled property value. Indicates whether the email notifications sent to end users to inform them if an email is phish, spam or junk is customized or not. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedMessageEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isCustomizedMessageEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCustomizedMessageEnabledForPhishing sets the isCustomizedMessageEnabledForPhishing property value. If enabled, customized message only shows when email is reported as phishing. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedMessageEnabledForPhishing(value *bool)() {
    err := m.GetBackingStore().Set("isCustomizedMessageEnabledForPhishing", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCustomizedNotificationSenderEnabled sets the isCustomizedNotificationSenderEnabled property value. Indicates whether to use the sender email address set using customizedNotificationSenderEmailAddress for sending email notifications to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedNotificationSenderEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isCustomizedNotificationSenderEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsNeverReportEnabledForUsers sets the isNeverReportEnabledForUsers property value. Indicates whether end users can simply move the message from one folder to another based on the action of spam, phish or not junk without actually reporting it. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsNeverReportEnabledForUsers(value *bool)() {
    err := m.GetBackingStore().Set("isNeverReportEnabledForUsers", value)
    if err != nil {
        panic(err)
    }
}
// SetIsOrganizationBrandingEnabled sets the isOrganizationBrandingEnabled property value. Indicates whether the branding logo should be used in the email notifications sent to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsOrganizationBrandingEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isOrganizationBrandingEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsReportFromQuarantineEnabled sets the isReportFromQuarantineEnabled property value. Indicates whether end users can submit from the quarantine page. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReportFromQuarantineEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isReportFromQuarantineEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsReportToCustomizedEmailAddressEnabled sets the isReportToCustomizedEmailAddressEnabled property value. Indicates whether emails reported by end users should be send to the custom mailbox configured using customizedReportRecipientEmailAddress.  The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReportToCustomizedEmailAddressEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isReportToCustomizedEmailAddressEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsReportToMicrosoftEnabled sets the isReportToMicrosoftEnabled property value. If enabled, the email will be sent to Microsoft for analysis. The default value is false. Required for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReportToMicrosoftEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isReportToMicrosoftEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsReviewEmailNotificationEnabled sets the isReviewEmailNotificationEnabled property value. Indicates whether an email notification is sent to the end user who reported the email when it has been reviewed by the admin. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReviewEmailNotificationEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isReviewEmailNotificationEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EmailThreatSubmissionPolicy) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// EmailThreatSubmissionPolicyable 
type EmailThreatSubmissionPolicyable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomizedNotificationSenderEmailAddress()(*string)
    GetCustomizedReportRecipientEmailAddress()(*string)
    GetIsAlwaysReportEnabledForUsers()(*bool)
    GetIsAskMeEnabledForUsers()(*bool)
    GetIsCustomizedMessageEnabled()(*bool)
    GetIsCustomizedMessageEnabledForPhishing()(*bool)
    GetIsCustomizedNotificationSenderEnabled()(*bool)
    GetIsNeverReportEnabledForUsers()(*bool)
    GetIsOrganizationBrandingEnabled()(*bool)
    GetIsReportFromQuarantineEnabled()(*bool)
    GetIsReportToCustomizedEmailAddressEnabled()(*bool)
    GetIsReportToMicrosoftEnabled()(*bool)
    GetIsReviewEmailNotificationEnabled()(*bool)
    GetOdataType()(*string)
    SetCustomizedNotificationSenderEmailAddress(value *string)()
    SetCustomizedReportRecipientEmailAddress(value *string)()
    SetIsAlwaysReportEnabledForUsers(value *bool)()
    SetIsAskMeEnabledForUsers(value *bool)()
    SetIsCustomizedMessageEnabled(value *bool)()
    SetIsCustomizedMessageEnabledForPhishing(value *bool)()
    SetIsCustomizedNotificationSenderEnabled(value *bool)()
    SetIsNeverReportEnabledForUsers(value *bool)()
    SetIsOrganizationBrandingEnabled(value *bool)()
    SetIsReportFromQuarantineEnabled(value *bool)()
    SetIsReportToCustomizedEmailAddressEnabled(value *bool)()
    SetIsReportToMicrosoftEnabled(value *bool)()
    SetIsReviewEmailNotificationEnabled(value *bool)()
    SetOdataType(value *string)()
}
