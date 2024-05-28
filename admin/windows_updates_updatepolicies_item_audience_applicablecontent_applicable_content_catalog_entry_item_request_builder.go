package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder provides operations to manage the applicableContent property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderGetQueryParameters content eligible to deploy to devices in the audience. Not nullable. Read-only.
type WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderGetQueryParameters
}
// WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CatalogEntry provides operations to manage the catalogEntry property of the microsoft.graph.windowsUpdates.applicableContent entity.
// returns a *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentItemCatalogentryCatalogEntryRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) CatalogEntry()(*WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentItemCatalogentryCatalogEntryRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentItemCatalogentryCatalogEntryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderInternal instantiates a new WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) {
    m := &WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/applicableContent/{applicableContent%2DcatalogEntryId}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder instantiates a new WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property applicableContent for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get content eligible to deploy to devices in the audience. Not nullable. Read-only.
// returns a ApplicableContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateApplicableContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable), nil
}
// MatchedDevices provides operations to manage the matchedDevices property of the microsoft.graph.windowsUpdates.applicableContent entity.
// returns a *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) MatchedDevices()(*WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentItemMatcheddevicesMatchedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property applicableContent in admin
// returns a ApplicableContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.CreateApplicableContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable), nil
}
// ToDeleteRequestInformation delete navigation property applicableContent for admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation content eligible to deploy to devices in the audience. Not nullable. Read-only.
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property applicableContent in admin
// returns a *RequestInformation when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, requestConfiguration *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder when successful
func (m *WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder) {
    return NewWindowsUpdatesUpdatepoliciesItemAudienceApplicablecontentApplicableContentCatalogEntryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
