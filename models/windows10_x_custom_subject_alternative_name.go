package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XCustomSubjectAlternativeName base Profile Type for Authentication Certificates (SCEP or PFX Create)
type Windows10XCustomSubjectAlternativeName struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Custom SAN Name
    name *string
    // Subject Alternative Name Options.
    sanType *SubjectAlternativeNameType
}
// NewWindows10XCustomSubjectAlternativeName instantiates a new windows10XCustomSubjectAlternativeName and sets the default values.
func NewWindows10XCustomSubjectAlternativeName()(*Windows10XCustomSubjectAlternativeName) {
    m := &Windows10XCustomSubjectAlternativeName{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWindows10XCustomSubjectAlternativeNameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XCustomSubjectAlternativeNameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10XCustomSubjectAlternativeName(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Windows10XCustomSubjectAlternativeName) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XCustomSubjectAlternativeName) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["sanType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectAlternativeNameType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSanType(val.(*SubjectAlternativeNameType))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Custom SAN Name
func (m *Windows10XCustomSubjectAlternativeName) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetSanType gets the sanType property value. Subject Alternative Name Options.
func (m *Windows10XCustomSubjectAlternativeName) GetSanType()(*SubjectAlternativeNameType) {
    if m == nil {
        return nil
    } else {
        return m.sanType
    }
}
// Serialize serializes information the current object
func (m *Windows10XCustomSubjectAlternativeName) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetSanType() != nil {
        cast := (*m.GetSanType()).String()
        err := writer.WriteStringValue("sanType", &cast)
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
func (m *Windows10XCustomSubjectAlternativeName) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. Custom SAN Name
func (m *Windows10XCustomSubjectAlternativeName) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetSanType sets the sanType property value. Subject Alternative Name Options.
func (m *Windows10XCustomSubjectAlternativeName) SetSanType(value *SubjectAlternativeNameType)() {
    if m != nil {
        m.sanType = value
    }
}
