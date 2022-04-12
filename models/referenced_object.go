package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReferencedObject 
type ReferencedObject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of the referenced object. Must match one of the objects in the directory definition.
    referencedObjectName *string
    // Currently not supported. Name of the property in the referenced object, the value for which is used as the reference.
    referencedProperty *string
}
// NewReferencedObject instantiates a new referencedObject and sets the default values.
func NewReferencedObject()(*ReferencedObject) {
    m := &ReferencedObject{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateReferencedObjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateReferencedObjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewReferencedObject(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ReferencedObject) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ReferencedObject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["referencedObjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferencedObjectName(val)
        }
        return nil
    }
    res["referencedProperty"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReferencedProperty(val)
        }
        return nil
    }
    return res
}
// GetReferencedObjectName gets the referencedObjectName property value. Name of the referenced object. Must match one of the objects in the directory definition.
func (m *ReferencedObject) GetReferencedObjectName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referencedObjectName
    }
}
// GetReferencedProperty gets the referencedProperty property value. Currently not supported. Name of the property in the referenced object, the value for which is used as the reference.
func (m *ReferencedObject) GetReferencedProperty()(*string) {
    if m == nil {
        return nil
    } else {
        return m.referencedProperty
    }
}
// Serialize serializes information the current object
func (m *ReferencedObject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("referencedObjectName", m.GetReferencedObjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("referencedProperty", m.GetReferencedProperty())
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
func (m *ReferencedObject) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReferencedObjectName sets the referencedObjectName property value. Name of the referenced object. Must match one of the objects in the directory definition.
func (m *ReferencedObject) SetReferencedObjectName(value *string)() {
    if m != nil {
        m.referencedObjectName = value
    }
}
// SetReferencedProperty sets the referencedProperty property value. Currently not supported. Name of the property in the referenced object, the value for which is used as the reference.
func (m *ReferencedObject) SetReferencedProperty(value *string)() {
    if m != nil {
        m.referencedProperty = value
    }
}
