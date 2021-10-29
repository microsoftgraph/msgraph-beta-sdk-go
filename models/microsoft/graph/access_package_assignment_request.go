package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AccessPackageAssignmentRequest struct {
    Entity
    // The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
    accessPackage *AccessPackage;
    // For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
    accessPackageAssignment *AccessPackageAssignment;
    // Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
    answers []AccessPackageAnswer;
    // The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    completedDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // True if the request is not to be processed for assignment.
    isValidationOnly *bool;
    // The requestor's supplied justification.
    justification *string;
    // The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
    requestor *AccessPackageSubject;
    // One of PendingApproval, Canceled,  Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted or Scheduled. Read-only.
    requestState *string;
    // More information on the request processing status. Read-only.
    requestStatus *string;
    // One of UserAdd, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd or UserRemove. Read-only.
    requestType *string;
    // The range of dates that access is to be assigned to the requestor. Read-only.
    schedule *RequestSchedule;
}
// Instantiates a new accessPackageAssignmentRequest and sets the default values.
func NewAccessPackageAssignmentRequest()(*AccessPackageAssignmentRequest) {
    m := &AccessPackageAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAccessPackage()(*AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
// Gets the accessPackageAssignment property value. For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAccessPackageAssignment()(*AccessPackageAssignment) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignment
    }
}
// Gets the answers property value. Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
func (m *AccessPackageAssignmentRequest) GetAnswers()([]AccessPackageAnswer) {
    if m == nil {
        return nil
    } else {
        return m.answers
    }
}
// Gets the completedDate property value. The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) GetCompletedDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completedDate
    }
}
// Gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the expirationDateTime property value. 
func (m *AccessPackageAssignmentRequest) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// Gets the isValidationOnly property value. True if the request is not to be processed for assignment.
func (m *AccessPackageAssignmentRequest) GetIsValidationOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidationOnly
    }
}
// Gets the justification property value. The requestor's supplied justification.
func (m *AccessPackageAssignmentRequest) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// Gets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) GetRequestor()(*AccessPackageSubject) {
    if m == nil {
        return nil
    } else {
        return m.requestor
    }
}
// Gets the requestState property value. One of PendingApproval, Canceled,  Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted or Scheduled. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestState
    }
}
// Gets the requestStatus property value. More information on the request processing status. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestStatus
    }
}
// Gets the requestType property value. One of UserAdd, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd or UserRemove. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestType
    }
}
// Gets the schedule property value. The range of dates that access is to be assigned to the requestor. Read-only.
func (m *AccessPackageAssignmentRequest) GetSchedule()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// The deserialization information for the current model
func (m *AccessPackageAssignmentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        m.SetAccessPackage(val.(*AccessPackage))
        return nil
    }
    res["accessPackageAssignment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignment() })
        if err != nil {
            return err
        }
        m.SetAccessPackageAssignment(val.(*AccessPackageAssignment))
        return nil
    }
    res["answers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAnswer() })
        if err != nil {
            return err
        }
        res := make([]AccessPackageAnswer, len(val))
        for i, v := range val {
            res[i] = *(v.(*AccessPackageAnswer))
        }
        m.SetAnswers(res)
        return nil
    }
    res["completedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCompletedDate(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["expirationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetExpirationDateTime(val)
        return nil
    }
    res["isValidationOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsValidationOnly(val)
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJustification(val)
        return nil
    }
    res["requestor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageSubject() })
        if err != nil {
            return err
        }
        m.SetRequestor(val.(*AccessPackageSubject))
        return nil
    }
    res["requestState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestState(val)
        return nil
    }
    res["requestStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestStatus(val)
        return nil
    }
    res["requestType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRequestType(val)
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequestSchedule() })
        if err != nil {
            return err
        }
        m.SetSchedule(val.(*RequestSchedule))
        return nil
    }
    return res
}
func (m *AccessPackageAssignmentRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AccessPackageAssignmentRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackage", m.GetAccessPackage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageAssignment", m.GetAccessPackageAssignment())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAnswers()))
        for i, v := range m.GetAnswers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("answers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("completedDate", m.GetCompletedDate())
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
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isValidationOnly", m.GetIsValidationOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestor", m.GetRequestor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestState", m.GetRequestState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestStatus", m.GetRequestStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestType", m.GetRequestType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("schedule", m.GetSchedule())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
// Parameters:
//  - value : Value to set for the accessPackage property.
func (m *AccessPackageAssignmentRequest) SetAccessPackage(value *AccessPackage)() {
    m.accessPackage = value
}
// Sets the accessPackageAssignment property value. For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand.
// Parameters:
//  - value : Value to set for the accessPackageAssignment property.
func (m *AccessPackageAssignmentRequest) SetAccessPackageAssignment(value *AccessPackageAssignment)() {
    m.accessPackageAssignment = value
}
// Sets the answers property value. Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
// Parameters:
//  - value : Value to set for the answers property.
func (m *AccessPackageAssignmentRequest) SetAnswers(value []AccessPackageAnswer)() {
    m.answers = value
}
// Sets the completedDate property value. The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the completedDate property.
func (m *AccessPackageAssignmentRequest) SetCompletedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDate = value
}
// Sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *AccessPackageAssignmentRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the expirationDateTime property value. 
// Parameters:
//  - value : Value to set for the expirationDateTime property.
func (m *AccessPackageAssignmentRequest) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// Sets the isValidationOnly property value. True if the request is not to be processed for assignment.
// Parameters:
//  - value : Value to set for the isValidationOnly property.
func (m *AccessPackageAssignmentRequest) SetIsValidationOnly(value *bool)() {
    m.isValidationOnly = value
}
// Sets the justification property value. The requestor's supplied justification.
// Parameters:
//  - value : Value to set for the justification property.
func (m *AccessPackageAssignmentRequest) SetJustification(value *string)() {
    m.justification = value
}
// Sets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
// Parameters:
//  - value : Value to set for the requestor property.
func (m *AccessPackageAssignmentRequest) SetRequestor(value *AccessPackageSubject)() {
    m.requestor = value
}
// Sets the requestState property value. One of PendingApproval, Canceled,  Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted or Scheduled. Read-only.
// Parameters:
//  - value : Value to set for the requestState property.
func (m *AccessPackageAssignmentRequest) SetRequestState(value *string)() {
    m.requestState = value
}
// Sets the requestStatus property value. More information on the request processing status. Read-only.
// Parameters:
//  - value : Value to set for the requestStatus property.
func (m *AccessPackageAssignmentRequest) SetRequestStatus(value *string)() {
    m.requestStatus = value
}
// Sets the requestType property value. One of UserAdd, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd or UserRemove. Read-only.
// Parameters:
//  - value : Value to set for the requestType property.
func (m *AccessPackageAssignmentRequest) SetRequestType(value *string)() {
    m.requestType = value
}
// Sets the schedule property value. The range of dates that access is to be assigned to the requestor. Read-only.
// Parameters:
//  - value : Value to set for the schedule property.
func (m *AccessPackageAssignmentRequest) SetSchedule(value *RequestSchedule)() {
    m.schedule = value
}
