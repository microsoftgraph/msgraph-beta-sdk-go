// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProcessFileMetadata struct {
    ProcessContentMetadataBase
}
// NewProcessFileMetadata instantiates a new ProcessFileMetadata and sets the default values.
func NewProcessFileMetadata()(*ProcessFileMetadata) {
    m := &ProcessFileMetadata{
        ProcessContentMetadataBase: *NewProcessContentMetadataBase(),
    }
    odataTypeValue := "#microsoft.graph.processFileMetadata"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateProcessFileMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProcessFileMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProcessFileMetadata(), nil
}
// GetCustomProperties gets the customProperties property value. A dictionary containing custom metadata associated with the file, potentially extracted by the calling application.
// returns a CustomMetadataDictionaryable when successful
func (m *ProcessFileMetadata) GetCustomProperties()(CustomMetadataDictionaryable) {
    val, err := m.GetBackingStore().Get("customProperties")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CustomMetadataDictionaryable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProcessFileMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProcessContentMetadataBase.GetFieldDeserializers()
    res["customProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCustomMetadataDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomProperties(val.(CustomMetadataDictionaryable))
        }
        return nil
    }
    res["ownerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwnerId(val)
        }
        return nil
    }
    return res
}
// GetOwnerId gets the ownerId property value. The unique identifier (for example, Entra User ID or UPN) of the owner of the file.
// returns a *string when successful
func (m *ProcessFileMetadata) GetOwnerId()(*string) {
    val, err := m.GetBackingStore().Get("ownerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ProcessFileMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProcessContentMetadataBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("customProperties", m.GetCustomProperties())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ownerId", m.GetOwnerId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomProperties sets the customProperties property value. A dictionary containing custom metadata associated with the file, potentially extracted by the calling application.
func (m *ProcessFileMetadata) SetCustomProperties(value CustomMetadataDictionaryable)() {
    err := m.GetBackingStore().Set("customProperties", value)
    if err != nil {
        panic(err)
    }
}
// SetOwnerId sets the ownerId property value. The unique identifier (for example, Entra User ID or UPN) of the owner of the file.
func (m *ProcessFileMetadata) SetOwnerId(value *string)() {
    err := m.GetBackingStore().Set("ownerId", value)
    if err != nil {
        panic(err)
    }
}
type ProcessFileMetadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProcessContentMetadataBaseable
    GetCustomProperties()(CustomMetadataDictionaryable)
    GetOwnerId()(*string)
    SetCustomProperties(value CustomMetadataDictionaryable)()
    SetOwnerId(value *string)()
}
