package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSetting entity.
type CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetQueryParameters list of related Setting Definitions. This property is read-only.
type CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetQueryParameters struct {
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
// CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetQueryParameters
}
// ByDeviceManagementConfigurationSettingDefinitionId provides operations to manage the settingDefinitions property of the microsoft.graph.deviceManagementConfigurationSetting entity.
// returns a *CompliancePoliciesItemSettingsItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder when successful
func (m *CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder) ByDeviceManagementConfigurationSettingDefinitionId(deviceManagementConfigurationSettingDefinitionId string)(*CompliancePoliciesItemSettingsItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementConfigurationSettingDefinitionId != "" {
        urlTplParams["deviceManagementConfigurationSettingDefinition%2Did"] = deviceManagementConfigurationSettingDefinitionId
    }
    return NewCompliancePoliciesItemSettingsItemSettingDefinitionsDeviceManagementConfigurationSettingDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilderInternal instantiates a new CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder and sets the default values.
func NewCompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder) {
    m := &CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/compliancePolicies/{deviceManagementCompliancePolicy%2Did}/settings/{deviceManagementConfigurationSetting%2Did}/settingDefinitions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder instantiates a new CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder and sets the default values.
func NewCompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompliancePoliciesItemSettingsItemSettingDefinitionsCountRequestBuilder when successful
func (m *CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder) Count()(*CompliancePoliciesItemSettingsItemSettingDefinitionsCountRequestBuilder) {
    return NewCompliancePoliciesItemSettingsItemSettingDefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list of related Setting Definitions. This property is read-only.
// returns a DeviceManagementConfigurationSettingDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementConfigurationSettingDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementConfigurationSettingDefinitionCollectionResponseable), nil
}
// ToGetRequestInformation list of related Setting Definitions. This property is read-only.
// returns a *RequestInformation when successful
func (m *CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder when successful
func (m *CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder) WithUrl(rawUrl string)(*CompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder) {
    return NewCompliancePoliciesItemSettingsItemSettingDefinitionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
