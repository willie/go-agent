// Agent configuration package.

// This package includes both compile-time and run-time configuration of the
// agent. Variables are made configurable at run-time when necessary for users.

package config

import (
	"net"
	"net/http"
	"time"

	"github.com/sqreen/go-agent/agent/plog"

	"github.com/spf13/viper"
)

type HTTPAPIEndpoint struct {
	Method, URL string
}

const (
	// Default value of network timeouts.
	DefaultNetworkTimeout = 5 * time.Second
)

// Backend client configuration.
var (
	// Timeout value of a HTTP request. See http.Client.Timeout.
	BackendHTTPAPIRequestTimeout = DefaultNetworkTimeout

	// List of endpoint addresses, relative to the base URL.
	BackendHTTPAPIEndpoint = struct {
		AppLogin, AppLogout, AppBeat, Batch HTTPAPIEndpoint
	}{
		AppLogin:  HTTPAPIEndpoint{http.MethodPost, "/sqreen/v1/app-login"},
		AppLogout: HTTPAPIEndpoint{http.MethodGet, "/sqreen/v0/app-logout"},
		AppBeat:   HTTPAPIEndpoint{http.MethodPost, "/sqreen/v1/app-beat"},
		Batch:     HTTPAPIEndpoint{http.MethodPost, "/sqreen/v0/batch"},
	}

	// Header name of the API token.
	BackendHTTPAPIHeaderToken = "X-Api-Key"

	// Header name of the API session.
	BackendHTTPAPIHeaderSession = "X-Session-Key"

	// BackendHTTPAPIRequestRetryPeriod is the time period to retry failed backend
	// HTTP requests.
	BackendHTTPAPIRequestRetryPeriod = time.Minute

	// BackendHTTPAPIBackoffRate is the backoff rate to compute the next sleep
	// duration before retrying the failed request.
	BackendHTTPAPIBackoffRate = 2.0

	// BackendHTTPAPIBackoffMaxDuration is the maximum backoff's sleep duration.
	BackendHTTPAPIBackoffMaxDuration = 30 * time.Minute

	// BackendHTTPAPIBackoffMaxDuration is the minimum backoff's sleep duration.
	BackendHTTPAPIBackoffMinDuration = time.Millisecond

	// BackendHTTPAPIDefaultHeartbeatDelay is the default heartbeat delay when not
	// correctly provided by the backend.
	BackendHTTPAPIDefaultHeartbeatDelay = time.Minute
)

const (
	MaxEventsPerHeatbeat = 1000
)

var (
	TrackedHTTPHeaders = []string{
		"X-Forwarded-For",
		"X-Forwarded-Host",
		"X-Forwarded-Proto",
		"X-Client-Ip",
		"X-Real-Ip",
		"X-Forwarded",
		"X-Cluster-Client-Ip",
		"Forwarded-For",
		"Forwarded",
		"Via",
		"User-Agent",
		"Content-Type",
		"Content-Length",
		"Host",
		"X-Requested-With",
		"X-Request-Id",
		"HTTP_X_FORWARDED_FOR",
		"HTTP_X_REAL_IP",
		"HTTP_CLIENT_IP",
		"HTTP_X_FORWARDED",
		"HTTP_X_CLUSTER_CLIENT_IP",
		"HTTP_FORWARDED_FOR",
		"HTTP_FORWARDED",
		"HTTP_VIA",
	}

	IPRelatedHTTPHeaders = []string{
		"X-Forwarded-For",
		"X-Client-Ip",
		"X-Real-Ip",
		"X-Forwarded",
		"X-Cluster-Client-Ip",
		"Forwarded-For",
		"Forwarded",
		"Via",
		"HTTP_X_FORWARDED_FOR",
		"HTTP_X_REAL_IP",
		"HTTP_CLIENT_IP",
		"HTTP_X_FORWARDED",
		"HTTP_X_CLUSTER_CLIENT_IP",
		"HTTP_FORWARDED_FOR",
		"HTTP_FORWARDED",
		"HTTP_VIA",
	}
)

// Helper function to return the IP network out of a string.
func ipnet(s string) *net.IPNet {
	_, n, _ := net.ParseCIDR(s)
	return n
}

// IP networks allowing to compute whether to
var (
	IPPrivateNetworks = []*net.IPNet{
		ipnet("0.0.0.0/8"),
		ipnet("10.0.0.0/8"),
		ipnet("127.0.0.0/8"),
		ipnet("169.254.0.0/16"),
		ipnet("172.16.0.0/12"),
		ipnet("192.0.0.0/29"),
		ipnet("192.0.0.170/31"),
		ipnet("192.0.2.0/24"),
		ipnet("192.168.0.0/16"),
		ipnet("198.18.0.0/15"),
		ipnet("198.51.100.0/24"),
		ipnet("203.0.113.0/24"),
		ipnet("240.0.0.0/4"),
		ipnet("255.255.255.255/32"),
		ipnet("::1/128"),
		ipnet("::/128"),
		ipnet("::ffff:0:0/96"),
		ipnet("100::/64"),
		ipnet("2001::/23"),
		ipnet("2001:2::/48"),
		ipnet("2001:db8::/32"),
		ipnet("2001:10::/28"),
		ipnet("fc00::/7"),
		ipnet("fe80::/10"),
	}

	IPv4PublicNetwork = ipnet("100.64.0.0/10")
)

const (
	configEnvPrefix    = `sqreen`
	configFileBasename = `sqreen`
	configFilePath     = `.`
)

const (
	configKeyBackendHTTPAPIBaseURL = `url`
	configKeyBackendHTTPAPIToken   = `token`
	configKeyLogLevel              = `log_level`
	configKeyAppName               = `app_name`
	configKeyHTTPClientIPHeader    = `ip_header`
	configKeyDisable               = `disable`
)

// User configuration's default values.
const (
	configDefaultBackendHTTPAPIBaseURL = `https://back.sqreen.com`
	configDefaultLogLevel              = `warn`
)

func init() {
	viper.SetEnvPrefix(configEnvPrefix)
	viper.AutomaticEnv()
	viper.SetConfigName(configFileBasename)
	viper.AddConfigPath(configFilePath)

	viper.SetDefault(configKeyBackendHTTPAPIBaseURL, configDefaultBackendHTTPAPIBaseURL)
	viper.SetDefault(configKeyLogLevel, configDefaultLogLevel)
	viper.SetDefault(configKeyAppName, "")
	viper.SetDefault(configKeyHTTPClientIPHeader, "")
	viper.SetDefault(configKeyDisable, "")

	logger := plog.NewLogger("sqreen/agent/config")

	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("configuration file read error:", err)
	}
}

// BackendHTTPAPIBaseURL returns the base URL of the backend HTTP API.
func BackendHTTPAPIBaseURL() string {
	return viper.GetString(configKeyBackendHTTPAPIBaseURL)
}

// BackendHTTPAPIToken returns the access token to the backend API.
func BackendHTTPAPIToken() string {
	return viper.GetString(configKeyBackendHTTPAPIToken)
}

// LogLevel returns the default log level.
func LogLevel() string {
	return viper.GetString(configKeyLogLevel)
}

// AppName returns the default log level.
func AppName() string {
	return viper.GetString(configKeyAppName)
}

// HTTPClientIPHeader IPHeader returns the header to first lookup to find the client ip of a HTTP request.
func HTTPClientIPHeader() string {
	return viper.GetString(configKeyHTTPClientIPHeader)
}

func Disable() bool {
	disable := viper.GetString(configKeyDisable)
	return disable != "" || BackendHTTPAPIToken() == ""
}
