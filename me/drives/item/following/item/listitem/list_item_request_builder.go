package listitem

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0200062a34da2e84da1ba38ed53716a6c335e29393ff891ed46e7e1998d7e337 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/listitem/analytics"
    i029e39e1a489a361591dd80f03eafd7563b0f80ff694c6163f1f2c6a4adfe817 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/listitem/fields"
    i8850711d1b0d99eed71bf65c8a4094b44165a3efd97988c645ce0ad8a1ff1016 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/listitem/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i94f2aa247f8c68fa04f71e891c7160553d0a921cf78b96328cff4b1e2128bf6b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/listitem/versions"
    iab8ade9795bb0109e89b249fb54e31570dc4adda833905738c24457e467c0021 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/listitem/driveitem"
    id8456d02ae67e5174e4636d0aeb1affdf3ef84fb6b3c4dc4626e2d9e03da9530 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/listitem/activities"
    iec364c7e023b567ed86ea8e2b9f19747e41a6f4feab7d1ea79c32ef1ec546e81 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/listitem/createlink"
    i367f76f144ed2ea129ecbec2b4a381499d73810e16d2932164382706f7cf132a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/listitem/activities/item"
    if17a2c1be25f601966e2ffef30c4847cc60f1e5d7639fedbaa462f9173d97a79 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/following/item/listitem/versions/item"
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
func (m *ListItemRequestBuilder) Activities()(*id8456d02ae67e5174e4636d0aeb1affdf3ef84fb6b3c4dc4626e2d9e03da9530.ActivitiesRequestBuilder) {
    return id8456d02ae67e5174e4636d0aeb1affdf3ef84fb6b3c4dc4626e2d9e03da9530.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.following.item.listItem.activities.item collection
func (m *ListItemRequestBuilder) ActivitiesById(id string)(*i367f76f144ed2ea129ecbec2b4a381499d73810e16d2932164382706f7cf132a.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD_id"] = id
    }
    return i367f76f144ed2ea129ecbec2b4a381499d73810e16d2932164382706f7cf132a.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Analytics()(*i0200062a34da2e84da1ba38ed53716a6c335e29393ff891ed46e7e1998d7e337.AnalyticsRequestBuilder) {
    return i0200062a34da2e84da1ba38ed53716a6c335e29393ff891ed46e7e1998d7e337.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemRequestBuilderInternal instantiates a new ListItemRequestBuilder and sets the default values.
func NewListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ListItemRequestBuilder) {
    m := &ListItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/following/{driveItem_id}/listItem{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property listItem for me
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
func (m *ListItemRequestBuilder) CreateLink()(*iec364c7e023b567ed86ea8e2b9f19747e41a6f4feab7d1ea79c32ef1ec546e81.CreateLinkRequestBuilder) {
    return iec364c7e023b567ed86ea8e2b9f19747e41a6f4feab7d1ea79c32ef1ec546e81.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property listItem in me
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
// Delete delete navigation property listItem for me
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
func (m *ListItemRequestBuilder) DriveItem()(*iab8ade9795bb0109e89b249fb54e31570dc4adda833905738c24457e467c0021.DriveItemRequestBuilder) {
    return iab8ade9795bb0109e89b249fb54e31570dc4adda833905738c24457e467c0021.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ListItemRequestBuilder) Fields()(*i029e39e1a489a361591dd80f03eafd7563b0f80ff694c6163f1f2c6a4adfe817.FieldsRequestBuilder) {
    return i029e39e1a489a361591dd80f03eafd7563b0f80ff694c6163f1f2c6a4adfe817.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ListItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(startDateTime *string, interval *string, endDateTime *string)(*i8850711d1b0d99eed71bf65c8a4094b44165a3efd97988c645ce0ad8a1ff1016.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i8850711d1b0d99eed71bf65c8a4094b44165a3efd97988c645ce0ad8a1ff1016.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, startDateTime, interval, endDateTime);
}
// Patch update the navigation property listItem in me
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
func (m *ListItemRequestBuilder) Versions()(*i94f2aa247f8c68fa04f71e891c7160553d0a921cf78b96328cff4b1e2128bf6b.VersionsRequestBuilder) {
    return i94f2aa247f8c68fa04f71e891c7160553d0a921cf78b96328cff4b1e2128bf6b.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.following.item.listItem.versions.item collection
func (m *ListItemRequestBuilder) VersionsById(id string)(*if17a2c1be25f601966e2ffef30c4847cc60f1e5d7639fedbaa462f9173d97a79.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion_id"] = id
    }
    return if17a2c1be25f601966e2ffef30c4847cc60f1e5d7639fedbaa462f9173d97a79.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
