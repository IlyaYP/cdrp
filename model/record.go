package model

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var CdrHeader = [...]string{
	"cdrRecordType",
	"globalCallID_callManagerId",
	"globalCallID_callId",
	"origLegCallIdentifier",
	"dateTimeOrigination",
	"origNodeId",
	"origSpan",
	"origIpAddr",
	"callingPartyNumber",
	"callingPartyUnicodeLoginUserID",
	"origCause_location",
	"origCause_value",
	"origPrecedenceLevel",
	"origMediaTransportAddress_IP",
	"origMediaTransportAddress_Port",
	"origMediaCap_payloadCapability",
	"origMediaCap_maxFramesPerPacket",
	"origMediaCap_g723BitRate",
	"origVideoCap_Codec",
	"origVideoCap_Bandwidth",
	"origVideoCap_Resolution",
	"origVideoTransportAddress_IP",
	"origVideoTransportAddress_Port",
	"origRSVPAudioStat",
	"origRSVPVideoStat",
	"destLegIdentifier",
	"destNodeId",
	"destSpan",
	"destIpAddr",
	"originalCalledPartyNumber",
	"finalCalledPartyNumber",
	"finalCalledPartyUnicodeLoginUserID",
	"destCause_location",
	"destCause_value",
	"destPrecedenceLevel",
	"destMediaTransportAddress_IP",
	"destMediaTransportAddress_Port",
	"destMediaCap_payloadCapability",
	"destMediaCap_maxFramesPerPacket",
	"destMediaCap_g723BitRate",
	"destVideoCap_Codec",
	"destVideoCap_Bandwidth",
	"destVideoCap_Resolution",
	"destVideoTransportAddress_IP",
	"destVideoTransportAddress_Port",
	"destRSVPAudioStat",
	"destRSVPVideoStat",
	"dateTimeConnect",
	"dateTimeDisconnect",
	"lastRedirectDn",
	"pkid",
	"originalCalledPartyNumberPartition",
	"callingPartyNumberPartition",
	"finalCalledPartyNumberPartition",
	"lastRedirectDnPartition",
	"duration",
	"origDeviceName",
	"destDeviceName",
	"origCallTerminationOnBehalfOf",
	"destCallTerminationOnBehalfOf",
	"origCalledPartyRedirectOnBehalfOf",
	"lastRedirectRedirectOnBehalfOf",
	"origCalledPartyRedirectReason",
	"lastRedirectRedirectReason",
	"destConversationId",
	"globalCallId_ClusterID",
	"joinOnBehalfOf",
	"comment",
	"authCodeDescription",
	"authorizationLevel",
	"clientMatterCode",
	"origDTMFMethod",
	"destDTMFMethod",
	"callSecuredStatus",
	"origConversationId",
	"origMediaCap_Bandwidth",
	"destMediaCap_Bandwidth",
	"authorizationCodeValue",
	"outpulsedCallingPartyNumber",
	"outpulsedCalledPartyNumber",
	"origIpv4v6Addr",
	"destIpv4v6Addr",
	"origVideoCap_Codec_Channel2",
	"origVideoCap_Bandwidth_Channel2",
	"origVideoCap_Resolution_Channel2",
	"origVideoTransportAddress_IP_Channel2",
	"origVideoTransportAddress_Port_Channel2",
	"origVideoChannel_Role_Channel2",
	"destVideoCap_Codec_Channel2",
	"destVideoCap_Bandwidth_Channel2",
	"destVideoCap_Resolution_Channel2",
	"destVideoTransportAddress_IP_Channel2",
	"destVideoTransportAddress_Port_Channel2",
	"destVideoChannel_Role_Channel2",
	"IncomingProtocolID",
	"IncomingProtocolCallRef",
	"OutgoingProtocolID",
	"OutgoingProtocolCallRef",
	"currentRoutingReason",
	"origRoutingReason",
	"lastRedirectingRoutingReason",
	"huntPilotPartition",
	"huntPilotDN",
	"calledPartyPatternUsage",
	"IncomingICID",
	"IncomingOrigIOI",
	"IncomingTermIOI",
	"OutgoingICID",
	"OutgoingOrigIOI",
	"OutgoingTermIOI",
	"outpulsedOriginalCalledPartyNumber",
	"outpulsedLastRedirectingNumber",
	"wasCallQueued",
	"totalWaitTimeInQueue",
	"callingPartyNumber_uri",
	"originalCalledPartyNumber_uri",
	"finalCalledPartyNumber_uri",
	"lastRedirectDn_uri",
	"mobileCallingPartyNumber",
	"finalMobileCalledPartyNumber",
	"origMobileDeviceName",
	"destMobileDeviceName",
	"origMobileCallDuration",
	"destMobileCallDuration",
	"mobileCallType",
	"originalCalledPartyPattern",
	"finalCalledPartyPattern",
	"lastRedirectingPartyPattern",
	"huntPilotPattern",
	"origDeviceType",
	"destDeviceType",
}

