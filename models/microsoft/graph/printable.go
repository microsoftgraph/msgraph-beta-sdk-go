package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Printable 
type Printable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetConnectors()([]PrintConnectorable)
    GetOperations()([]PrintOperationable)
    GetPrinters()([]Printerable)
    GetPrinterShares()([]PrinterShareable)
    GetReports()(ReportRootable)
    GetServices()([]PrintServiceable)
    GetSettings()(PrintSettingsable)
    GetShares()([]PrinterShareable)
    GetTaskDefinitions()([]PrintTaskDefinitionable)
    SetConnectors(value []PrintConnectorable)()
    SetOperations(value []PrintOperationable)()
    SetPrinters(value []Printerable)()
    SetPrinterShares(value []PrinterShareable)()
    SetReports(value ReportRootable)()
    SetServices(value []PrintServiceable)()
    SetSettings(value PrintSettingsable)()
    SetShares(value []PrinterShareable)()
    SetTaskDefinitions(value []PrintTaskDefinitionable)()
}
