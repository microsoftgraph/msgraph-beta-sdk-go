package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSSingleSignOnExtension 
type MacOSSingleSignOnExtension struct {
    SingleSignOnExtension
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewMacOSSingleSignOnExtension instantiates a new MacOSSingleSignOnExtension and sets the default values.
func NewMacOSSingleSignOnExtension()(*MacOSSingleSignOnExtension) {
    m := &MacOSSingleSignOnExtension{
        SingleSignOnExtension: *NewSingleSignOnExtension(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
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
                mappingStr := *mappingValue
                switch mappingStr {
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
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSSingleSignOnExtension) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
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
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSSingleSignOnExtension) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
