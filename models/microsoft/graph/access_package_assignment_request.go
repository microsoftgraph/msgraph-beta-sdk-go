package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentRequest 
type AccessPackageAssignmentRequest struct {
    Entity
    // The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
    accessPackage *AccessPackage;
    // For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
    accessPackageAssignment *AccessPackageAssignment;
    // Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
    answers []AccessPackageAnswer;
    // The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    completedDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    customExtensionHandlerInstances []CustomExtensionHandlerInstance;
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
    // The type of the request. The possible values are: notSpecified, userAdd, userUpdate, userRemove, adminAdd, adminUpdate, adminRemove, systemAdd, systemUpdate, systemRemove, onBehalfAdd, unknownFutureValue. A request from the user themselves would have requestType of UserAdd or UserRemove. This property cannot be changed once set.
    requestType *string;
    // The range of dates that access is to be assigned to the requestor. This property cannot be changed once set.
    schedule *RequestSchedule;
}
// NewAccessPackageAssignmentRequest instantiates a new accessPackageAssignmentRequest and sets the default values.
func NewAccessPackageAssignmentRequest()(*AccessPackageAssignmentRequest) {
    m := &AccessPackageAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackage gets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAccessPackage()(*AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackage
    }
}
// GetAccessPackageAssignment gets the accessPackageAssignment property value. For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAccessPackageAssignment()(*AccessPackageAssignment) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignment
    }
}
// GetAnswers gets the answers property value. Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
func (m *AccessPackageAssignmentRequest) GetAnswers()([]AccessPackageAnswer) {
    if m == nil {
        return nil
    } else {
        return m.answers
    }
}
// GetCompletedDate gets the completedDate property value. The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) GetCompletedDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completedDate
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetCustomExtensionHandlerInstances gets the customExtensionHandlerInstances property value. 
func (m *AccessPackageAssignmentRequest) GetCustomExtensionHandlerInstances()([]CustomExtensionHandlerInstance) {
    if m == nil {
        return nil
    } else {
        return m.customExtensionHandlerInstances
    }
}
// GetExpirationDateTime gets the expirationDateTime property value. 
func (m *AccessPackageAssignmentRequest) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.expirationDateTime
    }
}
// GetIsValidationOnly gets the isValidationOnly property value. True if the request is not to be processed for assignment.
func (m *AccessPackageAssignmentRequest) GetIsValidationOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isValidationOnly
    }
}
// GetJustification gets the justification property value. The requestor's supplied justification.
func (m *AccessPackageAssignmentRequest) GetJustification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justification
    }
}
// GetRequestor gets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) GetRequestor()(*AccessPackageSubject) {
    if m == nil {
        return nil
    } else {
        return m.requestor
    }
}
// GetRequestState gets the requestState property value. One of PendingApproval, Canceled,  Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted or Scheduled. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestState
    }
}
// GetRequestStatus gets the requestStatus property value. More information on the request processing status. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestStatus
    }
}
// GetRequestType gets the requestType property value. The type of the request. The possible values are: notSpecified, userAdd, userUpdate, userRemove, adminAdd, adminUpdate, adminRemove, systemAdd, systemUpdate, systemRemove, onBehalfAdd, unknownFutureValue. A request from the user themselves would have requestType of UserAdd or UserRemove. This property cannot be changed once set.
func (m *AccessPackageAssignmentRequest) GetRequestType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.requestType
    }
}
// GetSchedule gets the schedule property value. The range of dates that access is to be assigned to the requestor. This property cannot be changed once set.
func (m *AccessPackageAssignmentRequest) GetSchedule()(*RequestSchedule) {
    if m == nil {
        return nil
    } else {
        return m.schedule
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackage(val.(*AccessPackage))
        }
        return nil
    }
    res["accessPackageAssignment"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageAssignment(val.(*AccessPackageAssignment))
        }
        return nil
    }
    res["answers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAnswer() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAnswer, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageAnswer))
            }
            m.SetAnswers(res)
        }
        return nil
    }
    res["completedDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDate(val)
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
    res["customExtensionHandlerInstances"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCustomExtensionHandlerInstance() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomExtensionHandlerInstance, len(val))
            for i, v := range val {
                res[i] = *(v.(*CustomExtensionHandlerInstance))
            }
            m.SetCustomExtensionHandlerInstances(res)
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
    res["isValidationOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidationOnly(val)
        }
        return nil
    }
    res["justification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["requestor"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageSubject() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestor(val.(*AccessPackageSubject))
        }
        return nil
    }
    res["requestState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestState(val)
        }
        return nil
    }
    res["requestStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestStatus(val)
        }
        return nil
    }
    res["requestType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestType(val)
        }
        return nil
    }
    res["schedule"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRequestSchedule() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(*RequestSchedule))
        }
        return nil
    }
    return res
}
func (m *AccessPackageAssignmentRequest) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetAnswers() != nil {
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
    if m.GetCustomExtensionHandlerInstances() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomExtensionHandlerInstances()))
        for i, v := range m.GetCustomExtensionHandlerInstances() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("customExtensionHandlerInstances", cast)
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
// SetAccessPackage sets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand.
func (m *AccessPackageAssignmentRequest) SetAccessPackage(value *AccessPackage)() {
    if m != nil {
        m.accessPackage = value
    }
}
// SetAccessPackageAssignment sets the accessPackageAssignment property value. For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
func (m *AccessPackageAssignmentRequest) SetAccessPackageAssignment(value *AccessPackageAssignment)() {
    if m != nil {
        m.accessPackageAssignment = value
    }
}
// SetAnswers sets the answers property value. Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
func (m *AccessPackageAssignmentRequest) SetAnswers(value []AccessPackageAnswer)() {
    if m != nil {
        m.answers = value
    }
}
// SetCompletedDate sets the completedDate property value. The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) SetCompletedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.completedDate = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetCustomExtensionHandlerInstances sets the customExtensionHandlerInstances property value. 
func (m *AccessPackageAssignmentRequest) SetCustomExtensionHandlerInstances(value []CustomExtensionHandlerInstance)() {
    if m != nil {
        m.customExtensionHandlerInstances = value
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. 
func (m *AccessPackageAssignmentRequest) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.expirationDateTime = value
    }
}
// SetIsValidationOnly sets the isValidationOnly property value. True if the request is not to be processed for assignment.
func (m *AccessPackageAssignmentRequest) SetIsValidationOnly(value *bool)() {
    if m != nil {
        m.isValidationOnly = value
    }
}
// SetJustification sets the justification property value. The requestor's supplied justification.
func (m *AccessPackageAssignmentRequest) SetJustification(value *string)() {
    if m != nil {
        m.justification = value
    }
}
// SetRequestor sets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) SetRequestor(value *AccessPackageSubject)() {
    if m != nil {
        m.requestor = value
    }
}
// SetRequestState sets the requestState property value. One of PendingApproval, Canceled,  Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted or Scheduled. Read-only.
func (m *AccessPackageAssignmentRequest) SetRequestState(value *string)() {
    if m != nil {
        m.requestState = value
    }
}
// SetRequestStatus sets the requestStatus property value. More information on the request processing status. Read-only.
func (m *AccessPackageAssignmentRequest) SetRequestStatus(value *string)() {
    if m != nil {
        m.requestStatus = value
    }
}
// SetRequestType sets the requestType property value. The type of the request. The possible values are: notSpecified, userAdd, userUpdate, userRemove, adminAdd, adminUpdate, adminRemove, systemAdd, systemUpdate, systemRemove, onBehalfAdd, unknownFutureValue. A request from the user themselves would have requestType of UserAdd or UserRemove. This property cannot be changed once set.
func (m *AccessPackageAssignmentRequest) SetRequestType(value *string)() {
    if m != nil {
        m.requestType = value
    }
}
// SetSchedule sets the schedule property value. The range of dates that access is to be assigned to the requestor. This property cannot be changed once set.
func (m *AccessPackageAssignmentRequest) SetSchedule(value *RequestSchedule)() {
    if m != nil {
        m.schedule = value
    }
}
