package approveapps

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApproveAppsRequestBody struct {
    additionalData map[string]interface{};
    approveAllPermissions *bool;
    packageIds []string;
}
func NewApproveAppsRequestBody()(*ApproveAppsRequestBody) {
    m := &ApproveAppsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ApproveAppsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ApproveAppsRequestBody) GetApproveAllPermissions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.approveAllPermissions
    }
}
func (m *ApproveAppsRequestBody) GetPackageIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.packageIds
    }
}
func (m *ApproveAppsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["approveAllPermissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetApproveAllPermissions(val)
        return nil
    }
    res["packageIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetPackageIds(res)
        return nil
    }
    return res
}
func (m *ApproveAppsRequestBody) IsNil()(bool) {
    return m == nil
}
func (m *ApproveAppsRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("approveAllPermissions", m.GetApproveAllPermissions())
        if err != nil {
            return err
        }
    }
    {
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
func (m *ApproveAppsRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ApproveAppsRequestBody) SetApproveAllPermissions(value *bool)() {
    m.approveAllPermissions = value
}
func (m *ApproveAppsRequestBody) SetPackageIds(value []string)() {
    m.packageIds = value
}
