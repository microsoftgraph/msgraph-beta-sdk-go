package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DelegatedAdminRelationshipOperation provides operations to manage the tenantRelationship singleton.
type DelegatedAdminRelationshipOperation struct {
    Entity
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    data *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    operationType *DelegatedAdminRelationshipOperationType;
    // 
    status *DelegatedAdminRelationshipOperationStatus;
}
// NewDelegatedAdminRelationshipOperation instantiates a new delegatedAdminRelationshipOperation and sets the default values.
func NewDelegatedAdminRelationshipOperation()(*DelegatedAdminRelationshipOperation) {
    m := &DelegatedAdminRelationshipOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDelegatedAdminRelationshipOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminRelationshipOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDelegatedAdminRelationshipOperation(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *DelegatedAdminRelationshipOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetData gets the data property value. 
func (m *DelegatedAdminRelationshipOperation) GetData()(*string) {
    if m == nil {
        return nil
    } else {
        return m.data
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminRelationshipOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["data"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val)
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
    res["operationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminRelationshipOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*DelegatedAdminRelationshipOperationType))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminRelationshipOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DelegatedAdminRelationshipOperationStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *DelegatedAdminRelationshipOperation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetOperationType gets the operationType property value. 
func (m *DelegatedAdminRelationshipOperation) GetOperationType()(*DelegatedAdminRelationshipOperationType) {
    if m == nil {
        return nil
    } else {
        return m.operationType
    }
}
// GetStatus gets the status property value. 
func (m *DelegatedAdminRelationshipOperation) GetStatus()(*DelegatedAdminRelationshipOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *DelegatedAdminRelationshipOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DelegatedAdminRelationshipOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("data", m.GetData())
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
    if m.GetOperationType() != nil {
        cast := (*m.GetOperationType()).String()
        err = writer.WriteStringValue("operationType", &cast)
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
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *DelegatedAdminRelationshipOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetData sets the data property value. 
func (m *DelegatedAdminRelationshipOperation) SetData(value *string)() {
    if m != nil {
        m.data = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *DelegatedAdminRelationshipOperation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetOperationType sets the operationType property value. 
func (m *DelegatedAdminRelationshipOperation) SetOperationType(value *DelegatedAdminRelationshipOperationType)() {
    if m != nil {
        m.operationType = value
    }
}
// SetStatus sets the status property value. 
func (m *DelegatedAdminRelationshipOperation) SetStatus(value *DelegatedAdminRelationshipOperationStatus)() {
    if m != nil {
        m.status = value
    }
}
