package createmigrationreport

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
)

// CreateMigrationReportRequestBody provides operations to call the createMigrationReport method.
type CreateMigrationReportRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The groupPolicyObjectFile property
    groupPolicyObjectFile ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyObjectFileable
}
// NewCreateMigrationReportRequestBody instantiates a new createMigrationReportRequestBody and sets the default values.
func NewCreateMigrationReportRequestBody()(*CreateMigrationReportRequestBody) {
    m := &CreateMigrationReportRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCreateMigrationReportRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreateMigrationReportRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateMigrationReportRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateMigrationReportRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateMigrationReportRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupPolicyObjectFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyObjectFileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupPolicyObjectFile(val.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyObjectFileable))
        }
        return nil
    }
    return res
}
// GetGroupPolicyObjectFile gets the groupPolicyObjectFile property value. The groupPolicyObjectFile property
func (m *CreateMigrationReportRequestBody) GetGroupPolicyObjectFile()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyObjectFileable) {
    if m == nil {
        return nil
    } else {
        return m.groupPolicyObjectFile
    }
}
// Serialize serializes information the current object
func (m *CreateMigrationReportRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("groupPolicyObjectFile", m.GetGroupPolicyObjectFile())
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
func (m *CreateMigrationReportRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGroupPolicyObjectFile sets the groupPolicyObjectFile property value. The groupPolicyObjectFile property
func (m *CreateMigrationReportRequestBody) SetGroupPolicyObjectFile(value ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyObjectFileable)() {
    if m != nil {
        m.groupPolicyObjectFile = value
    }
}
