package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DocumentProcessingJob struct {
    Entity
}
// NewDocumentProcessingJob instantiates a new DocumentProcessingJob and sets the default values.
func NewDocumentProcessingJob()(*DocumentProcessingJob) {
    m := &DocumentProcessingJob{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDocumentProcessingJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDocumentProcessingJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDocumentProcessingJob(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of item creation. Read-only.
// returns a *Time when successful
func (m *DocumentProcessingJob) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("createdDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DocumentProcessingJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["jobType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDocumentProcessingJobType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobType(val.(*DocumentProcessingJobType))
        }
        return nil
    }
    res["listItemUniqueId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetListItemUniqueId(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDocumentProcessingJobStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DocumentProcessingJobStatus))
        }
        return nil
    }
    return res
}
// GetJobType gets the jobType property value. The document processing job type. The possible values are: file, folder
// returns a *DocumentProcessingJobType when successful
func (m *DocumentProcessingJob) GetJobType()(*DocumentProcessingJobType) {
    val, err := m.GetBackingStore().Get("jobType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DocumentProcessingJobType)
    }
    return nil
}
// GetListItemUniqueId gets the listItemUniqueId property value. The listItemUniqueId of the file, or folder to process. Use GET driveItem resource operation and read  sharepointIds property to get listItemUniqueId.
// returns a *string when successful
func (m *DocumentProcessingJob) GetListItemUniqueId()(*string) {
    val, err := m.GetBackingStore().Get("listItemUniqueId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The document processing Job status. The possible values are: inProgress, completed, failed, unknownFutureValue.
// returns a *DocumentProcessingJobStatus when successful
func (m *DocumentProcessingJob) GetStatus()(*DocumentProcessingJobStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DocumentProcessingJobStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DocumentProcessingJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetJobType() != nil {
        cast := (*m.GetJobType()).String()
        err = writer.WriteStringValue("jobType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("listItemUniqueId", m.GetListItemUniqueId())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of item creation. Read-only.
func (m *DocumentProcessingJob) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("createdDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetJobType sets the jobType property value. The document processing job type. The possible values are: file, folder
func (m *DocumentProcessingJob) SetJobType(value *DocumentProcessingJobType)() {
    err := m.GetBackingStore().Set("jobType", value)
    if err != nil {
        panic(err)
    }
}
// SetListItemUniqueId sets the listItemUniqueId property value. The listItemUniqueId of the file, or folder to process. Use GET driveItem resource operation and read  sharepointIds property to get listItemUniqueId.
func (m *DocumentProcessingJob) SetListItemUniqueId(value *string)() {
    err := m.GetBackingStore().Set("listItemUniqueId", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The document processing Job status. The possible values are: inProgress, completed, failed, unknownFutureValue.
func (m *DocumentProcessingJob) SetStatus(value *DocumentProcessingJobStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
type DocumentProcessingJobable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetJobType()(*DocumentProcessingJobType)
    GetListItemUniqueId()(*string)
    GetStatus()(*DocumentProcessingJobStatus)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetJobType(value *DocumentProcessingJobType)()
    SetListItemUniqueId(value *string)()
    SetStatus(value *DocumentProcessingJobStatus)()
}
