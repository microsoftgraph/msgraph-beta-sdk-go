package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ThreatSubmission provides operations to manage the collection of accessReview entities.
type ThreatSubmission struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // Specifies the admin review property which constitutes of who reviewed the user submission, when and what was it identified as.
    adminReview SubmissionAdminReviewable
    // The category property
    category *SubmissionCategory
    // Specifies the source of the submission. The possible values are: microsoft,  other and unkownFutureValue.
    clientSource *SubmissionClientSource
    // Specifies the type of content being submitted. The possible values are: email, url, file, app and unkownFutureValue.
    contentType *SubmissionContentType
    // Specifies who submitted the email as a threat. Supports $filter = createdBy/email eq 'value'.
    createdBy SubmissionUserIdentityable
    // Specifies when the threat submission was created. Supports $filter = createdDateTime ge 2022-01-01T00:00:00Z and createdDateTime lt 2022-01-02T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the result of the analysis performed by Microsoft.
    result SubmissionResultable
    // Specifies the role of the submitter. Supports $filter = source eq 'value'. The possible values are: administrator,  user and unkownFutureValue.
    source *SubmissionSource
    // Indicates whether the threat submission has been analyzed by Microsoft. Supports $filter = status eq 'value'. The possible values are: notStarted, running, succeeded, failed, skipped and unkownFutureValue.
    status *LongRunningOperationStatus
    // Indicates the tenant id of the submitter. Not required when created using a POST operation. It is extracted from the token of the post API call.
    tenantId *string
}
// NewThreatSubmission instantiates a new threatSubmission and sets the default values.
func NewThreatSubmission()(*ThreatSubmission) {
    m := &ThreatSubmission{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    return m
}
// CreateThreatSubmissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
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
// GetAdminReview gets the adminReview property value. Specifies the admin review property which constitutes of who reviewed the user submission, when and what was it identified as.
func (m *ThreatSubmission) GetAdminReview()(SubmissionAdminReviewable) {
    return m.adminReview
}
// GetCategory gets the category property value. The category property
func (m *ThreatSubmission) GetCategory()(*SubmissionCategory) {
    return m.category
}
// GetClientSource gets the clientSource property value. Specifies the source of the submission. The possible values are: microsoft,  other and unkownFutureValue.
func (m *ThreatSubmission) GetClientSource()(*SubmissionClientSource) {
    return m.clientSource
}
// GetContentType gets the contentType property value. Specifies the type of content being submitted. The possible values are: email, url, file, app and unkownFutureValue.
func (m *ThreatSubmission) GetContentType()(*SubmissionContentType) {
    return m.contentType
}
// GetCreatedBy gets the createdBy property value. Specifies who submitted the email as a threat. Supports $filter = createdBy/email eq 'value'.
func (m *ThreatSubmission) GetCreatedBy()(SubmissionUserIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Specifies when the threat submission was created. Supports $filter = createdDateTime ge 2022-01-01T00:00:00Z and createdDateTime lt 2022-01-02T00:00:00Z.
func (m *ThreatSubmission) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ThreatSubmission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["adminReview"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSubmissionAdminReviewFromDiscriminatorValue , m.SetAdminReview)
    res["category"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSubmissionCategory , m.SetCategory)
    res["clientSource"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSubmissionClientSource , m.SetClientSource)
    res["contentType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSubmissionContentType , m.SetContentType)
    res["createdBy"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSubmissionUserIdentityFromDiscriminatorValue , m.SetCreatedBy)
    res["createdDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreatedDateTime)
    res["result"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSubmissionResultFromDiscriminatorValue , m.SetResult)
    res["source"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSubmissionSource , m.SetSource)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseLongRunningOperationStatus , m.SetStatus)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    return res
}
// GetResult gets the result property value. Specifies the result of the analysis performed by Microsoft.
func (m *ThreatSubmission) GetResult()(SubmissionResultable) {
    return m.result
}
// GetSource gets the source property value. Specifies the role of the submitter. Supports $filter = source eq 'value'. The possible values are: administrator,  user and unkownFutureValue.
func (m *ThreatSubmission) GetSource()(*SubmissionSource) {
    return m.source
}
// GetStatus gets the status property value. Indicates whether the threat submission has been analyzed by Microsoft. Supports $filter = status eq 'value'. The possible values are: notStarted, running, succeeded, failed, skipped and unkownFutureValue.
func (m *ThreatSubmission) GetStatus()(*LongRunningOperationStatus) {
    return m.status
}
// GetTenantId gets the tenantId property value. Indicates the tenant id of the submitter. Not required when created using a POST operation. It is extracted from the token of the post API call.
func (m *ThreatSubmission) GetTenantId()(*string) {
    return m.tenantId
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
// SetAdminReview sets the adminReview property value. Specifies the admin review property which constitutes of who reviewed the user submission, when and what was it identified as.
func (m *ThreatSubmission) SetAdminReview(value SubmissionAdminReviewable)() {
    m.adminReview = value
}
// SetCategory sets the category property value. The category property
func (m *ThreatSubmission) SetCategory(value *SubmissionCategory)() {
    m.category = value
}
// SetClientSource sets the clientSource property value. Specifies the source of the submission. The possible values are: microsoft,  other and unkownFutureValue.
func (m *ThreatSubmission) SetClientSource(value *SubmissionClientSource)() {
    m.clientSource = value
}
// SetContentType sets the contentType property value. Specifies the type of content being submitted. The possible values are: email, url, file, app and unkownFutureValue.
func (m *ThreatSubmission) SetContentType(value *SubmissionContentType)() {
    m.contentType = value
}
// SetCreatedBy sets the createdBy property value. Specifies who submitted the email as a threat. Supports $filter = createdBy/email eq 'value'.
func (m *ThreatSubmission) SetCreatedBy(value SubmissionUserIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Specifies when the threat submission was created. Supports $filter = createdDateTime ge 2022-01-01T00:00:00Z and createdDateTime lt 2022-01-02T00:00:00Z.
func (m *ThreatSubmission) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetResult sets the result property value. Specifies the result of the analysis performed by Microsoft.
func (m *ThreatSubmission) SetResult(value SubmissionResultable)() {
    m.result = value
}
// SetSource sets the source property value. Specifies the role of the submitter. Supports $filter = source eq 'value'. The possible values are: administrator,  user and unkownFutureValue.
func (m *ThreatSubmission) SetSource(value *SubmissionSource)() {
    m.source = value
}
// SetStatus sets the status property value. Indicates whether the threat submission has been analyzed by Microsoft. Supports $filter = status eq 'value'. The possible values are: notStarted, running, succeeded, failed, skipped and unkownFutureValue.
func (m *ThreatSubmission) SetStatus(value *LongRunningOperationStatus)() {
    m.status = value
}
// SetTenantId sets the tenantId property value. Indicates the tenant id of the submitter. Not required when created using a POST operation. It is extracted from the token of the post API call.
func (m *ThreatSubmission) SetTenantId(value *string)() {
    m.tenantId = value
}
