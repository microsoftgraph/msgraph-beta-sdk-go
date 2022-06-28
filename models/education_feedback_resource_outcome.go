package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationFeedbackResourceOutcome 
type EducationFeedbackResourceOutcome struct {
    EducationOutcome
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The actual feedback resource.
    feedbackResource EducationResourceable
    // The status of the feedback resource. The possible values are: notPublished, pendingPublish, published, failedPublish, and unknownFutureValue.
    resourceStatus *EducationFeedbackResourceOutcomeStatus
}
// NewEducationFeedbackResourceOutcome instantiates a new EducationFeedbackResourceOutcome and sets the default values.
func NewEducationFeedbackResourceOutcome()(*EducationFeedbackResourceOutcome) {
    m := &EducationFeedbackResourceOutcome{
        EducationOutcome: *NewEducationOutcome(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEducationFeedbackResourceOutcomeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationFeedbackResourceOutcomeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationFeedbackResourceOutcome(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationFeedbackResourceOutcome) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFeedbackResource gets the feedbackResource property value. The actual feedback resource.
func (m *EducationFeedbackResourceOutcome) GetFeedbackResource()(EducationResourceable) {
    if m == nil {
        return nil
    } else {
        return m.feedbackResource
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationFeedbackResourceOutcome) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationOutcome.GetFieldDeserializers()
    res["feedbackResource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedbackResource(val.(EducationResourceable))
        }
        return nil
    }
    res["resourceStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationFeedbackResourceOutcomeStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceStatus(val.(*EducationFeedbackResourceOutcomeStatus))
        }
        return nil
    }
    return res
}
// GetResourceStatus gets the resourceStatus property value. The status of the feedback resource. The possible values are: notPublished, pendingPublish, published, failedPublish, and unknownFutureValue.
func (m *EducationFeedbackResourceOutcome) GetResourceStatus()(*EducationFeedbackResourceOutcomeStatus) {
    if m == nil {
        return nil
    } else {
        return m.resourceStatus
    }
}
// Serialize serializes information the current object
func (m *EducationFeedbackResourceOutcome) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationOutcome.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("feedbackResource", m.GetFeedbackResource())
        if err != nil {
            return err
        }
    }
    if m.GetResourceStatus() != nil {
        cast := (*m.GetResourceStatus()).String()
        err = writer.WriteStringValue("resourceStatus", &cast)
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
func (m *EducationFeedbackResourceOutcome) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetFeedbackResource sets the feedbackResource property value. The actual feedback resource.
func (m *EducationFeedbackResourceOutcome) SetFeedbackResource(value EducationResourceable)() {
    if m != nil {
        m.feedbackResource = value
    }
}
// SetResourceStatus sets the resourceStatus property value. The status of the feedback resource. The possible values are: notPublished, pendingPublish, published, failedPublish, and unknownFutureValue.
func (m *EducationFeedbackResourceOutcome) SetResourceStatus(value *EducationFeedbackResourceOutcomeStatus)() {
    if m != nil {
        m.resourceStatus = value
    }
}
