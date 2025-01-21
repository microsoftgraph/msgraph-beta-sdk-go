package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EducationSynchronizationError struct {
    Entity
}
// NewEducationSynchronizationError instantiates a new EducationSynchronizationError and sets the default values.
func NewEducationSynchronizationError()(*EducationSynchronizationError) {
    m := &EducationSynchronizationError{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationSynchronizationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationSynchronizationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationError(), nil
}
// GetEntryType gets the entryType property value. The entryType property
// returns a *string when successful
func (m *EducationSynchronizationError) GetEntryType()(*string) {
    val, err := m.GetBackingStore().Get("entryType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetErrorCode gets the errorCode property value. The errorCode property
// returns a *string when successful
func (m *EducationSynchronizationError) GetErrorCode()(*string) {
    val, err := m.GetBackingStore().Get("errorCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetErrorMessage gets the errorMessage property value. The errorMessage property
// returns a *string when successful
func (m *EducationSynchronizationError) GetErrorMessage()(*string) {
    val, err := m.GetBackingStore().Get("errorMessage")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationSynchronizationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["entryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntryType(val)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["errorMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorMessage(val)
        }
        return nil
    }
    res["joiningValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoiningValue(val)
        }
        return nil
    }
    res["recordedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordedDateTime(val)
        }
        return nil
    }
    res["reportableIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetJoiningValue gets the joiningValue property value. The joiningValue property
// returns a *string when successful
func (m *EducationSynchronizationError) GetJoiningValue()(*string) {
    val, err := m.GetBackingStore().Get("joiningValue")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRecordedDateTime gets the recordedDateTime property value. The recordedDateTime property
// returns a *Time when successful
func (m *EducationSynchronizationError) GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("recordedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetReportableIdentifier gets the reportableIdentifier property value. The reportableIdentifier property
// returns a *string when successful
func (m *EducationSynchronizationError) GetReportableIdentifier()(*string) {
    val, err := m.GetBackingStore().Get("reportableIdentifier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationSynchronizationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetEntryType sets the entryType property value. The entryType property
func (m *EducationSynchronizationError) SetEntryType(value *string)() {
    err := m.GetBackingStore().Set("entryType", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorCode sets the errorCode property value. The errorCode property
func (m *EducationSynchronizationError) SetErrorCode(value *string)() {
    err := m.GetBackingStore().Set("errorCode", value)
    if err != nil {
        panic(err)
    }
}
// SetErrorMessage sets the errorMessage property value. The errorMessage property
func (m *EducationSynchronizationError) SetErrorMessage(value *string)() {
    err := m.GetBackingStore().Set("errorMessage", value)
    if err != nil {
        panic(err)
    }
}
// SetJoiningValue sets the joiningValue property value. The joiningValue property
func (m *EducationSynchronizationError) SetJoiningValue(value *string)() {
    err := m.GetBackingStore().Set("joiningValue", value)
    if err != nil {
        panic(err)
    }
}
// SetRecordedDateTime sets the recordedDateTime property value. The recordedDateTime property
func (m *EducationSynchronizationError) SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("recordedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetReportableIdentifier sets the reportableIdentifier property value. The reportableIdentifier property
func (m *EducationSynchronizationError) SetReportableIdentifier(value *string)() {
    err := m.GetBackingStore().Set("reportableIdentifier", value)
    if err != nil {
        panic(err)
    }
}
type EducationSynchronizationErrorable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEntryType()(*string)
    GetErrorCode()(*string)
    GetErrorMessage()(*string)
    GetJoiningValue()(*string)
    GetRecordedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReportableIdentifier()(*string)
    SetEntryType(value *string)()
    SetErrorCode(value *string)()
    SetErrorMessage(value *string)()
    SetJoiningValue(value *string)()
    SetRecordedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReportableIdentifier(value *string)()
}
