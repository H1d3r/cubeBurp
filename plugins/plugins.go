package Plugins

import (
	"cube/model"
	"cube/plugins/crack"
	"cube/plugins/probe"
	"cube/plugins/sqlcmd"
	"cube/util"
)

type ProbeFunc func(task model.ProbeTask) (taskResult model.ProbeTaskResult)
type SqlcmdFunc func(task model.SqlcmdTask) (taskResult model.SqlcmdTaskResult)
type CrackFunc func(task model.CrackTask) (taskResult model.CrackTaskResult)

//type CrackFunc func(task model.Task) (taskResult model.TaskResult)

var (
	ProbeFuncMap  map[string]ProbeFunc
	SqlcmdFuncMap map[string]SqlcmdFunc
	CrackFuncMap  map[string]CrackFunc

	ProbeKeys        []string
	SqlcmdKeys       []string
	CrackKeys        []string
	CrackFuncExclude []string
	//CrackFuncMap map[string]CrackFunc
)

func init() {
	ProbeFuncMap = make(map[string]ProbeFunc)
	ProbeFuncMap["oxid"] = probe.OxidProbe
	//ProbeFuncMap["smb"] = probe.SmbProbe
	ProbeFuncMap["ms17010"] = probe.Ms17010Probe
	ProbeFuncMap["zookeeper"] = probe.ZookeeperProbe
	ProbeKeys = append(ProbeKeys, "ALL")
	for k := range ProbeFuncMap {
		ProbeKeys = append(ProbeKeys, k)
	}

	SqlcmdFuncMap = make(map[string]SqlcmdFunc)
	SqlcmdFuncMap["ssh"] = sqlcmd.SshCmd
	for k := range SqlcmdFuncMap {
		SqlcmdKeys = append(SqlcmdKeys, k)
	}

	CrackFuncMap = make(map[string]CrackFunc)
	CrackFuncMap["ssh"] = crack.SshCrack
	CrackFuncMap["mysql"] = crack.MysqlCrack
	CrackFuncMap["redis"] = crack.RedisCrack
	CrackFuncMap["ftp"] = crack.FtpCrack
	CrackFuncMap["smb"] = crack.SmbCrack
	CrackFuncMap["mongo"] = crack.MongoCrack
	CrackFuncMap["elastic"] = crack.ElasticCrack
	CrackFuncMap["postgres"] = crack.PostgresCrack
	CrackFuncMap["mssql"] = crack.MssqlCrack
	CrackFuncMap["phpmyadmin"] = crack.PhpmyadminCrack
	CrackFuncMap["httpbasic"] = crack.HttpBasicCrack

	CrackFuncExclude = []string{"phpmyadmin", "httpbasic"} //去除phpmyadmin这类单独使用的

	for k := range CrackFuncMap {
		if !util.SliceContain(k, CrackFuncExclude) {
			CrackKeys = append(CrackKeys, k)
		}
	}

	CrackPluginKeys := make(map[string][]string)
	CrackPluginKeys["ALL"] = CrackKeys

}
