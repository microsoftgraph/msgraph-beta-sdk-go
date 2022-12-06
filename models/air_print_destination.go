package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AirPrintDestination represents an AirPrint destination.
type AirPrintDestination struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // If true AirPrint connections are secured by Transport Layer Security (TLS). Default is false. Available in iOS 11.0 and later.
    forceTls *bool
    // The IP Address of the AirPrint destination.
    ipAddress *string
    // The OdataType property
    odataType *string
    // The listening port of the AirPrint destination. If this key is not specified AirPrint will use the default port. Available in iOS 11.0 and later.
    port *int32
    // The Resource Path associated with the printer. This corresponds to the rp parameter of the ipps.tcp Bonjour record. For example: printers/Canon_MG5300_series, printers/Xerox_Phaser_7600, ipp/print, Epson_IPP_Printer.
    resourcePath *string
}
// NewAirPrintDestination instantiates a new airPrintDestination and sets the default values.
func NewAirPrintDestination()(*AirPrintDestination) {
    m := &AirPrintDestination{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAirPrintDestinationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAirPrintDestinationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAirPrintDestination(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AirPrintDestination) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AirPrintDestination) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["forceTls"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetForceTls)
    res["ipAddress"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetIpAddress)
    res["@odata.type"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOdataType)
    res["port"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPort)
    res["resourcePath"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetResourcePath)
    return res
}
// GetForceTls gets the forceTls property value. If true AirPrint connections are secured by Transport Layer Security (TLS). Default is false. Available in iOS 11.0 and later.
func (m *AirPrintDestination) GetForceTls()(*bool) {
    return m.forceTls
}
// GetIpAddress gets the ipAddress property value. The IP Address of the AirPrint destination.
func (m *AirPrintDestination) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AirPrintDestination) GetOdataType()(*string) {
    return m.odataType
}
// GetPort gets the port property value. The listening port of the AirPrint destination. If this key is not specified AirPrint will use the default port. Available in iOS 11.0 and later.
func (m *AirPrintDestination) GetPort()(*int32) {
    return m.port
}
// GetResourcePath gets the resourcePath property value. The Resource Path associated with the printer. This corresponds to the rp parameter of the ipps.tcp Bonjour record. For example: printers/Canon_MG5300_series, printers/Xerox_Phaser_7600, ipp/print, Epson_IPP_Printer.
func (m *AirPrintDestination) GetResourcePath()(*string) {
    return m.resourcePath
}
// Serialize serializes information the current object
func (m *AirPrintDestination) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("forceTls", m.GetForceTls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
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
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourcePath", m.GetResourcePath())
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
func (m *AirPrintDestination) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetForceTls sets the forceTls property value. If true AirPrint connections are secured by Transport Layer Security (TLS). Default is false. Available in iOS 11.0 and later.
func (m *AirPrintDestination) SetForceTls(value *bool)() {
    m.forceTls = value
}
// SetIpAddress sets the ipAddress property value. The IP Address of the AirPrint destination.
func (m *AirPrintDestination) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AirPrintDestination) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPort sets the port property value. The listening port of the AirPrint destination. If this key is not specified AirPrint will use the default port. Available in iOS 11.0 and later.
func (m *AirPrintDestination) SetPort(value *int32)() {
    m.port = value
}
// SetResourcePath sets the resourcePath property value. The Resource Path associated with the printer. This corresponds to the rp parameter of the ipps.tcp Bonjour record. For example: printers/Canon_MG5300_series, printers/Xerox_Phaser_7600, ipp/print, Epson_IPP_Printer.
func (m *AirPrintDestination) SetResourcePath(value *string)() {
    m.resourcePath = value
}
