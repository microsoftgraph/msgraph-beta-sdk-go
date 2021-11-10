package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type KerberosSignOnSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
    kerberosServicePrincipalName *string;
    // The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
    kerberosSignOnMappingAttributeType *KerberosSignOnMappingAttributeType;
}
// Instantiates a new kerberosSignOnSettings and sets the default values.
func NewKerberosSignOnSettings()(*KerberosSignOnSettings) {
    m := &KerberosSignOnSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *KerberosSignOnSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the kerberosServicePrincipalName property value. The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
func (m *KerberosSignOnSettings) GetKerberosServicePrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.kerberosServicePrincipalName
    }
}
// Gets the kerberosSignOnMappingAttributeType property value. The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
func (m *KerberosSignOnSettings) GetKerberosSignOnMappingAttributeType()(*KerberosSignOnMappingAttributeType) {
    if m == nil {
        return nil
    } else {
        return m.kerberosSignOnMappingAttributeType
    }
}
// The deserialization information for the current model
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
            cast := val.(KerberosSignOnMappingAttributeType)
            m.SetKerberosSignOnMappingAttributeType(&cast)
        }
        return nil
    }
    return res
}
func (m *KerberosSignOnSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *KerberosSignOnSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("kerberosServicePrincipalName", m.GetKerberosServicePrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetKerberosSignOnMappingAttributeType() != nil {
        cast := m.GetKerberosSignOnMappingAttributeType().String()
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *KerberosSignOnSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the kerberosServicePrincipalName property value. The Internal Application SPN of the application server. This SPN needs to be in the list of services to which the connector can present delegated credentials.
// Parameters:
//  - value : Value to set for the kerberosServicePrincipalName property.
func (m *KerberosSignOnSettings) SetKerberosServicePrincipalName(value *string)() {
    m.kerberosServicePrincipalName = value
}
// Sets the kerberosSignOnMappingAttributeType property value. The Delegated Login Identity for the connector to use on behalf of your users. For more information, see Working with different on-premises and cloud identities . Possible values are: userPrincipalName, onPremisesUserPrincipalName, userPrincipalUsername, onPremisesUserPrincipalUsername, onPremisesSAMAccountName.
// Parameters:
//  - value : Value to set for the kerberosSignOnMappingAttributeType property.
func (m *KerberosSignOnSettings) SetKerberosSignOnMappingAttributeType(value *KerberosSignOnMappingAttributeType)() {
    m.kerberosSignOnMappingAttributeType = value
}
