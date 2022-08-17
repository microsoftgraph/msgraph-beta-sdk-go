package assign

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// AssignPostRequestBody provides operations to call the assign method.
type AssignPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The appProvisioningConfigurationGroupAssignments property
    appProvisioningConfigurationGroupAssignments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable
    // The iOSLobAppProvisioningConfigAssignments property
    iOSLobAppProvisioningConfigAssignments []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable
}
// NewAssignPostRequestBody instantiates a new assignPostRequestBody and sets the default values.
func NewAssignPostRequestBody()(*AssignPostRequestBody) {
    m := &AssignPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAppProvisioningConfigurationGroupAssignments gets the appProvisioningConfigurationGroupAssignments property value. The appProvisioningConfigurationGroupAssignments property
func (m *AssignPostRequestBody) GetAppProvisioningConfigurationGroupAssignments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable) {
    return m.appProvisioningConfigurationGroupAssignments
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appProvisioningConfigurationGroupAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMobileAppProvisioningConfigGroupAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable)
            }
            m.SetAppProvisioningConfigurationGroupAssignments(res)
        }
        return nil
    }
    res["iOSLobAppProvisioningConfigAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable)
            }
            m.SetIOSLobAppProvisioningConfigAssignments(res)
        }
        return nil
    }
    return res
}
// GetIOSLobAppProvisioningConfigAssignments gets the iOSLobAppProvisioningConfigAssignments property value. The iOSLobAppProvisioningConfigAssignments property
func (m *AssignPostRequestBody) GetIOSLobAppProvisioningConfigAssignments()([]ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable) {
    return m.iOSLobAppProvisioningConfigAssignments
}
// Serialize serializes information the current object
func (m *AssignPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAppProvisioningConfigurationGroupAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppProvisioningConfigurationGroupAssignments()))
        for i, v := range m.GetAppProvisioningConfigurationGroupAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("appProvisioningConfigurationGroupAssignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIOSLobAppProvisioningConfigAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIOSLobAppProvisioningConfigAssignments()))
        for i, v := range m.GetIOSLobAppProvisioningConfigAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("iOSLobAppProvisioningConfigAssignments", cast)
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
func (m *AssignPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAppProvisioningConfigurationGroupAssignments sets the appProvisioningConfigurationGroupAssignments property value. The appProvisioningConfigurationGroupAssignments property
func (m *AssignPostRequestBody) SetAppProvisioningConfigurationGroupAssignments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MobileAppProvisioningConfigGroupAssignmentable)() {
    m.appProvisioningConfigurationGroupAssignments = value
}
// SetIOSLobAppProvisioningConfigAssignments sets the iOSLobAppProvisioningConfigAssignments property value. The iOSLobAppProvisioningConfigAssignments property
func (m *AssignPostRequestBody) SetIOSLobAppProvisioningConfigAssignments(value []ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable)() {
    m.iOSLobAppProvisioningConfigAssignments = value
}
