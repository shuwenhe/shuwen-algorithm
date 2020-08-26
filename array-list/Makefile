# 配置
service_name = antdate-service
repository = registry.cn-beijing.aliyuncs.com/antdate/antdate-service
docker_file = Dockerfile
config_path = configs/dev.yaml
tag = dev

.PHONY:help
help: ##@other Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

.PHONY:run
run: ##@run 启动服务.
	@echo "启动服务..."
	go run -race . --config=$(config_path)


.PHONY:build-linux
build-linux: ##@build 构建二进制文件.
	@echo "构建二进制文件"
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(service_name) .
	upx $(service_name)

# 镜像构建
.PHONY:docker-build
docker-build:build-linux ##@docker 构建镜像.
	docker build -t $(repository):$(tag) -f $(docker_file) .
	@echo "build success"


# 镜像构建
.PHONY:docker-push
docker-push: docker-build##@docker 推送镜像.
	docker push $(repository):$(tag)
	@echo "push success"

# 部署
.PHONY:deploy
deploy:docker-push  ##@docker 推送镜像.
	curl -X POST -v http://portainer.antdate.cn/api/webhooks/3008ef6e-d826-46c8-82ff-9e19c80cb22c
	@echo "push success"

HELP_FUN = \
	%help; \
	while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
	print "usage: make [target]\n\n"; \
	for (sort keys %help) { \
		print "${WHITE}$$_:${RESET}\n"; \
		for (@{$$help{$$_}}) { \
			$$sep = " " x (32 - length $$_->[0]); \
			print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
	}; \
	print "\n"; }
