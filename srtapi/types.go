package srtapi

// #cgo LDFLAGS: -lsrt
// #include <srt/srt.h>
import "C"
import (
	"syscall"
	"unsafe"
)

type _Socklen C.int

var rsa syscall.RawSockaddrAny
var rs4 syscall.RawSockaddrInet4
var rs6 syscall.RawSockaddrInet6

// Size of raw sock addr structures
const (
	SizeofSockaddrAny   = _Socklen(unsafe.Sizeof(rsa))
	SizeofSockaddrInet4 = _Socklen(unsafe.Sizeof(rs4))
	SizeofSockaddrInet6 = _Socklen(unsafe.Sizeof(rs6))
)

// SRT socket status
const (
	StatusInit       = C.SRTS_INIT
	StatusOpened     = C.SRTS_OPENED
	StatusListening  = C.SRTS_LISTENING
	StatusConnecting = C.SRTS_CONNECTING
	StatusConnected  = C.SRTS_CONNECTED
	StatusBroken     = C.SRTS_BROKEN
	StatusClosing    = C.SRTS_CLOSING
	StatusClosed     = C.SRTS_CLOSED
	StatusNonexist   = C.SRTS_NONEXIST
)

// SRT socket options
const (
	OptionMss            = C.SRTO_MSS
	OptionSndsyn         = C.SRTO_SNDSYN
	OptionRcvsyn         = C.SRTO_RCVSYN
	OptionIsn            = C.SRTO_ISN
	OptionFc             = C.SRTO_FC
	OptionSndbuf         = C.SRTO_SNDBUF
	OptionRcvbuf         = C.SRTO_RCVBUF
	OptionLinger         = C.SRTO_LINGER
	OptionUDPSndbuf      = C.SRTO_UDP_SNDBUF
	OptionUDPRcvbuf      = C.SRTO_UDP_RCVBUF
	OptionRendezvous     = C.SRTO_RENDEZVOUS
	OptionSndtimeo       = C.SRTO_SNDTIMEO
	OptionRcvtimeo       = C.SRTO_RCVTIMEO
	OptionReuseaddr      = C.SRTO_REUSEADDR
	OptionMaxbw          = C.SRTO_MAXBW
	OptionState          = C.SRTO_STATE
	OptionEvent          = C.SRTO_EVENT
	OptionSnddata        = C.SRTO_SNDDATA
	OptionRcvdata        = C.SRTO_RCVDATA
	OptionSender         = C.SRTO_SENDER
	OptionTsbpdmode      = C.SRTO_TSBPDMODE
	OptionLatency        = C.SRTO_LATENCY
	OptionTsbpddelay     = C.SRTO_TSBPDDELAY
	OptionInputbw        = C.SRTO_INPUTBW
	OptionOheadbw        = C.SRTO_OHEADBW
	OptionPassphrase     = C.SRTO_PASSPHRASE
	OptionPbkeylen       = C.SRTO_PBKEYLEN
	OptionKmstate        = C.SRTO_KMSTATE
	OptionIpttl          = C.SRTO_IPTTL
	OptionIptos          = C.SRTO_IPTOS
	OptionTlpktdrop      = C.SRTO_TLPKTDROP
	OptionNakreport      = C.SRTO_NAKREPORT
	OptionVersion        = C.SRTO_VERSION
	OptionPeerversion    = C.SRTO_PEERVERSION
	OptionConntimeo      = C.SRTO_CONNTIMEO
	OptionSndpeerkmstate = C.SRTO_SNDPEERKMSTATE
	OptionRcvkmstate     = C.SRTO_RCVKMSTATE
	OptionLossmaxttl     = C.SRTO_LOSSMAXTTL
	OptionRcvlatency     = C.SRTO_RCVLATENCY
	OptionPeerlatency    = C.SRTO_PEERLATENCY
	OptionMinversion     = C.SRTO_MINVERSION
	OptionStreamid       = C.SRTO_STREAMID
	OptionSmoother       = C.SRTO_SMOOTHER
	OptionMessageapi     = C.SRTO_MESSAGEAPI
	OptionPayloadsize    = C.SRTO_PAYLOADSIZE
	OptionTranstype      = C.SRTO_TRANSTYPE
)

