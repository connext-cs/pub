package config

import (
	"errors"
	"fmt"
	"time"
)

//一些系统变量
const (
	dZipkinAddr                 = ""
	dConsulAddr                 = ""
	dFileServerStaticPath       = ""
	dLocalIP                    = ""
	dFileserverIP               = ""
	dFileserverPort             = 0
	dCIDir                      = ""
	dLocalSSHPort               = 0
	dLocalSSHUser               = ""
	dLocalSSHPass               = ""
	dRedisHost                  = ""
	dRedisPort                  = 0
	dNatsAddr                   = ""
	dOsExitSignal               = -1
	dConfigPrefix               = ""
	dRegisterTtL                = 0
	dRegisterInterval           = 0
	dManagementTimeOutSecond    = 0
	dClientRequestTimeOutSecond = 0
	dClientRequestPoolNum       = 0
	dClientDialTimeOutSecond    = 0

	dCloudMysqlHost     = ""
	dCloudMysqlPort     = 0
	dCloudMysqlUserName = ""
	dCloudMysqlPasswd   = ""
	dCloudMysqlDatabase = ""

	dMysqlHost             = ""
	dMysqlPort             = 0
	dMysqlUserName         = ""
	dMysqlPasswd           = ""
	dMysqlDatabase         = ""
	dMysqlMaxOpen          = 0
	dMysqlMaxIdle          = 0
	dMaxLifeTime           = 0
	dExecLogLocationPrefix = ""

	dSftpAddr     = ""
	dSftpPort     = 0
	dSftpUserName = ""
	dSftpPassword = ""
	dSftpPath     = ""

	// dUserServiceAddr   = ""
	// dUserServicePrefix = ""

	// dCloudServiceAddr   = ""
	// dVmServicePrefix    = ""
	dGrafanaPrefix      = ""
	dGraylogPrefix      = ""
	dAlertManagerPrefix = ""
	dFileServerPrefix   = ""
	dFileServiceAddr    = ""
)

const (
	dDefaultHost = ""
	dAppName     = ""

	dLogServiceHost = dDefaultHost
	dLogServicePort = 0

	dLdapHost     = ""
	dLdapPort     = 0
	dLdapUser     = ""
	dLdapPassword = ""
	dLdapSearchdn = ""

	dK8sUser     = ""
	dK8sPassword = ""
	dMasterip    = ""

	dELKHost                 = ""
	dELKPort                 = 0
	dELKUser                 = ""
	dELKPassword             = ""
	dELKIndexMaxResultWindow = 0

	dClusterServiceControllerURL = ""
	dClusterExternalNetworkPort  = 0
	dClusterHarbor               = ""

	dVMEtcdHost = ""
	dVMEtcdPort = 0

	dVMJenkinsHost     = ""
	dVMJenkinsPort     = 0
	dVMJenkinsUser     = ""
	dVMJenkinsPassword = ""

	dVMSonarHost     = ""
	dVMSonarPort     = 0
	dVMSonarUser     = ""
	dVMSonarPassword = ""

	dVMSonarEMailHost         = ""
	dVMSonarEMailPort         = 0
	dVMSonarEMailUser         = ""
	dVMSonarEMailPassword     = ""
	dVMSonarEMailDefaultTitle = ""

	dSonarAccessHost = ""
	dSonarAccessPort = 0
	dSonarUser       = ""
	dSonarPassword   = ""
	dwarURL          = ""

	dConsulHost = ""
	dConsulPort = 0

	//dAgentInstallCMD           string = ""
	dLinuxAgentInstallCMD        string = ""
	dLinuxAgentDownloadURL       string = ""
	dWindowsAgentDownloadURL     string = ""
	dWindowsAgentDownloadURLShow string = ""
	dWindowsAgentInstallCMD      string = ""
	dLinuxAgentLogPath           string = ""
	dWindowsAgentLogPath         string = ""
	dLinuxAgentPath              string = ""
	dWindowsAgentPath            string = ""
	dSuperPassword               string = ""
)

