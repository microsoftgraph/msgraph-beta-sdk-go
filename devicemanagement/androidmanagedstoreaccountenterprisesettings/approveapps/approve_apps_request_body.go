package approveapps

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApproveAppsRequestBody 
type ApproveAppsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    approveAllPermissions *bool;
    // 
    packageIds []string;
}
// NewApproveAppsRequestBody instantiates a new approveAppsRequestBody and sets the default values.
func NewApproveAppsRequestBody()(*ApproveAppsRequestBody) {
    m := &ApproveAppsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ApproveAppsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetApproveAllPermissions gets the approveAllPermissions property value. 
func (m *ApproveAppsRequestBody) GetApproveAllPermissions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.approveAllPermissions
    }
}
// GetPackageIds gets the packageIds property value. 
func (m *ApproveAppsRequestBody) GetPackageIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.packageIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApproveAppsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["approveAllPermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApproveAllPermissions(val)
        }
        return nil
    }
    res["packageIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *ApproveAppsRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApproveAppsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// SetApproveAllPermissions sets the approveAllPermissions property value. 
func (m *ApproveAppsRequestBody) SetApproveAllPermissions(value *bool)() {
    if m != nil {
        m.approveAllPermissions = value
    }
}
// SetPackageIds sets the packageIds property value. 
func (m *ApproveAppsRequestBody) SetPackageIds(value []string)() {
    if m != nil {
        m.packageIds = value
    }
}
