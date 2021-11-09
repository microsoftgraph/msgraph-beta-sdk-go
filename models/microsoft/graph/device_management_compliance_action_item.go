package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new deviceManagementComplianceActionItem and sets the default values.
func NewDeviceManagementComplianceActionItem()(*DeviceManagementComplianceActionItem) {
    m := &DeviceManagementComplianceActionItem{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the actionType property value. What action to take. Possible values are: noAction, notification, block, retire, wipe, removeResourceAccessProfiles, pushNotification, remoteLock.
func (m *DeviceManagementComplianceActionItem) GetActionType()(*DeviceManagementComplianceActionType) {
    if m == nil {
        return nil
    } else {
        return m.actionType
    }
}
// Gets the gracePeriodHours property value. Number of hours to wait till the action will be enforced. Valid values 0 to 8760
func (m *DeviceManagementComplianceActionItem) GetGracePeriodHours()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.gracePeriodHours
    }
}
// Gets the notificationMessageCCList property value. A list of group IDs to speicify who to CC this notification message to. This collection can contain a maximum of 100 elements.
func (m *DeviceManagementComplianceActionItem) GetNotificationMessageCCList()([]string) {
    if m == nil {
        return nil
    } else {
        return m.notificationMessageCCList
    }
}
// Gets the notificationTemplateId property value. What notification Message template to use
func (m *DeviceManagementComplianceActionItem) GetNotificationTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationTemplateId
    }
}
// The deserialization information for the current model
func (m *DeviceManagementComplianceActionItem) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementComplianceActionType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(DeviceManagementComplianceActionType)
            m.SetActionType(&cast)
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagementComplianceActionItem) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActionType() != nil {
        cast := m.GetActionType().String()
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
    {
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
// Sets the actionType property value. What action to take. Possible values are: noAction, notification, block, retire, wipe, removeResourceAccessProfiles, pushNotification, remoteLock.
// Parameters:
//  - value : Value to set for the actionType property.
func (m *DeviceManagementComplianceActionItem) SetActionType(value *DeviceManagementComplianceActionType)() {
    m.actionType = value
}
// Sets the gracePeriodHours property value. Number of hours to wait till the action will be enforced. Valid values 0 to 8760
// Parameters:
//  - value : Value to set for the gracePeriodHours property.
func (m *DeviceManagementComplianceActionItem) SetGracePeriodHours(value *int32)() {
    m.gracePeriodHours = value
}
// Sets the notificationMessageCCList property value. A list of group IDs to speicify who to CC this notification message to. This collection can contain a maximum of 100 elements.
// Parameters:
//  - value : Value to set for the notificationMessageCCList property.
func (m *DeviceManagementComplianceActionItem) SetNotificationMessageCCList(value []string)() {
    m.notificationMessageCCList = value
}
// Sets the notificationTemplateId property value. What notification Message template to use
// Parameters:
//  - value : Value to set for the notificationTemplateId property.
func (m *DeviceManagementComplianceActionItem) SetNotificationTemplateId(value *string)() {
    m.notificationTemplateId = value
}
