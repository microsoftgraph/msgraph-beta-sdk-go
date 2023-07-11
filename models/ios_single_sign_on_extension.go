package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosSingleSignOnExtension an abstract base class for all iOS-specific single sign-on extension types.
type IosSingleSignOnExtension struct {
    SingleSignOnExtension
}
// NewIosSingleSignOnExtension instantiates a new iosSingleSignOnExtension and sets the default values.
func NewIosSingleSignOnExtension()(*IosSingleSignOnExtension) {
    m := &IosSingleSignOnExtension{
        SingleSignOnExtension: *NewSingleSignOnExtension(),
    }
    odataTypeValue := "#microsoft.graph.iosSingleSignOnExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.iosAzureAdSingleSignOnExtension":
                        return NewIosAzureAdSingleSignOnExtension(), nil
                    case "#microsoft.graph.iosCredentialSingleSignOnExtension":
                        return NewIosCredentialSingleSignOnExtension(), nil
                    case "#microsoft.graph.iosKerberosSingleSignOnExtension":
                        return NewIosKerberosSingleSignOnExtension(), nil
                    case "#microsoft.graph.iosRedirectSingleSignOnExtension":
                        return NewIosRedirectSingleSignOnExtension(), nil
                }
            }
        }
    }
    return NewIosSingleSignOnExtension(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SingleSignOnExtension.GetFieldDeserializers()
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosSingleSignOnExtension) GetOdataType()(*string) {
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
func (m *IosSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosSingleSignOnExtension) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// IosSingleSignOnExtensionable 
type IosSingleSignOnExtensionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SingleSignOnExtensionable
    GetOdataType()(*string)
    SetOdataType(value *string)()
}
