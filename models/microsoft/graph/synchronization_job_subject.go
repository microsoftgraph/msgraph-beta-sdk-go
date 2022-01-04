package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SynchronizationJobSubject 
type SynchronizationJobSubject struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    links *SynchronizationLinkedObjects;
    // The identifier of an object to which a synchronizationJob  is to be applied.
    objectId *string;
    // The type of the object to which a synchronizationJob  is to be applied.
    objectTypeName *string;
}
// NewSynchronizationJobSubject instantiates a new synchronizationJobSubject and sets the default values.
func NewSynchronizationJobSubject()(*SynchronizationJobSubject) {
    m := &SynchronizationJobSubject{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SynchronizationJobSubject) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetLinks gets the links property value. 
func (m *SynchronizationJobSubject) GetLinks()(*SynchronizationLinkedObjects) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *SynchronizationJobSubject) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["links"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSynchronizationLinkedObjects() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(*SynchronizationLinkedObjects))
        }
        return nil
    }
    res["objectId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectId(val)
        }
        return nil
    }
    res["objectTypeName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *SynchronizationJobSubject) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SynchronizationJobSubject) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetLinks sets the links property value. 
func (m *SynchronizationJobSubject) SetLinks(value *SynchronizationLinkedObjects)() {
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
