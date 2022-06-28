package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassificationError 
type ClassificationError struct {
    ClassifcationErrorBase
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The details property
    details []ClassifcationErrorBaseable
}
// NewClassificationError instantiates a new ClassificationError and sets the default values.
func NewClassificationError()(*ClassificationError) {
    m := &ClassificationError{
        ClassifcationErrorBase: *NewClassifcationErrorBase(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateClassificationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassificationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassificationError(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassificationError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDetails gets the details property value. The details property
func (m *ClassificationError) GetDetails()([]ClassifcationErrorBaseable) {
    if m == nil {
        return nil
    } else {
        return m.details
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassificationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ClassifcationErrorBase.GetFieldDeserializers()
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateClassifcationErrorBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ClassifcationErrorBaseable, len(val))
            for i, v := range val {
                res[i] = v.(ClassifcationErrorBaseable)
            }
            m.SetDetails(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ClassificationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ClassifcationErrorBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("details", cast)
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
func (m *ClassificationError) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDetails sets the details property value. The details property
func (m *ClassificationError) SetDetails(value []ClassifcationErrorBaseable)() {
    if m != nil {
        m.details = value
    }
}
