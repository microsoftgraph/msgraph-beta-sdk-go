package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsCloudManagementDevicesSummary the user experience work from anywhere Cloud management devices summary.
type UserExperienceAnalyticsCloudManagementDevicesSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Total number of  co-managed devices.
    coManagedDeviceCount *int32
    // The count of intune devices that are not autopilot registerd.
    intuneDeviceCount *int32
    // The OdataType property
    odataType *string
    // Total count of tenant attach devices.
    tenantAttachDeviceCount *int32
}
// NewUserExperienceAnalyticsCloudManagementDevicesSummary instantiates a new userExperienceAnalyticsCloudManagementDevicesSummary and sets the default values.
func NewUserExperienceAnalyticsCloudManagementDevicesSummary()(*UserExperienceAnalyticsCloudManagementDevicesSummary) {
    m := &UserExperienceAnalyticsCloudManagementDevicesSummary{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.userExperienceAnalyticsCloudManagementDevicesSummary";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateUserExperienceAnalyticsCloudManagementDevicesSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserExperienceAnalyticsCloudManagementDevicesSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserExperienceAnalyticsCloudManagementDevicesSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCoManagedDeviceCount gets the coManagedDeviceCount property value. Total number of  co-managed devices.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetCoManagedDeviceCount()(*int32) {
    return m.coManagedDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["coManagedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoManagedDeviceCount(val)
        }
        return nil
    }
    res["intuneDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntuneDeviceCount(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["tenantAttachDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantAttachDeviceCount(val)
        }
        return nil
    }
    return res
}
// GetIntuneDeviceCount gets the intuneDeviceCount property value. The count of intune devices that are not autopilot registerd.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetIntuneDeviceCount()(*int32) {
    return m.intuneDeviceCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetTenantAttachDeviceCount gets the tenantAttachDeviceCount property value. Total count of tenant attach devices.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) GetTenantAttachDeviceCount()(*int32) {
    return m.tenantAttachDeviceCount
}
// Serialize serializes information the current object
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("coManagedDeviceCount", m.GetCoManagedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("intuneDeviceCount", m.GetIntuneDeviceCount())
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
        err := writer.WriteInt32Value("tenantAttachDeviceCount", m.GetTenantAttachDeviceCount())
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
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCoManagedDeviceCount sets the coManagedDeviceCount property value. Total number of  co-managed devices.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetCoManagedDeviceCount(value *int32)() {
    m.coManagedDeviceCount = value
}
// SetIntuneDeviceCount sets the intuneDeviceCount property value. The count of intune devices that are not autopilot registerd.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetIntuneDeviceCount(value *int32)() {
    m.intuneDeviceCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTenantAttachDeviceCount sets the tenantAttachDeviceCount property value. Total count of tenant attach devices.
func (m *UserExperienceAnalyticsCloudManagementDevicesSummary) SetTenantAttachDeviceCount(value *int32)() {
    m.tenantAttachDeviceCount = value
}
