package goway_couchbase_store

import (
	"github.com/andrepinto/goway/product"
	"github.com/couchbase/gocb"
	"fmt"
)

type CouchbaseRepository struct {
	Cluster 	*gocb.Cluster
	Bucket 		*gocb.Bucket
}

type CouchbaseRepositoryOptions struct {
	ConnectionString 	string
	BucketName 		string
	BucketPassword 		string
}



func NewCouchbaseRepository(options *CouchbaseRepositoryOptions) *CouchbaseRepository{

	cluster, err := gocb.Connect(options.ConnectionString)

	if(err!=nil){
		panic("ERROR_ON_OPEN_CLUSTER")
	}

	bucket, err := cluster.OpenBucket(options.BucketName, "")

	if(err!=nil){
		panic("ERROR_ON_OPEN_BUCKET")
	}

	manager := bucket.Manager("","")

	manager.CreatePrimaryIndex("gateway",true, false)

	manager.CreateIndex("gateway_products",[]string{"_type"},true, false)


	return &CouchbaseRepository{
		Cluster: cluster,
		Bucket: bucket,
	}
}


func(l *CouchbaseRepository) CreateAndGet(){


	l.Bucket.Upsert("product:4",
		product.Product_v1{
			Type: "product",
			Id:"4",
			Code: "customer",
			Name: "Customer Api",
			Version: "1",
			Routes: []product.Routes_v1{
				product.Routes_v1{
					ListenPath: "/api/facets",
					Verb: "GET",
					ServiceName: "authentication",
					Handlers: []string{"AUTHENTICATION", "METRICS"},
					Code:"auth_by_email",
					Roles: []string{},
					InjectData: []product.InjectData_v1{
						product.InjectData_v1{
							Where:"params",
							Code: "paramId",
							Value: "123456789",
						},
						product.InjectData_v1{
							Where:"header",
							Code: "headerId",
							Value: "9999999999",
						},
					},
					InjectGlobalData:true,
				},
				product.Routes_v1{
					ListenPath: "/api/facets",
					Verb: "POST",
					ServiceName: "authentication",
					Handlers: []string{"AUTHENTICATION", "METRICS"},
					Code:"auth_by_email",
					Roles: []string{},
					InjectData: []product.InjectData_v1{
						product.InjectData_v1{
							Where:"params",
							Code: "paramId",
							Value: "123456789",
						},
						product.InjectData_v1{
							Where:"header",
							Code: "headerId",
							Value: "9999999999",
						},
					},
					InjectGlobalData:true,
				},
			},
		},
		0)

	var inUser product.Product_v1
	l.Bucket.Get("product:4", &inUser)
	fmt.Printf("User: %v\n", inUser)
}


func(l *CouchbaseRepository) GetAllProducts() []product.Product_v1 {

	query := gocb.NewN1qlQuery("SELECT * FROM gateway AS product WHERE _type = $1")
	rows, _ := l.Bucket.ExecuteN1qlQuery(query, []interface{}{"product"})

	var row interface{}
	var result []product.Product_v1

	rows.One(&row)
	fmt.Printf("Query1: %+v\n", row)

	//
	//for rows.Next(&row) {
	//	fmt.Println("%v", row)
	//
	//}
	rows.Close()

	return result
}

func(l *CouchbaseRepository) GetAllClients() []product.Client_v1 {
	return []product.Client_v1{
		product.Client_v1{
			ApiPath:"12124578",
			Product:"customer",
			Client:"myorg",
			RemoveApiPath: true,
			Version: "1",
			GlobalInjectData:[]product.InjectData_v1{

			},
			Routes: []product.Routes_v1{
				product.Routes_v1{
					ListenPath: "/auth/byemail",
					Verb: "GET",
					ServiceName: "authentication-custom",
					Handlers: []string{"AUTHENTICATION", "METRICS"},
					Code:"auth_by_email",
					Roles: []string{},
				},
			},
		},
		product.Client_v1{
			ApiPath:"121245782",
			Product:"customer",
			Client:"myorg2",
			RemoveApiPath: true,
			Version: "1",
			GlobalInjectData:[]product.InjectData_v1{
				product.InjectData_v1{
					Where:"url",
					Code: "orgs",
					Value: "tlantic",
				},
				product.InjectData_v1{
					Where:"url",
					Code: "apps",
					Value: "customer",
				},
			},
			Routes: []product.Routes_v1{
				product.Routes_v1{
					ListenPath: "/auth/byemail",
					Verb: "GET",
					ServiceName: "authentication-custom",
					Handlers: []string{"AUTHENTICATION", "METRICS"},
					Code:"auth_by_email",
					Roles: []string{},
				},
			},
		},
	}
}
