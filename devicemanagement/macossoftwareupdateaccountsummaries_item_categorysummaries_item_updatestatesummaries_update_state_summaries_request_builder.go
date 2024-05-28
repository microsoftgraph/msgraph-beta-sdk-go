package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder provides operations to manage the updateStateSummaries property of the microsoft.graph.macOSSoftwareUpdateCategorySummary entity.
type MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderGetQueryParameters summary of the update states.
type MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderGetQueryParameters
}
// MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByMacOSSoftwareUpdateStateSummaryId provides operations to manage the updateStateSummaries property of the microsoft.graph.macOSSoftwareUpdateCategorySummary entity.
// returns a *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) ByMacOSSoftwareUpdateStateSummaryId(macOSSoftwareUpdateStateSummaryId string)(*MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if macOSSoftwareUpdateStateSummaryId != "" {
        urlTplParams["macOSSoftwareUpdateStateSummary%2Did"] = macOSSoftwareUpdateStateSummaryId
    }
    return NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesMacOSSoftwareUpdateStateSummaryItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderInternal instantiates a new MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder and sets the default values.
func NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) {
    m := &MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/macOSSoftwareUpdateAccountSummaries/{macOSSoftwareUpdateAccountSummary%2Did}/categorySummaries/{macOSSoftwareUpdateCategorySummary%2Did}/updateStateSummaries{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder instantiates a new MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder and sets the default values.
func NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesCountRequestBuilder when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) Count()(*MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesCountRequestBuilder) {
    return NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get summary of the update states.
// returns a MacOSSoftwareUpdateStateSummaryCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) Get(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateMacOSSoftwareUpdateStateSummaryCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryCollectionResponseable), nil
}
// Post create new navigation property to updateStateSummaries for deviceManagement
// returns a MacOSSoftwareUpdateStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation summary of the update states.
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to updateStateSummaries for deviceManagement
// returns a *RequestInformation when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.MacOSSoftwareUpdateStateSummaryable, requestConfiguration *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder when successful
func (m *MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) WithUrl(rawUrl string)(*MacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder) {
    return NewMacossoftwareupdateaccountsummariesItemCategorysummariesItemUpdatestatesummariesUpdateStateSummariesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
