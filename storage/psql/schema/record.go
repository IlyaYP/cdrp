package schema

import (
	"time"

	"github.com/IlyaYP/cdrp/model"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Record struct {
	bun.BaseModel                           `bun:"table:records,alias:r"`
	CdrRecordType                           int       `bun:"type:integer"`
	GlobalCallID_callManagerId              int       `bun:"type:integer"`
	GlobalCallID_callId                     int       `bun:"type:integer"`
	OrigLegCallIdentifier                   int       `bun:"type:integer"`
	DateTimeOrigination                     time.Time //`bun:"type:integer"`
	OrigNodeId                              int       `bun:"type:integer"`
	OrigSpan                                int       `bun:"type:integer"`
	OrigIpAddr                              int       `bun:"type:integer"`
	CallingPartyNumber                      string    `bun:"type:varchar(50)"`
	CallingPartyUnicodeLoginUserID          string    `bun:"type:varchar(128)"`
	OrigCause_location                      int       `bun:"type:integer"`
	OrigCause_value                         int       `bun:"type:integer"`
	OrigPrecedenceLevel                     int       `bun:"type:integer"`
	OrigMediaTransportAddress_IP            int       `bun:"type:integer"`
	OrigMediaTransportAddress_Port          int       `bun:"type:integer"`
	OrigMediaCap_payloadCapability          int       `bun:"type:integer"`
	OrigMediaCap_maxFramesPerPacket         int       `bun:"type:integer"`
	OrigMediaCap_g723BitRate                int       `bun:"type:integer"`
	OrigVideoCap_Codec                      int       `bun:"type:integer"`
	OrigVideoCap_Bandwidth                  int       `bun:"type:integer"`
	OrigVideoCap_Resolution                 int       `bun:"type:integer"`
	OrigVideoTransportAddress_IP            int       `bun:"type:integer"`
	OrigVideoTransportAddress_Port          int       `bun:"type:integer"`
	OrigRSVPAudioStat                       string    `bun:"type:varchar(64)"`
	OrigRSVPVideoStat                       string    `bun:"type:varchar(64)"`
	DestLegIdentifier                       int       `bun:"type:integer"`
	DestNodeId                              int       `bun:"type:integer"`
	DestSpan                                int       `bun:"type:integer"`
	DestIpAddr                              int       `bun:"type:integer"`
	OriginalCalledPartyNumber               string    `bun:"type:varchar(50)"`
	FinalCalledPartyNumber                  string    `bun:"type:varchar(50)"`
	FinalCalledPartyUnicodeLoginUserID      string    `bun:"type:varchar(128)"`
	DestCause_location                      int       `bun:"type:integer"`
	DestCause_value                         int       `bun:"type:integer"`
	DestPrecedenceLevel                     int       `bun:"type:integer"`
	DestMediaTransportAddress_IP            int       `bun:"type:integer"`
	DestMediaTransportAddress_Port          int       `bun:"type:integer"`
	DestMediaCap_payloadCapability          int       `bun:"type:integer"`
	DestMediaCap_maxFramesPerPacket         int       `bun:"type:integer"`
	DestMediaCap_g723BitRate                int       `bun:"type:integer"`
	DestVideoCap_Codec                      int       `bun:"type:integer"`
	DestVideoCap_Bandwidth                  int       `bun:"type:integer"`
	DestVideoCap_Resolution                 int       `bun:"type:integer"`
	DestVideoTransportAddress_IP            int       `bun:"type:integer"`
	DestVideoTransportAddress_Port          int       `bun:"type:integer"`
	DestRSVPAudioStat                       string    `bun:"type:varchar(64)"`
	DestRSVPVideoStat                       string    `bun:"type:varchar(64)"`
	DateTimeConnect                         time.Time //`bun:"type:integer"`
	DateTimeDisconnect                      time.Time //`bun:"type:integer"`
	LastRedirectDn                          string    `bun:"type:varchar(50)"`
	Pkid                                    uuid.UUID `bun:"type:uuid,pk"` //UNIQUEIDENTIFIER
	OriginalCalledPartyNumberPartition      string    `bun:"type:varchar(50)"`
	CallingPartyNumberPartition             string    `bun:"type:varchar(50)"`
	FinalCalledPartyNumberPartition         string    `bun:"type:varchar(50)"`
	LastRedirectDnPartition                 string    `bun:"type:varchar(50)"`
	Duration                                int       `bun:"type:integer"`
	OrigDeviceName                          string    `bun:"type:varchar(129)"`
	DestDeviceName                          string    `bun:"type:varchar(129)"`
	OrigCallTerminationOnBehalfOf           int       `bun:"type:integer"`
	DestCallTerminationOnBehalfOf           int       `bun:"type:integer"`
	OrigCalledPartyRedirectOnBehalfOf       int       `bun:"type:integer"`
	LastRedirectRedirectOnBehalfOf          int       `bun:"type:integer"`
	OrigCalledPartyRedirectReason           int       `bun:"type:integer"`
	LastRedirectRedirectReason              int       `bun:"type:integer"`
	DestConversationId                      int       `bun:"type:integer"`
	GlobalCallId_ClusterID                  string    `bun:"type:varchar(50)"`
	JoinOnBehalfOf                          int       `bun:"type:integer"`
	Comment                                 string    `bun:"type:varchar(2048)"`
	AuthCodeDescription                     string    `bun:"type:varchar(50)"`
	AuthorizationLevel                      int       `bun:"type:integer"`
	ClientMatterCode                        string    `bun:"type:varchar(32)"`
	OrigDTMFMethod                          int       `bun:"type:integer"`
	DestDTMFMethod                          int       `bun:"type:integer"`
	CallSecuredStatus                       int       `bun:"type:integer"`
	OrigConversationId                      int       `bun:"type:integer"`
	OrigMediaCap_Bandwidth                  int       `bun:"type:integer"`
	DestMediaCap_Bandwidth                  int       `bun:"type:integer"`
	AuthorizationCodeValue                  string    `bun:"type:varchar(32)"`
	OutpulsedCallingPartyNumber             string    `bun:"type:varchar(50)"`
	OutpulsedCalledPartyNumber              string    `bun:"type:varchar(50)"`
	OrigIpv4v6Addr                          string    `bun:"type:varchar(64)"`
	DestIpv4v6Addr                          string    `bun:"type:varchar(64)"`
	OrigVideoCap_Codec_Channel2             int       `bun:"type:integer"`
	OrigVideoCap_Bandwidth_Channel2         int       `bun:"type:integer"`
	OrigVideoCap_Resolution_Channel2        int       `bun:"type:integer"`
	OrigVideoTransportAddress_IP_Channel2   int       `bun:"type:integer"`
	OrigVideoTransportAddress_Port_Channel2 int       `bun:"type:integer"`
	OrigVideoChannel_Role_Channel2          int       `bun:"type:integer"`
	DestVideoCap_Codec_Channel2             int       `bun:"type:integer"`
	DestVideoCap_Bandwidth_Channel2         int       `bun:"type:integer"`
	DestVideoCap_Resolution_Channel2        int       `bun:"type:integer"`
	DestVideoTransportAddress_IP_Channel2   int       `bun:"type:integer"`
	DestVideoTransportAddress_Port_Channel2 int       `bun:"type:integer"`
	DestVideoChannel_Role_Channel2          int       `bun:"type:integer"`
	IncomingProtocolID                      int       `bun:"type:integer"`
	IncomingProtocolCallRef                 string    `bun:"type:varchar(32)"`
	OutgoingProtocolID                      int       `bun:"type:integer"`
	OutgoingProtocolCallRef                 string    `bun:"type:varchar(32)"`
	CurrentRoutingReason                    int       `bun:"type:integer"`
	OrigRoutingReason                       int       `bun:"type:integer"`
	LastRedirectingRoutingReason            int       `bun:"type:integer"`
	HuntPilotPartition                      string    `bun:"type:varchar(50)"`
	HuntPilotDN                             string    `bun:"type:varchar(50)"`
	CalledPartyPatternUsage                 int       `bun:"type:integer"`
	IncomingICID                            string    `bun:"type:varchar(50)"`
	IncomingOrigIOI                         string    `bun:"type:varchar(50)"`
	IncomingTermIOI                         string    `bun:"type:varchar(50)"`
	OutgoingICID                            string    `bun:"type:varchar(50)"`
	OutgoingOrigIOI                         string    `bun:"type:varchar(50)"`
	OutgoingTermIOI                         string    `bun:"type:varchar(50)"`
	OutpulsedOriginalCalledPartyNumber      string    `bun:"type:varchar(50)"`
	OutpulsedLastRedirectingNumber          string    `bun:"type:varchar(50)"`
	WasCallQueued                           int       `bun:"type:integer"`
	TotalWaitTimeInQueue                    int       `bun:"type:integer"`
	CallingPartyNumber_uri                  string    `bun:"type:varchar(255)"`
	OriginalCalledPartyNumber_uri           string    `bun:"type:varchar(255)"`
	FinalCalledPartyNumber_uri              string    `bun:"type:varchar(255)"`
	LastRedirectDn_uri                      string    `bun:"type:varchar(255)"`
	MobileCallingPartyNumber                string    `bun:"type:varchar(50)"`
	FinalMobileCalledPartyNumber            string    `bun:"type:varchar(50)"`
	OrigMobileDeviceName                    string    `bun:"type:varchar(129)"`
	DestMobileDeviceName                    string    `bun:"type:varchar(129)"`
	OrigMobileCallDuration                  int       `bun:"type:integer"`
	DestMobileCallDuration                  int       `bun:"type:integer"`
	MobileCallType                          int       `bun:"type:integer"`
	OriginalCalledPartyPattern              string    `bun:"type:varchar(50)"`
	FinalCalledPartyPattern                 string    `bun:"type:varchar(50)"`
	LastRedirectingPartyPattern             string    `bun:"type:varchar(50)"`
	HuntPilotPattern                        string    `bun:"type:varchar(50)"`
	OrigDeviceType                          string    `bun:"type:varchar(100)"`
	DestDeviceType                          string    `bun:"type:varchar(100)"`
}

// NewRecordFromCanonical creates a new Record DB object from canonical model.
func NewRecordFromCanonical(obj model.Record) Record {
	return Record{
		CdrRecordType:                           obj.CdrRecordType,
		GlobalCallID_callManagerId:              obj.GlobalCallID_callManagerId,
		GlobalCallID_callId:                     obj.GlobalCallID_callId,
		OrigLegCallIdentifier:                   obj.OrigLegCallIdentifier,
		DateTimeOrigination:                     obj.DateTimeOrigination,
		OrigNodeId:                              obj.OrigNodeId,
		OrigSpan:                                obj.OrigSpan,
		OrigIpAddr:                              obj.OrigIpAddr,
		CallingPartyNumber:                      obj.CallingPartyNumber,
		CallingPartyUnicodeLoginUserID:          obj.CallingPartyUnicodeLoginUserID,
		OrigCause_location:                      obj.OrigCause_location,
		OrigCause_value:                         obj.OrigCause_value,
		OrigPrecedenceLevel:                     obj.OrigPrecedenceLevel,
		OrigMediaTransportAddress_IP:            obj.OrigMediaTransportAddress_IP,
		OrigMediaTransportAddress_Port:          obj.OrigMediaTransportAddress_Port,
		OrigMediaCap_payloadCapability:          obj.OrigMediaCap_payloadCapability,
		OrigMediaCap_maxFramesPerPacket:         obj.OrigMediaCap_maxFramesPerPacket,
		OrigMediaCap_g723BitRate:                obj.OrigMediaCap_g723BitRate,
		OrigVideoCap_Codec:                      obj.OrigVideoCap_Codec,
		OrigVideoCap_Bandwidth:                  obj.OrigVideoCap_Bandwidth,
		OrigVideoCap_Resolution:                 obj.OrigVideoCap_Resolution,
		OrigVideoTransportAddress_IP:            obj.OrigVideoTransportAddress_IP,
		OrigVideoTransportAddress_Port:          obj.OrigVideoTransportAddress_Port,
		OrigRSVPAudioStat:                       obj.OrigRSVPAudioStat,
		OrigRSVPVideoStat:                       obj.OrigRSVPVideoStat,
		DestLegIdentifier:                       obj.DestLegIdentifier,
		DestNodeId:                              obj.DestNodeId,
		DestSpan:                                obj.DestSpan,
		DestIpAddr:                              obj.DestIpAddr,
		OriginalCalledPartyNumber:               obj.OriginalCalledPartyNumber,
		FinalCalledPartyNumber:                  obj.FinalCalledPartyNumber,
		FinalCalledPartyUnicodeLoginUserID:      obj.FinalCalledPartyUnicodeLoginUserID,
		DestCause_location:                      obj.DestCause_location,
		DestCause_value:                         obj.DestCause_value,
		DestPrecedenceLevel:                     obj.DestPrecedenceLevel,
		DestMediaTransportAddress_IP:            obj.DestMediaTransportAddress_IP,
		DestMediaTransportAddress_Port:          obj.DestMediaTransportAddress_Port,
		DestMediaCap_payloadCapability:          obj.DestMediaCap_payloadCapability,
		DestMediaCap_maxFramesPerPacket:         obj.DestMediaCap_maxFramesPerPacket,
		DestMediaCap_g723BitRate:                obj.DestMediaCap_g723BitRate,
		DestVideoCap_Codec:                      obj.DestVideoCap_Codec,
		DestVideoCap_Bandwidth:                  obj.DestVideoCap_Bandwidth,
		DestVideoCap_Resolution:                 obj.DestVideoCap_Resolution,
		DestVideoTransportAddress_IP:            obj.DestVideoTransportAddress_IP,
		DestVideoTransportAddress_Port:          obj.DestVideoTransportAddress_Port,
		DestRSVPAudioStat:                       obj.DestRSVPAudioStat,
		DestRSVPVideoStat:                       obj.DestRSVPVideoStat,
		DateTimeConnect:                         obj.DateTimeConnect,
		DateTimeDisconnect:                      obj.DateTimeDisconnect,
		LastRedirectDn:                          obj.LastRedirectDn,
		Pkid:                                    obj.Pkid,
		OriginalCalledPartyNumberPartition:      obj.OriginalCalledPartyNumberPartition,
		CallingPartyNumberPartition:             obj.CallingPartyNumberPartition,
		FinalCalledPartyNumberPartition:         obj.FinalCalledPartyNumberPartition,
		LastRedirectDnPartition:                 obj.LastRedirectDnPartition,
		Duration:                                obj.Duration,
		OrigDeviceName:                          obj.OrigDeviceName,
		DestDeviceName:                          obj.DestDeviceName,
		OrigCallTerminationOnBehalfOf:           obj.OrigCallTerminationOnBehalfOf,
		DestCallTerminationOnBehalfOf:           obj.DestCallTerminationOnBehalfOf,
		OrigCalledPartyRedirectOnBehalfOf:       obj.OrigCalledPartyRedirectOnBehalfOf,
		LastRedirectRedirectOnBehalfOf:          obj.LastRedirectRedirectOnBehalfOf,
		OrigCalledPartyRedirectReason:           obj.OrigCalledPartyRedirectReason,
		LastRedirectRedirectReason:              obj.LastRedirectRedirectReason,
		DestConversationId:                      obj.DestConversationId,
		GlobalCallId_ClusterID:                  obj.GlobalCallId_ClusterID,
		JoinOnBehalfOf:                          obj.JoinOnBehalfOf,
		Comment:                                 obj.Comment,
		AuthCodeDescription:                     obj.AuthCodeDescription,
		AuthorizationLevel:                      obj.AuthorizationLevel,
		ClientMatterCode:                        obj.ClientMatterCode,
		OrigDTMFMethod:                          obj.OrigDTMFMethod,
		DestDTMFMethod:                          obj.DestDTMFMethod,
		CallSecuredStatus:                       obj.CallSecuredStatus,
		OrigConversationId:                      obj.OrigConversationId,
		OrigMediaCap_Bandwidth:                  obj.OrigMediaCap_Bandwidth,
		DestMediaCap_Bandwidth:                  obj.DestMediaCap_Bandwidth,
		AuthorizationCodeValue:                  obj.AuthorizationCodeValue,
		OutpulsedCallingPartyNumber:             obj.OutpulsedCallingPartyNumber,
		OutpulsedCalledPartyNumber:              obj.OutpulsedCalledPartyNumber,
		OrigIpv4v6Addr:                          obj.OrigIpv4v6Addr,
		DestIpv4v6Addr:                          obj.DestIpv4v6Addr,
		OrigVideoCap_Codec_Channel2:             obj.OrigVideoCap_Codec_Channel2,
		OrigVideoCap_Bandwidth_Channel2:         obj.OrigVideoCap_Bandwidth_Channel2,
		OrigVideoCap_Resolution_Channel2:        obj.OrigVideoCap_Resolution_Channel2,
		OrigVideoTransportAddress_IP_Channel2:   obj.OrigVideoTransportAddress_IP_Channel2,
		OrigVideoTransportAddress_Port_Channel2: obj.OrigVideoTransportAddress_Port_Channel2,
		OrigVideoChannel_Role_Channel2:          obj.OrigVideoChannel_Role_Channel2,
		DestVideoCap_Codec_Channel2:             obj.DestVideoCap_Codec_Channel2,
		DestVideoCap_Bandwidth_Channel2:         obj.DestVideoCap_Bandwidth_Channel2,
		DestVideoCap_Resolution_Channel2:        obj.DestVideoCap_Resolution_Channel2,
		DestVideoTransportAddress_IP_Channel2:   obj.DestVideoTransportAddress_IP_Channel2,
		DestVideoTransportAddress_Port_Channel2: obj.DestVideoTransportAddress_Port_Channel2,
		DestVideoChannel_Role_Channel2:          obj.DestVideoChannel_Role_Channel2,
		IncomingProtocolID:                      obj.IncomingProtocolID,
		IncomingProtocolCallRef:                 obj.IncomingProtocolCallRef,
		OutgoingProtocolID:                      obj.OutgoingProtocolID,
		OutgoingProtocolCallRef:                 obj.OutgoingProtocolCallRef,
		CurrentRoutingReason:                    obj.CurrentRoutingReason,
		OrigRoutingReason:                       obj.OrigRoutingReason,
		LastRedirectingRoutingReason:            obj.LastRedirectingRoutingReason,
		HuntPilotPartition:                      obj.HuntPilotPartition,
		HuntPilotDN:                             obj.HuntPilotDN,
		CalledPartyPatternUsage:                 obj.CalledPartyPatternUsage,
		IncomingICID:                            obj.IncomingICID,
		IncomingOrigIOI:                         obj.IncomingOrigIOI,
		IncomingTermIOI:                         obj.IncomingTermIOI,
		OutgoingICID:                            obj.OutgoingICID,
		OutgoingOrigIOI:                         obj.OutgoingOrigIOI,
		OutgoingTermIOI:                         obj.OutgoingTermIOI,
		OutpulsedOriginalCalledPartyNumber:      obj.OutpulsedOriginalCalledPartyNumber,
		OutpulsedLastRedirectingNumber:          obj.OutpulsedLastRedirectingNumber,
		WasCallQueued:                           obj.WasCallQueued,
		TotalWaitTimeInQueue:                    obj.TotalWaitTimeInQueue,
		CallingPartyNumber_uri:                  obj.CallingPartyNumber_uri,
		OriginalCalledPartyNumber_uri:           obj.OriginalCalledPartyNumber_uri,
		FinalCalledPartyNumber_uri:              obj.FinalCalledPartyNumber_uri,
		LastRedirectDn_uri:                      obj.LastRedirectDn_uri,
		MobileCallingPartyNumber:                obj.MobileCallingPartyNumber,
		FinalMobileCalledPartyNumber:            obj.FinalMobileCalledPartyNumber,
		OrigMobileDeviceName:                    obj.OrigMobileDeviceName,
		DestMobileDeviceName:                    obj.DestMobileDeviceName,
		OrigMobileCallDuration:                  obj.OrigMobileCallDuration,
		DestMobileCallDuration:                  obj.DestMobileCallDuration,
		MobileCallType:                          obj.MobileCallType,
		OriginalCalledPartyPattern:              obj.OriginalCalledPartyPattern,
		FinalCalledPartyPattern:                 obj.FinalCalledPartyPattern,
		LastRedirectingPartyPattern:             obj.LastRedirectingPartyPattern,
		HuntPilotPattern:                        obj.HuntPilotPattern,
		OrigDeviceType:                          obj.OrigDeviceType,
		DestDeviceType:                          obj.DestDeviceType,
	}
}

// ToCanonical converts a DB object to canonical model.
func (r Record) ToCanonical() (model.Record, error) {
	return model.Record{
		CdrRecordType:                           r.CdrRecordType,
		GlobalCallID_callManagerId:              r.GlobalCallID_callManagerId,
		GlobalCallID_callId:                     r.GlobalCallID_callId,
		OrigLegCallIdentifier:                   r.OrigLegCallIdentifier,
		DateTimeOrigination:                     r.DateTimeOrigination,
		OrigNodeId:                              r.OrigNodeId,
		OrigSpan:                                r.OrigSpan,
		OrigIpAddr:                              r.OrigIpAddr,
		CallingPartyNumber:                      r.CallingPartyNumber,
		CallingPartyUnicodeLoginUserID:          r.CallingPartyUnicodeLoginUserID,
		OrigCause_location:                      r.OrigCause_location,
		OrigCause_value:                         r.OrigCause_value,
		OrigPrecedenceLevel:                     r.OrigPrecedenceLevel,
		OrigMediaTransportAddress_IP:            r.OrigMediaTransportAddress_IP,
		OrigMediaTransportAddress_Port:          r.OrigMediaTransportAddress_Port,
		OrigMediaCap_payloadCapability:          r.OrigMediaCap_payloadCapability,
		OrigMediaCap_maxFramesPerPacket:         r.OrigMediaCap_maxFramesPerPacket,
		OrigMediaCap_g723BitRate:                r.OrigMediaCap_g723BitRate,
		OrigVideoCap_Codec:                      r.OrigVideoCap_Codec,
		OrigVideoCap_Bandwidth:                  r.OrigVideoCap_Bandwidth,
		OrigVideoCap_Resolution:                 r.OrigVideoCap_Resolution,
		OrigVideoTransportAddress_IP:            r.OrigVideoTransportAddress_IP,
		OrigVideoTransportAddress_Port:          r.OrigVideoTransportAddress_Port,
		OrigRSVPAudioStat:                       r.OrigRSVPAudioStat,
		OrigRSVPVideoStat:                       r.OrigRSVPVideoStat,
		DestLegIdentifier:                       r.DestLegIdentifier,
		DestNodeId:                              r.DestNodeId,
		DestSpan:                                r.DestSpan,
		DestIpAddr:                              r.DestIpAddr,
		OriginalCalledPartyNumber:               r.OriginalCalledPartyNumber,
		FinalCalledPartyNumber:                  r.FinalCalledPartyNumber,
		FinalCalledPartyUnicodeLoginUserID:      r.FinalCalledPartyUnicodeLoginUserID,
		DestCause_location:                      r.DestCause_location,
		DestCause_value:                         r.DestCause_value,
		DestPrecedenceLevel:                     r.DestPrecedenceLevel,
		DestMediaTransportAddress_IP:            r.DestMediaTransportAddress_IP,
		DestMediaTransportAddress_Port:          r.DestMediaTransportAddress_Port,
		DestMediaCap_payloadCapability:          r.DestMediaCap_payloadCapability,
		DestMediaCap_maxFramesPerPacket:         r.DestMediaCap_maxFramesPerPacket,
		DestMediaCap_g723BitRate:                r.DestMediaCap_g723BitRate,
		DestVideoCap_Codec:                      r.DestVideoCap_Codec,
		DestVideoCap_Bandwidth:                  r.DestVideoCap_Bandwidth,
		DestVideoCap_Resolution:                 r.DestVideoCap_Resolution,
		DestVideoTransportAddress_IP:            r.DestVideoTransportAddress_IP,
		DestVideoTransportAddress_Port:          r.DestVideoTransportAddress_Port,
		DestRSVPAudioStat:                       r.DestRSVPAudioStat,
		DestRSVPVideoStat:                       r.DestRSVPVideoStat,
		DateTimeConnect:                         r.DateTimeConnect,
		DateTimeDisconnect:                      r.DateTimeDisconnect,
		LastRedirectDn:                          r.LastRedirectDn,
		Pkid:                                    r.Pkid,
		OriginalCalledPartyNumberPartition:      r.OriginalCalledPartyNumberPartition,
		CallingPartyNumberPartition:             r.CallingPartyNumberPartition,
		FinalCalledPartyNumberPartition:         r.FinalCalledPartyNumberPartition,
		LastRedirectDnPartition:                 r.LastRedirectDnPartition,
		Duration:                                r.Duration,
		OrigDeviceName:                          r.OrigDeviceName,
		DestDeviceName:                          r.DestDeviceName,
		OrigCallTerminationOnBehalfOf:           r.OrigCallTerminationOnBehalfOf,
		DestCallTerminationOnBehalfOf:           r.DestCallTerminationOnBehalfOf,
		OrigCalledPartyRedirectOnBehalfOf:       r.OrigCalledPartyRedirectOnBehalfOf,
		LastRedirectRedirectOnBehalfOf:          r.LastRedirectRedirectOnBehalfOf,
		OrigCalledPartyRedirectReason:           r.OrigCalledPartyRedirectReason,
		LastRedirectRedirectReason:              r.LastRedirectRedirectReason,
		DestConversationId:                      r.DestConversationId,
		GlobalCallId_ClusterID:                  r.GlobalCallId_ClusterID,
		JoinOnBehalfOf:                          r.JoinOnBehalfOf,
		Comment:                                 r.Comment,
		AuthCodeDescription:                     r.AuthCodeDescription,
		AuthorizationLevel:                      r.AuthorizationLevel,
		ClientMatterCode:                        r.ClientMatterCode,
		OrigDTMFMethod:                          r.OrigDTMFMethod,
		DestDTMFMethod:                          r.DestDTMFMethod,
		CallSecuredStatus:                       r.CallSecuredStatus,
		OrigConversationId:                      r.OrigConversationId,
		OrigMediaCap_Bandwidth:                  r.OrigMediaCap_Bandwidth,
		DestMediaCap_Bandwidth:                  r.DestMediaCap_Bandwidth,
		AuthorizationCodeValue:                  r.AuthorizationCodeValue,
		OutpulsedCallingPartyNumber:             r.OutpulsedCallingPartyNumber,
		OutpulsedCalledPartyNumber:              r.OutpulsedCalledPartyNumber,
		OrigIpv4v6Addr:                          r.OrigIpv4v6Addr,
		DestIpv4v6Addr:                          r.DestIpv4v6Addr,
		OrigVideoCap_Codec_Channel2:             r.OrigVideoCap_Codec_Channel2,
		OrigVideoCap_Bandwidth_Channel2:         r.OrigVideoCap_Bandwidth_Channel2,
		OrigVideoCap_Resolution_Channel2:        r.OrigVideoCap_Resolution_Channel2,
		OrigVideoTransportAddress_IP_Channel2:   r.OrigVideoTransportAddress_IP_Channel2,
		OrigVideoTransportAddress_Port_Channel2: r.OrigVideoTransportAddress_Port_Channel2,
		OrigVideoChannel_Role_Channel2:          r.OrigVideoChannel_Role_Channel2,
		DestVideoCap_Codec_Channel2:             r.DestVideoCap_Codec_Channel2,
		DestVideoCap_Bandwidth_Channel2:         r.DestVideoCap_Bandwidth_Channel2,
		DestVideoCap_Resolution_Channel2:        r.DestVideoCap_Resolution_Channel2,
		DestVideoTransportAddress_IP_Channel2:   r.DestVideoTransportAddress_IP_Channel2,
		DestVideoTransportAddress_Port_Channel2: r.DestVideoTransportAddress_Port_Channel2,
		DestVideoChannel_Role_Channel2:          r.DestVideoChannel_Role_Channel2,
		IncomingProtocolID:                      r.IncomingProtocolID,
		IncomingProtocolCallRef:                 r.IncomingProtocolCallRef,
		OutgoingProtocolID:                      r.OutgoingProtocolID,
		OutgoingProtocolCallRef:                 r.OutgoingProtocolCallRef,
		CurrentRoutingReason:                    r.CurrentRoutingReason,
		OrigRoutingReason:                       r.OrigRoutingReason,
		LastRedirectingRoutingReason:            r.LastRedirectingRoutingReason,
		HuntPilotPartition:                      r.HuntPilotPartition,
		HuntPilotDN:                             r.HuntPilotDN,
		CalledPartyPatternUsage:                 r.CalledPartyPatternUsage,
		IncomingICID:                            r.IncomingICID,
		IncomingOrigIOI:                         r.IncomingOrigIOI,
		IncomingTermIOI:                         r.IncomingTermIOI,
		OutgoingICID:                            r.OutgoingICID,
		OutgoingOrigIOI:                         r.OutgoingOrigIOI,
		OutgoingTermIOI:                         r.OutgoingTermIOI,
		OutpulsedOriginalCalledPartyNumber:      r.OutpulsedOriginalCalledPartyNumber,
		OutpulsedLastRedirectingNumber:          r.OutpulsedLastRedirectingNumber,
		WasCallQueued:                           r.WasCallQueued,
		TotalWaitTimeInQueue:                    r.TotalWaitTimeInQueue,
		CallingPartyNumber_uri:                  r.CallingPartyNumber_uri,
		OriginalCalledPartyNumber_uri:           r.OriginalCalledPartyNumber_uri,
		FinalCalledPartyNumber_uri:              r.FinalCalledPartyNumber_uri,
		LastRedirectDn_uri:                      r.LastRedirectDn_uri,
		MobileCallingPartyNumber:                r.MobileCallingPartyNumber,
		FinalMobileCalledPartyNumber:            r.FinalMobileCalledPartyNumber,
		OrigMobileDeviceName:                    r.OrigMobileDeviceName,
		DestMobileDeviceName:                    r.DestMobileDeviceName,
		OrigMobileCallDuration:                  r.OrigMobileCallDuration,
		DestMobileCallDuration:                  r.DestMobileCallDuration,
		MobileCallType:                          r.MobileCallType,
		OriginalCalledPartyPattern:              r.OriginalCalledPartyPattern,
		FinalCalledPartyPattern:                 r.FinalCalledPartyPattern,
		LastRedirectingPartyPattern:             r.LastRedirectingPartyPattern,
		HuntPilotPattern:                        r.HuntPilotPattern,
		OrigDeviceType:                          r.OrigDeviceType,
		DestDeviceType:                          r.DestDeviceType,
	}, nil
}
