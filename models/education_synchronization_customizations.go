package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSynchronizationCustomizations 
type EducationSynchronizationCustomizations struct {
    EducationSynchronizationCustomizationsBase
    // Customizations for School entities.
    school EducationSynchronizationCustomizationable
    // Customizations for Section entities.
    section EducationSynchronizationCustomizationable
    // Customizations for Student entities.
    student EducationSynchronizationCustomizationable
    // Customizations for Student Enrollments.
    studentEnrollment EducationSynchronizationCustomizationable
    // Customizations for Teacher entities.
    teacher EducationSynchronizationCustomizationable
    // Customizations for Teacher Rosters.
    teacherRoster EducationSynchronizationCustomizationable
}
// NewEducationSynchronizationCustomizations instantiates a new EducationSynchronizationCustomizations and sets the default values.
func NewEducationSynchronizationCustomizations()(*EducationSynchronizationCustomizations) {
    m := &EducationSynchronizationCustomizations{
        EducationSynchronizationCustomizationsBase: *NewEducationSynchronizationCustomizationsBase(),
    }
    odataTypeValue := "#microsoft.graph.educationSynchronizationCustomizations";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEducationSynchronizationCustomizationsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSynchronizationCustomizationsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSynchronizationCustomizations(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSynchronizationCustomizations) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationSynchronizationCustomizationsBase.GetFieldDeserializers()
    res["school"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue , m.SetSchool)
    res["section"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue , m.SetSection)
    res["student"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue , m.SetStudent)
    res["studentEnrollment"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue , m.SetStudentEnrollment)
    res["teacher"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue , m.SetTeacher)
    res["teacherRoster"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEducationSynchronizationCustomizationFromDiscriminatorValue , m.SetTeacherRoster)
    return res
}
// GetSchool gets the school property value. Customizations for School entities.
func (m *EducationSynchronizationCustomizations) GetSchool()(EducationSynchronizationCustomizationable) {
    return m.school
}
// GetSection gets the section property value. Customizations for Section entities.
func (m *EducationSynchronizationCustomizations) GetSection()(EducationSynchronizationCustomizationable) {
    return m.section
}
// GetStudent gets the student property value. Customizations for Student entities.
func (m *EducationSynchronizationCustomizations) GetStudent()(EducationSynchronizationCustomizationable) {
    return m.student
}
// GetStudentEnrollment gets the studentEnrollment property value. Customizations for Student Enrollments.
func (m *EducationSynchronizationCustomizations) GetStudentEnrollment()(EducationSynchronizationCustomizationable) {
    return m.studentEnrollment
}
// GetTeacher gets the teacher property value. Customizations for Teacher entities.
func (m *EducationSynchronizationCustomizations) GetTeacher()(EducationSynchronizationCustomizationable) {
    return m.teacher
}
// GetTeacherRoster gets the teacherRoster property value. Customizations for Teacher Rosters.
func (m *EducationSynchronizationCustomizations) GetTeacherRoster()(EducationSynchronizationCustomizationable) {
    return m.teacherRoster
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
// SetSchool sets the school property value. Customizations for School entities.
func (m *EducationSynchronizationCustomizations) SetSchool(value EducationSynchronizationCustomizationable)() {
    m.school = value
}
// SetSection sets the section property value. Customizations for Section entities.
func (m *EducationSynchronizationCustomizations) SetSection(value EducationSynchronizationCustomizationable)() {
    m.section = value
}
// SetStudent sets the student property value. Customizations for Student entities.
func (m *EducationSynchronizationCustomizations) SetStudent(value EducationSynchronizationCustomizationable)() {
    m.student = value
}
// SetStudentEnrollment sets the studentEnrollment property value. Customizations for Student Enrollments.
func (m *EducationSynchronizationCustomizations) SetStudentEnrollment(value EducationSynchronizationCustomizationable)() {
    m.studentEnrollment = value
}
// SetTeacher sets the teacher property value. Customizations for Teacher entities.
func (m *EducationSynchronizationCustomizations) SetTeacher(value EducationSynchronizationCustomizationable)() {
    m.teacher = value
}
// SetTeacherRoster sets the teacherRoster property value. Customizations for Teacher Rosters.
func (m *EducationSynchronizationCustomizations) SetTeacherRoster(value EducationSynchronizationCustomizationable)() {
    m.teacherRoster = value
}
