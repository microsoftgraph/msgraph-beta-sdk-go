package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EncryptContent 
type EncryptContent struct {
    LabelActionBase
    // The OdataType property
    OdataType *string
}
// NewEncryptContent instantiates a new encryptContent and sets the default values.
func NewEncryptContent()(*EncryptContent) {
    m := &EncryptContent{
        LabelActionBase: *NewLabelActionBase(),
    }
    odataTypeValue := "#microsoft.graph.encryptContent"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEncryptContentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEncryptContentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.encryptWithTemplate":
                        return NewEncryptWithTemplate(), nil
                    case "#microsoft.graph.encryptWithUserDefinedRights":
                        return NewEncryptWithUserDefinedRights(), nil
                }
            }
        }
    }
    return NewEncryptContent(), nil
}
// GetEncryptWith gets the encryptWith property value. The encryptWith property
func (m *EncryptContent) GetEncryptWith()(*EncryptWith) {
    val, err := m.GetBackingStore().Get("encryptWith")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*EncryptWith)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EncryptContent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LabelActionBase.GetFieldDeserializers()
    res["encryptWith"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEncryptWith)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptWith(val.(*EncryptWith))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EncryptContent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LabelActionBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEncryptWith() != nil {
        cast := (*m.GetEncryptWith()).String()
        err = writer.WriteStringValue("encryptWith", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEncryptWith sets the encryptWith property value. The encryptWith property
func (m *EncryptContent) SetEncryptWith(value *EncryptWith)() {
    err := m.GetBackingStore().Set("encryptWith", value)
    if err != nil {
        panic(err)
    }
}
// EncryptContentable 
type EncryptContentable interface {
    LabelActionBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEncryptWith()(*EncryptWith)
    SetEncryptWith(value *EncryptWith)()
}
