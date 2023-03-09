package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// PrinterUsageSummary 
type PrinterUsageSummary struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewPrinterUsageSummary instantiates a new printerUsageSummary and sets the default values.
func NewPrinterUsageSummary()(*PrinterUsageSummary) {
    m := &PrinterUsageSummary{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrinterUsageSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterUsageSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterUsageSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterUsageSummary) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *PrinterUsageSummary) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCompletedJobCount gets the completedJobCount property value. The completedJobCount property
func (m *PrinterUsageSummary) GetCompletedJobCount()(*int32) {
    val, err := m.GetBackingStore().Get("completedJobCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterUsageSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["completedJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedJobCount(val)
        }
        return nil
    }
    res["incompleteJobCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncompleteJobCount(val)
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
    res["printer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinter(val.(DirectoryObjectable))
        }
        return nil
    }
    res["printerDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterDisplayName(val)
        }
        return nil
    }
    res["printerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterId(val)
        }
        return nil
    }
    res["printerManufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterManufacturer(val)
        }
        return nil
    }
    res["printerModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrinterModel(val)
        }
        return nil
    }
    return res
}
// GetIncompleteJobCount gets the incompleteJobCount property value. The incompleteJobCount property
func (m *PrinterUsageSummary) GetIncompleteJobCount()(*int32) {
    val, err := m.GetBackingStore().Get("incompleteJobCount")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrinterUsageSummary) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrinter gets the printer property value. The printer property
func (m *PrinterUsageSummary) GetPrinter()(DirectoryObjectable) {
    val, err := m.GetBackingStore().Get("printer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DirectoryObjectable)
    }
    return nil
}
// GetPrinterDisplayName gets the printerDisplayName property value. The printerDisplayName property
func (m *PrinterUsageSummary) GetPrinterDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("printerDisplayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrinterId gets the printerId property value. The printerId property
func (m *PrinterUsageSummary) GetPrinterId()(*string) {
    val, err := m.GetBackingStore().Get("printerId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrinterManufacturer gets the printerManufacturer property value. The printerManufacturer property
func (m *PrinterUsageSummary) GetPrinterManufacturer()(*string) {
    val, err := m.GetBackingStore().Get("printerManufacturer")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPrinterModel gets the printerModel property value. The printerModel property
func (m *PrinterUsageSummary) GetPrinterModel()(*string) {
    val, err := m.GetBackingStore().Get("printerModel")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *PrinterUsageSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("completedJobCount", m.GetCompletedJobCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("incompleteJobCount", m.GetIncompleteJobCount())
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
        err := writer.WriteObjectValue("printer", m.GetPrinter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("printerDisplayName", m.GetPrinterDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("printerId", m.GetPrinterId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("printerManufacturer", m.GetPrinterManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("printerModel", m.GetPrinterModel())
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
func (m *PrinterUsageSummary) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *PrinterUsageSummary) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetCompletedJobCount sets the completedJobCount property value. The completedJobCount property
func (m *PrinterUsageSummary) SetCompletedJobCount(value *int32)() {
    err := m.GetBackingStore().Set("completedJobCount", value)
    if err != nil {
        panic(err)
    }
}
// SetIncompleteJobCount sets the incompleteJobCount property value. The incompleteJobCount property
func (m *PrinterUsageSummary) SetIncompleteJobCount(value *int32)() {
    err := m.GetBackingStore().Set("incompleteJobCount", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrinterUsageSummary) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetPrinter sets the printer property value. The printer property
func (m *PrinterUsageSummary) SetPrinter(value DirectoryObjectable)() {
    err := m.GetBackingStore().Set("printer", value)
    if err != nil {
        panic(err)
    }
}
// SetPrinterDisplayName sets the printerDisplayName property value. The printerDisplayName property
func (m *PrinterUsageSummary) SetPrinterDisplayName(value *string)() {
    err := m.GetBackingStore().Set("printerDisplayName", value)
    if err != nil {
        panic(err)
    }
}
// SetPrinterId sets the printerId property value. The printerId property
func (m *PrinterUsageSummary) SetPrinterId(value *string)() {
    err := m.GetBackingStore().Set("printerId", value)
    if err != nil {
        panic(err)
    }
}
// SetPrinterManufacturer sets the printerManufacturer property value. The printerManufacturer property
func (m *PrinterUsageSummary) SetPrinterManufacturer(value *string)() {
    err := m.GetBackingStore().Set("printerManufacturer", value)
    if err != nil {
        panic(err)
    }
}
// SetPrinterModel sets the printerModel property value. The printerModel property
func (m *PrinterUsageSummary) SetPrinterModel(value *string)() {
    err := m.GetBackingStore().Set("printerModel", value)
    if err != nil {
        panic(err)
    }
}
// PrinterUsageSummaryable 
type PrinterUsageSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetCompletedJobCount()(*int32)
    GetIncompleteJobCount()(*int32)
    GetOdataType()(*string)
    GetPrinter()(DirectoryObjectable)
    GetPrinterDisplayName()(*string)
    GetPrinterId()(*string)
    GetPrinterManufacturer()(*string)
    GetPrinterModel()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetCompletedJobCount(value *int32)()
    SetIncompleteJobCount(value *int32)()
    SetOdataType(value *string)()
    SetPrinter(value DirectoryObjectable)()
    SetPrinterDisplayName(value *string)()
    SetPrinterId(value *string)()
    SetPrinterManufacturer(value *string)()
    SetPrinterModel(value *string)()
}
