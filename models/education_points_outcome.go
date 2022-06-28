package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationPointsOutcome 
type EducationPointsOutcome struct {
    EducationOutcome
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The numeric grade the teacher has given the student for this assignment.
    points EducationAssignmentPointsGradeable
    // A copy of the points property that is made when the grade is released to the student.
    publishedPoints EducationAssignmentPointsGradeable
}
// NewEducationPointsOutcome instantiates a new EducationPointsOutcome and sets the default values.
func NewEducationPointsOutcome()(*EducationPointsOutcome) {
    m := &EducationPointsOutcome{
        EducationOutcome: *NewEducationOutcome(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationPointsOutcomeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationPointsOutcomeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationPointsOutcome(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationPointsOutcome) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationPointsOutcome) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationOutcome.GetFieldDeserializers()
    res["points"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationAssignmentPointsGradeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPoints(val.(EducationAssignmentPointsGradeable))
        }
        return nil
    }
    res["publishedPoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationAssignmentPointsGradeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedPoints(val.(EducationAssignmentPointsGradeable))
        }
        return nil
    }
    return res
}
// GetPoints gets the points property value. The numeric grade the teacher has given the student for this assignment.
func (m *EducationPointsOutcome) GetPoints()(EducationAssignmentPointsGradeable) {
    if m == nil {
        return nil
    } else {
        return m.points
    }
}
// GetPublishedPoints gets the publishedPoints property value. A copy of the points property that is made when the grade is released to the student.
func (m *EducationPointsOutcome) GetPublishedPoints()(EducationAssignmentPointsGradeable) {
    if m == nil {
        return nil
    } else {
        return m.publishedPoints
    }
}
// Serialize serializes information the current object
func (m *EducationPointsOutcome) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationOutcome.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("points", m.GetPoints())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publishedPoints", m.GetPublishedPoints())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationPointsOutcome) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPoints sets the points property value. The numeric grade the teacher has given the student for this assignment.
func (m *EducationPointsOutcome) SetPoints(value EducationAssignmentPointsGradeable)() {
    if m != nil {
        m.points = value
    }
}
// SetPublishedPoints sets the publishedPoints property value. A copy of the points property that is made when the grade is released to the student.
func (m *EducationPointsOutcome) SetPublishedPoints(value EducationAssignmentPointsGradeable)() {
    if m != nil {
        m.publishedPoints = value
    }
}
