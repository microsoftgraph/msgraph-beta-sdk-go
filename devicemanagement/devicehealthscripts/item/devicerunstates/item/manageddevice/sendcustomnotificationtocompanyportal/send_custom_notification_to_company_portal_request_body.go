package sendcustomnotificationtocompanyportal

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SendCustomNotificationToCompanyPortalRequestBody struct {
    additionalData map[string]interface{};
    notificationBody *string;
    notificationTitle *string;
}
func NewSendCustomNotificationToCompanyPortalRequestBody()(*SendCustomNotificationToCompanyPortalRequestBody) {
    m := &SendCustomNotificationToCompanyPortalRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SendCustomNotificationToCompanyPortalRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SendCustomNotificationToCompanyPortalRequestBody) GetNotificationBody()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationBody
    }
}
func (m *SendCustomNotificationToCompanyPortalRequestBody) GetNotificationTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTitle
    }
}
func (m *SendCustomNotificationToCompanyPortalRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["notificationBody"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationBody(val)
        return nil
    }
    res["notificationTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetNotificationTitle(val)
        return nil
    }
    return res
}
func (m *SendCustomNotificationToCompanyPortalRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *SendCustomNotificationToCompanyPortalRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("notificationBody", m.GetNotificationBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("notificationTitle", m.GetNotificationTitle())
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
func (m *SendCustomNotificationToCompanyPortalRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SendCustomNotificationToCompanyPortalRequestBody) SetNotificationBody(value *string)() {
    m.notificationBody = value
}
func (m *SendCustomNotificationToCompanyPortalRequestBody) SetNotificationTitle(value *string)() {
    m.notificationTitle = value
}
