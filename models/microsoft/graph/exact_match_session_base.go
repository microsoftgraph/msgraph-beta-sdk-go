package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new exactMatchSessionBase and sets the default values.
func NewExactMatchSessionBase()(*ExactMatchSessionBase) {
    m := &ExactMatchSessionBase{
        ExactMatchJobBase: *NewExactMatchJobBase(),
    }
    return m
}
// Gets the dataStoreId property value. 
func (m *ExactMatchSessionBase) GetDataStoreId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataStoreId
    }
}
// Gets the processingCompletionDateTime property value. 
func (m *ExactMatchSessionBase) GetProcessingCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.processingCompletionDateTime
    }
}
// Gets the remainingBlockCount property value. 
func (m *ExactMatchSessionBase) GetRemainingBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remainingBlockCount
    }
}
// Gets the remainingJobCount property value. 
func (m *ExactMatchSessionBase) GetRemainingJobCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.remainingJobCount
    }
}
// Gets the state property value. 
func (m *ExactMatchSessionBase) GetState()(*string) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the totalBlockCount property value. 
func (m *ExactMatchSessionBase) GetTotalBlockCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalBlockCount
    }
}
// Gets the totalJobCount property value. 
func (m *ExactMatchSessionBase) GetTotalJobCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalJobCount
    }
}
// Gets the uploadCompletionDateTime property value. 
func (m *ExactMatchSessionBase) GetUploadCompletionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.uploadCompletionDateTime
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the dataStoreId property value. 
// Parameters:
//  - value : Value to set for the dataStoreId property.
func (m *ExactMatchSessionBase) SetDataStoreId(value *string)() {
    m.dataStoreId = value
}
// Sets the processingCompletionDateTime property value. 
// Parameters:
//  - value : Value to set for the processingCompletionDateTime property.
func (m *ExactMatchSessionBase) SetProcessingCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.processingCompletionDateTime = value
}
// Sets the remainingBlockCount property value. 
// Parameters:
//  - value : Value to set for the remainingBlockCount property.
func (m *ExactMatchSessionBase) SetRemainingBlockCount(value *int32)() {
    m.remainingBlockCount = value
}
// Sets the remainingJobCount property value. 
// Parameters:
//  - value : Value to set for the remainingJobCount property.
func (m *ExactMatchSessionBase) SetRemainingJobCount(value *int32)() {
    m.remainingJobCount = value
}
// Sets the state property value. 
// Parameters:
//  - value : Value to set for the state property.
func (m *ExactMatchSessionBase) SetState(value *string)() {
    m.state = value
}
// Sets the totalBlockCount property value. 
// Parameters:
//  - value : Value to set for the totalBlockCount property.
func (m *ExactMatchSessionBase) SetTotalBlockCount(value *int32)() {
    m.totalBlockCount = value
}
// Sets the totalJobCount property value. 
// Parameters:
//  - value : Value to set for the totalJobCount property.
func (m *ExactMatchSessionBase) SetTotalJobCount(value *int32)() {
    m.totalJobCount = value
}
// Sets the uploadCompletionDateTime property value. 
// Parameters:
//  - value : Value to set for the uploadCompletionDateTime property.
func (m *ExactMatchSessionBase) SetUploadCompletionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.uploadCompletionDateTime = value
}
