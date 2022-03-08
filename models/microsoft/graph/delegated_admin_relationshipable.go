package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DelegatedAdminRelationshipable 
type DelegatedAdminRelationshipable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessAssignments()([]DelegatedAdminAccessAssignmentable)
    GetAccessDetails()(DelegatedAdminAccessDetailsable)
    GetActivatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomer()(DelegatedAdminRelationshipCustomerParticipantable)
    GetDisplayName()(*string)
    GetDuration()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOperations()([]DelegatedAdminRelationshipOperationable)
    GetPartner()(DelegatedAdminRelationshipParticipantable)
    GetRequests()([]DelegatedAdminRelationshipRequestable)
    GetStatus()(*DelegatedAdminRelationshipStatus)
    SetAccessAssignments(value []DelegatedAdminAccessAssignmentable)()
    SetAccessDetails(value DelegatedAdminAccessDetailsable)()
    SetActivatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomer(value DelegatedAdminRelationshipCustomerParticipantable)()
    SetDisplayName(value *string)()
    SetDuration(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOperations(value []DelegatedAdminRelationshipOperationable)()
    SetPartner(value DelegatedAdminRelationshipParticipantable)()
    SetRequests(value []DelegatedAdminRelationshipRequestable)()
    SetStatus(value *DelegatedAdminRelationshipStatus)()
}
