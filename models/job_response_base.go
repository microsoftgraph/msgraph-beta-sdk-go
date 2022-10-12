package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// JobResponseBase 
type JobResponseBase struct {
    Entity
    // The creationDateTime property
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The endDateTime property
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The error property
    error ClassificationErrorable
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *string
    // The tenantId property
    tenantId *string
    // The type property
    type_escaped *string
}
// NewJobResponseBase instantiates a new JobResponseBase and sets the default values.
func NewJobResponseBase()(*JobResponseBase) {
    m := &JobResponseBase{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.jobResponseBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateJobResponseBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateJobResponseBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.classificationJobResponse":
                        return NewClassificationJobResponse(), nil
                    case "#microsoft.graph.dlpEvaluatePoliciesJobResponse":
                        return NewDlpEvaluatePoliciesJobResponse(), nil
                    case "#microsoft.graph.evaluateLabelJobResponse":
                        return NewEvaluateLabelJobResponse(), nil
                }
            }
        }
    }
    return NewJobResponseBase(), nil
}
// GetCreationDateTime gets the creationDateTime property value. The creationDateTime property
func (m *JobResponseBase) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.creationDateTime
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *JobResponseBase) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetError gets the error property value. The error property
func (m *JobResponseBase) GetError()(ClassificationErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *JobResponseBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["creationDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetCreationDateTime)
    res["endDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetEndDateTime)
    res["error"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateClassificationErrorFromDiscriminatorValue , m.SetError)
    res["startDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetStartDateTime)
    res["status"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetStatus)
    res["tenantId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTenantId)
    res["type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetType)
    return res
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *JobResponseBase) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatus gets the status property value. The status property
func (m *JobResponseBase) GetStatus()(*string) {
    return m.status
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *JobResponseBase) GetTenantId()(*string) {
    return m.tenantId
}
// GetType gets the type property value. The type property
func (m *JobResponseBase) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *JobResponseBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
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
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
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
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreationDateTime sets the creationDateTime property value. The creationDateTime property
func (m *JobResponseBase) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *JobResponseBase) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetError sets the error property value. The error property
func (m *JobResponseBase) SetError(value ClassificationErrorable)() {
    m.error = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *JobResponseBase) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *JobResponseBase) SetStatus(value *string)() {
    m.status = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *JobResponseBase) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetType sets the type property value. The type property
func (m *JobResponseBase) SetType(value *string)() {
    m.type_escaped = value
}