const (
	CdrRecordType = iota
	GlobalCallID_callManagerId
	GlobalCallID_callId
	OrigLegCallIdentifier
	DateTimeOrigination
	OrigNodeId
	OrigSpan
	OrigIpAddr
	CallingPartyNumber
	CallingPartyUnicodeLoginUserID
	OrigCause_location
	OrigCause_value
	OrigPrecedenceLevel
	OrigMediaTransportAddress_IP
	OrigMediaTransportAddress_Port
	OrigMediaCap_payloadCapability
	OrigMediaCap_maxFramesPerPacket
	OrigMediaCap_g723BitRate
	OrigVideoCap_Codec
	OrigVideoCap_Bandwidth
	OrigVideoCap_Resolution
	OrigVideoTransportAddress_IP
	OrigVideoTransportAddress_Port
	OrigRSVPAudioStat
	OrigRSVPVideoStat
	DestLegIdentifier
	DestNodeId
	DestSpan
	DestIpAddr
	OriginalCalledPartyNumber
	FinalCalledPartyNumber
	FinalCalledPartyUnicodeLoginUserID
	DestCause_location
	DestCause_value
	DestPrecedenceLevel
	DestMediaTransportAddress_IP
	DestMediaTransportAddress_Port
	DestMediaCap_payloadCapability
	DestMediaCap_maxFramesPerPacket
	DestMediaCap_g723BitRate
	DestVideoCap_Codec
	DestVideoCap_Bandwidth
	DestVideoCap_Resolution
	DestVideoTransportAddress_IP
	DestVideoTransportAddress_Port
	DestRSVPAudioStat
	DestRSVPVideoStat
	DateTimeConnect
	DateTimeDisconnect
	LastRedirectDn
	Pkid
	OriginalCalledPartyNumberPartition
	CallingPartyNumberPartition
	FinalCalledPartyNumberPartition
	LastRedirectDnPartition
	Duration
	OrigDeviceName
	DestDeviceName
	OrigCallTerminationOnBehalfOf
	DestCallTerminationOnBehalfOf
	OrigCalledPartyRedirectOnBehalfOf
	LastRedirectRedirectOnBehalfOf
	OrigCalledPartyRedirectReason
	LastRedirectRedirectReason
	DestConversationId
	GlobalCallId_ClusterID
	JoinOnBehalfOf
	Comment
	AuthCodeDescription
	AuthorizationLevel
	ClientMatterCode
	OrigDTMFMethod
	DestDTMFMethod
	CallSecuredStatus
	OrigConversationId
	OrigMediaCap_Bandwidth
	DestMediaCap_Bandwidth
	AuthorizationCodeValue
	OutpulsedCallingPartyNumber
	OutpulsedCalledPartyNumber
	OrigIpv4v6Addr
	DestIpv4v6Addr
	OrigVideoCap_Codec_Channel2
	OrigVideoCap_Bandwidth_Channel2
	OrigVideoCap_Resolution_Channel2
	OrigVideoTransportAddress_IP_Channel2
	OrigVideoTransportAddress_Port_Channel2
	OrigVideoChannel_Role_Channel2
	DestVideoCap_Codec_Channel2
	DestVideoCap_Bandwidth_Channel2
	DestVideoCap_Resolution_Channel2
	DestVideoTransportAddress_IP_Channel2
	DestVideoTransportAddress_Port_Channel2
	DestVideoChannel_Role_Channel2
	IncomingProtocolID
	IncomingProtocolCallRef
	OutgoingProtocolID
	OutgoingProtocolCallRef
	CurrentRoutingReason
	OrigRoutingReason
	LastRedirectingRoutingReason
	HuntPilotPartition
	HuntPilotDN
	CalledPartyPatternUsage
	IncomingICID
	IncomingOrigIOI
	IncomingTermIOI
	OutgoingICID
	OutgoingOrigIOI
	OutgoingTermIOI
	OutpulsedOriginalCalledPartyNumber
	OutpulsedLastRedirectingNumber
	WasCallQueued
	TotalWaitTimeInQueue
	CallingPartyNumber_uri
	OriginalCalledPartyNumber_uri
	FinalCalledPartyNumber_uri
	LastRedirectDn_uri
	MobileCallingPartyNumber
	FinalMobileCalledPartyNumber
	OrigMobileDeviceName
	DestMobileDeviceName
	OrigMobileCallDuration
	DestMobileCallDuration
	MobileCallType
	OriginalCalledPartyPattern
	FinalCalledPartyPattern
	LastRedirectingPartyPattern
	HuntPilotPattern
	OrigDeviceType
	DestDeviceType
	RecordLen
)

