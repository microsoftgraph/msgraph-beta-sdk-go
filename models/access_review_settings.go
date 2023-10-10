package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AccessReviewSettings 
type AccessReviewSettings struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAccessReviewSettings instantiates a new accessReviewSettings and sets the default values.
func NewAccessReviewSettings()(*AccessReviewSettings) {
    m := &AccessReviewSettings{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessReviewSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.businessFlowSettings":
                        return NewBusinessFlowSettings(), nil
                }
            }
        }
    }
    return NewAccessReviewSettings(), nil
}
// GetAccessRecommendationsEnabled gets the accessRecommendationsEnabled property value. Indicates whether showing recommendations to reviewers is enabled.
func (m *AccessReviewSettings) GetAccessRecommendationsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("accessRecommendationsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetActivityDurationInDays gets the activityDurationInDays property value. The number of days of user activities to show to reviewers.
func (m *AccessReviewSettings) GetActivityDurationInDays()(*int32) {
    val, err := m.GetBackingStore().Get("activityDurationInDays")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewSettings) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAutoApplyReviewResultsEnabled gets the autoApplyReviewResultsEnabled property value. Indicates whether the auto-apply capability, to automatically change the target object access resource, is enabled.  If not enabled, a user must, after the review completes, apply the access review.
func (m *AccessReviewSettings) GetAutoApplyReviewResultsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("autoApplyReviewResultsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAutoReviewEnabled gets the autoReviewEnabled property value. Indicates whether a decision should be set if the reviewer didn't supply one. For use when, auto-apply is enabled. If you don't want to have a review decision recorded unless the reviewer makes an explicit choice, set it to false.
func (m *AccessReviewSettings) GetAutoReviewEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("autoReviewEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetAutoReviewSettings gets the autoReviewSettings property value. Detailed settings for how the feature should set the review decision. For use when, auto-apply is enabled.
func (m *AccessReviewSettings) GetAutoReviewSettings()(AutoReviewSettingsable) {
    val, err := m.GetBackingStore().Get("autoReviewSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AutoReviewSettingsable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AccessReviewSettings) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accessRecommendationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessRecommendationsEnabled(val)
        }
        return nil
    }
    res["activityDurationInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDurationInDays(val)
        }
        return nil
    }
    res["autoApplyReviewResultsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoApplyReviewResultsEnabled(val)
        }
        return nil
    }
    res["autoReviewEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoReviewEnabled(val)
        }
        return nil
    }
    res["autoReviewSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutoReviewSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoReviewSettings(val.(AutoReviewSettingsable))
        }
        return nil
    }
    res["justificationRequiredOnApproval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustificationRequiredOnApproval(val)
        }
        return nil
    }
    res["mailNotificationsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNotificationsEnabled(val)
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
    res["recurrenceSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewRecurrenceSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrenceSettings(val.(AccessReviewRecurrenceSettingsable))
        }
        return nil
    }
    res["remindersEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetJustificationRequiredOnApproval gets the justificationRequiredOnApproval property value. Indicates whether reviewers are required to provide a justification when reviewing access.
func (m *AccessReviewSettings) GetJustificationRequiredOnApproval()(*bool) {
    val, err := m.GetBackingStore().Get("justificationRequiredOnApproval")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMailNotificationsEnabled gets the mailNotificationsEnabled property value. Indicates whether sending mails to reviewers and the review creator is enabled.
func (m *AccessReviewSettings) GetMailNotificationsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("mailNotificationsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessReviewSettings) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecurrenceSettings gets the recurrenceSettings property value. Detailed settings for recurrence.
func (m *AccessReviewSettings) GetRecurrenceSettings()(AccessReviewRecurrenceSettingsable) {
    val, err := m.GetBackingStore().Get("recurrenceSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessReviewRecurrenceSettingsable)
    }
    return nil
}
// GetRemindersEnabled gets the remindersEnabled property value. Indicates whether sending reminder emails to reviewers is enabled.
func (m *AccessReviewSettings) GetRemindersEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("remindersEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessReviewSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    err := m.GetBackingStore().Set("accessRecommendationsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetActivityDurationInDays sets the activityDurationInDays property value. The number of days of user activities to show to reviewers.
func (m *AccessReviewSettings) SetActivityDurationInDays(value *int32)() {
    err := m.GetBackingStore().Set("activityDurationInDays", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewSettings) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoApplyReviewResultsEnabled sets the autoApplyReviewResultsEnabled property value. Indicates whether the auto-apply capability, to automatically change the target object access resource, is enabled.  If not enabled, a user must, after the review completes, apply the access review.
func (m *AccessReviewSettings) SetAutoApplyReviewResultsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("autoApplyReviewResultsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoReviewEnabled sets the autoReviewEnabled property value. Indicates whether a decision should be set if the reviewer didn't supply one. For use when, auto-apply is enabled. If you don't want to have a review decision recorded unless the reviewer makes an explicit choice, set it to false.
func (m *AccessReviewSettings) SetAutoReviewEnabled(value *bool)() {
    err := m.GetBackingStore().Set("autoReviewEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetAutoReviewSettings sets the autoReviewSettings property value. Detailed settings for how the feature should set the review decision. For use when, auto-apply is enabled.
func (m *AccessReviewSettings) SetAutoReviewSettings(value AutoReviewSettingsable)() {
    err := m.GetBackingStore().Set("autoReviewSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AccessReviewSettings) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetJustificationRequiredOnApproval sets the justificationRequiredOnApproval property value. Indicates whether reviewers are required to provide a justification when reviewing access.
func (m *AccessReviewSettings) SetJustificationRequiredOnApproval(value *bool)() {
    err := m.GetBackingStore().Set("justificationRequiredOnApproval", value)
    if err != nil {
        panic(err)
    }
}
// SetMailNotificationsEnabled sets the mailNotificationsEnabled property value. Indicates whether sending mails to reviewers and the review creator is enabled.
func (m *AccessReviewSettings) SetMailNotificationsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("mailNotificationsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessReviewSettings) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRecurrenceSettings sets the recurrenceSettings property value. Detailed settings for recurrence.
func (m *AccessReviewSettings) SetRecurrenceSettings(value AccessReviewRecurrenceSettingsable)() {
    err := m.GetBackingStore().Set("recurrenceSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetRemindersEnabled sets the remindersEnabled property value. Indicates whether sending reminder emails to reviewers is enabled.
func (m *AccessReviewSettings) SetRemindersEnabled(value *bool)() {
    err := m.GetBackingStore().Set("remindersEnabled", value)
    if err != nil {
        panic(err)
    }
}
// AccessReviewSettingsable 
type AccessReviewSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessRecommendationsEnabled()(*bool)
    GetActivityDurationInDays()(*int32)
    GetAutoApplyReviewResultsEnabled()(*bool)
    GetAutoReviewEnabled()(*bool)
    GetAutoReviewSettings()(AutoReviewSettingsable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetJustificationRequiredOnApproval()(*bool)
    GetMailNotificationsEnabled()(*bool)
    GetOdataType()(*string)
    GetRecurrenceSettings()(AccessReviewRecurrenceSettingsable)
    GetRemindersEnabled()(*bool)
    SetAccessRecommendationsEnabled(value *bool)()
    SetActivityDurationInDays(value *int32)()
    SetAutoApplyReviewResultsEnabled(value *bool)()
    SetAutoReviewEnabled(value *bool)()
    SetAutoReviewSettings(value AutoReviewSettingsable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetJustificationRequiredOnApproval(value *bool)()
    SetMailNotificationsEnabled(value *bool)()
    SetOdataType(value *string)()
    SetRecurrenceSettings(value AccessReviewRecurrenceSettingsable)()
    SetRemindersEnabled(value *bool)()
}