//一些错误变量
const (
	dNoneRowsAffect      = ""
	ZeroRows             = 0
	ZeroId               = 0
	HostInfoLogTable     = ""
	HostInfoTable        = ""
	SoftwareInfoTable    = ""
	SoftwareInfoLogTable = ""
	SoftwareConfigTable  = ""

	TClusterNodeTable          = ""
	TClusterNodeGroupTable     = ""
	TClusterNodeGroupLinkTable = ""
	TClusterNodeJobTable       = ""

	ManageTable            = ""
	ManageJobListTable     = ""
	ManageTableItemsColoum = ""
	UnkownColoumIndex      = 0
	Deleted                = 0
	NoDeleted              = 0
	EmptyStr               = ""
)
const KIBANA_ADDR = ""

func CDefaultHost() string {
	return cStr("default_host", dDefaultHost)
}

func CAppName() string {
	return cStr("app_name", dAppName)
}

func CLogServiceHost() string {
	return cStr("log_service_Host", dLogServiceHost)
}

func CLogServicePort() int {
	return cInt("log_service_port", dLogServicePort)
}

func CLdapHost() string {
	return cStr("ldap_host", dLdapHost)
}

func CLdapPort() int {
	return cInt("ldap_port", dLdapPort)
}

func CLdapUser() string {
	return cStr("ldap_user", dLdapUser)
}

func CLdapPassword() string {
	return cStr("ldap_password", dLdapPassword)
}

func CLdapSearchdn() string {
	return cStr("ldap_searchdn", dLdapSearchdn)
}
func CK8sUser() string {
	return cStr("k8s_user", dK8sUser)
}

func CK8sMasterip() string {
	return cStr("k8s_masterip", dMasterip)
}

func CK8sPassword() string {
	return cStr("k8s_password", dK8sPassword)
}

func CELKHost() string {
	return cStr("elk_Host", dELKHost)
}

func CRedisHost() string {
	return cStr("redis_host", dRedisHost)
}

func CRedisPort() int {
	return cInt("redis_port", dRedisPort)
}

func CELKPort() int {
	return cInt("elk_port", dELKPort)
}

func CELKUser() string {
	return cStr("elk_user", dELKUser)
}

func CELKPassword() string {
	return cStr("elk_password", dELKPassword)
}

func CELKIndexMaxResultWindow() int {
	return cInt("elk_index_max_result_window", dELKIndexMaxResultWindow)
}

func CZipkinAddr() string {
	return cStr("zipkin_addr", dZipkinAddr)
}

func CConsulAddr() string {
	return cStr("consul_addr", dConsulAddr)
}

func CLocalIP() string {
	return cStr("local_ip", dLocalIP)
}

func CFileserverIP() string {
	return cStr("fileserver_ip", dFileserverIP)
}

func CFileserverPort() int {
	return cInt("fileserver_port", dFileserverPort)
}

func CCIDir() string {
	return cStr("ci_dir", dCIDir)
}
func CLocalSSHPort() int {
	return cInt("local_ssh_port", dLocalSSHPort)
}

func CFileServerStaticPath() string {
	return cStr("file_server_static_path", dFileServerStaticPath)
}
func CLocalSSHUser() string {
	return cStr("local_ssh_user", dLocalSSHUser)
}

func CLocalSSHPass() string {
	return cStr("local_ssh_pass", dLocalSSHPass)
}

func CNatsAddr() string {
	return cStr("nats_addr", dNatsAddr)
}

func COsExitSignal() int {
	return cInt("os_exit_signal", dOsExitSignal)
}

func CConfigPrefix() string {
	return cStr("config_prefix", dConfigPrefix)
}

func CCloudMysqlHost() string {
	return cStr("cloud_mysql_host", dCloudMysqlHost)
}

func CCloudMysqlPort() int {
	return cInt("cloud_mysql_port", dCloudMysqlPort)
}

func CCloudMysqlUserName() string {
	return cStr("cloud_mysql_user_name", dCloudMysqlUserName)
}

func CCloudMysqlPasswd() string {
	return cStr("cloud_mysql_passwd", dCloudMysqlPasswd)
}

func CCloudMysqlDatabase() string {
	return cStr("cloud_mysql_database", dCloudMysqlDatabase)
}

