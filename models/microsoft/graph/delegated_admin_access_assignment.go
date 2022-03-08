package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DelegatedAdminAccessAssignment provides operations to manage the tenantRelationship singleton.
type DelegatedAdminAccessAssignment struct {
    Entity
    // 
    accessContainer DelegatedAdminAccessContainerable;
    // 
    accessDetails DelegatedAdminAccessDetailsable;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    status *DelegatedAdminAccessAssignmentStatus;
}
// NewDelegatedAdminAccessAssignment instantiates a new delegatedAdminAccessAssignment and sets the default values.
func NewDelegatedAdminAccessAssignment()(*DelegatedAdminAccessAssignment) {
    m := &DelegatedAdminAccessAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDelegatedAdminAccessAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminAccessAssignmentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDelegatedAdminAccessAssignment(), nil
}
// GetAccessContainer gets the accessContainer property value. 
func (m *DelegatedAdminAccessAssignment) GetAccessContainer()(DelegatedAdminAccessContainerable) {
    if m == nil {
        return nil
    } else {
        return m.accessContainer
    }
}
// GetAccessDetails gets the accessDetails property value. 
func (m *DelegatedAdminAccessAssignment) GetAccessDetails()(DelegatedAdminAccessDetailsable) {
    if m == nil {
        return nil
    } else {
        return m.accessDetails
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *DelegatedAdminAccessAssignment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminAccessAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessContainer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDelegatedAdminAccessContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessContainer(val.(DelegatedAdminAccessContainerable))
        }
        return nil
    }
    res["accessDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDelegatedAdminAccessDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessDetails(val.(DelegatedAdminAccessDetailsable))
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
        val, err := n.GetEnumValue(ParseDelegatedAdminAccessAssignmentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DelegatedAdminAccessAssignmentStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *DelegatedAdminAccessAssignment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetStatus gets the status property value. 
func (m *DelegatedAdminAccessAssignment) GetStatus()(*DelegatedAdminAccessAssignmentStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *DelegatedAdminAccessAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DelegatedAdminAccessAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessContainer", m.GetAccessContainer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessDetails", m.GetAccessDetails())
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
// SetAccessContainer sets the accessContainer property value. 
func (m *DelegatedAdminAccessAssignment) SetAccessContainer(value DelegatedAdminAccessContainerable)() {
    if m != nil {
        m.accessContainer = value
    }
}
// SetAccessDetails sets the accessDetails property value. 
func (m *DelegatedAdminAccessAssignment) SetAccessDetails(value DelegatedAdminAccessDetailsable)() {
    if m != nil {
        m.accessDetails = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *DelegatedAdminAccessAssignment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *DelegatedAdminAccessAssignment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetStatus sets the status property value. 
func (m *DelegatedAdminAccessAssignment) SetStatus(value *DelegatedAdminAccessAssignmentStatus)() {
    if m != nil {
        m.status = value
    }
}
