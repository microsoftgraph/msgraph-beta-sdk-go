package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterCreateOperation 
type PrinterCreateOperation struct {
    PrintOperation
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The signed certificate created during the registration process. Read-only.
    certificate *string
    // The created printer entity. Read-only.
    printer Printerable
}
// NewPrinterCreateOperation instantiates a new PrinterCreateOperation and sets the default values.
func NewPrinterCreateOperation()(*PrinterCreateOperation) {
    m := &PrinterCreateOperation{
        PrintOperation: *NewPrintOperation(),
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePrinterCreateOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterCreateOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterCreateOperation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterCreateOperation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificate gets the certificate property value. The signed certificate created during the registration process. Read-only.
func (m *PrinterCreateOperation) GetCertificate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.certificate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterCreateOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PrintOperation.GetFieldDeserializers()
    res["certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["printer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinter(val.(Printerable))
        }
        return nil
    }
    return res
}
// GetPrinter gets the printer property value. The created printer entity. Read-only.
func (m *PrinterCreateOperation) GetPrinter()(Printerable) {
    if m == nil {
        return nil
    } else {
        return m.printer
    }
}
// Serialize serializes information the current object
func (m *PrinterCreateOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PrintOperation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("printer", m.GetPrinter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterCreateOperation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertificate sets the certificate property value. The signed certificate created during the registration process. Read-only.
func (m *PrinterCreateOperation) SetCertificate(value *string)() {
    if m != nil {
        m.certificate = value
    }
}
// SetPrinter sets the printer property value. The created printer entity. Read-only.
func (m *PrinterCreateOperation) SetPrinter(value Printerable)() {
    if m != nil {
        m.printer = value
    }
}
