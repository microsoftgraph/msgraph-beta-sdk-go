package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DlpActionInfo struct {
    action *DlpAction;
    additionalData map[string]interface{};
}
func NewDlpActionInfo()(*DlpActionInfo) {
    m := &DlpActionInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DlpActionInfo) GetAction()(*DlpAction) {
    if m == nil {
        return nil
    } else {
        return m.action
    }
}
func (m *DlpActionInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DlpActionInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["action"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDlpAction)
        if err != nil {
            return err
        }
        cast := val.(DlpAction)
        m.SetAction(&cast)
        return nil
    }
    return res
}
func (m *DlpActionInfo) IsNil()(bool) {
    return m == nil
}
func (m *DlpActionInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAction() != nil {
        cast := m.GetAction().String()
        err := writer.WriteStringValue("action", &cast)
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
func (m *DlpActionInfo) SetAction(value *DlpAction)() {
    m.action = value
}
func (m *DlpActionInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
