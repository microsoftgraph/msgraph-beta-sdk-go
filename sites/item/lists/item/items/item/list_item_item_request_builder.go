package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i01322f6c445fc8cfdd645f8a58a2e9dc2aae2538661d3565a27837d1aa1a3bf3 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/analytics"
    i0d9fa4e07c80bd530f75b82be901e491731fc66da0aea17d127c28e0b96cdd3f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/createlink"
    i70f1efce412deb28482396740c683ea9f4b411c123301ce44883bab61b141a65 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/versions"
    i71f902ff67b40c4be62d504c1d2c0f525058bba7aa6c03f0ad5fea4098c14d1f "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/driveitem"
    i9237c91883a88a40d038ddbc99b525375d2ffafd6ef05bfcdf2ce7a59c7b967d "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/activities"
    ia0af2bad6bfa2334cda9b36762c7754513aad438f8dad54fe36023b80b1086e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/fields"
    ia8cb9d287c711158fa129209470a23d78a03e7501d70219257f78200123e41e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i2b138b210334df4f8b6453d94f09db4128ad42a5036fede3e837fa332b533270 "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/versions/item"
    ieddbfd259f1fbc5f85a15691baa1cba1c9fc88a0b59e0d161ce0c97e208a77bc "github.com/microsoftgraph/msgraph-beta-sdk-go/sites/item/lists/item/items/item/activities/item"
)

// ListItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.list entity.
type ListItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ListItemItemRequestBuilderDeleteOptions options for Delete
type ListItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemItemRequestBuilderGetOptions options for Get
type ListItemItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ListItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ListItemItemRequestBuilderGetQueryParameters all items contained in the list.
type ListItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListItemItemRequestBuilderPatchOptions options for Patch
type ListItemItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListItemItemRequestBuilder) Activities()(*i9237c91883a88a40d038ddbc99b525375d2ffafd6ef05bfcdf2ce7a59c7b967d.ActivitiesRequestBuilder) {
    return i9237c91883a88a40d038ddbc99b525375d2ffafd6ef05bfcdf2ce7a59c7b967d.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.items.item.activities.item collection
func (m *ListItemItemRequestBuilder) ActivitiesById(id string)(*ieddbfd259f1fbc5f85a15691baa1cba1c9fc88a0b59e0d161ce0c97e208a77bc.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ieddbfd259f1fbc5f85a15691baa1cba1c9fc88a0b59e0d161ce0c97e208a77bc.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemItemRequestBuilder) Analytics()(*i01322f6c445fc8cfdd645f8a58a2e9dc2aae2538661d3565a27837d1aa1a3bf3.AnalyticsRequestBuilder) {
    return i01322f6c445fc8cfdd645f8a58a2e9dc2aae2538661d3565a27837d1aa1a3bf3.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemItemRequestBuilderInternal instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemItemRequestBuilder) {
    m := &ListItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/lists/{list_id}/items/{listItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemItemRequestBuilder instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property items for sites
func (m *ListItemItemRequestBuilder) CreateDeleteRequestInformation(options *ListItemItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation all items contained in the list.
func (m *ListItemItemRequestBuilder) CreateGetRequestInformation(options *ListItemItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ListItemItemRequestBuilder) CreateLink()(*i0d9fa4e07c80bd530f75b82be901e491731fc66da0aea17d127c28e0b96cdd3f.CreateLinkRequestBuilder) {
    return i0d9fa4e07c80bd530f75b82be901e491731fc66da0aea17d127c28e0b96cdd3f.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in sites
func (m *ListItemItemRequestBuilder) CreatePatchRequestInformation(options *ListItemItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property items for sites
func (m *ListItemItemRequestBuilder) Delete(options *ListItemItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemItemRequestBuilder) DriveItem()(*i71f902ff67b40c4be62d504c1d2c0f525058bba7aa6c03f0ad5fea4098c14d1f.DriveItemRequestBuilder) {
    return i71f902ff67b40c4be62d504c1d2c0f525058bba7aa6c03f0ad5fea4098c14d1f.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemItemRequestBuilder) Fields()(*ia0af2bad6bfa2334cda9b36762c7754513aad438f8dad54fe36023b80b1086e4.FieldsRequestBuilder) {
    return ia0af2bad6bfa2334cda9b36762c7754513aad438f8dad54fe36023b80b1086e4.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the list.
func (m *ListItemItemRequestBuilder) Get(options *ListItemItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateListItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, interval *string, endDateTime *string)(*ia8cb9d287c711158fa129209470a23d78a03e7501d70219257f78200123e41e5.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return ia8cb9d287c711158fa129209470a23d78a03e7501d70219257f78200123e41e5.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, interval, endDateTime);
}
// Patch update the navigation property items in sites
func (m *ListItemItemRequestBuilder) Patch(options *ListItemItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
        "5XX": i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemItemRequestBuilder) Versions()(*i70f1efce412deb28482396740c683ea9f4b411c123301ce44883bab61b141a65.VersionsRequestBuilder) {
    return i70f1efce412deb28482396740c683ea9f4b411c123301ce44883bab61b141a65.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.sites.item.lists.item.items.item.versions.item collection
func (m *ListItemItemRequestBuilder) VersionsById(id string)(*i2b138b210334df4f8b6453d94f09db4128ad42a5036fede3e837fa332b533270.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return i2b138b210334df4f8b6453d94f09db4128ad42a5036fede3e837fa332b533270.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
