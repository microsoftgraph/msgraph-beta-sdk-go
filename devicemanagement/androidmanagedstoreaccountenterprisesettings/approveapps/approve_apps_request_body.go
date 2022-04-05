package approveapps

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApproveAppsRequestBody provides operations to call the approveApps method.
type ApproveAppsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The approveAllPermissions property
    approveAllPermissions *bool;
    // The packageIds property
    packageIds []string;
}
// NewApproveAppsRequestBody instantiates a new approveAppsRequestBody and sets the default values.
func NewApproveAppsRequestBody()(*ApproveAppsRequestBody) {
    m := &ApproveAppsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateApproveAppsRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApproveAppsRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApproveAppsRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApproveAppsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApproveAllPermissions gets the approveAllPermissions property value. The approveAllPermissions property
func (m *ApproveAppsRequestBody) GetApproveAllPermissions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.approveAllPermissions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApproveAppsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["approveAllPermissions"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproveAllPermissions(val)
        }
        return nil
    }
    res["packageIds"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetPackageIds(res)
        }
        return nil
    }
    return res
}
// GetPackageIds gets the packageIds property value. The packageIds property
func (m *ApproveAppsRequestBody) GetPackageIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.packageIds
    }
}
// Serialize serializes information the current object
func (m *ApproveAppsRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("approveAllPermissions", m.GetApproveAllPermissions())
        if err != nil {
            return err
        }
    }
    if m.GetPackageIds() != nil {
        err := writer.WriteCollectionOfStringValues("packageIds", m.GetPackageIds())
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
func (m *ApproveAppsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetApproveAllPermissions sets the approveAllPermissions property value. The approveAllPermissions property
func (m *ApproveAppsRequestBody) SetApproveAllPermissions(value *bool)() {
    if m != nil {
        m.approveAllPermissions = value
    }
}
// SetPackageIds sets the packageIds property value. The packageIds property
func (m *ApproveAppsRequestBody) SetPackageIds(value []string)() {
    if m != nil {
        m.packageIds = value
    }
}
