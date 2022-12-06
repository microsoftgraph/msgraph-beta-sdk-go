package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationStatus 
type SynchronizationStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The code property
    code *SynchronizationStatusCode
    // Number of consecutive times this job failed.
    countSuccessiveCompleteFailures *int64
    // true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
    escrowsPruned *bool
    // Details of the last execution of the job.
    lastExecution SynchronizationTaskExecutionable
    // Details of the last execution of this job, which didn't have any errors.
    lastSuccessfulExecution SynchronizationTaskExecutionable
    // Details of the last execution of the job, which exported objects into the target directory.
    lastSuccessfulExecutionWithExports SynchronizationTaskExecutionable
    // The OdataType property
    odataType *string
    // Details of the progress of a job toward completion.
    progress []SynchronizationProgressable
    // If job is in quarantine, quarantine details.
    quarantine SynchronizationQuarantineable
    // The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    steadyStateFirstAchievedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    steadyStateLastAchievedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Count of synchronized objects, listed by object type.
    synchronizedEntryCountByType []StringKeyLongValuePairable
    // In the event of an error, the URL with the troubleshooting steps for the issue.
    troubleshootingUrl *string
}
// NewSynchronizationStatus instantiates a new synchronizationStatus and sets the default values.
func NewSynchronizationStatus()(*SynchronizationStatus) {
    m := &SynchronizationStatus{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSynchronizationStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationStatus) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
func (m *SynchronizationStatus) GetCode()(*SynchronizationStatusCode) {
    return m.code
}
// GetCountSuccessiveCompleteFailures gets the countSuccessiveCompleteFailures property value. Number of consecutive times this job failed.
func (m *SynchronizationStatus) GetCountSuccessiveCompleteFailures()(*int64) {
    return m.countSuccessiveCompleteFailures
}
// GetEscrowsPruned gets the escrowsPruned property value. true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
func (m *SynchronizationStatus) GetEscrowsPruned()(*bool) {
    return m.escrowsPruned
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSynchronizationStatusCode , m.SetCode)
    res["countSuccessiveCompleteFailures"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetCountSuccessiveCompleteFailures)
    res["escrowsPruned"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetEscrowsPruned)
    res["lastExecution"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSynchronizationTaskExecutionFromDiscriminatorValue , m.SetLastExecution)
    res["lastSuccessfulExecution"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSynchronizationTaskExecutionFromDiscriminatorValue , m.SetLastSuccessfulExecution)
    res["lastSuccessfulExecutionWithExports"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSynchronizationTaskExecutionFromDiscriminatorValue , m.SetLastSuccessfulExecutionWithExports)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["progress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateSynchronizationProgressFromDiscriminatorValue , m.SetProgress)
    res["quarantine"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSynchronizationQuarantineFromDiscriminatorValue , m.SetQuarantine)
    res["steadyStateFirstAchievedTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetSteadyStateFirstAchievedTime)
    res["steadyStateLastAchievedTime"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetSteadyStateLastAchievedTime)
    res["synchronizedEntryCountByType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateStringKeyLongValuePairFromDiscriminatorValue , m.SetSynchronizedEntryCountByType)
    res["troubleshootingUrl"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetTroubleshootingUrl)
    return res
}
// GetLastExecution gets the lastExecution property value. Details of the last execution of the job.
func (m *SynchronizationStatus) GetLastExecution()(SynchronizationTaskExecutionable) {
    return m.lastExecution
}
// GetLastSuccessfulExecution gets the lastSuccessfulExecution property value. Details of the last execution of this job, which didn't have any errors.
func (m *SynchronizationStatus) GetLastSuccessfulExecution()(SynchronizationTaskExecutionable) {
    return m.lastSuccessfulExecution
}
// GetLastSuccessfulExecutionWithExports gets the lastSuccessfulExecutionWithExports property value. Details of the last execution of the job, which exported objects into the target directory.
func (m *SynchronizationStatus) GetLastSuccessfulExecutionWithExports()(SynchronizationTaskExecutionable) {
    return m.lastSuccessfulExecutionWithExports
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SynchronizationStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetProgress gets the progress property value. Details of the progress of a job toward completion.
func (m *SynchronizationStatus) GetProgress()([]SynchronizationProgressable) {
    return m.progress
}
// GetQuarantine gets the quarantine property value. If job is in quarantine, quarantine details.
func (m *SynchronizationStatus) GetQuarantine()(SynchronizationQuarantineable) {
    return m.quarantine
}
// GetSteadyStateFirstAchievedTime gets the steadyStateFirstAchievedTime property value. The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) GetSteadyStateFirstAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.steadyStateFirstAchievedTime
}
// GetSteadyStateLastAchievedTime gets the steadyStateLastAchievedTime property value. The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) GetSteadyStateLastAchievedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.steadyStateLastAchievedTime
}
// GetSynchronizedEntryCountByType gets the synchronizedEntryCountByType property value. Count of synchronized objects, listed by object type.
func (m *SynchronizationStatus) GetSynchronizedEntryCountByType()([]StringKeyLongValuePairable) {
    return m.synchronizedEntryCountByType
}
// GetTroubleshootingUrl gets the troubleshootingUrl property value. In the event of an error, the URL with the troubleshooting steps for the issue.
func (m *SynchronizationStatus) GetTroubleshootingUrl()(*string) {
    return m.troubleshootingUrl
}
// Serialize serializes information the current object
func (m *SynchronizationStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetProgress() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetProgress())
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
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetSynchronizedEntryCountByType())
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
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *SynchronizationStatus) SetCode(value *SynchronizationStatusCode)() {
    m.code = value
}
// SetCountSuccessiveCompleteFailures sets the countSuccessiveCompleteFailures property value. Number of consecutive times this job failed.
func (m *SynchronizationStatus) SetCountSuccessiveCompleteFailures(value *int64)() {
    m.countSuccessiveCompleteFailures = value
}
// SetEscrowsPruned sets the escrowsPruned property value. true if the job's escrows (object-level errors) were pruned during initial synchronization. Escrows can be pruned if during the initial synchronization, you reach the threshold of errors that would normally put the job in quarantine. Instead of going into quarantine, the synchronization process clears the job's errors and continues until the initial synchronization is completed. When the initial synchronization is completed, the job will pause and wait for the customer to clean up the errors.
func (m *SynchronizationStatus) SetEscrowsPruned(value *bool)() {
    m.escrowsPruned = value
}
// SetLastExecution sets the lastExecution property value. Details of the last execution of the job.
func (m *SynchronizationStatus) SetLastExecution(value SynchronizationTaskExecutionable)() {
    m.lastExecution = value
}
// SetLastSuccessfulExecution sets the lastSuccessfulExecution property value. Details of the last execution of this job, which didn't have any errors.
func (m *SynchronizationStatus) SetLastSuccessfulExecution(value SynchronizationTaskExecutionable)() {
    m.lastSuccessfulExecution = value
}
// SetLastSuccessfulExecutionWithExports sets the lastSuccessfulExecutionWithExports property value. Details of the last execution of the job, which exported objects into the target directory.
func (m *SynchronizationStatus) SetLastSuccessfulExecutionWithExports(value SynchronizationTaskExecutionable)() {
    m.lastSuccessfulExecutionWithExports = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SynchronizationStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProgress sets the progress property value. Details of the progress of a job toward completion.
func (m *SynchronizationStatus) SetProgress(value []SynchronizationProgressable)() {
    m.progress = value
}
// SetQuarantine sets the quarantine property value. If job is in quarantine, quarantine details.
func (m *SynchronizationStatus) SetQuarantine(value SynchronizationQuarantineable)() {
    m.quarantine = value
}
// SetSteadyStateFirstAchievedTime sets the steadyStateFirstAchievedTime property value. The time when steady state (no more changes to the process) was first achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) SetSteadyStateFirstAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.steadyStateFirstAchievedTime = value
}
// SetSteadyStateLastAchievedTime sets the steadyStateLastAchievedTime property value. The time when steady state (no more changes to the process) was last achieved. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationStatus) SetSteadyStateLastAchievedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.steadyStateLastAchievedTime = value
}
// SetSynchronizedEntryCountByType sets the synchronizedEntryCountByType property value. Count of synchronized objects, listed by object type.
func (m *SynchronizationStatus) SetSynchronizedEntryCountByType(value []StringKeyLongValuePairable)() {
    m.synchronizedEntryCountByType = value
}
// SetTroubleshootingUrl sets the troubleshootingUrl property value. In the event of an error, the URL with the troubleshooting steps for the issue.
func (m *SynchronizationStatus) SetTroubleshootingUrl(value *string)() {
    m.troubleshootingUrl = value
}