type Record struct {
	CdrRecordType                           int       //INTEGER
	GlobalCallID_callManagerId              int       //INTEGER
	GlobalCallID_callId                     int       //INTEGER
	OrigLegCallIdentifier                   int       //INTEGER
	DateTimeOrigination                     time.Time //INTEGER
	OrigNodeId                              int       //INTEGER
	OrigSpan                                int       //INTEGER
	OrigIpAddr                              int       //INTEGER
	CallingPartyNumber                      string    //VARCHAR(50)
	CallingPartyUnicodeLoginUserID          string    //VARCHAR(128)
	OrigCause_location                      int       //INTEGER
	OrigCause_value                         int       //INTEGER
	OrigPrecedenceLevel                     int       //INTEGER
	OrigMediaTransportAddress_IP            int       //INTEGER
	OrigMediaTransportAddress_Port          int       //INTEGER
	OrigMediaCap_payloadCapability          int       //INTEGER
	OrigMediaCap_maxFramesPerPacket         int       //INTEGER
	OrigMediaCap_g723BitRate                int       //INTEGER
	OrigVideoCap_Codec                      int       //INTEGER
	OrigVideoCap_Bandwidth                  int       //INTEGER
	OrigVideoCap_Resolution                 int       //INTEGER
	OrigVideoTransportAddress_IP            int       //INTEGER
	OrigVideoTransportAddress_Port          int       //INTEGER
	OrigRSVPAudioStat                       string    //VARCHAR(64)
	OrigRSVPVideoStat                       string    //VARCHAR(64)
	DestLegIdentifier                       int       //INTEGER
	DestNodeId                              int       //INTEGER
	DestSpan                                int       //INTEGER
	DestIpAddr                              int       //INTEGER
	OriginalCalledPartyNumber               string    //VARCHAR(50)
	FinalCalledPartyNumber                  string    //VARCHAR(50)
	FinalCalledPartyUnicodeLoginUserID      string    //VARCHAR(128)
	DestCause_location                      int       //INTEGER
	DestCause_value                         int       //INTEGER
	DestPrecedenceLevel                     int       //INTEGER
	DestMediaTransportAddress_IP            int       //INTEGER
	DestMediaTransportAddress_Port          int       //INTEGER
	DestMediaCap_payloadCapability          int       //INTEGER
	DestMediaCap_maxFramesPerPacket         int       //INTEGER
	DestMediaCap_g723BitRate                int       //INTEGER
	DestVideoCap_Codec                      int       //INTEGER
	DestVideoCap_Bandwidth                  int       //INTEGER
	DestVideoCap_Resolution                 int       //INTEGER
	DestVideoTransportAddress_IP            int       //INTEGER
	DestVideoTransportAddress_Port          int       //INTEGER
	DestRSVPAudioStat                       string    //VARCHAR(64)
	DestRSVPVideoStat                       string    //VARCHAR(64)
	DateTimeConnect                         time.Time //INTEGER
	DateTimeDisconnect                      time.Time //INTEGER
	LastRedirectDn                          string    //VARCHAR(50)
	Pkid                                    uuid.UUID //UNIQUEIDENTIFIER
	OriginalCalledPartyNumberPartition      string    //VARCHAR(50)
	CallingPartyNumberPartition             string    //VARCHAR(50)
	FinalCalledPartyNumberPartition         string    //VARCHAR(50)
	LastRedirectDnPartition                 string    //VARCHAR(50)
	Duration                                int       //INTEGER
	OrigDeviceName                          string    //VARCHAR(129)
	DestDeviceName                          string    //VARCHAR(129)
	OrigCallTerminationOnBehalfOf           int       //INTEGER
	DestCallTerminationOnBehalfOf           int       //INTEGER
	OrigCalledPartyRedirectOnBehalfOf       int       //INTEGER
	LastRedirectRedirectOnBehalfOf          int       //INTEGER
	OrigCalledPartyRedirectReason           int       //INTEGER
	LastRedirectRedirectReason              int       //INTEGER
	DestConversationId                      int       //INTEGER
	GlobalCallId_ClusterID                  string    //VARCHAR(50)
	JoinOnBehalfOf                          int       //INTEGER
	Comment                                 string    //VARCHAR(2048)
	AuthCodeDescription                     string    //VARCHAR(50)
	AuthorizationLevel                      int       //INTEGER
	ClientMatterCode                        string    //VARCHAR(32)
	OrigDTMFMethod                          int       //INTEGER
	DestDTMFMethod                          int       //INTEGER
	CallSecuredStatus                       int       //INTEGER
	OrigConversationId                      int       //INTEGER
	OrigMediaCap_Bandwidth                  int       //INTEGER
	DestMediaCap_Bandwidth                  int       //INTEGER
	AuthorizationCodeValue                  string    //VARCHAR(32)
	OutpulsedCallingPartyNumber             string    //VARCHAR(50)
	OutpulsedCalledPartyNumber              string    //VARCHAR(50)
	OrigIpv4v6Addr                          string    //VARCHAR(64)
	DestIpv4v6Addr                          string    //VARCHAR(64)
	OrigVideoCap_Codec_Channel2             int       //INTEGER
	OrigVideoCap_Bandwidth_Channel2         int       //INTEGER
	OrigVideoCap_Resolution_Channel2        int       //INTEGER
	OrigVideoTransportAddress_IP_Channel2   int       //INTEGER
	OrigVideoTransportAddress_Port_Channel2 int       //INTEGER
	OrigVideoChannel_Role_Channel2          int       //INTEGER
	DestVideoCap_Codec_Channel2             int       //INTEGER
	DestVideoCap_Bandwidth_Channel2         int       //INTEGER
	DestVideoCap_Resolution_Channel2        int       //INTEGER
	DestVideoTransportAddress_IP_Channel2   int       //INTEGER
	DestVideoTransportAddress_Port_Channel2 int       //INTEGER
	DestVideoChannel_Role_Channel2          int       //INTEGER
	IncomingProtocolID                      int       //INTEGER
	IncomingProtocolCallRef                 string    //VARCHAR(32)
	OutgoingProtocolID                      int       //INTEGER
	OutgoingProtocolCallRef                 string    //VARCHAR(32)
	CurrentRoutingReason                    int       //INTEGER
	OrigRoutingReason                       int       //INTEGER
	LastRedirectingRoutingReason            int       //INTEGER
	HuntPilotPartition                      string    //VARCHAR(50)
	HuntPilotDN                             string    //VARCHAR(50)
	CalledPartyPatternUsage                 int       //INTEGER
	IncomingICID                            string    //VARCHAR(50)
	IncomingOrigIOI                         string    //VARCHAR(50)
	IncomingTermIOI                         string    //VARCHAR(50)
	OutgoingICID                            string    //VARCHAR(50)
	OutgoingOrigIOI                         string    //VARCHAR(50)
	OutgoingTermIOI                         string    //VARCHAR(50)
	OutpulsedOriginalCalledPartyNumber      string    //VARCHAR(50)
	OutpulsedLastRedirectingNumber          string    //VARCHAR(50)
	WasCallQueued                           int       //INTEGER
	TotalWaitTimeInQueue                    int       //INTEGER
	CallingPartyNumber_uri                  string    //VARCHAR(255)
	OriginalCalledPartyNumber_uri           string    //VARCHAR(255)
	FinalCalledPartyNumber_uri              string    //VARCHAR(255)
	LastRedirectDn_uri                      string    //VARCHAR(255)
	MobileCallingPartyNumber                string    //VARCHAR(50)
	FinalMobileCalledPartyNumber            string    //VARCHAR(50)
	OrigMobileDeviceName                    string    //VARCHAR(129)
	DestMobileDeviceName                    string    //VARCHAR(129)
	OrigMobileCallDuration                  int       //INTEGER
	DestMobileCallDuration                  int       //INTEGER
	MobileCallType                          int       //INTEGER
	OriginalCalledPartyPattern              string    //VARCHAR(50)
	FinalCalledPartyPattern                 string    //VARCHAR(50)
	LastRedirectingPartyPattern             string    //VARCHAR(50)
	HuntPilotPattern                        string    //VARCHAR(50)
	OrigDeviceType                          string    //VARCHAR(100)
	DestDeviceType                          string    //VARCHAR(100)
}

