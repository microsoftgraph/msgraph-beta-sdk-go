package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Bundle struct {
    additionalData map[string]interface{};
    album *Album;
    childCount *int32;
}
func NewBundle()(*Bundle) {
    m := &Bundle{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Bundle) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Bundle) GetAlbum()(*Album) {
    if m == nil {
        return nil
    } else {
        return m.album
    }
}
func (m *Bundle) GetChildCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.childCount
    }
}
func (m *Bundle) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["album"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAlbum() })
        if err != nil {
            return err
        }
        m.SetAlbum(val.(*Album))
        return nil
    }
    res["childCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetChildCount(val)
        return nil
    }
    return res
}
func (m *Bundle) IsNil()(bool) {
    return m == nil
}
func (m *Bundle) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("album", m.GetAlbum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("childCount", m.GetChildCount())
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
func (m *Bundle) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Bundle) SetAlbum(value *Album)() {
    m.album = value
}
func (m *Bundle) SetChildCount(value *int32)() {
    m.childCount = value
}
