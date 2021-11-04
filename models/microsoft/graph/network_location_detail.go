package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type NetworkLocationDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Provides the name of the network used when signing in.
    networkNames []string;
    // Provides the type of network used when signing in. Possible values are: intranet, extranet, namedNetwork, trusted, unknownFutureValue.
    networkType *NetworkType;
}
// Instantiates a new networkLocationDetail and sets the default values.
func NewNetworkLocationDetail()(*NetworkLocationDetail) {
    m := &NetworkLocationDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkLocationDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the networkNames property value. Provides the name of the network used when signing in.
func (m *NetworkLocationDetail) GetNetworkNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.networkNames
    }
}
// Gets the networkType property value. Provides the type of network used when signing in. Possible values are: intranet, extranet, namedNetwork, trusted, unknownFutureValue.
func (m *NetworkLocationDetail) GetNetworkType()(*NetworkType) {
    if m == nil {
        return nil
    } else {
        return m.networkType
    }
}
// The deserialization information for the current model
func (m *NetworkLocationDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["networkNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetNetworkNames(res)
        return nil
    }
    res["networkType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkType)
        if err != nil {
            return err
        }
        cast := val.(NetworkType)
        m.SetNetworkType(&cast)
        return nil
    }
    return res
}
func (m *NetworkLocationDetail) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *NetworkLocationDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("networkNames", m.GetNetworkNames())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkType() != nil {
        cast := m.GetNetworkType().String()
        err := writer.WriteStringValue("networkType", &cast)
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
func (m *NetworkLocationDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the networkNames property value. Provides the name of the network used when signing in.
// Parameters:
//  - value : Value to set for the networkNames property.
func (m *NetworkLocationDetail) SetNetworkNames(value []string)() {
    m.networkNames = value
}
// Sets the networkType property value. Provides the type of network used when signing in. Possible values are: intranet, extranet, namedNetwork, trusted, unknownFutureValue.
// Parameters:
//  - value : Value to set for the networkType property.
func (m *NetworkLocationDetail) SetNetworkType(value *NetworkType)() {
    m.networkType = value
}
