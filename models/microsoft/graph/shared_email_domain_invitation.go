package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SharedEmailDomainInvitation 
type SharedEmailDomainInvitation struct {
    Entity
    // 
    expiryTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    invitationDomain *string;
    // 
    invitationStatus *string;
}
// NewSharedEmailDomainInvitation instantiates a new sharedEmailDomainInvitation and sets the default values.
func NewSharedEmailDomainInvitation()(*SharedEmailDomainInvitation) {
    m := &SharedEmailDomainInvitation{
        Entity: *NewEntity(),
    }
    return m
}
// GetExpiryTime gets the expiryTime property value. 
func (m *SharedEmailDomainInvitation) GetExpiryTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiryTime
    }
}
// GetInvitationDomain gets the invitationDomain property value. 
func (m *SharedEmailDomainInvitation) GetInvitationDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitationDomain
    }
}
// GetInvitationStatus gets the invitationStatus property value. 
func (m *SharedEmailDomainInvitation) GetInvitationStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitationStatus
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedEmailDomainInvitation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expiryTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiryTime(val)
        }
        return nil
    }
    res["invitationDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitationDomain(val)
        }
        return nil
    }
    res["invitationStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitationStatus(val)
        }
        return nil
    }
    return res
}
func (m *SharedEmailDomainInvitation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SharedEmailDomainInvitation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expiryTime", m.GetExpiryTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invitationDomain", m.GetInvitationDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invitationStatus", m.GetInvitationStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpiryTime sets the expiryTime property value. 
func (m *SharedEmailDomainInvitation) SetExpiryTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expiryTime = value
    }
}
// SetInvitationDomain sets the invitationDomain property value. 
func (m *SharedEmailDomainInvitation) SetInvitationDomain(value *string)() {
    if m != nil {
        m.invitationDomain = value
    }
}
// SetInvitationStatus sets the invitationStatus property value. 
func (m *SharedEmailDomainInvitation) SetInvitationStatus(value *string)() {
    if m != nil {
        m.invitationStatus = value
    }
}
