package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubmissionAdminReview 
type SubmissionAdminReview struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The reviewBy property
    reviewBy *string
    // The reviewDateTime property
    reviewDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The reviewResult property
    reviewResult *SubmissionResultCategory
}
// NewSubmissionAdminReview instantiates a new submissionAdminReview and sets the default values.
func NewSubmissionAdminReview()(*SubmissionAdminReview) {
    m := &SubmissionAdminReview{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSubmissionAdminReviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubmissionAdminReviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubmissionAdminReview(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubmissionAdminReview) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubmissionAdminReview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["reviewBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewBy(val)
        }
        return nil
    }
    res["reviewDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewDateTime(val)
        }
        return nil
    }
    res["reviewResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubmissionResultCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewResult(val.(*SubmissionResultCategory))
        }
        return nil
    }
    return res
}
// GetReviewBy gets the reviewBy property value. The reviewBy property
func (m *SubmissionAdminReview) GetReviewBy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reviewBy
    }
}
// GetReviewDateTime gets the reviewDateTime property value. The reviewDateTime property
func (m *SubmissionAdminReview) GetReviewDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.reviewDateTime
    }
}
// GetReviewResult gets the reviewResult property value. The reviewResult property
func (m *SubmissionAdminReview) GetReviewResult()(*SubmissionResultCategory) {
    if m == nil {
        return nil
    } else {
        return m.reviewResult
    }
}
// Serialize serializes information the current object
func (m *SubmissionAdminReview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("reviewBy", m.GetReviewBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("reviewDateTime", m.GetReviewDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetReviewResult() != nil {
        cast := (*m.GetReviewResult()).String()
        err := writer.WriteStringValue("reviewResult", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubmissionAdminReview) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReviewBy sets the reviewBy property value. The reviewBy property
func (m *SubmissionAdminReview) SetReviewBy(value *string)() {
    if m != nil {
        m.reviewBy = value
    }
}
// SetReviewDateTime sets the reviewDateTime property value. The reviewDateTime property
func (m *SubmissionAdminReview) SetReviewDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.reviewDateTime = value
    }
}
// SetReviewResult sets the reviewResult property value. The reviewResult property
func (m *SubmissionAdminReview) SetReviewResult(value *SubmissionResultCategory)() {
    if m != nil {
        m.reviewResult = value
    }
}
