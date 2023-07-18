package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CategoryTemplate 
type CategoryTemplate struct {
    FilePlanDescriptorTemplate
}
// NewCategoryTemplate instantiates a new categoryTemplate and sets the default values.
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CategoryTemplate) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
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
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
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
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CategoryTemplate) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
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
    GetOdataType()(*string)
    GetSubCategories()([]SubCategoryTemplateable)
    SetOdataType(value *string)()
    SetSubCategories(value []SubCategoryTemplateable)()
}
