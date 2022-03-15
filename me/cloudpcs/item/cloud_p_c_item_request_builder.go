package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i0543ab08206c491078113b4b6ec824e6ea4dc3659ea1392f05efc163b50bac03 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/changeuseraccounttype"
    i3fc1b0e9b520291a737c5354479f5ce10f5650dd33030dfefdc4b6ee1d1af0a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/reboot"
    i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph/odataerrors"
    i48833d47a3b49cc67b08b792c187d1d2e754160e7d01e8291dfc3500786630e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/troubleshoot"
    ibd3ed8e86fd7912f38df4a5ee348162bbd2f3faee2c4688adbdf41aed4563b2d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/getcloudpclaunchinfo"
    idd692d448454691ef652a1555dc60f805306ee63076a45f51d1f0e45253c9964 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/rename"
    ie1b3b83917cdcc75511ad5eb6d6dbf8d5c29ca730fb1a3e4c356599c7261e842 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/endgraceperiod"
    ie4696cde156288c3f3bd9ecdc93adfc93dc5d8beeb9cceac7e81177b09b725a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/reprovision"
)

// CloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
type CloudPCItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CloudPCItemRequestBuilderDeleteOptions options for Delete
type CloudPCItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPCItemRequestBuilderGetOptions options for Get
type CloudPCItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CloudPCItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPCItemRequestBuilderGetQueryParameters get cloudPCs from me
type CloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CloudPCItemRequestBuilderPatchOptions options for Patch
type CloudPCItemRequestBuilderPatchOptions struct {
    // 
    Body i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPCable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CloudPCItemRequestBuilder) ChangeUserAccountType()(*i0543ab08206c491078113b4b6ec824e6ea4dc3659ea1392f05efc163b50bac03.ChangeUserAccountTypeRequestBuilder) {
    return i0543ab08206c491078113b4b6ec824e6ea4dc3659ea1392f05efc163b50bac03.NewChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPCItemRequestBuilder) {
    m := &CloudPCItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/cloudPCs/{cloudPC_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cloudPCs for me
func (m *CloudPCItemRequestBuilder) CreateDeleteRequestInformation(options *CloudPCItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get cloudPCs from me
func (m *CloudPCItemRequestBuilder) CreateGetRequestInformation(options *CloudPCItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property cloudPCs in me
func (m *CloudPCItemRequestBuilder) CreatePatchRequestInformation(options *CloudPCItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property cloudPCs for me
func (m *CloudPCItemRequestBuilder) Delete(options *CloudPCItemRequestBuilderDeleteOptions)(error) {
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
func (m *CloudPCItemRequestBuilder) EndGracePeriod()(*ie1b3b83917cdcc75511ad5eb6d6dbf8d5c29ca730fb1a3e4c356599c7261e842.EndGracePeriodRequestBuilder) {
    return ie1b3b83917cdcc75511ad5eb6d6dbf8d5c29ca730fb1a3e4c356599c7261e842.NewEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get cloudPCs from me
func (m *CloudPCItemRequestBuilder) Get(options *CloudPCItemRequestBuilderGetOptions)(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPCable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
        "5XX": i428a28d14ab585560ab266716b214a45f45f18468b52fdb0f932c81a7f9706e4.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CreateCloudPCFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPCable), nil
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *CloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*ibd3ed8e86fd7912f38df4a5ee348162bbd2f3faee2c4688adbdf41aed4563b2d.GetCloudPcLaunchInfoRequestBuilder) {
    return ibd3ed8e86fd7912f38df4a5ee348162bbd2f3faee2c4688adbdf41aed4563b2d.NewGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property cloudPCs in me
func (m *CloudPCItemRequestBuilder) Patch(options *CloudPCItemRequestBuilderPatchOptions)(error) {
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
func (m *CloudPCItemRequestBuilder) Reboot()(*i3fc1b0e9b520291a737c5354479f5ce10f5650dd33030dfefdc4b6ee1d1af0a6.RebootRequestBuilder) {
    return i3fc1b0e9b520291a737c5354479f5ce10f5650dd33030dfefdc4b6ee1d1af0a6.NewRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCItemRequestBuilder) Rename()(*idd692d448454691ef652a1555dc60f805306ee63076a45f51d1f0e45253c9964.RenameRequestBuilder) {
    return idd692d448454691ef652a1555dc60f805306ee63076a45f51d1f0e45253c9964.NewRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCItemRequestBuilder) Reprovision()(*ie4696cde156288c3f3bd9ecdc93adfc93dc5d8beeb9cceac7e81177b09b725a3.ReprovisionRequestBuilder) {
    return ie4696cde156288c3f3bd9ecdc93adfc93dc5d8beeb9cceac7e81177b09b725a3.NewReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPCItemRequestBuilder) Troubleshoot()(*i48833d47a3b49cc67b08b792c187d1d2e754160e7d01e8291dfc3500786630e6.TroubleshootRequestBuilder) {
    return i48833d47a3b49cc67b08b792c187d1d2e754160e7d01e8291dfc3500786630e6.NewTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
