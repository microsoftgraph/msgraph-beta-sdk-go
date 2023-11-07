package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegeManagementElevationRequest these are elevation approval requests for EPM support arbitrated scenario initiated by IW user that admins can take action on.
type PrivilegeManagementElevationRequest struct {
    Entity
}
// NewPrivilegeManagementElevationRequest instantiates a new privilegeManagementElevationRequest and sets the default values.
func NewPrivilegeManagementElevationRequest()(*PrivilegeManagementElevationRequest) {
    m := &PrivilegeManagementElevationRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegeManagementElevationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegeManagementElevationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegeManagementElevationRequest(), nil
}
// GetApplicationDetail gets the applicationDetail property value. Details of the application which is being requested to elevate, allowing the admin to understand the identity of the application. It includes file info such as FilePath, FileHash, FilePublisher, and etc. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) GetApplicationDetail()(ApplicationDetailable) {
    val, err := m.GetBackingStore().Get("applicationDetail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ApplicationDetailable)
    }
    return nil
}
// GetDeviceName gets the deviceName property value. The device name used to initiate the elevation request. For example: 'cotonso-laptop'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) GetDeviceName()(*string) {
    val, err := m.GetBackingStore().Get("deviceName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegeManagementElevationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApplicationDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationDetail(val.(ApplicationDetailable))
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["requestCreatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestCreatedDateTime(val)
        }
        return nil
    }
    res["requestedByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedByUserId(val)
        }
        return nil
    }
    res["requestedByUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedByUserPrincipalName(val)
        }
        return nil
    }
    res["requestedOnDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestedOnDeviceId(val)
        }
        return nil
    }
    res["requestExpiryDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestExpiryDateTime(val)
        }
        return nil
    }
    res["requestJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestJustification(val)
        }
        return nil
    }
    res["requestLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestLastModifiedDateTime(val)
        }
        return nil
    }
    res["reviewCompletedByUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewCompletedByUserId(val)
        }
        return nil
    }
    res["reviewCompletedByUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewCompletedByUserPrincipalName(val)
        }
        return nil
    }
    res["reviewCompletedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewCompletedDateTime(val)
        }
        return nil
    }
    res["reviewerJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewerJustification(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseElevationRequestState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ElevationRequestState))
        }
        return nil
    }
    return res
}
// GetRequestCreatedDateTime gets the requestCreatedDateTime property value. The date and time when the elevation request was submitted/created. The value cannot be modified and is automatically populated when the elevation request is submitted/created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) GetRequestCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestCreatedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRequestedByUserId gets the requestedByUserId property value. The Azure Active Directory (AAD) identifier of the end user who is requesting this elevation. For example: 'F1A57311-B9EB-45B7-9415-8555E68EDC9E'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) GetRequestedByUserId()(*string) {
    val, err := m.GetBackingStore().Get("requestedByUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestedByUserPrincipalName gets the requestedByUserPrincipalName property value. The User Principal Name (UPN) of the end user who requested this elevation. For example: 'user1@contoso.com'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) GetRequestedByUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("requestedByUserPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestedOnDeviceId gets the requestedOnDeviceId property value. The Intune Device Identifier of the managed device used to initiate the elevation request. For example: '90F5F6E8-CA09-4811-97F6-4D0DD532D916'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) GetRequestedOnDeviceId()(*string) {
    val, err := m.GetBackingStore().Get("requestedOnDeviceId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestExpiryDateTime gets the requestExpiryDateTime property value. Expiration set for the request when it was created, regardless of approved or denied status. For example: '2023-08-03T14:24:22Z'. Returned by default. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) GetRequestExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestExpiryDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRequestJustification gets the requestJustification property value. Justification provided by the end user for the elevation request. For example :'Need to elevate to install microsoft word'. Read-only.
func (m *PrivilegeManagementElevationRequest) GetRequestJustification()(*string) {
    val, err := m.GetBackingStore().Get("requestJustification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRequestLastModifiedDateTime gets the requestLastModifiedDateTime property value. The date and time when the elevation request was either submitted/created or approved/denied. The value cannot be modified and is automatically populated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) GetRequestLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("requestLastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetReviewCompletedByUserId gets the reviewCompletedByUserId property value. This is the Azure Active Directory (AAD) user id of the administrator who approved or denied the request. For example: 'F1A57311-B9EB-45B7-9415-8555E68EDC9E'. This field would be String.Empty before the request is either approved or denied. Read-only.
func (m *PrivilegeManagementElevationRequest) GetReviewCompletedByUserId()(*string) {
    val, err := m.GetBackingStore().Get("reviewCompletedByUserId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReviewCompletedByUserPrincipalName gets the reviewCompletedByUserPrincipalName property value. This is the User Principal Name (UPN) of the administrator who approved or denied the request. For example: 'admin@contoso.com'. This field would be String.Empty before the request is either approved or denied. Read-only.
func (m *PrivilegeManagementElevationRequest) GetReviewCompletedByUserPrincipalName()(*string) {
    val, err := m.GetBackingStore().Get("reviewCompletedByUserPrincipalName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetReviewCompletedDateTime gets the reviewCompletedDateTime property value. The DateTime for which the request was approved or denied. For example, midnight UTC on August 3rd, 2023 would look like this: '2023-08-03T00:00:00Z'. Read-only.
func (m *PrivilegeManagementElevationRequest) GetReviewCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("reviewCompletedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetReviewerJustification gets the reviewerJustification property value. An optional justification provided by approver at approval or denied time. This field will be String.Empty if approver decides to not provide a justification. For example: 'Run this installer today'
func (m *PrivilegeManagementElevationRequest) GetReviewerJustification()(*string) {
    val, err := m.GetBackingStore().Get("reviewerJustification")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. Indicates state of elevation request
func (m *PrivilegeManagementElevationRequest) GetStatus()(*ElevationRequestState) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ElevationRequestState)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrivilegeManagementElevationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("applicationDetail", m.GetApplicationDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestCreatedDateTime", m.GetRequestCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestedByUserId", m.GetRequestedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestedByUserPrincipalName", m.GetRequestedByUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestedOnDeviceId", m.GetRequestedOnDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestExpiryDateTime", m.GetRequestExpiryDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("requestJustification", m.GetRequestJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("requestLastModifiedDateTime", m.GetRequestLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reviewCompletedByUserId", m.GetReviewCompletedByUserId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reviewCompletedByUserPrincipalName", m.GetReviewCompletedByUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewCompletedDateTime", m.GetReviewCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reviewerJustification", m.GetReviewerJustification())
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
// SetApplicationDetail sets the applicationDetail property value. Details of the application which is being requested to elevate, allowing the admin to understand the identity of the application. It includes file info such as FilePath, FileHash, FilePublisher, and etc. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) SetApplicationDetail(value ApplicationDetailable)() {
    err := m.GetBackingStore().Set("applicationDetail", value)
    if err != nil {
        panic(err)
    }
}
// SetDeviceName sets the deviceName property value. The device name used to initiate the elevation request. For example: 'cotonso-laptop'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) SetDeviceName(value *string)() {
    err := m.GetBackingStore().Set("deviceName", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestCreatedDateTime sets the requestCreatedDateTime property value. The date and time when the elevation request was submitted/created. The value cannot be modified and is automatically populated when the elevation request is submitted/created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) SetRequestCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestCreatedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedByUserId sets the requestedByUserId property value. The Azure Active Directory (AAD) identifier of the end user who is requesting this elevation. For example: 'F1A57311-B9EB-45B7-9415-8555E68EDC9E'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) SetRequestedByUserId(value *string)() {
    err := m.GetBackingStore().Set("requestedByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedByUserPrincipalName sets the requestedByUserPrincipalName property value. The User Principal Name (UPN) of the end user who requested this elevation. For example: 'user1@contoso.com'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) SetRequestedByUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("requestedByUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestedOnDeviceId sets the requestedOnDeviceId property value. The Intune Device Identifier of the managed device used to initiate the elevation request. For example: '90F5F6E8-CA09-4811-97F6-4D0DD532D916'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) SetRequestedOnDeviceId(value *string)() {
    err := m.GetBackingStore().Set("requestedOnDeviceId", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestExpiryDateTime sets the requestExpiryDateTime property value. Expiration set for the request when it was created, regardless of approved or denied status. For example: '2023-08-03T14:24:22Z'. Returned by default. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) SetRequestExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestExpiryDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestJustification sets the requestJustification property value. Justification provided by the end user for the elevation request. For example :'Need to elevate to install microsoft word'. Read-only.
func (m *PrivilegeManagementElevationRequest) SetRequestJustification(value *string)() {
    err := m.GetBackingStore().Set("requestJustification", value)
    if err != nil {
        panic(err)
    }
}
// SetRequestLastModifiedDateTime sets the requestLastModifiedDateTime property value. The date and time when the elevation request was either submitted/created or approved/denied. The value cannot be modified and is automatically populated. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only.
func (m *PrivilegeManagementElevationRequest) SetRequestLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("requestLastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewCompletedByUserId sets the reviewCompletedByUserId property value. This is the Azure Active Directory (AAD) user id of the administrator who approved or denied the request. For example: 'F1A57311-B9EB-45B7-9415-8555E68EDC9E'. This field would be String.Empty before the request is either approved or denied. Read-only.
func (m *PrivilegeManagementElevationRequest) SetReviewCompletedByUserId(value *string)() {
    err := m.GetBackingStore().Set("reviewCompletedByUserId", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewCompletedByUserPrincipalName sets the reviewCompletedByUserPrincipalName property value. This is the User Principal Name (UPN) of the administrator who approved or denied the request. For example: 'admin@contoso.com'. This field would be String.Empty before the request is either approved or denied. Read-only.
func (m *PrivilegeManagementElevationRequest) SetReviewCompletedByUserPrincipalName(value *string)() {
    err := m.GetBackingStore().Set("reviewCompletedByUserPrincipalName", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewCompletedDateTime sets the reviewCompletedDateTime property value. The DateTime for which the request was approved or denied. For example, midnight UTC on August 3rd, 2023 would look like this: '2023-08-03T00:00:00Z'. Read-only.
func (m *PrivilegeManagementElevationRequest) SetReviewCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("reviewCompletedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetReviewerJustification sets the reviewerJustification property value. An optional justification provided by approver at approval or denied time. This field will be String.Empty if approver decides to not provide a justification. For example: 'Run this installer today'
func (m *PrivilegeManagementElevationRequest) SetReviewerJustification(value *string)() {
    err := m.GetBackingStore().Set("reviewerJustification", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Indicates state of elevation request
func (m *PrivilegeManagementElevationRequest) SetStatus(value *ElevationRequestState)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// PrivilegeManagementElevationRequestable 
type PrivilegeManagementElevationRequestable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationDetail()(ApplicationDetailable)
    GetDeviceName()(*string)
    GetRequestCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestedByUserId()(*string)
    GetRequestedByUserPrincipalName()(*string)
    GetRequestedOnDeviceId()(*string)
    GetRequestExpiryDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestJustification()(*string)
    GetRequestLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewCompletedByUserId()(*string)
    GetReviewCompletedByUserPrincipalName()(*string)
    GetReviewCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewerJustification()(*string)
    GetStatus()(*ElevationRequestState)
    SetApplicationDetail(value ApplicationDetailable)()
    SetDeviceName(value *string)()
    SetRequestCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestedByUserId(value *string)()
    SetRequestedByUserPrincipalName(value *string)()
    SetRequestedOnDeviceId(value *string)()
    SetRequestExpiryDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestJustification(value *string)()
    SetRequestLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewCompletedByUserId(value *string)()
    SetReviewCompletedByUserPrincipalName(value *string)()
    SetReviewCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewerJustification(value *string)()
    SetStatus(value *ElevationRequestState)()
}
