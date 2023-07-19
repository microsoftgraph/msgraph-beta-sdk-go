package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdgeHomeButtonOpensCustomURL show the home button; clicking the home button loads a specific URL.
type EdgeHomeButtonOpensCustomURL struct {
    EdgeHomeButtonConfiguration
}
// NewEdgeHomeButtonOpensCustomURL instantiates a new edgeHomeButtonOpensCustomURL and sets the default values.
func NewEdgeHomeButtonOpensCustomURL()(*EdgeHomeButtonOpensCustomURL) {
    m := &EdgeHomeButtonOpensCustomURL{
        EdgeHomeButtonConfiguration: *NewEdgeHomeButtonConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.edgeHomeButtonOpensCustomURL"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEdgeHomeButtonOpensCustomURLFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdgeHomeButtonOpensCustomURLFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdgeHomeButtonOpensCustomURL(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdgeHomeButtonOpensCustomURL) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EdgeHomeButtonConfiguration.GetFieldDeserializers()
    res["homeButtonCustomURL"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHomeButtonCustomURL(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetHomeButtonCustomURL gets the homeButtonCustomURL property value. The specific URL to load.
func (m *EdgeHomeButtonOpensCustomURL) GetHomeButtonCustomURL()(*string) {
    val, err := m.GetBackingStore().Get("homeButtonCustomURL")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EdgeHomeButtonOpensCustomURL) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdgeHomeButtonOpensCustomURL) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EdgeHomeButtonConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("homeButtonCustomURL", m.GetHomeButtonCustomURL())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHomeButtonCustomURL sets the homeButtonCustomURL property value. The specific URL to load.
func (m *EdgeHomeButtonOpensCustomURL) SetHomeButtonCustomURL(value *string)() {
    err := m.GetBackingStore().Set("homeButtonCustomURL", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EdgeHomeButtonOpensCustomURL) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// EdgeHomeButtonOpensCustomURLable 
type EdgeHomeButtonOpensCustomURLable interface {
    EdgeHomeButtonConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHomeButtonCustomURL()(*string)
    GetOdataType()(*string)
    SetHomeButtonCustomURL(value *string)()
    SetOdataType(value *string)()
}
