package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder provides operations to call the importOffice365DeviceConfigurationPolicies method.
type TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal instantiates a new TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    m := &TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/templates/importOffice365DeviceConfigurationPolicies", pathParameters),
    }
    return m
}
// NewTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder instantiates a new TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action importOffice365DeviceConfigurationPolicies
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a TemplatesImportOffice365DeviceConfigurationPoliciesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) Post(ctx context.Context, requestConfiguration *TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(TemplatesImportOffice365DeviceConfigurationPoliciesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTemplatesImportOffice365DeviceConfigurationPoliciesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TemplatesImportOffice365DeviceConfigurationPoliciesResponseable), nil
}
// PostAsImportOffice365DeviceConfigurationPoliciesPostResponse invoke action importOffice365DeviceConfigurationPolicies
// returns a TemplatesImportOffice365DeviceConfigurationPoliciesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) PostAsImportOffice365DeviceConfigurationPoliciesPostResponse(ctx context.Context, requestConfiguration *TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(TemplatesImportOffice365DeviceConfigurationPoliciesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTemplatesImportOffice365DeviceConfigurationPoliciesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TemplatesImportOffice365DeviceConfigurationPoliciesPostResponseable), nil
}
// ToPostRequestInformation invoke action importOffice365DeviceConfigurationPolicies
// returns a *RequestInformation when successful
func (m *TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder when successful
func (m *TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) WithUrl(rawUrl string)(*TemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    return NewTemplatesImportOffice365DeviceConfigurationPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
