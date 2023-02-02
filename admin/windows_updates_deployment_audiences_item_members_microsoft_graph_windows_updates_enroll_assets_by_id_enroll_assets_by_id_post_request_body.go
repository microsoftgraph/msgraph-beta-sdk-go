package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
)

// WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody 
type WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ids property
    ids []string
    // The memberEntityType property
    memberEntityType *string
    // The updateCategory property
    updateCategory *i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory
}
// NewWindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody instantiates a new WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody and sets the default values.
func NewWindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody()(*WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) {
    m := &WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIds(res)
        }
        return nil
    }
    res["memberEntityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMemberEntityType(val)
        }
        return nil
    }
    res["updateCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ParseUpdateCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateCategory(val.(*i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory))
        }
        return nil
    }
    return res
}
// GetIds gets the ids property value. The ids property
func (m *WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) GetIds()([]string) {
    return m.ids
}
// GetMemberEntityType gets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) GetMemberEntityType()(*string) {
    return m.memberEntityType
}
// GetUpdateCategory gets the updateCategory property value. The updateCategory property
func (m *WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) GetUpdateCategory()(*i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory) {
    return m.updateCategory
}
// Serialize serializes information the current object
func (m *WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("memberEntityType", m.GetMemberEntityType())
        if err != nil {
            return err
        }
    }
    if m.GetUpdateCategory() != nil {
        cast := (*m.GetUpdateCategory()).String()
        err := writer.WriteStringValue("updateCategory", &cast)
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
func (m *WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIds sets the ids property value. The ids property
func (m *WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
// SetMemberEntityType sets the memberEntityType property value. The memberEntityType property
func (m *WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) SetMemberEntityType(value *string)() {
    m.memberEntityType = value
}
// SetUpdateCategory sets the updateCategory property value. The updateCategory property
func (m *WindowsUpdatesDeploymentAudiencesItemMembersMicrosoftGraphWindowsUpdatesEnrollAssetsByIdEnrollAssetsByIdPostRequestBody) SetUpdateCategory(value *i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory)() {
    m.updateCategory = value
}
