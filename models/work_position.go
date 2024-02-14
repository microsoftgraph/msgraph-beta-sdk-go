package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WorkPosition struct {
    ItemFacet
}
// NewWorkPosition instantiates a new WorkPosition and sets the default values.
func NewWorkPosition()(*WorkPosition) {
    m := &WorkPosition{
        ItemFacet: *NewItemFacet(),
    }
    odataTypeValue := "#microsoft.graph.workPosition"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWorkPositionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWorkPositionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkPosition(), nil
}
// GetCategories gets the categories property value. Categories that the user has associated with this position.
// returns a []string when successful
func (m *WorkPosition) GetCategories()([]string) {
    val, err := m.GetBackingStore().Get("categories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetColleagues gets the colleagues property value. Colleagues that are associated with this position.
// returns a []RelatedPersonable when successful
func (m *WorkPosition) GetColleagues()([]RelatedPersonable) {
    val, err := m.GetBackingStore().Get("colleagues")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]RelatedPersonable)
    }
    return nil
}
// GetDetail gets the detail property value. The detail property
// returns a PositionDetailable when successful
func (m *WorkPosition) GetDetail()(PositionDetailable) {
    val, err := m.GetBackingStore().Get("detail")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(PositionDetailable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WorkPosition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ItemFacet.GetFieldDeserializers()
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["colleagues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRelatedPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedPersonable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(RelatedPersonable)
                }
            }
            m.SetColleagues(res)
        }
        return nil
    }
    res["detail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePositionDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetail(val.(PositionDetailable))
        }
        return nil
    }
    res["isCurrent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCurrent(val)
        }
        return nil
    }
    res["manager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRelatedPersonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(RelatedPersonable))
        }
        return nil
    }
    return res
}
// GetIsCurrent gets the isCurrent property value. Denotes whether or not the position is current.
// returns a *bool when successful
func (m *WorkPosition) GetIsCurrent()(*bool) {
    val, err := m.GetBackingStore().Get("isCurrent")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetManager gets the manager property value. Contains detail of the user's manager in this position.
// returns a RelatedPersonable when successful
func (m *WorkPosition) GetManager()(RelatedPersonable) {
    val, err := m.GetBackingStore().Get("manager")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(RelatedPersonable)
    }
    return nil
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
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetColleagues()))
        for i, v := range m.GetColleagues() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
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
    err := m.GetBackingStore().Set("categories", value)
    if err != nil {
        panic(err)
    }
}
// SetColleagues sets the colleagues property value. Colleagues that are associated with this position.
func (m *WorkPosition) SetColleagues(value []RelatedPersonable)() {
    err := m.GetBackingStore().Set("colleagues", value)
    if err != nil {
        panic(err)
    }
}
// SetDetail sets the detail property value. The detail property
func (m *WorkPosition) SetDetail(value PositionDetailable)() {
    err := m.GetBackingStore().Set("detail", value)
    if err != nil {
        panic(err)
    }
}
// SetIsCurrent sets the isCurrent property value. Denotes whether or not the position is current.
func (m *WorkPosition) SetIsCurrent(value *bool)() {
    err := m.GetBackingStore().Set("isCurrent", value)
    if err != nil {
        panic(err)
    }
}
// SetManager sets the manager property value. Contains detail of the user's manager in this position.
func (m *WorkPosition) SetManager(value RelatedPersonable)() {
    err := m.GetBackingStore().Set("manager", value)
    if err != nil {
        panic(err)
    }
}
type WorkPositionable interface {
    ItemFacetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategories()([]string)
    GetColleagues()([]RelatedPersonable)
    GetDetail()(PositionDetailable)
    GetIsCurrent()(*bool)
    GetManager()(RelatedPersonable)
    SetCategories(value []string)()
    SetColleagues(value []RelatedPersonable)()
    SetDetail(value PositionDetailable)()
    SetIsCurrent(value *bool)()
    SetManager(value RelatedPersonable)()
}
