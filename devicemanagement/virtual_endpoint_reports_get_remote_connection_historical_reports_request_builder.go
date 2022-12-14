package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder provides operations to call the getRemoteConnectionHistoricalReports method.
type VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderInternal instantiates a new GetRemoteConnectionHistoricalReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) {
    m := &VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/reports/microsoft.graph.getRemoteConnectionHistoricalReports";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder instantiates a new GetRemoteConnectionHistoricalReportsRequestBuilder and sets the default values.
func NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation get the remote connection history records of a Cloud PC during a given period.
func (m *VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) CreatePostRequestInformation(ctx context.Context, body VirtualEndpointReportsGetRemoteConnectionHistoricalReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post get the remote connection history records of a Cloud PC during a given period.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudpcreports-getremoteconnectionhistoricalreports?view=graph-rest-1.0
func (m *VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilder) Post(ctx context.Context, body VirtualEndpointReportsGetRemoteConnectionHistoricalReportsPostRequestBodyable, requestConfiguration *VirtualEndpointReportsGetRemoteConnectionHistoricalReportsRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
