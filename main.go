package main

import (
	"flag"
	"net"

	"google.golang.org/grpc"

	"github.com/zhiruchen/archi/conf"
	"github.com/zhiruchen/archi/infrastructure"
	"github.com/zhiruchen/archi/interfaces"
	pb "github.com/zhiruchen/archi/pb"
	"github.com/zhiruchen/archi/usecases"
)

func main() {
	configPath := flag.String("config", "conf/config.toml", "config file's path")
	flag.Parse()

	conf.LoadConfig(*configPath)
	infrastructure.InitMysql(conf.Conf.MysqlConf)

	listen, err := net.Listen("tcp", conf.Conf.ListenAddr)
	if err != nil {
		panic(err)
	}

	server := &interfaces.RPCHandler{
		QuestionInteractor: &usecases.QuestionInteractor{
			QuestionStore: interfaces.NewDBQuestion(infrastructure.Db),
		},
	}

	s := grpc.NewServer()
	pb.RegisterArchiServer(s, server)
	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
