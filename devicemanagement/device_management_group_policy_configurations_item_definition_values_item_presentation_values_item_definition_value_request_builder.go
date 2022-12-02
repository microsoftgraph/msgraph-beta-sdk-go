package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder provides operations to manage the definitionValue property of the microsoft.graph.groupPolicyPresentationValue entity.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderGetQueryParameters the group policy definition value associated with the presentation value.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderGetQueryParameters
}
// NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderInternal instantiates a new DefinitionValueRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder) {
    m := &DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration%2Did}/definitionValues/{groupPolicyDefinitionValue%2Did}/presentationValues/{groupPolicyPresentationValue%2Did}/definitionValue{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder instantiates a new DefinitionValueRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation the group policy definition value associated with the presentation value.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the group policy definition value associated with the presentation value.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable), nil
}
