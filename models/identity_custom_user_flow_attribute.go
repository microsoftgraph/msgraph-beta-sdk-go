package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityCustomUserFlowAttribute 
type IdentityCustomUserFlowAttribute struct {
    IdentityUserFlowAttribute
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewIdentityCustomUserFlowAttribute instantiates a new IdentityCustomUserFlowAttribute and sets the default values.
func NewIdentityCustomUserFlowAttribute()(*IdentityCustomUserFlowAttribute) {
    m := &IdentityCustomUserFlowAttribute{
        IdentityUserFlowAttribute: *NewIdentityUserFlowAttribute(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIdentityCustomUserFlowAttributeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityCustomUserFlowAttributeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityCustomUserFlowAttribute(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentityCustomUserFlowAttribute) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityCustomUserFlowAttribute) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityUserFlowAttribute.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *IdentityCustomUserFlowAttribute) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityUserFlowAttribute.Serialize(writer)
    if err != nil {
        return err
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
func (m *IdentityCustomUserFlowAttribute) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
