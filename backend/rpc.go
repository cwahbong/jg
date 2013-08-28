package backend

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
)

const (
	dbName = "jg_test"
)

type Jg struct{}

type PrimaryTagsArgs struct{}

type PrimaryTagsReply struct {
	Tags []Tag `json:"tags"`
}

func (*Jg) PrimaryTags(request *http.Request, args *PrimaryTagsArgs, reply *PrimaryTagsReply) error {
	session := defaultDial()
	defer session.Close()
	c := session.DB(dbName).C("tag")
	c.Find(nil).All(&reply.Tags)
	return nil
}

type TagByIdArgs struct {
	Id string `json:"id"`
}

type TagByIdReply struct {
	Tags     []Tag     `json:"tags"`
	Problems []Problem `json:"problems"`
}

func (*Jg) TagById(request *http.Request, args *TagByIdArgs, reply *TagByIdReply) error {
	// TODO query database and return result
	reply.Tags = []Tag{{"blablaid", "blablaname"}}
	return nil
}

type ProblemByIdArgs struct {
}

type ProblemByIdReply struct {
}

func (*Jg) ProblemById(request *http.Request, args *ProblemByIdArgs, replay ProblemByIdReply) error {
	return nil
}

// http.HandleFunc("/j/submit", SubmitHandler)
// http.HandleFunc() submission result

func RpcServer() *rpc.Server {
	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(json.NewCodec(), "application/json")
	rpcServer.RegisterService(new(Jg), "")
	return rpcServer
}
