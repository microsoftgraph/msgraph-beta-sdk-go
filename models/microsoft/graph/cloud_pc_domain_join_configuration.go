package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type CloudPcDomainJoinConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The on-premises connection ID that matches the virtual network IT admins want the provisioning policy to use when they create Cloud PCs. You can use this property in both domain join types: Azure AD joined or Hybrid Azure AD joined. If you enter an onPremisesConnectionId, leave regionName as empty.
    onPremisesConnectionId *string;
    // The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs. The underlying virtual network will be created and managed by the Windows 365 service. This can only be entered if the IT admin chooses Azure AD joined as the domain join type. If you enter a regionName, leave onPremisesConnectionId as empty.
    regionName *string;
    // Specifies how the provisioned Cloud PC will be joined to Azure AD. If you choose the hybridAzureADJoin type, only provide a value for the onPremisesConnectionId property and leave regionName as empty. If you choose the azureADJoin type, provide a value for either onPremisesConnectionId or regionName. The possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
    type_escaped *CloudPcDomainJoinType;
}
// Instantiates a new cloudPcDomainJoinConfiguration and sets the default values.
func NewCloudPcDomainJoinConfiguration()(*CloudPcDomainJoinConfiguration) {
    m := &CloudPcDomainJoinConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CloudPcDomainJoinConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the onPremisesConnectionId property value. The on-premises connection ID that matches the virtual network IT admins want the provisioning policy to use when they create Cloud PCs. You can use this property in both domain join types: Azure AD joined or Hybrid Azure AD joined. If you enter an onPremisesConnectionId, leave regionName as empty.
func (m *CloudPcDomainJoinConfiguration) GetOnPremisesConnectionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.onPremisesConnectionId
    }
}
// Gets the regionName property value. The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs. The underlying virtual network will be created and managed by the Windows 365 service. This can only be entered if the IT admin chooses Azure AD joined as the domain join type. If you enter a regionName, leave onPremisesConnectionId as empty.
func (m *CloudPcDomainJoinConfiguration) GetRegionName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.regionName
    }
}
// Gets the type_escaped property value. Specifies how the provisioned Cloud PC will be joined to Azure AD. If you choose the hybridAzureADJoin type, only provide a value for the onPremisesConnectionId property and leave regionName as empty. If you choose the azureADJoin type, provide a value for either onPremisesConnectionId or regionName. The possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
func (m *CloudPcDomainJoinConfiguration) GetType_escaped()(*CloudPcDomainJoinType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *CloudPcDomainJoinConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["onPremisesConnectionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesConnectionId(val)
        }
        return nil
    }
    res["regionName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegionName(val)
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudPcDomainJoinType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CloudPcDomainJoinType)
            m.SetType_escaped(&cast)
        }
        return nil
    }
    return res
}
func (m *CloudPcDomainJoinConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *CloudPcDomainJoinConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("onPremisesConnectionId", m.GetOnPremisesConnectionId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("regionName", m.GetRegionName())
        if err != nil {
            return err
        }
    }
    if m.GetType_escaped() != nil {
        cast := m.GetType_escaped().String()
        err := writer.WriteStringValue("type_escaped", &cast)
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
func (m *CloudPcDomainJoinConfiguration) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the onPremisesConnectionId property value. The on-premises connection ID that matches the virtual network IT admins want the provisioning policy to use when they create Cloud PCs. You can use this property in both domain join types: Azure AD joined or Hybrid Azure AD joined. If you enter an onPremisesConnectionId, leave regionName as empty.
// Parameters:
//  - value : Value to set for the onPremisesConnectionId property.
func (m *CloudPcDomainJoinConfiguration) SetOnPremisesConnectionId(value *string)() {
    m.onPremisesConnectionId = value
}
// Sets the regionName property value. The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs. The underlying virtual network will be created and managed by the Windows 365 service. This can only be entered if the IT admin chooses Azure AD joined as the domain join type. If you enter a regionName, leave onPremisesConnectionId as empty.
// Parameters:
//  - value : Value to set for the regionName property.
func (m *CloudPcDomainJoinConfiguration) SetRegionName(value *string)() {
    m.regionName = value
}
// Sets the type_escaped property value. Specifies how the provisioned Cloud PC will be joined to Azure AD. If you choose the hybridAzureADJoin type, only provide a value for the onPremisesConnectionId property and leave regionName as empty. If you choose the azureADJoin type, provide a value for either onPremisesConnectionId or regionName. The possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *CloudPcDomainJoinConfiguration) SetType_escaped(value *CloudPcDomainJoinType)() {
    m.type_escaped = value
}
