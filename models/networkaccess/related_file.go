package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RelatedFile struct {
    RelatedResource
}
// NewRelatedFile instantiates a new RelatedFile and sets the default values.
func NewRelatedFile()(*RelatedFile) {
    m := &RelatedFile{
        RelatedResource: *NewRelatedResource(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.relatedFile"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRelatedFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelatedFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelatedFile(), nil
}
// GetDirectory gets the directory property value. The directory property
// returns a *string when successful
func (m *RelatedFile) GetDirectory()(*string) {
    val, err := m.GetBackingStore().Get("directory")
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
func (m *RelatedFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RelatedResource.GetFieldDeserializers()
    res["directory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectory(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["sizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInBytes(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *RelatedFile) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSizeInBytes gets the sizeInBytes property value. The sizeInBytes property
// returns a *int64 when successful
func (m *RelatedFile) GetSizeInBytes()(*int64) {
    val, err := m.GetBackingStore().Get("sizeInBytes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RelatedFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RelatedResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("directory", m.GetDirectory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sizeInBytes", m.GetSizeInBytes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDirectory sets the directory property value. The directory property
func (m *RelatedFile) SetDirectory(value *string)() {
    err := m.GetBackingStore().Set("directory", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *RelatedFile) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetSizeInBytes sets the sizeInBytes property value. The sizeInBytes property
func (m *RelatedFile) SetSizeInBytes(value *int64)() {
    err := m.GetBackingStore().Set("sizeInBytes", value)
    if err != nil {
        panic(err)
    }
}
type RelatedFileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RelatedResourceable
    GetDirectory()(*string)
    GetName()(*string)
    GetSizeInBytes()(*int64)
    SetDirectory(value *string)()
    SetName(value *string)()
    SetSizeInBytes(value *int64)()
}
