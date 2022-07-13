package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// EmailThreatSubmissionPolicy provides operations to manage the collection of activityStatistics entities.
type EmailThreatSubmissionPolicy struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The customizedNotificationSenderEmailAddress property
    customizedNotificationSenderEmailAddress *string
    // The customizedReportRecipientEmailAddress property
    customizedReportRecipientEmailAddress *string
    // The isAlwaysReportEnabledForUsers property
    isAlwaysReportEnabledForUsers *bool
    // The isAskMeEnabledForUsers property
    isAskMeEnabledForUsers *bool
    // The isCustomizedMessageEnabled property
    isCustomizedMessageEnabled *bool
    // The isCustomizedMessageEnabledForPhishing property
    isCustomizedMessageEnabledForPhishing *bool
    // The isCustomizedNotificationSenderEnabled property
    isCustomizedNotificationSenderEnabled *bool
    // The isNeverReportEnabledForUsers property
    isNeverReportEnabledForUsers *bool
    // The isOrganizationBrandingEnabled property
    isOrganizationBrandingEnabled *bool
    // The isReportFromQuarantineEnabled property
    isReportFromQuarantineEnabled *bool
    // The isReportToCustomizedEmailAddressEnabled property
    isReportToCustomizedEmailAddressEnabled *bool
    // The isReportToMicrosoftEnabled property
    isReportToMicrosoftEnabled *bool
    // The isReviewEmailNotificationEnabled property
    isReviewEmailNotificationEnabled *bool
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
// GetCustomizedNotificationSenderEmailAddress gets the customizedNotificationSenderEmailAddress property value. The customizedNotificationSenderEmailAddress property
func (m *EmailThreatSubmissionPolicy) GetCustomizedNotificationSenderEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customizedNotificationSenderEmailAddress
    }
}
// GetCustomizedReportRecipientEmailAddress gets the customizedReportRecipientEmailAddress property value. The customizedReportRecipientEmailAddress property
func (m *EmailThreatSubmissionPolicy) GetCustomizedReportRecipientEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customizedReportRecipientEmailAddress
    }
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
    return res
}
// GetIsAlwaysReportEnabledForUsers gets the isAlwaysReportEnabledForUsers property value. The isAlwaysReportEnabledForUsers property
func (m *EmailThreatSubmissionPolicy) GetIsAlwaysReportEnabledForUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAlwaysReportEnabledForUsers
    }
}
// GetIsAskMeEnabledForUsers gets the isAskMeEnabledForUsers property value. The isAskMeEnabledForUsers property
func (m *EmailThreatSubmissionPolicy) GetIsAskMeEnabledForUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAskMeEnabledForUsers
    }
}
// GetIsCustomizedMessageEnabled gets the isCustomizedMessageEnabled property value. The isCustomizedMessageEnabled property
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedMessageEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCustomizedMessageEnabled
    }
}
// GetIsCustomizedMessageEnabledForPhishing gets the isCustomizedMessageEnabledForPhishing property value. The isCustomizedMessageEnabledForPhishing property
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedMessageEnabledForPhishing()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCustomizedMessageEnabledForPhishing
    }
}
// GetIsCustomizedNotificationSenderEnabled gets the isCustomizedNotificationSenderEnabled property value. The isCustomizedNotificationSenderEnabled property
func (m *EmailThreatSubmissionPolicy) GetIsCustomizedNotificationSenderEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isCustomizedNotificationSenderEnabled
    }
}
// GetIsNeverReportEnabledForUsers gets the isNeverReportEnabledForUsers property value. The isNeverReportEnabledForUsers property
func (m *EmailThreatSubmissionPolicy) GetIsNeverReportEnabledForUsers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isNeverReportEnabledForUsers
    }
}
// GetIsOrganizationBrandingEnabled gets the isOrganizationBrandingEnabled property value. The isOrganizationBrandingEnabled property
func (m *EmailThreatSubmissionPolicy) GetIsOrganizationBrandingEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOrganizationBrandingEnabled
    }
}
// GetIsReportFromQuarantineEnabled gets the isReportFromQuarantineEnabled property value. The isReportFromQuarantineEnabled property
func (m *EmailThreatSubmissionPolicy) GetIsReportFromQuarantineEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReportFromQuarantineEnabled
    }
}
// GetIsReportToCustomizedEmailAddressEnabled gets the isReportToCustomizedEmailAddressEnabled property value. The isReportToCustomizedEmailAddressEnabled property
func (m *EmailThreatSubmissionPolicy) GetIsReportToCustomizedEmailAddressEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReportToCustomizedEmailAddressEnabled
    }
}
// GetIsReportToMicrosoftEnabled gets the isReportToMicrosoftEnabled property value. The isReportToMicrosoftEnabled property
func (m *EmailThreatSubmissionPolicy) GetIsReportToMicrosoftEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReportToMicrosoftEnabled
    }
}
// GetIsReviewEmailNotificationEnabled gets the isReviewEmailNotificationEnabled property value. The isReviewEmailNotificationEnabled property
func (m *EmailThreatSubmissionPolicy) GetIsReviewEmailNotificationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReviewEmailNotificationEnabled
    }
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
// SetCustomizedNotificationSenderEmailAddress sets the customizedNotificationSenderEmailAddress property value. The customizedNotificationSenderEmailAddress property
func (m *EmailThreatSubmissionPolicy) SetCustomizedNotificationSenderEmailAddress(value *string)() {
    if m != nil {
        m.customizedNotificationSenderEmailAddress = value
    }
}
// SetCustomizedReportRecipientEmailAddress sets the customizedReportRecipientEmailAddress property value. The customizedReportRecipientEmailAddress property
func (m *EmailThreatSubmissionPolicy) SetCustomizedReportRecipientEmailAddress(value *string)() {
    if m != nil {
        m.customizedReportRecipientEmailAddress = value
    }
}
// SetIsAlwaysReportEnabledForUsers sets the isAlwaysReportEnabledForUsers property value. The isAlwaysReportEnabledForUsers property
func (m *EmailThreatSubmissionPolicy) SetIsAlwaysReportEnabledForUsers(value *bool)() {
    if m != nil {
        m.isAlwaysReportEnabledForUsers = value
    }
}
// SetIsAskMeEnabledForUsers sets the isAskMeEnabledForUsers property value. The isAskMeEnabledForUsers property
func (m *EmailThreatSubmissionPolicy) SetIsAskMeEnabledForUsers(value *bool)() {
    if m != nil {
        m.isAskMeEnabledForUsers = value
    }
}
// SetIsCustomizedMessageEnabled sets the isCustomizedMessageEnabled property value. The isCustomizedMessageEnabled property
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedMessageEnabled(value *bool)() {
    if m != nil {
        m.isCustomizedMessageEnabled = value
    }
}
// SetIsCustomizedMessageEnabledForPhishing sets the isCustomizedMessageEnabledForPhishing property value. The isCustomizedMessageEnabledForPhishing property
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedMessageEnabledForPhishing(value *bool)() {
    if m != nil {
        m.isCustomizedMessageEnabledForPhishing = value
    }
}
// SetIsCustomizedNotificationSenderEnabled sets the isCustomizedNotificationSenderEnabled property value. The isCustomizedNotificationSenderEnabled property
func (m *EmailThreatSubmissionPolicy) SetIsCustomizedNotificationSenderEnabled(value *bool)() {
    if m != nil {
        m.isCustomizedNotificationSenderEnabled = value
    }
}
// SetIsNeverReportEnabledForUsers sets the isNeverReportEnabledForUsers property value. The isNeverReportEnabledForUsers property
func (m *EmailThreatSubmissionPolicy) SetIsNeverReportEnabledForUsers(value *bool)() {
    if m != nil {
        m.isNeverReportEnabledForUsers = value
    }
}
// SetIsOrganizationBrandingEnabled sets the isOrganizationBrandingEnabled property value. The isOrganizationBrandingEnabled property
func (m *EmailThreatSubmissionPolicy) SetIsOrganizationBrandingEnabled(value *bool)() {
    if m != nil {
        m.isOrganizationBrandingEnabled = value
    }
}
// SetIsReportFromQuarantineEnabled sets the isReportFromQuarantineEnabled property value. The isReportFromQuarantineEnabled property
func (m *EmailThreatSubmissionPolicy) SetIsReportFromQuarantineEnabled(value *bool)() {
    if m != nil {
        m.isReportFromQuarantineEnabled = value
    }
}
// SetIsReportToCustomizedEmailAddressEnabled sets the isReportToCustomizedEmailAddressEnabled property value. The isReportToCustomizedEmailAddressEnabled property
func (m *EmailThreatSubmissionPolicy) SetIsReportToCustomizedEmailAddressEnabled(value *bool)() {
    if m != nil {
        m.isReportToCustomizedEmailAddressEnabled = value
    }
}
// SetIsReportToMicrosoftEnabled sets the isReportToMicrosoftEnabled property value. The isReportToMicrosoftEnabled property
func (m *EmailThreatSubmissionPolicy) SetIsReportToMicrosoftEnabled(value *bool)() {
    if m != nil {
        m.isReportToMicrosoftEnabled = value
    }
}
// SetIsReviewEmailNotificationEnabled sets the isReviewEmailNotificationEnabled property value. The isReviewEmailNotificationEnabled property
func (m *EmailThreatSubmissionPolicy) SetIsReviewEmailNotificationEnabled(value *bool)() {
    if m != nil {
        m.isReviewEmailNotificationEnabled = value
    }
}
