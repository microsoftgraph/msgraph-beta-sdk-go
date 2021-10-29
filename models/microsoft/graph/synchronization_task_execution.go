package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SynchronizationTaskExecution struct {
    // Identifier of the job run.
    activityIdentifier *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Count of processed entries that were assigned for this application.
    countEntitled *int64;
    // Count of processed entries that were assigned for provisioning.
    countEntitledForProvisioning *int64;
    // Count of entries that were escrowed (errors).
    countEscrowed *int64;
    // Count of entries that were escrowed, including system-generated escrows.
    countEscrowedRaw *int64;
    // Count of exported entries.
    countExported *int64;
    // Count of entries that were expected to be exported.
    countExports *int64;
    // Count of imported entries.
    countImported *int64;
    // Count of imported delta-changes.
    countImportedDeltas *int64;
    // Count of imported delta-changes pertaining to reference changes.
    countImportedReferenceDeltas *int64;
    // If an error was encountered, contains a synchronizationError object with details.
    error *SynchronizationError;
    // Code summarizing the result of this run. Possible values are: Succeeded, Failed, EntryLevelErrors.
    state *SynchronizationTaskExecutionResult;
    // Time when this job run began. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    timeBegan *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Time when this job run ended. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    timeEnded *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new synchronizationTaskExecution and sets the default values.
func NewSynchronizationTaskExecution()(*SynchronizationTaskExecution) {
    m := &SynchronizationTaskExecution{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the activityIdentifier property value. Identifier of the job run.
func (m *SynchronizationTaskExecution) GetActivityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityIdentifier
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationTaskExecution) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the countEntitled property value. Count of processed entries that were assigned for this application.
func (m *SynchronizationTaskExecution) GetCountEntitled()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countEntitled
    }
}
// Gets the countEntitledForProvisioning property value. Count of processed entries that were assigned for provisioning.
func (m *SynchronizationTaskExecution) GetCountEntitledForProvisioning()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countEntitledForProvisioning
    }
}
// Gets the countEscrowed property value. Count of entries that were escrowed (errors).
func (m *SynchronizationTaskExecution) GetCountEscrowed()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countEscrowed
    }
}
// Gets the countEscrowedRaw property value. Count of entries that were escrowed, including system-generated escrows.
func (m *SynchronizationTaskExecution) GetCountEscrowedRaw()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countEscrowedRaw
    }
}
// Gets the countExported property value. Count of exported entries.
func (m *SynchronizationTaskExecution) GetCountExported()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countExported
    }
}
// Gets the countExports property value. Count of entries that were expected to be exported.
func (m *SynchronizationTaskExecution) GetCountExports()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countExports
    }
}
// Gets the countImported property value. Count of imported entries.
func (m *SynchronizationTaskExecution) GetCountImported()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countImported
    }
}
// Gets the countImportedDeltas property value. Count of imported delta-changes.
func (m *SynchronizationTaskExecution) GetCountImportedDeltas()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countImportedDeltas
    }
}
// Gets the countImportedReferenceDeltas property value. Count of imported delta-changes pertaining to reference changes.
func (m *SynchronizationTaskExecution) GetCountImportedReferenceDeltas()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countImportedReferenceDeltas
    }
}
// Gets the error property value. If an error was encountered, contains a synchronizationError object with details.
func (m *SynchronizationTaskExecution) GetError()(*SynchronizationError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// Gets the state property value. Code summarizing the result of this run. Possible values are: Succeeded, Failed, EntryLevelErrors.
func (m *SynchronizationTaskExecution) GetState()(*SynchronizationTaskExecutionResult) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// Gets the timeBegan property value. Time when this job run began. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationTaskExecution) GetTimeBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.timeBegan
    }
}
// Gets the timeEnded property value. Time when this job run ended. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *SynchronizationTaskExecution) GetTimeEnded()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.timeEnded
    }
}
// The deserialization information for the current model
func (m *SynchronizationTaskExecution) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["activityIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivityIdentifier(val)
        return nil
    }
    res["countEntitled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCountEntitled(val)
        return nil
    }
    res["countEntitledForProvisioning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCountEntitledForProvisioning(val)
        return nil
    }
    res["countEscrowed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCountEscrowed(val)
        return nil
    }
    res["countEscrowedRaw"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCountEscrowedRaw(val)
        return nil
    }
    res["countExported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCountExported(val)
        return nil
    }
    res["countExports"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCountExports(val)
        return nil
    }
    res["countImported"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCountImported(val)
        return nil
    }
    res["countImportedDeltas"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCountImportedDeltas(val)
        return nil
    }
    res["countImportedReferenceDeltas"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetCountImportedReferenceDeltas(val)
        return nil
    }
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationError() })
        if err != nil {
            return err
        }
        m.SetError(val.(*SynchronizationError))
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSynchronizationTaskExecutionResult)
        if err != nil {
            return err
        }
        cast := val.(SynchronizationTaskExecutionResult)
        m.SetState(&cast)
        return nil
    }
    res["timeBegan"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetTimeBegan(val)
        return nil
    }
    res["timeEnded"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetTimeEnded(val)
        return nil
    }
    return res
}
func (m *SynchronizationTaskExecution) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SynchronizationTaskExecution) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    if m.GetState() != nil {
        cast := m.GetState().String()
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
// Sets the activityIdentifier property value. Identifier of the job run.
// Parameters:
//  - value : Value to set for the activityIdentifier property.
func (m *SynchronizationTaskExecution) SetActivityIdentifier(value *string)() {
    m.activityIdentifier = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SynchronizationTaskExecution) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the countEntitled property value. Count of processed entries that were assigned for this application.
// Parameters:
//  - value : Value to set for the countEntitled property.
func (m *SynchronizationTaskExecution) SetCountEntitled(value *int64)() {
    m.countEntitled = value
}
// Sets the countEntitledForProvisioning property value. Count of processed entries that were assigned for provisioning.
// Parameters:
//  - value : Value to set for the countEntitledForProvisioning property.
func (m *SynchronizationTaskExecution) SetCountEntitledForProvisioning(value *int64)() {
    m.countEntitledForProvisioning = value
}
// Sets the countEscrowed property value. Count of entries that were escrowed (errors).
// Parameters:
//  - value : Value to set for the countEscrowed property.
func (m *SynchronizationTaskExecution) SetCountEscrowed(value *int64)() {
    m.countEscrowed = value
}
// Sets the countEscrowedRaw property value. Count of entries that were escrowed, including system-generated escrows.
// Parameters:
//  - value : Value to set for the countEscrowedRaw property.
func (m *SynchronizationTaskExecution) SetCountEscrowedRaw(value *int64)() {
    m.countEscrowedRaw = value
}
// Sets the countExported property value. Count of exported entries.
// Parameters:
//  - value : Value to set for the countExported property.
func (m *SynchronizationTaskExecution) SetCountExported(value *int64)() {
    m.countExported = value
}
// Sets the countExports property value. Count of entries that were expected to be exported.
// Parameters:
//  - value : Value to set for the countExports property.
func (m *SynchronizationTaskExecution) SetCountExports(value *int64)() {
    m.countExports = value
}
// Sets the countImported property value. Count of imported entries.
// Parameters:
//  - value : Value to set for the countImported property.
func (m *SynchronizationTaskExecution) SetCountImported(value *int64)() {
    m.countImported = value
}
// Sets the countImportedDeltas property value. Count of imported delta-changes.
// Parameters:
//  - value : Value to set for the countImportedDeltas property.
func (m *SynchronizationTaskExecution) SetCountImportedDeltas(value *int64)() {
    m.countImportedDeltas = value
}
// Sets the countImportedReferenceDeltas property value. Count of imported delta-changes pertaining to reference changes.
// Parameters:
//  - value : Value to set for the countImportedReferenceDeltas property.
func (m *SynchronizationTaskExecution) SetCountImportedReferenceDeltas(value *int64)() {
    m.countImportedReferenceDeltas = value
}
// Sets the error property value. If an error was encountered, contains a synchronizationError object with details.
// Parameters:
//  - value : Value to set for the error property.
func (m *SynchronizationTaskExecution) SetError(value *SynchronizationError)() {
    m.error = value
}
// Sets the state property value. Code summarizing the result of this run. Possible values are: Succeeded, Failed, EntryLevelErrors.
// Parameters:
//  - value : Value to set for the state property.
func (m *SynchronizationTaskExecution) SetState(value *SynchronizationTaskExecutionResult)() {
    m.state = value
}
// Sets the timeBegan property value. Time when this job run began. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the timeBegan property.
func (m *SynchronizationTaskExecution) SetTimeBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timeBegan = value
}
// Sets the timeEnded property value. Time when this job run ended. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the timeEnded property.
func (m *SynchronizationTaskExecution) SetTimeEnded(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timeEnded = value
}
