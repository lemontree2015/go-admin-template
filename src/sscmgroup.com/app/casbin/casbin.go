package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"sscmgroup.com/app/config"
	"sscmgroup.com/app/logger"
	"time"
)

var Enforcer *casbin.SyncedEnforcer
var chCasbinPolicy chan *chCasbinPolicyItem

type chCasbinPolicyItem struct {
	e *casbin.SyncedEnforcer
}

func InitCasbin(adapter persist.Adapter) (fn func(), err error) { //*casbin.SyncedEnforcer,
	cfg := config.Conf.Casbin
	logger.Logger.Debug("InitCasbin casbin.......")
	if cfg.Model == "" {
		return nil, nil
	}

	Enforcer, err = casbin.NewSyncedEnforcer(cfg.Model)
	if err != nil {
		logger.Logger.Error("InitCasbin casbin step1 error.......", err.Error())
		return nil, err
	}
	Enforcer.EnableLog(cfg.Debug)

	err = Enforcer.InitWithModelAndAdapter(Enforcer.GetModel(), adapter)
	if err != nil {
		logger.Logger.Error("InitCasbin casbin step2 error.......", err.Error())
		return nil, err
	}
	Enforcer.EnableEnforce(cfg.Enable)

	cleanFunc := func() {}
	if cfg.AutoLoad {
		Enforcer.StartAutoLoadPolicy(time.Duration(cfg.AutoLoadInternal) * time.Second)
		cleanFunc = func() {
			Enforcer.StopAutoLoadPolicy()
		}
	}

	chCasbinPolicy = make(chan *chCasbinPolicyItem, 1)
	go func() {
		logger.Logger.Debug("load casbin policy manu")
		for item := range chCasbinPolicy {
			err := item.e.LoadPolicy()
			if err != nil {
				logger.Logger.Error("load casbin policy manu", err.Error())
			}
		}
	}()
	return cleanFunc, nil
}

// LoadCasbinPolicy 异步加载casbin权限策略
func LoadCasbinPolicy(e *casbin.SyncedEnforcer) {
	cfg := config.Conf.Casbin
	if !cfg.Enable {
		return
	}

	if e == nil {
		return
	}

	if len(chCasbinPolicy) > 0 {
		logger.Logger.Info("The load casbin policy is already in the wait queue")
		return
	}

	chCasbinPolicy <- &chCasbinPolicyItem{
		e: e,
	}
}
