package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder provides operations to manage the macOSSoftwareUpdateAccountSummaries property of the microsoft.graph.deviceManagement entity.
type MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters the MacOS software update account summaries for this account.
type MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetQueryParameters
}
// MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CategorySummaries provides operations to manage the categorySummaries property of the microsoft.graph.macOSSoftwareUpdateAccountSummary entity.
// returns a *MacossoftwareupdateaccountsummariesItemCategorysummariesCategorySummariesRequestBuilder when successful
func (m *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) CategorySummaries()(*MacossoftwareupdateaccountsummariesItemCategorysummariesCategorySummariesRequestBuilder) {
    return NewMacossoftwareupdateaccountsummariesItemCategorysummariesCategorySummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal instantiates a new MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder and sets the default values.
func NewMacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    m := &MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder instantiates a new MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder and sets the default values.
func NewMacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property macOSSoftwareUpdateAccountSummaries for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the MacOS software update account summaries for this account.
// returns a MacOSSoftwareUpdateAccountSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable), nil
}
// Patch update the navigation property macOSSoftwareUpdateAccountSummaries in deviceManagement
// returns a MacOSSoftwareUpdateAccountSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, requestConfiguration *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateAccountSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property macOSSoftwareUpdateAccountSummaries for deviceManagement
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the MacOS software update account summaries for this account.
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property macOSSoftwareUpdateAccountSummaries in deviceManagement
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateAccountSummaryable, requestConfiguration *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder when successful
func (m *MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) WithUrl(rawUrl string)(*MacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder) {
    return NewMacossoftwareupdateaccountsummariesMacOSSoftwareUpdateAccountSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
