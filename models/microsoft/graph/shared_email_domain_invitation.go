package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SharedEmailDomainInvitation struct {
    Entity
    expiryTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    invitationDomain *string;
    invitationStatus *string;
}
func NewSharedEmailDomainInvitation()(*SharedEmailDomainInvitation) {
    m := &SharedEmailDomainInvitation{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SharedEmailDomainInvitation) GetExpiryTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expiryTime
    }
}
func (m *SharedEmailDomainInvitation) GetInvitationDomain()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitationDomain
    }
}
func (m *SharedEmailDomainInvitation) GetInvitationStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.invitationStatus
    }
}
func (m *SharedEmailDomainInvitation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expiryTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpiryTime(val)
        return nil
    }
    res["invitationDomain"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInvitationDomain(val)
        return nil
    }
    res["invitationStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetInvitationStatus(val)
        return nil
    }
    return res
}
func (m *SharedEmailDomainInvitation) IsNil()(bool) {
    return m == nil
}
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
func (m *SharedEmailDomainInvitation) SetExpiryTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiryTime = value
}
func (m *SharedEmailDomainInvitation) SetInvitationDomain(value *string)() {
    m.invitationDomain = value
}
func (m *SharedEmailDomainInvitation) SetInvitationStatus(value *string)() {
    m.invitationStatus = value
}
