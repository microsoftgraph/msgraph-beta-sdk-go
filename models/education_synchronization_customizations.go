package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EducationSynchronizationCustomizations struct {
    EducationSynchronizationCustomizationsBase
}
// NewEducationSynchronizationCustomizations instantiates a new EducationSynchronizationCustomizations and sets the default values.
func NewEducationSynchronizationCustomizations()(*EducationSynchronizationCustomizations) {
    m := &EducationSynchronizationCustomizations{
        EducationSynchronizationCustomizationsBase: *NewEducationSynchronizationCustomizationsBase(),
    }
    odataTypeValue := "#microsoft.graph.educationSynchronizationCustomizations"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationSynchronizationCustomizationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEducationSynchronizationCustomizationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationCustomizations(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EducationSynchronizationCustomizations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationCustomizationsBase.GetFieldDeserializers()
    res["school"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchool(val.(EducationSynchronizationCustomizationable))
        }
        return nil
    }
    res["section"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSection(val.(EducationSynchronizationCustomizationable))
        }
        return nil
    }
    res["student"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudent(val.(EducationSynchronizationCustomizationable))
        }
        return nil
    }
    res["studentEnrollment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudentEnrollment(val.(EducationSynchronizationCustomizationable))
        }
        return nil
    }
    res["teacher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeacher(val.(EducationSynchronizationCustomizationable))
        }
        return nil
    }
    res["teacherRoster"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeacherRoster(val.(EducationSynchronizationCustomizationable))
        }
        return nil
    }
    return res
}
// GetSchool gets the school property value. The school property
// returns a EducationSynchronizationCustomizationable when successful
func (m *EducationSynchronizationCustomizations) GetSchool()(EducationSynchronizationCustomizationable) {
    val, err := m.GetBackingStore().Get("school")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationSynchronizationCustomizationable)
    }
    return nil
}
// GetSection gets the section property value. The section property
// returns a EducationSynchronizationCustomizationable when successful
func (m *EducationSynchronizationCustomizations) GetSection()(EducationSynchronizationCustomizationable) {
    val, err := m.GetBackingStore().Get("section")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationSynchronizationCustomizationable)
    }
    return nil
}
// GetStudent gets the student property value. The student property
// returns a EducationSynchronizationCustomizationable when successful
func (m *EducationSynchronizationCustomizations) GetStudent()(EducationSynchronizationCustomizationable) {
    val, err := m.GetBackingStore().Get("student")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationSynchronizationCustomizationable)
    }
    return nil
}
// GetStudentEnrollment gets the studentEnrollment property value. The studentEnrollment property
// returns a EducationSynchronizationCustomizationable when successful
func (m *EducationSynchronizationCustomizations) GetStudentEnrollment()(EducationSynchronizationCustomizationable) {
    val, err := m.GetBackingStore().Get("studentEnrollment")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationSynchronizationCustomizationable)
    }
    return nil
}
// GetTeacher gets the teacher property value. The teacher property
// returns a EducationSynchronizationCustomizationable when successful
func (m *EducationSynchronizationCustomizations) GetTeacher()(EducationSynchronizationCustomizationable) {
    val, err := m.GetBackingStore().Get("teacher")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationSynchronizationCustomizationable)
    }
    return nil
}
// GetTeacherRoster gets the teacherRoster property value. The teacherRoster property
// returns a EducationSynchronizationCustomizationable when successful
func (m *EducationSynchronizationCustomizations) GetTeacherRoster()(EducationSynchronizationCustomizationable) {
    val, err := m.GetBackingStore().Get("teacherRoster")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationSynchronizationCustomizationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationSynchronizationCustomizations) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationSynchronizationCustomizationsBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("school", m.GetSchool())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("section", m.GetSection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("student", m.GetStudent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("studentEnrollment", m.GetStudentEnrollment())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teacher", m.GetTeacher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teacherRoster", m.GetTeacherRoster())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSchool sets the school property value. The school property
func (m *EducationSynchronizationCustomizations) SetSchool(value EducationSynchronizationCustomizationable)() {
    err := m.GetBackingStore().Set("school", value)
    if err != nil {
        panic(err)
    }
}
// SetSection sets the section property value. The section property
func (m *EducationSynchronizationCustomizations) SetSection(value EducationSynchronizationCustomizationable)() {
    err := m.GetBackingStore().Set("section", value)
    if err != nil {
        panic(err)
    }
}
// SetStudent sets the student property value. The student property
func (m *EducationSynchronizationCustomizations) SetStudent(value EducationSynchronizationCustomizationable)() {
    err := m.GetBackingStore().Set("student", value)
    if err != nil {
        panic(err)
    }
}
// SetStudentEnrollment sets the studentEnrollment property value. The studentEnrollment property
func (m *EducationSynchronizationCustomizations) SetStudentEnrollment(value EducationSynchronizationCustomizationable)() {
    err := m.GetBackingStore().Set("studentEnrollment", value)
    if err != nil {
        panic(err)
    }
}
// SetTeacher sets the teacher property value. The teacher property
func (m *EducationSynchronizationCustomizations) SetTeacher(value EducationSynchronizationCustomizationable)() {
    err := m.GetBackingStore().Set("teacher", value)
    if err != nil {
        panic(err)
    }
}
// SetTeacherRoster sets the teacherRoster property value. The teacherRoster property
func (m *EducationSynchronizationCustomizations) SetTeacherRoster(value EducationSynchronizationCustomizationable)() {
    err := m.GetBackingStore().Set("teacherRoster", value)
    if err != nil {
        panic(err)
    }
}
type EducationSynchronizationCustomizationsable interface {
    EducationSynchronizationCustomizationsBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSchool()(EducationSynchronizationCustomizationable)
    GetSection()(EducationSynchronizationCustomizationable)
    GetStudent()(EducationSynchronizationCustomizationable)
    GetStudentEnrollment()(EducationSynchronizationCustomizationable)
    GetTeacher()(EducationSynchronizationCustomizationable)
    GetTeacherRoster()(EducationSynchronizationCustomizationable)
    SetSchool(value EducationSynchronizationCustomizationable)()
    SetSection(value EducationSynchronizationCustomizationable)()
    SetStudent(value EducationSynchronizationCustomizationable)()
    SetStudentEnrollment(value EducationSynchronizationCustomizationable)()
    SetTeacher(value EducationSynchronizationCustomizationable)()
    SetTeacherRoster(value EducationSynchronizationCustomizationable)()
}
