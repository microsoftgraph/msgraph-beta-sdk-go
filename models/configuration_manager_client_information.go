package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConfigurationManagerClientInformation configuration Manager client information synced from SCCM
type ConfigurationManagerClientInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Configuration Manager Client Id from SCCM
    clientIdentifier *string;
    // Configuration Manager Client blocked status from SCCM
    isBlocked *bool;
}
// NewConfigurationManagerClientInformation instantiates a new configurationManagerClientInformation and sets the default values.
func NewConfigurationManagerClientInformation()(*ConfigurationManagerClientInformation) {
    m := &ConfigurationManagerClientInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConfigurationManagerClientInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConfigurationManagerClientInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConfigurationManagerClientInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConfigurationManagerClientInformation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClientIdentifier gets the clientIdentifier property value. Configuration Manager Client Id from SCCM
func (m *ConfigurationManagerClientInformation) GetClientIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.clientIdentifier
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConfigurationManagerClientInformation) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientIdentifier"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientIdentifier(val)
        }
        return nil
    }
    res["isBlocked"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBlocked(val)
        }
        return nil
    }
    return res
}
// GetIsBlocked gets the isBlocked property value. Configuration Manager Client blocked status from SCCM
func (m *ConfigurationManagerClientInformation) GetIsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBlocked
    }
}
// Serialize serializes information the current object
func (m *ConfigurationManagerClientInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientIdentifier", m.GetClientIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isBlocked", m.GetIsBlocked())
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
func (m *ConfigurationManagerClientInformation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClientIdentifier sets the clientIdentifier property value. Configuration Manager Client Id from SCCM
func (m *ConfigurationManagerClientInformation) SetClientIdentifier(value *string)() {
    if m != nil {
        m.clientIdentifier = value
    }
}
// SetIsBlocked sets the isBlocked property value. Configuration Manager Client blocked status from SCCM
func (m *ConfigurationManagerClientInformation) SetIsBlocked(value *bool)() {
    if m != nil {
        m.isBlocked = value
    }
}
