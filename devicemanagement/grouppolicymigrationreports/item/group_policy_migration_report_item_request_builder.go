package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i38dcbc387461ec5f887af4b1883963552ff5d81b6dc6f5aa2953f7f8f864121b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/grouppolicysettingmappings"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    ie4e5feb9f5b1a4b9a016d60fb9ac47087a8d705ab16fc0ae43f56fae55d28e34 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/unsupportedgrouppolicyextensions"
    i2d8bacdd344be782bb5743b2f7edb3872ac2f58bbb54c2a247c2bacadd8de35f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/grouppolicysettingmappings/item"
    i8771848a49e50ae62c54e7781ba80e77c5b2ae75c24c00f4ca58eac952cbd206 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/grouppolicymigrationreports/item/unsupportedgrouppolicyextensions/item"
)

// GroupPolicyMigrationReportItemRequestBuilder provides operations to manage the groupPolicyMigrationReports property of the microsoft.graph.deviceManagement entity.
type GroupPolicyMigrationReportItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GroupPolicyMigrationReportItemRequestBuilderDeleteOptions options for Delete
type GroupPolicyMigrationReportItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyMigrationReportItemRequestBuilderGetOptions options for Get
type GroupPolicyMigrationReportItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *GroupPolicyMigrationReportItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// GroupPolicyMigrationReportItemRequestBuilderGetQueryParameters a list of Group Policy migration reports.
type GroupPolicyMigrationReportItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// GroupPolicyMigrationReportItemRequestBuilderPatchOptions options for Patch
type GroupPolicyMigrationReportItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyMigrationReportable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGroupPolicyMigrationReportItemRequestBuilderInternal instantiates a new GroupPolicyMigrationReportItemRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyMigrationReportItemRequestBuilder) {
    m := &GroupPolicyMigrationReportItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyMigrationReports/{groupPolicyMigrationReport_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupPolicyMigrationReportItemRequestBuilder instantiates a new GroupPolicyMigrationReportItemRequestBuilder and sets the default values.
func NewGroupPolicyMigrationReportItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GroupPolicyMigrationReportItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupPolicyMigrationReportItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groupPolicyMigrationReports for deviceManagement
func (m *GroupPolicyMigrationReportItemRequestBuilder) CreateDeleteRequestInformation(options *GroupPolicyMigrationReportItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation a list of Group Policy migration reports.
func (m *GroupPolicyMigrationReportItemRequestBuilder) CreateGetRequestInformation(options *GroupPolicyMigrationReportItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property groupPolicyMigrationReports in deviceManagement
func (m *GroupPolicyMigrationReportItemRequestBuilder) CreatePatchRequestInformation(options *GroupPolicyMigrationReportItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property groupPolicyMigrationReports for deviceManagement
func (m *GroupPolicyMigrationReportItemRequestBuilder) Delete(options *GroupPolicyMigrationReportItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get a list of Group Policy migration reports.
func (m *GroupPolicyMigrationReportItemRequestBuilder) Get(options *GroupPolicyMigrationReportItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyMigrationReportable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateGroupPolicyMigrationReportFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.GroupPolicyMigrationReportable), nil
}
func (m *GroupPolicyMigrationReportItemRequestBuilder) GroupPolicySettingMappings()(*i38dcbc387461ec5f887af4b1883963552ff5d81b6dc6f5aa2953f7f8f864121b.GroupPolicySettingMappingsRequestBuilder) {
    return i38dcbc387461ec5f887af4b1883963552ff5d81b6dc6f5aa2953f7f8f864121b.NewGroupPolicySettingMappingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GroupPolicySettingMappingsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyMigrationReports.item.groupPolicySettingMappings.item collection
func (m *GroupPolicyMigrationReportItemRequestBuilder) GroupPolicySettingMappingsById(id string)(*i2d8bacdd344be782bb5743b2f7edb3872ac2f58bbb54c2a247c2bacadd8de35f.GroupPolicySettingMappingItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicySettingMapping_id"] = id
    }
    return i2d8bacdd344be782bb5743b2f7edb3872ac2f58bbb54c2a247c2bacadd8de35f.NewGroupPolicySettingMappingItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property groupPolicyMigrationReports in deviceManagement
func (m *GroupPolicyMigrationReportItemRequestBuilder) Patch(options *GroupPolicyMigrationReportItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *GroupPolicyMigrationReportItemRequestBuilder) UnsupportedGroupPolicyExtensions()(*ie4e5feb9f5b1a4b9a016d60fb9ac47087a8d705ab16fc0ae43f56fae55d28e34.UnsupportedGroupPolicyExtensionsRequestBuilder) {
    return ie4e5feb9f5b1a4b9a016d60fb9ac47087a8d705ab16fc0ae43f56fae55d28e34.NewUnsupportedGroupPolicyExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UnsupportedGroupPolicyExtensionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.groupPolicyMigrationReports.item.unsupportedGroupPolicyExtensions.item collection
func (m *GroupPolicyMigrationReportItemRequestBuilder) UnsupportedGroupPolicyExtensionsById(id string)(*i8771848a49e50ae62c54e7781ba80e77c5b2ae75c24c00f4ca58eac952cbd206.UnsupportedGroupPolicyExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["unsupportedGroupPolicyExtension_id"] = id
    }
    return i8771848a49e50ae62c54e7781ba80e77c5b2ae75c24c00f4ca58eac952cbd206.NewUnsupportedGroupPolicyExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
