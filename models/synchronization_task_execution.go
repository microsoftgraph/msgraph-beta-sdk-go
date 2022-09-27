package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationTaskExecution 
type SynchronizationTaskExecution struct {
    // Identifier of the job run.
    activityIdentifier *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Count of processed entries that were assigned for this application.
    countEntitled *int64
    // Count of processed entries that were assigned for provisioning.
    countEntitledForProvisioning *int64
    // Count of entries that were escrowed (errors).
    countEscrowed *int64
    // Count of entries that were escrowed, including system-generated escrows.
    countEscrowedRaw *int64
    // Count of exported entries.
    countExported *int64
    // Count of entries that were expected to be exported.
    countExports *int64
    // Count of imported entries.
    countImported *int64
    // Count of imported delta-changes.
    countImportedDeltas *int64
    // Count of imported delta-changes pertaining to reference changes.
    countImportedReferenceDeltas *int64
    // If an error was encountered, contains a synchronizationError object with details.
    error SynchronizationErrorable
    // The OdataType property
    odataType *string
    // The state property
    state *SynchronizationTaskExecutionResult
    // Time when this job run began. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    timeBegan *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Time when this job run ended. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    timeEnded *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewSynchronizationTaskExecution instantiates a new synchronizationTaskExecution and sets the default values.
func NewSynchronizationTaskExecution()(*SynchronizationTaskExecution) {
    m := &SynchronizationTaskExecution{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.synchronizationTaskExecution";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateSynchronizationTaskExecutionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationTaskExecutionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationTaskExecution(), nil
}
// GetActivityIdentifier gets the activityIdentifier property value. Identifier of the job run.
func (m *SynchronizationTaskExecution) GetActivityIdentifier()(*string) {
    return m.activityIdentifier
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationTaskExecution) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCountEntitled gets the countEntitled property value. Count of processed entries that were assigned for this application.
func (m *SynchronizationTaskExecution) GetCountEntitled()(*int64) {
    return m.countEntitled
}
// GetCountEntitledForProvisioning gets the countEntitledForProvisioning property value. Count of processed entries that were assigned for provisioning.
func (m *SynchronizationTaskExecution) GetCountEntitledForProvisioning()(*int64) {
    return m.countEntitledForProvisioning
}
// GetCountEscrowed gets the countEscrowed property value. Count of entries that were escrowed (errors).
func (m *SynchronizationTaskExecution) GetCountEscrowed()(*int64) {
    return m.countEscrowed
}
// GetCountEscrowedRaw gets the countEscrowedRaw property value. Count of entries that were escrowed, including system-generated escrows.
func (m *SynchronizationTaskExecution) GetCountEscrowedRaw()(*int64) {
    return m.countEscrowedRaw
}
// GetCountExported gets the countExported property value. Count of exported entries.
func (m *SynchronizationTaskExecution) GetCountExported()(*int64) {
    return m.countExported
}
// GetCountExports gets the countExports property value. Count of entries that were expected to be exported.
func (m *SynchronizationTaskExecution) GetCountExports()(*int64) {
    return m.countExports
}
// GetCountImported gets the countImported property value. Count of imported entries.
func (m *SynchronizationTaskExecution) GetCountImported()(*int64) {
    return m.countImported
}
// GetCountImportedDeltas gets the countImportedDeltas property value. Count of imported delta-changes.
func (m *SynchronizationTaskExecution) GetCountImportedDeltas()(*int64) {
    return m.countImportedDeltas
}
// GetCountImportedReferenceDeltas gets the countImportedReferenceDeltas property value. Count of imported delta-changes pertaining to reference changes.
func (m *SynchronizationTaskExecution) GetCountImportedReferenceDeltas()(*int64) {
    return m.countImportedReferenceDeltas
}
// GetError gets the error property value. If an error was encountered, contains a synchronizationError object with details.
func (m *SynchronizationTaskExecution) GetError()(SynchronizationErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationTaskExecution) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activityIdentifier"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetActivityIdentifier)
    res["countEntitled"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetCountEntitled)
    res["countEntitledForProvisioning"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetCountEntitledForProvisioning)
    res["countEscrowed"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetCountEscrowed)
    res["countEscrowedRaw"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetCountEscrowedRaw)
    res["countExported"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetCountExported)
    res["countExports"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetCountExports)
    res["countImported"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetCountImported)
    res["countImportedDeltas"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetCountImportedDeltas)
    res["countImportedReferenceDeltas"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt64Value(m.SetCountImportedReferenceDeltas)
    res["error"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSynchronizationErrorFromDiscriminatorValue , m.SetError)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["state"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseSynchronizationTaskExecutionResult , m.SetState)
    res["timeBegan"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetTimeBegan)
    res["timeEnded"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetTimeValue(m.SetTimeEnded)
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SynchronizationTaskExecution) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. The state property
func (m *SynchronizationTaskExecution) GetState()(*SynchronizationTaskExecutionResult) {
    return m.state
}
// GetTimeBegan gets the timeBegan property value. Time when this job run began. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationTaskExecution) GetTimeBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.timeBegan
}
// GetTimeEnded gets the timeEnded property value. Time when this job run ended. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationTaskExecution) GetTimeEnded()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.timeEnded
}
// Serialize serializes information the current object
func (m *SynchronizationTaskExecution) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityIdentifier", m.GetActivityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countEntitled", m.GetCountEntitled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countEntitledForProvisioning", m.GetCountEntitledForProvisioning())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countEscrowed", m.GetCountEscrowed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countEscrowedRaw", m.GetCountEscrowedRaw())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countExported", m.GetCountExported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countExports", m.GetCountExports())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countImported", m.GetCountImported())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countImportedDeltas", m.GetCountImportedDeltas())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("countImportedReferenceDeltas", m.GetCountImportedReferenceDeltas())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("timeBegan", m.GetTimeBegan())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("timeEnded", m.GetTimeEnded())
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
// SetActivityIdentifier sets the activityIdentifier property value. Identifier of the job run.
func (m *SynchronizationTaskExecution) SetActivityIdentifier(value *string)() {
    m.activityIdentifier = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationTaskExecution) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCountEntitled sets the countEntitled property value. Count of processed entries that were assigned for this application.
