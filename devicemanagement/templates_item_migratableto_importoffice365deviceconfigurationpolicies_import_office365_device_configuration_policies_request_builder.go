package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder provides operations to call the importOffice365DeviceConfigurationPolicies method.
type TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal instantiates a new TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewTemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    m := &TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/templates/{deviceManagementTemplate%2Did}/migratableTo/importOffice365DeviceConfigurationPolicies", pathParameters),
    }
    return m
}
// NewTemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder instantiates a new TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder and sets the default values.
func NewTemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action importOffice365DeviceConfigurationPolicies
// Deprecated: This method is obsolete. Use PostAsImportOffice365DeviceConfigurationPoliciesPostResponse instead.
// returns a TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder) Post(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesResponseable), nil
}
// PostAsImportOffice365DeviceConfigurationPoliciesPostResponse invoke action importOffice365DeviceConfigurationPolicies
// returns a TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder) PostAsImportOffice365DeviceConfigurationPoliciesPostResponse(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesPostResponseable), nil
}
// ToPostRequestInformation invoke action importOffice365DeviceConfigurationPolicies
// returns a *RequestInformation when successful
func (m *TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder when successful
func (m *TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder) WithUrl(rawUrl string)(*TemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder) {
    return NewTemplatesItemMigratabletoImportoffice365deviceconfigurationpoliciesImportOffice365DeviceConfigurationPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
