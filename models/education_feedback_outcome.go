package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationFeedbackOutcome 
type EducationFeedbackOutcome struct {
    EducationOutcome
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Teacher's written feedback to the student.
    feedback EducationFeedbackable
    // A copy of the feedback property that is made when the grade is released to the student.
    publishedFeedback EducationFeedbackable
}
// NewEducationFeedbackOutcome instantiates a new EducationFeedbackOutcome and sets the default values.
func NewEducationFeedbackOutcome()(*EducationFeedbackOutcome) {
    m := &EducationFeedbackOutcome{
        EducationOutcome: *NewEducationOutcome(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationFeedbackOutcomeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationFeedbackOutcomeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationFeedbackOutcome(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationFeedbackOutcome) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFeedback gets the feedback property value. Teacher's written feedback to the student.
func (m *EducationFeedbackOutcome) GetFeedback()(EducationFeedbackable) {
    if m == nil {
        return nil
    } else {
        return m.feedback
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationFeedbackOutcome) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationOutcome.GetFieldDeserializers()
    res["feedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationFeedbackFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedback(val.(EducationFeedbackable))
        }
        return nil
    }
    res["publishedFeedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationFeedbackFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishedFeedback(val.(EducationFeedbackable))
        }
        return nil
    }
    return res
}
// GetPublishedFeedback gets the publishedFeedback property value. A copy of the feedback property that is made when the grade is released to the student.
func (m *EducationFeedbackOutcome) GetPublishedFeedback()(EducationFeedbackable) {
    if m == nil {
        return nil
    } else {
        return m.publishedFeedback
    }
}
// Serialize serializes information the current object
func (m *EducationFeedbackOutcome) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationOutcome.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("feedback", m.GetFeedback())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publishedFeedback", m.GetPublishedFeedback())
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
func (m *EducationFeedbackOutcome) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFeedback sets the feedback property value. Teacher's written feedback to the student.
func (m *EducationFeedbackOutcome) SetFeedback(value EducationFeedbackable)() {
    if m != nil {
        m.feedback = value
    }
}
// SetPublishedFeedback sets the publishedFeedback property value. A copy of the feedback property that is made when the grade is released to the student.
func (m *EducationFeedbackOutcome) SetPublishedFeedback(value EducationFeedbackable)() {
    if m != nil {
        m.publishedFeedback = value
    }
}