func CMysqlHost() string {
	return cStr("mysql_host", dMysqlHost)
}

func CMysqlPort() int {
	return cInt("mysql_port", dMysqlPort)
}

func CMysqlUserName() string {
	return cStr("mysql_user_name", dMysqlUserName)
}

func CMysqlPasswd() string {
	return cStr("mysql_passwd", dMysqlPasswd)
}

func CMysqlDatabase() string {
	return cStr("mysql_database", dMysqlDatabase)
}

func CMysqlMaxOpen() int {
	return cInt("mysql_max_open", dMysqlMaxOpen)
}

func CMysqlMaxIdle() int {
	return cInt("mysql_max_idle", dMysqlMaxIdle)
}

func CMysqlMaxLifeTime() time.Duration {
	return time.Duration(cInt("mysql_max_life_time", dMaxLifeTime)) * time.Second
}

func CExecLogLocationPrefix() string {
	return cStr("exec_log_location_prefix", dExecLogLocationPrefix)
}

func CSftpAddr() string {
	return cStr("sftp_addr", dSftpAddr)
}

func CSftpPort() int {
	return cInt("sftp_port", dSftpPort)
}

func CSftpUserName() string {
	return cStr("sftp_name", dSftpUserName)
}

func CSftpPassword() string {
	return cStr("sftp_password", dSftpPassword)
}

func CSftpPrefixPath() string {
	return cStr("sftp_path", dSftpPath)
}

func CManagementTimeOutSecond() int {
	return cInt("management_time_out_second", dManagementTimeOutSecond)
}

func CClientRequestTimeOutSecond() time.Duration {
	return time.Duration(cInt("client_request_time_out_second", dClientRequestTimeOutSecond)) * time.Second
}

func CClientRequestPoolNum() int {
	return cInt("client_request_pool_num", dClientRequestPoolNum)
}

func CClientDialTimeOutSecond() time.Duration {
	return time.Duration(cInt("client_dial_time_out_second", dClientDialTimeOutSecond))
}

// func CUserServiceAddr() *url.URL {
// 	if addrUrl, err := url.Parse(cStr("user_service_addr", dUserServiceAddr)); err != nil {
// 		return nil
// 	} else {
// 		return addrUrl
// 	}
// }

// func CCloudServiceAddr() *url.URL {
// 	if addrUrl, err := url.Parse(cStr("cloud_service_addr", dCloudServiceAddr)); err != nil {
// 		return nil
// 	} else {
// 		return addrUrl
// 	}
// }

// func CVmServicePrefix() string {
// 	return cStr("vm_service_prefix", dVmServicePrefix)
// }

// func CUserServicePrefix() string {
// 	return cStr("user_service_prefix", dUserServicePrefix)
// }

func CGrafanaPrefix() string {
	return cStr("grafana_prefix", dGrafanaPrefix)
}

func CGraylogPrefix() string {
	return cStr("graylog_prefix", dGraylogPrefix)
}

func CAlertManagerPrefix() string {
	return cStr("alert_manager_prefix", dAlertManagerPrefix)
}

func CFileServerPrefix() string {
	return cStr("file_server_prefix", dFileServerPrefix)
}

func CFileServiceAddr() string {
	return cStr("file_service_addr", dFileServiceAddr)
}

//一些额外的Get
func CNoneRowsAffect(tableName string) error {
	return errors.New(fmt.Sprintf("table:%s %s", tableName, dNoneRowsAffect))
}

func CNotFind(tableName string, condition interface{}) error {
	return errors.New(fmt.Sprintf("can not find item(table:%s  condition:%+v", tableName, condition))
}

