package goway_couchbase_store

import "github.com/andrepinto/goway/product"

type LocalProductRepository struct {

}

func NewLocalRepository() *LocalProductRepository{
	return &LocalProductRepository{}
}

func(l *LocalProductRepository) GetAllProducts() []product.Product_v1 {

	return []product.Product_v1{

		product.Product_v1{
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
		product.Product_v1{
			Code: "cockpit",
			Name: "cockpit Api",
			Version: "1",
			Routes: []product.Routes_v1{
				product.Routes_v1{
					ListenPath: "/auth/byemail",
					Verb: "GET",
					ServiceName: "authentication",
					Handlers: []string{"AUTHENTICATION", "METRICS"},
					Code:"auth_by_email",
					Roles: []string{},
				},
				product.Routes_v1{
					ListenPath: "/auth/renew",
					Verb: "GET",
					ServiceName: "authentication",
					Handlers:  []string{"AUTHENTICATION", "METRICS"},
					Code:"renew_token",
					Roles: []string{},
				},
			},
		},

	}
}

func(l *LocalProductRepository) GetAllClients() []product.Client_v1 {
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

func(l *LocalProductRepository) CreateProduct() ( bool, product.Product_v1 ) {
	return false, nil
}

func(l *LocalProductRepository) CreateClient() ( bool, product.Client ) {
	return false, nil
}