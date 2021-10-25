package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SynchronizationError struct {
    additionalData map[string]interface{};
    code *string;
    message *string;
    tenantActionable *bool;
}
func NewSynchronizationError()(*SynchronizationError) {
    m := &SynchronizationError{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *SynchronizationError) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *SynchronizationError) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
func (m *SynchronizationError) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
func (m *SynchronizationError) GetTenantActionable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.tenantActionable
    }
}
func (m *SynchronizationError) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCode(val)
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    res["tenantActionable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetTenantActionable(val)
        return nil
    }
    return res
}
func (m *SynchronizationError) IsNil()(bool) {
    return m == nil
}
func (m *SynchronizationError) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("tenantActionable", m.GetTenantActionable())
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
func (m *SynchronizationError) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *SynchronizationError) SetCode(value *string)() {
    m.code = value
}
func (m *SynchronizationError) SetMessage(value *string)() {
    m.message = value
}
func (m *SynchronizationError) SetTenantActionable(value *bool)() {
    m.tenantActionable = value
}
