package windowsupdates

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExpeditedQualityUpdateReference 
type ExpeditedQualityUpdateReference struct {
    QualityUpdateReference
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Specifies other content to consider as equivalent. Supports a subset of the values for equivalentContentOption. Default value is latestSecurity. Possible values are: latestSecurity, unknownFutureValue.
    equivalentContent *EquivalentContentOption
}
// NewExpeditedQualityUpdateReference instantiates a new ExpeditedQualityUpdateReference and sets the default values.
func NewExpeditedQualityUpdateReference()(*ExpeditedQualityUpdateReference) {
    m := &ExpeditedQualityUpdateReference{
        QualityUpdateReference: *NewQualityUpdateReference(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateExpeditedQualityUpdateReferenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExpeditedQualityUpdateReferenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExpeditedQualityUpdateReference(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ExpeditedQualityUpdateReference) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEquivalentContent gets the equivalentContent property value. Specifies other content to consider as equivalent. Supports a subset of the values for equivalentContentOption. Default value is latestSecurity. Possible values are: latestSecurity, unknownFutureValue.
func (m *ExpeditedQualityUpdateReference) GetEquivalentContent()(*EquivalentContentOption) {
    if m == nil {
        return nil
    } else {
        return m.equivalentContent
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExpeditedQualityUpdateReference) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.QualityUpdateReference.GetFieldDeserializers()
    res["equivalentContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEquivalentContentOption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEquivalentContent(val.(*EquivalentContentOption))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ExpeditedQualityUpdateReference) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.QualityUpdateReference.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEquivalentContent() != nil {
        cast := (*m.GetEquivalentContent()).String()
        err = writer.WriteStringValue("equivalentContent", &cast)
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
func (m *ExpeditedQualityUpdateReference) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEquivalentContent sets the equivalentContent property value. Specifies other content to consider as equivalent. Supports a subset of the values for equivalentContentOption. Default value is latestSecurity. Possible values are: latestSecurity, unknownFutureValue.
func (m *ExpeditedQualityUpdateReference) SetEquivalentContent(value *EquivalentContentOption)() {
    if m != nil {
        m.equivalentContent = value
    }
}
