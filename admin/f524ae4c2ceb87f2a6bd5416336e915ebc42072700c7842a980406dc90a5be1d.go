package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c "github.com/microsoftgraph/msgraph-beta-sdk-go/models/windowsupdates"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder provides operations to manage the applicableContent property of the microsoft.graph.windowsUpdates.deploymentAudience entity.
type WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderGetQueryParameters content eligible to deploy to devices in the audience. Not nullable. Read-only.
type WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderGetQueryParameters
}
// WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CatalogEntry provides operations to manage the catalogEntry property of the microsoft.graph.windowsUpdates.applicableContent entity.
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemCatalogEntryRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) CatalogEntry()(*WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemCatalogEntryRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemCatalogEntryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderInternal instantiates a new WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) {
    m := &WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/windows/updates/updatePolicies/{updatePolicy%2Did}/audience/applicableContent/{applicableContent%2DcatalogEntryId}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder instantiates a new WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder and sets the default values.
func NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property applicableContent for admin
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderGetRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, error) {
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
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) MatchedDevices()(*WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentItemMatchedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property applicableContent in admin
// returns a ApplicableContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) Patch(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderPatchRequestConfiguration)(i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, error) {
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
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body i17376df570f19ff3c32da2d66a677d31250ed0ff64059351645f48a152316b3c.ApplicableContentable, requestConfiguration *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder when successful
func (m *WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) WithUrl(rawUrl string)(*WindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder) {
    return NewWindowsUpdatesUpdatePoliciesItemAudienceApplicableContentApplicableContentCatalogEntryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
