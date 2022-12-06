package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkPosition 
type WorkPosition struct {
    ItemFacet
    // Categories that the user has associated with this position.
    categories []string
    // Colleagues that are associated with this position.
    colleagues []RelatedPersonable
    // The detail property
    detail PositionDetailable
    // Denotes whether or not the position is current.
    isCurrent *bool
    // Contains detail of the user's manager in this position.
    manager RelatedPersonable
}
// NewWorkPosition instantiates a new WorkPosition and sets the default values.
func NewWorkPosition()(*WorkPosition) {
    m := &WorkPosition{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.workPosition";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWorkPositionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkPositionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkPosition(), nil
}
// GetCategories gets the categories property value. Categories that the user has associated with this position.
func (m *WorkPosition) GetCategories()([]string) {
    return m.categories
}
// GetColleagues gets the colleagues property value. Colleagues that are associated with this position.
func (m *WorkPosition) GetColleagues()([]RelatedPersonable) {
    return m.colleagues
}
// GetDetail gets the detail property value. The detail property
func (m *WorkPosition) GetDetail()(PositionDetailable) {
    return m.detail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkPosition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetCategories)
    res["colleagues"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateRelatedPersonFromDiscriminatorValue , m.SetColleagues)
    res["detail"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreatePositionDetailFromDiscriminatorValue , m.SetDetail)
    res["isCurrent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetIsCurrent)
    res["manager"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateRelatedPersonFromDiscriminatorValue , m.SetManager)
    return res
}
// GetIsCurrent gets the isCurrent property value. Denotes whether or not the position is current.
func (m *WorkPosition) GetIsCurrent()(*bool) {
    return m.isCurrent
}
// GetManager gets the manager property value. Contains detail of the user's manager in this position.
func (m *WorkPosition) GetManager()(RelatedPersonable) {
    return m.manager
}
// Serialize serializes information the current object
func (m *WorkPosition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ItemFacet.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCategories() != nil {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    if m.GetColleagues() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetColleagues())
        err = writer.WriteCollectionOfObjectValues("colleagues", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("detail", m.GetDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCurrent", m.GetIsCurrent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCategories sets the categories property value. Categories that the user has associated with this position.
func (m *WorkPosition) SetCategories(value []string)() {
    m.categories = value
}
// SetColleagues sets the colleagues property value. Colleagues that are associated with this position.
func (m *WorkPosition) SetColleagues(value []RelatedPersonable)() {
    m.colleagues = value
}
// SetDetail sets the detail property value. The detail property
func (m *WorkPosition) SetDetail(value PositionDetailable)() {
    m.detail = value
}
// SetIsCurrent sets the isCurrent property value. Denotes whether or not the position is current.
func (m *WorkPosition) SetIsCurrent(value *bool)() {
    m.isCurrent = value
}
// SetManager sets the manager property value. Contains detail of the user's manager in this position.
func (m *WorkPosition) SetManager(value RelatedPersonable)() {
    m.manager = value
}
