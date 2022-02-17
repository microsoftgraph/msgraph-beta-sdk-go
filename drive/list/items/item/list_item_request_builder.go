package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i3494ba260c7e5cb742941c392b00916ff749fe19697d06484c3b7083eabe7c49 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/fields"
    i4b6b4c3281893ceaa0871aa95c8b6c1ae470e33bc991bfd9ff48b8116029e68a "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i58c8d16874aac3827dc40a3dbf2328fea48f08417c9c41debf67a58f919c2720 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/driveitem"
    i5ba57fe34d693b4cded862b4e7ada5fe09f96c4be8b524c005dc2feda37966d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/activities"
    ia2b700d43ac5cd289f1eae1bb612360c49638c62dfb54b8c7e8228d6fd206e25 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/analytics"
    ic5c7e5c1201df8e08dd3c31a725a1b6b5c4b8c93b46ad9390ad32b9a00aa10a1 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/versions"
    icf63e1dabde99d58f91beae56e8ed47fc1a61a3ec747a3742ff506ffd784cb39 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/createlink"
    id15ff120ad8f053cfd1691895b73c567a4d9203be255a2354e8e1c0a99cce095 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/activities/item"
    ie00e09a3d9b974ec973d38ac015be59a09a07678fa7822e52487a59cc84f7610 "github.com/microsoftgraph/msgraph-beta-sdk-go/drive/list/items/item/versions/item"
)

// ListItemRequestBuilder builds and executes requests for operations under \drive\list\items\{listItem-id}
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
// ListItemRequestBuilderGetQueryParameters all items contained in the list.
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
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
func (m *ListItemRequestBuilder) Activities()(*i5ba57fe34d693b4cded862b4e7ada5fe09f96c4be8b524c005dc2feda37966d7.ActivitiesRequestBuilder) {
    return i5ba57fe34d693b4cded862b4e7ada5fe09f96c4be8b524c005dc2feda37966d7.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.items.item.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*id15ff120ad8f053cfd1691895b73c567a4d9203be255a2354e8e1c0a99cce095.ItemActivityOLDRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return id15ff120ad8f053cfd1691895b73c567a4d9203be255a2354e8e1c0a99cce095.NewItemActivityOLDRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*ia2b700d43ac5cd289f1eae1bb612360c49638c62dfb54b8c7e8228d6fd206e25.AnalyticsRequestBuilder) {
    return ia2b700d43ac5cd289f1eae1bb612360c49638c62dfb54b8c7e8228d6fd206e25.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/list/items/{listItem_id}{?select,expand}";
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
// CreateDeleteRequestInformation all items contained in the list.
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
// CreateGetRequestInformation all items contained in the list.
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
func (m *ListItemRequestBuilder) CreateLink()(*icf63e1dabde99d58f91beae56e8ed47fc1a61a3ec747a3742ff506ffd784cb39.CreateLinkRequestBuilder) {
    return icf63e1dabde99d58f91beae56e8ed47fc1a61a3ec747a3742ff506ffd784cb39.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation all items contained in the list.
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
// Delete all items contained in the list.
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) DriveItem()(*i58c8d16874aac3827dc40a3dbf2328fea48f08417c9c41debf67a58f919c2720.DriveItemRequestBuilder) {
    return i58c8d16874aac3827dc40a3dbf2328fea48f08417c9c41debf67a58f919c2720.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i3494ba260c7e5cb742941c392b00916ff749fe19697d06484c3b7083eabe7c49.FieldsRequestBuilder) {
    return i3494ba260c7e5cb742941c392b00916ff749fe19697d06484c3b7083eabe7c49.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the list.
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewListItem() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItem), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval builds and executes requests for operations under \drive\list\items\{listItem-id}\microsoft.graph.getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}')
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, endDateTime *string, interval *string)(*i4b6b4c3281893ceaa0871aa95c8b6c1ae470e33bc991bfd9ff48b8116029e68a.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i4b6b4c3281893ceaa0871aa95c8b6c1ae470e33bc991bfd9ff48b8116029e68a.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, endDateTime, interval);
}
// Patch all items contained in the list.
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ListItemRequestBuilder) Versions()(*ic5c7e5c1201df8e08dd3c31a725a1b6b5c4b8c93b46ad9390ad32b9a00aa10a1.VersionsRequestBuilder) {
    return ic5c7e5c1201df8e08dd3c31a725a1b6b5c4b8c93b46ad9390ad32b9a00aa10a1.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.drive.list.items.item.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*ie00e09a3d9b974ec973d38ac015be59a09a07678fa7822e52487a59cc84f7610.ListItemVersionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return ie00e09a3d9b974ec973d38ac015be59a09a07678fa7822e52487a59cc84f7610.NewListItemVersionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