func CompareHeader(h []string) bool {
	if len(h) != len(CdrHeader) {
		return false
	}

	for i, v := range h {
		if v != CdrHeader[i] {
			return false
		}
	}
	return true
}

func NewRecordFromStrSlice(rec []string) (*Record, error) {
	if len(rec) != RecordLen {
		return nil, errors.New("wrong slice")
	}
	r := Record{}
	var err error
	r.CdrRecordType, err = strconv.Atoi(rec[CdrRecordType])
	if err != nil {
		return nil, err
	}
	if r.CdrRecordType != 1 {
		return nil, errors.New("wrong record type")
	}
	r.GlobalCallID_callManagerId, _ = strconv.Atoi(rec[GlobalCallID_callManagerId])
	r.GlobalCallID_callId, _ = strconv.Atoi(rec[GlobalCallID_callId])
	r.OrigLegCallIdentifier, _ = strconv.Atoi(rec[OrigLegCallIdentifier])
	var t1, _ = strconv.Atoi(rec[DateTimeOrigination])
	r.DateTimeOrigination = time.Unix(int64(t1), 0)
	r.OrigNodeId, _ = strconv.Atoi(rec[OrigNodeId])
	r.OrigSpan, _ = strconv.Atoi(rec[OrigSpan])
	r.OrigIpAddr, _ = strconv.Atoi(rec[OrigIpAddr])
	r.CallingPartyNumber = rec[CallingPartyNumber]
	r.CallingPartyUnicodeLoginUserID = rec[CallingPartyUnicodeLoginUserID]
	r.OrigCause_location, _ = strconv.Atoi(rec[OrigCause_location])
	r.OrigCause_value, _ = strconv.Atoi(rec[OrigCause_value])
	r.OrigPrecedenceLevel, _ = strconv.Atoi(rec[OrigPrecedenceLevel])
	r.OrigMediaTransportAddress_IP, _ = strconv.Atoi(rec[OrigMediaTransportAddress_IP])
	r.OrigMediaTransportAddress_Port, _ = strconv.Atoi(rec[OrigMediaTransportAddress_Port])
	r.OrigMediaCap_payloadCapability, _ = strconv.Atoi(rec[OrigMediaCap_payloadCapability])
	r.OrigMediaCap_maxFramesPerPacket, _ = strconv.Atoi(rec[OrigMediaCap_maxFramesPerPacket])
	r.OrigMediaCap_g723BitRate, _ = strconv.Atoi(rec[OrigMediaCap_g723BitRate])
	r.OrigVideoCap_Codec, _ = strconv.Atoi(rec[OrigVideoCap_Codec])
	r.OrigVideoCap_Bandwidth, _ = strconv.Atoi(rec[OrigVideoCap_Bandwidth])
	r.OrigVideoCap_Resolution, _ = strconv.Atoi(rec[OrigVideoCap_Resolution])
	r.OrigVideoTransportAddress_IP, _ = strconv.Atoi(rec[OrigVideoTransportAddress_IP])
	r.OrigVideoTransportAddress_Port, _ = strconv.Atoi(rec[OrigVideoTransportAddress_Port])
	r.OrigRSVPAudioStat = rec[OrigRSVPAudioStat]
	r.OrigRSVPVideoStat = rec[OrigRSVPVideoStat]
	r.DestLegIdentifier, _ = strconv.Atoi(rec[DestLegIdentifier])
	r.DestNodeId, _ = strconv.Atoi(rec[DestNodeId])
	r.DestSpan, _ = strconv.Atoi(rec[DestSpan])
	r.DestIpAddr, _ = strconv.Atoi(rec[DestIpAddr])
	r.OriginalCalledPartyNumber = rec[OriginalCalledPartyNumber]
	r.FinalCalledPartyNumber = rec[FinalCalledPartyNumber]
	r.FinalCalledPartyUnicodeLoginUserID = rec[FinalCalledPartyUnicodeLoginUserID]
	r.DestCause_location, _ = strconv.Atoi(rec[DestCause_location])
	r.DestCause_value, _ = strconv.Atoi(rec[DestCause_value])
	r.DestPrecedenceLevel, _ = strconv.Atoi(rec[DestPrecedenceLevel])
	r.DestMediaTransportAddress_IP, _ = strconv.Atoi(rec[DestMediaTransportAddress_IP])
	r.DestMediaTransportAddress_Port, _ = strconv.Atoi(rec[DestMediaTransportAddress_Port])
	r.DestMediaCap_payloadCapability, _ = strconv.Atoi(rec[DestMediaCap_payloadCapability])
	r.DestMediaCap_maxFramesPerPacket, _ = strconv.Atoi(rec[DestMediaCap_maxFramesPerPacket])
	r.DestMediaCap_g723BitRate, _ = strconv.Atoi(rec[DestMediaCap_g723BitRate])
	r.DestVideoCap_Codec, _ = strconv.Atoi(rec[DestVideoCap_Codec])
	r.DestVideoCap_Bandwidth, _ = strconv.Atoi(rec[DestVideoCap_Bandwidth])
	r.DestVideoCap_Resolution, _ = strconv.Atoi(rec[DestVideoCap_Resolution])
	r.DestVideoTransportAddress_IP, _ = strconv.Atoi(rec[DestVideoTransportAddress_IP])
	r.DestVideoTransportAddress_Port, _ = strconv.Atoi(rec[DestVideoTransportAddress_Port])
	r.DestRSVPAudioStat = rec[DestRSVPAudioStat]
	r.DestRSVPVideoStat = rec[DestRSVPVideoStat]
	t1, _ = strconv.Atoi(rec[DateTimeConnect])
	r.DateTimeConnect = time.Unix(int64(t1), 0)
	t1, _ = strconv.Atoi(rec[DateTimeDisconnect])
	r.DateTimeDisconnect = time.Unix(int64(t1), 0)
	r.LastRedirectDn = rec[LastRedirectDn]
	r.Pkid, _ = uuid.Parse(rec[Pkid])
	r.OriginalCalledPartyNumberPartition = rec[OriginalCalledPartyNumberPartition]
	r.CallingPartyNumberPartition = rec[CallingPartyNumberPartition]
	r.FinalCalledPartyNumberPartition = rec[FinalCalledPartyNumberPartition]
	r.LastRedirectDnPartition = rec[LastRedirectDnPartition]
	r.Duration, _ = strconv.Atoi(rec[Duration])
	r.OrigDeviceName = rec[OrigDeviceName]
	r.DestDeviceName = rec[DestDeviceName]
	r.OrigCallTerminationOnBehalfOf, _ = strconv.Atoi(rec[OrigCallTerminationOnBehalfOf])
	r.DestCallTerminationOnBehalfOf, _ = strconv.Atoi(rec[DestCallTerminationOnBehalfOf])
	r.OrigCalledPartyRedirectOnBehalfOf, _ = strconv.Atoi(rec[OrigCalledPartyRedirectOnBehalfOf])
	r.LastRedirectRedirectOnBehalfOf, _ = strconv.Atoi(rec[LastRedirectRedirectOnBehalfOf])
	r.OrigCalledPartyRedirectReason, _ = strconv.Atoi(rec[OrigCalledPartyRedirectReason])
	r.LastRedirectRedirectReason, _ = strconv.Atoi(rec[LastRedirectRedirectReason])
	r.DestConversationId, _ = strconv.Atoi(rec[DestConversationId])
	r.GlobalCallId_ClusterID = rec[GlobalCallId_ClusterID]
	r.JoinOnBehalfOf, _ = strconv.Atoi(rec[JoinOnBehalfOf])
	r.Comment = rec[Comment]
	r.AuthCodeDescription = rec[AuthCodeDescription]
	r.AuthorizationLevel, _ = strconv.Atoi(rec[AuthorizationLevel])
	r.ClientMatterCode = rec[ClientMatterCode]
	r.OrigDTMFMethod, _ = strconv.Atoi(rec[OrigDTMFMethod])
	r.DestDTMFMethod, _ = strconv.Atoi(rec[DestDTMFMethod])
	r.CallSecuredStatus, _ = strconv.Atoi(rec[CallSecuredStatus])
	r.OrigConversationId, _ = strconv.Atoi(rec[OrigConversationId])
	r.OrigMediaCap_Bandwidth, _ = strconv.Atoi(rec[OrigMediaCap_Bandwidth])
	r.DestMediaCap_Bandwidth, _ = strconv.Atoi(rec[DestMediaCap_Bandwidth])
	r.AuthorizationCodeValue = rec[AuthorizationCodeValue]
	r.OutpulsedCallingPartyNumber = rec[OutpulsedCallingPartyNumber]
	r.OutpulsedCalledPartyNumber = rec[OutpulsedCalledPartyNumber]
	r.OrigIpv4v6Addr = rec[OrigIpv4v6Addr]
	r.DestIpv4v6Addr = rec[DestIpv4v6Addr]
	r.OrigVideoCap_Codec_Channel2, _ = strconv.Atoi(rec[OrigVideoCap_Codec_Channel2])
	r.OrigVideoCap_Bandwidth_Channel2, _ = strconv.Atoi(rec[OrigVideoCap_Bandwidth_Channel2])
	r.OrigVideoCap_Resolution_Channel2, _ = strconv.Atoi(rec[OrigVideoCap_Resolution_Channel2])
	r.OrigVideoTransportAddress_IP_Channel2, _ = strconv.Atoi(rec[OrigVideoTransportAddress_IP_Channel2])
	r.OrigVideoTransportAddress_Port_Channel2, _ = strconv.Atoi(rec[OrigVideoTransportAddress_Port_Channel2])
	r.OrigVideoChannel_Role_Channel2, _ = strconv.Atoi(rec[OrigVideoChannel_Role_Channel2])
	r.DestVideoCap_Codec_Channel2, _ = strconv.Atoi(rec[DestVideoCap_Codec_Channel2])
	r.DestVideoCap_Bandwidth_Channel2, _ = strconv.Atoi(rec[DestVideoCap_Bandwidth_Channel2])
	r.DestVideoCap_Resolution_Channel2, _ = strconv.Atoi(rec[DestVideoCap_Resolution_Channel2])
	r.DestVideoTransportAddress_IP_Channel2, _ = strconv.Atoi(rec[DestVideoTransportAddress_IP_Channel2])
	r.DestVideoTransportAddress_Port_Channel2, _ = strconv.Atoi(rec[DestVideoTransportAddress_Port_Channel2])
	r.DestVideoChannel_Role_Channel2, _ = strconv.Atoi(rec[DestVideoChannel_Role_Channel2])
	r.IncomingProtocolID, _ = strconv.Atoi(rec[IncomingProtocolID])
	r.IncomingProtocolCallRef = rec[IncomingProtocolCallRef]
	r.OutgoingProtocolID, _ = strconv.Atoi(rec[OutgoingProtocolID])
	r.OutgoingProtocolCallRef = rec[OutgoingProtocolCallRef]
	r.CurrentRoutingReason, _ = strconv.Atoi(rec[CurrentRoutingReason])
	r.OrigRoutingReason, _ = strconv.Atoi(rec[OrigRoutingReason])
	r.LastRedirectingRoutingReason, _ = strconv.Atoi(rec[LastRedirectingRoutingReason])
	r.HuntPilotPartition = rec[HuntPilotPartition]
	r.HuntPilotDN = rec[HuntPilotDN]
	r.CalledPartyPatternUsage, _ = strconv.Atoi(rec[CalledPartyPatternUsage])
	r.IncomingICID = rec[IncomingICID]
	r.IncomingOrigIOI = rec[IncomingOrigIOI]
	r.IncomingTermIOI = rec[IncomingTermIOI]
	r.OutgoingICID = rec[OutgoingICID]
	r.OutgoingOrigIOI = rec[OutgoingOrigIOI]
	r.OutgoingTermIOI = rec[OutgoingTermIOI]
	r.OutpulsedOriginalCalledPartyNumber = rec[OutpulsedOriginalCalledPartyNumber]
	r.OutpulsedLastRedirectingNumber = rec[OutpulsedLastRedirectingNumber]
	r.WasCallQueued, _ = strconv.Atoi(rec[WasCallQueued])
	r.TotalWaitTimeInQueue, _ = strconv.Atoi(rec[TotalWaitTimeInQueue])
	r.CallingPartyNumber_uri = rec[CallingPartyNumber_uri]
	r.OriginalCalledPartyNumber_uri = rec[OriginalCalledPartyNumber_uri]
	r.FinalCalledPartyNumber_uri = rec[FinalCalledPartyNumber_uri]
	r.LastRedirectDn_uri = rec[LastRedirectDn_uri]
	r.MobileCallingPartyNumber = rec[MobileCallingPartyNumber]
	r.FinalMobileCalledPartyNumber = rec[FinalMobileCalledPartyNumber]
	r.OrigMobileDeviceName = rec[OrigMobileDeviceName]
	r.DestMobileDeviceName = rec[DestMobileDeviceName]
	r.OrigMobileCallDuration, _ = strconv.Atoi(rec[OrigMobileCallDuration])
	r.DestMobileCallDuration, _ = strconv.Atoi(rec[DestMobileCallDuration])
	r.MobileCallType, _ = strconv.Atoi(rec[MobileCallType])
	r.OriginalCalledPartyPattern = rec[OriginalCalledPartyPattern]
	r.FinalCalledPartyPattern = rec[FinalCalledPartyPattern]
	r.LastRedirectingPartyPattern = rec[LastRedirectingPartyPattern]
	r.HuntPilotPattern = rec[HuntPilotPattern]
	r.OrigDeviceType = rec[OrigDeviceType]
	r.DestDeviceType = rec[DestDeviceType]

	return &r, nil
}
