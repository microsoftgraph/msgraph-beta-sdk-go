package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ExactMatchDataStoreBase 
type ExactMatchDataStoreBase struct {
    Entity
    // 
    columns []ExactDataMatchStoreColumn;
    // 
    dataLastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    description *string;
    // 
    displayName *string;
}
// NewExactMatchDataStoreBase instantiates a new exactMatchDataStoreBase and sets the default values.
func NewExactMatchDataStoreBase()(*ExactMatchDataStoreBase) {
    m := &ExactMatchDataStoreBase{
        Entity: *NewEntity(),
    }
    return m
}
// GetColumns gets the columns property value. 
func (m *ExactMatchDataStoreBase) GetColumns()([]ExactDataMatchStoreColumn) {
    if m == nil {
        return nil
    } else {
        return m.columns
    }
}
// GetDataLastUpdatedDateTime gets the dataLastUpdatedDateTime property value. 
func (m *ExactMatchDataStoreBase) GetDataLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.dataLastUpdatedDateTime
    }
}
// GetDescription gets the description property value. 
func (m *ExactMatchDataStoreBase) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. 
func (m *ExactMatchDataStoreBase) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExactMatchDataStoreBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["columns"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExactDataMatchStoreColumn() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactDataMatchStoreColumn, len(val))
            for i, v := range val {
                res[i] = *(v.(*ExactDataMatchStoreColumn))
            }
            m.SetColumns(res)
        }
        return nil
    }
    res["dataLastUpdatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataLastUpdatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    return res
}
func (m *ExactMatchDataStoreBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ExactMatchDataStoreBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColumns() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetColumns()))
        for i, v := range m.GetColumns() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("columns", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("dataLastUpdatedDateTime", m.GetDataLastUpdatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    return nil
}
// SetColumns sets the columns property value. 
func (m *ExactMatchDataStoreBase) SetColumns(value []ExactDataMatchStoreColumn)() {
    if m != nil {
        m.columns = value
    }
}
// SetDataLastUpdatedDateTime sets the dataLastUpdatedDateTime property value. 
func (m *ExactMatchDataStoreBase) SetDataLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.dataLastUpdatedDateTime = value
    }
}
// SetDescription sets the description property value. 
func (m *ExactMatchDataStoreBase) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *ExactMatchDataStoreBase) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
