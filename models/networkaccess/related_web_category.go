package networkaccess

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RelatedWebCategory struct {
    RelatedResource
}
// NewRelatedWebCategory instantiates a new RelatedWebCategory and sets the default values.
func NewRelatedWebCategory()(*RelatedWebCategory) {
    m := &RelatedWebCategory{
        RelatedResource: *NewRelatedResource(),
    }
    odataTypeValue := "#microsoft.graph.networkaccess.relatedWebCategory"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateRelatedWebCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRelatedWebCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRelatedWebCategory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RelatedWebCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RelatedResource.GetFieldDeserializers()
    res["webCategoryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebCategoryName(val)
        }
        return nil
    }
    return res
}
// GetWebCategoryName gets the webCategoryName property value. The webCategoryName property
// returns a *string when successful
func (m *RelatedWebCategory) GetWebCategoryName()(*string) {
    val, err := m.GetBackingStore().Get("webCategoryName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *RelatedWebCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RelatedResource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("webCategoryName", m.GetWebCategoryName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetWebCategoryName sets the webCategoryName property value. The webCategoryName property
func (m *RelatedWebCategory) SetWebCategoryName(value *string)() {
    err := m.GetBackingStore().Set("webCategoryName", value)
    if err != nil {
        panic(err)
    }
}
type RelatedWebCategoryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RelatedResourceable
    GetWebCategoryName()(*string)
    SetWebCategoryName(value *string)()
}
