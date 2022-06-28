package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeliveryOptimizationGroupIdSourceOptions 
type DeliveryOptimizationGroupIdSourceOptions struct {
    DeliveryOptimizationGroupIdSource
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Set this policy to restrict peer selection to a specific source. Possible values are: notConfigured, adSite, authenticatedDomainSid, dhcpUserOption, dnsSuffix.
    groupIdSourceOption *DeliveryOptimizationGroupIdOptionsType
}
// NewDeliveryOptimizationGroupIdSourceOptions instantiates a new DeliveryOptimizationGroupIdSourceOptions and sets the default values.
func NewDeliveryOptimizationGroupIdSourceOptions()(*DeliveryOptimizationGroupIdSourceOptions) {
    m := &DeliveryOptimizationGroupIdSourceOptions{
        DeliveryOptimizationGroupIdSource: *NewDeliveryOptimizationGroupIdSource(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeliveryOptimizationGroupIdSourceOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeliveryOptimizationGroupIdSourceOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeliveryOptimizationGroupIdSourceOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeliveryOptimizationGroupIdSourceOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeliveryOptimizationGroupIdSourceOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeliveryOptimizationGroupIdSource.GetFieldDeserializers()
    res["groupIdSourceOption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeliveryOptimizationGroupIdOptionsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupIdSourceOption(val.(*DeliveryOptimizationGroupIdOptionsType))
        }
        return nil
    }
    return res
}
// GetGroupIdSourceOption gets the groupIdSourceOption property value. Set this policy to restrict peer selection to a specific source. Possible values are: notConfigured, adSite, authenticatedDomainSid, dhcpUserOption, dnsSuffix.
func (m *DeliveryOptimizationGroupIdSourceOptions) GetGroupIdSourceOption()(*DeliveryOptimizationGroupIdOptionsType) {
    if m == nil {
        return nil
    } else {
        return m.groupIdSourceOption
    }
}
// Serialize serializes information the current object
func (m *DeliveryOptimizationGroupIdSourceOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeliveryOptimizationGroupIdSource.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetGroupIdSourceOption() != nil {
        cast := (*m.GetGroupIdSourceOption()).String()
        err = writer.WriteStringValue("groupIdSourceOption", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeliveryOptimizationGroupIdSourceOptions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGroupIdSourceOption sets the groupIdSourceOption property value. Set this policy to restrict peer selection to a specific source. Possible values are: notConfigured, adSite, authenticatedDomainSid, dhcpUserOption, dnsSuffix.
func (m *DeliveryOptimizationGroupIdSourceOptions) SetGroupIdSourceOption(value *DeliveryOptimizationGroupIdOptionsType)() {
    if m != nil {
        m.groupIdSourceOption = value
    }
}
