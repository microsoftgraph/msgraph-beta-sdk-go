package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SynchronizationJobSubject 
type SynchronizationJobSubject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The links property
    links SynchronizationLinkedObjectsable
    // The identifier of an object to which a synchronizationJob  is to be applied.
    objectId *string
    // The type of the object to which a synchronizationJob  is to be applied.
    objectTypeName *string
}
// NewSynchronizationJobSubject instantiates a new synchronizationJobSubject and sets the default values.
func NewSynchronizationJobSubject()(*SynchronizationJobSubject) {
    m := &SynchronizationJobSubject{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSynchronizationJobSubjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSynchronizationJobSubjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSynchronizationJobSubject(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationJobSubject) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJobSubject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSynchronizationLinkedObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(SynchronizationLinkedObjectsable))
        }
        return nil
    }
    res["objectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectId(val)
        }
        return nil
    }
    res["objectTypeName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectTypeName(val)
        }
        return nil
    }
    return res
}
// GetLinks gets the links property value. The links property
func (m *SynchronizationJobSubject) GetLinks()(SynchronizationLinkedObjectsable) {
    if m == nil {
        return nil
    } else {
        return m.links
    }
}
// GetObjectId gets the objectId property value. The identifier of an object to which a synchronizationJob  is to be applied.
func (m *SynchronizationJobSubject) GetObjectId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectId
    }
}
// GetObjectTypeName gets the objectTypeName property value. The type of the object to which a synchronizationJob  is to be applied.
func (m *SynchronizationJobSubject) GetObjectTypeName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectTypeName
    }
}
// Serialize serializes information the current object
func (m *SynchronizationJobSubject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("links", m.GetLinks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("objectId", m.GetObjectId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("objectTypeName", m.GetObjectTypeName())
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
func (m *SynchronizationJobSubject) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLinks sets the links property value. The links property
func (m *SynchronizationJobSubject) SetLinks(value SynchronizationLinkedObjectsable)() {
    if m != nil {
        m.links = value
    }
}
// SetObjectId sets the objectId property value. The identifier of an object to which a synchronizationJob  is to be applied.
func (m *SynchronizationJobSubject) SetObjectId(value *string)() {
    if m != nil {
        m.objectId = value
    }
}
// SetObjectTypeName sets the objectTypeName property value. The type of the object to which a synchronizationJob  is to be applied.
func (m *SynchronizationJobSubject) SetObjectTypeName(value *string)() {
    if m != nil {
        m.objectTypeName = value
    }
}
