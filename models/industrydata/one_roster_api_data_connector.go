package industrydata

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OneRosterApiDataConnector struct {
    ApiDataConnector
}
// NewOneRosterApiDataConnector instantiates a new OneRosterApiDataConnector and sets the default values.
func NewOneRosterApiDataConnector()(*OneRosterApiDataConnector) {
    m := &OneRosterApiDataConnector{
        ApiDataConnector: *NewApiDataConnector(),
    }
    odataTypeValue := "#microsoft.graph.industryData.oneRosterApiDataConnector"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOneRosterApiDataConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOneRosterApiDataConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOneRosterApiDataConnector(), nil
}
// GetApiVersion gets the apiVersion property value. The API version of the OneRoster source. Example: 1.1, 1.2
// returns a *string when successful
func (m *OneRosterApiDataConnector) GetApiVersion()(*string) {
    val, err := m.GetBackingStore().Get("apiVersion")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OneRosterApiDataConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ApiDataConnector.GetFieldDeserializers()
    res["apiVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiVersion(val)
        }
        return nil
    }
    res["isContactsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContactsEnabled(val)
        }
        return nil
    }
    res["isDemographicsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDemographicsEnabled(val)
        }
        return nil
    }
    res["isFlagsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFlagsEnabled(val)
        }
        return nil
    }
    return res
}
// GetIsContactsEnabled gets the isContactsEnabled property value. Indicates whether the user specified to import optional contacts data.
// returns a *bool when successful
func (m *OneRosterApiDataConnector) GetIsContactsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isContactsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsDemographicsEnabled gets the isDemographicsEnabled property value. Indicates whether the user specified to import optional demographics data.
// returns a *bool when successful
func (m *OneRosterApiDataConnector) GetIsDemographicsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isDemographicsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetIsFlagsEnabled gets the isFlagsEnabled property value. Indicates whether the user specified to import optional flags data.
// returns a *bool when successful
func (m *OneRosterApiDataConnector) GetIsFlagsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isFlagsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *OneRosterApiDataConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ApiDataConnector.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("apiVersion", m.GetApiVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isContactsEnabled", m.GetIsContactsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDemographicsEnabled", m.GetIsDemographicsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFlagsEnabled", m.GetIsFlagsEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApiVersion sets the apiVersion property value. The API version of the OneRoster source. Example: 1.1, 1.2
func (m *OneRosterApiDataConnector) SetApiVersion(value *string)() {
    err := m.GetBackingStore().Set("apiVersion", value)
    if err != nil {
        panic(err)
    }
}
// SetIsContactsEnabled sets the isContactsEnabled property value. Indicates whether the user specified to import optional contacts data.
func (m *OneRosterApiDataConnector) SetIsContactsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isContactsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsDemographicsEnabled sets the isDemographicsEnabled property value. Indicates whether the user specified to import optional demographics data.
func (m *OneRosterApiDataConnector) SetIsDemographicsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isDemographicsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsFlagsEnabled sets the isFlagsEnabled property value. Indicates whether the user specified to import optional flags data.
func (m *OneRosterApiDataConnector) SetIsFlagsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isFlagsEnabled", value)
    if err != nil {
        panic(err)
    }
}
type OneRosterApiDataConnectorable interface {
    ApiDataConnectorable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiVersion()(*string)
    GetIsContactsEnabled()(*bool)
    GetIsDemographicsEnabled()(*bool)
    GetIsFlagsEnabled()(*bool)
    SetApiVersion(value *string)()
    SetIsContactsEnabled(value *bool)()
    SetIsDemographicsEnabled(value *bool)()
    SetIsFlagsEnabled(value *bool)()
}
