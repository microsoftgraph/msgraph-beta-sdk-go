package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NetworkLocationDetail 
type NetworkLocationDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Provides the name of the network used when signing in.
    networkNames []string
    // Provides the type of network used when signing in. Possible values are: intranet, extranet, namedNetwork, trusted, unknownFutureValue.
    networkType *NetworkType
    // The OdataType property
    odataType *string
}
// NewNetworkLocationDetail instantiates a new networkLocationDetail and sets the default values.
func NewNetworkLocationDetail()(*NetworkLocationDetail) {
    m := &NetworkLocationDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.networkLocationDetail";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateNetworkLocationDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNetworkLocationDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNetworkLocationDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkLocationDetail) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NetworkLocationDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["networkNames"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetNetworkNames)
    res["networkType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseNetworkType , m.SetNetworkType)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetNetworkNames gets the networkNames property value. Provides the name of the network used when signing in.
func (m *NetworkLocationDetail) GetNetworkNames()([]string) {
    return m.networkNames
}
// GetNetworkType gets the networkType property value. Provides the type of network used when signing in. Possible values are: intranet, extranet, namedNetwork, trusted, unknownFutureValue.
func (m *NetworkLocationDetail) GetNetworkType()(*NetworkType) {
    return m.networkType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NetworkLocationDetail) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *NetworkLocationDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
    m.additionalData = value
}
// SetNetworkNames sets the networkNames property value. Provides the name of the network used when signing in.
func (m *NetworkLocationDetail) SetNetworkNames(value []string)() {
    m.networkNames = value
}
// SetNetworkType sets the networkType property value. Provides the type of network used when signing in. Possible values are: intranet, extranet, namedNetwork, trusted, unknownFutureValue.
func (m *NetworkLocationDetail) SetNetworkType(value *NetworkType)() {
    m.networkType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NetworkLocationDetail) SetOdataType(value *string)() {
    m.odataType = value
}
