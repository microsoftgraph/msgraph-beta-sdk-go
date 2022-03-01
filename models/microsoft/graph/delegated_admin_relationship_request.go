package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DelegatedAdminRelationshipRequest 
type DelegatedAdminRelationshipRequest struct {
    Entity
    // 
    action *DelegatedAdminRelationshipRequestAction;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    status *DelegatedAdminRelationshipRequestStatus;
}
// NewDelegatedAdminRelationshipRequest instantiates a new delegatedAdminRelationshipRequest and sets the default values.
func NewDelegatedAdminRelationshipRequest()(*DelegatedAdminRelationshipRequest) {
    m := &DelegatedAdminRelationshipRequest{
        Entity: *NewEntity(),
    }
    return m
}
// GetAction gets the action property value. 
func (m *DelegatedAdminRelationshipRequest) GetAction()(*DelegatedAdminRelationshipRequestAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *DelegatedAdminRelationshipRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *DelegatedAdminRelationshipRequest) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetStatus gets the status property value. 
func (m *DelegatedAdminRelationshipRequest) GetStatus()(*DelegatedAdminRelationshipRequestStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminRelationshipRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminRelationshipRequestAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*DelegatedAdminRelationshipRequestAction))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminRelationshipRequestStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DelegatedAdminRelationshipRequestStatus))
        }
        return nil
    }
    return res
}
func (m *DelegatedAdminRelationshipRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DelegatedAdminRelationshipRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. 
func (m *DelegatedAdminRelationshipRequest) SetAction(value *DelegatedAdminRelationshipRequestAction)() {
    if m != nil {
        m.action = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *DelegatedAdminRelationshipRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *DelegatedAdminRelationshipRequest) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetStatus sets the status property value. 
func (m *DelegatedAdminRelationshipRequest) SetStatus(value *DelegatedAdminRelationshipRequestStatus)() {
    if m != nil {
        m.status = value
    }
}
