package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AccessReviewNotificationRecipientItem struct {
    additionalData map[string]interface{};
    notificationRecipientScope *AccessReviewNotificationRecipientScope;
    notificationTemplateType *string;
}
func NewAccessReviewNotificationRecipientItem()(*AccessReviewNotificationRecipientItem) {
    m := &AccessReviewNotificationRecipientItem{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AccessReviewNotificationRecipientItem) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AccessReviewNotificationRecipientItem) GetNotificationRecipientScope()(*AccessReviewNotificationRecipientScope) {
    if m == nil {
        return nil
    } else {
        return m.notificationRecipientScope
    }
}
func (m *AccessReviewNotificationRecipientItem) GetNotificationTemplateType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTemplateType
    }
}
func (m *AccessReviewNotificationRecipientItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["notificationRecipientScope"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessReviewNotificationRecipientScope() })
        if err != nil {
            return err
        }
        m.SetNotificationRecipientScope(val.(*AccessReviewNotificationRecipientScope))
        return nil
    }
    res["notificationTemplateType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationTemplateType(val)
        return nil
    }
    return res
}
func (m *AccessReviewNotificationRecipientItem) IsNil()(bool) {
    return m == nil
}
func (m *AccessReviewNotificationRecipientItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("notificationRecipientScope", m.GetNotificationRecipientScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationTemplateType", m.GetNotificationTemplateType())
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
func (m *AccessReviewNotificationRecipientItem) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AccessReviewNotificationRecipientItem) SetNotificationRecipientScope(value *AccessReviewNotificationRecipientScope)() {
    m.notificationRecipientScope = value
}
func (m *AccessReviewNotificationRecipientItem) SetNotificationTemplateType(value *string)() {
    m.notificationTemplateType = value
}
