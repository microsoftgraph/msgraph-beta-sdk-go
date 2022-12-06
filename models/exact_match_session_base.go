package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExactMatchSessionBase 
type ExactMatchSessionBase struct {
    ExactMatchJobBase
    // The dataStoreId property
    dataStoreId *string
    // The processingCompletionDateTime property
    processingCompletionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The remainingBlockCount property
    remainingBlockCount *int32
    // The remainingJobCount property
    remainingJobCount *int32
    // The state property
    state *string
    // The totalBlockCount property
    totalBlockCount *int32
    // The totalJobCount property
    totalJobCount *int32
    // The uploadCompletionDateTime property
    uploadCompletionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewExactMatchSessionBase instantiates a new ExactMatchSessionBase and sets the default values.
func NewExactMatchSessionBase()(*ExactMatchSessionBase) {
    m := &ExactMatchSessionBase{
        ExactMatchJobBase: *NewExactMatchJobBase(),
    }
    odataTypeValue := "#microsoft.graph.exactMatchSessionBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExactMatchSessionBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExactMatchSessionBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.exactMatchSession":
                        return NewExactMatchSession(), nil
                }
            }
        }
    }
    return NewExactMatchSessionBase(), nil
}
// GetDataStoreId gets the dataStoreId property value. The dataStoreId property
func (m *ExactMatchSessionBase) GetDataStoreId()(*string) {
    return m.dataStoreId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchSessionBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExactMatchJobBase.GetFieldDeserializers()
    res["dataStoreId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDataStoreId)
    res["processingCompletionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetProcessingCompletionDateTime)
    res["remainingBlockCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRemainingBlockCount)
    res["remainingJobCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetRemainingJobCount)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetState)
    res["totalBlockCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalBlockCount)
    res["totalJobCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetTotalJobCount)
    res["uploadCompletionDateTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetUploadCompletionDateTime)
    return res
}
// GetProcessingCompletionDateTime gets the processingCompletionDateTime property value. The processingCompletionDateTime property
func (m *ExactMatchSessionBase) GetProcessingCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.processingCompletionDateTime
}
// GetRemainingBlockCount gets the remainingBlockCount property value. The remainingBlockCount property
func (m *ExactMatchSessionBase) GetRemainingBlockCount()(*int32) {
    return m.remainingBlockCount
}
// GetRemainingJobCount gets the remainingJobCount property value. The remainingJobCount property
func (m *ExactMatchSessionBase) GetRemainingJobCount()(*int32) {
    return m.remainingJobCount
}
// GetState gets the state property value. The state property
func (m *ExactMatchSessionBase) GetState()(*string) {
    return m.state
}
// GetTotalBlockCount gets the totalBlockCount property value. The totalBlockCount property
func (m *ExactMatchSessionBase) GetTotalBlockCount()(*int32) {
    return m.totalBlockCount
}
// GetTotalJobCount gets the totalJobCount property value. The totalJobCount property
func (m *ExactMatchSessionBase) GetTotalJobCount()(*int32) {
    return m.totalJobCount
}
// GetUploadCompletionDateTime gets the uploadCompletionDateTime property value. The uploadCompletionDateTime property
func (m *ExactMatchSessionBase) GetUploadCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.uploadCompletionDateTime
}
// Serialize serializes information the current object
func (m *ExactMatchSessionBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ExactMatchJobBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("dataStoreId", m.GetDataStoreId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("processingCompletionDateTime", m.GetProcessingCompletionDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remainingBlockCount", m.GetRemainingBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remainingJobCount", m.GetRemainingJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalBlockCount", m.GetTotalBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalJobCount", m.GetTotalJobCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("uploadCompletionDateTime", m.GetUploadCompletionDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDataStoreId sets the dataStoreId property value. The dataStoreId property
func (m *ExactMatchSessionBase) SetDataStoreId(value *string)() {
    m.dataStoreId = value
}
// SetProcessingCompletionDateTime sets the processingCompletionDateTime property value. The processingCompletionDateTime property
func (m *ExactMatchSessionBase) SetProcessingCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.processingCompletionDateTime = value
}
// SetRemainingBlockCount sets the remainingBlockCount property value. The remainingBlockCount property
func (m *ExactMatchSessionBase) SetRemainingBlockCount(value *int32)() {
    m.remainingBlockCount = value
}
// SetRemainingJobCount sets the remainingJobCount property value. The remainingJobCount property
func (m *ExactMatchSessionBase) SetRemainingJobCount(value *int32)() {
    m.remainingJobCount = value
}
// SetState sets the state property value. The state property
func (m *ExactMatchSessionBase) SetState(value *string)() {
    m.state = value
}
// SetTotalBlockCount sets the totalBlockCount property value. The totalBlockCount property
func (m *ExactMatchSessionBase) SetTotalBlockCount(value *int32)() {
    m.totalBlockCount = value
}
// SetTotalJobCount sets the totalJobCount property value. The totalJobCount property
func (m *ExactMatchSessionBase) SetTotalJobCount(value *int32)() {
    m.totalJobCount = value
}
// SetUploadCompletionDateTime sets the uploadCompletionDateTime property value. The uploadCompletionDateTime property
func (m *ExactMatchSessionBase) SetUploadCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.uploadCompletionDateTime = value
}
