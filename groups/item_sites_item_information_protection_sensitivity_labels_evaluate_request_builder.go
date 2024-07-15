package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder provides operations to call the evaluate method.
type ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilderInternal instantiates a new ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder) {
    m := &ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/sensitivityLabels/evaluate", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder instantiates a new ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder and sets the default values.
func NewItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action evaluate
// returns a EvaluateLabelJobResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder) Post(ctx context.Context, body ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EvaluateLabelJobResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateEvaluateLabelJobResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EvaluateLabelJobResponseable), nil
}
// ToPostRequestInformation invoke action evaluate
// returns a *RequestInformation when successful
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBodyable, requestConfiguration *ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder when successful
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder) {
    return NewItemSitesItemInformationProtectionSensitivityLabelsEvaluateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
