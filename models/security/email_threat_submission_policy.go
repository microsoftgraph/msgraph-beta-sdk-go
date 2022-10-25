package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// EmailThreatSubmissionPolicy provides operations to manage the collection of accessReview entities.
type EmailThreatSubmissionPolicy struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Specifies the email address of the sender from which email notifications will be sent to end users to inform them whether an email is spam, phish or clean. The default value is null. Optional for creation.
    customizedNotificationSenderEmailAddress *string
    // Specifies the destination where the reported messages from end users will land whenever they report something as phish, junk or not junk. The default value is null. Optional for creation.
    customizedReportRecipientEmailAddress *string
    // Indicates whether end users can report a message as spam, phish or junk directly without a confirmation(popup). The default value is true.  Optional for creation.
    isAlwaysReportEnabledForUsers *bool
    // Indicates whether end users can confirm using a popup before reporting messages as spam, phish or not junk. The default value is true.  Optional for creation.
    isAskMeEnabledForUsers *bool
    // Indicates whether the email notifications sent to end users to inform them if an email is phish, spam or junk is customized or not. The default value is false. Optional for creation.
    isCustomizedMessageEnabled *bool
    // If enabled, customized message only shows when email is reported as phishing. The default value is false. Optional for creation.
    isCustomizedMessageEnabledForPhishing *bool
    // Indicates whether to use the sender email address set using customizedNotificationSenderEmailAddress for sending email notifications to end users. The default value is false. Optional for creation.
    isCustomizedNotificationSenderEnabled *bool
    // Indicates whether end users can simply move the message from one folder to another based on the action of spam, phish or not junk without actually reporting it. The default value is true. Optional for creation.
    isNeverReportEnabledForUsers *bool
    // Indicates whether the branding logo should be used in the email notifications sent to end users. The default value is false. Optional for creation.
    isOrganizationBrandingEnabled *bool
    // Indicates whether end users can submit from the quarantine page. The default value is true. Optional for creation.
    isReportFromQuarantineEnabled *bool
    // Indicates whether emails reported by end users should be send to the custom mailbox configured using customizedReportRecipientEmailAddress.  The default value is false. Optional for creation.
    isReportToCustomizedEmailAddressEnabled *bool
    // If enabled, the email will be sent to Microsoft for analysis. The default value is false. Required for creation.
    isReportToMicrosoftEnabled *bool
    // Indicates whether an email notification is sent to the end user who reported the email when it has been reviewed by the admin. The default value is false. Optional for creation.
    isReviewEmailNotificationEnabled *bool
}
// NewEmailThreatSubmissionPolicy instantiates a new emailThreatSubmissionPolicy and sets the default values.
func NewEmailThreatSubmissionPolicy()(*EmailThreatSubmissionPolicy) {
    m := &EmailThreatSubmissionPolicy{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.security.emailThreatSubmissionPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEmailThreatSubmissionPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailThreatSubmissionPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailThreatSubmissionPolicy(), nil
}
// GetCustomizedNotificationSenderEmailAddress gets the customizedNotificationSenderEmailAddress property value. Specifies the email address of the sender from which email notifications will be sent to end users to inform them whether an email is spam, phish or clean. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetCustomizedNotificationSenderEmailAddress()(*string) {
    return m.customizedNotificationSenderEmailAddress
}
// GetCustomizedReportRecipientEmailAddress gets the customizedReportRecipientEmailAddress property value. Specifies the destination where the reported messages from end users will land whenever they report something as phish, junk or not junk. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetCustomizedReportRecipientEmailAddress()(*string) {
    return m.customizedReportRecipientEmailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailThreatSubmissionPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["customizedNotificationSenderEmailAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomizedNotificationSenderEmailAddress)
    res["customizedReportRecipientEmailAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetCustomizedReportRecipientEmailAddress)
    res["isAlwaysReportEnabledForUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAlwaysReportEnabledForUsers)
    res["isAskMeEnabledForUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsAskMeEnabledForUsers)
    res["isCustomizedMessageEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsCustomizedMessageEnabled)
    res["isCustomizedMessageEnabledForPhishing"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsCustomizedMessageEnabledForPhishing)
    res["isCustomizedNotificationSenderEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsCustomizedNotificationSenderEnabled)
    res["isNeverReportEnabledForUsers"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsNeverReportEnabledForUsers)
    res["isOrganizationBrandingEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsOrganizationBrandingEnabled)
    res["isReportFromQuarantineEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsReportFromQuarantineEnabled)
    res["isReportToCustomizedEmailAddressEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsReportToCustomizedEmailAddressEnabled)
    res["isReportToMicrosoftEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsReportToMicrosoftEnabled)
    res["isReviewEmailNotificationEnabled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsReviewEmailNotificationEnabled)
    return res
}
// GetIsAlwaysReportEnabledForUsers gets the isAlwaysReportEnabledForUsers property value. Indicates whether end users can report a message as spam, phish or junk directly without a confirmation(popup). The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsAlwaysReportEnabledForUsers()(*bool) {
    return m.isAlwaysReportEnabledForUsers
}
// GetIsAskMeEnabledForUsers gets the isAskMeEnabledForUsers property value. Indicates whether end users can confirm using a popup before reporting messages as spam, phish or not junk. The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsAskMeEnabledForUsers()(*bool) {
    return m.isAskMeEnabledForUsers
}
// GetIsCustomizedMessageEnabled gets the isCustomizedMessageEnabled property value. Indicates whether the email notifications sent to end users to inform them if an email is phish, spam or junk is customized or not. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedMessageEnabled()(*bool) {
    return m.isCustomizedMessageEnabled
}
// GetIsCustomizedMessageEnabledForPhishing gets the isCustomizedMessageEnabledForPhishing property value. If enabled, customized message only shows when email is reported as phishing. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedMessageEnabledForPhishing()(*bool) {
    return m.isCustomizedMessageEnabledForPhishing
}
// GetIsCustomizedNotificationSenderEnabled gets the isCustomizedNotificationSenderEnabled property value. Indicates whether to use the sender email address set using customizedNotificationSenderEmailAddress for sending email notifications to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedNotificationSenderEnabled()(*bool) {
    return m.isCustomizedNotificationSenderEnabled
}
// GetIsNeverReportEnabledForUsers gets the isNeverReportEnabledForUsers property value. Indicates whether end users can simply move the message from one folder to another based on the action of spam, phish or not junk without actually reporting it. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsNeverReportEnabledForUsers()(*bool) {
    return m.isNeverReportEnabledForUsers
}
// GetIsOrganizationBrandingEnabled gets the isOrganizationBrandingEnabled property value. Indicates whether the branding logo should be used in the email notifications sent to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsOrganizationBrandingEnabled()(*bool) {
    return m.isOrganizationBrandingEnabled
}
// GetIsReportFromQuarantineEnabled gets the isReportFromQuarantineEnabled property value. Indicates whether end users can submit from the quarantine page. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReportFromQuarantineEnabled()(*bool) {
    return m.isReportFromQuarantineEnabled
}
// GetIsReportToCustomizedEmailAddressEnabled gets the isReportToCustomizedEmailAddressEnabled property value. Indicates whether emails reported by end users should be send to the custom mailbox configured using customizedReportRecipientEmailAddress.  The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReportToCustomizedEmailAddressEnabled()(*bool) {
    return m.isReportToCustomizedEmailAddressEnabled
}
// GetIsReportToMicrosoftEnabled gets the isReportToMicrosoftEnabled property value. If enabled, the email will be sent to Microsoft for analysis. The default value is false. Required for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReportToMicrosoftEnabled()(*bool) {
    return m.isReportToMicrosoftEnabled
}
// GetIsReviewEmailNotificationEnabled gets the isReviewEmailNotificationEnabled property value. Indicates whether an email notification is sent to the end user who reported the email when it has been reviewed by the admin. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) GetIsReviewEmailNotificationEnabled()(*bool) {
    return m.isReviewEmailNotificationEnabled
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
    return nil
}
// SetCustomizedNotificationSenderEmailAddress sets the customizedNotificationSenderEmailAddress property value. Specifies the email address of the sender from which email notifications will be sent to end users to inform them whether an email is spam, phish or clean. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetCustomizedNotificationSenderEmailAddress(value *string)() {
    m.customizedNotificationSenderEmailAddress = value
}
// SetCustomizedReportRecipientEmailAddress sets the customizedReportRecipientEmailAddress property value. Specifies the destination where the reported messages from end users will land whenever they report something as phish, junk or not junk. The default value is null. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetCustomizedReportRecipientEmailAddress(value *string)() {
    m.customizedReportRecipientEmailAddress = value
}
// SetIsAlwaysReportEnabledForUsers sets the isAlwaysReportEnabledForUsers property value. Indicates whether end users can report a message as spam, phish or junk directly without a confirmation(popup). The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsAlwaysReportEnabledForUsers(value *bool)() {
    m.isAlwaysReportEnabledForUsers = value
}
// SetIsAskMeEnabledForUsers sets the isAskMeEnabledForUsers property value. Indicates whether end users can confirm using a popup before reporting messages as spam, phish or not junk. The default value is true.  Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsAskMeEnabledForUsers(value *bool)() {
    m.isAskMeEnabledForUsers = value
}
// SetIsCustomizedMessageEnabled sets the isCustomizedMessageEnabled property value. Indicates whether the email notifications sent to end users to inform them if an email is phish, spam or junk is customized or not. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedMessageEnabled(value *bool)() {
    m.isCustomizedMessageEnabled = value
}
// SetIsCustomizedMessageEnabledForPhishing sets the isCustomizedMessageEnabledForPhishing property value. If enabled, customized message only shows when email is reported as phishing. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedMessageEnabledForPhishing(value *bool)() {
    m.isCustomizedMessageEnabledForPhishing = value
}
// SetIsCustomizedNotificationSenderEnabled sets the isCustomizedNotificationSenderEnabled property value. Indicates whether to use the sender email address set using customizedNotificationSenderEmailAddress for sending email notifications to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedNotificationSenderEnabled(value *bool)() {
    m.isCustomizedNotificationSenderEnabled = value
}
// SetIsNeverReportEnabledForUsers sets the isNeverReportEnabledForUsers property value. Indicates whether end users can simply move the message from one folder to another based on the action of spam, phish or not junk without actually reporting it. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsNeverReportEnabledForUsers(value *bool)() {
    m.isNeverReportEnabledForUsers = value
}
// SetIsOrganizationBrandingEnabled sets the isOrganizationBrandingEnabled property value. Indicates whether the branding logo should be used in the email notifications sent to end users. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsOrganizationBrandingEnabled(value *bool)() {
    m.isOrganizationBrandingEnabled = value
}
// SetIsReportFromQuarantineEnabled sets the isReportFromQuarantineEnabled property value. Indicates whether end users can submit from the quarantine page. The default value is true. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReportFromQuarantineEnabled(value *bool)() {
    m.isReportFromQuarantineEnabled = value
}
// SetIsReportToCustomizedEmailAddressEnabled sets the isReportToCustomizedEmailAddressEnabled property value. Indicates whether emails reported by end users should be send to the custom mailbox configured using customizedReportRecipientEmailAddress.  The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReportToCustomizedEmailAddressEnabled(value *bool)() {
    m.isReportToCustomizedEmailAddressEnabled = value
}
// SetIsReportToMicrosoftEnabled sets the isReportToMicrosoftEnabled property value. If enabled, the email will be sent to Microsoft for analysis. The default value is false. Required for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReportToMicrosoftEnabled(value *bool)() {
    m.isReportToMicrosoftEnabled = value
}
// SetIsReviewEmailNotificationEnabled sets the isReviewEmailNotificationEnabled property value. Indicates whether an email notification is sent to the end user who reported the email when it has been reviewed by the admin. The default value is false. Optional for creation.
func (m *EmailThreatSubmissionPolicy) SetIsReviewEmailNotificationEnabled(value *bool)() {
    m.isReviewEmailNotificationEnabled = value
}
