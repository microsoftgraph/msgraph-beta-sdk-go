package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder provides operations to call the extractSensitivityLabels method.
type ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderInternal instantiates a new ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder and sets the default values.
func NewItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) {
    m := &ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/extractSensitivityLabels", pathParameters),
    }
    return m
}
// NewItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder instantiates a new ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder and sets the default values.
func NewItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action extractSensitivityLabels
// returns a ExtractSensitivityLabelsResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExtractSensitivityLabelsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateExtractSensitivityLabelsResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ExtractSensitivityLabelsResultable), nil
}
// ToPostRequestInformation invoke action extractSensitivityLabels
// returns a *RequestInformation when successful
func (m *ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder when successful
func (m *ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder) {
    return NewItemItemsItemExtractsensitivitylabelsExtractSensitivityLabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
