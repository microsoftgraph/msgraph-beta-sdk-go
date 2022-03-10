package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceComplianceScriptRuleable 
type DeviceComplianceScriptRuleable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDataType()(*DataType)
    GetDeviceComplianceScriptRuleDataType()(*DeviceComplianceScriptRuleDataType)
    GetDeviceComplianceScriptRulOperator()(*DeviceComplianceScriptRulOperator)
    GetOperand()(*string)
    GetOperator()(*Operator)
    GetSettingName()(*string)
    SetDataType(value *DataType)()
    SetDeviceComplianceScriptRuleDataType(value *DeviceComplianceScriptRuleDataType)()
    SetDeviceComplianceScriptRulOperator(value *DeviceComplianceScriptRulOperator)()
    SetOperand(value *string)()
    SetOperator(value *Operator)()
    SetSettingName(value *string)()
}
