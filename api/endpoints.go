package api

import (
	"net/url"

	"github.com/chanbakjsd/gotrix/event"
	"github.com/chanbakjsd/gotrix/matrix"
)

// Version is the Client-Server API version implemented by Gotrix.
var Version = "v3"

// Known Matrix Client-Server API endpoints.
//nolint:lll
var (
	EndpointBase = "_matrix/client/" + Version

	EndpointSupportedVersions = "_matrix/client/versions"

	EndpointUser     = func(id matrix.UserID) string { return EndpointBase + "/user/" + url.PathEscape(string(id)) }
	EndpointRoom     = func(id matrix.RoomID) string { return EndpointBase + "/rooms/" + url.PathEscape(string(id)) }
	EndpointUserRoom = func(userID matrix.UserID, roomID matrix.RoomID) string {
		return EndpointUser(userID) + "/rooms/" + url.PathEscape(string(roomID))
	}

	EndpointLogin     = EndpointBase + "/login"
	EndpointLogout    = EndpointBase + "/logout"
	EndpointLogoutAll = EndpointLogout + "/all"

	EndpointRegister             = EndpointBase + "/register"
	EndpointRegisterAvailable    = EndpointRegister + "/available"
	EndpointRegisterRequestToken = func(authType string) string { return EndpointRegister + "/" + authType + "/requestToken" }

	EndpointAccount                     = EndpointBase + "/account"
	EndpointAccountWhoami               = EndpointAccount + "/whoami"
	EndpointAccountDeactivate           = EndpointAccount + "/deactivate"
	EndpointAccountPassword             = EndpointAccount + "/password"
	EndpointAccountPasswordRequestToken = func(authType string) string { return EndpointAccountPassword + "/" + authType + "/requestToken" }

	EndpointAccount3PID             = EndpointAccount + "/3pid"
	EndpointAccount3PIDAdd          = EndpointAccount3PID + "/add"
	EndpointAccount3PIDBind         = EndpointAccount3PID + "/bind"
	EndpointAccount3PIDRequestToken = func(authType string) string { return EndpointAccount3PID + "/" + authType + "/requestToken" }

	EndpointCapabilities = EndpointBase + "/capabilities"
	EndpointJoinedRooms  = EndpointBase + "/joined_rooms"
	EndpointPublicRooms  = EndpointBase + "/publicRooms"
	EndpointSync         = EndpointBase + "/sync"

	EndpointFilter = func(userID matrix.UserID) string {
		return EndpointUser(userID) + "/filter"
	}
	EndpointFilterGet = func(userID matrix.UserID, filterID string) string {
		return EndpointFilter(userID) + "/" + url.PathEscape(filterID)
	}

	EndpointRoomCreate        = EndpointBase + "/createRoom"
	EndpointRoomAliases       = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/aliases" }
	EndpointRoomBan           = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/ban" }
	EndpointRoomForget        = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/forget" }
	EndpointRoomInvite        = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/invite" }
	EndpointRoomJoin          = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/join" }
	EndpointRoomJoinedMembers = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/joined_members" }
	EndpointRoomKick          = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/kick" }
	EndpointRoomLeave         = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/leave" }
	EndpointRoomMembers       = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/members" }
	EndpointRoomMessages      = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/messages" }
	EndpointRoomState         = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/state" }
	EndpointRoomUnban         = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/unban" }
	EndpointRoomUpgrade       = func(roomID matrix.RoomID) string { return EndpointRoom(roomID) + "/upgrade" }
	EndpointRoomEvent         = func(roomID matrix.RoomID, eventID matrix.EventID) string {
		return EndpointRoom(roomID) + "/event/" + url.PathEscape(string(eventID))
	}
	EndpointRoomStateExact = func(roomID matrix.RoomID, eventType event.Type, stateKey string) string {
		return EndpointRoomState(roomID) + "/" + url.PathEscape(string(eventType)) + "/" + url.PathEscape(stateKey)
	}
	EndpointRoomReceipt = func(roomID matrix.RoomID, receiptType ReceiptType, eventID matrix.EventID) string {
		return EndpointRoomState(roomID) + "/" + url.PathEscape(string(receiptType)) + "/" + url.PathEscape(string(eventID))
	}
	EndpointRoomRedact = func(roomID matrix.RoomID, eventID matrix.EventID, transactionID string) string {
		return EndpointRoom(roomID) + "/redact/" + url.PathEscape(string(eventID)) + "/" + url.PathEscape(transactionID)
	}
	EndpointRoomSend = func(roomID matrix.RoomID, eventType event.Type, transactionID string) string {
		return EndpointRoom(roomID) + "/send/" + url.PathEscape(string(eventType)) + "/" + url.PathEscape(transactionID)
	}
	EndpointRoomTyping = func(roomID matrix.RoomID, userID matrix.UserID) string {
		return EndpointRoom(roomID) + "/typing/" + url.PathEscape(string(userID))
	}

	EndpointDirectory          = EndpointBase + "/directory"
	EndpointDirectoryRoomAlias = func(roomAlias string) string { return EndpointDirectory + "/room/" + url.PathEscape(roomAlias) }
	EndpointDirectoryListRoom  = func(roomID matrix.RoomID) string {
		return EndpointDirectory + "/list/room/" + url.PathEscape(string(roomID))
	}

	EndpointUserDirectorySearch = EndpointBase + "/user_directory/search"

	EndpointProfile            = func(userID matrix.UserID) string { return EndpointBase + "/profile/" + url.PathEscape(string(userID)) }
	EndpointProfileAvatarURL   = func(userID matrix.UserID) string { return EndpointProfile(userID) + "/avatar_url" }
	EndpointProfileDisplayName = func(userID matrix.UserID) string { return EndpointProfile(userID) + "/displayname" }

	EndpointMedia           = "_matrix/media/" + Version
	EndpointMediaConfig     = EndpointMedia + "/config"
	EndpointMediaPreviewURL = EndpointMedia + "/preview_url"
	EndpointMediaUpload     = EndpointMedia + "/upload"
	EndpointMediaDownload   = func(serverName string, mediaID string, fileName string) string {
		return EndpointMedia + "/download/" + url.PathEscape(serverName) + "/" + url.PathEscape(mediaID) + "/" +
			url.PathEscape(fileName)
	}
	EndpointMediaThumbnail = func(serverName string, mediaID string) string {
		return EndpointMedia + "/thumbnail/" + url.PathEscape(serverName) + "/" + url.PathEscape(mediaID)
	}

	EndpointVOIPTURNServers = EndpointMedia + "/voip/turnServer"

	EndpointPresenceStatus = func(userID matrix.UserID) string {
		return EndpointBase + "/presence/" + url.PathEscape(string(userID)) + "/status"
	}

	EndpointAccountDataGlobal = func(userID matrix.UserID, dataType string) string {
		return EndpointUser(userID) + "/account_data/" + url.PathEscape(dataType)
	}
	EndpointAccountDataRoom = func(userID matrix.UserID, roomID matrix.RoomID, dataType string) string {
		return EndpointUserRoom(userID, roomID) + "/account_data/" + url.PathEscape(dataType)
	}

	EndpointSendToDevice = func(eventType event.Type, transactionID string) string {
		return EndpointBase + "/sendToDevice/" + url.PathEscape(string(eventType)) + "/" + url.PathEscape(transactionID)
	}

	EndpointTags = func(userID matrix.UserID, roomID matrix.RoomID) string {
		return EndpointUserRoom(userID, roomID) + "/tags"
	}

	EndpointTag = func(userID matrix.UserID, roomID matrix.RoomID, name matrix.TagName) string {
		return EndpointTags(userID, roomID) + "/" + url.PathEscape(string(name))
	}

	EndpointSSOLogin = func(redirectURL string) string {
		return EndpointBase + "/login/sso/redirect?redirectUrl=" + url.QueryEscape(redirectURL)
	}
)
