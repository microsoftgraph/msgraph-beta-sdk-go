package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i010dc49dba19772e8fc5717aeb2c8bd3769d793dba9b6ac97365ac0e7e64aaec "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities/item/listitem/createlink"
    i0d7f387a0091b172eb2422701cd6963fa0ec0fce5de947565300cdf581fbf834 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities/item/listitem/analytics"
    i379ea2863762eee877ab46370c46b080c4a48588a229303d134a3921152e41db "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities/item/listitem/activities"
    i6cbfdf2d216812f73fe28cae3af73af840ecd7bedd5b1520855fc52cc5871728 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i79f654a6ed4a5d0b3725dd2e301afc339010f5caa6c02b1a2b6f14ac1f706344 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities/item/listitem/fields"
    ia2667dacc41db9c50c92d995c549b6d06f59673da12e451540a5344df523db05 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities/item/listitem/driveitem"
    iad938253dd15bf1ad93d1333b01566a5bb673e6d21b1bc572d9fbb4237a5c28f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities/item/listitem/versions"
    i0c4a10f20b2a807a30807663746302f83008d86c56901414353878e74f1c4d3d "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities/item/listitem/versions/item"
    i60926d3835f9a98f47052035c075913a9c9073f1489d335f12324cd004d64dda "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/activities/item/listitem/activities/item"
)

// listItemRequestBuilder builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\activities\{itemActivityOLD-id}\listItem
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListItemRequestBuilderDeleteOptions options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemRequestBuilderGetOptions options for Get
type ListItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// listItemRequestBuilderGetQueryParameters get listItem from sites
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// ListItemRequestBuilderPatchOptions options for Patch
type ListItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListItemRequestBuilder) Activities()(*i379ea2863762eee877ab46370c46b080c4a48588a229303d134a3921152e41db.ActivitiesRequestBuilder) {
    return i379ea2863762eee877ab46370c46b080c4a48588a229303d134a3921152e41db.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.activities.item.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*i60926d3835f9a98f47052035c075913a9c9073f1489d335f12324cd004d64dda.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id1"] = id
    }
    return i60926d3835f9a98f47052035c075913a9c9073f1489d335f12324cd004d64dda.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*i0d7f387a0091b172eb2422701cd6963fa0ec0fce5de947565300cdf581fbf834.AnalyticsRequestBuilder) {
    return i0d7f387a0091b172eb2422701cd6963fa0ec0fce5de947565300cdf581fbf834.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/lists/{list_id}/activities/{itemActivityOLD_id}/listItem{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemRequestBuilder instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property listItem for sites
func (m *ListItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get listItem from sites
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
func (m *ListItemRequestBuilder) CreateLink()(*i010dc49dba19772e8fc5717aeb2c8bd3769d793dba9b6ac97365ac0e7e64aaec.CreateLinkRequestBuilder) {
    return i010dc49dba19772e8fc5717aeb2c8bd3769d793dba9b6ac97365ac0e7e64aaec.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property listItem in sites
func (m *ListItemRequestBuilder) CreatePatchRequestInformation(options *ListItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property listItem for sites
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) DriveItem()(*ia2667dacc41db9c50c92d995c549b6d06f59673da12e451540a5344df523db05.DriveItemRequestBuilder) {
    return ia2667dacc41db9c50c92d995c549b6d06f59673da12e451540a5344df523db05.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i79f654a6ed4a5d0b3725dd2e301afc339010f5caa6c02b1a2b6f14ac1f706344.FieldsRequestBuilder) {
    return i79f654a6ed4a5d0b3725dd2e301afc339010f5caa6c02b1a2b6f14ac1f706344.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get listItem from sites
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewListItem() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval builds and executes requests for operations under \sites\{site-id}\lists\{list-id}\activities\{itemActivityOLD-id}\listItem\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*i6cbfdf2d216812f73fe28cae3af73af840ecd7bedd5b1520855fc52cc5871728.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i6cbfdf2d216812f73fe28cae3af73af840ecd7bedd5b1520855fc52cc5871728.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// Patch update the navigation property listItem in sites
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) Versions()(*iad938253dd15bf1ad93d1333b01566a5bb673e6d21b1bc572d9fbb4237a5c28f.VersionsRequestBuilder) {
    return iad938253dd15bf1ad93d1333b01566a5bb673e6d21b1bc572d9fbb4237a5c28f.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.activities.item.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*i0c4a10f20b2a807a30807663746302f83008d86c56901414353878e74f1c4d3d.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return i0c4a10f20b2a807a30807663746302f83008d86c56901414353878e74f1c4d3d.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
