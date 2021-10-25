package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationTaskExecution struct {
    activityIdentifier *string;
    additionalData map[string]interface{};
    countEntitled *int64;
    countEntitledForProvisioning *int64;
    countEscrowed *int64;
    countEscrowedRaw *int64;
    countExported *int64;
    countExports *int64;
    countImported *int64;
    countImportedDeltas *int64;
    countImportedReferenceDeltas *int64;
    error *SynchronizationError;
    state *SynchronizationTaskExecutionResult;
    timeBegan *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    timeEnded *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
func NewSynchronizationTaskExecution()(*SynchronizationTaskExecution) {
    m := &SynchronizationTaskExecution{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SynchronizationTaskExecution) GetActivityIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityIdentifier
    }
}
func (m *SynchronizationTaskExecution) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SynchronizationTaskExecution) GetCountEntitled()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countEntitled
    }
}
func (m *SynchronizationTaskExecution) GetCountEntitledForProvisioning()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countEntitledForProvisioning
    }
}
func (m *SynchronizationTaskExecution) GetCountEscrowed()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countEscrowed
    }
}
func (m *SynchronizationTaskExecution) GetCountEscrowedRaw()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countEscrowedRaw
    }
}
func (m *SynchronizationTaskExecution) GetCountExported()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countExported
    }
}
func (m *SynchronizationTaskExecution) GetCountExports()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countExports
    }
}
func (m *SynchronizationTaskExecution) GetCountImported()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countImported
    }
}
func (m *SynchronizationTaskExecution) GetCountImportedDeltas()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countImportedDeltas
    }
}
func (m *SynchronizationTaskExecution) GetCountImportedReferenceDeltas()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.countImportedReferenceDeltas
    }
}
func (m *SynchronizationTaskExecution) GetError()(*SynchronizationError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
func (m *SynchronizationTaskExecution) GetState()(*SynchronizationTaskExecutionResult) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *SynchronizationTaskExecution) GetTimeBegan()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.timeBegan
    }
}
func (m *SynchronizationTaskExecution) GetTimeEnded()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.timeEnded
    }
}
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
func (m *SynchronizationTaskExecution) SetActivityIdentifier(value *string)() {
    m.activityIdentifier = value
}
func (m *SynchronizationTaskExecution) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SynchronizationTaskExecution) SetCountEntitled(value *int64)() {
    m.countEntitled = value
}
func (m *SynchronizationTaskExecution) SetCountEntitledForProvisioning(value *int64)() {
    m.countEntitledForProvisioning = value
}
func (m *SynchronizationTaskExecution) SetCountEscrowed(value *int64)() {
    m.countEscrowed = value
}
func (m *SynchronizationTaskExecution) SetCountEscrowedRaw(value *int64)() {
    m.countEscrowedRaw = value
}
func (m *SynchronizationTaskExecution) SetCountExported(value *int64)() {
    m.countExported = value
}
func (m *SynchronizationTaskExecution) SetCountExports(value *int64)() {
    m.countExports = value
}
func (m *SynchronizationTaskExecution) SetCountImported(value *int64)() {
    m.countImported = value
}
func (m *SynchronizationTaskExecution) SetCountImportedDeltas(value *int64)() {
    m.countImportedDeltas = value
}
func (m *SynchronizationTaskExecution) SetCountImportedReferenceDeltas(value *int64)() {
    m.countImportedReferenceDeltas = value
}
func (m *SynchronizationTaskExecution) SetError(value *SynchronizationError)() {
    m.error = value
}
func (m *SynchronizationTaskExecution) SetState(value *SynchronizationTaskExecutionResult)() {
    m.state = value
}
func (m *SynchronizationTaskExecution) SetTimeBegan(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timeBegan = value
}
func (m *SynchronizationTaskExecution) SetTimeEnded(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.timeEnded = value
}
