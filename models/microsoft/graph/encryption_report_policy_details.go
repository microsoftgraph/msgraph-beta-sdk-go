package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type EncryptionReportPolicyDetails struct {
    additionalData map[string]interface{};
    policyId *string;
    policyName *string;
}
func NewEncryptionReportPolicyDetails()(*EncryptionReportPolicyDetails) {
    m := &EncryptionReportPolicyDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *EncryptionReportPolicyDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *EncryptionReportPolicyDetails) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
func (m *EncryptionReportPolicyDetails) GetPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyName
    }
}
func (m *EncryptionReportPolicyDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["policyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyId(val)
        return nil
    }
    res["policyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyName(val)
        return nil
    }
    return res
}
func (m *EncryptionReportPolicyDetails) IsNil()(bool) {
    return m == nil
}
func (m *EncryptionReportPolicyDetails) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("policyId", m.GetPolicyId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("policyName", m.GetPolicyName())
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
func (m *EncryptionReportPolicyDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *EncryptionReportPolicyDetails) SetPolicyId(value *string)() {
    m.policyId = value
}
func (m *EncryptionReportPolicyDetails) SetPolicyName(value *string)() {
    m.policyName = value
}
