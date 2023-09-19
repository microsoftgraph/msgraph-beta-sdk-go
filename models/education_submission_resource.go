package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSubmissionResource 
type EducationSubmissionResource struct {
    Entity
}
// NewEducationSubmissionResource instantiates a new educationSubmissionResource and sets the default values.
func NewEducationSubmissionResource()(*EducationSubmissionResource) {
    m := &EducationSubmissionResource{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationSubmissionResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSubmissionResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSubmissionResource(), nil
}
// GetAssignmentResourceUrl gets the assignmentResourceUrl property value. Pointer to the assignment from which the resource was copied. If the value is null, the student uploaded the resource.
func (m *EducationSubmissionResource) GetAssignmentResourceUrl()(*string) {
    val, err := m.GetBackingStore().Get("assignmentResourceUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDependentResources gets the dependentResources property value. The dependentResources property
func (m *EducationSubmissionResource) GetDependentResources()([]EducationSubmissionResourceable) {
    val, err := m.GetBackingStore().Get("dependentResources")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]EducationSubmissionResourceable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSubmissionResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentResourceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentResourceUrl(val)
        }
        return nil
    }
    res["dependentResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSubmissionResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSubmissionResourceable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(EducationSubmissionResourceable)
                }
            }
            m.SetDependentResources(res)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(EducationResourceable))
        }
        return nil
    }
    return res
}
// GetResource gets the resource property value. Resource object.
func (m *EducationSubmissionResource) GetResource()(EducationResourceable) {
    val, err := m.GetBackingStore().Get("resource")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EducationResourceable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EducationSubmissionResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignmentResourceUrl", m.GetAssignmentResourceUrl())
        if err != nil {
            return err
        }
    }
    if m.GetDependentResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDependentResources()))
        for i, v := range m.GetDependentResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dependentResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentResourceUrl sets the assignmentResourceUrl property value. Pointer to the assignment from which the resource was copied. If the value is null, the student uploaded the resource.
func (m *EducationSubmissionResource) SetAssignmentResourceUrl(value *string)() {
    err := m.GetBackingStore().Set("assignmentResourceUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetDependentResources sets the dependentResources property value. The dependentResources property
func (m *EducationSubmissionResource) SetDependentResources(value []EducationSubmissionResourceable)() {
    err := m.GetBackingStore().Set("dependentResources", value)
    if err != nil {
        panic(err)
    }
}
// SetResource sets the resource property value. Resource object.
func (m *EducationSubmissionResource) SetResource(value EducationResourceable)() {
    err := m.GetBackingStore().Set("resource", value)
    if err != nil {
        panic(err)
    }
}
// EducationSubmissionResourceable 
type EducationSubmissionResourceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentResourceUrl()(*string)
    GetDependentResources()([]EducationSubmissionResourceable)
    GetResource()(EducationResourceable)
    SetAssignmentResourceUrl(value *string)()
    SetDependentResources(value []EducationSubmissionResourceable)()
    SetResource(value EducationResourceable)()
}
