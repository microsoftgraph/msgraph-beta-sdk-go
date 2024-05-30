package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder provides operations to manage the updateStateSummaries property of the microsoft.graph.macOSSoftwareUpdateCategorySummary entity.
type MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetQueryParameters summary of the update states.
type MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetQueryParameters
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderInternal instantiates a new MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder and sets the default values.
func NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) {
    m := &MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary%2Did}/categorySummaries/{macOSSoftwareUpdateCategorySummary%2Did}/updateStateSummaries/{macOSSoftwareUpdateStateSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder instantiates a new MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder and sets the default values.
func NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property updateStateSummaries for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get summary of the update states.
// returns a MacOSSoftwareUpdateStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable), nil
}
// Patch update the navigation property updateStateSummaries in deviceManagement
// returns a MacOSSoftwareUpdateStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property updateStateSummaries for deviceManagement
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation summary of the update states.
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property updateStateSummaries in deviceManagement
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) WithUrl(rawUrl string)(*MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) {
    return NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
