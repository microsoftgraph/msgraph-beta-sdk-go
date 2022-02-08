package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Presence 
type Presence struct {
    Entity
    // The supplemental information to a user's availability. Possible values are Available, Away, BeRightBack, Busy, DoNotDisturb, InACall, InAConferenceCall, Inactive, InAMeeting, Offline, OffWork, OutOfOffice, PresenceUnknown, Presenting, UrgentInterruptionsOnly.
    activity *string;
    // The base presence information for a user. Possible values are Available, AvailableIdle,  Away, BeRightBack, Busy, BusyIdle, DoNotDisturb, Offline, PresenceUnknown
    availability *string;
    // The out of office settings for a user.
    outOfOfficeSettings *OutOfOfficeSettings;
}
// NewPresence instantiates a new presence and sets the default values.
func NewPresence()(*Presence) {
    m := &Presence{
        Entity: *NewEntity(),
    }
    return m
}
// GetActivity gets the activity property value. The supplemental information to a user's availability. Possible values are Available, Away, BeRightBack, Busy, DoNotDisturb, InACall, InAConferenceCall, Inactive, InAMeeting, Offline, OffWork, OutOfOffice, PresenceUnknown, Presenting, UrgentInterruptionsOnly.
func (m *Presence) GetActivity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activity
    }
}
// GetAvailability gets the availability property value. The base presence information for a user. Possible values are Available, AvailableIdle,  Away, BeRightBack, Busy, BusyIdle, DoNotDisturb, Offline, PresenceUnknown
func (m *Presence) GetAvailability()(*string) {
    if m == nil {
        return nil
    } else {
        return m.availability
    }
}
// GetOutOfOfficeSettings gets the outOfOfficeSettings property value. The out of office settings for a user.
func (m *Presence) GetOutOfOfficeSettings()(*OutOfOfficeSettings) {
    if m == nil {
        return nil
    } else {
        return m.outOfOfficeSettings
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Presence) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val)
        }
        return nil
    }
    res["availability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailability(val)
        }
        return nil
    }
    res["outOfOfficeSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOutOfOfficeSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOutOfOfficeSettings(val.(*OutOfOfficeSettings))
        }
        return nil
    }
    return res
}
func (m *Presence) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Presence) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("availability", m.GetAvailability())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("outOfOfficeSettings", m.GetOutOfOfficeSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivity sets the activity property value. The supplemental information to a user's availability. Possible values are Available, Away, BeRightBack, Busy, DoNotDisturb, InACall, InAConferenceCall, Inactive, InAMeeting, Offline, OffWork, OutOfOffice, PresenceUnknown, Presenting, UrgentInterruptionsOnly.
func (m *Presence) SetActivity(value *string)() {
    if m != nil {
        m.activity = value
    }
}
// SetAvailability sets the availability property value. The base presence information for a user. Possible values are Available, AvailableIdle,  Away, BeRightBack, Busy, BusyIdle, DoNotDisturb, Offline, PresenceUnknown
func (m *Presence) SetAvailability(value *string)() {
    if m != nil {
        m.availability = value
    }
}
// SetOutOfOfficeSettings sets the outOfOfficeSettings property value. The out of office settings for a user.
func (m *Presence) SetOutOfOfficeSettings(value *OutOfOfficeSettings)() {
    if m != nil {
        m.outOfOfficeSettings = value
    }
}
