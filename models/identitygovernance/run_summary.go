package identitygovernance

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RunSummary 
type RunSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The number of failed workflow runs.
    failedRuns *int32
    // The number of failed tasks of a workflow.
    failedTasks *int32
    // The OdataType property
    odataType *string
    // The number of successful workflow runs.
    successfulRuns *int32
    // The total number of runs for a workflow.
    totalRuns *int32
    // The totalTasks property
    totalTasks *int32
    // The totalUsers property
    totalUsers *int32
}
// NewRunSummary instantiates a new runSummary and sets the default values.
func NewRunSummary()(*RunSummary) {
    m := &RunSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.identityGovernance.runSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateRunSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRunSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRunSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RunSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFailedRuns gets the failedRuns property value. The number of failed workflow runs.
func (m *RunSummary) GetFailedRuns()(*int32) {
    return m.failedRuns
}
// GetFailedTasks gets the failedTasks property value. The number of failed tasks of a workflow.
func (m *RunSummary) GetFailedTasks()(*int32) {
    return m.failedTasks
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RunSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["failedRuns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedRuns(val)
        }
        return nil
    }
    res["failedTasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedTasks(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["successfulRuns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulRuns(val)
        }
        return nil
    }
    res["totalRuns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalRuns(val)
        }
        return nil
    }
    res["totalTasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalTasks(val)
        }
        return nil
    }
    res["totalUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalUsers(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RunSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetSuccessfulRuns gets the successfulRuns property value. The number of successful workflow runs.
func (m *RunSummary) GetSuccessfulRuns()(*int32) {
    return m.successfulRuns
}
// GetTotalRuns gets the totalRuns property value. The total number of runs for a workflow.
func (m *RunSummary) GetTotalRuns()(*int32) {
    return m.totalRuns
}
// GetTotalTasks gets the totalTasks property value. The totalTasks property
func (m *RunSummary) GetTotalTasks()(*int32) {
    return m.totalTasks
}
// GetTotalUsers gets the totalUsers property value. The totalUsers property
func (m *RunSummary) GetTotalUsers()(*int32) {
    return m.totalUsers
}
// Serialize serializes information the current object
func (m *RunSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("failedRuns", m.GetFailedRuns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("failedTasks", m.GetFailedTasks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("successfulRuns", m.GetSuccessfulRuns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalRuns", m.GetTotalRuns())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalTasks", m.GetTotalTasks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalUsers", m.GetTotalUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RunSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetFailedRuns sets the failedRuns property value. The number of failed workflow runs.
func (m *RunSummary) SetFailedRuns(value *int32)() {
    m.failedRuns = value
}
// SetFailedTasks sets the failedTasks property value. The number of failed tasks of a workflow.
func (m *RunSummary) SetFailedTasks(value *int32)() {
    m.failedTasks = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RunSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSuccessfulRuns sets the successfulRuns property value. The number of successful workflow runs.
func (m *RunSummary) SetSuccessfulRuns(value *int32)() {
    m.successfulRuns = value
}
// SetTotalRuns sets the totalRuns property value. The total number of runs for a workflow.
func (m *RunSummary) SetTotalRuns(value *int32)() {
    m.totalRuns = value
}
// SetTotalTasks sets the totalTasks property value. The totalTasks property
func (m *RunSummary) SetTotalTasks(value *int32)() {
    m.totalTasks = value
}
// SetTotalUsers sets the totalUsers property value. The totalUsers property
func (m *RunSummary) SetTotalUsers(value *int32)() {
    m.totalUsers = value
}
