package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserConfiguration 
type UserConfiguration struct {
    Entity
    // 
    binaryData []byte;
}
// NewUserConfiguration instantiates a new userConfiguration and sets the default values.
func NewUserConfiguration()(*UserConfiguration) {
    m := &UserConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// GetBinaryData gets the binaryData property value. 
func (m *UserConfiguration) GetBinaryData()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.binaryData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["binaryData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBinaryData(val)
        }
        return nil
    }
    return res
}
func (m *UserConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("binaryData", m.GetBinaryData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBinaryData sets the binaryData property value. 
func (m *UserConfiguration) SetBinaryData(value []byte)() {
    if m != nil {
        m.binaryData = value
    }
}
