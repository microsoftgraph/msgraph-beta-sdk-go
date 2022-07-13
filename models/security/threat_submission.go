package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// ThreatSubmission provides operations to manage the collection of accessReviewDecision entities.
type ThreatSubmission struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The adminReview property
    adminReview SubmissionAdminReviewable
    // The category property
    category *SubmissionCategory
    // The clientSource property
    clientSource *SubmissionClientSource
    // The contentType property
    contentType *SubmissionContentType
    // The createdBy property
    createdBy SubmissionUserIdentityable
    // The createdDateTime property
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The result property
    result SubmissionResultable
    // The source property
    source *SubmissionSource
    // The status property
    status *LongRunningOperationStatus
    // The tenantId property
    tenantId *string
}
// NewThreatSubmission instantiates a new threatSubmission and sets the default values.
func NewThreatSubmission()(*ThreatSubmission) {
    m := &ThreatSubmission{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odatatypeValue := "#microsoft.graph.security.threatSubmission";
    m.SetOdatatype(&odatatypeValue);
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
                mappingStr := *mappingValue
                switch mappingStr {
                    case "#microsoft.graph.security.emailThreatSubmission":
                        return NewEmailThreatSubmission(), nil
                    case "#microsoft.graph.security.fileThreatSubmission":
                        return NewFileThreatSubmission(), nil
                    case "#microsoft.graph.security.urlThreatSubmission":
                        return NewUrlThreatSubmission(), nil
                }
            }
        }
    }
    return NewThreatSubmission(), nil
}
// GetAdminReview gets the adminReview property value. The adminReview property
func (m *ThreatSubmission) GetAdminReview()(SubmissionAdminReviewable) {
    if m == nil {
        return nil
    } else {
        return m.adminReview
    }
}
// GetCategory gets the category property value. The category property
func (m *ThreatSubmission) GetCategory()(*SubmissionCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetClientSource gets the clientSource property value. The clientSource property
func (m *ThreatSubmission) GetClientSource()(*SubmissionClientSource) {
    if m == nil {
        return nil
    } else {
        return m.clientSource
    }
}
// GetContentType gets the contentType property value. The contentType property
func (m *ThreatSubmission) GetContentType()(*SubmissionContentType) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *ThreatSubmission) GetCreatedBy()(SubmissionUserIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The createdDateTime property
func (m *ThreatSubmission) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// GetResult gets the result property value. The result property
func (m *ThreatSubmission) GetResult()(SubmissionResultable) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
// GetSource gets the source property value. The source property
func (m *ThreatSubmission) GetSource()(*SubmissionSource) {
    if m == nil {
        return nil
    } else {
        return m.source
    }
}
// GetStatus gets the status property value. The status property
func (m *ThreatSubmission) GetStatus()(*LongRunningOperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *ThreatSubmission) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
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
// SetAdminReview sets the adminReview property value. The adminReview property
func (m *ThreatSubmission) SetAdminReview(value SubmissionAdminReviewable)() {
    if m != nil {
        m.adminReview = value
    }
}
// SetCategory sets the category property value. The category property
func (m *ThreatSubmission) SetCategory(value *SubmissionCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetClientSource sets the clientSource property value. The clientSource property
func (m *ThreatSubmission) SetClientSource(value *SubmissionClientSource)() {
    if m != nil {
        m.clientSource = value
    }
}
// SetContentType sets the contentType property value. The contentType property
func (m *ThreatSubmission) SetContentType(value *SubmissionContentType)() {
    if m != nil {
        m.contentType = value
    }
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *ThreatSubmission) SetCreatedBy(value SubmissionUserIdentityable)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The createdDateTime property
func (m *ThreatSubmission) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetResult sets the result property value. The result property
func (m *ThreatSubmission) SetResult(value SubmissionResultable)() {
    if m != nil {
        m.result = value
    }
}
// SetSource sets the source property value. The source property
func (m *ThreatSubmission) SetSource(value *SubmissionSource)() {
    if m != nil {
        m.source = value
    }
}
// SetStatus sets the status property value. The status property
func (m *ThreatSubmission) SetStatus(value *LongRunningOperationStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *ThreatSubmission) SetTenantId(value *string)() {
    if m != nil {
        m.tenantId = value
    }
}
