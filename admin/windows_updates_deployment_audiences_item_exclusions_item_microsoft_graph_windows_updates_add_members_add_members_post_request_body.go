package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
)

// WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody 
type WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The assets property
    assets []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable
}
// NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody instantiates a new WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody()(*WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody) {
    m := &WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssets gets the assets property value. The assets property
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody) GetAssets()([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable) {
    return m.assets
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateUpdatableAssetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable, len(val))
            for i, v := range val {
                res[i] = v.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)
            }
            m.SetAssets(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssets()))
        for i, v := range m.GetAssets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("assets", cast)
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
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssets sets the assets property value. The assets property
func (m *WindowsUpdatesDeploymentAudiencesItemExclusionsItemMicrosoftGraphWindowsUpdatesAddMembersAddMembersPostRequestBody) SetAssets(value []i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdatableAssetable)() {
    m.assets = value
}
