package unenrollassetsbyid

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
)

// UnenrollAssetsByIdRequestBody provides operations to call the unenrollAssetsById method.
type UnenrollAssetsByIdRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The ids property
    ids []string
    // The memberEntityType property
    memberEntityType *string
    // The updateCategory property
    updateCategory *i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory
}
// NewUnenrollAssetsByIdRequestBody instantiates a new unenrollAssetsByIdRequestBody and sets the default values.
func NewUnenrollAssetsByIdRequestBody()(*UnenrollAssetsByIdRequestBody) {
    m := &UnenrollAssetsByIdRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateUnenrollAssetsByIdRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnenrollAssetsByIdRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnenrollAssetsByIdRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnenrollAssetsByIdRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnenrollAssetsByIdRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *UnenrollAssetsByIdRequestBody) GetIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.ids
    }
}
// GetMemberEntityType gets the memberEntityType property value. The memberEntityType property
func (m *UnenrollAssetsByIdRequestBody) GetMemberEntityType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.memberEntityType
    }
}
// GetUpdateCategory gets the updateCategory property value. The updateCategory property
func (m *UnenrollAssetsByIdRequestBody) GetUpdateCategory()(*i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory) {
    if m == nil {
        return nil
    } else {
        return m.updateCategory
    }
}
// Serialize serializes information the current object
func (m *UnenrollAssetsByIdRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *UnenrollAssetsByIdRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIds sets the ids property value. The ids property
func (m *UnenrollAssetsByIdRequestBody) SetIds(value []string)() {
    if m != nil {
        m.ids = value
    }
}
// SetMemberEntityType sets the memberEntityType property value. The memberEntityType property
func (m *UnenrollAssetsByIdRequestBody) SetMemberEntityType(value *string)() {
    if m != nil {
        m.memberEntityType = value
    }
}
// SetUpdateCategory sets the updateCategory property value. The updateCategory property
func (m *UnenrollAssetsByIdRequestBody) SetUpdateCategory(value *i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.UpdateCategory)() {
    if m != nil {
        m.updateCategory = value
    }
}
