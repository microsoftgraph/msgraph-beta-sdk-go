package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CategoryTemplate 
type CategoryTemplate struct {
    FilePlanDescriptorTemplate
}
// NewCategoryTemplate instantiates a new CategoryTemplate and sets the default values.
func NewCategoryTemplate()(*CategoryTemplate) {
    m := &CategoryTemplate{
        FilePlanDescriptorTemplate: *NewFilePlanDescriptorTemplate(),
    }
    return m
}
// CreateCategoryTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCategoryTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCategoryTemplate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CategoryTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FilePlanDescriptorTemplate.GetFieldDeserializers()
    res["subCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubCategoryTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubCategoryTemplateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SubCategoryTemplateable)
                }
            }
            m.SetSubCategories(res)
        }
        return nil
    }
    return res
}
// GetSubCategories gets the subCategories property value. Represents all subcategories under a particular category.
func (m *CategoryTemplate) GetSubCategories()([]SubCategoryTemplateable) {
    val, err := m.GetBackingStore().Get("subCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SubCategoryTemplateable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *CategoryTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FilePlanDescriptorTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSubCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubCategories()))
        for i, v := range m.GetSubCategories() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("subCategories", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSubCategories sets the subCategories property value. Represents all subcategories under a particular category.
func (m *CategoryTemplate) SetSubCategories(value []SubCategoryTemplateable)() {
    err := m.GetBackingStore().Set("subCategories", value)
    if err != nil {
        panic(err)
    }
}
// CategoryTemplateable 
type CategoryTemplateable interface {
    FilePlanDescriptorTemplateable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetSubCategories()([]SubCategoryTemplateable)
    SetSubCategories(value []SubCategoryTemplateable)()
}
