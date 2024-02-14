package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ExactMatchSessionBase struct {
    ExactMatchJobBase
}
// NewExactMatchSessionBase instantiates a new ExactMatchSessionBase and sets the default values.
func NewExactMatchSessionBase()(*ExactMatchSessionBase) {
    m := &ExactMatchSessionBase{
        ExactMatchJobBase: *NewExactMatchJobBase(),
    }
    odataTypeValue := "#microsoft.graph.exactMatchSessionBase"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateExactMatchSessionBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
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
// returns a *string when successful
func (m *ExactMatchSessionBase) GetDataStoreId()(*string) {
    val, err := m.GetBackingStore().Get("dataStoreId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ExactMatchSessionBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ExactMatchJobBase.GetFieldDeserializers()
    res["dataStoreId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataStoreId(val)
        }
        return nil
    }
    res["processingCompletionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingCompletionDateTime(val)
        }
        return nil
    }
    res["remainingBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemainingBlockCount(val)
        }
        return nil
    }
    res["remainingJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemainingJobCount(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["totalBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBlockCount(val)
        }
        return nil
    }
    res["totalJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalJobCount(val)
        }
        return nil
    }
    res["uploadCompletionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadCompletionDateTime(val)
        }
        return nil
    }
    return res
}
// GetProcessingCompletionDateTime gets the processingCompletionDateTime property value. The processingCompletionDateTime property
// returns a *Time when successful
func (m *ExactMatchSessionBase) GetProcessingCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("processingCompletionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRemainingBlockCount gets the remainingBlockCount property value. The remainingBlockCount property
// returns a *int32 when successful
func (m *ExactMatchSessionBase) GetRemainingBlockCount()(*int32) {
    val, err := m.GetBackingStore().Get("remainingBlockCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRemainingJobCount gets the remainingJobCount property value. The remainingJobCount property
// returns a *int32 when successful
func (m *ExactMatchSessionBase) GetRemainingJobCount()(*int32) {
    val, err := m.GetBackingStore().Get("remainingJobCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetState gets the state property value. The state property
// returns a *string when successful
func (m *ExactMatchSessionBase) GetState()(*string) {
    val, err := m.GetBackingStore().Get("state")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTotalBlockCount gets the totalBlockCount property value. The totalBlockCount property
// returns a *int32 when successful
func (m *ExactMatchSessionBase) GetTotalBlockCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalBlockCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetTotalJobCount gets the totalJobCount property value. The totalJobCount property
// returns a *int32 when successful
func (m *ExactMatchSessionBase) GetTotalJobCount()(*int32) {
    val, err := m.GetBackingStore().Get("totalJobCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetUploadCompletionDateTime gets the uploadCompletionDateTime property value. The uploadCompletionDateTime property
// returns a *Time when successful
func (m *ExactMatchSessionBase) GetUploadCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("uploadCompletionDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
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
    err := m.GetBackingStore().Set("dataStoreId", value)
    if err != nil {
        panic(err)
    }
}
// SetProcessingCompletionDateTime sets the processingCompletionDateTime property value. The processingCompletionDateTime property
func (m *ExactMatchSessionBase) SetProcessingCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("processingCompletionDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRemainingBlockCount sets the remainingBlockCount property value. The remainingBlockCount property
func (m *ExactMatchSessionBase) SetRemainingBlockCount(value *int32)() {
    err := m.GetBackingStore().Set("remainingBlockCount", value)
    if err != nil {
        panic(err)
    }
}
// SetRemainingJobCount sets the remainingJobCount property value. The remainingJobCount property
func (m *ExactMatchSessionBase) SetRemainingJobCount(value *int32)() {
    err := m.GetBackingStore().Set("remainingJobCount", value)
    if err != nil {
        panic(err)
    }
}
// SetState sets the state property value. The state property
func (m *ExactMatchSessionBase) SetState(value *string)() {
    err := m.GetBackingStore().Set("state", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalBlockCount sets the totalBlockCount property value. The totalBlockCount property
func (m *ExactMatchSessionBase) SetTotalBlockCount(value *int32)() {
    err := m.GetBackingStore().Set("totalBlockCount", value)
    if err != nil {
        panic(err)
    }
}
// SetTotalJobCount sets the totalJobCount property value. The totalJobCount property
func (m *ExactMatchSessionBase) SetTotalJobCount(value *int32)() {
    err := m.GetBackingStore().Set("totalJobCount", value)
    if err != nil {
        panic(err)
    }
}
// SetUploadCompletionDateTime sets the uploadCompletionDateTime property value. The uploadCompletionDateTime property
func (m *ExactMatchSessionBase) SetUploadCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("uploadCompletionDateTime", value)
    if err != nil {
        panic(err)
    }
}
type ExactMatchSessionBaseable interface {
    ExactMatchJobBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataStoreId()(*string)
    GetProcessingCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRemainingBlockCount()(*int32)
    GetRemainingJobCount()(*int32)
    GetState()(*string)
    GetTotalBlockCount()(*int32)
    GetTotalJobCount()(*int32)
    GetUploadCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDataStoreId(value *string)()
    SetProcessingCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRemainingBlockCount(value *int32)()
    SetRemainingJobCount(value *int32)()
    SetState(value *string)()
    SetTotalBlockCount(value *int32)()
    SetTotalJobCount(value *int32)()
    SetUploadCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
