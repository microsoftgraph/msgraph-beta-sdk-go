package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DirectorySizeQuota struct {
    additionalData map[string]interface{};
    total *int32;
    used *int32;
}
func NewDirectorySizeQuota()(*DirectorySizeQuota) {
    m := &DirectorySizeQuota{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DirectorySizeQuota) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DirectorySizeQuota) GetTotal()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.total
    }
}
func (m *DirectorySizeQuota) GetUsed()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.used
    }
}
func (m *DirectorySizeQuota) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["total"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetTotal(val)
        return nil
    }
    res["used"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetUsed(val)
        return nil
    }
    return res
}
func (m *DirectorySizeQuota) IsNil()(bool) {
    return m == nil
}
func (m *DirectorySizeQuota) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("total", m.GetTotal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("used", m.GetUsed())
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
func (m *DirectorySizeQuota) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DirectorySizeQuota) SetTotal(value *int32)() {
    m.total = value
}
func (m *DirectorySizeQuota) SetUsed(value *int32)() {
    m.used = value
}
