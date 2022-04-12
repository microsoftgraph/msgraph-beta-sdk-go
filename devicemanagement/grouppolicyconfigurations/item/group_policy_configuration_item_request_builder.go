package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i16db673b6039e5f54f2159476d6177eaa1e9bfedc38455788aab318c22dffc2e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/updatedefinitionvalues"
    i2573cd7f18c0bef6763b0cd5d5daf2b302474f359820e0e9a5b01b4456663a1b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues"
    ib82d0e0e2e13d1b253232a5b5c89e5d5704adda9d4f7e9e1e53d75f5f967101d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/assignments"
    if10892f25c880cff6f351597e09e418fb5b4358d3bea03afb26c9d1e6f52d6e1 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/assign"
    i0cb6dc720cdfefbc227fbe58a3f3d6fae0bd97db3acd773de3b5dc3e12dd5792 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/assignments/item"
    id5d6fab6089b28b24032ceb492d4726a5f5b3dd8719e632a9f878e2fd9a9c9e8 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicyconfigurations/item/definitionvalues/item"
)

// GroupPolicyConfigurationItemRequestBuilder provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
type GroupPolicyConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupPolicyConfigurationItemRequestBuilderDeleteOptions options for Delete
type GroupPolicyConfigurationItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// GroupPolicyConfigurationItemRequestBuilderGetOptions options for Get
type GroupPolicyConfigurationItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupPolicyConfigurationItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// GroupPolicyConfigurationItemRequestBuilderGetQueryParameters the group policy configurations created by this account.
type GroupPolicyConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupPolicyConfigurationItemRequestBuilderPatchOptions options for Patch
type GroupPolicyConfigurationItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyConfigurationable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Assign the assign property
func (m *GroupPolicyConfigurationItemRequestBuilder) Assign()(*if10892f25c880cff6f351597e09e418fb5b4358d3bea03afb26c9d1e6f52d6e1.AssignRequestBuilder) {
    return if10892f25c880cff6f351597e09e418fb5b4358d3bea03afb26c9d1e6f52d6e1.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *GroupPolicyConfigurationItemRequestBuilder) Assignments()(*ib82d0e0e2e13d1b253232a5b5c89e5d5704adda9d4f7e9e1e53d75f5f967101d.AssignmentsRequestBuilder) {
    return ib82d0e0e2e13d1b253232a5b5c89e5d5704adda9d4f7e9e1e53d75f5f967101d.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyConfigurations.item.assignments.item collection
func (m *GroupPolicyConfigurationItemRequestBuilder) AssignmentsById(id string)(*i0cb6dc720cdfefbc227fbe58a3f3d6fae0bd97db3acd773de3b5dc3e12dd5792.GroupPolicyConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyConfigurationAssignment%2Did"] = id
    }
    return i0cb6dc720cdfefbc227fbe58a3f3d6fae0bd97db3acd773de3b5dc3e12dd5792.NewGroupPolicyConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewGroupPolicyConfigurationItemRequestBuilderInternal instantiates a new GroupPolicyConfigurationItemRequestBuilder and sets the default values.
func NewGroupPolicyConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyConfigurationItemRequestBuilder) {
    m := &GroupPolicyConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyConfigurationItemRequestBuilder instantiates a new GroupPolicyConfigurationItemRequestBuilder and sets the default values.
func NewGroupPolicyConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupPolicyConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groupPolicyConfigurations for deviceManagement
func (m *GroupPolicyConfigurationItemRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyConfigurationItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the group policy configurations created by this account.
func (m *GroupPolicyConfigurationItemRequestBuilder) CreateGetRequestInformation(options *GroupPolicyConfigurationItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property groupPolicyConfigurations in deviceManagement
func (m *GroupPolicyConfigurationItemRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyConfigurationItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// DefinitionValues the definitionValues property
func (m *GroupPolicyConfigurationItemRequestBuilder) DefinitionValues()(*i2573cd7f18c0bef6763b0cd5d5daf2b302474f359820e0e9a5b01b4456663a1b.DefinitionValuesRequestBuilder) {
    return i2573cd7f18c0bef6763b0cd5d5daf2b302474f359820e0e9a5b01b4456663a1b.NewDefinitionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefinitionValuesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyConfigurations.item.definitionValues.item collection
func (m *GroupPolicyConfigurationItemRequestBuilder) DefinitionValuesById(id string)(*id5d6fab6089b28b24032ceb492d4726a5f5b3dd8719e632a9f878e2fd9a9c9e8.GroupPolicyDefinitionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyDefinitionValue%2Did"] = id
    }
    return id5d6fab6089b28b24032ceb492d4726a5f5b3dd8719e632a9f878e2fd9a9c9e8.NewGroupPolicyDefinitionValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property groupPolicyConfigurations for deviceManagement
func (m *GroupPolicyConfigurationItemRequestBuilder) Delete(options *GroupPolicyConfigurationItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the group policy configurations created by this account.
func (m *GroupPolicyConfigurationItemRequestBuilder) Get(options *GroupPolicyConfigurationItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyConfigurationFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyConfigurationable), nil
}
// Patch update the navigation property groupPolicyConfigurations in deviceManagement
func (m *GroupPolicyConfigurationItemRequestBuilder) Patch(options *GroupPolicyConfigurationItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// UpdateDefinitionValues the updateDefinitionValues property
func (m *GroupPolicyConfigurationItemRequestBuilder) UpdateDefinitionValues()(*i16db673b6039e5f54f2159476d6177eaa1e9bfedc38455788aab318c22dffc2e.UpdateDefinitionValuesRequestBuilder) {
    return i16db673b6039e5f54f2159476d6177eaa1e9bfedc38455788aab318c22dffc2e.NewUpdateDefinitionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
