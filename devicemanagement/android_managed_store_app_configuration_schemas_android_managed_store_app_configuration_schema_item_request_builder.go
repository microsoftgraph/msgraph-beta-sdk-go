package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder provides operations to manage the androidManagedStoreAppConfigurationSchemas property of the microsoft.graph.deviceManagement entity.
type AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderGetQueryParameters android Enterprise app configuration schema entities.
type AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderGetQueryParameters
}
// AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderInternal instantiates a new AndroidManagedStoreAppConfigurationSchemaItemRequestBuilder and sets the default values.
func NewAndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) {
    m := &AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/androidManagedStoreAppConfigurationSchemas/{androidManagedStoreAppConfigurationSchema%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder instantiates a new AndroidManagedStoreAppConfigurationSchemaItemRequestBuilder and sets the default values.
func NewAndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property androidManagedStoreAppConfigurationSchemas for deviceManagement
func (m *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get android Enterprise app configuration schema entities.
func (m *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAppConfigurationSchemaFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable), nil
}
// Patch update the navigation property androidManagedStoreAppConfigurationSchemas in deviceManagement
func (m *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable, requestConfiguration *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAndroidManagedStoreAppConfigurationSchemaFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable), nil
}
// ToDeleteRequestInformation delete navigation property androidManagedStoreAppConfigurationSchemas for deviceManagement
func (m *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation android Enterprise app configuration schema entities.
func (m *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property androidManagedStoreAppConfigurationSchemas in deviceManagement
func (m *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AndroidManagedStoreAppConfigurationSchemaable, requestConfiguration *AndroidManagedStoreAppConfigurationSchemasAndroidManagedStoreAppConfigurationSchemaItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
