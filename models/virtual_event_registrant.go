package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventRegistrant 
type VirtualEventRegistrant struct {
    Entity
}
// NewVirtualEventRegistrant instantiates a new virtualEventRegistrant and sets the default values.
func NewVirtualEventRegistrant()(*VirtualEventRegistrant) {
    m := &VirtualEventRegistrant{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEventRegistrantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventRegistrantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventRegistrant(), nil
}
// GetCancelationDateTime gets the cancelationDateTime property value. Time in UTC when the registrant cancels their registration for the virtual event. Only appears when applicable.
func (m *VirtualEventRegistrant) GetCancelationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("cancelationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetEmail gets the email property value. Email address of the registrant.
func (m *VirtualEventRegistrant) GetEmail()(*string) {
    val, err := m.GetBackingStore().Get("email")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEventRegistrant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["cancelationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCancelationDateTime(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["firstName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstName(val)
        }
        return nil
    }
    res["lastName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["registrationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationDateTime(val)
        }
        return nil
    }
    res["registrationQuestionAnswers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventRegistrationQuestionAnswerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventRegistrationQuestionAnswerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventRegistrationQuestionAnswerable)
                }
            }
            m.SetRegistrationQuestionAnswers(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVirtualEventAttendeeRegistrationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*VirtualEventAttendeeRegistrationStatus))
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetFirstName gets the firstName property value. First name of the registrant.
func (m *VirtualEventRegistrant) GetFirstName()(*string) {
    val, err := m.GetBackingStore().Get("firstName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastName gets the lastName property value. Last name of the registrant.
func (m *VirtualEventRegistrant) GetLastName()(*string) {
    val, err := m.GetBackingStore().Get("lastName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VirtualEventRegistrant) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRegistrationDateTime gets the registrationDateTime property value. Time in UTC when the registrant registers for the virtual event.
func (m *VirtualEventRegistrant) GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("registrationDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetRegistrationQuestionAnswers gets the registrationQuestionAnswers property value. The registrant's answer to the registration questions.
func (m *VirtualEventRegistrant) GetRegistrationQuestionAnswers()([]VirtualEventRegistrationQuestionAnswerable) {
    val, err := m.GetBackingStore().Get("registrationQuestionAnswers")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualEventRegistrationQuestionAnswerable)
    }
    return nil
}
// GetStatus gets the status property value. Registration status of the registrant. Read-only.
func (m *VirtualEventRegistrant) GetStatus()(*VirtualEventAttendeeRegistrationStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VirtualEventAttendeeRegistrationStatus)
    }
    return nil
}
// GetUserId gets the userId property value. The registrant's AAD user ID. Only appears when the registrant is registered in AAD.
func (m *VirtualEventRegistrant) GetUserId()(*string) {
    val, err := m.GetBackingStore().Get("userId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEventRegistrant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("cancelationDateTime", m.GetCancelationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("firstName", m.GetFirstName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lastName", m.GetLastName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("registrationDateTime", m.GetRegistrationDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRegistrationQuestionAnswers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegistrationQuestionAnswers()))
        for i, v := range m.GetRegistrationQuestionAnswers() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("registrationQuestionAnswers", cast)
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
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCancelationDateTime sets the cancelationDateTime property value. Time in UTC when the registrant cancels their registration for the virtual event. Only appears when applicable.
func (m *VirtualEventRegistrant) SetCancelationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("cancelationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetEmail sets the email property value. Email address of the registrant.
func (m *VirtualEventRegistrant) SetEmail(value *string)() {
    err := m.GetBackingStore().Set("email", value)
    if err != nil {
        panic(err)
    }
}
// SetFirstName sets the firstName property value. First name of the registrant.
func (m *VirtualEventRegistrant) SetFirstName(value *string)() {
    err := m.GetBackingStore().Set("firstName", value)
    if err != nil {
        panic(err)
    }
}
// SetLastName sets the lastName property value. Last name of the registrant.
func (m *VirtualEventRegistrant) SetLastName(value *string)() {
    err := m.GetBackingStore().Set("lastName", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VirtualEventRegistrant) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistrationDateTime sets the registrationDateTime property value. Time in UTC when the registrant registers for the virtual event.
func (m *VirtualEventRegistrant) SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("registrationDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistrationQuestionAnswers sets the registrationQuestionAnswers property value. The registrant's answer to the registration questions.
func (m *VirtualEventRegistrant) SetRegistrationQuestionAnswers(value []VirtualEventRegistrationQuestionAnswerable)() {
    err := m.GetBackingStore().Set("registrationQuestionAnswers", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Registration status of the registrant. Read-only.
func (m *VirtualEventRegistrant) SetStatus(value *VirtualEventAttendeeRegistrationStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// SetUserId sets the userId property value. The registrant's AAD user ID. Only appears when the registrant is registered in AAD.
func (m *VirtualEventRegistrant) SetUserId(value *string)() {
    err := m.GetBackingStore().Set("userId", value)
    if err != nil {
        panic(err)
    }
}
// VirtualEventRegistrantable 
type VirtualEventRegistrantable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCancelationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEmail()(*string)
    GetFirstName()(*string)
    GetLastName()(*string)
    GetOdataType()(*string)
    GetRegistrationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRegistrationQuestionAnswers()([]VirtualEventRegistrationQuestionAnswerable)
    GetStatus()(*VirtualEventAttendeeRegistrationStatus)
    GetUserId()(*string)
    SetCancelationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEmail(value *string)()
    SetFirstName(value *string)()
    SetLastName(value *string)()
    SetOdataType(value *string)()
    SetRegistrationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRegistrationQuestionAnswers(value []VirtualEventRegistrationQuestionAnswerable)()
    SetStatus(value *VirtualEventAttendeeRegistrationStatus)()
    SetUserId(value *string)()
}
