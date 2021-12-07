package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeleteAction 
type DeleteAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The name of the item that was deleted.
    name *string;
    // File or Folder, depending on the type of the deleted item.
    objectType *string;
}
// NewDeleteAction instantiates a new deleteAction and sets the default values.
func NewDeleteAction()(*DeleteAction) {
    m := &DeleteAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeleteAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetName gets the name property value. The name of the item that was deleted.
func (m *DeleteAction) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetObjectType gets the objectType property value. File or Folder, depending on the type of the deleted item.
func (m *DeleteAction) GetObjectType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.objectType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeleteAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["objectType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectType(val)
        }
        return nil
    }
    return res
}
func (m *DeleteAction) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeleteAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("objectType", m.GetObjectType())
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
func (m *DeleteAction) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetName sets the name property value. The name of the item that was deleted.
func (m *DeleteAction) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetObjectType sets the objectType property value. File or Folder, depending on the type of the deleted item.
func (m *DeleteAction) SetObjectType(value *string)() {
    if m != nil {
        m.objectType = value
    }
}