func (m *SynchronizationTaskExecution) SetCountEntitled(value *int64)() {
    m.countEntitled = value
}
// SetCountEntitledForProvisioning sets the countEntitledForProvisioning property value. Count of processed entries that were assigned for provisioning.
func (m *SynchronizationTaskExecution) SetCountEntitledForProvisioning(value *int64)() {
    m.countEntitledForProvisioning = value
}
// SetCountEscrowed sets the countEscrowed property value. Count of entries that were escrowed (errors).
func (m *SynchronizationTaskExecution) SetCountEscrowed(value *int64)() {
    m.countEscrowed = value
}
// SetCountEscrowedRaw sets the countEscrowedRaw property value. Count of entries that were escrowed, including system-generated escrows.
func (m *SynchronizationTaskExecution) SetCountEscrowedRaw(value *int64)() {
    m.countEscrowedRaw = value
}
// SetCountExported sets the countExported property value. Count of exported entries.
func (m *SynchronizationTaskExecution) SetCountExported(value *int64)() {
    m.countExported = value
}
// SetCountExports sets the countExports property value. Count of entries that were expected to be exported.
func (m *SynchronizationTaskExecution) SetCountExports(value *int64)() {
    m.countExports = value
}
// SetCountImported sets the countImported property value. Count of imported entries.
func (m *SynchronizationTaskExecution) SetCountImported(value *int64)() {
    m.countImported = value
}
// SetCountImportedDeltas sets the countImportedDeltas property value. Count of imported delta-changes.
func (m *SynchronizationTaskExecution) SetCountImportedDeltas(value *int64)() {
    m.countImportedDeltas = value
}
// SetCountImportedReferenceDeltas sets the countImportedReferenceDeltas property value. Count of imported delta-changes pertaining to reference changes.
func (m *SynchronizationTaskExecution) SetCountImportedReferenceDeltas(value *int64)() {
    m.countImportedReferenceDeltas = value
}
// SetError sets the error property value. If an error was encountered, contains a synchronizationError object with details.
func (m *SynchronizationTaskExecution) SetError(value SynchronizationErrorable)() {
    m.error = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SynchronizationTaskExecution) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. The state property
func (m *SynchronizationTaskExecution) SetState(value *SynchronizationTaskExecutionResult)() {
    m.state = value
}
// SetTimeBegan sets the timeBegan property value. Time when this job run began. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationTaskExecution) SetTimeBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timeBegan = value
}
// SetTimeEnded sets the timeEnded property value. Time when this job run ended. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationTaskExecution) SetTimeEnded(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timeEnded = value
}
