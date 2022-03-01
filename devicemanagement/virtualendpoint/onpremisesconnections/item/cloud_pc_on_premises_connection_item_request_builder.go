package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i747483174c0e61425ce1a618854ba8ca6d202a33166c1334d75d025ab520f858 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/onpremisesconnections/item/runhealthchecks"
    if44fdd10cc1f6f4d9584fbcc1424ba5392d6f3eda6f11879e38028e5a1ea53bf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/onpremisesconnections/item/updateaddomainpassword"
)

// CloudPcOnPremisesConnectionItemRequestBuilder builds and executes requests for operations under \deviceManagement\virtualEndpoint\onPremisesConnections\{cloudPcOnPremisesConnection-id}
type CloudPcOnPremisesConnectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CloudPcOnPremisesConnectionItemRequestBuilderDeleteOptions options for Delete
type CloudPcOnPremisesConnectionItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPcOnPremisesConnectionItemRequestBuilderGetOptions options for Get
type CloudPcOnPremisesConnectionItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
type CloudPcOnPremisesConnectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CloudPcOnPremisesConnectionItemRequestBuilderPatchOptions options for Patch
type CloudPcOnPremisesConnectionItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcOnPremisesConnection;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewCloudPcOnPremisesConnectionItemRequestBuilderInternal instantiates a new CloudPcOnPremisesConnectionItemRequestBuilder and sets the default values.
func NewCloudPcOnPremisesConnectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPcOnPremisesConnectionItemRequestBuilder) {
    m := &CloudPcOnPremisesConnectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/onPremisesConnections/{cloudPcOnPremisesConnection_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCloudPcOnPremisesConnectionItemRequestBuilder instantiates a new CloudPcOnPremisesConnectionItemRequestBuilder and sets the default values.
func NewCloudPcOnPremisesConnectionItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CloudPcOnPremisesConnectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPcOnPremisesConnectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) CreateDeleteRequestInformation(options *CloudPcOnPremisesConnectionItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) CreateGetRequestInformation(options *CloudPcOnPremisesConnectionItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) CreatePatchRequestInformation(options *CloudPcOnPremisesConnectionItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) Delete(options *CloudPcOnPremisesConnectionItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) Get(options *CloudPcOnPremisesConnectionItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcOnPremisesConnection, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCloudPcOnPremisesConnection() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CloudPcOnPremisesConnection), nil
}
// Patch a defined collection of Azure resource information that can be used to establish on-premises network connectivity for Cloud PCs.
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) Patch(options *CloudPcOnPremisesConnectionItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) RunHealthChecks()(*i747483174c0e61425ce1a618854ba8ca6d202a33166c1334d75d025ab520f858.RunHealthChecksRequestBuilder) {
    return i747483174c0e61425ce1a618854ba8ca6d202a33166c1334d75d025ab520f858.NewRunHealthChecksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CloudPcOnPremisesConnectionItemRequestBuilder) UpdateAdDomainPassword()(*if44fdd10cc1f6f4d9584fbcc1424ba5392d6f3eda6f11879e38028e5a1ea53bf.UpdateAdDomainPasswordRequestBuilder) {
    return if44fdd10cc1f6f4d9584fbcc1424ba5392d6f3eda6f11879e38028e5a1ea53bf.NewUpdateAdDomainPasswordRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}