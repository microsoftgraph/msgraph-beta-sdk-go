package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EducationSynchronizationError struct {
    Entity
    entryType *string;
    errorCode *string;
    errorMessage *string;
    joiningValue *string;
    recordedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    reportableIdentifier *string;
}
func NewEducationSynchronizationError()(*EducationSynchronizationError) {
    m := &EducationSynchronizationError{
        Entity: *NewEntity(),
    }
    return m
}
func (m *EducationSynchronizationError) GetEntryType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.entryType
    }
}
func (m *EducationSynchronizationError) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
func (m *EducationSynchronizationError) GetErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorMessage
    }
}
func (m *EducationSynchronizationError) GetJoiningValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joiningValue
    }
}
func (m *EducationSynchronizationError) GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.recordedDateTime
    }
}
func (m *EducationSynchronizationError) GetReportableIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportableIdentifier
    }
}
func (m *EducationSynchronizationError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["entryType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEntryType(val)
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorCode(val)
        return nil
    }
    res["errorMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetErrorMessage(val)
        return nil
    }
    res["joiningValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJoiningValue(val)
        return nil
    }
    res["recordedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetRecordedDateTime(val)
        return nil
    }
    res["reportableIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReportableIdentifier(val)
        return nil
    }
    return res
}
func (m *EducationSynchronizationError) IsNil()(bool) {
    return m == nil
}
func (m *EducationSynchronizationError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("entryType", m.GetEntryType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorMessage", m.GetErrorMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joiningValue", m.GetJoiningValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("recordedDateTime", m.GetRecordedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("reportableIdentifier", m.GetReportableIdentifier())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *EducationSynchronizationError) SetEntryType(value *string)() {
    m.entryType = value
}
func (m *EducationSynchronizationError) SetErrorCode(value *string)() {
    m.errorCode = value
}
func (m *EducationSynchronizationError) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
func (m *EducationSynchronizationError) SetJoiningValue(value *string)() {
    m.joiningValue = value
}
func (m *EducationSynchronizationError) SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.recordedDateTime = value
}
func (m *EducationSynchronizationError) SetReportableIdentifier(value *string)() {
    m.reportableIdentifier = value
}
