package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FileClassificationRequest struct {
    Entity
    file []byte;
    sensitiveTypeIds []string;
}
func NewFileClassificationRequest()(*FileClassificationRequest) {
    m := &FileClassificationRequest{
        Entity: *NewEntity(),
    }
    return m
}
func (m *FileClassificationRequest) GetFile()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.file
    }
}
func (m *FileClassificationRequest) GetSensitiveTypeIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.sensitiveTypeIds
    }
}
func (m *FileClassificationRequest) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["file"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetFile(val)
        return nil
    }
    res["sensitiveTypeIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSensitiveTypeIds(res)
        return nil
    }
    return res
}
func (m *FileClassificationRequest) IsNil()(bool) {
    return m == nil
}
func (m *FileClassificationRequest) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("sensitiveTypeIds", m.GetSensitiveTypeIds())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *FileClassificationRequest) SetFile(value []byte)() {
    m.file = value
}
func (m *FileClassificationRequest) SetSensitiveTypeIds(value []string)() {
    m.sensitiveTypeIds = value
}
