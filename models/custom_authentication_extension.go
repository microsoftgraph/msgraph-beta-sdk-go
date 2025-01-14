package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CustomAuthenticationExtension struct {
    CustomCalloutExtension
}
// NewCustomAuthenticationExtension instantiates a new CustomAuthenticationExtension and sets the default values.
func NewCustomAuthenticationExtension()(*CustomAuthenticationExtension) {
    m := &CustomAuthenticationExtension{
        CustomCalloutExtension: *NewCustomCalloutExtension(),
    }
    odataTypeValue := "#microsoft.graph.customAuthenticationExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCustomAuthenticationExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCustomAuthenticationExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.onAttributeCollectionStartCustomExtension":
                        return NewOnAttributeCollectionStartCustomExtension(), nil
                    case "#microsoft.graph.onAttributeCollectionSubmitCustomExtension":
                        return NewOnAttributeCollectionSubmitCustomExtension(), nil
                    case "#microsoft.graph.onOtpSendCustomExtension":
                        return NewOnOtpSendCustomExtension(), nil
                    case "#microsoft.graph.onTokenIssuanceStartCustomExtension":
                        return NewOnTokenIssuanceStartCustomExtension(), nil
                }
            }
        }
    }
    return NewCustomAuthenticationExtension(), nil
}
// GetBehaviorOnError gets the behaviorOnError property value. The behaviour on error for the custom authentication extension.
// returns a CustomExtensionBehaviorOnErrorable when successful
func (m *CustomAuthenticationExtension) GetBehaviorOnError()(CustomExtensionBehaviorOnErrorable) {
    val, err := m.GetBackingStore().Get("behaviorOnError")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomExtensionBehaviorOnErrorable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CustomAuthenticationExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CustomCalloutExtension.GetFieldDeserializers()
    res["behaviorOnError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomExtensionBehaviorOnErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBehaviorOnError(val.(CustomExtensionBehaviorOnErrorable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CustomAuthenticationExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CustomCalloutExtension.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("behaviorOnError", m.GetBehaviorOnError())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBehaviorOnError sets the behaviorOnError property value. The behaviour on error for the custom authentication extension.
func (m *CustomAuthenticationExtension) SetBehaviorOnError(value CustomExtensionBehaviorOnErrorable)() {
    err := m.GetBackingStore().Set("behaviorOnError", value)
    if err != nil {
        panic(err)
    }
}
type CustomAuthenticationExtensionable interface {
    CustomCalloutExtensionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBehaviorOnError()(CustomExtensionBehaviorOnErrorable)
    SetBehaviorOnError(value CustomExtensionBehaviorOnErrorable)()
}
