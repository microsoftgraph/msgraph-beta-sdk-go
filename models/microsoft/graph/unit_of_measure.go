package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnitOfMeasure 
type UnitOfMeasure struct {
    Entity
    // 
    code *string;
    // 
    displayName *string;
    // 
    internationalStandardCode *string;
    // 
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// NewUnitOfMeasure instantiates a new unitOfMeasure and sets the default values.
func NewUnitOfMeasure()(*UnitOfMeasure) {
    m := &UnitOfMeasure{
        Entity: *NewEntity(),
    }
    return m
}
// GetCode gets the code property value. 
func (m *UnitOfMeasure) GetCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.code
    }
}
// GetDisplayName gets the displayName property value. 
func (m *UnitOfMeasure) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetInternationalStandardCode gets the internationalStandardCode property value. 
func (m *UnitOfMeasure) GetInternationalStandardCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internationalStandardCode
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. 
func (m *UnitOfMeasure) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnitOfMeasure) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["code"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
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
    res["internationalStandardCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternationalStandardCode(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    return res
}
func (m *UnitOfMeasure) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UnitOfMeasure) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("code", m.GetCode())
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
        err = writer.WriteStringValue("internationalStandardCode", m.GetInternationalStandardCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCode sets the code property value. 
func (m *UnitOfMeasure) SetCode(value *string)() {
    if m != nil {
        m.code = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *UnitOfMeasure) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetInternationalStandardCode sets the internationalStandardCode property value. 
func (m *UnitOfMeasure) SetInternationalStandardCode(value *string)() {
    if m != nil {
        m.internationalStandardCode = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. 
func (m *UnitOfMeasure) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
