package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// AndroidFotaDeploymentAssignment describes deployment security group to assign a deployment to. The backend will expand the security Group ID to extract device serial numbers prior sending a create deployment request to Zebra.
type AndroidFotaDeploymentAssignment struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewAndroidFotaDeploymentAssignment instantiates a new androidFotaDeploymentAssignment and sets the default values.
func NewAndroidFotaDeploymentAssignment()(*AndroidFotaDeploymentAssignment) {
    m := &AndroidFotaDeploymentAssignment{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidFotaDeploymentAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidFotaDeploymentAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidFotaDeploymentAssignment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidFotaDeploymentAssignment) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAssignmentTarget gets the assignmentTarget property value. The Azure Active Directory (Azure AD) we are deploying firmware updates to (e.g.: d93c8f48-bd42-4514-ba40-bc6b84780930). NOTE: Use this property moving forward because the existing property, target, is deprecated.
func (m *AndroidFotaDeploymentAssignment) GetAssignmentTarget()(DeviceAndAppManagementAssignmentTargetable) {
    val, err := m.GetBackingStore().Get("assignmentTarget")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DeviceAndAppManagementAssignmentTargetable)
    }
    return nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
func (m *AndroidFotaDeploymentAssignment) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDisplayName gets the displayName property value. The display name of the Azure AD security group used for the assignment.
func (m *AndroidFotaDeploymentAssignment) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidFotaDeploymentAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidFotaDeploymentAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(AndroidFotaDeploymentAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetId gets the id property value. A unique identifier assigned to each Android FOTA Assignment entity
func (m *AndroidFotaDeploymentAssignment) GetId()(*string) {
    val, err := m.GetBackingStore().Get("id")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidFotaDeploymentAssignment) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTarget gets the target property value. The AAD Group we are deploying firmware updates to
func (m *AndroidFotaDeploymentAssignment) GetTarget()(AndroidFotaDeploymentAssignmentTargetable) {
    val, err := m.GetBackingStore().Get("target")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(AndroidFotaDeploymentAssignmentTargetable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *AndroidFotaDeploymentAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("assignmentTarget", m.GetAssignmentTarget())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidFotaDeploymentAssignment) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignmentTarget sets the assignmentTarget property value. The Azure Active Directory (Azure AD) we are deploying firmware updates to (e.g.: d93c8f48-bd42-4514-ba40-bc6b84780930). NOTE: Use this property moving forward because the existing property, target, is deprecated.
func (m *AndroidFotaDeploymentAssignment) SetAssignmentTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    err := m.GetBackingStore().Set("assignmentTarget", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *AndroidFotaDeploymentAssignment) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDisplayName sets the displayName property value. The display name of the Azure AD security group used for the assignment.
func (m *AndroidFotaDeploymentAssignment) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetId sets the id property value. A unique identifier assigned to each Android FOTA Assignment entity
func (m *AndroidFotaDeploymentAssignment) SetId(value *string)() {
    err := m.GetBackingStore().Set("id", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidFotaDeploymentAssignment) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetTarget sets the target property value. The AAD Group we are deploying firmware updates to
func (m *AndroidFotaDeploymentAssignment) SetTarget(value AndroidFotaDeploymentAssignmentTargetable)() {
    err := m.GetBackingStore().Set("target", value)
    if err != nil {
        panic(err)
    }
}
// AndroidFotaDeploymentAssignmentable 
type AndroidFotaDeploymentAssignmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentTarget()(DeviceAndAppManagementAssignmentTargetable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDisplayName()(*string)
    GetId()(*string)
    GetOdataType()(*string)
    GetTarget()(AndroidFotaDeploymentAssignmentTargetable)
    SetAssignmentTarget(value DeviceAndAppManagementAssignmentTargetable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDisplayName(value *string)()
    SetId(value *string)()
    SetOdataType(value *string)()
    SetTarget(value AndroidFotaDeploymentAssignmentTargetable)()
}
