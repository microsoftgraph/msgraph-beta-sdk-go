package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceComplianceScriptRule 
type DeviceComplianceScriptRule struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
    dataType *DataType;
    // Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
    deviceComplianceScriptRuleDataType *DeviceComplianceScriptRuleDataType;
    // Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
    deviceComplianceScriptRulOperator *DeviceComplianceScriptRulOperator;
    // Operand specified in the rule.
    operand *string;
    // Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
    operator *Operator;
    // Setting name specified in the rule.
    settingName *string;
}
// NewDeviceComplianceScriptRule instantiates a new deviceComplianceScriptRule and sets the default values.
func NewDeviceComplianceScriptRule()(*DeviceComplianceScriptRule) {
    m := &DeviceComplianceScriptRule{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptRule) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDataType gets the dataType property value. Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
func (m *DeviceComplianceScriptRule) GetDataType()(*DataType) {
    if m == nil {
        return nil
    } else {
        return m.dataType
    }
}
// GetDeviceComplianceScriptRuleDataType gets the deviceComplianceScriptRuleDataType property value. Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
func (m *DeviceComplianceScriptRule) GetDeviceComplianceScriptRuleDataType()(*DeviceComplianceScriptRuleDataType) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScriptRuleDataType
    }
}
// GetDeviceComplianceScriptRulOperator gets the deviceComplianceScriptRulOperator property value. Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
func (m *DeviceComplianceScriptRule) GetDeviceComplianceScriptRulOperator()(*DeviceComplianceScriptRulOperator) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceScriptRulOperator
    }
}
// GetOperand gets the operand property value. Operand specified in the rule.
func (m *DeviceComplianceScriptRule) GetOperand()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operand
    }
}
// GetOperator gets the operator property value. Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
func (m *DeviceComplianceScriptRule) GetOperator()(*Operator) {
    if m == nil {
        return nil
    } else {
        return m.operator
    }
}
// GetSettingName gets the settingName property value. Setting name specified in the rule.
func (m *DeviceComplianceScriptRule) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceScriptRule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDataType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataType(val.(*DataType))
        }
        return nil
    }
    res["deviceComplianceScriptRuleDataType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceScriptRuleDataType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceScriptRuleDataType(val.(*DeviceComplianceScriptRuleDataType))
        }
        return nil
    }
    res["deviceComplianceScriptRulOperator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceComplianceScriptRulOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceScriptRulOperator(val.(*DeviceComplianceScriptRulOperator))
        }
        return nil
    }
    res["operand"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperand(val)
        }
        return nil
    }
    res["operator"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*Operator))
        }
        return nil
    }
    res["settingName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    return res
}
func (m *DeviceComplianceScriptRule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceComplianceScriptRule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDataType() != nil {
        cast := (*m.GetDataType()).String()
        err := writer.WriteStringValue("dataType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScriptRuleDataType() != nil {
        cast := (*m.GetDeviceComplianceScriptRuleDataType()).String()
        err := writer.WriteStringValue("deviceComplianceScriptRuleDataType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceComplianceScriptRulOperator() != nil {
        cast := (*m.GetDeviceComplianceScriptRulOperator()).String()
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
        cast := (*m.GetOperator()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceComplianceScriptRule) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDataType sets the dataType property value. Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
func (m *DeviceComplianceScriptRule) SetDataType(value *DataType)() {
    if m != nil {
        m.dataType = value
    }
}
// SetDeviceComplianceScriptRuleDataType sets the deviceComplianceScriptRuleDataType property value. Data type specified in the rule. Possible values are: none, boolean, int64, double, string, dateTime, version, base64, xml, booleanArray, int64Array, doubleArray, stringArray, dateTimeArray, versionArray.
func (m *DeviceComplianceScriptRule) SetDeviceComplianceScriptRuleDataType(value *DeviceComplianceScriptRuleDataType)() {
    if m != nil {
        m.deviceComplianceScriptRuleDataType = value
    }
}
// SetDeviceComplianceScriptRulOperator sets the deviceComplianceScriptRulOperator property value. Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
func (m *DeviceComplianceScriptRule) SetDeviceComplianceScriptRulOperator(value *DeviceComplianceScriptRulOperator)() {
    if m != nil {
        m.deviceComplianceScriptRulOperator = value
    }
}
// SetOperand sets the operand property value. Operand specified in the rule.
func (m *DeviceComplianceScriptRule) SetOperand(value *string)() {
    if m != nil {
        m.operand = value
    }
}
// SetOperator sets the operator property value. Operator specified in the rule. Possible values are: none, and, or, isEquals, notEquals, greaterThan, lessThan, between, notBetween, greaterEquals, lessEquals, dayTimeBetween, beginsWith, notBeginsWith, endsWith, notEndsWith, contains, notContains, allOf, oneOf, noneOf, setEquals, orderedSetEquals, subsetOf, excludesAll.
func (m *DeviceComplianceScriptRule) SetOperator(value *Operator)() {
    if m != nil {
        m.operator = value
    }
}
// SetSettingName sets the settingName property value. Setting name specified in the rule.
func (m *DeviceComplianceScriptRule) SetSettingName(value *string)() {
    if m != nil {
        m.settingName = value
    }
}
