package enrollassetsbyid

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
)

// EnrollAssetsByIdPostRequestBody provides operations to call the enrollAssetsById method.
type EnrollAssetsByIdPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ids property
    ids []string
    // The memberEntityType property
    memberEntityType *string
    // The updateCategory property
    updateCategory *i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory
}
// NewEnrollAssetsByIdPostRequestBody instantiates a new enrollAssetsByIdPostRequestBody and sets the default values.
func NewEnrollAssetsByIdPostRequestBody()(*EnrollAssetsByIdPostRequestBody) {
    m := &EnrollAssetsByIdPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateEnrollAssetsByIdPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrollAssetsByIdPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrollAssetsByIdPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EnrollAssetsByIdPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollAssetsByIdPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfPrimitiveValues("string" , m.SetIds)
    res["memberEntityType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetMemberEntityType)
    res["updateCategory"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ParseUpdateCategory , m.SetUpdateCategory)
    return res
}
// GetIds gets the ids property value. The ids property
func (m *EnrollAssetsByIdPostRequestBody) GetIds()([]string) {
    return m.ids
}
// GetMemberEntityType gets the memberEntityType property value. The memberEntityType property
func (m *EnrollAssetsByIdPostRequestBody) GetMemberEntityType()(*string) {
    return m.memberEntityType
}
// GetUpdateCategory gets the updateCategory property value. The updateCategory property
func (m *EnrollAssetsByIdPostRequestBody) GetUpdateCategory()(*i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory) {
    return m.updateCategory
}
// Serialize serializes information the current object
func (m *EnrollAssetsByIdPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *EnrollAssetsByIdPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetIds sets the ids property value. The ids property
func (m *EnrollAssetsByIdPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
// SetMemberEntityType sets the memberEntityType property value. The memberEntityType property
func (m *EnrollAssetsByIdPostRequestBody) SetMemberEntityType(value *string)() {
    m.memberEntityType = value
}
// SetUpdateCategory sets the updateCategory property value. The updateCategory property
func (m *EnrollAssetsByIdPostRequestBody) SetUpdateCategory(value *i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory)() {
    m.updateCategory = value
}
