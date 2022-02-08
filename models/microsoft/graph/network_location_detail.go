package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// NetworkLocationDetail 
type NetworkLocationDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Provides the name of the network used when signing in.
    networkNames []string;
    // Provides the type of network used when signing in. Possible values are: intranet, extranet, namedNetwork, trusted, unknownFutureValue.
    networkType *NetworkType;
}
// NewNetworkLocationDetail instantiates a new networkLocationDetail and sets the default values.
func NewNetworkLocationDetail()(*NetworkLocationDetail) {
    m := &NetworkLocationDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkLocationDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetNetworkNames gets the networkNames property value. Provides the name of the network used when signing in.
func (m *NetworkLocationDetail) GetNetworkNames()([]string) {
    if m == nil {
        return nil
    } else {
        return m.networkNames
    }
}
// GetNetworkType gets the networkType property value. Provides the type of network used when signing in. Possible values are: intranet, extranet, namedNetwork, trusted, unknownFutureValue.
func (m *NetworkLocationDetail) GetNetworkType()(*NetworkType) {
    if m == nil {
        return nil
    } else {
        return m.networkType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NetworkLocationDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["networkNames"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNetworkNames(res)
        }
        return nil
    }
    res["networkType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseNetworkType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkType(val.(*NetworkType))
        }
        return nil
    }
    return res
}
func (m *NetworkLocationDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *NetworkLocationDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetNetworkNames() != nil {
        err := writer.WriteCollectionOfStringValues("networkNames", m.GetNetworkNames())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkType() != nil {
        cast := (*m.GetNetworkType()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkLocationDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetNetworkNames sets the networkNames property value. Provides the name of the network used when signing in.
func (m *NetworkLocationDetail) SetNetworkNames(value []string)() {
    if m != nil {
        m.networkNames = value
    }
}
// SetNetworkType sets the networkType property value. Provides the type of network used when signing in. Possible values are: intranet, extranet, namedNetwork, trusted, unknownFutureValue.
func (m *NetworkLocationDetail) SetNetworkType(value *NetworkType)() {
    if m != nil {
        m.networkType = value
    }
}
