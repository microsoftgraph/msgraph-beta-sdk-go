package financials

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilder provides operations to count the resources in the collection.
type CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilderGetQueryParameters get the number of the resource
type CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilderGetQueryParameters
}
// NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilder) {
    m := &CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/financials/companies/{company%2Did}/salesCreditMemos/{salesCreditMemo%2Did}/salesCreditMemoLines/{salesCreditMemoLine%2Did}/item/picture/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesItemSalesCreditMemosItemSalesCreditMemoLinesItemItemPictureCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "text/plain")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
