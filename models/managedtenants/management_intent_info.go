package managedtenants

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementIntentInfo 
type ManagementIntentInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The display name for the management intent. Optional. Read-only.
    managementIntentDisplayName *string
    // The identifier for the management intent. Required. Read-only.
    managementIntentId *string
    // The collection of management template information associated with the management intent. Optional. Read-only.
    managementTemplates []ManagementTemplateDetailedInfoable
    // The OdataType property
    odataType *string
}
// NewManagementIntentInfo instantiates a new managementIntentInfo and sets the default values.
func NewManagementIntentInfo()(*ManagementIntentInfo) {
    m := &ManagementIntentInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagementIntentInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementIntentInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementIntentInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementIntentInfo) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementIntentInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["managementIntentDisplayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagementIntentDisplayName)
    res["managementIntentId"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetManagementIntentId)
    res["managementTemplates"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateManagementTemplateDetailedInfoFromDiscriminatorValue , m.SetManagementTemplates)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    return res
}
// GetManagementIntentDisplayName gets the managementIntentDisplayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntentInfo) GetManagementIntentDisplayName()(*string) {
    return m.managementIntentDisplayName
}
// GetManagementIntentId gets the managementIntentId property value. The identifier for the management intent. Required. Read-only.
func (m *ManagementIntentInfo) GetManagementIntentId()(*string) {
    return m.managementIntentId
}
// GetManagementTemplates gets the managementTemplates property value. The collection of management template information associated with the management intent. Optional. Read-only.
func (m *ManagementIntentInfo) GetManagementTemplates()([]ManagementTemplateDetailedInfoable) {
    return m.managementTemplates
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ManagementIntentInfo) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ManagementIntentInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("managementIntentDisplayName", m.GetManagementIntentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("managementIntentId", m.GetManagementIntentId())
        if err != nil {
            return err
        }
    }
    if m.GetManagementTemplates() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetManagementTemplates())
        err := writer.WriteCollectionOfObjectValues("managementTemplates", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementIntentInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetManagementIntentDisplayName sets the managementIntentDisplayName property value. The display name for the management intent. Optional. Read-only.
func (m *ManagementIntentInfo) SetManagementIntentDisplayName(value *string)() {
    m.managementIntentDisplayName = value
}
// SetManagementIntentId sets the managementIntentId property value. The identifier for the management intent. Required. Read-only.
func (m *ManagementIntentInfo) SetManagementIntentId(value *string)() {
    m.managementIntentId = value
}
// SetManagementTemplates sets the managementTemplates property value. The collection of management template information associated with the management intent. Optional. Read-only.
func (m *ManagementIntentInfo) SetManagementTemplates(value []ManagementTemplateDetailedInfoable)() {
    m.managementTemplates = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ManagementIntentInfo) SetOdataType(value *string)() {
    m.odataType = value
}
