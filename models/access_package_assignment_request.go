package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequest 
type AccessPackageAssignmentRequest struct {
    Entity
}
// NewAccessPackageAssignmentRequest instantiates a new accessPackageAssignmentRequest and sets the default values.
func NewAccessPackageAssignmentRequest()(*AccessPackageAssignmentRequest) {
    m := &AccessPackageAssignmentRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageAssignmentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentRequest(), nil
}
// GetAccessPackage gets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAccessPackage()(AccessPackageable) {
    val, err := m.GetBackingStore().Get("accessPackage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageable)
    }
    return nil
}
// GetAccessPackageAssignment gets the accessPackageAssignment property value. For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
func (m *AccessPackageAssignmentRequest) GetAccessPackageAssignment()(AccessPackageAssignmentable) {
    val, err := m.GetBackingStore().Get("accessPackageAssignment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageAssignmentable)
    }
    return nil
}
// GetAnswers gets the answers property value. Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
func (m *AccessPackageAssignmentRequest) GetAnswers()([]AccessPackageAnswerable) {
    val, err := m.GetBackingStore().Get("answers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageAnswerable)
    }
    return nil
}
// GetCompletedDate gets the completedDate property value. The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) GetCompletedDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("completedDate")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetCustomExtensionCalloutInstances gets the customExtensionCalloutInstances property value. Information about all the custom extension calls that were made during the access package assignment request workflow.
func (m *AccessPackageAssignmentRequest) GetCustomExtensionCalloutInstances()([]CustomExtensionCalloutInstanceable) {
    val, err := m.GetBackingStore().Get("customExtensionCalloutInstances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomExtensionCalloutInstanceable)
    }
    return nil
}
// GetCustomExtensionHandlerInstances gets the customExtensionHandlerInstances property value. A collection of custom workflow extension instances being run on an assignment request. Read-only.
func (m *AccessPackageAssignmentRequest) GetCustomExtensionHandlerInstances()([]CustomExtensionHandlerInstanceable) {
    val, err := m.GetBackingStore().Get("customExtensionHandlerInstances")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]CustomExtensionHandlerInstanceable)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The expirationDateTime property
