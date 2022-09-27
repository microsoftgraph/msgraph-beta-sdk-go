package assignedaccessmultimodeprofiles

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AssignedAccessMultiModeProfilesPostRequestBody provides operations to call the assignedAccessMultiModeProfiles method.
type AssignedAccessMultiModeProfilesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The assignedAccessMultiModeProfiles property
    assignedAccessMultiModeProfiles []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable
}
// NewAssignedAccessMultiModeProfilesPostRequestBody instantiates a new assignedAccessMultiModeProfilesPostRequestBody and sets the default values.
func NewAssignedAccessMultiModeProfilesPostRequestBody()(*AssignedAccessMultiModeProfilesPostRequestBody) {
    m := &AssignedAccessMultiModeProfilesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignedAccessMultiModeProfilesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignedAccessMultiModeProfilesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignedAccessMultiModeProfilesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignedAccessMultiModeProfilesPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignedAccessMultiModeProfiles gets the assignedAccessMultiModeProfiles property value. The assignedAccessMultiModeProfiles property
func (m *AssignedAccessMultiModeProfilesPostRequestBody) GetAssignedAccessMultiModeProfiles()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable) {
    return m.assignedAccessMultiModeProfiles
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignedAccessMultiModeProfilesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedAccessMultiModeProfiles"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWindowsAssignedAccessProfileFromDiscriminatorValue , m.SetAssignedAccessMultiModeProfiles)
    return res
}
// Serialize serializes information the current object
func (m *AssignedAccessMultiModeProfilesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignedAccessMultiModeProfiles() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetAssignedAccessMultiModeProfiles())
        err := writer.WriteCollectionOfObjectValues("assignedAccessMultiModeProfiles", cast)
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
func (m *AssignedAccessMultiModeProfilesPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignedAccessMultiModeProfiles sets the assignedAccessMultiModeProfiles property value. The assignedAccessMultiModeProfiles property
func (m *AssignedAccessMultiModeProfilesPostRequestBody) SetAssignedAccessMultiModeProfiles(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WindowsAssignedAccessProfileable)() {
    m.assignedAccessMultiModeProfiles = value
}
