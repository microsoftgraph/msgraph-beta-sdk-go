package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EncryptionReportPolicyDetails 
type EncryptionReportPolicyDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Policy Id for Encryption Report
    policyId *string;
    // Policy Name for Encryption Report
    policyName *string;
}
// NewEncryptionReportPolicyDetails instantiates a new encryptionReportPolicyDetails and sets the default values.
func NewEncryptionReportPolicyDetails()(*EncryptionReportPolicyDetails) {
    m := &EncryptionReportPolicyDetails{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EncryptionReportPolicyDetails) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPolicyId gets the policyId property value. Policy Id for Encryption Report
func (m *EncryptionReportPolicyDetails) GetPolicyId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyId
    }
}
// GetPolicyName gets the policyName property value. Policy Name for Encryption Report
func (m *EncryptionReportPolicyDetails) GetPolicyName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EncryptionReportPolicyDetails) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["policyId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyId(val)
        }
        return nil
    }
    res["policyName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyName(val)
        }
        return nil
    }
    return res
}
func (m *EncryptionReportPolicyDetails) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EncryptionReportPolicyDetails) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetPolicyId sets the policyId property value. Policy Id for Encryption Report
func (m *EncryptionReportPolicyDetails) SetPolicyId(value *string)() {
    if m != nil {
        m.policyId = value
    }
}
// SetPolicyName sets the policyName property value. Policy Name for Encryption Report
func (m *EncryptionReportPolicyDetails) SetPolicyName(value *string)() {
    if m != nil {
        m.policyName = value
    }
}
