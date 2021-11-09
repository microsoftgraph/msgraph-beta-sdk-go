package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type RenameAction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The new name of the item.
    newName *string;
    // The previous name of the item.
    oldName *string;
}
// Instantiates a new renameAction and sets the default values.
func NewRenameAction()(*RenameAction) {
    m := &RenameAction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RenameAction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the newName property value. The new name of the item.
func (m *RenameAction) GetNewName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.newName
    }
}
// Gets the oldName property value. The previous name of the item.
func (m *RenameAction) GetOldName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oldName
    }
}
// The deserialization information for the current model
func (m *RenameAction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["newName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewName(val)
        }
        return nil
    }
    res["oldName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldName(val)
        }
        return nil
    }
    return res
}
func (m *RenameAction) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *RenameAction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("newName", m.GetNewName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oldName", m.GetOldName())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *RenameAction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the newName property value. The new name of the item.
// Parameters:
//  - value : Value to set for the newName property.
func (m *RenameAction) SetNewName(value *string)() {
    m.newName = value
}
// Sets the oldName property value. The previous name of the item.
// Parameters:
//  - value : Value to set for the oldName property.
func (m *RenameAction) SetOldName(value *string)() {
    m.oldName = value
}
