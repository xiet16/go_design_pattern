package proxy_pattern

func ProxyTest() {
	vehicle := &TransportByVehicle{
		Name: "自行车出行方式",
	}

	proxy := &TransportProxy{
		Name:             "第三方代理出行方式",
		TransportHandler: vehicle,
	}

	proxy.GoOut()
}