// SRT trans type
const (
	TypeLive    = C.SRTT_LIVE
	TypeFile    = C.SRTT_FILE
	TypeInvalid = C.SRTT_INVALID
)

// SRT errno
const (
	EUNKNOWN        = C.SRT_EUNKNOWN
	SUCCESS         = C.SRT_SUCCESS
	ECONNSETUP      = C.SRT_ECONNSETUP
	ENOSERVER       = C.SRT_ENOSERVER
	ECONNREJ        = C.SRT_ECONNREJ
	ESOCKFAIL       = C.SRT_ESOCKFAIL
	ESECFAIL        = C.SRT_ESECFAIL
	ECONNFAIL       = C.SRT_ECONNFAIL
	ECONNLOST       = C.SRT_ECONNLOST
	ENOCONN         = C.SRT_ENOCONN
	ERESOURCE       = C.SRT_ERESOURCE
	ETHREAD         = C.SRT_ETHREAD
	ENOBUF          = C.SRT_ENOBUF
	EFILE           = C.SRT_EFILE
	EINVRDOFF       = C.SRT_EINVRDOFF
	ERDPERM         = C.SRT_ERDPERM
	EINVWROFF       = C.SRT_EINVWROFF
	EWRPERM         = C.SRT_EWRPERM
	EINVOP          = C.SRT_EINVOP
	EBOUNDSOCK      = C.SRT_EBOUNDSOCK
	ECONNSOCK       = C.SRT_ECONNSOCK
	EINVPARAM       = C.SRT_EINVPARAM
	EINVSOCK        = C.SRT_EINVSOCK
	EUNBOUNDSOCK    = C.SRT_EUNBOUNDSOCK
	ENOLISTEN       = C.SRT_ENOLISTEN
	ERDVNOSERV      = C.SRT_ERDVNOSERV
	ERDVUNBOUND     = C.SRT_ERDVUNBOUND
	EINVALMSGAPI    = C.SRT_EINVALMSGAPI
	EINVALBUFFERAPI = C.SRT_EINVALBUFFERAPI
	EDUPLISTEN      = C.SRT_EDUPLISTEN
	ELARGEMSG       = C.SRT_ELARGEMSG
	EINVPOLLID      = C.SRT_EINVPOLLID
	EASYNCFAIL      = C.SRT_EASYNCFAIL
	EASYNCSND       = C.SRT_EASYNCSND
	EASYNCRCV       = C.SRT_EASYNCRCV
	ETIMEOUT        = C.SRT_ETIMEOUT
	ECONGEST        = C.SRT_ECONGEST
	EPEERERR        = C.SRT_EPEERERR
)

// SRT log level
const (
	LogFatal   = 2
	LogError   = 3
	LogWarning = 4
	LogNote    = 5
	LogDebug   = 7
)

// SRT log FA
const (
	LogFAGeneral  = 0
	LogFABstats   = 1
	LogFAControl  = 2
	LogFAData     = 3
	LogFATsbpd    = 4
	LogFARexmit   = 5
	LogFALastnone = 31
)

// SRT log flags
const (
	LogFlagDisableTime       = 1
	LogFlagDisableThreadname = 2
	LogFlagDisableSeverity   = 4
	LogFlagDisableEOF        = 8
)

// SRT epoll opt
const (
	EpollIn  = C.SRT_EPOLL_IN
	EpollOut = C.SRT_EPOLL_OUT
	EpollErr = C.SRT_EPOLL_ERR
)

// SRT const
const (
	InvalidSock = -1
	APIError    = -1
)