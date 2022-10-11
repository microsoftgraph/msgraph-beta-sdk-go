package devicemanagement

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AlertRecord provides operations to manage the collection of accessReviewDecision entities.
type AlertRecord struct {
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Entity
    // The alertImpact property
    alertImpact AlertImpactable
    // The alertRuleId property
    alertRuleId *string
    // The alertRuleTemplate property
    alertRuleTemplate *AlertRuleTemplate
    // The detectedDateTime property
    detectedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The displayName property
    displayName *string
    // The lastUpdatedDateTime property
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The resolvedDateTime property
    resolvedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The severity property
    severity *RuleSeverityType
    // The status property
    status *AlertStatusType
}
// NewAlertRecord instantiates a new alertRecord and sets the default values.
func NewAlertRecord()(*AlertRecord) {
    m := &AlertRecord{
        Entity: *ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagement.alertRecord";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAlertRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAlertRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAlertRecord(), nil
}
// GetAlertImpact gets the alertImpact property value. The alertImpact property
func (m *AlertRecord) GetAlertImpact()(AlertImpactable) {
    return m.alertImpact
}
// GetAlertRuleId gets the alertRuleId property value. The alertRuleId property
func (m *AlertRecord) GetAlertRuleId()(*string) {
    return m.alertRuleId
}
// GetAlertRuleTemplate gets the alertRuleTemplate property value. The alertRuleTemplate property
func (m *AlertRecord) GetAlertRuleTemplate()(*AlertRuleTemplate) {
    return m.alertRuleTemplate
}
// GetDetectedDateTime gets the detectedDateTime property value. The detectedDateTime property
func (m *AlertRecord) GetDetectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.detectedDateTime
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *AlertRecord) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlertRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["alertImpact"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAlertImpactFromDiscriminatorValue , m.SetAlertImpact)
    res["alertRuleId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetAlertRuleId)
    res["alertRuleTemplate"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAlertRuleTemplate , m.SetAlertRuleTemplate)
    res["detectedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetDetectedDateTime)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["lastUpdatedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetLastUpdatedDateTime)
    res["resolvedDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetResolvedDateTime)
    res["severity"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRuleSeverityType , m.SetSeverity)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseAlertStatusType , m.SetStatus)
    return res
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *AlertRecord) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetResolvedDateTime gets the resolvedDateTime property value. The resolvedDateTime property
func (m *AlertRecord) GetResolvedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.resolvedDateTime
}
// GetSeverity gets the severity property value. The severity property
func (m *AlertRecord) GetSeverity()(*RuleSeverityType) {
    return m.severity
}
// GetStatus gets the status property value. The status property
func (m *AlertRecord) GetStatus()(*AlertStatusType) {
    return m.status
}
// Serialize serializes information the current object
func (m *AlertRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("alertImpact", m.GetAlertImpact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("alertRuleId", m.GetAlertRuleId())
        if err != nil {
            return err
        }
    }
    if m.GetAlertRuleTemplate() != nil {
        cast := (*m.GetAlertRuleTemplate()).String()
        err = writer.WriteStringValue("alertRuleTemplate", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("detectedDateTime", m.GetDetectedDateTime())
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
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("resolvedDateTime", m.GetResolvedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSeverity() != nil {
        cast := (*m.GetSeverity()).String()
        err = writer.WriteStringValue("severity", &cast)
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
// SetAlertImpact sets the alertImpact property value. The alertImpact property
func (m *AlertRecord) SetAlertImpact(value AlertImpactable)() {
    m.alertImpact = value
}
// SetAlertRuleId sets the alertRuleId property value. The alertRuleId property
func (m *AlertRecord) SetAlertRuleId(value *string)() {
    m.alertRuleId = value
}
// SetAlertRuleTemplate sets the alertRuleTemplate property value. The alertRuleTemplate property
func (m *AlertRecord) SetAlertRuleTemplate(value *AlertRuleTemplate)() {
    m.alertRuleTemplate = value
}
// SetDetectedDateTime sets the detectedDateTime property value. The detectedDateTime property
func (m *AlertRecord) SetDetectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.detectedDateTime = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AlertRecord) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. The lastUpdatedDateTime property
func (m *AlertRecord) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetResolvedDateTime sets the resolvedDateTime property value. The resolvedDateTime property
func (m *AlertRecord) SetResolvedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.resolvedDateTime = value
}
// SetSeverity sets the severity property value. The severity property
func (m *AlertRecord) SetSeverity(value *RuleSeverityType)() {
    m.severity = value
}
// SetStatus sets the status property value. The status property
func (m *AlertRecord) SetStatus(value *AlertStatusType)() {
    m.status = value
}
