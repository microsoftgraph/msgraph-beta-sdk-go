package managedtenants

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagementTemplateDetailedInfo 
type ManagementTemplateDetailedInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
    category *ManagementCategory
    // The display name for the management template. Required. Read-only.
    displayName *string
    // The unique identifier for the management template. Required. Read-only.
    managementTemplateId *string
    // The version property
    version *int32
}
// NewManagementTemplateDetailedInfo instantiates a new managementTemplateDetailedInfo and sets the default values.
func NewManagementTemplateDetailedInfo()(*ManagementTemplateDetailedInfo) {
    m := &ManagementTemplateDetailedInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateManagementTemplateDetailedInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagementTemplateDetailedInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagementTemplateDetailedInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagementTemplateDetailedInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCategory gets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) GetCategory()(*ManagementCategory) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// GetDisplayName gets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagementTemplateDetailedInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["category"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagementCategory)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory(val.(*ManagementCategory))
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
    res["managementTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagementTemplateId(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetManagementTemplateId gets the managementTemplateId property value. The unique identifier for the management template. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) GetManagementTemplateId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.managementTemplateId
    }
}
// GetVersion gets the version property value. The version property
func (m *ManagementTemplateDetailedInfo) GetVersion()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// Serialize serializes information the current object
func (m *ManagementTemplateDetailedInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCategory() != nil {
        cast := (*m.GetCategory()).String()
        err := writer.WriteStringValue("category", &cast)
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
        err := writer.WriteStringValue("managementTemplateId", m.GetManagementTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("version", m.GetVersion())
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
func (m *ManagementTemplateDetailedInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCategory sets the category property value. The management category for the management template. Possible values are: custom, devices, identity, unknownFutureValue. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) SetCategory(value *ManagementCategory)() {
    if m != nil {
        m.category = value
    }
}
// SetDisplayName sets the displayName property value. The display name for the management template. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetManagementTemplateId sets the managementTemplateId property value. The unique identifier for the management template. Required. Read-only.
func (m *ManagementTemplateDetailedInfo) SetManagementTemplateId(value *string)() {
    if m != nil {
        m.managementTemplateId = value
    }
}
// SetVersion sets the version property value. The version property
func (m *ManagementTemplateDetailedInfo) SetVersion(value *int32)() {
    if m != nil {
        m.version = value
    }
}
