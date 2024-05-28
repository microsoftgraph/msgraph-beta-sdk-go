package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder provides operations to call the evaluate method.
type ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilderInternal instantiates a new ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder) {
    m := &ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/informationProtection/sensitivityLabels/{sensitivityLabel%2Did}/sublabels/evaluate", pathParameters),
    }
    return m
}
// NewItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder instantiates a new ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder and sets the default values.
func NewItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action evaluate
// returns a EvaluateLabelJobResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder) Post(ctx context.Context, body ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluatePostRequestBodyable, requestConfiguration *ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.EvaluateLabelJobResponseable, error) {
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
func (m *ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluatePostRequestBodyable, requestConfiguration *ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder when successful
func (m *ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder) {
    return NewItemSitesItemInformationprotectionSensitivitylabelsItemSublabelsEvaluateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
