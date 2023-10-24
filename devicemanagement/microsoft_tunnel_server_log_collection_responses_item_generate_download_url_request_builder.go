package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder provides operations to call the generateDownloadUrl method.
type MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilderInternal instantiates a new GenerateDownloadUrlRequestBuilder and sets the default values.
func NewMicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder) {
    m := &MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/microsoftTunnelServerLogCollectionResponses/{microsoftTunnelServerLogCollectionResponse%2Did}/generateDownloadUrl", pathParameters),
    }
    return m
}
// NewMicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder instantiates a new GenerateDownloadUrlRequestBuilder and sets the default values.
func NewMicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action generateDownloadUrl
// Deprecated: This method is obsolete. Use PostAsGenerateDownloadUrlPostResponse instead.
func (m *MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder) Post(ctx context.Context, requestConfiguration *MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilderPostRequestConfiguration)(MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlResponseable), nil
}
// PostAsGenerateDownloadUrlPostResponse invoke action generateDownloadUrl
func (m *MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder) PostAsGenerateDownloadUrlPostResponse(ctx context.Context, requestConfiguration *MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilderPostRequestConfiguration)(MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlPostResponseable), nil
}
// ToPostRequestInformation invoke action generateDownloadUrl
func (m *MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder) WithUrl(rawUrl string)(*MicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder) {
    return NewMicrosoftTunnelServerLogCollectionResponsesItemGenerateDownloadUrlRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
