package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GoalsExportJob struct {
    LongRunningOperation
}
// NewGoalsExportJob instantiates a new GoalsExportJob and sets the default values.
func NewGoalsExportJob()(*GoalsExportJob) {
    m := &GoalsExportJob{
        LongRunningOperation: *NewLongRunningOperation(),
    }
    return m
}
// CreateGoalsExportJobFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGoalsExportJobFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGoalsExportJob(), nil
}
// GetContent gets the content property value. The content of the goalsExportJob.
// returns a []byte when successful
func (m *GoalsExportJob) GetContent()([]byte) {
    val, err := m.GetBackingStore().Get("content")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]byte)
    }
    return nil
}
// GetExpirationDateTime gets the expirationDateTime property value. The date and time of expiry of the result of the operation.
// returns a *Time when successful
func (m *GoalsExportJob) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("expirationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetExplorerViewId gets the explorerViewId property value. The unique identifier of the explorer view to be exported.
// returns a *string when successful
func (m *GoalsExportJob) GetExplorerViewId()(*string) {
    val, err := m.GetBackingStore().Get("explorerViewId")
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
func (m *GoalsExportJob) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LongRunningOperation.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["explorerViewId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExplorerViewId(val)
        }
        return nil
    }
    res["goalsOrganizationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGoalsOrganizationId(val)
        }
        return nil
    }
    return res
}
// GetGoalsOrganizationId gets the goalsOrganizationId property value. The unique identifier of the viva goals organization.
// returns a *string when successful
func (m *GoalsExportJob) GetGoalsOrganizationId()(*string) {
    val, err := m.GetBackingStore().Get("goalsOrganizationId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *GoalsExportJob) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LongRunningOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("explorerViewId", m.GetExplorerViewId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("goalsOrganizationId", m.GetGoalsOrganizationId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The content of the goalsExportJob.
func (m *GoalsExportJob) SetContent(value []byte)() {
    err := m.GetBackingStore().Set("content", value)
    if err != nil {
        panic(err)
    }
}
// SetExpirationDateTime sets the expirationDateTime property value. The date and time of expiry of the result of the operation.
func (m *GoalsExportJob) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("expirationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetExplorerViewId sets the explorerViewId property value. The unique identifier of the explorer view to be exported.
func (m *GoalsExportJob) SetExplorerViewId(value *string)() {
    err := m.GetBackingStore().Set("explorerViewId", value)
    if err != nil {
        panic(err)
    }
}
// SetGoalsOrganizationId sets the goalsOrganizationId property value. The unique identifier of the viva goals organization.
func (m *GoalsExportJob) SetGoalsOrganizationId(value *string)() {
    err := m.GetBackingStore().Set("goalsOrganizationId", value)
    if err != nil {
        panic(err)
    }
}
type GoalsExportJobable interface {
    LongRunningOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExplorerViewId()(*string)
    GetGoalsOrganizationId()(*string)
    SetContent(value []byte)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExplorerViewId(value *string)()
    SetGoalsOrganizationId(value *string)()
}
