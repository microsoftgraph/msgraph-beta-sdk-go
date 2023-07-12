package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VirtualEventRegistration 
type VirtualEventRegistration struct {
    Entity
    // The OdataType property
    OdataType *string
}
// NewVirtualEventRegistration instantiates a new virtualEventRegistration and sets the default values.
func NewVirtualEventRegistration()(*VirtualEventRegistration) {
    m := &VirtualEventRegistration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEventRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVirtualEventRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVirtualEventRegistration(), nil
}
// GetCapacity gets the capacity property value. Total capacity of the virtual event.
func (m *VirtualEventRegistration) GetCapacity()(*int32) {
    val, err := m.GetBackingStore().Get("capacity")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VirtualEventRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["capacity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapacity(val)
        }
        return nil
    }
    res["questions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventRegistrationQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventRegistrationQuestionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventRegistrationQuestionable)
                }
            }
            m.SetQuestions(res)
        }
        return nil
    }
    res["registrants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventRegistrantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventRegistrantable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventRegistrantable)
                }
            }
            m.SetRegistrants(res)
        }
        return nil
    }
    res["registrationWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationWebUrl(val)
        }
        return nil
    }
    return res
}
// GetQuestions gets the questions property value. Registration questions.
func (m *VirtualEventRegistration) GetQuestions()([]VirtualEventRegistrationQuestionable) {
    val, err := m.GetBackingStore().Get("questions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualEventRegistrationQuestionable)
    }
    return nil
}
// GetRegistrants gets the registrants property value. Information of attendees who have registered for the virtual event.
func (m *VirtualEventRegistration) GetRegistrants()([]VirtualEventRegistrantable) {
    val, err := m.GetBackingStore().Get("registrants")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualEventRegistrantable)
    }
    return nil
}
// GetRegistrationWebUrl gets the registrationWebUrl property value. Registration URL of the virtual event.
func (m *VirtualEventRegistration) GetRegistrationWebUrl()(*string) {
    val, err := m.GetBackingStore().Get("registrationWebUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEventRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("capacity", m.GetCapacity())
        if err != nil {
            return err
        }
    }
    if m.GetQuestions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetQuestions()))
        for i, v := range m.GetQuestions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("questions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRegistrants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRegistrants()))
        for i, v := range m.GetRegistrants() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("registrants", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registrationWebUrl", m.GetRegistrationWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCapacity sets the capacity property value. Total capacity of the virtual event.
func (m *VirtualEventRegistration) SetCapacity(value *int32)() {
    err := m.GetBackingStore().Set("capacity", value)
    if err != nil {
        panic(err)
    }
}
// SetQuestions sets the questions property value. Registration questions.
func (m *VirtualEventRegistration) SetQuestions(value []VirtualEventRegistrationQuestionable)() {
    err := m.GetBackingStore().Set("questions", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistrants sets the registrants property value. Information of attendees who have registered for the virtual event.
func (m *VirtualEventRegistration) SetRegistrants(value []VirtualEventRegistrantable)() {
    err := m.GetBackingStore().Set("registrants", value)
    if err != nil {
        panic(err)
    }
}
// SetRegistrationWebUrl sets the registrationWebUrl property value. Registration URL of the virtual event.
func (m *VirtualEventRegistration) SetRegistrationWebUrl(value *string)() {
    err := m.GetBackingStore().Set("registrationWebUrl", value)
    if err != nil {
        panic(err)
    }
}
// VirtualEventRegistrationable 
type VirtualEventRegistrationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapacity()(*int32)
    GetQuestions()([]VirtualEventRegistrationQuestionable)
    GetRegistrants()([]VirtualEventRegistrantable)
    GetRegistrationWebUrl()(*string)
    SetCapacity(value *int32)()
    SetQuestions(value []VirtualEventRegistrationQuestionable)()
    SetRegistrants(value []VirtualEventRegistrantable)()
    SetRegistrationWebUrl(value *string)()
}
