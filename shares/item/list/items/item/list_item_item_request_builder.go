package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0df30f15a37bda4c75a746f1d3a78aeb6d7fedcda8fa91a25630e37a46bb63d3 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/createlink"
    i3d8467aa52be83bf5e27b9b74f21a76f1a5c28b249229ca2c40ba746dc31f6cb "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/analytics"
    i426197eeda24a2ef666da9df0ddff559ab7003700bb784ad8da8009e9521a011 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/driveitem"
    i5245c95090e9c6e9cde87b4772a00a7124828683c592734b734a62638f6f0df8 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/getactivitiesbyintervalwithstartdatetimewithenddatetimewithinterval"
    i5ce7a932ecf70db030b8364f7699aa3822b2ace9797a93c63743a24c7b9c40e2 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/activities"
    i6a33e6735c94a4eec0295b3339373e3d3d94a57d512c4c1297f7ec69a5bcace5 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/fields"
    i7e01e46513a14884280655da10ff1bf036a7401d1f1dd3ee9dbcbe02ffe6757e "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/versions"
    id63b80b47ec3f1a4dd0724f6f1236b0d27e2a5216fe80c55230633ece44c6914 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/documentsetversions"
    i201a560e0b52f390181eee26b5bb385daeff5fecb17cf5d3d44e3139031471e5 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/activities/item"
    idbb42c40e539b6a5f59145da8ae59d604202f26279325ad5c46744efa09d6027 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/documentsetversions/item"
    iec1083d88bf75a5c5d9ec0cd2e11b051c52addbc2efb1c26db7c0e9b60aab3d7 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item/versions/item"
)

// ListItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.list entity.
type ListItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ListItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ListItemItemRequestBuilderGetQueryParameters all items contained in the list.
type ListItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ListItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ListItemItemRequestBuilderGetQueryParameters
}
// ListItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) Activities()(*i5ce7a932ecf70db030b8364f7699aa3822b2ace9797a93c63743a24c7b9c40e2.ActivitiesRequestBuilder) {
    return i5ce7a932ecf70db030b8364f7699aa3822b2ace9797a93c63743a24c7b9c40e2.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) ActivitiesById(id string)(*i201a560e0b52f390181eee26b5bb385daeff5fecb17cf5d3d44e3139031471e5.ItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["itemActivityOLD%2Did"] = id
    }
    return i201a560e0b52f390181eee26b5bb385daeff5fecb17cf5d3d44e3139031471e5.NewItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) Analytics()(*i3d8467aa52be83bf5e27b9b74f21a76f1a5c28b249229ca2c40ba746dc31f6cb.AnalyticsRequestBuilder) {
    return i3d8467aa52be83bf5e27b9b74f21a76f1a5c28b249229ca2c40ba746dc31f6cb.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewListItemItemRequestBuilderInternal instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemItemRequestBuilder) {
    m := &ListItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem%2Did}/list/items/{listItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListItemItemRequestBuilder instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewListItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property items for shares
func (m *ListItemItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ListItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation all items contained in the list.
func (m *ListItemItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ListItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateLink provides operations to call the createLink method.
func (m *ListItemItemRequestBuilder) CreateLink()(*i0df30f15a37bda4c75a746f1d3a78aeb6d7fedcda8fa91a25630e37a46bb63d3.CreateLinkRequestBuilder) {
    return i0df30f15a37bda4c75a746f1d3a78aeb6d7fedcda8fa91a25630e37a46bb63d3.NewCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreatePatchRequestInformation update the navigation property items in shares
func (m *ListItemItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property items for shares
func (m *ListItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ListItemItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DocumentSetVersions provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) DocumentSetVersions()(*id63b80b47ec3f1a4dd0724f6f1236b0d27e2a5216fe80c55230633ece44c6914.DocumentSetVersionsRequestBuilder) {
    return id63b80b47ec3f1a4dd0724f6f1236b0d27e2a5216fe80c55230633ece44c6914.NewDocumentSetVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DocumentSetVersionsById provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) DocumentSetVersionsById(id string)(*idbb42c40e539b6a5f59145da8ae59d604202f26279325ad5c46744efa09d6027.DocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["documentSetVersion%2Did"] = id
    }
    return idbb42c40e539b6a5f59145da8ae59d604202f26279325ad5c46744efa09d6027.NewDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// DriveItem provides operations to manage the driveItem property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) DriveItem()(*i426197eeda24a2ef666da9df0ddff559ab7003700bb784ad8da8009e9521a011.DriveItemRequestBuilder) {
    return i426197eeda24a2ef666da9df0ddff559ab7003700bb784ad8da8009e9521a011.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Fields provides operations to manage the fields property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) Fields()(*i6a33e6735c94a4eec0295b3339373e3d3d94a57d512c4c1297f7ec69a5bcace5.FieldsRequestBuilder) {
    return i6a33e6735c94a4eec0295b3339373e3d3d94a57d512c4c1297f7ec69a5bcace5.NewFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get all items contained in the list.
func (m *ListItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ListItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ListItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*i5245c95090e9c6e9cde87b4772a00a7124828683c592734b734a62638f6f0df8.GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return i5245c95090e9c6e9cde87b4772a00a7124828683c592734b734a62638f6f0df8.NewGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime);
}
// Patch update the navigation property items in shares
func (m *ListItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ListItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable), nil
}
// Versions provides operations to manage the versions property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) Versions()(*i7e01e46513a14884280655da10ff1bf036a7401d1f1dd3ee9dbcbe02ffe6757e.VersionsRequestBuilder) {
    return i7e01e46513a14884280655da10ff1bf036a7401d1f1dd3ee9dbcbe02ffe6757e.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.listItem entity.
func (m *ListItemItemRequestBuilder) VersionsById(id string)(*iec1083d88bf75a5c5d9ec0cd2e11b051c52addbc2efb1c26db7c0e9b60aab3d7.ListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItemVersion%2Did"] = id
    }
    return iec1083d88bf75a5c5d9ec0cd2e11b051c52addbc2efb1c26db7c0e9b60aab3d7.NewListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
