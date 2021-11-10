package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SharedEmailDomainInvitation struct {
    Entity
    // 
    expiryTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    invitationDomain *string;
    // 
    invitationStatus *string;
}
// Instantiates a new sharedEmailDomainInvitation and sets the default values.
func NewSharedEmailDomainInvitation()(*SharedEmailDomainInvitation) {
    m := &SharedEmailDomainInvitation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the expiryTime property value. 
func (m *SharedEmailDomainInvitation) GetExpiryTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiryTime
    }
}
// Gets the invitationDomain property value. 
func (m *SharedEmailDomainInvitation) GetInvitationDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitationDomain
    }
}
// Gets the invitationStatus property value. 
func (m *SharedEmailDomainInvitation) GetInvitationStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitationStatus
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the expiryTime property value. 
// Parameters:
//  - value : Value to set for the expiryTime property.
func (m *SharedEmailDomainInvitation) SetExpiryTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiryTime = value
}
// Sets the invitationDomain property value. 
// Parameters:
//  - value : Value to set for the invitationDomain property.
func (m *SharedEmailDomainInvitation) SetInvitationDomain(value *string)() {
    m.invitationDomain = value
}
// Sets the invitationStatus property value. 
// Parameters:
//  - value : Value to set for the invitationStatus property.
func (m *SharedEmailDomainInvitation) SetInvitationStatus(value *string)() {
    m.invitationStatus = value
}
