package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactMatchSessionBase 
type ExactMatchSessionBase struct {
    ExactMatchJobBase
    // 
    dataStoreId *string;
    // 
    processingCompletionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    remainingBlockCount *int32;
    // 
    remainingJobCount *int32;
    // 
    state *string;
    // 
    totalBlockCount *int32;
    // 
    totalJobCount *int32;
    // 
    uploadCompletionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewExactMatchSessionBase instantiates a new exactMatchSessionBase and sets the default values.
func NewExactMatchSessionBase()(*ExactMatchSessionBase) {
    m := &ExactMatchSessionBase{
        ExactMatchJobBase: *NewExactMatchJobBase(),
    }
    return m
}
// GetDataStoreId gets the dataStoreId property value. 
func (m *ExactMatchSessionBase) GetDataStoreId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataStoreId
    }
}
// GetProcessingCompletionDateTime gets the processingCompletionDateTime property value. 
func (m *ExactMatchSessionBase) GetProcessingCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.processingCompletionDateTime
    }
}
// GetRemainingBlockCount gets the remainingBlockCount property value. 
func (m *ExactMatchSessionBase) GetRemainingBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remainingBlockCount
    }
}
// GetRemainingJobCount gets the remainingJobCount property value. 
func (m *ExactMatchSessionBase) GetRemainingJobCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remainingJobCount
    }
}
// GetState gets the state property value. 
func (m *ExactMatchSessionBase) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetTotalBlockCount gets the totalBlockCount property value. 
func (m *ExactMatchSessionBase) GetTotalBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalBlockCount
    }
}
// GetTotalJobCount gets the totalJobCount property value. 
func (m *ExactMatchSessionBase) GetTotalJobCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalJobCount
    }
}
// GetUploadCompletionDateTime gets the uploadCompletionDateTime property value. 
func (m *ExactMatchSessionBase) GetUploadCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.uploadCompletionDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchSessionBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ExactMatchJobBase.GetFieldDeserializers()
    res["dataStoreId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataStoreId(val)
        }
        return nil
    }
    res["processingCompletionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessingCompletionDateTime(val)
        }
        return nil
    }
    res["remainingBlockCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemainingBlockCount(val)
        }
        return nil
    }
    res["remainingJobCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemainingJobCount(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["totalBlockCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalBlockCount(val)
        }
        return nil
    }
    res["totalJobCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalJobCount(val)
        }
        return nil
    }
    res["uploadCompletionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *ExactMatchSessionBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExactMatchSessionBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetDataStoreId sets the dataStoreId property value. 
func (m *ExactMatchSessionBase) SetDataStoreId(value *string)() {
    if m != nil {
        m.dataStoreId = value
    }
}
// SetProcessingCompletionDateTime sets the processingCompletionDateTime property value. 
func (m *ExactMatchSessionBase) SetProcessingCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.processingCompletionDateTime = value
    }
}
// SetRemainingBlockCount sets the remainingBlockCount property value. 
func (m *ExactMatchSessionBase) SetRemainingBlockCount(value *int32)() {
    if m != nil {
        m.remainingBlockCount = value
    }
}
// SetRemainingJobCount sets the remainingJobCount property value. 
func (m *ExactMatchSessionBase) SetRemainingJobCount(value *int32)() {
    if m != nil {
        m.remainingJobCount = value
    }
}
// SetState sets the state property value. 
func (m *ExactMatchSessionBase) SetState(value *string)() {
    if m != nil {
        m.state = value
    }
}
// SetTotalBlockCount sets the totalBlockCount property value. 
func (m *ExactMatchSessionBase) SetTotalBlockCount(value *int32)() {
    if m != nil {
        m.totalBlockCount = value
    }
}
// SetTotalJobCount sets the totalJobCount property value. 
func (m *ExactMatchSessionBase) SetTotalJobCount(value *int32)() {
    if m != nil {
        m.totalJobCount = value
    }
}
// SetUploadCompletionDateTime sets the uploadCompletionDateTime property value. 
func (m *ExactMatchSessionBase) SetUploadCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.uploadCompletionDateTime = value
    }
}
