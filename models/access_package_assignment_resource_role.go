package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentResourceRole 
type AccessPackageAssignmentResourceRole struct {
    Entity
}
// NewAccessPackageAssignmentResourceRole instantiates a new accessPackageAssignmentResourceRole and sets the default values.
func NewAccessPackageAssignmentResourceRole()(*AccessPackageAssignmentResourceRole) {
    m := &AccessPackageAssignmentResourceRole{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentResourceRoleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentResourceRole(), nil
}
// GetAccessPackageAssignments gets the accessPackageAssignments property value. The access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageAssignments()([]AccessPackageAssignmentable) {
    val, err := m.GetBackingStore().Get("accessPackageAssignments")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AccessPackageAssignmentable)
    }
    return nil
}
// GetAccessPackageResourceRole gets the accessPackageResourceRole property value. The accessPackageResourceRole property
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageResourceRole()(AccessPackageResourceRoleable) {
    val, err := m.GetBackingStore().Get("accessPackageResourceRole")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageResourceRoleable)
    }
    return nil
}
// GetAccessPackageResourceScope gets the accessPackageResourceScope property value. The accessPackageResourceScope property
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageResourceScope()(AccessPackageResourceScopeable) {
    val, err := m.GetBackingStore().Get("accessPackageResourceScope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageResourceScopeable)
    }
    return nil
}
// GetAccessPackageSubject gets the accessPackageSubject property value. Read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
func (m *AccessPackageAssignmentResourceRole) GetAccessPackageSubject()(AccessPackageSubjectable) {
    val, err := m.GetBackingStore().Get("accessPackageSubject")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AccessPackageSubjectable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentResourceRole) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AccessPackageAssignmentable)
                }
            }
            m.SetAccessPackageAssignments(res)
        }
        return nil
    }
    res["accessPackageResourceRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResourceRole(val.(AccessPackageResourceRoleable))
        }
        return nil
    }
    res["accessPackageResourceScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageResourceScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageResourceScope(val.(AccessPackageResourceScopeable))
        }
        return nil
    }
    res["accessPackageSubject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageSubjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackageSubject(val.(AccessPackageSubjectable))
        }
        return nil
    }
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
    res["originId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginId(val)
        }
        return nil
    }
    res["originSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginSystem(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessPackageAssignmentResourceRole) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOriginId gets the originId property value. A unique identifier relative to the origin system, corresponding to the originId property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) GetOriginId()(*string) {
    val, err := m.GetBackingStore().Get("originId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOriginSystem gets the originSystem property value. The system where the role assignment is to be created or has been created for an access package assignment, such as SharePointOnline, AadGroup or AadApplication, corresponding to the originSystem property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) GetOriginSystem()(*string) {
    val, err := m.GetBackingStore().Get("originSystem")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetStatus gets the status property value. The value is PendingFulfillment when the access package assignment has not yet been delivered to the origin system, and Fulfilled when the access package assignment has been delivered to the origin system.
func (m *AccessPackageAssignmentResourceRole) GetStatus()(*string) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentResourceRole) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessPackageAssignments()))
        for i, v := range m.GetAccessPackageAssignments() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageResourceRole", m.GetAccessPackageResourceRole())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageResourceScope", m.GetAccessPackageResourceScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessPackageSubject", m.GetAccessPackageSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originId", m.GetOriginId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originSystem", m.GetOriginSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackageAssignments sets the accessPackageAssignments property value. The access package assignments resulting in this role assignment. Read-only. Nullable.
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageAssignments(value []AccessPackageAssignmentable)() {
    err := m.GetBackingStore().Set("accessPackageAssignments", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageResourceRole sets the accessPackageResourceRole property value. The accessPackageResourceRole property
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageResourceRole(value AccessPackageResourceRoleable)() {
    err := m.GetBackingStore().Set("accessPackageResourceRole", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageResourceScope sets the accessPackageResourceScope property value. The accessPackageResourceScope property
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageResourceScope(value AccessPackageResourceScopeable)() {
    err := m.GetBackingStore().Set("accessPackageResourceScope", value)
    if err != nil {
        panic(err)
    }
}
// SetAccessPackageSubject sets the accessPackageSubject property value. Read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
func (m *AccessPackageAssignmentResourceRole) SetAccessPackageSubject(value AccessPackageSubjectable)() {
    err := m.GetBackingStore().Set("accessPackageSubject", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessPackageAssignmentResourceRole) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetOriginId sets the originId property value. A unique identifier relative to the origin system, corresponding to the originId property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) SetOriginId(value *string)() {
    err := m.GetBackingStore().Set("originId", value)
    if err != nil {
        panic(err)
    }
}
// SetOriginSystem sets the originSystem property value. The system where the role assignment is to be created or has been created for an access package assignment, such as SharePointOnline, AadGroup or AadApplication, corresponding to the originSystem property of the accessPackageResourceRole.
func (m *AccessPackageAssignmentResourceRole) SetOriginSystem(value *string)() {
    err := m.GetBackingStore().Set("originSystem", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. The value is PendingFulfillment when the access package assignment has not yet been delivered to the origin system, and Fulfilled when the access package assignment has been delivered to the origin system.
func (m *AccessPackageAssignmentResourceRole) SetStatus(value *string)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
// AccessPackageAssignmentResourceRoleable 
type AccessPackageAssignmentResourceRoleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccessPackageAssignments()([]AccessPackageAssignmentable)
    GetAccessPackageResourceRole()(AccessPackageResourceRoleable)
    GetAccessPackageResourceScope()(AccessPackageResourceScopeable)
    GetAccessPackageSubject()(AccessPackageSubjectable)
    GetOdataType()(*string)
    GetOriginId()(*string)
    GetOriginSystem()(*string)
    GetStatus()(*string)
    SetAccessPackageAssignments(value []AccessPackageAssignmentable)()
    SetAccessPackageResourceRole(value AccessPackageResourceRoleable)()
    SetAccessPackageResourceScope(value AccessPackageResourceScopeable)()
    SetAccessPackageSubject(value AccessPackageSubjectable)()
    SetOdataType(value *string)()
    SetOriginId(value *string)()
    SetOriginSystem(value *string)()
    SetStatus(value *string)()
}
