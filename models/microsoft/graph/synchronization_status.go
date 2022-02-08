package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SynchronizationStatus 
type SynchronizationStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // High-level status code of the synchronization job. Possible values are: NotConfigured, NotRun, Active, Paused, Quarantine.
    code *SynchronizationStatusCode;
    // Number of consecutive times this job failed.
    countSuccessiveCompleteFailures *int64;
    // true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
    escrowsPruned *bool;
    // Details of the last execution of the job.
    lastExecution *SynchronizationTaskExecution;
    // Details of the last execution of this job, which didn't have any errors.
    lastSuccessfulExecution *SynchronizationTaskExecution;
    // Details of the last execution of the job, which exported objects into the target directory.
    lastSuccessfulExecutionWithExports *SynchronizationTaskExecution;
    // Details of the progress of a job toward completion.
    progress []SynchronizationProgress;
    // If job is in quarantine, quarantine details.
    quarantine *SynchronizationQuarantine;
    // The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    steadyStateFirstAchievedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    steadyStateLastAchievedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Count of synchronized objects, listed by object type.
    synchronizedEntryCountByType []StringKeyLongValuePair;
    // In the event of an error, the URL with the troubleshooting steps for the issue.
    troubleshootingUrl *string;
}
// NewSynchronizationStatus instantiates a new synchronizationStatus and sets the default values.
func NewSynchronizationStatus()(*SynchronizationStatus) {
    m := &SynchronizationStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationStatus) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCode gets the code property value. High-level status code of the synchronization job. Possible values are: NotConfigured, NotRun, Active, Paused, Quarantine.
