package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Admin struct {
    additionalData map[string]interface{};
    serviceAnnouncement *ServiceAnnouncement;
    windows *Windows;
}
func NewAdmin()(*Admin) {
    m := &Admin{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Admin) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Admin) GetServiceAnnouncement()(*ServiceAnnouncement) {
    if m == nil {
        return nil
    } else {
        return m.serviceAnnouncement
    }
}
func (m *Admin) GetWindows()(*Windows) {
    if m == nil {
        return nil
    } else {
        return m.windows
    }
}
func (m *Admin) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["serviceAnnouncement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceAnnouncement() })
        if err != nil {
            return err
        }
        m.SetServiceAnnouncement(val.(*ServiceAnnouncement))
        return nil
    }
    res["windows"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindows() })
        if err != nil {
            return err
        }
        m.SetWindows(val.(*Windows))
        return nil
    }
    return res
}
func (m *Admin) IsNil()(bool) {
    return m == nil
}
func (m *Admin) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("serviceAnnouncement", m.GetServiceAnnouncement())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("windows", m.GetWindows())
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
func (m *Admin) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Admin) SetServiceAnnouncement(value *ServiceAnnouncement)() {
    m.serviceAnnouncement = value
}
func (m *Admin) SetWindows(value *Windows)() {
    m.windows = value
}
