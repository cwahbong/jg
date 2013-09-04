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
	Tags []Named `json:"tags"`
}

func (*Jg) PrimaryTags(request *http.Request, args *PrimaryTagsArgs, reply *PrimaryTagsReply) error {
	session := DefaultDial()
	defer session.Close()
	c := session.DB(dbName).C("tag")
	c.Find(bson.M{"primary": true}).All(&reply.Tags)
	return nil
}

type TagByIdArgs struct {
	Id string `json:"id"`
}

type TagByIdReply struct {
	Name     string  `json:"name"`
	Tags     []Named `json:"tags"`
	Problems []Named `json:"problems"`
	Tagged   []Named `json:"tagged"`
}

func (*Jg) TagById(request *http.Request, args *TagByIdArgs, reply *TagByIdReply) error {
	session := DefaultDial()
	defer session.Close()
	ct := session.DB(dbName).C("tag")

	var tag Tag
	ct.Find(bson.M{"id": args.Id}).One(&tag)
	reply.Name = tag.Name
	reply.Tagged = make([]Named, len(tag.Tagged))
	for idx, objId := range tag.Tagged {
		ct.Find(bson.M{"_id": objId}).One(&reply.Tagged[idx])
	}

	ct.Find(bson.M{"tagged": tag.ObjId}).All(&reply.Tags)
	cp := session.DB(dbName).C("problem")
	cp.Find(bson.M{"tagged": tag.ObjId}).All(&reply.Problems)
	return nil
}

type ProblemByIdArgs struct {
	Id string `json:"id"`
}

type ProblemByIdReply struct {
	Name   string  `json:"name"`
	Brief  string  `json:"brief"`
	Tagged []Named `json:"tagged"`
	// TODO supported method
	// TODO downloadable resource
}

func (*Jg) ProblemById(request *http.Request, args *ProblemByIdArgs, reply *ProblemByIdReply) error {
	session := DefaultDial()
	defer session.Close()
	c := session.DB(dbName).C("problem")

	var problem Problem
	c.Find(bson.M{"id": args.Id}).One(&problem)

	reply.Name = problem.Name
	reply.Brief = problem.Brief
	reply.Tagged = make([]Named, len(problem.Tagged))

	ct := session.DB(dbName).C("tag")
	for idx, objId := range problem.Tagged {
		ct.Find(bson.M{"_id": objId}).One(&reply.Tagged[idx])
	}
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
