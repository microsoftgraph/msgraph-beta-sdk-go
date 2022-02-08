package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceManagementComplianceActionItem 
type DeviceManagementComplianceActionItem struct {
    Entity
    // What action to take. Possible values are: noAction, notification, block, retire, wipe, removeResourceAccessProfiles, pushNotification, remoteLock.
    actionType *DeviceManagementComplianceActionType;
    // Number of hours to wait till the action will be enforced. Valid values 0 to 8760
    gracePeriodHours *int32;
    // A list of group IDs to speicify who to CC this notification message to. This collection can contain a maximum of 100 elements.
    notificationMessageCCList []string;
    // What notification Message template to use
    notificationTemplateId *string;
}
// NewDeviceManagementComplianceActionItem instantiates a new deviceManagementComplianceActionItem and sets the default values.
func NewDeviceManagementComplianceActionItem()(*DeviceManagementComplianceActionItem) {
    m := &DeviceManagementComplianceActionItem{
        Entity: *NewEntity(),
    }
    return m
}
// GetActionType gets the actionType property value. What action to take. Possible values are: noAction, notification, block, retire, wipe, removeResourceAccessProfiles, pushNotification, remoteLock.
func (m *DeviceManagementComplianceActionItem) GetActionType()(*DeviceManagementComplianceActionType) {
    if m == nil {
        return nil
    } else {
        return m.actionType
    }
}
// GetGracePeriodHours gets the gracePeriodHours property value. Number of hours to wait till the action will be enforced. Valid values 0 to 8760
func (m *DeviceManagementComplianceActionItem) GetGracePeriodHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.gracePeriodHours
    }
}
// GetNotificationMessageCCList gets the notificationMessageCCList property value. A list of group IDs to speicify who to CC this notification message to. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceActionItem) GetNotificationMessageCCList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notificationMessageCCList
    }
}
// GetNotificationTemplateId gets the notificationTemplateId property value. What notification Message template to use
func (m *DeviceManagementComplianceActionItem) GetNotificationTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTemplateId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComplianceActionItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementComplianceActionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionType(val.(*DeviceManagementComplianceActionType))
        }
        return nil
    }
    res["gracePeriodHours"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGracePeriodHours(val)
        }
        return nil
    }
    res["notificationMessageCCList"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotificationMessageCCList(res)
        }
        return nil
    }
    res["notificationTemplateId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationTemplateId(val)
        }
        return nil
    }
    return res
}
func (m *DeviceManagementComplianceActionItem) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceManagementComplianceActionItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionType() != nil {
        cast := (*m.GetActionType()).String()
        err = writer.WriteStringValue("actionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("gracePeriodHours", m.GetGracePeriodHours())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationMessageCCList() != nil {
        err = writer.WriteCollectionOfStringValues("notificationMessageCCList", m.GetNotificationMessageCCList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationTemplateId", m.GetNotificationTemplateId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionType sets the actionType property value. What action to take. Possible values are: noAction, notification, block, retire, wipe, removeResourceAccessProfiles, pushNotification, remoteLock.
func (m *DeviceManagementComplianceActionItem) SetActionType(value *DeviceManagementComplianceActionType)() {
    if m != nil {
        m.actionType = value
    }
}
// SetGracePeriodHours sets the gracePeriodHours property value. Number of hours to wait till the action will be enforced. Valid values 0 to 8760
func (m *DeviceManagementComplianceActionItem) SetGracePeriodHours(value *int32)() {
    if m != nil {
        m.gracePeriodHours = value
    }
}
// SetNotificationMessageCCList sets the notificationMessageCCList property value. A list of group IDs to speicify who to CC this notification message to. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceActionItem) SetNotificationMessageCCList(value []string)() {
    if m != nil {
        m.notificationMessageCCList = value
    }
}
// SetNotificationTemplateId sets the notificationTemplateId property value. What notification Message template to use
func (m *DeviceManagementComplianceActionItem) SetNotificationTemplateId(value *string)() {
    if m != nil {
        m.notificationTemplateId = value
    }
}
