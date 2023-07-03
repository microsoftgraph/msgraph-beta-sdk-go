package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSingleSignOnExtension 
type MacOSSingleSignOnExtension struct {
    SingleSignOnExtension
}
// NewMacOSSingleSignOnExtension instantiates a new macOSSingleSignOnExtension and sets the default values.
func NewMacOSSingleSignOnExtension()(*MacOSSingleSignOnExtension) {
    m := &MacOSSingleSignOnExtension{
        SingleSignOnExtension: *NewSingleSignOnExtension(),
    }
    odataTypeValue := "#microsoft.graph.macOSSingleSignOnExtension"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSSingleSignOnExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSSingleSignOnExtensionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.macOSAzureAdSingleSignOnExtension":
                        return NewMacOSAzureAdSingleSignOnExtension(), nil
                    case "#microsoft.graph.macOSCredentialSingleSignOnExtension":
                        return NewMacOSCredentialSingleSignOnExtension(), nil
                    case "#microsoft.graph.macOSKerberosSingleSignOnExtension":
                        return NewMacOSKerberosSingleSignOnExtension(), nil
                    case "#microsoft.graph.macOSRedirectSingleSignOnExtension":
                        return NewMacOSRedirectSingleSignOnExtension(), nil
                }
            }
        }
    }
    return NewMacOSSingleSignOnExtension(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSSingleSignOnExtension) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SingleSignOnExtension.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *MacOSSingleSignOnExtension) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SingleSignOnExtension.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// MacOSSingleSignOnExtensionable 
type MacOSSingleSignOnExtensionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SingleSignOnExtensionable
}
