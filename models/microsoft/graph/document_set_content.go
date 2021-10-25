package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DocumentSetContent struct {
    additionalData map[string]interface{};
    contentType *ContentTypeInfo;
    fileName *string;
    folderName *string;
}
func NewDocumentSetContent()(*DocumentSetContent) {
    m := &DocumentSetContent{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DocumentSetContent) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DocumentSetContent) GetContentType()(*ContentTypeInfo) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
func (m *DocumentSetContent) GetFileName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.fileName
    }
}
func (m *DocumentSetContent) GetFolderName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.folderName
    }
}
func (m *DocumentSetContent) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewContentTypeInfo() })
        if err != nil {
            return err
        }
        m.SetContentType(val.(*ContentTypeInfo))
        return nil
    }
    res["fileName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFileName(val)
        return nil
    }
    res["folderName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetFolderName(val)
        return nil
    }
    return res
}
func (m *DocumentSetContent) IsNil()(bool) {
    return m == nil
}
func (m *DocumentSetContent) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("folderName", m.GetFolderName())
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
func (m *DocumentSetContent) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DocumentSetContent) SetContentType(value *ContentTypeInfo)() {
    m.contentType = value
}
func (m *DocumentSetContent) SetFileName(value *string)() {
    m.fileName = value
}
func (m *DocumentSetContent) SetFolderName(value *string)() {
    m.folderName = value
}
