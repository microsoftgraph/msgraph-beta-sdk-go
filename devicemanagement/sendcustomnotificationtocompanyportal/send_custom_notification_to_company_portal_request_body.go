package sendcustomnotificationtocompanyportal

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SendCustomNotificationToCompanyPortalRequestBody 
type SendCustomNotificationToCompanyPortalRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    groupsToNotify []string;
    // 
    notificationBody *string;
    // 
    notificationTitle *string;
}
// NewSendCustomNotificationToCompanyPortalRequestBody instantiates a new sendCustomNotificationToCompanyPortalRequestBody and sets the default values.
func NewSendCustomNotificationToCompanyPortalRequestBody()(*SendCustomNotificationToCompanyPortalRequestBody) {
    m := &SendCustomNotificationToCompanyPortalRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SendCustomNotificationToCompanyPortalRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetGroupsToNotify gets the groupsToNotify property value. 
func (m *SendCustomNotificationToCompanyPortalRequestBody) GetGroupsToNotify()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupsToNotify
    }
}
// GetNotificationBody gets the notificationBody property value. 
func (m *SendCustomNotificationToCompanyPortalRequestBody) GetNotificationBody()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationBody
    }
}
// GetNotificationTitle gets the notificationTitle property value. 
func (m *SendCustomNotificationToCompanyPortalRequestBody) GetNotificationTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTitle
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SendCustomNotificationToCompanyPortalRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["groupsToNotify"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupsToNotify(res)
        }
        return nil
    }
    res["notificationBody"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationBody(val)
        }
        return nil
    }
    res["notificationTitle"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTitle(val)
        }
        return nil
    }
    return res
}
func (m *SendCustomNotificationToCompanyPortalRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SendCustomNotificationToCompanyPortalRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetGroupsToNotify() != nil {
        err := writer.WriteCollectionOfStringValues("groupsToNotify", m.GetGroupsToNotify())
        if err != nil {
            return err
        }
    }
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SendCustomNotificationToCompanyPortalRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGroupsToNotify sets the groupsToNotify property value. 
func (m *SendCustomNotificationToCompanyPortalRequestBody) SetGroupsToNotify(value []string)() {
    if m != nil {
        m.groupsToNotify = value
    }
}
// SetNotificationBody sets the notificationBody property value. 
func (m *SendCustomNotificationToCompanyPortalRequestBody) SetNotificationBody(value *string)() {
    if m != nil {
        m.notificationBody = value
    }
}
// SetNotificationTitle sets the notificationTitle property value. 
func (m *SendCustomNotificationToCompanyPortalRequestBody) SetNotificationTitle(value *string)() {
    if m != nil {
        m.notificationTitle = value
    }
}
