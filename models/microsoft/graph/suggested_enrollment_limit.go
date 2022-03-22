package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SuggestedEnrollmentLimit the suggestedEnrollmentLimit resource represents the suggested enrollment limit when given an enrollment type.
type SuggestedEnrollmentLimit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The suggested enrollment limit within a day
    suggestedDailyLimit *int32;
}
// NewSuggestedEnrollmentLimit instantiates a new suggestedEnrollmentLimit and sets the default values.
func NewSuggestedEnrollmentLimit()(*SuggestedEnrollmentLimit) {
    m := &SuggestedEnrollmentLimit{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSuggestedEnrollmentLimitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSuggestedEnrollmentLimitFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSuggestedEnrollmentLimit(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SuggestedEnrollmentLimit) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SuggestedEnrollmentLimit) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["suggestedDailyLimit"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestedDailyLimit(val)
        }
        return nil
    }
    return res
}
// GetSuggestedDailyLimit gets the suggestedDailyLimit property value. The suggested enrollment limit within a day
func (m *SuggestedEnrollmentLimit) GetSuggestedDailyLimit()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.suggestedDailyLimit
    }
}
// Serialize serializes information the current object
func (m *SuggestedEnrollmentLimit) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("suggestedDailyLimit", m.GetSuggestedDailyLimit())
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
func (m *SuggestedEnrollmentLimit) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetSuggestedDailyLimit sets the suggestedDailyLimit property value. The suggested enrollment limit within a day
func (m *SuggestedEnrollmentLimit) SetSuggestedDailyLimit(value *int32)() {
    if m != nil {
        m.suggestedDailyLimit = value
    }
}
