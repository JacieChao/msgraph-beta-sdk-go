package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i9d482566fdf9936391b4484a35db36ee2c9ea057c528e4906b93646a759d755c "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/remoteassistancepartners/item/disconnect"
    id0d53dab897cd68adbdccb455d85697abd98f9e5792c35eb1a92375c1451257b "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/remoteassistancepartners/item/beginonboarding"
)

// RemoteAssistancePartnerItemRequestBuilder builds and executes requests for operations under \deviceManagement\remoteAssistancePartners\{remoteAssistancePartner-id}
type RemoteAssistancePartnerItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RemoteAssistancePartnerItemRequestBuilderDeleteOptions options for Delete
type RemoteAssistancePartnerItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RemoteAssistancePartnerItemRequestBuilderGetOptions options for Get
type RemoteAssistancePartnerItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RemoteAssistancePartnerItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RemoteAssistancePartnerItemRequestBuilderGetQueryParameters the remote assist partners.
type RemoteAssistancePartnerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RemoteAssistancePartnerItemRequestBuilderPatchOptions options for Patch
type RemoteAssistancePartnerItemRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RemoteAssistancePartner;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RemoteAssistancePartnerItemRequestBuilder) BeginOnboarding()(*id0d53dab897cd68adbdccb455d85697abd98f9e5792c35eb1a92375c1451257b.BeginOnboardingRequestBuilder) {
    return id0d53dab897cd68adbdccb455d85697abd98f9e5792c35eb1a92375c1451257b.NewBeginOnboardingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewRemoteAssistancePartnerItemRequestBuilderInternal instantiates a new RemoteAssistancePartnerItemRequestBuilder and sets the default values.
func NewRemoteAssistancePartnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RemoteAssistancePartnerItemRequestBuilder) {
    m := &RemoteAssistancePartnerItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/remoteAssistancePartners/{remoteAssistancePartner_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRemoteAssistancePartnerItemRequestBuilder instantiates a new RemoteAssistancePartnerItemRequestBuilder and sets the default values.
func NewRemoteAssistancePartnerItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RemoteAssistancePartnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRemoteAssistancePartnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the remote assist partners.
func (m *RemoteAssistancePartnerItemRequestBuilder) CreateDeleteRequestInformation(options *RemoteAssistancePartnerItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the remote assist partners.
func (m *RemoteAssistancePartnerItemRequestBuilder) CreateGetRequestInformation(options *RemoteAssistancePartnerItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the remote assist partners.
func (m *RemoteAssistancePartnerItemRequestBuilder) CreatePatchRequestInformation(options *RemoteAssistancePartnerItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the remote assist partners.
func (m *RemoteAssistancePartnerItemRequestBuilder) Delete(options *RemoteAssistancePartnerItemRequestBuilderDeleteOptions)(error) {
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
func (m *RemoteAssistancePartnerItemRequestBuilder) Disconnect()(*i9d482566fdf9936391b4484a35db36ee2c9ea057c528e4906b93646a759d755c.DisconnectRequestBuilder) {
    return i9d482566fdf9936391b4484a35db36ee2c9ea057c528e4906b93646a759d755c.NewDisconnectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the remote assist partners.
func (m *RemoteAssistancePartnerItemRequestBuilder) Get(options *RemoteAssistancePartnerItemRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RemoteAssistancePartner, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewRemoteAssistancePartner() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.RemoteAssistancePartner), nil
}
// Patch the remote assist partners.
func (m *RemoteAssistancePartnerItemRequestBuilder) Patch(options *RemoteAssistancePartnerItemRequestBuilderPatchOptions)(error) {
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