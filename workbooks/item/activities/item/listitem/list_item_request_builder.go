package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i011721b3aff2a1bec1bc4b8072b16e57a077e5bd0ad29d5b0310c08c3e900631 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities/item/listitem/activities"
    i02a07f3a8801be734911d74955aa42fc86122bf2d14a4d0dabf35a1d7e53aa5b "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities/item/listitem/driveitem"
    i25761131b8b9ca038d18339b1aa724218b546344b13c6144db482184368fa588 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities/item/listitem/versions"
    i49da27924e3ddd967033dde973ebce1b68a92a5555db0b6d65fc090332cd4d52 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities/item/listitem/fields"
    i4b6753ea407909cd2e216cd0a71d140437a42e7f41bd28c6175cd361503dcdcb "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities/item/listitem/createlink"
    ib638bfada559e2a13f0d917e127d91685fab15904951d74fa236ad66a4435989 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    ibf619fd0b56b0c1f14c07f1b35180efd40f127d0a2b9522d387c5f734df30007 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities/item/listitem/analytics"
    i60be518f8b88206e19347e0572e5587873a5539ceeee1870ec5955cf06e5dd6c "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities/item/listitem/activities/item"
    ia509ef7f1584295aa65b396b975476f40b1c8a7ca06dfb53d72bc2f9f6cf87c1 "github.com/microsoftgraph/msgraph-beta-sdk-go/workbooks/item/activities/item/listitem/versions/item"
)

// Builds and executes requests for operations under \workbooks\{driveItem-id}\activities\{itemActivityOLD-id}\listItem
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
// Get listItem from workbooks
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
func (m *ListItemRequestBuilder) Activities()(*i011721b3aff2a1bec1bc4b8072b16e57a077e5bd0ad29d5b0310c08c3e900631.ActivitiesRequestBuilder) {
    return i011721b3aff2a1bec1bc4b8072b16e57a077e5bd0ad29d5b0310c08c3e900631.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.activities.item.listItem.activities.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*i60be518f8b88206e19347e0572e5587873a5539ceeee1870ec5955cf06e5dd6c.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id1"] = id
    }
    return i60be518f8b88206e19347e0572e5587873a5539ceeee1870ec5955cf06e5dd6c.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*ibf619fd0b56b0c1f14c07f1b35180efd40f127d0a2b9522d387c5f734df30007.AnalyticsRequestBuilder) {
    return ibf619fd0b56b0c1f14c07f1b35180efd40f127d0a2b9522d387c5f734df30007.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ListItemRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/workbooks/{driveItem_id}/activities/{itemActivityOLD_id}/listItem{?select,expand}";
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
// Delete navigation property listItem for workbooks
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
// Get listItem from workbooks
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
func (m *ListItemRequestBuilder) CreateLink()(*i4b6753ea407909cd2e216cd0a71d140437a42e7f41bd28c6175cd361503dcdcb.CreateLinkRequestBuilder) {
    return i4b6753ea407909cd2e216cd0a71d140437a42e7f41bd28c6175cd361503dcdcb.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update the navigation property listItem in workbooks
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
// Delete navigation property listItem for workbooks
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
func (m *ListItemRequestBuilder) DriveItem()(*i02a07f3a8801be734911d74955aa42fc86122bf2d14a4d0dabf35a1d7e53aa5b.DriveItemRequestBuilder) {
    return i02a07f3a8801be734911d74955aa42fc86122bf2d14a4d0dabf35a1d7e53aa5b.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i49da27924e3ddd967033dde973ebce1b68a92a5555db0b6d65fc090332cd4d52.FieldsRequestBuilder) {
    return i49da27924e3ddd967033dde973ebce1b68a92a5555db0b6d65fc090332cd4d52.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get listItem from workbooks
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
// Builds and executes requests for operations under \workbooks\{driveItem-id}\activities\{itemActivityOLD-id}\listItem\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
// Parameters:
//  - endDateTime : Usage: endDateTime={endDateTime}
//  - interval : Usage: interval={interval}
//  - startDateTime : Usage: startDateTime={startDateTime}
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*ib638bfada559e2a13f0d917e127d91685fab15904951d74fa236ad66a4435989.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ib638bfada559e2a13f0d917e127d91685fab15904951d74fa236ad66a4435989.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// Update the navigation property listItem in workbooks
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
func (m *ListItemRequestBuilder) Versions()(*i25761131b8b9ca038d18339b1aa724218b546344b13c6144db482184368fa588.VersionsRequestBuilder) {
    return i25761131b8b9ca038d18339b1aa724218b546344b13c6144db482184368fa588.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.workbooks.item.activities.item.listItem.versions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ListItemRequestBuilder) VersionsById(id string)(*ia509ef7f1584295aa65b396b975476f40b1c8a7ca06dfb53d72bc2f9f6cf87c1.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return ia509ef7f1584295aa65b396b975476f40b1c8a7ca06dfb53d72bc2f9f6cf87c1.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
