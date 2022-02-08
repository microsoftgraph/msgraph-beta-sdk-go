package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// KerberosSignOnSettings 
type KerberosSignOnSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
    kerberosServicePrincipalName *string;
    // The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
    kerberosSignOnMappingAttributeType *KerberosSignOnMappingAttributeType;
}
// NewKerberosSignOnSettings instantiates a new kerberosSignOnSettings and sets the default values.
func NewKerberosSignOnSettings()(*KerberosSignOnSettings) {
    m := &KerberosSignOnSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KerberosSignOnSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetKerberosServicePrincipalName gets the kerberosServicePrincipalName property value. The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
func (m *KerberosSignOnSettings) GetKerberosServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kerberosServicePrincipalName
    }
}
// GetKerberosSignOnMappingAttributeType gets the kerberosSignOnMappingAttributeType property value. The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
func (m *KerberosSignOnSettings) GetKerberosSignOnMappingAttributeType()(*KerberosSignOnMappingAttributeType) {
    if m == nil {
        return nil
    } else {
        return m.kerberosSignOnMappingAttributeType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *KerberosSignOnSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["kerberosServicePrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosServicePrincipalName(val)
        }
        return nil
    }
    res["kerberosSignOnMappingAttributeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseKerberosSignOnMappingAttributeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKerberosSignOnMappingAttributeType(val.(*KerberosSignOnMappingAttributeType))
        }
        return nil
    }
    return res
}
func (m *KerberosSignOnSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *KerberosSignOnSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("kerberosServicePrincipalName", m.GetKerberosServicePrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetKerberosSignOnMappingAttributeType() != nil {
        cast := (*m.GetKerberosSignOnMappingAttributeType()).String()
        err := writer.WriteStringValue("kerberosSignOnMappingAttributeType", &cast)
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
func (m *KerberosSignOnSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetKerberosServicePrincipalName sets the kerberosServicePrincipalName property value. The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
func (m *KerberosSignOnSettings) SetKerberosServicePrincipalName(value *string)() {
    if m != nil {
        m.kerberosServicePrincipalName = value
    }
}
// SetKerberosSignOnMappingAttributeType sets the kerberosSignOnMappingAttributeType property value. The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
func (m *KerberosSignOnSettings) SetKerberosSignOnMappingAttributeType(value *KerberosSignOnMappingAttributeType)() {
    if m != nil {
        m.kerberosSignOnMappingAttributeType = value
    }
}