func (m *SynchronizationStatus) GetCode()(*SynchronizationStatusCode) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetCountSuccessiveCompleteFailures gets the countSuccessiveCompleteFailures property value. Number of consecutive times this job failed.
func (m *SynchronizationStatus) GetCountSuccessiveCompleteFailures()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countSuccessiveCompleteFailures
    }
}
// GetEscrowsPruned gets the escrowsPruned property value. true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
func (m *SynchronizationStatus) GetEscrowsPruned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.escrowsPruned
    }
}
// GetLastExecution gets the lastExecution property value. Details of the last execution of the job.
func (m *SynchronizationStatus) GetLastExecution()(*SynchronizationTaskExecution) {
    if m == nil {
        return nil
    } else {
        return m.lastExecution
    }
}
// GetLastSuccessfulExecution gets the lastSuccessfulExecution property value. Details of the last execution of this job, which didn't have any errors.
func (m *SynchronizationStatus) GetLastSuccessfulExecution()(*SynchronizationTaskExecution) {
    if m == nil {
        return nil
    } else {
        return m.lastSuccessfulExecution
    }
}
// GetLastSuccessfulExecutionWithExports gets the lastSuccessfulExecutionWithExports property value. Details of the last execution of the job, which exported objects into the target directory.
func (m *SynchronizationStatus) GetLastSuccessfulExecutionWithExports()(*SynchronizationTaskExecution) {
    if m == nil {
        return nil
    } else {
        return m.lastSuccessfulExecutionWithExports
    }
}
// GetProgress gets the progress property value. Details of the progress of a job toward completion.
func (m *SynchronizationStatus) GetProgress()([]SynchronizationProgress) {
    if m == nil {
        return nil
    } else {
        return m.progress
    }
}
// GetQuarantine gets the quarantine property value. If job is in quarantine, quarantine details.
func (m *SynchronizationStatus) GetQuarantine()(*SynchronizationQuarantine) {
    if m == nil {
        return nil
    } else {
        return m.quarantine
    }
}
// GetSteadyStateFirstAchievedTime gets the steadyStateFirstAchievedTime property value. The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) GetSteadyStateFirstAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.steadyStateFirstAchievedTime
    }
}
// GetSteadyStateLastAchievedTime gets the steadyStateLastAchievedTime property value. The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) GetSteadyStateLastAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.steadyStateLastAchievedTime
    }
}
// GetSynchronizedEntryCountByType gets the synchronizedEntryCountByType property value. Count of synchronized objects, listed by object type.
func (m *SynchronizationStatus) GetSynchronizedEntryCountByType()([]StringKeyLongValuePair) {
    if m == nil {
        return nil
    } else {
        return m.synchronizedEntryCountByType
    }
}
// GetTroubleshootingUrl gets the troubleshootingUrl property value. In the event of an error, the URL with the troubleshooting steps for the issue.
func (m *SynchronizationStatus) GetTroubleshootingUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationStatusCode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val.(*SynchronizationStatusCode))
        }
        return nil
    }
    res["countSuccessiveCompleteFailures"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountSuccessiveCompleteFailures(val)
        }
        return nil
    }
    res["escrowsPruned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEscrowsPruned(val)
        }
        return nil
    }
    res["lastExecution"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationTaskExecution() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastExecution(val.(*SynchronizationTaskExecution))
        }
        return nil
    }
    res["lastSuccessfulExecution"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationTaskExecution() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulExecution(val.(*SynchronizationTaskExecution))
        }
        return nil
    }
    res["lastSuccessfulExecutionWithExports"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationTaskExecution() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSuccessfulExecutionWithExports(val.(*SynchronizationTaskExecution))
        }
        return nil
    }
    res["progress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationProgress() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SynchronizationProgress, len(val))
            for i, v := range val {
                res[i] = *(v.(*SynchronizationProgress))
            }
            m.SetProgress(res)
        }
        return nil
    }
    res["quarantine"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationQuarantine() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuarantine(val.(*SynchronizationQuarantine))
        }
        return nil
    }
    res["steadyStateFirstAchievedTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSteadyStateFirstAchievedTime(val)
        }
        return nil
    }
    res["steadyStateLastAchievedTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSteadyStateLastAchievedTime(val)
        }
        return nil
    }
    res["synchronizedEntryCountByType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewStringKeyLongValuePair() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]StringKeyLongValuePair, len(val))
            for i, v := range val {
                res[i] = *(v.(*StringKeyLongValuePair))
            }
            m.SetSynchronizedEntryCountByType(res)
        }
        return nil
    }
    res["troubleshootingUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTroubleshootingUrl(val)
        }
        return nil
    }
    return res
}
func (m *SynchronizationStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SynchronizationStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetCode() != nil {
        cast := (*m.GetCode()).String()
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
    if m.GetProgress() != nil {
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
    if m.GetSynchronizedEntryCountByType() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationStatus) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCode sets the code property value. High-level status code of the synchronization job. Possible values are: NotConfigured, NotRun, Active, Paused, Quarantine.
func (m *SynchronizationStatus) SetCode(value *SynchronizationStatusCode)() {
    if m != nil {
        m.code = value
    }
}
// SetCountSuccessiveCompleteFailures sets the countSuccessiveCompleteFailures property value. Number of consecutive times this job failed.
func (m *SynchronizationStatus) SetCountSuccessiveCompleteFailures(value *int64)() {
    if m != nil {
        m.countSuccessiveCompleteFailures = value
    }
}
// SetEscrowsPruned sets the escrowsPruned property value. true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
func (m *SynchronizationStatus) SetEscrowsPruned(value *bool)() {
    if m != nil {
        m.escrowsPruned = value
    }
}
// SetLastExecution sets the lastExecution property value. Details of the last execution of the job.
func (m *SynchronizationStatus) SetLastExecution(value *SynchronizationTaskExecution)() {
    if m != nil {
        m.lastExecution = value
    }
}
// SetLastSuccessfulExecution sets the lastSuccessfulExecution property value. Details of the last execution of this job, which didn't have any errors.
func (m *SynchronizationStatus) SetLastSuccessfulExecution(value *SynchronizationTaskExecution)() {
    if m != nil {
        m.lastSuccessfulExecution = value
    }
}
// SetLastSuccessfulExecutionWithExports sets the lastSuccessfulExecutionWithExports property value. Details of the last execution of the job, which exported objects into the target directory.
func (m *SynchronizationStatus) SetLastSuccessfulExecutionWithExports(value *SynchronizationTaskExecution)() {
    if m != nil {
        m.lastSuccessfulExecutionWithExports = value
    }
}
// SetProgress sets the progress property value. Details of the progress of a job toward completion.
func (m *SynchronizationStatus) SetProgress(value []SynchronizationProgress)() {
    if m != nil {
        m.progress = value
    }
}
// SetQuarantine sets the quarantine property value. If job is in quarantine, quarantine details.
func (m *SynchronizationStatus) SetQuarantine(value *SynchronizationQuarantine)() {
    if m != nil {
        m.quarantine = value
    }
}
// SetSteadyStateFirstAchievedTime sets the steadyStateFirstAchievedTime property value. The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) SetSteadyStateFirstAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.steadyStateFirstAchievedTime = value
    }
}
// SetSteadyStateLastAchievedTime sets the steadyStateLastAchievedTime property value. The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) SetSteadyStateLastAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.steadyStateLastAchievedTime = value
    }
}
// SetSynchronizedEntryCountByType sets the synchronizedEntryCountByType property value. Count of synchronized objects, listed by object type.
func (m *SynchronizationStatus) SetSynchronizedEntryCountByType(value []StringKeyLongValuePair)() {
    if m != nil {
        m.synchronizedEntryCountByType = value
    }
}
// SetTroubleshootingUrl sets the troubleshootingUrl property value. In the event of an error, the URL with the troubleshooting steps for the issue.
func (m *SynchronizationStatus) SetTroubleshootingUrl(value *string)() {
    if m != nil {
        m.troubleshootingUrl = value
    }
}
