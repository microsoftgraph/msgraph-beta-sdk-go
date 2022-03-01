package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DelegatedAdminRelationship 
type DelegatedAdminRelationship struct {
    Entity
    // 
    accessAssignments []DelegatedAdminAccessAssignment;
    // 
    accessDetails *DelegatedAdminAccessDetails;
    // 
    activatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    customer *DelegatedAdminRelationshipCustomerParticipant;
    // 
    displayName *string;
    // 
    duration *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // 
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    operations []DelegatedAdminRelationshipOperation;
    // 
    partner *DelegatedAdminRelationshipParticipant;
    // 
    requests []DelegatedAdminRelationshipRequest;
    // 
    status *DelegatedAdminRelationshipStatus;
}
// NewDelegatedAdminRelationship instantiates a new delegatedAdminRelationship and sets the default values.
func NewDelegatedAdminRelationship()(*DelegatedAdminRelationship) {
    m := &DelegatedAdminRelationship{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessAssignments gets the accessAssignments property value. 
func (m *DelegatedAdminRelationship) GetAccessAssignments()([]DelegatedAdminAccessAssignment) {
    if m == nil {
        return nil
    } else {
        return m.accessAssignments
    }
}
// GetAccessDetails gets the accessDetails property value. 
func (m *DelegatedAdminRelationship) GetAccessDetails()(*DelegatedAdminAccessDetails) {
    if m == nil {
        return nil
    } else {
        return m.accessDetails
    }
}
// GetActivatedDateTime gets the activatedDateTime property value. 
func (m *DelegatedAdminRelationship) GetActivatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activatedDateTime
    }
}
// GetCreatedDateTime gets the createdDateTime property value. 
func (m *DelegatedAdminRelationship) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCustomer gets the customer property value. 
func (m *DelegatedAdminRelationship) GetCustomer()(*DelegatedAdminRelationshipCustomerParticipant) {
    if m == nil {
        return nil
    } else {
        return m.customer
    }
}
// GetDisplayName gets the displayName property value. 
func (m *DelegatedAdminRelationship) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetDuration gets the duration property value. 
func (m *DelegatedAdminRelationship) GetDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.duration
    }
}
// GetEndDateTime gets the endDateTime property value. 
func (m *DelegatedAdminRelationship) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *DelegatedAdminRelationship) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetOperations gets the operations property value. 
func (m *DelegatedAdminRelationship) GetOperations()([]DelegatedAdminRelationshipOperation) {
    if m == nil {
        return nil
    } else {
        return m.operations
    }
}
// GetPartner gets the partner property value. 
func (m *DelegatedAdminRelationship) GetPartner()(*DelegatedAdminRelationshipParticipant) {
    if m == nil {
        return nil
    } else {
        return m.partner
    }
}
// GetRequests gets the requests property value. 
func (m *DelegatedAdminRelationship) GetRequests()([]DelegatedAdminRelationshipRequest) {
    if m == nil {
        return nil
    } else {
        return m.requests
    }
}
// GetStatus gets the status property value. 
func (m *DelegatedAdminRelationship) GetStatus()(*DelegatedAdminRelationshipStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminRelationship) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDelegatedAdminAccessAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminAccessAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*DelegatedAdminAccessAssignment))
            }
            m.SetAccessAssignments(res)
        }
        return nil
    }
    res["accessDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDelegatedAdminAccessDetails() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessDetails(val.(*DelegatedAdminAccessDetails))
        }
        return nil
    }
    res["activatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedDateTime(val)
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
    res["customer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDelegatedAdminRelationshipCustomerParticipant() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomer(val.(*DelegatedAdminRelationshipCustomerParticipant))
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["duration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
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
    res["operations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDelegatedAdminRelationshipOperation() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminRelationshipOperation, len(val))
            for i, v := range val {
                res[i] = *(v.(*DelegatedAdminRelationshipOperation))
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["partner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDelegatedAdminRelationshipParticipant() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartner(val.(*DelegatedAdminRelationshipParticipant))
        }
        return nil
    }
    res["requests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDelegatedAdminRelationshipRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminRelationshipRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*DelegatedAdminRelationshipRequest))
            }
            m.SetRequests(res)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminRelationshipStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DelegatedAdminRelationshipStatus))
        }
        return nil
    }
    return res
}
func (m *DelegatedAdminRelationship) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DelegatedAdminRelationship) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessAssignments()))
        for i, v := range m.GetAccessAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessAssignments", cast)
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
        err = writer.WriteTimeValue("activatedDateTime", m.GetActivatedDateTime())
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
        err = writer.WriteObjectValue("customer", m.GetCustomer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
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
    if m.GetOperations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("partner", m.GetPartner())
        if err != nil {
            return err
        }
    }
    if m.GetRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRequests()))
        for i, v := range m.GetRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("requests", cast)
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
// SetAccessAssignments sets the accessAssignments property value. 
func (m *DelegatedAdminRelationship) SetAccessAssignments(value []DelegatedAdminAccessAssignment)() {
    if m != nil {
        m.accessAssignments = value
    }
}
// SetAccessDetails sets the accessDetails property value. 
func (m *DelegatedAdminRelationship) SetAccessDetails(value *DelegatedAdminAccessDetails)() {
    if m != nil {
        m.accessDetails = value
    }
}
// SetActivatedDateTime sets the activatedDateTime property value. 
func (m *DelegatedAdminRelationship) SetActivatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.activatedDateTime = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. 
func (m *DelegatedAdminRelationship) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCustomer sets the customer property value. 
func (m *DelegatedAdminRelationship) SetCustomer(value *DelegatedAdminRelationshipCustomerParticipant)() {
    if m != nil {
        m.customer = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *DelegatedAdminRelationship) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetDuration sets the duration property value. 
func (m *DelegatedAdminRelationship) SetDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.duration = value
    }
}
// SetEndDateTime sets the endDateTime property value. 
func (m *DelegatedAdminRelationship) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.endDateTime = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *DelegatedAdminRelationship) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetOperations sets the operations property value. 
func (m *DelegatedAdminRelationship) SetOperations(value []DelegatedAdminRelationshipOperation)() {
    if m != nil {
        m.operations = value
    }
}
// SetPartner sets the partner property value. 
func (m *DelegatedAdminRelationship) SetPartner(value *DelegatedAdminRelationshipParticipant)() {
    if m != nil {
        m.partner = value
    }
}
// SetRequests sets the requests property value. 
func (m *DelegatedAdminRelationship) SetRequests(value []DelegatedAdminRelationshipRequest)() {
    if m != nil {
        m.requests = value
    }
}
// SetStatus sets the status property value. 
func (m *DelegatedAdminRelationship) SetStatus(value *DelegatedAdminRelationshipStatus)() {
    if m != nil {
        m.status = value
    }
}
