package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationalActivity 
type EducationalActivity struct {
    ItemFacet
}
// NewEducationalActivity instantiates a new educationalActivity and sets the default values.
func NewEducationalActivity()(*EducationalActivity) {
    m := &EducationalActivity{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.educationalActivity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationalActivityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationalActivityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationalActivity(), nil
}
// GetCompletionMonthYear gets the completionMonthYear property value. The month and year the user graduated or completed the activity.
func (m *EducationalActivity) GetCompletionMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("completionMonthYear")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetEndMonthYear gets the endMonthYear property value. The month and year the user completed the educational activity referenced.
func (m *EducationalActivity) GetEndMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("endMonthYear")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationalActivity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["completionMonthYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletionMonthYear(val)
        }
        return nil
    }
    res["endMonthYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndMonthYear(val)
        }
        return nil
    }
    res["institution"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInstitutionDataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstitution(val.(InstitutionDataable))
        }
        return nil
    }
    res["program"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationalActivityDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgram(val.(EducationalActivityDetailable))
        }
        return nil
    }
    res["startMonthYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMonthYear(val)
        }
        return nil
    }
    return res
}
// GetInstitution gets the institution property value. The institution property
func (m *EducationalActivity) GetInstitution()(InstitutionDataable) {
    val, err := m.GetBackingStore().Get("institution")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(InstitutionDataable)
    }
    return nil
}
// GetProgram gets the program property value. The program property
func (m *EducationalActivity) GetProgram()(EducationalActivityDetailable) {
    val, err := m.GetBackingStore().Get("program")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationalActivityDetailable)
    }
    return nil
}
// GetStartMonthYear gets the startMonthYear property value. The month and year the user commenced the activity referenced.
func (m *EducationalActivity) GetStartMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    val, err := m.GetBackingStore().Get("startMonthYear")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationalActivity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteDateOnlyValue("completionMonthYear", m.GetCompletionMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("endMonthYear", m.GetEndMonthYear())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("institution", m.GetInstitution())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("program", m.GetProgram())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("startMonthYear", m.GetStartMonthYear())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompletionMonthYear sets the completionMonthYear property value. The month and year the user graduated or completed the activity.
func (m *EducationalActivity) SetCompletionMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("completionMonthYear", value)
    if err != nil {
        panic(err)
    }
}
// SetEndMonthYear sets the endMonthYear property value. The month and year the user completed the educational activity referenced.
func (m *EducationalActivity) SetEndMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("endMonthYear", value)
    if err != nil {
        panic(err)
    }
}
// SetInstitution sets the institution property value. The institution property
func (m *EducationalActivity) SetInstitution(value InstitutionDataable)() {
    err := m.GetBackingStore().Set("institution", value)
    if err != nil {
        panic(err)
    }
}
// SetProgram sets the program property value. The program property
func (m *EducationalActivity) SetProgram(value EducationalActivityDetailable)() {
    err := m.GetBackingStore().Set("program", value)
    if err != nil {
        panic(err)
    }
}
// SetStartMonthYear sets the startMonthYear property value. The month and year the user commenced the activity referenced.
func (m *EducationalActivity) SetStartMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    err := m.GetBackingStore().Set("startMonthYear", value)
    if err != nil {
        panic(err)
    }
}
// EducationalActivityable 
type EducationalActivityable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCompletionMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetEndMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetInstitution()(InstitutionDataable)
    GetProgram()(EducationalActivityDetailable)
    GetStartMonthYear()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    SetCompletionMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetEndMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetInstitution(value InstitutionDataable)()
    SetProgram(value EducationalActivityDetailable)()
    SetStartMonthYear(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
}
