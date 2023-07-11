package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppliedCategory 
type AppliedCategory struct {
    FilePlanDescriptorBase
}
// NewAppliedCategory instantiates a new appliedCategory and sets the default values.
func NewAppliedCategory()(*AppliedCategory) {
    m := &AppliedCategory{
        FilePlanDescriptorBase: *NewFilePlanDescriptorBase(),
    }
    return m
}
// CreateAppliedCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppliedCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppliedCategory(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppliedCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.FilePlanDescriptorBase.GetFieldDeserializers()
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
    res["subCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSubCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubCategory(val.(SubCategoryable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AppliedCategory) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSubCategory gets the subCategory property value. Represents the file plan descriptor for a subcategory under a specific category, which has been assigned to a particular retention label.
func (m *AppliedCategory) GetSubCategory()(SubCategoryable) {
    val, err := m.GetBackingStore().Get("subCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SubCategoryable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AppliedCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.FilePlanDescriptorBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("subCategory", m.GetSubCategory())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AppliedCategory) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetSubCategory sets the subCategory property value. Represents the file plan descriptor for a subcategory under a specific category, which has been assigned to a particular retention label.
func (m *AppliedCategory) SetSubCategory(value SubCategoryable)() {
    err := m.GetBackingStore().Set("subCategory", value)
    if err != nil {
        panic(err)
    }
}
// AppliedCategoryable 
type AppliedCategoryable interface {
    FilePlanDescriptorBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetSubCategory()(SubCategoryable)
    SetOdataType(value *string)()
    SetSubCategory(value SubCategoryable)()
}
