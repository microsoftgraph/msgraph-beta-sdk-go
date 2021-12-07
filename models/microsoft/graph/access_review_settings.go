package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewSettings 
type AccessReviewSettings struct {
    // Indicates whether showing recommendations to reviewers is enabled.
    accessRecommendationsEnabled *bool;
    // The number of days of user activities to show to reviewers.
    activityDurationInDays *int32;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether the auto-apply capability, to automatically change the target object access resource, is enabled.  If not enabled, a user must, after the review completes, apply the access review.
    autoApplyReviewResultsEnabled *bool;
    // Indicates whether a decision should be set if the reviewer did not supply one. For use when auto-apply is enabled. If you don't want to have a review decision recorded unless the reviewer makes an explicit choice, set it to false.
    autoReviewEnabled *bool;
    // Detailed settings for how the feature should set the review decision. For use when auto-apply is enabled.
    autoReviewSettings *AutoReviewSettings;
    // Indicates whether reviewers are required to provide a justification when reviewing access.
    justificationRequiredOnApproval *bool;
    // Indicates whether sending mails to reviewers and the review creator is enabled.
    mailNotificationsEnabled *bool;
    // Detailed settings for recurrence.
    recurrenceSettings *AccessReviewRecurrenceSettings;
    // Indicates whether sending reminder emails to reviewers is enabled.
    remindersEnabled *bool;
}
// NewAccessReviewSettings instantiates a new accessReviewSettings and sets the default values.
func NewAccessReviewSettings()(*AccessReviewSettings) {
    m := &AccessReviewSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAccessRecommendationsEnabled gets the accessRecommendationsEnabled property value. Indicates whether showing recommendations to reviewers is enabled.
func (m *AccessReviewSettings) GetAccessRecommendationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accessRecommendationsEnabled
    }
}
// GetActivityDurationInDays gets the activityDurationInDays property value. The number of days of user activities to show to reviewers.
func (m *AccessReviewSettings) GetActivityDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activityDurationInDays
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAutoApplyReviewResultsEnabled gets the autoApplyReviewResultsEnabled property value. Indicates whether the auto-apply capability, to automatically change the target object access resource, is enabled.  If not enabled, a user must, after the review completes, apply the access review.
func (m *AccessReviewSettings) GetAutoApplyReviewResultsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoApplyReviewResultsEnabled
    }
}
// GetAutoReviewEnabled gets the autoReviewEnabled property value. Indicates whether a decision should be set if the reviewer did not supply one. For use when auto-apply is enabled. If you don't want to have a review decision recorded unless the reviewer makes an explicit choice, set it to false.
func (m *AccessReviewSettings) GetAutoReviewEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoReviewEnabled
    }
}
// GetAutoReviewSettings gets the autoReviewSettings property value. Detailed settings for how the feature should set the review decision. For use when auto-apply is enabled.
func (m *AccessReviewSettings) GetAutoReviewSettings()(*AutoReviewSettings) {
    if m == nil {
        return nil
    } else {
        return m.autoReviewSettings
    }
}
// GetJustificationRequiredOnApproval gets the justificationRequiredOnApproval property value. Indicates whether reviewers are required to provide a justification when reviewing access.
func (m *AccessReviewSettings) GetJustificationRequiredOnApproval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.justificationRequiredOnApproval
    }
}
// GetMailNotificationsEnabled gets the mailNotificationsEnabled property value. Indicates whether sending mails to reviewers and the review creator is enabled.
func (m *AccessReviewSettings) GetMailNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mailNotificationsEnabled
    }
}
// GetRecurrenceSettings gets the recurrenceSettings property value. Detailed settings for recurrence.
func (m *AccessReviewSettings) GetRecurrenceSettings()(*AccessReviewRecurrenceSettings) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceSettings
    }
}
// GetRemindersEnabled gets the remindersEnabled property value. Indicates whether sending reminder emails to reviewers is enabled.
func (m *AccessReviewSettings) GetRemindersEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.remindersEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessRecommendationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessRecommendationsEnabled(val)
        }
        return nil
    }
    res["activityDurationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDurationInDays(val)
        }
        return nil
    }
    res["autoApplyReviewResultsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoApplyReviewResultsEnabled(val)
        }
        return nil
    }
    res["autoReviewEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoReviewEnabled(val)
        }
        return nil
    }
    res["autoReviewSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAutoReviewSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoReviewSettings(val.(*AutoReviewSettings))
        }
        return nil
    }
    res["justificationRequiredOnApproval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustificationRequiredOnApproval(val)
        }
        return nil
    }
    res["mailNotificationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNotificationsEnabled(val)
        }
        return nil
    }
    res["recurrenceSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewRecurrenceSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceSettings(val.(*AccessReviewRecurrenceSettings))
        }
        return nil
    }
    res["remindersEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemindersEnabled(val)
        }
        return nil
    }
    return res
}
func (m *AccessReviewSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AccessReviewSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("accessRecommendationsEnabled", m.GetAccessRecommendationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("activityDurationInDays", m.GetActivityDurationInDays())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("autoApplyReviewResultsEnabled", m.GetAutoApplyReviewResultsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("autoReviewEnabled", m.GetAutoReviewEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("autoReviewSettings", m.GetAutoReviewSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("justificationRequiredOnApproval", m.GetJustificationRequiredOnApproval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("mailNotificationsEnabled", m.GetMailNotificationsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recurrenceSettings", m.GetRecurrenceSettings())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("remindersEnabled", m.GetRemindersEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessRecommendationsEnabled sets the accessRecommendationsEnabled property value. Indicates whether showing recommendations to reviewers is enabled.
func (m *AccessReviewSettings) SetAccessRecommendationsEnabled(value *bool)() {
    if m != nil {
        m.accessRecommendationsEnabled = value
    }
}
// SetActivityDurationInDays sets the activityDurationInDays property value. The number of days of user activities to show to reviewers.
func (m *AccessReviewSettings) SetActivityDurationInDays(value *int32)() {
    if m != nil {
        m.activityDurationInDays = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAutoApplyReviewResultsEnabled sets the autoApplyReviewResultsEnabled property value. Indicates whether the auto-apply capability, to automatically change the target object access resource, is enabled.  If not enabled, a user must, after the review completes, apply the access review.
func (m *AccessReviewSettings) SetAutoApplyReviewResultsEnabled(value *bool)() {
    if m != nil {
        m.autoApplyReviewResultsEnabled = value
    }
}
// SetAutoReviewEnabled sets the autoReviewEnabled property value. Indicates whether a decision should be set if the reviewer did not supply one. For use when auto-apply is enabled. If you don't want to have a review decision recorded unless the reviewer makes an explicit choice, set it to false.
func (m *AccessReviewSettings) SetAutoReviewEnabled(value *bool)() {
    if m != nil {
        m.autoReviewEnabled = value
    }
}
// SetAutoReviewSettings sets the autoReviewSettings property value. Detailed settings for how the feature should set the review decision. For use when auto-apply is enabled.
func (m *AccessReviewSettings) SetAutoReviewSettings(value *AutoReviewSettings)() {
    if m != nil {
        m.autoReviewSettings = value
    }
}
// SetJustificationRequiredOnApproval sets the justificationRequiredOnApproval property value. Indicates whether reviewers are required to provide a justification when reviewing access.
func (m *AccessReviewSettings) SetJustificationRequiredOnApproval(value *bool)() {
    if m != nil {
        m.justificationRequiredOnApproval = value
    }
}
// SetMailNotificationsEnabled sets the mailNotificationsEnabled property value. Indicates whether sending mails to reviewers and the review creator is enabled.
func (m *AccessReviewSettings) SetMailNotificationsEnabled(value *bool)() {
    if m != nil {
        m.mailNotificationsEnabled = value
    }
}
// SetRecurrenceSettings sets the recurrenceSettings property value. Detailed settings for recurrence.
func (m *AccessReviewSettings) SetRecurrenceSettings(value *AccessReviewRecurrenceSettings)() {
    if m != nil {
        m.recurrenceSettings = value
    }
}
// SetRemindersEnabled sets the remindersEnabled property value. Indicates whether sending reminder emails to reviewers is enabled.
func (m *AccessReviewSettings) SetRemindersEnabled(value *bool)() {
    if m != nil {
        m.remindersEnabled = value
    }
}
