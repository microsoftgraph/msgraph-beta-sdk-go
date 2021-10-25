package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ExactMatchSessionBase struct {
    ExactMatchJobBase
    dataStoreId *string;
    processingCompletionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    remainingBlockCount *int32;
    remainingJobCount *int32;
    state *string;
    totalBlockCount *int32;
    totalJobCount *int32;
    uploadCompletionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewExactMatchSessionBase()(*ExactMatchSessionBase) {
    m := &ExactMatchSessionBase{
        ExactMatchJobBase: *NewExactMatchJobBase(),
    }
    return m
}
func (m *ExactMatchSessionBase) GetDataStoreId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataStoreId
    }
}
func (m *ExactMatchSessionBase) GetProcessingCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.processingCompletionDateTime
    }
}
func (m *ExactMatchSessionBase) GetRemainingBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remainingBlockCount
    }
}
func (m *ExactMatchSessionBase) GetRemainingJobCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remainingJobCount
    }
}
func (m *ExactMatchSessionBase) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *ExactMatchSessionBase) GetTotalBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalBlockCount
    }
}
func (m *ExactMatchSessionBase) GetTotalJobCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalJobCount
    }
}
func (m *ExactMatchSessionBase) GetUploadCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.uploadCompletionDateTime
    }
}
func (m *ExactMatchSessionBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ExactMatchJobBase.GetFieldDeserializers()
    res["dataStoreId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDataStoreId(val)
        return nil
    }
    res["processingCompletionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetProcessingCompletionDateTime(val)
        return nil
    }
    res["remainingBlockCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemainingBlockCount(val)
        return nil
    }
    res["remainingJobCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRemainingJobCount(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetState(val)
        return nil
    }
    res["totalBlockCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalBlockCount(val)
        return nil
    }
    res["totalJobCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotalJobCount(val)
        return nil
    }
    res["uploadCompletionDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetUploadCompletionDateTime(val)
        return nil
    }
    return res
}
func (m *ExactMatchSessionBase) IsNil()(bool) {
    return m == nil
}
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
func (m *ExactMatchSessionBase) SetDataStoreId(value *string)() {
    m.dataStoreId = value
}
func (m *ExactMatchSessionBase) SetProcessingCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.processingCompletionDateTime = value
}
func (m *ExactMatchSessionBase) SetRemainingBlockCount(value *int32)() {
    m.remainingBlockCount = value
}
func (m *ExactMatchSessionBase) SetRemainingJobCount(value *int32)() {
    m.remainingJobCount = value
}
func (m *ExactMatchSessionBase) SetState(value *string)() {
    m.state = value
}
func (m *ExactMatchSessionBase) SetTotalBlockCount(value *int32)() {
    m.totalBlockCount = value
}
func (m *ExactMatchSessionBase) SetTotalJobCount(value *int32)() {
    m.totalJobCount = value
}
func (m *ExactMatchSessionBase) SetUploadCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.uploadCompletionDateTime = value
}
