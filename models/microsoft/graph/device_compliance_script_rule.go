package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceComplianceScriptRule struct {
    additionalData map[string]interface{};
    dataType *DataType;
    deviceComplianceScriptRuleDataType *DeviceComplianceScriptRuleDataType;
    deviceComplianceScriptRulOperator *DeviceComplianceScriptRulOperator;
    operand *string;
    operator *Operator;
    settingName *string;
}
func NewDeviceComplianceScriptRule()(*DeviceComplianceScriptRule) {
    m := &DeviceComplianceScriptRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceComplianceScriptRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceComplianceScriptRule) GetDataType()(*DataType) {
    if m == nil {
        return nil
    } else {
        return m.dataType
    }
}
func (m *DeviceComplianceScriptRule) GetDeviceComplianceScriptRuleDataType()(*DeviceComplianceScriptRuleDataType) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScriptRuleDataType
    }
}
func (m *DeviceComplianceScriptRule) GetDeviceComplianceScriptRulOperator()(*DeviceComplianceScriptRulOperator) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScriptRulOperator
    }
}
func (m *DeviceComplianceScriptRule) GetOperand()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operand
    }
}
func (m *DeviceComplianceScriptRule) GetOperator()(*Operator) {
    if m == nil {
        return nil
    } else {
        return m.operator
    }
}
func (m *DeviceComplianceScriptRule) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
func (m *DeviceComplianceScriptRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataType)
        if err != nil {
            return err
        }
        cast := val.(DataType)
        m.SetDataType(&cast)
        return nil
    }
    res["deviceComplianceScriptRuleDataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceScriptRuleDataType)
        if err != nil {
            return err
        }
        cast := val.(DeviceComplianceScriptRuleDataType)
        m.SetDeviceComplianceScriptRuleDataType(&cast)
        return nil
    }
    res["deviceComplianceScriptRulOperator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceScriptRulOperator)
        if err != nil {
            return err
        }
        cast := val.(DeviceComplianceScriptRulOperator)
        m.SetDeviceComplianceScriptRulOperator(&cast)
        return nil
    }
    res["operand"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperand(val)
        return nil
    }
    res["operator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperator)
        if err != nil {
            return err
        }
        cast := val.(Operator)
        m.SetOperator(&cast)
        return nil
    }
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSettingName(val)
        return nil
    }
    return res
}
func (m *DeviceComplianceScriptRule) IsNil()(bool) {
    return m == nil
}
func (m *DeviceComplianceScriptRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDataType() != nil {
        cast := m.GetDataType().String()
        err := writer.WriteStringValue("dataType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScriptRuleDataType() != nil {
        cast := m.GetDeviceComplianceScriptRuleDataType().String()
        err := writer.WriteStringValue("deviceComplianceScriptRuleDataType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScriptRulOperator() != nil {
        cast := m.GetDeviceComplianceScriptRulOperator().String()
        err := writer.WriteStringValue("deviceComplianceScriptRulOperator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operand", m.GetOperand())
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := m.GetOperator().String()
        err := writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingName", m.GetSettingName())
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
func (m *DeviceComplianceScriptRule) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceComplianceScriptRule) SetDataType(value *DataType)() {
    m.dataType = value
}
func (m *DeviceComplianceScriptRule) SetDeviceComplianceScriptRuleDataType(value *DeviceComplianceScriptRuleDataType)() {
    m.deviceComplianceScriptRuleDataType = value
}
func (m *DeviceComplianceScriptRule) SetDeviceComplianceScriptRulOperator(value *DeviceComplianceScriptRulOperator)() {
    m.deviceComplianceScriptRulOperator = value
}
func (m *DeviceComplianceScriptRule) SetOperand(value *string)() {
    m.operand = value
}
func (m *DeviceComplianceScriptRule) SetOperator(value *Operator)() {
    m.operator = value
}
func (m *DeviceComplianceScriptRule) SetSettingName(value *string)() {
    m.settingName = value
}
