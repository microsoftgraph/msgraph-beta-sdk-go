package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewSettings struct {
    accessRecommendationsEnabled *bool;
    activityDurationInDays *int32;
    additionalData map[string]interface{};
    autoApplyReviewResultsEnabled *bool;
    autoReviewEnabled *bool;
    autoReviewSettings *AutoReviewSettings;
    justificationRequiredOnApproval *bool;
    mailNotificationsEnabled *bool;
    recurrenceSettings *AccessReviewRecurrenceSettings;
    remindersEnabled *bool;
}
func NewAccessReviewSettings()(*AccessReviewSettings) {
    m := &AccessReviewSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AccessReviewSettings) GetAccessRecommendationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.accessRecommendationsEnabled
    }
}
func (m *AccessReviewSettings) GetActivityDurationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.activityDurationInDays
    }
}
func (m *AccessReviewSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AccessReviewSettings) GetAutoApplyReviewResultsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoApplyReviewResultsEnabled
    }
}
func (m *AccessReviewSettings) GetAutoReviewEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.autoReviewEnabled
    }
}
func (m *AccessReviewSettings) GetAutoReviewSettings()(*AutoReviewSettings) {
    if m == nil {
        return nil
    } else {
        return m.autoReviewSettings
    }
}
func (m *AccessReviewSettings) GetJustificationRequiredOnApproval()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.justificationRequiredOnApproval
    }
}
func (m *AccessReviewSettings) GetMailNotificationsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.mailNotificationsEnabled
    }
}
func (m *AccessReviewSettings) GetRecurrenceSettings()(*AccessReviewRecurrenceSettings) {
    if m == nil {
        return nil
    } else {
        return m.recurrenceSettings
    }
}
func (m *AccessReviewSettings) GetRemindersEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.remindersEnabled
    }
}
func (m *AccessReviewSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accessRecommendationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAccessRecommendationsEnabled(val)
        return nil
    }
    res["activityDurationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetActivityDurationInDays(val)
        return nil
    }
    res["autoApplyReviewResultsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAutoApplyReviewResultsEnabled(val)
        return nil
    }
    res["autoReviewEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAutoReviewEnabled(val)
        return nil
    }
    res["autoReviewSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAutoReviewSettings() })
        if err != nil {
            return err
        }
        m.SetAutoReviewSettings(val.(*AutoReviewSettings))
        return nil
    }
    res["justificationRequiredOnApproval"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetJustificationRequiredOnApproval(val)
        return nil
    }
    res["mailNotificationsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetMailNotificationsEnabled(val)
        return nil
    }
    res["recurrenceSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewRecurrenceSettings() })
        if err != nil {
            return err
        }
        m.SetRecurrenceSettings(val.(*AccessReviewRecurrenceSettings))
        return nil
    }
    res["remindersEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRemindersEnabled(val)
        return nil
    }
    return res
}
func (m *AccessReviewSettings) IsNil()(bool) {
    return m == nil
}
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
func (m *AccessReviewSettings) SetAccessRecommendationsEnabled(value *bool)() {
    m.accessRecommendationsEnabled = value
}
func (m *AccessReviewSettings) SetActivityDurationInDays(value *int32)() {
    m.activityDurationInDays = value
}
func (m *AccessReviewSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AccessReviewSettings) SetAutoApplyReviewResultsEnabled(value *bool)() {
    m.autoApplyReviewResultsEnabled = value
}
func (m *AccessReviewSettings) SetAutoReviewEnabled(value *bool)() {
    m.autoReviewEnabled = value
}
func (m *AccessReviewSettings) SetAutoReviewSettings(value *AutoReviewSettings)() {
    m.autoReviewSettings = value
}
func (m *AccessReviewSettings) SetJustificationRequiredOnApproval(value *bool)() {
    m.justificationRequiredOnApproval = value
}
func (m *AccessReviewSettings) SetMailNotificationsEnabled(value *bool)() {
    m.mailNotificationsEnabled = value
}
func (m *AccessReviewSettings) SetRecurrenceSettings(value *AccessReviewRecurrenceSettings)() {
    m.recurrenceSettings = value
}
func (m *AccessReviewSettings) SetRemindersEnabled(value *bool)() {
    m.remindersEnabled = value
}
