package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationStatus struct {
    additionalData map[string]interface{};
    code *SynchronizationStatusCode;
    countSuccessiveCompleteFailures *int64;
    escrowsPruned *bool;
    lastExecution *SynchronizationTaskExecution;
    lastSuccessfulExecution *SynchronizationTaskExecution;
    lastSuccessfulExecutionWithExports *SynchronizationTaskExecution;
    progress []SynchronizationProgress;
    quarantine *SynchronizationQuarantine;
    steadyStateFirstAchievedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    steadyStateLastAchievedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    synchronizedEntryCountByType []StringKeyLongValuePair;
    troubleshootingUrl *string;
}
func NewSynchronizationStatus()(*SynchronizationStatus) {
    m := &SynchronizationStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SynchronizationStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SynchronizationStatus) GetCode()(*SynchronizationStatusCode) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
func (m *SynchronizationStatus) GetCountSuccessiveCompleteFailures()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countSuccessiveCompleteFailures
    }
}
func (m *SynchronizationStatus) GetEscrowsPruned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.escrowsPruned
    }
}
func (m *SynchronizationStatus) GetLastExecution()(*SynchronizationTaskExecution) {
    if m == nil {
        return nil
    } else {
        return m.lastExecution
    }
}
func (m *SynchronizationStatus) GetLastSuccessfulExecution()(*SynchronizationTaskExecution) {
    if m == nil {
        return nil
    } else {
        return m.lastSuccessfulExecution
    }
}
func (m *SynchronizationStatus) GetLastSuccessfulExecutionWithExports()(*SynchronizationTaskExecution) {
    if m == nil {
        return nil
    } else {
        return m.lastSuccessfulExecutionWithExports
    }
}
func (m *SynchronizationStatus) GetProgress()([]SynchronizationProgress) {
    if m == nil {
        return nil
    } else {
        return m.progress
    }
}
func (m *SynchronizationStatus) GetQuarantine()(*SynchronizationQuarantine) {
    if m == nil {
        return nil
    } else {
        return m.quarantine
    }
}
func (m *SynchronizationStatus) GetSteadyStateFirstAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.steadyStateFirstAchievedTime
    }
}
func (m *SynchronizationStatus) GetSteadyStateLastAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.steadyStateLastAchievedTime
    }
}
func (m *SynchronizationStatus) GetSynchronizedEntryCountByType()([]StringKeyLongValuePair) {
    if m == nil {
        return nil
    } else {
        return m.synchronizedEntryCountByType
    }
}
func (m *SynchronizationStatus) GetTroubleshootingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingUrl
    }
}
func (m *SynchronizationStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationStatusCode)
        if err != nil {
            return err
        }
        cast := val.(SynchronizationStatusCode)
        m.SetCode(&cast)
        return nil
    }
    res["countSuccessiveCompleteFailures"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCountSuccessiveCompleteFailures(val)
        return nil
    }
    res["escrowsPruned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetEscrowsPruned(val)
        return nil
    }
    res["lastExecution"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationTaskExecution() })
        if err != nil {
            return err
        }
        m.SetLastExecution(val.(*SynchronizationTaskExecution))
        return nil
    }
    res["lastSuccessfulExecution"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationTaskExecution() })
        if err != nil {
            return err
        }
        m.SetLastSuccessfulExecution(val.(*SynchronizationTaskExecution))
        return nil
    }
    res["lastSuccessfulExecutionWithExports"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationTaskExecution() })
        if err != nil {
            return err
        }
        m.SetLastSuccessfulExecutionWithExports(val.(*SynchronizationTaskExecution))
        return nil
    }
    res["progress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationProgress() })
        if err != nil {
            return err
        }
        res := make([]SynchronizationProgress, len(val))
        for i, v := range val {
            res[i] = *(v.(*SynchronizationProgress))
        }
        m.SetProgress(res)
        return nil
    }
    res["quarantine"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationQuarantine() })
        if err != nil {
            return err
        }
        m.SetQuarantine(val.(*SynchronizationQuarantine))
        return nil
    }
    res["steadyStateFirstAchievedTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSteadyStateFirstAchievedTime(val)
        return nil
    }
    res["steadyStateLastAchievedTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetSteadyStateLastAchievedTime(val)
        return nil
    }
    res["synchronizedEntryCountByType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStringKeyLongValuePair() })
        if err != nil {
            return err
        }
        res := make([]StringKeyLongValuePair, len(val))
        for i, v := range val {
            res[i] = *(v.(*StringKeyLongValuePair))
        }
        m.SetSynchronizedEntryCountByType(res)
        return nil
    }
    res["troubleshootingUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTroubleshootingUrl(val)
        return nil
    }
    return res
}
func (m *SynchronizationStatus) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetCode() != nil {
        cast := m.GetCode().String()
        err := writer.WriteStringValue("code", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countSuccessiveCompleteFailures", m.GetCountSuccessiveCompleteFailures())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("escrowsPruned", m.GetEscrowsPruned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastExecution", m.GetLastExecution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastSuccessfulExecution", m.GetLastSuccessfulExecution())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastSuccessfulExecutionWithExports", m.GetLastSuccessfulExecutionWithExports())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProgress()))
        for i, v := range m.GetProgress() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("progress", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("quarantine", m.GetQuarantine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("steadyStateFirstAchievedTime", m.GetSteadyStateFirstAchievedTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("steadyStateLastAchievedTime", m.GetSteadyStateLastAchievedTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSynchronizedEntryCountByType()))
        for i, v := range m.GetSynchronizedEntryCountByType() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("synchronizedEntryCountByType", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("troubleshootingUrl", m.GetTroubleshootingUrl())
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
func (m *SynchronizationStatus) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SynchronizationStatus) SetCode(value *SynchronizationStatusCode)() {
    m.code = value
}
func (m *SynchronizationStatus) SetCountSuccessiveCompleteFailures(value *int64)() {
    m.countSuccessiveCompleteFailures = value
}
func (m *SynchronizationStatus) SetEscrowsPruned(value *bool)() {
    m.escrowsPruned = value
}
func (m *SynchronizationStatus) SetLastExecution(value *SynchronizationTaskExecution)() {
    m.lastExecution = value
}
func (m *SynchronizationStatus) SetLastSuccessfulExecution(value *SynchronizationTaskExecution)() {
    m.lastSuccessfulExecution = value
}
func (m *SynchronizationStatus) SetLastSuccessfulExecutionWithExports(value *SynchronizationTaskExecution)() {
    m.lastSuccessfulExecutionWithExports = value
}
func (m *SynchronizationStatus) SetProgress(value []SynchronizationProgress)() {
    m.progress = value
}
func (m *SynchronizationStatus) SetQuarantine(value *SynchronizationQuarantine)() {
    m.quarantine = value
}
func (m *SynchronizationStatus) SetSteadyStateFirstAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.steadyStateFirstAchievedTime = value
}
func (m *SynchronizationStatus) SetSteadyStateLastAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.steadyStateLastAchievedTime = value
}
func (m *SynchronizationStatus) SetSynchronizedEntryCountByType(value []StringKeyLongValuePair)() {
    m.synchronizedEntryCountByType = value
}
func (m *SynchronizationStatus) SetTroubleshootingUrl(value *string)() {
    m.troubleshootingUrl = value
}
