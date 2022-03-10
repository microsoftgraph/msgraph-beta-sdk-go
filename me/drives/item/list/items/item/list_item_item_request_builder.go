package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i3953319bb4f6946dded4210b880b12146634491db703c93f7c4a1a89f3116dc4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item/versions"
    i5dc11f251693b55b65c4b33be26a2f295faee4dd38566d35f34d4b4db8eafd48 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item/fields"
    i6be348558e2e2939c7c51954f64ae097874959d2eb8ec667664468ac5a1c14b0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i770d17a7d84a22fa51739215adb2fa1c51e3dd0cd9708191b7da5634d727e0c6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item/createlink"
    i9f0cb343a75a3e872e96ef1b16ce4d1e8a2bba59df4a7b37178a7781d2af40c4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item/activities"
    i9f72ad44bfad94c65e465baf7c19ecb61b5360563d4e653434fe9d61005b39d6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item/driveitem"
    ifa894bccfd80baedc75d8f728ba78f8f9c25d619dd0f9746dc0ed73c62923429 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item/analytics"
    i208e797228f4a34e9090e121412a62a8ddd78d9f86717f935b72540268b8b613 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item/activities/item"
    iac4d302efed8ae0cdc34fd3f0abda2c5025b48759343d2a7a1b4b83d581b4875 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item/versions/item"
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
func (m *ListItemItemRequestBuilder) Activities()(*i9f0cb343a75a3e872e96ef1b16ce4d1e8a2bba59df4a7b37178a7781d2af40c4.ActivitiesRequestBuilder) {
    return i9f0cb343a75a3e872e96ef1b16ce4d1e8a2bba59df4a7b37178a7781d2af40c4.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.items.item.activities.item collection
func (m *ListItemItemRequestBuilder) ActivitiesById(id string)(*i208e797228f4a34e9090e121412a62a8ddd78d9f86717f935b72540268b8b613.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i208e797228f4a34e9090e121412a62a8ddd78d9f86717f935b72540268b8b613.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemItemRequestBuilder) Analytics()(*ifa894bccfd80baedc75d8f728ba78f8f9c25d619dd0f9746dc0ed73c62923429.AnalyticsRequestBuilder) {
    return ifa894bccfd80baedc75d8f728ba78f8f9c25d619dd0f9746dc0ed73c62923429.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemItemRequestBuilderInternal instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemItemRequestBuilder) {
    m := &ListItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/list/items/{listItem_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property items for me
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
func (m *ListItemItemRequestBuilder) CreateLink()(*i770d17a7d84a22fa51739215adb2fa1c51e3dd0cd9708191b7da5634d727e0c6.CreateLinkRequestBuilder) {
    return i770d17a7d84a22fa51739215adb2fa1c51e3dd0cd9708191b7da5634d727e0c6.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in me
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
// Delete delete navigation property items for me
func (m *ListItemItemRequestBuilder) Delete(options *ListItemItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemItemRequestBuilder) DriveItem()(*i9f72ad44bfad94c65e465baf7c19ecb61b5360563d4e653434fe9d61005b39d6.DriveItemRequestBuilder) {
    return i9f72ad44bfad94c65e465baf7c19ecb61b5360563d4e653434fe9d61005b39d6.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemItemRequestBuilder) Fields()(*i5dc11f251693b55b65c4b33be26a2f295faee4dd38566d35f34d4b4db8eafd48.FieldsRequestBuilder) {
    return i5dc11f251693b55b65c4b33be26a2f295faee4dd38566d35f34d4b4db8eafd48.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the list.
func (m *ListItemItemRequestBuilder) Get(options *ListItemItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateListItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i6be348558e2e2939c7c51954f64ae097874959d2eb8ec667664468ac5a1c14b0.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i6be348558e2e2939c7c51954f64ae097874959d2eb8ec667664468ac5a1c14b0.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Patch update the navigation property items in me
func (m *ListItemItemRequestBuilder) Patch(options *ListItemItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemItemRequestBuilder) Versions()(*i3953319bb4f6946dded4210b880b12146634491db703c93f7c4a1a89f3116dc4.VersionsRequestBuilder) {
    return i3953319bb4f6946dded4210b880b12146634491db703c93f7c4a1a89f3116dc4.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.items.item.versions.item collection
func (m *ListItemItemRequestBuilder) VersionsById(id string)(*iac4d302efed8ae0cdc34fd3f0abda2c5025b48759343d2a7a1b4b83d581b4875.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return iac4d302efed8ae0cdc34fd3f0abda2c5025b48759343d2a7a1b4b83d581b4875.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
