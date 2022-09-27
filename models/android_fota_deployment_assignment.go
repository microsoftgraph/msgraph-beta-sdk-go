package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidFotaDeploymentAssignment describes deployment security group to assign a deployment to. The backend will expand the security Group ID to extract device serial numbers prior sending a create deployment request to Zebra.
type AndroidFotaDeploymentAssignment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The display name of the Azure AD security group used for the assignment.
    displayName *string
    // A unique identifier assigned to each Android FOTA Assignment entity
    id *string
    // The OdataType property
    odataType *string
    // The AAD Group we are deploying firmware updates to
    target AndroidFotaDeploymentAssignmentTargetable
}
// NewAndroidFotaDeploymentAssignment instantiates a new androidFotaDeploymentAssignment and sets the default values.
func NewAndroidFotaDeploymentAssignment()(*AndroidFotaDeploymentAssignment) {
    m := &AndroidFotaDeploymentAssignment{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.androidFotaDeploymentAssignment";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidFotaDeploymentAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidFotaDeploymentAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidFotaDeploymentAssignment(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidFotaDeploymentAssignment) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The display name of the Azure AD security group used for the assignment.
func (m *AndroidFotaDeploymentAssignment) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidFotaDeploymentAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["id"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetId)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["target"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateAndroidFotaDeploymentAssignmentTargetFromDiscriminatorValue , m.SetTarget)
    return res
}
// GetId gets the id property value. A unique identifier assigned to each Android FOTA Assignment entity
func (m *AndroidFotaDeploymentAssignment) GetId()(*string) {
    return m.id
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidFotaDeploymentAssignment) GetOdataType()(*string) {
    return m.odataType
}
// GetTarget gets the target property value. The AAD Group we are deploying firmware updates to
func (m *AndroidFotaDeploymentAssignment) GetTarget()(AndroidFotaDeploymentAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *AndroidFotaDeploymentAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidFotaDeploymentAssignment) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The display name of the Azure AD security group used for the assignment.
func (m *AndroidFotaDeploymentAssignment) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetId sets the id property value. A unique identifier assigned to each Android FOTA Assignment entity
func (m *AndroidFotaDeploymentAssignment) SetId(value *string)() {
    m.id = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidFotaDeploymentAssignment) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTarget sets the target property value. The AAD Group we are deploying firmware updates to
func (m *AndroidFotaDeploymentAssignment) SetTarget(value AndroidFotaDeploymentAssignmentTargetable)() {
    m.target = value
}
