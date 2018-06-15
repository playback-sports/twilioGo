package twilio

// Grant defines interface of twilio permission model.
// Any structures implement this interface can be bundled into a twilio access token.
type Grant interface {
	// Key is the identifier of a permission.
	Key() string
	// Payload defines how permission data is marshaled in an access token.
	Payload() interface{}
}

// NewVoiceGrant creates a new permission grant for twilio programmable voice SDK
func NewVoiceGrant(ApplicationSid, PushCredentialSid, EndpointID string, AllowIncoming bool, ApplicationParams map[string]interface{}) Grant {
	return &voiceGrant{
		allowIncoming:             AllowIncoming,
		outgoingApplicationSid:    ApplicationSid,
		outgoingApplicationParams: ApplicationParams,
		pushCredentialSid:         PushCredentialSid,
		endpointID:                EndpointID,
	}
}

// twilio voice grant
type voiceGrant struct {
	allowIncoming             bool
	outgoingApplicationSid    string
	outgoingApplicationParams map[string]interface{}
	pushCredentialSid         string
	endpointID                string
}

// Key implements Grant interface.
func (g *voiceGrant) Key() string {
	return "voice"
}

// Payload implements Grant interface.
func (g *voiceGrant) Payload() interface{} {
	return struct {
		Incoming          interface{} `json:"incoming,omitempty"`
		Outgoing          interface{} `json:"outgoing,omitempty"`
		PushCredentialSid string      `json:"push_credential_sid,omitempty"`
		EndpointID        string      `json:"endpoint_id,omitempty"`
	}{
		Incoming: struct {
			Allow bool `json:"allow,omitempty"`
		}{
			Allow: g.allowIncoming,
		},
		Outgoing: struct {
			ApplicationSid    string                 `json:"application_sid,omitempty"`
			ApplicationParams map[string]interface{} `json:"params,omitempty"`
		}{
			ApplicationSid:    g.outgoingApplicationSid,
			ApplicationParams: g.outgoingApplicationParams,
		},
		PushCredentialSid: g.pushCredentialSid,
		EndpointID:        g.endpointID,
	}
}

// NewChatGrant creates a new permission grant for twilio chat service.
func NewChatGrant(serviceSid, endpointID, deploymentRoleSid, pushCredentialSid string) Grant {
	return &chatGrant{
		serviceSid:        serviceSid,
		endpointID:        endpointID,
		deploymentRoleSid: deploymentRoleSid,
		pushCredentialSid: pushCredentialSid,
	}
}

// twilio chat grant
type chatGrant struct {
	serviceSid        string
	endpointID        string
	deploymentRoleSid string
	pushCredentialSid string
}

// Key implements Grant interface.
func (g *chatGrant) Key() string {
	return "chat"
}

// Payload implements Grant interface.
func (g *chatGrant) Payload() interface{} {
	return struct {
		ServiceSid        string `json:"service_sid,omitempty"`
		EndpointID        string `json:"endpoint_id,omitempty"`
		DeploymentRoleSid string `json:"deployment_role_sid,omitempty"`
		PushCredentialSid string `json:"push_credential_sid,omitempty"`
	}{
		ServiceSid:        g.serviceSid,
		EndpointID:        g.endpointID,
		DeploymentRoleSid: g.deploymentRoleSid,
		PushCredentialSid: g.pushCredentialSid,
	}
}

// NewConversationGrant creates a new permission grant for twilio video conversation service.
func NewConversationGrant(sid string) Grant {
	return &conversationGrant{sid: sid}
}

// twilio video conversation grant
type conversationGrant struct {
	sid string
}

// Key implements Grant interface.
func (g *conversationGrant) Key() string {
	return "rtc"
}

// Payload implements Grant interface.
func (g *conversationGrant) Payload() interface{} {
	return struct {
		ConfigurationProfileSid string `json:"configuration_profile_sid,omitempty"`
	}{
		ConfigurationProfileSid: g.sid,
	}
}

// NewVideoGrant creates a new video grant for twilio video conversation service.
func NewVideoGrant(room string) Grant {
	return &videoGrant{room: room}
}

// twilio video grant
type videoGrant struct {
	room string
}

// Key implements Grant interface.
func (g *videoGrant) Key() string {
	return "video"
}

// Payload implements Grant interface.
func (g *videoGrant) Payload() interface{} {
	return struct {
		Room string `json:"room,omitempty"`
	}{
		Room: g.room,
	}
}

// NewIPMessagingGrant creates a new permission for twilio ip messaging service.
func NewIPMessagingGrant(serviceSid, endpointID, deploymentRoleSid, pushCredentialSid string) Grant {
	return &ipMessagingGrant{
		serviceSid:        serviceSid,
		endpointID:        endpointID,
		deploymentRoleSid: deploymentRoleSid,
		pushCredentialSid: pushCredentialSid,
	}
}

// twilio ip messaging grant
type ipMessagingGrant struct {
	serviceSid        string
	endpointID        string
	deploymentRoleSid string
	pushCredentialSid string
}

// Key implements Grant interface.
func (g *ipMessagingGrant) Key() string {
	return "ip_messaging"
}

// Payload implements Grant interface.
func (g *ipMessagingGrant) Payload() interface{} {
	return struct {
		ServiceSid        string `json:"service_sid,omitempty"`
		EndpointID        string `json:"endpoint_id,omitempty"`
		DeploymentRoleSid string `json:"deployment_role_sid,omitempty"`
		PushCredentialSid string `json:"push_credential_sid,omitempty"`
	}{
		ServiceSid:        g.serviceSid,
		EndpointID:        g.endpointID,
		DeploymentRoleSid: g.deploymentRoleSid,
		PushCredentialSid: g.pushCredentialSid,
	}
}
