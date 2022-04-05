package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OfficeConfiguration 
type OfficeConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // List of office Client configuration.
    clientConfigurations []OfficeClientConfigurationable;
    // List of office Client check-in status.
    tenantCheckinStatuses []OfficeClientCheckinStatusable;
    // Entity that describes tenant check-in statues
    tenantUserCheckinSummary OfficeUserCheckinSummaryable;
}
// NewOfficeConfiguration instantiates a new OfficeConfiguration and sets the default values.
func NewOfficeConfiguration()(*OfficeConfiguration) {
    m := &OfficeConfiguration{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOfficeConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOfficeConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOfficeConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OfficeConfiguration) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClientConfigurations gets the clientConfigurations property value. List of office Client configuration.
func (m *OfficeConfiguration) GetClientConfigurations()([]OfficeClientConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.clientConfigurations
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OfficeConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientConfigurations"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOfficeClientConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(OfficeClientConfigurationable)
            }
            m.SetClientConfigurations(res)
        }
        return nil
    }
    res["tenantCheckinStatuses"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOfficeClientCheckinStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfficeClientCheckinStatusable, len(val))
            for i, v := range val {
                res[i] = v.(OfficeClientCheckinStatusable)
            }
            m.SetTenantCheckinStatuses(res)
        }
        return nil
    }
    res["tenantUserCheckinSummary"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOfficeUserCheckinSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantUserCheckinSummary(val.(OfficeUserCheckinSummaryable))
        }
        return nil
    }
    return res
}
// GetTenantCheckinStatuses gets the tenantCheckinStatuses property value. List of office Client check-in status.
func (m *OfficeConfiguration) GetTenantCheckinStatuses()([]OfficeClientCheckinStatusable) {
    if m == nil {
        return nil
    } else {
        return m.tenantCheckinStatuses
    }
}
// GetTenantUserCheckinSummary gets the tenantUserCheckinSummary property value. Entity that describes tenant check-in statues
func (m *OfficeConfiguration) GetTenantUserCheckinSummary()(OfficeUserCheckinSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.tenantUserCheckinSummary
    }
}
// Serialize serializes information the current object
func (m *OfficeConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClientConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClientConfigurations()))
        for i, v := range m.GetClientConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("clientConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTenantCheckinStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTenantCheckinStatuses()))
        for i, v := range m.GetTenantCheckinStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("tenantCheckinStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tenantUserCheckinSummary", m.GetTenantUserCheckinSummary())
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
func (m *OfficeConfiguration) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClientConfigurations sets the clientConfigurations property value. List of office Client configuration.
func (m *OfficeConfiguration) SetClientConfigurations(value []OfficeClientConfigurationable)() {
    if m != nil {
        m.clientConfigurations = value
    }
}
// SetTenantCheckinStatuses sets the tenantCheckinStatuses property value. List of office Client check-in status.
func (m *OfficeConfiguration) SetTenantCheckinStatuses(value []OfficeClientCheckinStatusable)() {
    if m != nil {
        m.tenantCheckinStatuses = value
    }
}
// SetTenantUserCheckinSummary sets the tenantUserCheckinSummary property value. Entity that describes tenant check-in statues
func (m *OfficeConfiguration) SetTenantUserCheckinSummary(value OfficeUserCheckinSummaryable)() {
    if m != nil {
        m.tenantUserCheckinSummary = value
    }
}
