package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AppLogCollectionRequest struct {
    Entity
    // Time at which the upload log request reached a terminal state
    completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // List of log folders.
    customLogFolders []string;
    // Error message if any during the upload process
    errorMessage *string;
    // Log upload status. Possible values are: pending, completed, failed.
    status *AppLogUploadState;
}
// Instantiates a new appLogCollectionRequest and sets the default values.
func NewAppLogCollectionRequest()(*AppLogCollectionRequest) {
    m := &AppLogCollectionRequest{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the completedDateTime property value. Time at which the upload log request reached a terminal state
func (m *AppLogCollectionRequest) GetCompletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// Gets the customLogFolders property value. List of log folders.
func (m *AppLogCollectionRequest) GetCustomLogFolders()([]string) {
    if m == nil {
        return nil
    } else {
        return m.customLogFolders
    }
}
// Gets the errorMessage property value. Error message if any during the upload process
func (m *AppLogCollectionRequest) GetErrorMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorMessage
    }
}
// Gets the status property value. Log upload status. Possible values are: pending, completed, failed.
func (m *AppLogCollectionRequest) GetStatus()(*AppLogUploadState) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *AppLogCollectionRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCompletedDateTime(val)
        return nil
    }
    res["customLogFolders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetCustomLogFolders(res)
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppLogUploadState)
        if err != nil {
            return err
        }
        cast := val.(AppLogUploadState)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *AppLogCollectionRequest) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AppLogCollectionRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("customLogFolders", m.GetCustomLogFolders())
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
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the completedDateTime property value. Time at which the upload log request reached a terminal state
// Parameters:
//  - value : Value to set for the completedDateTime property.
func (m *AppLogCollectionRequest) SetCompletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.completedDateTime = value
}
// Sets the customLogFolders property value. List of log folders.
// Parameters:
//  - value : Value to set for the customLogFolders property.
func (m *AppLogCollectionRequest) SetCustomLogFolders(value []string)() {
    m.customLogFolders = value
}
// Sets the errorMessage property value. Error message if any during the upload process
// Parameters:
//  - value : Value to set for the errorMessage property.
func (m *AppLogCollectionRequest) SetErrorMessage(value *string)() {
    m.errorMessage = value
}
// Sets the status property value. Log upload status. Possible values are: pending, completed, failed.
// Parameters:
//  - value : Value to set for the status property.
func (m *AppLogCollectionRequest) SetStatus(value *AppLogUploadState)() {
    m.status = value
}