func (m *AccessPackageAssignmentRequest) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackage(val.(AccessPackageable))
        }
        return nil
    }
    res["accessPackageAssignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageAssignment(val.(AccessPackageAssignmentable))
        }
        return nil
    }
    res["answers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAnswerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAnswerable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAnswerable)
            }
            m.SetAnswers(res)
        }
        return nil
    }
    res["completedDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDate(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["customExtensionCalloutInstances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomExtensionCalloutInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomExtensionCalloutInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(CustomExtensionCalloutInstanceable)
            }
            m.SetCustomExtensionCalloutInstances(res)
        }
        return nil
    }
    res["customExtensionHandlerInstances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomExtensionHandlerInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomExtensionHandlerInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(CustomExtensionHandlerInstanceable)
            }
            m.SetCustomExtensionHandlerInstances(res)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["isValidationOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidationOnly(val)
        }
        return nil
    }
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["requestor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestor(val.(AccessPackageSubjectable))
        }
        return nil
    }
    res["requestState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestState(val)
        }
        return nil
    }
    res["requestStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestStatus(val)
        }
        return nil
    }
    res["requestType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestType(val)
        }
        return nil
    }
    res["schedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedule(val.(RequestScheduleable))
        }
        return nil
    }
    res["verifiedCredentialsData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVerifiedCredentialDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VerifiedCredentialDataable, len(val))
            for i, v := range val {
                res[i] = v.(VerifiedCredentialDataable)
            }
            m.SetVerifiedCredentialsData(res)
        }
        return nil
    }
    return res
}
// GetIsValidationOnly gets the isValidationOnly property value. True if the request is not to be processed for assignment.
func (m *AccessPackageAssignmentRequest) GetIsValidationOnly()(*bool) {
    val, err := m.GetBackingStore().Get("isValidationOnly")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetJustification gets the justification property value. The requestor's supplied justification.
func (m *AccessPackageAssignmentRequest) GetJustification()(*string) {
    val, err := m.GetBackingStore().Get("justification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestor gets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) GetRequestor()(AccessPackageSubjectable) {
    val, err := m.GetBackingStore().Get("requestor")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageSubjectable)
    }
    return nil
}
// GetRequestState gets the requestState property value. One of PendingApproval, Canceled,  Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted or Scheduled. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestState()(*string) {
    val, err := m.GetBackingStore().Get("requestState")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestStatus gets the requestStatus property value. More information on the request processing status. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestStatus()(*string) {
    val, err := m.GetBackingStore().Get("requestStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestType gets the requestType property value. One of UserAdd, UserExtend, UserUpdate, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd, UserUpdate or UserRemove. Read-only.
func (m *AccessPackageAssignmentRequest) GetRequestType()(*string) {
    val, err := m.GetBackingStore().Get("requestType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSchedule gets the schedule property value. The range of dates that access is to be assigned to the requestor. Read-only.
func (m *AccessPackageAssignmentRequest) GetSchedule()(RequestScheduleable) {
    val, err := m.GetBackingStore().Get("schedule")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RequestScheduleable)
    }
    return nil
}
// GetVerifiedCredentialsData gets the verifiedCredentialsData property value. The details of the verifiable credential that was presented by the requestor, such as the issuer and claims. Read-only.
func (m *AccessPackageAssignmentRequest) GetVerifiedCredentialsData()([]VerifiedCredentialDataable) {
    val, err := m.GetBackingStore().Get("verifiedCredentialsData")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VerifiedCredentialDataable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAnswers()))
        for i, v := range m.GetAnswers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetCustomExtensionCalloutInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomExtensionCalloutInstances()))
        for i, v := range m.GetCustomExtensionCalloutInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customExtensionCalloutInstances", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomExtensionHandlerInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomExtensionHandlerInstances()))
        for i, v := range m.GetCustomExtensionHandlerInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetVerifiedCredentialsData() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVerifiedCredentialsData()))
        for i, v := range m.GetVerifiedCredentialsData() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("verifiedCredentialsData", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackage sets the accessPackage property value. The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) SetAccessPackage(value AccessPackageable)() {
    err := m.GetBackingStore().Set("accessPackage", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageAssignment sets the accessPackageAssignment property value. For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.  Supports $expand.
func (m *AccessPackageAssignmentRequest) SetAccessPackageAssignment(value AccessPackageAssignmentable)() {
    err := m.GetBackingStore().Set("accessPackageAssignment", value)
    if err != nil {
        panic(err)
    }
}
// SetAnswers sets the answers property value. Answers provided by the requestor to accessPackageQuestions asked of them at the time of request.
func (m *AccessPackageAssignmentRequest) SetAnswers(value []AccessPackageAnswerable)() {
    err := m.GetBackingStore().Set("answers", value)
    if err != nil {
        panic(err)
    }
}
// SetCompletedDate sets the completedDate property value. The date of the end of processing, either successful or failure, of a request. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) SetCompletedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("completedDate", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *AccessPackageAssignmentRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomExtensionCalloutInstances sets the customExtensionCalloutInstances property value. Information about all the custom extension calls that were made during the access package assignment request workflow.
func (m *AccessPackageAssignmentRequest) SetCustomExtensionCalloutInstances(value []CustomExtensionCalloutInstanceable)() {
    err := m.GetBackingStore().Set("customExtensionCalloutInstances", value)
    if err != nil {
        panic(err)
    }
}
// SetCustomExtensionHandlerInstances sets the customExtensionHandlerInstances property value. A collection of custom workflow extension instances being run on an assignment request. Read-only.
func (m *AccessPackageAssignmentRequest) SetCustomExtensionHandlerInstances(value []CustomExtensionHandlerInstanceable)() {
    err := m.GetBackingStore().Set("customExtensionHandlerInstances", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The expirationDateTime property
func (m *AccessPackageAssignmentRequest) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetIsValidationOnly sets the isValidationOnly property value. True if the request is not to be processed for assignment.
func (m *AccessPackageAssignmentRequest) SetIsValidationOnly(value *bool)() {
    err := m.GetBackingStore().Set("isValidationOnly", value)
    if err != nil {
        panic(err)
    }
}
// SetJustification sets the justification property value. The requestor's supplied justification.
func (m *AccessPackageAssignmentRequest) SetJustification(value *string)() {
    err := m.GetBackingStore().Set("justification", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestor sets the requestor property value. The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (m *AccessPackageAssignmentRequest) SetRequestor(value AccessPackageSubjectable)() {
    err := m.GetBackingStore().Set("requestor", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestState sets the requestState property value. One of PendingApproval, Canceled,  Denied, Delivering, Delivered, PartiallyDelivered, DeliveryFailed, Submitted or Scheduled. Read-only.
func (m *AccessPackageAssignmentRequest) SetRequestState(value *string)() {
    err := m.GetBackingStore().Set("requestState", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestStatus sets the requestStatus property value. More information on the request processing status. Read-only.
func (m *AccessPackageAssignmentRequest) SetRequestStatus(value *string)() {
    err := m.GetBackingStore().Set("requestStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestType sets the requestType property value. One of UserAdd, UserExtend, UserUpdate, UserRemove, AdminAdd, AdminRemove or SystemRemove. A request from the user themselves would have requestType of UserAdd, UserUpdate or UserRemove. Read-only.
func (m *AccessPackageAssignmentRequest) SetRequestType(value *string)() {
    err := m.GetBackingStore().Set("requestType", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedule sets the schedule property value. The range of dates that access is to be assigned to the requestor. Read-only.
func (m *AccessPackageAssignmentRequest) SetSchedule(value RequestScheduleable)() {
    err := m.GetBackingStore().Set("schedule", value)
    if err != nil {
        panic(err)
    }
}
// SetVerifiedCredentialsData sets the verifiedCredentialsData property value. The details of the verifiable credential that was presented by the requestor, such as the issuer and claims. Read-only.
func (m *AccessPackageAssignmentRequest) SetVerifiedCredentialsData(value []VerifiedCredentialDataable)() {
    err := m.GetBackingStore().Set("verifiedCredentialsData", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageAssignmentRequestable 
type AccessPackageAssignmentRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAccessPackageAssignment()(AccessPackageAssignmentable)
    GetAnswers()([]AccessPackageAnswerable)
    GetCompletedDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomExtensionCalloutInstances()([]CustomExtensionCalloutInstanceable)
    GetCustomExtensionHandlerInstances()([]CustomExtensionHandlerInstanceable)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsValidationOnly()(*bool)
    GetJustification()(*string)
    GetRequestor()(AccessPackageSubjectable)
    GetRequestState()(*string)
    GetRequestStatus()(*string)
    GetRequestType()(*string)
    GetSchedule()(RequestScheduleable)
    GetVerifiedCredentialsData()([]VerifiedCredentialDataable)
    SetAccessPackage(value AccessPackageable)()
    SetAccessPackageAssignment(value AccessPackageAssignmentable)()
    SetAnswers(value []AccessPackageAnswerable)()
    SetCompletedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomExtensionCalloutInstances(value []CustomExtensionCalloutInstanceable)()
    SetCustomExtensionHandlerInstances(value []CustomExtensionHandlerInstanceable)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsValidationOnly(value *bool)()
    SetJustification(value *string)()
    SetRequestor(value AccessPackageSubjectable)()
    SetRequestState(value *string)()
    SetRequestStatus(value *string)()
    SetRequestType(value *string)()
    SetSchedule(value RequestScheduleable)()
    SetVerifiedCredentialsData(value []VerifiedCredentialDataable)()
}
