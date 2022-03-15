package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintDocument provides operations to manage the print singleton.
type PrintDocument struct {
    Entity
    // 
    configuration PrinterDocumentConfigurationable;
    // The document's content (MIME) type. Read-only.
    contentType *string;
    // The document's name. Read-only.
    displayName *string;
    // The document's size in bytes. Read-only.
    size *int32;
}
// NewPrintDocument instantiates a new printDocument and sets the default values.
func NewPrintDocument()(*PrintDocument) {
    m := &PrintDocument{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintDocumentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintDocumentFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPrintDocument(), nil
}
// GetConfiguration gets the configuration property value. 
func (m *PrintDocument) GetConfiguration()(PrinterDocumentConfigurationable) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// GetContentType gets the contentType property value. The document's content (MIME) type. Read-only.
func (m *PrintDocument) GetContentType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentType
    }
}
// GetDisplayName gets the displayName property value. The document's name. Read-only.
func (m *PrintDocument) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintDocument) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrinterDocumentConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(PrinterDocumentConfigurationable))
        }
        return nil
    }
    res["contentType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["size"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetSize gets the size property value. The document's size in bytes. Read-only.
func (m *PrintDocument) GetSize()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.size
    }
}
func (m *PrintDocument) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PrintDocument) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("size", m.GetSize())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfiguration sets the configuration property value. 
func (m *PrintDocument) SetConfiguration(value PrinterDocumentConfigurationable)() {
    if m != nil {
        m.configuration = value
    }
}
// SetContentType sets the contentType property value. The document's content (MIME) type. Read-only.
func (m *PrintDocument) SetContentType(value *string)() {
    if m != nil {
        m.contentType = value
    }
}
// SetDisplayName sets the displayName property value. The document's name. Read-only.
func (m *PrintDocument) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetSize sets the size property value. The document's size in bytes. Read-only.
func (m *PrintDocument) SetSize(value *int32)() {
    if m != nil {
        m.size = value
    }
}
