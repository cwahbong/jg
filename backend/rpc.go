package backend

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"labix.org/v2/mgo/bson"
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
	c.Find(bson.M{"primary": true}).All(&reply.Tags)
	return nil
}

type TagByIdArgs struct {
	Id string `json:"id"`
}

type TagByIdReply struct {
	Tag      Tag       `json:"tag"`
	Tags     []Tag     `json:"tags"`
	Problems []Problem `json:"problems"`
}

func (*Jg) TagById(request *http.Request, args *TagByIdArgs, reply *TagByIdReply) error {
	session := defaultDial()
	defer session.Close()
	ct := session.DB(dbName).C("tag")
	ct.Find(bson.M{"id": args.Id}).One(&reply.Tag)
	ct.Find(bson.M{"tagged": args.Id}).All(&reply.Tags)
	cp := session.DB(dbName).C("problem")
	cp.Find(bson.M{"tagged": args.Id}).All(&reply.Problems)
	return nil
}

type ProblemByIdArgs struct {
	Id string `json:"id"`
}

type ProblemByIdReply struct {
	// TODO
}

func (*Jg) ProblemById(request *http.Request, args *ProblemByIdArgs, reply ProblemByIdReply) error {
	session := defaultDial()
	defer session.Close()
	// c := session.DB(dbName).C("problem")
	// c.Find(bson.M{"id"}).One(&reply.)
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
