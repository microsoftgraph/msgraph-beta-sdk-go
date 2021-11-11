package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0a7cbb87a4d2bdb67c30df880c3fc5cecbdf2330792709c09667244a8b58fac2 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item/listitem/activities"
    i2b3780d2c363353054ee38d1d97069a285589a556b8dcdc29729426763cc0f51 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item/listitem/driveitem"
    i35d4b8829c70ec95512664c442947e8bcc3bee584da93f2084f339c44dfe92ed "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item/listitem/versions"
    i4700a0531c15b1ed9f3ad94956909ca89dacbad7bd2fb658c4016bf845f43d44 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item/listitem/fields"
    ib5c409fb32e630ffe49eb014bc819923bbd4b77ae6fca8f04fb30b50aef0b9eb "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    id4937d65654e2b7792bf059f8f283aff66fde293bd5a22f99f5de24e41b50122 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item/listitem/analytics"
    ieee107689f5953959af784a086b085d414e204b6d27d5ed887d61be61bc7c50d "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item/listitem/createlink"
    i3e2a6f6faadefe21527ade5d3b9a7bf5290619478072b9301cb33c0ecf729e02 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item/listitem/activities/item"
    if33df3aa0a25e948de61aaffcbdc27f8be4574ea44225a216ecc03d2224bc519 "github.com/microsoftgraph/msgraph-beta-sdk-go/drives/item/activities/item/listitem/versions/item"
)

// Builds and executes requests for operations under \drives\{drive-id}\activities\{itemActivityOLD-id}\listItem
type ListItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ListItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
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
// Get listItem from drives
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
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
func (m *ListItemRequestBuilder) Activities()(*i0a7cbb87a4d2bdb67c30df880c3fc5cecbdf2330792709c09667244a8b58fac2.ActivitiesRequestBuilder) {
    return i0a7cbb87a4d2bdb67c30df880c3fc5cecbdf2330792709c09667244a8b58fac2.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.activities.item.listItem.activities.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*i3e2a6f6faadefe21527ade5d3b9a7bf5290619478072b9301cb33c0ecf729e02.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id1"] = id
    }
    return i3e2a6f6faadefe21527ade5d3b9a7bf5290619478072b9301cb33c0ecf729e02.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*id4937d65654e2b7792bf059f8f283aff66fde293bd5a22f99f5de24e41b50122.AnalyticsRequestBuilder) {
    return id4937d65654e2b7792bf059f8f283aff66fde293bd5a22f99f5de24e41b50122.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ListItemRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/activities/{itemActivityOLD_id}/listItem{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ListItemRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property listItem for drives
// Parameters:
//  - options : Options for the request
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
// Get listItem from drives
// Parameters:
//  - options : Options for the request
func (m *ListItemRequestBuilder) CreateGetRequestInformation(options *ListItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
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
func (m *ListItemRequestBuilder) CreateLink()(*ieee107689f5953959af784a086b085d414e204b6d27d5ed887d61be61bc7c50d.CreateLinkRequestBuilder) {
    return ieee107689f5953959af784a086b085d414e204b6d27d5ed887d61be61bc7c50d.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update the navigation property listItem in drives
// Parameters:
//  - options : Options for the request
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
// Delete navigation property listItem for drives
// Parameters:
//  - options : Options for the request
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
func (m *ListItemRequestBuilder) DriveItem()(*i2b3780d2c363353054ee38d1d97069a285589a556b8dcdc29729426763cc0f51.DriveItemRequestBuilder) {
    return i2b3780d2c363353054ee38d1d97069a285589a556b8dcdc29729426763cc0f51.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i4700a0531c15b1ed9f3ad94956909ca89dacbad7bd2fb658c4016bf845f43d44.FieldsRequestBuilder) {
    return i4700a0531c15b1ed9f3ad94956909ca89dacbad7bd2fb658c4016bf845f43d44.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get listItem from drives
// Parameters:
//  - options : Options for the request
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
// Builds and executes requests for operations under \drives\{drive-id}\activities\{itemActivityOLD-id}\listItem\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
// Parameters:
//  - endDateTime : Usage: endDateTime={endDateTime}
//  - interval : Usage: interval={interval}
//  - startDateTime : Usage: startDateTime={startDateTime}
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*ib5c409fb32e630ffe49eb014bc819923bbd4b77ae6fca8f04fb30b50aef0b9eb.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ib5c409fb32e630ffe49eb014bc819923bbd4b77ae6fca8f04fb30b50aef0b9eb.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// Update the navigation property listItem in drives
// Parameters:
//  - options : Options for the request
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
func (m *ListItemRequestBuilder) Versions()(*i35d4b8829c70ec95512664c442947e8bcc3bee584da93f2084f339c44dfe92ed.VersionsRequestBuilder) {
    return i35d4b8829c70ec95512664c442947e8bcc3bee584da93f2084f339c44dfe92ed.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drives.item.activities.item.listItem.versions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListItemRequestBuilder) VersionsById(id string)(*if33df3aa0a25e948de61aaffcbdc27f8be4574ea44225a216ecc03d2224bc519.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return if33df3aa0a25e948de61aaffcbdc27f8be4574ea44225a216ecc03d2224bc519.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
