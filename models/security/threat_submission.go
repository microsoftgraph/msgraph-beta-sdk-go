package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

type ThreatSubmission struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
}
// NewThreatSubmission instantiates a new ThreatSubmission and sets the default values.
func NewThreatSubmission()(*ThreatSubmission) {
    m := &ThreatSubmission{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateThreatSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateThreatSubmissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.security.emailContentThreatSubmission":
                        return NewEmailContentThreatSubmission(), nil
                    case "#microsoft.graph.security.emailThreatSubmission":
                        return NewEmailThreatSubmission(), nil
                    case "#microsoft.graph.security.emailUrlThreatSubmission":
                        return NewEmailUrlThreatSubmission(), nil
                    case "#microsoft.graph.security.fileContentThreatSubmission":
                        return NewFileContentThreatSubmission(), nil
                    case "#microsoft.graph.security.fileThreatSubmission":
                        return NewFileThreatSubmission(), nil
                    case "#microsoft.graph.security.fileUrlThreatSubmission":
                        return NewFileUrlThreatSubmission(), nil
                    case "#microsoft.graph.security.urlThreatSubmission":
                        return NewUrlThreatSubmission(), nil
                }
            }
        }
    }
    return NewThreatSubmission(), nil
}
// GetAdminReview gets the adminReview property value. Specifies the admin review property that constitutes of who reviewed the user submission, when and what was it identified as.
// returns a SubmissionAdminReviewable when successful
func (m *ThreatSubmission) GetAdminReview()(SubmissionAdminReviewable) {
    val, err := m.GetBackingStore().Get("adminReview")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SubmissionAdminReviewable)
    }
    return nil
}
// GetCategory gets the category property value. The category property
// returns a *SubmissionCategory when successful
func (m *ThreatSubmission) GetCategory()(*SubmissionCategory) {
    val, err := m.GetBackingStore().Get("category")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubmissionCategory)
    }
    return nil
}
// GetClientSource gets the clientSource property value. Specifies the source of the submission. The possible values are: microsoft, other, and unkownFutureValue.
// returns a *SubmissionClientSource when successful
func (m *ThreatSubmission) GetClientSource()(*SubmissionClientSource) {
    val, err := m.GetBackingStore().Get("clientSource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubmissionClientSource)
    }
    return nil
}
// GetContentType gets the contentType property value. Specifies the type of content being submitted. The possible values are: email, url, file, app, and unkownFutureValue.
// returns a *SubmissionContentType when successful
func (m *ThreatSubmission) GetContentType()(*SubmissionContentType) {
    val, err := m.GetBackingStore().Get("contentType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubmissionContentType)
    }
    return nil
}
// GetCreatedBy gets the createdBy property value. Specifies who submitted the email as a threat. Supports $filter = createdBy/email eq 'value'.
// returns a SubmissionUserIdentityable when successful
func (m *ThreatSubmission) GetCreatedBy()(SubmissionUserIdentityable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SubmissionUserIdentityable)
    }
    return nil
}
// GetCreatedDateTime gets the createdDateTime property value. Specifies when the threat submission was created. Supports $filter = createdDateTime ge 2022-01-01T00:00:00Z and createdDateTime lt 2022-01-02T00:00:00Z.
// returns a *Time when successful
func (m *ThreatSubmission) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ThreatSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adminReview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubmissionAdminReviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminReview(val.(SubmissionAdminReviewable))
        }
        return nil
    }
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubmissionCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*SubmissionCategory))
        }
        return nil
    }
    res["clientSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubmissionClientSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSource(val.(*SubmissionClientSource))
        }
        return nil
    }
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubmissionContentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val.(*SubmissionContentType))
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubmissionUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(SubmissionUserIdentityable))
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
    res["result"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubmissionResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResult(val.(SubmissionResultable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubmissionSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*SubmissionSource))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLongRunningOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*LongRunningOperationStatus))
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetResult gets the result property value. Specifies the result of the analysis performed by Microsoft.
// returns a SubmissionResultable when successful
func (m *ThreatSubmission) GetResult()(SubmissionResultable) {
    val, err := m.GetBackingStore().Get("result")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SubmissionResultable)
    }
    return nil
}
// GetSource gets the source property value. Specifies the role of the submitter. Supports $filter = source eq 'value'. The possible values are: administrator,  user, and unkownFutureValue.
// returns a *SubmissionSource when successful
func (m *ThreatSubmission) GetSource()(*SubmissionSource) {
    val, err := m.GetBackingStore().Get("source")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*SubmissionSource)
    }
    return nil
}
// GetStatus gets the status property value. Indicates whether the threat submission has been analyzed by Microsoft. Supports $filter = status eq 'value'. The possible values are: notStarted, running, succeeded, failed, skipped, and unkownFutureValue.
// returns a *LongRunningOperationStatus when successful
func (m *ThreatSubmission) GetStatus()(*LongRunningOperationStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*LongRunningOperationStatus)
    }
    return nil
}
// GetTenantId gets the tenantId property value. Indicates the tenant id of the submitter. Not required when created using a POST operation. It's extracted from the token of the post API call.
// returns a *string when successful
func (m *ThreatSubmission) GetTenantId()(*string) {
    val, err := m.GetBackingStore().Get("tenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ThreatSubmission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("adminReview", m.GetAdminReview())
        if err != nil {
            return err
        }
    }
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err = writer.WriteStringValue("category", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetClientSource() != nil {
        cast := (*m.GetClientSource()).String()
        err = writer.WriteStringValue("clientSource", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentType() != nil {
        cast := (*m.GetContentType()).String()
        err = writer.WriteStringValue("contentType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
        err = writer.WriteObjectValue("result", m.GetResult())
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
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
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdminReview sets the adminReview property value. Specifies the admin review property that constitutes of who reviewed the user submission, when and what was it identified as.
func (m *ThreatSubmission) SetAdminReview(value SubmissionAdminReviewable)() {
    err := m.GetBackingStore().Set("adminReview", value)
    if err != nil {
        panic(err)
    }
}
// SetCategory sets the category property value. The category property
func (m *ThreatSubmission) SetCategory(value *SubmissionCategory)() {
    err := m.GetBackingStore().Set("category", value)
    if err != nil {
        panic(err)
    }
}
// SetClientSource sets the clientSource property value. Specifies the source of the submission. The possible values are: microsoft, other, and unkownFutureValue.
func (m *ThreatSubmission) SetClientSource(value *SubmissionClientSource)() {
    err := m.GetBackingStore().Set("clientSource", value)
    if err != nil {
        panic(err)
    }
}
// SetContentType sets the contentType property value. Specifies the type of content being submitted. The possible values are: email, url, file, app, and unkownFutureValue.
func (m *ThreatSubmission) SetContentType(value *SubmissionContentType)() {
    err := m.GetBackingStore().Set("contentType", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedBy sets the createdBy property value. Specifies who submitted the email as a threat. Supports $filter = createdBy/email eq 'value'.
func (m *ThreatSubmission) SetCreatedBy(value SubmissionUserIdentityable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Specifies when the threat submission was created. Supports $filter = createdDateTime ge 2022-01-01T00:00:00Z and createdDateTime lt 2022-01-02T00:00:00Z.
func (m *ThreatSubmission) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetResult sets the result property value. Specifies the result of the analysis performed by Microsoft.
func (m *ThreatSubmission) SetResult(value SubmissionResultable)() {
    err := m.GetBackingStore().Set("result", value)
    if err != nil {
        panic(err)
    }
}
// SetSource sets the source property value. Specifies the role of the submitter. Supports $filter = source eq 'value'. The possible values are: administrator,  user, and unkownFutureValue.
func (m *ThreatSubmission) SetSource(value *SubmissionSource)() {
    err := m.GetBackingStore().Set("source", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Indicates whether the threat submission has been analyzed by Microsoft. Supports $filter = status eq 'value'. The possible values are: notStarted, running, succeeded, failed, skipped, and unkownFutureValue.
func (m *ThreatSubmission) SetStatus(value *LongRunningOperationStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetTenantId sets the tenantId property value. Indicates the tenant id of the submitter. Not required when created using a POST operation. It's extracted from the token of the post API call.
func (m *ThreatSubmission) SetTenantId(value *string)() {
    err := m.GetBackingStore().Set("tenantId", value)
    if err != nil {
        panic(err)
    }
}
type ThreatSubmissionable interface {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdminReview()(SubmissionAdminReviewable)
    GetCategory()(*SubmissionCategory)
    GetClientSource()(*SubmissionClientSource)
    GetContentType()(*SubmissionContentType)
    GetCreatedBy()(SubmissionUserIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetResult()(SubmissionResultable)
    GetSource()(*SubmissionSource)
    GetStatus()(*LongRunningOperationStatus)
    GetTenantId()(*string)
    SetAdminReview(value SubmissionAdminReviewable)()
    SetCategory(value *SubmissionCategory)()
    SetClientSource(value *SubmissionClientSource)()
    SetContentType(value *SubmissionContentType)()
    SetCreatedBy(value SubmissionUserIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetResult(value SubmissionResultable)()
    SetSource(value *SubmissionSource)()
    SetStatus(value *LongRunningOperationStatus)()
    SetTenantId(value *string)()
}
