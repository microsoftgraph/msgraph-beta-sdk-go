package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationSynchronizationError 
type EducationSynchronizationError struct {
    Entity
    // Represents the sync entity (school, section, student, teacher).
    entryType *string;
    // Represents the error code for this error.
    errorCode *string;
    // Contains a description of the error.
    errorMessage *string;
    // The unique identifier for the entry.
    joiningValue *string;
    // The time of occurrence of this error.
    recordedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The identifier of this error entry.
    reportableIdentifier *string;
}
// NewEducationSynchronizationError instantiates a new educationSynchronizationError and sets the default values.
func NewEducationSynchronizationError()(*EducationSynchronizationError) {
    m := &EducationSynchronizationError{
        Entity: *NewEntity(),
    }
    return m
}
// GetEntryType gets the entryType property value. Represents the sync entity (school, section, student, teacher).
func (m *EducationSynchronizationError) GetEntryType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.entryType
    }
}
// GetErrorCode gets the errorCode property value. Represents the error code for this error.
func (m *EducationSynchronizationError) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetErrorMessage gets the errorMessage property value. Contains a description of the error.
func (m *EducationSynchronizationError) GetErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorMessage
    }
}
// GetJoiningValue gets the joiningValue property value. The unique identifier for the entry.
func (m *EducationSynchronizationError) GetJoiningValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joiningValue
    }
}
// GetRecordedDateTime gets the recordedDateTime property value. The time of occurrence of this error.
func (m *EducationSynchronizationError) GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.recordedDateTime
    }
}
// GetReportableIdentifier gets the reportableIdentifier property value. The identifier of this error entry.
func (m *EducationSynchronizationError) GetReportableIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reportableIdentifier
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["entryType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntryType(val)
        }
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["errorMessage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorMessage(val)
        }
        return nil
    }
    res["joiningValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoiningValue(val)
        }
        return nil
    }
    res["recordedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordedDateTime(val)
        }
        return nil
    }
    res["reportableIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReportableIdentifier(val)
        }
        return nil
    }
    return res
}
func (m *EducationSynchronizationError) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetEntryType sets the entryType property value. Represents the sync entity (school, section, student, teacher).
func (m *EducationSynchronizationError) SetEntryType(value *string)() {
    if m != nil {
        m.entryType = value
    }
}
// SetErrorCode sets the errorCode property value. Represents the error code for this error.
func (m *EducationSynchronizationError) SetErrorCode(value *string)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetErrorMessage sets the errorMessage property value. Contains a description of the error.
func (m *EducationSynchronizationError) SetErrorMessage(value *string)() {
    if m != nil {
        m.errorMessage = value
    }
}
// SetJoiningValue sets the joiningValue property value. The unique identifier for the entry.
func (m *EducationSynchronizationError) SetJoiningValue(value *string)() {
    if m != nil {
        m.joiningValue = value
    }
}
// SetRecordedDateTime sets the recordedDateTime property value. The time of occurrence of this error.
func (m *EducationSynchronizationError) SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.recordedDateTime = value
    }
}
// SetReportableIdentifier sets the reportableIdentifier property value. The identifier of this error entry.
func (m *EducationSynchronizationError) SetReportableIdentifier(value *string)() {
    if m != nil {
        m.reportableIdentifier = value
    }
}
