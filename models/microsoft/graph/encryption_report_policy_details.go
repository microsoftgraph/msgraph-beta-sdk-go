package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EncryptionReportPolicyDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Policy Id for Encryption Report
    policyId *string;
    // Policy Name for Encryption Report
    policyName *string;
}
// Instantiates a new encryptionReportPolicyDetails and sets the default values.
func NewEncryptionReportPolicyDetails()(*EncryptionReportPolicyDetails) {
    m := &EncryptionReportPolicyDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EncryptionReportPolicyDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the policyId property value. Policy Id for Encryption Report
func (m *EncryptionReportPolicyDetails) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
// Gets the policyName property value. Policy Name for Encryption Report
func (m *EncryptionReportPolicyDetails) GetPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyName
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *EncryptionReportPolicyDetails) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the policyId property value. Policy Id for Encryption Report
// Parameters:
//  - value : Value to set for the policyId property.
func (m *EncryptionReportPolicyDetails) SetPolicyId(value *string)() {
    m.policyId = value
}
// Sets the policyName property value. Policy Name for Encryption Report
// Parameters:
//  - value : Value to set for the policyName property.
func (m *EncryptionReportPolicyDetails) SetPolicyName(value *string)() {
    m.policyName = value
}
