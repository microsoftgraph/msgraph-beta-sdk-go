package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder provides operations to call the updateDefinitionValues method.
type GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilderInternal instantiates a new GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder and sets the default values.
func NewGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder) {
    m := &GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration%2Did}/updateDefinitionValues", pathParameters),
    }
    return m
}
// NewGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder instantiates a new GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder and sets the default values.
func NewGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateDefinitionValues
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder) Post(ctx context.Context, body GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBodyable, requestConfiguration *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action updateDefinitionValues
// returns a *RequestInformation when successful
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder) ToPostRequestInformation(ctx context.Context, body GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesPostRequestBodyable, requestConfiguration *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder when successful
func (m *GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder) WithUrl(rawUrl string)(*GrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder) {
    return NewGrouppolicyconfigurationsItemUpdatedefinitionvaluesUpdateDefinitionValuesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
