package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ServiceUpdateMessageViewpoint struct {
    additionalData map[string]interface{};
    isArchived *bool;
    isFavorited *bool;
    isRead *bool;
}
func NewServiceUpdateMessageViewpoint()(*ServiceUpdateMessageViewpoint) {
    m := &ServiceUpdateMessageViewpoint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ServiceUpdateMessageViewpoint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ServiceUpdateMessageViewpoint) GetIsArchived()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isArchived
    }
}
func (m *ServiceUpdateMessageViewpoint) GetIsFavorited()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFavorited
    }
}
func (m *ServiceUpdateMessageViewpoint) GetIsRead()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRead
    }
}
func (m *ServiceUpdateMessageViewpoint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isArchived"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsArchived(val)
        return nil
    }
    res["isFavorited"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsFavorited(val)
        return nil
    }
    res["isRead"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRead(val)
        return nil
    }
    return res
}
func (m *ServiceUpdateMessageViewpoint) IsNil()(bool) {
    return m == nil
}
func (m *ServiceUpdateMessageViewpoint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isArchived", m.GetIsArchived())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isFavorited", m.GetIsFavorited())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRead", m.GetIsRead())
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
func (m *ServiceUpdateMessageViewpoint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ServiceUpdateMessageViewpoint) SetIsArchived(value *bool)() {
    m.isArchived = value
}
func (m *ServiceUpdateMessageViewpoint) SetIsFavorited(value *bool)() {
    m.isFavorited = value
}
func (m *ServiceUpdateMessageViewpoint) SetIsRead(value *bool)() {
    m.isRead = value
}
