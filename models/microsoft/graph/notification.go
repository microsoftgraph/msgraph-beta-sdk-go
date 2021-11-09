package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Notification struct {
    Entity
    // Sets how long (in seconds) this notification content will stay in each platform’s notification viewer. For example, when the notification is delivered to a Windows device, the value of this property is passed on to ToastNotification.ExpirationTime, which determines how long the toast notification will stay in the user’s Windows Action Center.
    displayTimeToLive *int32;
    // Sets a UTC expiration date and time on a user notification using ISO 8601 format (for example, midnight UTC on Jan 1, 2019 would look like this: '2019-01-01T00:00:00Z'). When time is up, the notification is removed from the Microsoft Graph notification feed store completely and is no longer part of notification history. Max value is 30 days.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The name of the group that this notification belongs to. It is set by the developer for the purpose of grouping notifications together.
    groupName *string;
    // 
    payload *PayloadTypes;
    // Indicates the priority of a raw user notification. Visual notifications are sent with high priority by default. Valid values are None, High and Low.
    priority *Priority;
    // Represents the host name of the app to which the calling service wants to post the notification, for the given user. If targeting web endpoints (see targetPolicy.platformTypes), ensure that targetHostName is the same as the name used when creating a subscription on the client side within the application JSON property.
    targetHostName *string;
    // Target policy object handles notification delivery policy for endpoint types that should be targeted (Windows, iOS, Android and WebPush) for the given user.
    targetPolicy *TargetPolicyEndpoints;
}
// Instantiates a new notification and sets the default values.
func NewNotification()(*Notification) {
    m := &Notification{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayTimeToLive property value. Sets how long (in seconds) this notification content will stay in each platform’s notification viewer. For example, when the notification is delivered to a Windows device, the value of this property is passed on to ToastNotification.ExpirationTime, which determines how long the toast notification will stay in the user’s Windows Action Center.
func (m *Notification) GetDisplayTimeToLive()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.displayTimeToLive
    }
}
// Gets the expirationDateTime property value. Sets a UTC expiration date and time on a user notification using ISO 8601 format (for example, midnight UTC on Jan 1, 2019 would look like this: '2019-01-01T00:00:00Z'). When time is up, the notification is removed from the Microsoft Graph notification feed store completely and is no longer part of notification history. Max value is 30 days.
func (m *Notification) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the groupName property value. The name of the group that this notification belongs to. It is set by the developer for the purpose of grouping notifications together.
func (m *Notification) GetGroupName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupName
    }
}
// Gets the payload property value. 
func (m *Notification) GetPayload()(*PayloadTypes) {
    if m == nil {
        return nil
    } else {
        return m.payload
    }
}
// Gets the priority property value. Indicates the priority of a raw user notification. Visual notifications are sent with high priority by default. Valid values are None, High and Low.
func (m *Notification) GetPriority()(*Priority) {
    if m == nil {
        return nil
    } else {
        return m.priority
    }
}
// Gets the targetHostName property value. Represents the host name of the app to which the calling service wants to post the notification, for the given user. If targeting web endpoints (see targetPolicy.platformTypes), ensure that targetHostName is the same as the name used when creating a subscription on the client side within the application JSON property.
func (m *Notification) GetTargetHostName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetHostName
    }
}
// Gets the targetPolicy property value. Target policy object handles notification delivery policy for endpoint types that should be targeted (Windows, iOS, Android and WebPush) for the given user.
func (m *Notification) GetTargetPolicy()(*TargetPolicyEndpoints) {
    if m == nil {
        return nil
    } else {
        return m.targetPolicy
    }
}
// The deserialization information for the current model
func (m *Notification) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayTimeToLive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayTimeToLive(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["groupName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupName(val)
        }
        return nil
    }
    res["payload"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPayloadTypes() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val.(*PayloadTypes))
        }
        return nil
    }
    res["priority"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePriority)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(Priority)
            m.SetPriority(&cast)
        }
        return nil
    }
    res["targetHostName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetHostName(val)
        }
        return nil
    }
    res["targetPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTargetPolicyEndpoints() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetPolicy(val.(*TargetPolicyEndpoints))
        }
        return nil
    }
    return res
}
func (m *Notification) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Notification) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("displayTimeToLive", m.GetDisplayTimeToLive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupName", m.GetGroupName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    if m.GetPriority() != nil {
        cast := m.GetPriority().String()
        err = writer.WriteStringValue("priority", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetHostName", m.GetTargetHostName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetPolicy", m.GetTargetPolicy())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayTimeToLive property value. Sets how long (in seconds) this notification content will stay in each platform’s notification viewer. For example, when the notification is delivered to a Windows device, the value of this property is passed on to ToastNotification.ExpirationTime, which determines how long the toast notification will stay in the user’s Windows Action Center.
// Parameters:
//  - value : Value to set for the displayTimeToLive property.
func (m *Notification) SetDisplayTimeToLive(value *int32)() {
    m.displayTimeToLive = value
}
// Sets the expirationDateTime property value. Sets a UTC expiration date and time on a user notification using ISO 8601 format (for example, midnight UTC on Jan 1, 2019 would look like this: '2019-01-01T00:00:00Z'). When time is up, the notification is removed from the Microsoft Graph notification feed store completely and is no longer part of notification history. Max value is 30 days.
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *Notification) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the groupName property value. The name of the group that this notification belongs to. It is set by the developer for the purpose of grouping notifications together.
// Parameters:
//  - value : Value to set for the groupName property.
func (m *Notification) SetGroupName(value *string)() {
    m.groupName = value
}
// Sets the payload property value. 
// Parameters:
//  - value : Value to set for the payload property.
func (m *Notification) SetPayload(value *PayloadTypes)() {
    m.payload = value
}
// Sets the priority property value. Indicates the priority of a raw user notification. Visual notifications are sent with high priority by default. Valid values are None, High and Low.
// Parameters:
//  - value : Value to set for the priority property.
func (m *Notification) SetPriority(value *Priority)() {
    m.priority = value
}
// Sets the targetHostName property value. Represents the host name of the app to which the calling service wants to post the notification, for the given user. If targeting web endpoints (see targetPolicy.platformTypes), ensure that targetHostName is the same as the name used when creating a subscription on the client side within the application JSON property.
// Parameters:
//  - value : Value to set for the targetHostName property.
func (m *Notification) SetTargetHostName(value *string)() {
    m.targetHostName = value
}
// Sets the targetPolicy property value. Target policy object handles notification delivery policy for endpoint types that should be targeted (Windows, iOS, Android and WebPush) for the given user.
// Parameters:
//  - value : Value to set for the targetPolicy property.
func (m *Notification) SetTargetPolicy(value *TargetPolicyEndpoints)() {
    m.targetPolicy = value
}
