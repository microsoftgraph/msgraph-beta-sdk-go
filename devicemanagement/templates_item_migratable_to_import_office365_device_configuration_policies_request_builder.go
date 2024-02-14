package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder provides operations to call the importOffice365DeviceConfigurationPolicies method.
type TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal instantiates a new TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewTemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    m := &TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/importOffice365DeviceConfigurationPolicies", pathParameters),
    }
    return m
}
// NewTemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder instantiates a new TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewTemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action importOffice365DeviceConfigurationPolicies
// Deprecated: This method is obsolete. Use {TypeName} instead.
// returns a TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder) Post(ctx context.Context, requestConfiguration *TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesResponseable), nil
}
// PostAsImportOffice365DeviceConfigurationPoliciesPostResponse invoke action importOffice365DeviceConfigurationPolicies
// returns a TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder) PostAsImportOffice365DeviceConfigurationPoliciesPostResponse(ctx context.Context, requestConfiguration *TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesPostResponseable), nil
}
// ToPostRequestInformation invoke action importOffice365DeviceConfigurationPolicies
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder when successful
func (m *TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder) WithUrl(rawUrl string)(*TemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    return NewTemplatesItemMigratableToImportOffice365DeviceConfigurationPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
