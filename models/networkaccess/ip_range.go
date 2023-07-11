package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpRange 
type IpRange struct {
    RuleDestination
    // The OdataType property
    OdataType *string
}
// NewIpRange instantiates a new ipRange and sets the default values.
func NewIpRange()(*IpRange) {
    m := &IpRange{
        RuleDestination: *NewRuleDestination(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.ipRange"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIpRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIpRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpRange(), nil
}
// GetBeginAddress gets the beginAddress property value. Specifies the starting IP address of the IP range.
func (m *IpRange) GetBeginAddress()(*string) {
    val, err := m.GetBackingStore().Get("beginAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndAddress gets the endAddress property value. Specifies the ending IP address of the IP range.
func (m *IpRange) GetEndAddress()(*string) {
    val, err := m.GetBackingStore().Get("endAddress")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RuleDestination.GetFieldDeserializers()
    res["beginAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBeginAddress(val)
        }
        return nil
    }
    res["endAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndAddress(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *IpRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RuleDestination.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("beginAddress", m.GetBeginAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("endAddress", m.GetEndAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBeginAddress sets the beginAddress property value. Specifies the starting IP address of the IP range.
func (m *IpRange) SetBeginAddress(value *string)() {
    err := m.GetBackingStore().Set("beginAddress", value)
    if err != nil {
        panic(err)
    }
}
// SetEndAddress sets the endAddress property value. Specifies the ending IP address of the IP range.
func (m *IpRange) SetEndAddress(value *string)() {
    err := m.GetBackingStore().Set("endAddress", value)
    if err != nil {
        panic(err)
    }
}
// IpRangeable 
type IpRangeable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RuleDestinationable
    GetBeginAddress()(*string)
    GetEndAddress()(*string)
    SetBeginAddress(value *string)()
    SetEndAddress(value *string)()
}
