package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder provides operations to count the resources in the collection.
type DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilderGetQueryParameters get the number of the resource
type DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilderGetQueryParameters
}
// NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder) {
    m := &DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/allowedMembers/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder) WithUrl(rawUrl string)(*DeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder) {
    return NewDeletedTeamsItemChannelsItemSharedWithTeamsItemAllowedMembersCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
