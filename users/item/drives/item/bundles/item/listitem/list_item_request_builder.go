package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i11a6f420c9474b8569f2b8f004572276bed5b944b6e266f484043d979e53cc5f "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/listitem/analytics"
    i75e65a14bab941fb97fc27bc4bdc083f66aaf94ac22be9ae200ae0b766580921 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i79881ff60809f61c1dbb019305f30cf4a69da4fde59e2d4b60be93001c9d413b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/listitem/fields"
    ic0d3cc30c65505d87a30eb7a35ea4eb294d8c6e27d3d0b6c8530fdbe82a34a2a "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/listitem/driveitem"
    idb4e9ed588ecf69e7a50c304ad3b4fcfe96a2976dcca7032495adfdae0b5c363 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/listitem/activities"
    ie2854b8fa35daa57302512503283e466068f22b5a1851f7a750b95cbf06cefb8 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/listitem/createlink"
    iea4588b50ad7daa2a6e2141429da859ed4d8c190214a708caffffb1471f5088d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/listitem/versions"
    i19bf9bb5f05ef964d88e844bf894258355a92e478573a7e44043d027892cb932 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/listitem/versions/item"
    ibbf8a953744fd0cbd48b460bec676bb52105cc8312dbc8ff0990a28585fa7f4b "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/drives/item/bundles/item/listitem/activities/item"
)

// ListItemRequestBuilder provides operations to manage the listItem property of the microsoft.graph.driveItem entity.
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
// ListItemRequestBuilderGetQueryParameters for drives in SharePoint, the associated document library list item. Read-only. Nullable.
type ListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ListItemRequestBuilderPatchOptions options for Patch
type ListItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ListItemRequestBuilder) Activities()(*idb4e9ed588ecf69e7a50c304ad3b4fcfe96a2976dcca7032495adfdae0b5c363.ActivitiesRequestBuilder) {
    return idb4e9ed588ecf69e7a50c304ad3b4fcfe96a2976dcca7032495adfdae0b5c363.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.bundles.item.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*ibbf8a953744fd0cbd48b460bec676bb52105cc8312dbc8ff0990a28585fa7f4b.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return ibbf8a953744fd0cbd48b460bec676bb52105cc8312dbc8ff0990a28585fa7f4b.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*i11a6f420c9474b8569f2b8f004572276bed5b944b6e266f484043d979e53cc5f.AnalyticsRequestBuilder) {
    return i11a6f420c9474b8569f2b8f004572276bed5b944b6e266f484043d979e53cc5f.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/bundles/{driveItem_id}/listItem{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemRequestBuilder instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property listItem for users
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
// CreateGetRequestInformation for drives in SharePoint, the associated document library list item. Read-only. Nullable.
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
func (m *ListItemRequestBuilder) CreateLink()(*ie2854b8fa35daa57302512503283e466068f22b5a1851f7a750b95cbf06cefb8.CreateLinkRequestBuilder) {
    return ie2854b8fa35daa57302512503283e466068f22b5a1851f7a750b95cbf06cefb8.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property listItem in users
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
// Delete delete navigation property listItem for users
func (m *ListItemRequestBuilder) Delete(options *ListItemRequestBuilderDeleteOptions)(error) {
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
func (m *ListItemRequestBuilder) DriveItem()(*ic0d3cc30c65505d87a30eb7a35ea4eb294d8c6e27d3d0b6c8530fdbe82a34a2a.DriveItemRequestBuilder) {
    return ic0d3cc30c65505d87a30eb7a35ea4eb294d8c6e27d3d0b6c8530fdbe82a34a2a.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i79881ff60809f61c1dbb019305f30cf4a69da4fde59e2d4b60be93001c9d413b.FieldsRequestBuilder) {
    return i79881ff60809f61c1dbb019305f30cf4a69da4fde59e2d4b60be93001c9d413b.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the associated document library list item. Read-only. Nullable.
func (m *ListItemRequestBuilder) Get(options *ListItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.ListItemable, error) {
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
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, interval *string, endDateTime *string)(*i75e65a14bab941fb97fc27bc4bdc083f66aaf94ac22be9ae200ae0b766580921.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i75e65a14bab941fb97fc27bc4bdc083f66aaf94ac22be9ae200ae0b766580921.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, interval, endDateTime);
}
// Patch update the navigation property listItem in users
func (m *ListItemRequestBuilder) Patch(options *ListItemRequestBuilderPatchOptions)(error) {
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
func (m *ListItemRequestBuilder) Versions()(*iea4588b50ad7daa2a6e2141429da859ed4d8c190214a708caffffb1471f5088d.VersionsRequestBuilder) {
    return iea4588b50ad7daa2a6e2141429da859ed4d8c190214a708caffffb1471f5088d.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.users.item.drives.item.bundles.item.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*i19bf9bb5f05ef964d88e844bf894258355a92e478573a7e44043d027892cb932.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return i19bf9bb5f05ef964d88e844bf894258355a92e478573a7e44043d027892cb932.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