func CVMEtcdHost() string {
	return cStr("VMEtcdHost", dVMEtcdHost)
}
func CVMEtcdPort() int {
	return cInt("VMEtcdPort", dVMEtcdPort)
}
func CVMJenkinsHost() string {
	return cStr("VMJenkinsHost", dVMJenkinsHost)
}
func CVMJenkinsPort() int {
	return cInt("VMJenkinsPort", dVMJenkinsPort)
}
func CVMJenkinsUser() string {
	return cStr("VMJenkinsUser", dVMJenkinsUser)
}
func CVMJenkinsPassword() string {
	return cStr("VMJenkinsPassword", dVMJenkinsPassword)
}
func CVMSonarHost() string {
	return cStr("VMSonarHost", dVMSonarHost)
}
func CVMSonarPort() int {
	return cInt("VMSonarPort", dVMSonarPort)
}
func CVMSonarUser() string {
	return cStr("VMSonarUser", dVMSonarUser)
}
func CVMSonarPassword() string {
	return cStr("VMSonarPassword", dVMSonarPassword)
}
func CVMSonarEMailHost() string {
	return cStr("VMSonarEMailHost", dVMSonarEMailHost)
}
func CVMSonarEMailPort() int {
	return cInt("VMSonarEMailPort", dVMSonarEMailPort)
}
func CVMSonarEMailUser() string {
	return cStr("VMSonarEMailUser", dVMSonarEMailUser)
}
func CVMSonarEMailPassword() string {
	return cStr("VMSonarEMailPassword", dVMSonarEMailPassword)
}
func CVMSonarEMailDefaultTitle() string {
	return cStr("VMSonarEMailDefaultTitle", dVMSonarEMailDefaultTitle)
}
func SonarAccessHost() string {
	return cStr("SonarAccessHost", dSonarAccessHost)
}
func SonarAccessPort() int {
	return cInt("SonarAccessPort", dSonarAccessPort)
}
func CSonarUser() string {
	return cStr("SonarUser", dSonarUser)
}
func CSonarPassword() string {
	return cStr("SonarPassword", dSonarPassword)
}

func WarURL() string {
	return cStr("war_url", dwarURL)
}

func CConsulHost() string {
	return cStr("ConsulHost", dConsulHost)
}

func CConsulPort() int {
	return cInt("ConsulPort", dConsulPort)
}

func ServiceControllerURL() string {
	return cStr("ClusterServiceControllerURL", dClusterServiceControllerURL)
}

func ExternalNetworkPort() int {
	return cInt("ClusterExternalNetworkPort", dClusterExternalNetworkPort)
}

func Harbor() string {
	return cStr("ClusterHarbor", dClusterHarbor)
}

//没用到
//func CAgentInstallCMD() string {
//	return cStr("agent_install_cmd", dAgentInstallCMD)
//}

func CLinuxAgentDownloadURL() string {
	return cStr("linux_agent_download_url", dLinuxAgentDownloadURL)
}

func CLinuxAgentInstallCMD() string {
	return cStr("linux_agent_install_cmd", dLinuxAgentInstallCMD)
}

func CWindowsAgentInstallCMD() string {
	return cStr("windows_agent_install_cmd", dWindowsAgentInstallCMD)
}

func CLinuxAgentLogPath() string {
	return cStr("linux_agent_log_path", dLinuxAgentLogPath)
}

func CWindowsAgentLogPath() string {
	return cStr("windows_agent_log_path", dWindowsAgentLogPath)
}

func CLinuxAgentPath() string {
	return cStr("linux_agent_path", dLinuxAgentPath)
}

func CWindowsAgentPath() string {
	return cStr("windows_agent_path", dWindowsAgentPath)
}

//dWindowsAgentDownloadURL
func CWindowsAgentDownloadURL() string {
	return cStr("windows_agent_download_url", dWindowsAgentDownloadURL)
}

func CWindowsAgentDownloadURLShow() string {
	return cStr("windows_agent_download_url_show", dWindowsAgentDownloadURLShow)
}

func CRegisterTtL() time.Duration {
	return time.Duration(cInt("register_ttL", dRegisterTtL)) * time.Second
}

func CRegisterInterval() time.Duration {
	return time.Duration(cInt("register_interval", dRegisterInterval)) * time.Second
}

func CKibanaAddr() string {
	return cStr("kibana_addr", KIBANA_ADDR)
}

func CKibanaAddrExternal() string {
	return cStr("kibana_addr_external", KIBANA_ADDR)
}

func SuperPassword() string {
	return cStr("super_password", dSuperPassword)
}
