package main

import (
	"net"

	"google.golang.org/grpc"

	"github.com/zhiruchen/archi/conf"
	"github.com/zhiruchen/archi/infrastructure"
	pb "github.com/zhiruchen/archi/pb"
)

func main() {
	configPath := flag.String("config", "conf/config.toml", "config file's path")
	flag.Parse()

	conf.LoadConfig(*configPath)
	infrastructure.InitMysql(conf.Config.MysqlConf)

	listen, err := net.Listen("tcp", conf.Config.ListenAddr)
	if err != nil {
		panic(err)
	}

	// server := &infrastructure.RPCHandler{QuestionInteractor: }

	s := grpc.NewServer()
	pb.RegisterArchiServer(s)

	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
