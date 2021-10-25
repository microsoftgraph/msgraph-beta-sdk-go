package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type StatusBase struct {
    additionalData map[string]interface{};
    status *ProvisioningResult;
}
func NewStatusBase()(*StatusBase) {
    m := &StatusBase{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *StatusBase) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *StatusBase) GetStatus()(*ProvisioningResult) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *StatusBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningResult)
        if err != nil {
            return err
        }
        cast := val.(ProvisioningResult)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *StatusBase) IsNil()(bool) {
    return m == nil
}
func (m *StatusBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *StatusBase) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *StatusBase) SetStatus(value *ProvisioningResult)() {
    m.status = value
}
