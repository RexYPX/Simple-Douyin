package dal

import (
	"Simple-Douyin/cmd/message/dal/db"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Init() {
	db.Init()
	klog.Debug("dal::init::init()", db.DB)
}
